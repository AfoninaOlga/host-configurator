// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v3.12.4
// source: service.proto

package servicepb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	Configurator_SetHostname_FullMethodName = "/service.Configurator/SetHostname"
)

// ConfiguratorClient is the client API for Configurator service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ConfiguratorClient interface {
	SetHostname(ctx context.Context, in *HostnameRequest, opts ...grpc.CallOption) (*HostnameReply, error)
}

type configuratorClient struct {
	cc grpc.ClientConnInterface
}

func NewConfiguratorClient(cc grpc.ClientConnInterface) ConfiguratorClient {
	return &configuratorClient{cc}
}

func (c *configuratorClient) SetHostname(ctx context.Context, in *HostnameRequest, opts ...grpc.CallOption) (*HostnameReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(HostnameReply)
	err := c.cc.Invoke(ctx, Configurator_SetHostname_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConfiguratorServer is the server API for Configurator service.
// All implementations must embed UnimplementedConfiguratorServer
// for forward compatibility
type ConfiguratorServer interface {
	SetHostname(context.Context, *HostnameRequest) (*HostnameReply, error)
	mustEmbedUnimplementedConfiguratorServer()
}

// UnimplementedConfiguratorServer must be embedded to have forward compatible implementations.
type UnimplementedConfiguratorServer struct {
}

func (UnimplementedConfiguratorServer) SetHostname(context.Context, *HostnameRequest) (*HostnameReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetHostname not implemented")
}
func (UnimplementedConfiguratorServer) mustEmbedUnimplementedConfiguratorServer() {}

// UnsafeConfiguratorServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ConfiguratorServer will
// result in compilation errors.
type UnsafeConfiguratorServer interface {
	mustEmbedUnimplementedConfiguratorServer()
}

func RegisterConfiguratorServer(s grpc.ServiceRegistrar, srv ConfiguratorServer) {
	s.RegisterService(&Configurator_ServiceDesc, srv)
}

func _Configurator_SetHostname_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HostnameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfiguratorServer).SetHostname(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Configurator_SetHostname_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfiguratorServer).SetHostname(ctx, req.(*HostnameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Configurator_ServiceDesc is the grpc.ServiceDesc for Configurator service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Configurator_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "service.Configurator",
	HandlerType: (*ConfiguratorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetHostname",
			Handler:    _Configurator_SetHostname_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}
