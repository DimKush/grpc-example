// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: grpc-example.proto

package grpc_example

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	GrpcExampleServiceV1_Alive_FullMethodName = "/DimKush.grpc_example.v1.GrpcExampleServiceV1/Alive"
)

// GrpcExampleServiceV1Client is the client API for GrpcExampleServiceV1 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GrpcExampleServiceV1Client interface {
	Alive(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type grpcExampleServiceV1Client struct {
	cc grpc.ClientConnInterface
}

func NewGrpcExampleServiceV1Client(cc grpc.ClientConnInterface) GrpcExampleServiceV1Client {
	return &grpcExampleServiceV1Client{cc}
}

func (c *grpcExampleServiceV1Client) Alive(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, GrpcExampleServiceV1_Alive_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GrpcExampleServiceV1Server is the server API for GrpcExampleServiceV1 service.
// All implementations must embed UnimplementedGrpcExampleServiceV1Server
// for forward compatibility
type GrpcExampleServiceV1Server interface {
	Alive(context.Context, *emptypb.Empty) (*emptypb.Empty, error)
	mustEmbedUnimplementedGrpcExampleServiceV1Server()
}

// UnimplementedGrpcExampleServiceV1Server must be embedded to have forward compatible implementations.
type UnimplementedGrpcExampleServiceV1Server struct {
}

func (UnimplementedGrpcExampleServiceV1Server) Alive(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Alive not implemented")
}
func (UnimplementedGrpcExampleServiceV1Server) mustEmbedUnimplementedGrpcExampleServiceV1Server() {}

// UnsafeGrpcExampleServiceV1Server may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GrpcExampleServiceV1Server will
// result in compilation errors.
type UnsafeGrpcExampleServiceV1Server interface {
	mustEmbedUnimplementedGrpcExampleServiceV1Server()
}

func RegisterGrpcExampleServiceV1Server(s grpc.ServiceRegistrar, srv GrpcExampleServiceV1Server) {
	s.RegisterService(&GrpcExampleServiceV1_ServiceDesc, srv)
}

func _GrpcExampleServiceV1_Alive_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GrpcExampleServiceV1Server).Alive(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GrpcExampleServiceV1_Alive_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GrpcExampleServiceV1Server).Alive(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// GrpcExampleServiceV1_ServiceDesc is the grpc.ServiceDesc for GrpcExampleServiceV1 service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GrpcExampleServiceV1_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "DimKush.grpc_example.v1.GrpcExampleServiceV1",
	HandlerType: (*GrpcExampleServiceV1Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Alive",
			Handler:    _GrpcExampleServiceV1_Alive_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc-example.proto",
}
