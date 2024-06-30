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
	"syscall"
)

type Server struct {
	servicepb.UnimplementedConfiguratorServer
	hostnamePath string
	resolvePath  string
	hostMtx      sync.Mutex
	dnsMtx       sync.RWMutex
}

func NewServer(hostNamepath string, resolvePath string) *Server {
	return &Server{hostnamePath: hostNamepath, resolvePath: resolvePath}
}

func (s *Server) SetHostname(ctx context.Context, in *servicepb.HostnameRequest) (*servicepb.HostnameReply, error) {
	if !isValidHostname(in.Hostname) {
		log.Println("Got invalid hostname to set")
		return nil, fmt.Errorf("invalid hostname")
	}
	s.hostMtx.Lock()
	defer s.hostMtx.Unlock()
	err := syscall.Sethostname([]byte(in.Hostname))
	if err != nil {
		return nil, err
	}
	log.Println("SetHostname called")
	return &servicepb.HostnameReply{Hostname: in.Hostname}, nil
}

func (s *Server) AddDnsServer(address string) error {
	if !isValidIpAddress(address) {
		return fmt.Errorf("invalid dns address")
	}
	return nil
}

func (s *Server) ListDnsServers(ctx context.Context, in *servicepb.Empty) (*servicepb.DnsListReply, error) {
	s.dnsMtx.RLock()
	s.dnsMtx.RUnlock()
	file, err := os.OpenFile(s.resolvePath, os.O_RDONLY, 0777)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	servers := make([]string, 0, 20)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		if len(line) == 2 && line[0] == "nameserver" && isValidIpAddress(line[1]) {
			servers = append(servers, line[1])
		}
	}
	return &servicepb.DnsListReply{Servers: servers}, nil
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
