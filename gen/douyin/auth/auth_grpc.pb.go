// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.9
// source: douyin/auth.proto

package auth

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

// LoginServiceClient is the client API for LoginService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LoginServiceClient interface {
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
}

type loginServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLoginServiceClient(cc grpc.ClientConnInterface) LoginServiceClient {
	return &loginServiceClient{cc}
}

func (c *loginServiceClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, "/douyin.LoginService/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LoginServiceServer is the server API for LoginService service.
// All implementations must embed UnimplementedLoginServiceServer
// for forward compatibility
type LoginServiceServer interface {
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	mustEmbedUnimplementedLoginServiceServer()
}

// UnimplementedLoginServiceServer must be embedded to have forward compatible implementations.
type UnimplementedLoginServiceServer struct {
}

func (UnimplementedLoginServiceServer) Login(context.Context, *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedLoginServiceServer) mustEmbedUnimplementedLoginServiceServer() {}

// UnsafeLoginServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LoginServiceServer will
// result in compilation errors.
type UnsafeLoginServiceServer interface {
	mustEmbedUnimplementedLoginServiceServer()
}

func RegisterLoginServiceServer(s grpc.ServiceRegistrar, srv LoginServiceServer) {
	s.RegisterService(&LoginService_ServiceDesc, srv)
}

func _LoginService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoginServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/douyin.LoginService/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoginServiceServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// LoginService_ServiceDesc is the grpc.ServiceDesc for LoginService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LoginService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "douyin.LoginService",
	HandlerType: (*LoginServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _LoginService_Login_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "douyin/auth.proto",
}

// RegisterServiceClient is the client API for RegisterService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RegisterServiceClient interface {
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error)
}

type registerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRegisterServiceClient(cc grpc.ClientConnInterface) RegisterServiceClient {
	return &registerServiceClient{cc}
}

func (c *registerServiceClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error) {
	out := new(RegisterResponse)
	err := c.cc.Invoke(ctx, "/douyin.RegisterService/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RegisterServiceServer is the server API for RegisterService service.
// All implementations must embed UnimplementedRegisterServiceServer
// for forward compatibility
type RegisterServiceServer interface {
	Register(context.Context, *RegisterRequest) (*RegisterResponse, error)
	mustEmbedUnimplementedRegisterServiceServer()
}

// UnimplementedRegisterServiceServer must be embedded to have forward compatible implementations.
type UnimplementedRegisterServiceServer struct {
}

func (UnimplementedRegisterServiceServer) Register(context.Context, *RegisterRequest) (*RegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedRegisterServiceServer) mustEmbedUnimplementedRegisterServiceServer() {}

// UnsafeRegisterServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RegisterServiceServer will
// result in compilation errors.
type UnsafeRegisterServiceServer interface {
	mustEmbedUnimplementedRegisterServiceServer()
}

func RegisterRegisterServiceServer(s grpc.ServiceRegistrar, srv RegisterServiceServer) {
	s.RegisterService(&RegisterService_ServiceDesc, srv)
}

func _RegisterService_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegisterServiceServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/douyin.RegisterService/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegisterServiceServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RegisterService_ServiceDesc is the grpc.ServiceDesc for RegisterService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RegisterService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "douyin.RegisterService",
	HandlerType: (*RegisterServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _RegisterService_Register_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "douyin/auth.proto",
}
