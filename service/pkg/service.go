package host_service

import (
	"bufio"
	"context"
	"fmt"
	servicepb "github.com/AfoninaOlga/hostname-configurator/gen"
	"log"
	"net"
	"os"
	"regexp"
	"strings"
	"sync"
)

type Server struct {
	servicepb.UnimplementedConfiguratorServer
	hostnamePath string
	resolvePath  string
	hostMtx      sync.Mutex
	dnsMtx       sync.RWMutex
	servers      []string
	linesToWrite []string
}

func NewServer(hostNamepath string, resolvePath string) *Server {
	file, err := os.OpenFile(resolvePath, os.O_RDONLY, 0777)
	if err != nil {
		return nil
	}
	defer file.Close()

	// Initialize dns servers list and get data to write to /etc/resolv.conf
	servers := make([]string, 0, 50)
	linesToWrite := make([]string, 0, 50)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fileLine := scanner.Text()
		line := strings.Split(fileLine, " ")
		if len(line) == 2 && line[0] == "nameserver" && isValidIpAddress(line[1]) {
			servers = append(servers, line[1])
		} else {
			linesToWrite = append(linesToWrite, fileLine)
		}
	}

	return &Server{hostnamePath: hostNamepath, resolvePath: resolvePath, servers: servers, linesToWrite: linesToWrite}
}

func (s *Server) SetHostname(ctx context.Context, in *servicepb.HostnameRequest) (*servicepb.HostnameReply, error) {
	if !isValidHostname(in.Hostname) {
		log.Println("Got invalid hostname to set")
		return nil, fmt.Errorf("invalid hostname")
	}
	s.hostMtx.Lock()
	defer s.hostMtx.Unlock()
	err := os.WriteFile(s.hostnamePath, []byte(in.Hostname+"\n"), 0644)
	if err != nil {
		return nil, err
	}

	return &servicepb.HostnameReply{Hostname: in.Hostname}, nil
}

func (s *Server) ListDnsServers(ctx context.Context, in *servicepb.Empty) (*servicepb.DnsListReply, error) {
	s.dnsMtx.RLock()
	defer s.dnsMtx.RUnlock()
	return &servicepb.DnsListReply{Servers: s.servers}, nil
}

func (s *Server) AddDnsServer(ctx context.Context, in *servicepb.AddDnsRequest) (*servicepb.Empty, error) {
	if !isValidIpAddress(in.Server) {
		log.Println("Got invalid address to add")
		return nil, fmt.Errorf("invalid DNS server")
	}
	s.dnsMtx.Lock()
	defer s.dnsMtx.Unlock()

	if s.has(in.Server) {
		return &servicepb.Empty{}, nil
	}
	s.servers = append(s.servers, in.Server)
	err := s.writeFile()
	if err != nil {
		return nil, err
	}
	return &servicepb.Empty{}, nil
}

func (s *Server) DeleteDnsServer(ctx context.Context, in *servicepb.DeleteDnsRequest) (*servicepb.Empty, error) {
	if !isValidIpAddress(in.Server) {
		log.Println("Got invalid address to delete")
		return nil, fmt.Errorf("invalid DNS server")
	}
	s.dnsMtx.Lock()
	defer s.dnsMtx.Unlock()

	if s.delete(in.Server) {
		err := s.writeFile()
		if err != nil {
			// returning address to list, because it wasn't actually deleted
			s.servers = append(s.servers, in.Server)
			return nil, err
		}
	}
	return &servicepb.Empty{}, nil
}

func isValidIpAddress(ip string) bool {
	return net.ParseIP(ip) != nil
}

func isValidHostname(hostname string) bool {
	hostRegex := regexp.MustCompile(`^(([a-zA-Z0-9]|[a-zA-Z0-9][a-zA-Z0-9-]*[a-zA-Z0-9])\.)*([A-Za-z0-9]|[A-Za-z0-9][A-Za-z0-9-]*[A-Za-z0-9])$`)
	if hostRegex.MatchString(hostname) {
		if len(hostname) > 255 {
			return false
		}
		for _, s := range strings.Split(hostname, ".") {
			if len(s) > 63 {
				return false
			}
		}
		return true
	}
	return false
}

func (s *Server) has(server string) bool {
	for _, srv := range s.servers {
		if srv == server {
			return true
		}
	}
	return false
}

func (s *Server) delete(server string) bool {
	n := len(s.servers)
	if n == 0 {
		return false
	}
	for i, srv := range s.servers {
		if server == srv {
			s.servers[i] = s.servers[n-1]
			s.servers = s.servers[:n-1]
			return true
		}
	}
	return false
}

func (s *Server) writeFile() error {
	content := strings.Join(s.linesToWrite, "\n") + "\n"
	if len(s.servers) > 0 {
		content += "nameserver " + strings.Join(s.servers, "\nnameserver ") + "\n"
	}
	err := os.WriteFile(s.resolvePath, []byte(content+"\n"), 0644)
	return err
}
