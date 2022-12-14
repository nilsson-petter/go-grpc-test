// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.9
// source: greeting.proto

package greeting

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

// GreetorClient is the client API for Greetor service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GreetorClient interface {
	SayHello(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type greetorClient struct {
	cc grpc.ClientConnInterface
}

func NewGreetorClient(cc grpc.ClientConnInterface) GreetorClient {
	return &greetorClient{cc}
}

func (c *greetorClient) SayHello(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/greeting.Greetor/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GreetorServer is the server API for Greetor service.
// All implementations must embed UnimplementedGreetorServer
// for forward compatibility
type GreetorServer interface {
	SayHello(context.Context, *Request) (*Response, error)
	mustEmbedUnimplementedGreetorServer()
}

// UnimplementedGreetorServer must be embedded to have forward compatible implementations.
type UnimplementedGreetorServer struct {
}

func (UnimplementedGreetorServer) SayHello(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (UnimplementedGreetorServer) mustEmbedUnimplementedGreetorServer() {}

// UnsafeGreetorServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GreetorServer will
// result in compilation errors.
type UnsafeGreetorServer interface {
	mustEmbedUnimplementedGreetorServer()
}

func RegisterGreetorServer(s grpc.ServiceRegistrar, srv GreetorServer) {
	s.RegisterService(&Greetor_ServiceDesc, srv)
}

func _Greetor_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreetorServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/greeting.Greetor/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreetorServer).SayHello(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

// Greetor_ServiceDesc is the grpc.ServiceDesc for Greetor service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Greetor_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "greeting.Greetor",
	HandlerType: (*GreetorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _Greetor_SayHello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "greeting.proto",
}
