package host_service

import (
	"context"
	servicepb "github.com/AfoninaOlga/hostname-configurator/proto"
	"log"
	"syscall"
)

type Server struct {
	servicepb.UnimplementedConfiguratorServer
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) SetHostname(ctx context.Context, in *servicepb.HostnameRequest) (*servicepb.HostnameReply, error) {
	err := syscall.Sethostname([]byte(in.Hostname))
	if err != nil {
		return nil, err
	}
	log.Println("SetHostname called")
	return &servicepb.HostnameReply{Hostname: in.Hostname}, nil
}
