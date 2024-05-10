// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.26.1
// source: nmap.proto

package netvuln

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// NetVulnServiceClient is the client API for NetVulnService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NetVulnServiceClient interface {
	CheckVuln(ctx context.Context, in *CheckVulnRequest, opts ...grpc.CallOption) (*CheckVulnResponse, error)
}

type netVulnServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewNetVulnServiceClient(cc grpc.ClientConnInterface) NetVulnServiceClient {
	return &netVulnServiceClient{cc}
}

func (c *netVulnServiceClient) CheckVuln(ctx context.Context, in *CheckVulnRequest, opts ...grpc.CallOption) (*CheckVulnResponse, error) {
	out := new(CheckVulnResponse)
	err := c.cc.Invoke(ctx, "/netvuln.NetVulnService/CheckVuln", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NetVulnServiceServer is the server API for NetVulnService service.
// All implementations must embed UnimplementedNetVulnServiceServer
// for forward compatibility
type NetVulnServiceServer interface {
	CheckVuln(context.Context, *CheckVulnRequest) (*CheckVulnResponse, error)
	mustEmbedUnimplementedNetVulnServiceServer()
}

// UnimplementedNetVulnServiceServer must be embedded to have forward compatible implementations.
type UnimplementedNetVulnServiceServer struct {
}

func (UnimplementedNetVulnServiceServer) CheckVuln(context.Context, *CheckVulnRequest) (*CheckVulnResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckVuln not implemented")
}
func (UnimplementedNetVulnServiceServer) mustEmbedUnimplementedNetVulnServiceServer() {}

// UnsafeNetVulnServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NetVulnServiceServer will
// result in compilation errors.
type UnsafeNetVulnServiceServer interface {
	mustEmbedUnimplementedNetVulnServiceServer()
}

func RegisterNetVulnServiceServer(s grpc.ServiceRegistrar, srv NetVulnServiceServer) {
	s.RegisterService(&NetVulnService_ServiceDesc, srv)
}

func _NetVulnService_CheckVuln_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckVulnRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetVulnServiceServer).CheckVuln(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/netvuln.NetVulnService/CheckVuln",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetVulnServiceServer).CheckVuln(ctx, req.(*CheckVulnRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// NetVulnService_ServiceDesc is the grpc.ServiceDesc for NetVulnService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NetVulnService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "netvuln.NetVulnService",
	HandlerType: (*NetVulnServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CheckVuln",
			Handler:    _NetVulnService_CheckVuln_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "nmap.proto",
}
