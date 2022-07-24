// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: proto/http_webapp.proto

package operations_ecosys

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

// WebAppServicesClient is the client API for WebAppServices service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WebAppServicesClient interface {
	// <><> Telegram WebApp - HTTP <><>
	GetRosterAssignmentsForWebApp(ctx context.Context, in *HTTPRosterMessage, opts ...grpc.CallOption) (*HTTPMessage, error)
	PostWReportFromWebApp(ctx context.Context, in *IncidentReport, opts ...grpc.CallOption) (*HTTPMessage, error)
}

type webAppServicesClient struct {
	cc grpc.ClientConnInterface
}

func NewWebAppServicesClient(cc grpc.ClientConnInterface) WebAppServicesClient {
	return &webAppServicesClient{cc}
}

func (c *webAppServicesClient) GetRosterAssignmentsForWebApp(ctx context.Context, in *HTTPRosterMessage, opts ...grpc.CallOption) (*HTTPMessage, error) {
	out := new(HTTPMessage)
	err := c.cc.Invoke(ctx, "/http_webapp.WebAppServices/GetRosterAssignmentsForWebApp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *webAppServicesClient) PostWReportFromWebApp(ctx context.Context, in *IncidentReport, opts ...grpc.CallOption) (*HTTPMessage, error) {
	out := new(HTTPMessage)
	err := c.cc.Invoke(ctx, "/http_webapp.WebAppServices/PostWReportFromWebApp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WebAppServicesServer is the server API for WebAppServices service.
// All implementations must embed UnimplementedWebAppServicesServer
// for forward compatibility
type WebAppServicesServer interface {
	// <><> Telegram WebApp - HTTP <><>
	GetRosterAssignmentsForWebApp(context.Context, *HTTPRosterMessage) (*HTTPMessage, error)
	PostWReportFromWebApp(context.Context, *IncidentReport) (*HTTPMessage, error)
	mustEmbedUnimplementedWebAppServicesServer()
}

// UnimplementedWebAppServicesServer must be embedded to have forward compatible implementations.
type UnimplementedWebAppServicesServer struct {
}

func (UnimplementedWebAppServicesServer) GetRosterAssignmentsForWebApp(context.Context, *HTTPRosterMessage) (*HTTPMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRosterAssignmentsForWebApp not implemented")
}
func (UnimplementedWebAppServicesServer) PostWReportFromWebApp(context.Context, *IncidentReport) (*HTTPMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostWReportFromWebApp not implemented")
}
func (UnimplementedWebAppServicesServer) mustEmbedUnimplementedWebAppServicesServer() {}

// UnsafeWebAppServicesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WebAppServicesServer will
// result in compilation errors.
type UnsafeWebAppServicesServer interface {
	mustEmbedUnimplementedWebAppServicesServer()
}

func RegisterWebAppServicesServer(s grpc.ServiceRegistrar, srv WebAppServicesServer) {
	s.RegisterService(&WebAppServices_ServiceDesc, srv)
}

func _WebAppServices_GetRosterAssignmentsForWebApp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HTTPRosterMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WebAppServicesServer).GetRosterAssignmentsForWebApp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/http_webapp.WebAppServices/GetRosterAssignmentsForWebApp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WebAppServicesServer).GetRosterAssignmentsForWebApp(ctx, req.(*HTTPRosterMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _WebAppServices_PostWReportFromWebApp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IncidentReport)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WebAppServicesServer).PostWReportFromWebApp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/http_webapp.WebAppServices/PostWReportFromWebApp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WebAppServicesServer).PostWReportFromWebApp(ctx, req.(*IncidentReport))
	}
	return interceptor(ctx, in, info, handler)
}

// WebAppServices_ServiceDesc is the grpc.ServiceDesc for WebAppServices service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WebAppServices_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "http_webapp.WebAppServices",
	HandlerType: (*WebAppServicesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetRosterAssignmentsForWebApp",
			Handler:    _WebAppServices_GetRosterAssignmentsForWebApp_Handler,
		},
		{
			MethodName: "PostWReportFromWebApp",
			Handler:    _WebAppServices_PostWReportFromWebApp_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/http_webapp.proto",
}
