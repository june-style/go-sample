// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: sign.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	SignService_In_FullMethodName = "/api.SignService/In"
	SignService_Up_FullMethodName = "/api.SignService/Up"
)

// SignServiceClient is the client API for SignService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SignServiceClient interface {
	In(ctx context.Context, in *SignInRequest, opts ...grpc.CallOption) (*SignInResponse, error)
	Up(ctx context.Context, in *SignUpRequest, opts ...grpc.CallOption) (*SignUpResponse, error)
}

type signServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSignServiceClient(cc grpc.ClientConnInterface) SignServiceClient {
	return &signServiceClient{cc}
}

func (c *signServiceClient) In(ctx context.Context, in *SignInRequest, opts ...grpc.CallOption) (*SignInResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SignInResponse)
	err := c.cc.Invoke(ctx, SignService_In_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *signServiceClient) Up(ctx context.Context, in *SignUpRequest, opts ...grpc.CallOption) (*SignUpResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SignUpResponse)
	err := c.cc.Invoke(ctx, SignService_Up_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SignServiceServer is the server API for SignService service.
// All implementations must embed UnimplementedSignServiceServer
// for forward compatibility.
type SignServiceServer interface {
	In(context.Context, *SignInRequest) (*SignInResponse, error)
	Up(context.Context, *SignUpRequest) (*SignUpResponse, error)
	mustEmbedUnimplementedSignServiceServer()
}

// UnimplementedSignServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedSignServiceServer struct{}

func (UnimplementedSignServiceServer) In(context.Context, *SignInRequest) (*SignInResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method In not implemented")
}
func (UnimplementedSignServiceServer) Up(context.Context, *SignUpRequest) (*SignUpResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Up not implemented")
}
func (UnimplementedSignServiceServer) mustEmbedUnimplementedSignServiceServer() {}
func (UnimplementedSignServiceServer) testEmbeddedByValue()                     {}

// UnsafeSignServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SignServiceServer will
// result in compilation errors.
type UnsafeSignServiceServer interface {
	mustEmbedUnimplementedSignServiceServer()
}

func RegisterSignServiceServer(s grpc.ServiceRegistrar, srv SignServiceServer) {
	// If the following call pancis, it indicates UnimplementedSignServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&SignService_ServiceDesc, srv)
}

func _SignService_In_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignInRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SignServiceServer).In(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SignService_In_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SignServiceServer).In(ctx, req.(*SignInRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SignService_Up_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignUpRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SignServiceServer).Up(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SignService_Up_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SignServiceServer).Up(ctx, req.(*SignUpRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SignService_ServiceDesc is the grpc.ServiceDesc for SignService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SignService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.SignService",
	HandlerType: (*SignServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "In",
			Handler:    _SignService_In_Handler,
		},
		{
			MethodName: "Up",
			Handler:    _SignService_Up_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sign.proto",
}
