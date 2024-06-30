package host_service

import (
	"context"
	"fmt"
	servicepb "github.com/AfoninaOlga/hostname-configurator/gen"
	"log"
	"net"
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
