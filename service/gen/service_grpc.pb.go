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
	Configurator_SetHostname_FullMethodName     = "/service.Configurator/SetHostname"
	Configurator_GetHostname_FullMethodName     = "/service.Configurator/GetHostname"
	Configurator_ListDnsServers_FullMethodName  = "/service.Configurator/ListDnsServers"
	Configurator_AddDnsServer_FullMethodName    = "/service.Configurator/AddDnsServer"
	Configurator_DeleteDnsServer_FullMethodName = "/service.Configurator/DeleteDnsServer"
)

// ConfiguratorClient is the client API for Configurator service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ConfiguratorClient interface {
	// Sets hostname
	SetHostname(ctx context.Context, in *HostnameRequest, opts ...grpc.CallOption) (*HostnameReply, error)
	// Gets hostname
	GetHostname(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*HostnameReply, error)
	// Returns list of DNS servers
	ListDnsServers(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*DnsListReply, error)
	// Adds DNS server to list
	AddDnsServer(ctx context.Context, in *AddDnsRequest, opts ...grpc.CallOption) (*Empty, error)
	// Deletes DNS server from list
	DeleteDnsServer(ctx context.Context, in *DeleteDnsRequest, opts ...grpc.CallOption) (*Empty, error)
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

func (c *configuratorClient) GetHostname(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*HostnameReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(HostnameReply)
	err := c.cc.Invoke(ctx, Configurator_GetHostname_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configuratorClient) ListDnsServers(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*DnsListReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DnsListReply)
	err := c.cc.Invoke(ctx, Configurator_ListDnsServers_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configuratorClient) AddDnsServer(ctx context.Context, in *AddDnsRequest, opts ...grpc.CallOption) (*Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Empty)
	err := c.cc.Invoke(ctx, Configurator_AddDnsServer_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configuratorClient) DeleteDnsServer(ctx context.Context, in *DeleteDnsRequest, opts ...grpc.CallOption) (*Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Empty)
	err := c.cc.Invoke(ctx, Configurator_DeleteDnsServer_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConfiguratorServer is the server API for Configurator service.
// All implementations must embed UnimplementedConfiguratorServer
// for forward compatibility
type ConfiguratorServer interface {
	// Sets hostname
	SetHostname(context.Context, *HostnameRequest) (*HostnameReply, error)
	// Gets hostname
	GetHostname(context.Context, *Empty) (*HostnameReply, error)
	// Returns list of DNS servers
	ListDnsServers(context.Context, *Empty) (*DnsListReply, error)
	// Adds DNS server to list
	AddDnsServer(context.Context, *AddDnsRequest) (*Empty, error)
	// Deletes DNS server from list
	DeleteDnsServer(context.Context, *DeleteDnsRequest) (*Empty, error)
	mustEmbedUnimplementedConfiguratorServer()
}

// UnimplementedConfiguratorServer must be embedded to have forward compatible implementations.
type UnimplementedConfiguratorServer struct {
}

func (UnimplementedConfiguratorServer) SetHostname(context.Context, *HostnameRequest) (*HostnameReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetHostname not implemented")
}
func (UnimplementedConfiguratorServer) GetHostname(context.Context, *Empty) (*HostnameReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHostname not implemented")
}
func (UnimplementedConfiguratorServer) ListDnsServers(context.Context, *Empty) (*DnsListReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListDnsServers not implemented")
}
func (UnimplementedConfiguratorServer) AddDnsServer(context.Context, *AddDnsRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddDnsServer not implemented")
}
func (UnimplementedConfiguratorServer) DeleteDnsServer(context.Context, *DeleteDnsRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteDnsServer not implemented")
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

func _Configurator_GetHostname_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfiguratorServer).GetHostname(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Configurator_GetHostname_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfiguratorServer).GetHostname(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Configurator_ListDnsServers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfiguratorServer).ListDnsServers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Configurator_ListDnsServers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfiguratorServer).ListDnsServers(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Configurator_AddDnsServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddDnsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfiguratorServer).AddDnsServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Configurator_AddDnsServer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfiguratorServer).AddDnsServer(ctx, req.(*AddDnsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Configurator_DeleteDnsServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteDnsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfiguratorServer).DeleteDnsServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Configurator_DeleteDnsServer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfiguratorServer).DeleteDnsServer(ctx, req.(*DeleteDnsRequest))
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
		{
			MethodName: "GetHostname",
			Handler:    _Configurator_GetHostname_Handler,
		},
		{
			MethodName: "ListDnsServers",
			Handler:    _Configurator_ListDnsServers_Handler,
		},
		{
			MethodName: "AddDnsServer",
			Handler:    _Configurator_AddDnsServer_Handler,
		},
		{
			MethodName: "DeleteDnsServer",
			Handler:    _Configurator_DeleteDnsServer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}
