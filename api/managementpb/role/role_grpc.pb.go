// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             (unknown)
// source: managementpb/role/role.proto

package rolev1beta1

import (
	context "context"

	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	Role_CreateRole_FullMethodName     = "/role.v1beta1.Role/CreateRole"
	Role_UpdateRole_FullMethodName     = "/role.v1beta1.Role/UpdateRole"
	Role_DeleteRole_FullMethodName     = "/role.v1beta1.Role/DeleteRole"
	Role_GetRole_FullMethodName        = "/role.v1beta1.Role/GetRole"
	Role_ListRoles_FullMethodName      = "/role.v1beta1.Role/ListRoles"
	Role_AssignRoles_FullMethodName    = "/role.v1beta1.Role/AssignRoles"
	Role_SetDefaultRole_FullMethodName = "/role.v1beta1.Role/SetDefaultRole"
)

// RoleClient is the client API for Role service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Service Role provides public methods for managing Roles.
type RoleClient interface {
	// CreateRole creates a new role.
	CreateRole(ctx context.Context, in *CreateRoleRequest, opts ...grpc.CallOption) (*CreateRoleResponse, error)
	// UpdateRole creates a new role.
	UpdateRole(ctx context.Context, in *UpdateRoleRequest, opts ...grpc.CallOption) (*UpdateRoleResponse, error)
	// DeleteRole creates a new role.
	DeleteRole(ctx context.Context, in *DeleteRoleRequest, opts ...grpc.CallOption) (*DeleteRoleResponse, error)
	// GetRole retrieves a single role.
	GetRole(ctx context.Context, in *GetRoleRequest, opts ...grpc.CallOption) (*GetRoleResponse, error)
	// ListRoles retrieves a roles.
	ListRoles(ctx context.Context, in *ListRolesRequest, opts ...grpc.CallOption) (*ListRolesResponse, error)
	// AssignRoles replaces all assigned roles for a user.
	AssignRoles(ctx context.Context, in *AssignRolesRequest, opts ...grpc.CallOption) (*AssignRolesResponse, error)
	// SetDefaultRole configures default role assigned to users.
	SetDefaultRole(ctx context.Context, in *SetDefaultRoleRequest, opts ...grpc.CallOption) (*SetDefaultRoleResponse, error)
}

type roleClient struct {
	cc grpc.ClientConnInterface
}

func NewRoleClient(cc grpc.ClientConnInterface) RoleClient {
	return &roleClient{cc}
}

func (c *roleClient) CreateRole(ctx context.Context, in *CreateRoleRequest, opts ...grpc.CallOption) (*CreateRoleResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateRoleResponse)
	err := c.cc.Invoke(ctx, Role_CreateRole_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleClient) UpdateRole(ctx context.Context, in *UpdateRoleRequest, opts ...grpc.CallOption) (*UpdateRoleResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateRoleResponse)
	err := c.cc.Invoke(ctx, Role_UpdateRole_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleClient) DeleteRole(ctx context.Context, in *DeleteRoleRequest, opts ...grpc.CallOption) (*DeleteRoleResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteRoleResponse)
	err := c.cc.Invoke(ctx, Role_DeleteRole_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleClient) GetRole(ctx context.Context, in *GetRoleRequest, opts ...grpc.CallOption) (*GetRoleResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetRoleResponse)
	err := c.cc.Invoke(ctx, Role_GetRole_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleClient) ListRoles(ctx context.Context, in *ListRolesRequest, opts ...grpc.CallOption) (*ListRolesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListRolesResponse)
	err := c.cc.Invoke(ctx, Role_ListRoles_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleClient) AssignRoles(ctx context.Context, in *AssignRolesRequest, opts ...grpc.CallOption) (*AssignRolesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AssignRolesResponse)
	err := c.cc.Invoke(ctx, Role_AssignRoles_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleClient) SetDefaultRole(ctx context.Context, in *SetDefaultRoleRequest, opts ...grpc.CallOption) (*SetDefaultRoleResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SetDefaultRoleResponse)
	err := c.cc.Invoke(ctx, Role_SetDefaultRole_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RoleServer is the server API for Role service.
// All implementations must embed UnimplementedRoleServer
// for forward compatibility
//
// Service Role provides public methods for managing Roles.
type RoleServer interface {
	// CreateRole creates a new role.
	CreateRole(context.Context, *CreateRoleRequest) (*CreateRoleResponse, error)
	// UpdateRole creates a new role.
	UpdateRole(context.Context, *UpdateRoleRequest) (*UpdateRoleResponse, error)
	// DeleteRole creates a new role.
	DeleteRole(context.Context, *DeleteRoleRequest) (*DeleteRoleResponse, error)
	// GetRole retrieves a single role.
	GetRole(context.Context, *GetRoleRequest) (*GetRoleResponse, error)
	// ListRoles retrieves a roles.
	ListRoles(context.Context, *ListRolesRequest) (*ListRolesResponse, error)
	// AssignRoles replaces all assigned roles for a user.
	AssignRoles(context.Context, *AssignRolesRequest) (*AssignRolesResponse, error)
	// SetDefaultRole configures default role assigned to users.
	SetDefaultRole(context.Context, *SetDefaultRoleRequest) (*SetDefaultRoleResponse, error)
	mustEmbedUnimplementedRoleServer()
}

// UnimplementedRoleServer must be embedded to have forward compatible implementations.
type UnimplementedRoleServer struct{}

func (UnimplementedRoleServer) CreateRole(context.Context, *CreateRoleRequest) (*CreateRoleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRole not implemented")
}

func (UnimplementedRoleServer) UpdateRole(context.Context, *UpdateRoleRequest) (*UpdateRoleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRole not implemented")
}

func (UnimplementedRoleServer) DeleteRole(context.Context, *DeleteRoleRequest) (*DeleteRoleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRole not implemented")
}

func (UnimplementedRoleServer) GetRole(context.Context, *GetRoleRequest) (*GetRoleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRole not implemented")
}

func (UnimplementedRoleServer) ListRoles(context.Context, *ListRolesRequest) (*ListRolesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListRoles not implemented")
}

func (UnimplementedRoleServer) AssignRoles(context.Context, *AssignRolesRequest) (*AssignRolesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AssignRoles not implemented")
}

func (UnimplementedRoleServer) SetDefaultRole(context.Context, *SetDefaultRoleRequest) (*SetDefaultRoleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetDefaultRole not implemented")
}
func (UnimplementedRoleServer) mustEmbedUnimplementedRoleServer() {}

// UnsafeRoleServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RoleServer will
// result in compilation errors.
type UnsafeRoleServer interface {
	mustEmbedUnimplementedRoleServer()
}

func RegisterRoleServer(s grpc.ServiceRegistrar, srv RoleServer) {
	s.RegisterService(&Role_ServiceDesc, srv)
}

func _Role_CreateRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleServer).CreateRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Role_CreateRole_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleServer).CreateRole(ctx, req.(*CreateRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Role_UpdateRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleServer).UpdateRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Role_UpdateRole_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleServer).UpdateRole(ctx, req.(*UpdateRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Role_DeleteRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleServer).DeleteRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Role_DeleteRole_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleServer).DeleteRole(ctx, req.(*DeleteRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Role_GetRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleServer).GetRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Role_GetRole_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleServer).GetRole(ctx, req.(*GetRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Role_ListRoles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRolesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleServer).ListRoles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Role_ListRoles_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleServer).ListRoles(ctx, req.(*ListRolesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Role_AssignRoles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AssignRolesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleServer).AssignRoles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Role_AssignRoles_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleServer).AssignRoles(ctx, req.(*AssignRolesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Role_SetDefaultRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetDefaultRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleServer).SetDefaultRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Role_SetDefaultRole_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleServer).SetDefaultRole(ctx, req.(*SetDefaultRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Role_ServiceDesc is the grpc.ServiceDesc for Role service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Role_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "role.v1beta1.Role",
	HandlerType: (*RoleServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateRole",
			Handler:    _Role_CreateRole_Handler,
		},
		{
			MethodName: "UpdateRole",
			Handler:    _Role_UpdateRole_Handler,
		},
		{
			MethodName: "DeleteRole",
			Handler:    _Role_DeleteRole_Handler,
		},
		{
			MethodName: "GetRole",
			Handler:    _Role_GetRole_Handler,
		},
		{
			MethodName: "ListRoles",
			Handler:    _Role_ListRoles_Handler,
		},
		{
			MethodName: "AssignRoles",
			Handler:    _Role_AssignRoles_Handler,
		},
		{
			MethodName: "SetDefaultRole",
			Handler:    _Role_SetDefaultRole_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "managementpb/role/role.proto",
}
