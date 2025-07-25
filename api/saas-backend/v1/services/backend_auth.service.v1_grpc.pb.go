// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v6.30.2
// source: api/saas-backend/v1/services/backend_auth.service.v1.proto

package servicev1

import (
	context "context"
	resources "github.com/go-micro-saas/service-api/api/account-service/v1/resources"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	SrvSaasBackendAuthV1_LoginByEmail_FullMethodName = "/saas.api.backend.servicev1.SrvSaasBackendAuthV1/LoginByEmail"
	SrvSaasBackendAuthV1_LoginByPhone_FullMethodName = "/saas.api.backend.servicev1.SrvSaasBackendAuthV1/LoginByPhone"
)

// SrvSaasBackendAuthV1Client is the client API for SrvSaasBackendAuthV1 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// SrvSaasBackendAuthV1 ...
type SrvSaasBackendAuthV1Client interface {
	// 身份验证-Email登录
	LoginByEmail(ctx context.Context, in *resources.LoginByEmailReq, opts ...grpc.CallOption) (*resources.LoginResp, error)
	// 身份验证-手机登录
	LoginByPhone(ctx context.Context, in *resources.LoginByPhoneReq, opts ...grpc.CallOption) (*resources.LoginResp, error)
}

type srvSaasBackendAuthV1Client struct {
	cc grpc.ClientConnInterface
}

func NewSrvSaasBackendAuthV1Client(cc grpc.ClientConnInterface) SrvSaasBackendAuthV1Client {
	return &srvSaasBackendAuthV1Client{cc}
}

func (c *srvSaasBackendAuthV1Client) LoginByEmail(ctx context.Context, in *resources.LoginByEmailReq, opts ...grpc.CallOption) (*resources.LoginResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(resources.LoginResp)
	err := c.cc.Invoke(ctx, SrvSaasBackendAuthV1_LoginByEmail_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *srvSaasBackendAuthV1Client) LoginByPhone(ctx context.Context, in *resources.LoginByPhoneReq, opts ...grpc.CallOption) (*resources.LoginResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(resources.LoginResp)
	err := c.cc.Invoke(ctx, SrvSaasBackendAuthV1_LoginByPhone_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SrvSaasBackendAuthV1Server is the server API for SrvSaasBackendAuthV1 service.
// All implementations must embed UnimplementedSrvSaasBackendAuthV1Server
// for forward compatibility.
//
// SrvSaasBackendAuthV1 ...
type SrvSaasBackendAuthV1Server interface {
	// 身份验证-Email登录
	LoginByEmail(context.Context, *resources.LoginByEmailReq) (*resources.LoginResp, error)
	// 身份验证-手机登录
	LoginByPhone(context.Context, *resources.LoginByPhoneReq) (*resources.LoginResp, error)
	mustEmbedUnimplementedSrvSaasBackendAuthV1Server()
}

// UnimplementedSrvSaasBackendAuthV1Server must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedSrvSaasBackendAuthV1Server struct{}

func (UnimplementedSrvSaasBackendAuthV1Server) LoginByEmail(context.Context, *resources.LoginByEmailReq) (*resources.LoginResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginByEmail not implemented")
}
func (UnimplementedSrvSaasBackendAuthV1Server) LoginByPhone(context.Context, *resources.LoginByPhoneReq) (*resources.LoginResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginByPhone not implemented")
}
func (UnimplementedSrvSaasBackendAuthV1Server) mustEmbedUnimplementedSrvSaasBackendAuthV1Server() {}
func (UnimplementedSrvSaasBackendAuthV1Server) testEmbeddedByValue()                              {}

// UnsafeSrvSaasBackendAuthV1Server may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SrvSaasBackendAuthV1Server will
// result in compilation errors.
type UnsafeSrvSaasBackendAuthV1Server interface {
	mustEmbedUnimplementedSrvSaasBackendAuthV1Server()
}

func RegisterSrvSaasBackendAuthV1Server(s grpc.ServiceRegistrar, srv SrvSaasBackendAuthV1Server) {
	// If the following call pancis, it indicates UnimplementedSrvSaasBackendAuthV1Server was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&SrvSaasBackendAuthV1_ServiceDesc, srv)
}

func _SrvSaasBackendAuthV1_LoginByEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(resources.LoginByEmailReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SrvSaasBackendAuthV1Server).LoginByEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SrvSaasBackendAuthV1_LoginByEmail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SrvSaasBackendAuthV1Server).LoginByEmail(ctx, req.(*resources.LoginByEmailReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SrvSaasBackendAuthV1_LoginByPhone_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(resources.LoginByPhoneReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SrvSaasBackendAuthV1Server).LoginByPhone(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SrvSaasBackendAuthV1_LoginByPhone_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SrvSaasBackendAuthV1Server).LoginByPhone(ctx, req.(*resources.LoginByPhoneReq))
	}
	return interceptor(ctx, in, info, handler)
}

// SrvSaasBackendAuthV1_ServiceDesc is the grpc.ServiceDesc for SrvSaasBackendAuthV1 service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SrvSaasBackendAuthV1_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "saas.api.backend.servicev1.SrvSaasBackendAuthV1",
	HandlerType: (*SrvSaasBackendAuthV1Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "LoginByEmail",
			Handler:    _SrvSaasBackendAuthV1_LoginByEmail_Handler,
		},
		{
			MethodName: "LoginByPhone",
			Handler:    _SrvSaasBackendAuthV1_LoginByPhone_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/saas-backend/v1/services/backend_auth.service.v1.proto",
}
