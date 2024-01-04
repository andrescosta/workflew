// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.1
// source: control.proto

package types

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

const (
	Control_GetTenants_FullMethodName             = "/Control/GetTenants"
	Control_AddTenant_FullMethodName              = "/Control/AddTenant"
	Control_AddPackage_FullMethodName             = "/Control/AddPackage"
	Control_GetAllPackages_FullMethodName         = "/Control/GetAllPackages"
	Control_GetPackages_FullMethodName            = "/Control/GetPackages"
	Control_UpdatePackage_FullMethodName          = "/Control/UpdatePackage"
	Control_DeletePackage_FullMethodName          = "/Control/DeletePackage"
	Control_UpdateToPackagesStr_FullMethodName    = "/Control/UpdateToPackagesStr"
	Control_GetEnvironment_FullMethodName         = "/Control/GetEnvironment"
	Control_UpdateToEnvironmentStr_FullMethodName = "/Control/UpdateToEnvironmentStr"
	Control_AddEnvironment_FullMethodName         = "/Control/AddEnvironment"
	Control_UpdateEnvironment_FullMethodName      = "/Control/UpdateEnvironment"
)

// ControlClient is the client API for Control service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ControlClient interface {
	GetTenants(ctx context.Context, in *GetTenantsRequest, opts ...grpc.CallOption) (*GetTenantsReply, error)
	AddTenant(ctx context.Context, in *AddTenantRequest, opts ...grpc.CallOption) (*AddTenantReply, error)
	AddPackage(ctx context.Context, in *AddJobPackageRequest, opts ...grpc.CallOption) (*AddJobPackageReply, error)
	GetAllPackages(ctx context.Context, in *GetAllJobPackagesRequest, opts ...grpc.CallOption) (*GetAllJobPackagesReply, error)
	GetPackages(ctx context.Context, in *GetJobPackagesRequest, opts ...grpc.CallOption) (*GetJobPackagesReply, error)
	UpdatePackage(ctx context.Context, in *UpdateJobPackageRequest, opts ...grpc.CallOption) (*UpdateJobPackageReply, error)
	DeletePackage(ctx context.Context, in *DeleteJobPackageRequest, opts ...grpc.CallOption) (*DeleteJobPackageReply, error)
	UpdateToPackagesStr(ctx context.Context, in *UpdateToPackagesStrRequest, opts ...grpc.CallOption) (Control_UpdateToPackagesStrClient, error)
	GetEnvironment(ctx context.Context, in *GetEnvironmentRequest, opts ...grpc.CallOption) (*GetEnvironmentReply, error)
	UpdateToEnvironmentStr(ctx context.Context, in *UpdateToEnvironmentStrRequest, opts ...grpc.CallOption) (Control_UpdateToEnvironmentStrClient, error)
	AddEnvironment(ctx context.Context, in *AddEnvironmentRequest, opts ...grpc.CallOption) (*AddEnvironmentReply, error)
	UpdateEnvironment(ctx context.Context, in *UpdateEnvironmentRequest, opts ...grpc.CallOption) (*UpdateEnvironmentReply, error)
}

type controlClient struct {
	cc grpc.ClientConnInterface
}

func NewControlClient(cc grpc.ClientConnInterface) ControlClient {
	return &controlClient{cc}
}

func (c *controlClient) GetTenants(ctx context.Context, in *GetTenantsRequest, opts ...grpc.CallOption) (*GetTenantsReply, error) {
	out := new(GetTenantsReply)
	err := c.cc.Invoke(ctx, Control_GetTenants_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *controlClient) AddTenant(ctx context.Context, in *AddTenantRequest, opts ...grpc.CallOption) (*AddTenantReply, error) {
	out := new(AddTenantReply)
	err := c.cc.Invoke(ctx, Control_AddTenant_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *controlClient) AddPackage(ctx context.Context, in *AddJobPackageRequest, opts ...grpc.CallOption) (*AddJobPackageReply, error) {
	out := new(AddJobPackageReply)
	err := c.cc.Invoke(ctx, Control_AddPackage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *controlClient) GetAllPackages(ctx context.Context, in *GetAllJobPackagesRequest, opts ...grpc.CallOption) (*GetAllJobPackagesReply, error) {
	out := new(GetAllJobPackagesReply)
	err := c.cc.Invoke(ctx, Control_GetAllPackages_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *controlClient) GetPackages(ctx context.Context, in *GetJobPackagesRequest, opts ...grpc.CallOption) (*GetJobPackagesReply, error) {
	out := new(GetJobPackagesReply)
	err := c.cc.Invoke(ctx, Control_GetPackages_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *controlClient) UpdatePackage(ctx context.Context, in *UpdateJobPackageRequest, opts ...grpc.CallOption) (*UpdateJobPackageReply, error) {
	out := new(UpdateJobPackageReply)
	err := c.cc.Invoke(ctx, Control_UpdatePackage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *controlClient) DeletePackage(ctx context.Context, in *DeleteJobPackageRequest, opts ...grpc.CallOption) (*DeleteJobPackageReply, error) {
	out := new(DeleteJobPackageReply)
	err := c.cc.Invoke(ctx, Control_DeletePackage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *controlClient) UpdateToPackagesStr(ctx context.Context, in *UpdateToPackagesStrRequest, opts ...grpc.CallOption) (Control_UpdateToPackagesStrClient, error) {
	stream, err := c.cc.NewStream(ctx, &Control_ServiceDesc.Streams[0], Control_UpdateToPackagesStr_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &controlUpdateToPackagesStrClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Control_UpdateToPackagesStrClient interface {
	Recv() (*UpdateToPackagesStrReply, error)
	grpc.ClientStream
}

type controlUpdateToPackagesStrClient struct {
	grpc.ClientStream
}

func (x *controlUpdateToPackagesStrClient) Recv() (*UpdateToPackagesStrReply, error) {
	m := new(UpdateToPackagesStrReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *controlClient) GetEnvironment(ctx context.Context, in *GetEnvironmentRequest, opts ...grpc.CallOption) (*GetEnvironmentReply, error) {
	out := new(GetEnvironmentReply)
	err := c.cc.Invoke(ctx, Control_GetEnvironment_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *controlClient) UpdateToEnvironmentStr(ctx context.Context, in *UpdateToEnvironmentStrRequest, opts ...grpc.CallOption) (Control_UpdateToEnvironmentStrClient, error) {
	stream, err := c.cc.NewStream(ctx, &Control_ServiceDesc.Streams[1], Control_UpdateToEnvironmentStr_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &controlUpdateToEnvironmentStrClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Control_UpdateToEnvironmentStrClient interface {
	Recv() (*UpdateToEnvironmentStrReply, error)
	grpc.ClientStream
}

type controlUpdateToEnvironmentStrClient struct {
	grpc.ClientStream
}

func (x *controlUpdateToEnvironmentStrClient) Recv() (*UpdateToEnvironmentStrReply, error) {
	m := new(UpdateToEnvironmentStrReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *controlClient) AddEnvironment(ctx context.Context, in *AddEnvironmentRequest, opts ...grpc.CallOption) (*AddEnvironmentReply, error) {
	out := new(AddEnvironmentReply)
	err := c.cc.Invoke(ctx, Control_AddEnvironment_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *controlClient) UpdateEnvironment(ctx context.Context, in *UpdateEnvironmentRequest, opts ...grpc.CallOption) (*UpdateEnvironmentReply, error) {
	out := new(UpdateEnvironmentReply)
	err := c.cc.Invoke(ctx, Control_UpdateEnvironment_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ControlServer is the server API for Control service.
// All implementations must embed UnimplementedControlServer
// for forward compatibility
type ControlServer interface {
	GetTenants(context.Context, *GetTenantsRequest) (*GetTenantsReply, error)
	AddTenant(context.Context, *AddTenantRequest) (*AddTenantReply, error)
	AddPackage(context.Context, *AddJobPackageRequest) (*AddJobPackageReply, error)
	GetAllPackages(context.Context, *GetAllJobPackagesRequest) (*GetAllJobPackagesReply, error)
	GetPackages(context.Context, *GetJobPackagesRequest) (*GetJobPackagesReply, error)
	UpdatePackage(context.Context, *UpdateJobPackageRequest) (*UpdateJobPackageReply, error)
	DeletePackage(context.Context, *DeleteJobPackageRequest) (*DeleteJobPackageReply, error)
	UpdateToPackagesStr(*UpdateToPackagesStrRequest, Control_UpdateToPackagesStrServer) error
	GetEnvironment(context.Context, *GetEnvironmentRequest) (*GetEnvironmentReply, error)
	UpdateToEnvironmentStr(*UpdateToEnvironmentStrRequest, Control_UpdateToEnvironmentStrServer) error
	AddEnvironment(context.Context, *AddEnvironmentRequest) (*AddEnvironmentReply, error)
	UpdateEnvironment(context.Context, *UpdateEnvironmentRequest) (*UpdateEnvironmentReply, error)
	mustEmbedUnimplementedControlServer()
}

// UnimplementedControlServer must be embedded to have forward compatible implementations.
type UnimplementedControlServer struct {
}

func (UnimplementedControlServer) GetTenants(context.Context, *GetTenantsRequest) (*GetTenantsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTenants not implemented")
}
func (UnimplementedControlServer) AddTenant(context.Context, *AddTenantRequest) (*AddTenantReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddTenant not implemented")
}
func (UnimplementedControlServer) AddPackage(context.Context, *AddJobPackageRequest) (*AddJobPackageReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddPackage not implemented")
}
func (UnimplementedControlServer) GetAllPackages(context.Context, *GetAllJobPackagesRequest) (*GetAllJobPackagesReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllPackages not implemented")
}
func (UnimplementedControlServer) GetPackages(context.Context, *GetJobPackagesRequest) (*GetJobPackagesReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPackages not implemented")
}
func (UnimplementedControlServer) UpdatePackage(context.Context, *UpdateJobPackageRequest) (*UpdateJobPackageReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePackage not implemented")
}
func (UnimplementedControlServer) DeletePackage(context.Context, *DeleteJobPackageRequest) (*DeleteJobPackageReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePackage not implemented")
}
func (UnimplementedControlServer) UpdateToPackagesStr(*UpdateToPackagesStrRequest, Control_UpdateToPackagesStrServer) error {
	return status.Errorf(codes.Unimplemented, "method UpdateToPackagesStr not implemented")
}
func (UnimplementedControlServer) GetEnvironment(context.Context, *GetEnvironmentRequest) (*GetEnvironmentReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEnvironment not implemented")
}
func (UnimplementedControlServer) UpdateToEnvironmentStr(*UpdateToEnvironmentStrRequest, Control_UpdateToEnvironmentStrServer) error {
	return status.Errorf(codes.Unimplemented, "method UpdateToEnvironmentStr not implemented")
}
func (UnimplementedControlServer) AddEnvironment(context.Context, *AddEnvironmentRequest) (*AddEnvironmentReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddEnvironment not implemented")
}
func (UnimplementedControlServer) UpdateEnvironment(context.Context, *UpdateEnvironmentRequest) (*UpdateEnvironmentReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateEnvironment not implemented")
}
func (UnimplementedControlServer) mustEmbedUnimplementedControlServer() {}

// UnsafeControlServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ControlServer will
// result in compilation errors.
type UnsafeControlServer interface {
	mustEmbedUnimplementedControlServer()
}

func RegisterControlServer(s grpc.ServiceRegistrar, srv ControlServer) {
	s.RegisterService(&Control_ServiceDesc, srv)
}

func _Control_GetTenants_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTenantsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ControlServer).GetTenants(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Control_GetTenants_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ControlServer).GetTenants(ctx, req.(*GetTenantsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Control_AddTenant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddTenantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ControlServer).AddTenant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Control_AddTenant_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ControlServer).AddTenant(ctx, req.(*AddTenantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Control_AddPackage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddJobPackageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ControlServer).AddPackage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Control_AddPackage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ControlServer).AddPackage(ctx, req.(*AddJobPackageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Control_GetAllPackages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllJobPackagesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ControlServer).GetAllPackages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Control_GetAllPackages_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ControlServer).GetAllPackages(ctx, req.(*GetAllJobPackagesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Control_GetPackages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetJobPackagesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ControlServer).GetPackages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Control_GetPackages_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ControlServer).GetPackages(ctx, req.(*GetJobPackagesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Control_UpdatePackage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateJobPackageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ControlServer).UpdatePackage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Control_UpdatePackage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ControlServer).UpdatePackage(ctx, req.(*UpdateJobPackageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Control_DeletePackage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteJobPackageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ControlServer).DeletePackage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Control_DeletePackage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ControlServer).DeletePackage(ctx, req.(*DeleteJobPackageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Control_UpdateToPackagesStr_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(UpdateToPackagesStrRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ControlServer).UpdateToPackagesStr(m, &controlUpdateToPackagesStrServer{stream})
}

type Control_UpdateToPackagesStrServer interface {
	Send(*UpdateToPackagesStrReply) error
	grpc.ServerStream
}

type controlUpdateToPackagesStrServer struct {
	grpc.ServerStream
}

func (x *controlUpdateToPackagesStrServer) Send(m *UpdateToPackagesStrReply) error {
	return x.ServerStream.SendMsg(m)
}

func _Control_GetEnvironment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEnvironmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ControlServer).GetEnvironment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Control_GetEnvironment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ControlServer).GetEnvironment(ctx, req.(*GetEnvironmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Control_UpdateToEnvironmentStr_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(UpdateToEnvironmentStrRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ControlServer).UpdateToEnvironmentStr(m, &controlUpdateToEnvironmentStrServer{stream})
}

type Control_UpdateToEnvironmentStrServer interface {
	Send(*UpdateToEnvironmentStrReply) error
	grpc.ServerStream
}

type controlUpdateToEnvironmentStrServer struct {
	grpc.ServerStream
}

func (x *controlUpdateToEnvironmentStrServer) Send(m *UpdateToEnvironmentStrReply) error {
	return x.ServerStream.SendMsg(m)
}

func _Control_AddEnvironment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddEnvironmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ControlServer).AddEnvironment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Control_AddEnvironment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ControlServer).AddEnvironment(ctx, req.(*AddEnvironmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Control_UpdateEnvironment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateEnvironmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ControlServer).UpdateEnvironment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Control_UpdateEnvironment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ControlServer).UpdateEnvironment(ctx, req.(*UpdateEnvironmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Control_ServiceDesc is the grpc.ServiceDesc for Control service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Control_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Control",
	HandlerType: (*ControlServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTenants",
			Handler:    _Control_GetTenants_Handler,
		},
		{
			MethodName: "AddTenant",
			Handler:    _Control_AddTenant_Handler,
		},
		{
			MethodName: "AddPackage",
			Handler:    _Control_AddPackage_Handler,
		},
		{
			MethodName: "GetAllPackages",
			Handler:    _Control_GetAllPackages_Handler,
		},
		{
			MethodName: "GetPackages",
			Handler:    _Control_GetPackages_Handler,
		},
		{
			MethodName: "UpdatePackage",
			Handler:    _Control_UpdatePackage_Handler,
		},
		{
			MethodName: "DeletePackage",
			Handler:    _Control_DeletePackage_Handler,
		},
		{
			MethodName: "GetEnvironment",
			Handler:    _Control_GetEnvironment_Handler,
		},
		{
			MethodName: "AddEnvironment",
			Handler:    _Control_AddEnvironment_Handler,
		},
		{
			MethodName: "UpdateEnvironment",
			Handler:    _Control_UpdateEnvironment_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "UpdateToPackagesStr",
			Handler:       _Control_UpdateToPackagesStr_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "UpdateToEnvironmentStr",
			Handler:       _Control_UpdateToEnvironmentStr_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "control.proto",
}
