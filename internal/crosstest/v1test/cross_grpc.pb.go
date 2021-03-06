// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package crosspb

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

// CrossServiceClient is the client API for CrossService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CrossServiceClient interface {
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error)
	Fail(ctx context.Context, in *FailRequest, opts ...grpc.CallOption) (*FailResponse, error)
}

type crossServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCrossServiceClient(cc grpc.ClientConnInterface) CrossServiceClient {
	return &crossServiceClient{cc}
}

func (c *crossServiceClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error) {
	out := new(PingResponse)
	err := c.cc.Invoke(ctx, "/internal.crosstest.v1test.CrossService/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *crossServiceClient) Fail(ctx context.Context, in *FailRequest, opts ...grpc.CallOption) (*FailResponse, error) {
	out := new(FailResponse)
	err := c.cc.Invoke(ctx, "/internal.crosstest.v1test.CrossService/Fail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CrossServiceServer is the server API for CrossService service.
// All implementations must embed UnimplementedCrossServiceServer
// for forward compatibility
type CrossServiceServer interface {
	Ping(context.Context, *PingRequest) (*PingResponse, error)
	Fail(context.Context, *FailRequest) (*FailResponse, error)
	mustEmbedUnimplementedCrossServiceServer()
}

// UnimplementedCrossServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCrossServiceServer struct {
}

func (UnimplementedCrossServiceServer) Ping(context.Context, *PingRequest) (*PingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedCrossServiceServer) Fail(context.Context, *FailRequest) (*FailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Fail not implemented")
}
func (UnimplementedCrossServiceServer) mustEmbedUnimplementedCrossServiceServer() {}

// UnsafeCrossServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CrossServiceServer will
// result in compilation errors.
type UnsafeCrossServiceServer interface {
	mustEmbedUnimplementedCrossServiceServer()
}

func RegisterCrossServiceServer(s grpc.ServiceRegistrar, srv CrossServiceServer) {
	s.RegisterService(&CrossService_ServiceDesc, srv)
}

func _CrossService_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CrossServiceServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/internal.crosstest.v1test.CrossService/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CrossServiceServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CrossService_Fail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CrossServiceServer).Fail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/internal.crosstest.v1test.CrossService/Fail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CrossServiceServer).Fail(ctx, req.(*FailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CrossService_ServiceDesc is the grpc.ServiceDesc for CrossService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CrossService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "internal.crosstest.v1test.CrossService",
	HandlerType: (*CrossServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _CrossService_Ping_Handler,
		},
		{
			MethodName: "Fail",
			Handler:    _CrossService_Fail_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "internal/crosstest/v1test/cross.proto",
}
