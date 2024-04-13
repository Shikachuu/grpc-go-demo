// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.26.1
// source: proto/templates.proto

package proto

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

// TemplateServiceClient is the client API for TemplateService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TemplateServiceClient interface {
	GetTemplateById(ctx context.Context, in *GetTemplateRequest, opts ...grpc.CallOption) (*TemplateResponse, error)
	CreateTemplate(ctx context.Context, in *TemplateRequest, opts ...grpc.CallOption) (*TemplateResponse, error)
}

type templateServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTemplateServiceClient(cc grpc.ClientConnInterface) TemplateServiceClient {
	return &templateServiceClient{cc}
}

func (c *templateServiceClient) GetTemplateById(ctx context.Context, in *GetTemplateRequest, opts ...grpc.CallOption) (*TemplateResponse, error) {
	out := new(TemplateResponse)
	err := c.cc.Invoke(ctx, "/template.TemplateService/GetTemplateById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *templateServiceClient) CreateTemplate(ctx context.Context, in *TemplateRequest, opts ...grpc.CallOption) (*TemplateResponse, error) {
	out := new(TemplateResponse)
	err := c.cc.Invoke(ctx, "/template.TemplateService/CreateTemplate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TemplateServiceServer is the server API for TemplateService service.
// All implementations must embed UnimplementedTemplateServiceServer
// for forward compatibility
type TemplateServiceServer interface {
	GetTemplateById(context.Context, *GetTemplateRequest) (*TemplateResponse, error)
	CreateTemplate(context.Context, *TemplateRequest) (*TemplateResponse, error)
	mustEmbedUnimplementedTemplateServiceServer()
}

// UnimplementedTemplateServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTemplateServiceServer struct {
}

func (UnimplementedTemplateServiceServer) GetTemplateById(context.Context, *GetTemplateRequest) (*TemplateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTemplateById not implemented")
}
func (UnimplementedTemplateServiceServer) CreateTemplate(context.Context, *TemplateRequest) (*TemplateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTemplate not implemented")
}
func (UnimplementedTemplateServiceServer) mustEmbedUnimplementedTemplateServiceServer() {}

// UnsafeTemplateServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TemplateServiceServer will
// result in compilation errors.
type UnsafeTemplateServiceServer interface {
	mustEmbedUnimplementedTemplateServiceServer()
}

func RegisterTemplateServiceServer(s grpc.ServiceRegistrar, srv TemplateServiceServer) {
	s.RegisterService(&TemplateService_ServiceDesc, srv)
}

func _TemplateService_GetTemplateById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTemplateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TemplateServiceServer).GetTemplateById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/template.TemplateService/GetTemplateById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TemplateServiceServer).GetTemplateById(ctx, req.(*GetTemplateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TemplateService_CreateTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TemplateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TemplateServiceServer).CreateTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/template.TemplateService/CreateTemplate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TemplateServiceServer).CreateTemplate(ctx, req.(*TemplateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TemplateService_ServiceDesc is the grpc.ServiceDesc for TemplateService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TemplateService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "template.TemplateService",
	HandlerType: (*TemplateServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTemplateById",
			Handler:    _TemplateService_GetTemplateById_Handler,
		},
		{
			MethodName: "CreateTemplate",
			Handler:    _TemplateService_CreateTemplate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/templates.proto",
}
