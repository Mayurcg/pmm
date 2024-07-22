// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             (unknown)
// source: managementpb/backup/locations.proto

package backupv1

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
	Locations_ListLocations_FullMethodName      = "/backup.v1.Locations/ListLocations"
	Locations_AddLocation_FullMethodName        = "/backup.v1.Locations/AddLocation"
	Locations_ChangeLocation_FullMethodName     = "/backup.v1.Locations/ChangeLocation"
	Locations_RemoveLocation_FullMethodName     = "/backup.v1.Locations/RemoveLocation"
	Locations_TestLocationConfig_FullMethodName = "/backup.v1.Locations/TestLocationConfig"
)

// LocationsClient is the client API for Locations service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Locations service provides access to Backup Locations.
type LocationsClient interface {
	// ListLocations returns a list of all backup locations.
	ListLocations(ctx context.Context, in *ListLocationsRequest, opts ...grpc.CallOption) (*ListLocationsResponse, error)
	// AddLocation adds backup location.
	AddLocation(ctx context.Context, in *AddLocationRequest, opts ...grpc.CallOption) (*AddLocationResponse, error)
	// ChangeLocation changes backup location.
	ChangeLocation(ctx context.Context, in *ChangeLocationRequest, opts ...grpc.CallOption) (*ChangeLocationResponse, error)
	// RemoveLocation removes existing backup location.
	RemoveLocation(ctx context.Context, in *RemoveLocationRequest, opts ...grpc.CallOption) (*RemoveLocationResponse, error)
	// TestLocationConfig tests backup location and credentials.
	TestLocationConfig(ctx context.Context, in *TestLocationConfigRequest, opts ...grpc.CallOption) (*TestLocationConfigResponse, error)
}

type locationsClient struct {
	cc grpc.ClientConnInterface
}

func NewLocationsClient(cc grpc.ClientConnInterface) LocationsClient {
	return &locationsClient{cc}
}

func (c *locationsClient) ListLocations(ctx context.Context, in *ListLocationsRequest, opts ...grpc.CallOption) (*ListLocationsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListLocationsResponse)
	err := c.cc.Invoke(ctx, Locations_ListLocations_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *locationsClient) AddLocation(ctx context.Context, in *AddLocationRequest, opts ...grpc.CallOption) (*AddLocationResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddLocationResponse)
	err := c.cc.Invoke(ctx, Locations_AddLocation_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *locationsClient) ChangeLocation(ctx context.Context, in *ChangeLocationRequest, opts ...grpc.CallOption) (*ChangeLocationResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ChangeLocationResponse)
	err := c.cc.Invoke(ctx, Locations_ChangeLocation_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *locationsClient) RemoveLocation(ctx context.Context, in *RemoveLocationRequest, opts ...grpc.CallOption) (*RemoveLocationResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RemoveLocationResponse)
	err := c.cc.Invoke(ctx, Locations_RemoveLocation_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *locationsClient) TestLocationConfig(ctx context.Context, in *TestLocationConfigRequest, opts ...grpc.CallOption) (*TestLocationConfigResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TestLocationConfigResponse)
	err := c.cc.Invoke(ctx, Locations_TestLocationConfig_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LocationsServer is the server API for Locations service.
// All implementations must embed UnimplementedLocationsServer
// for forward compatibility
//
// Locations service provides access to Backup Locations.
type LocationsServer interface {
	// ListLocations returns a list of all backup locations.
	ListLocations(context.Context, *ListLocationsRequest) (*ListLocationsResponse, error)
	// AddLocation adds backup location.
	AddLocation(context.Context, *AddLocationRequest) (*AddLocationResponse, error)
	// ChangeLocation changes backup location.
	ChangeLocation(context.Context, *ChangeLocationRequest) (*ChangeLocationResponse, error)
	// RemoveLocation removes existing backup location.
	RemoveLocation(context.Context, *RemoveLocationRequest) (*RemoveLocationResponse, error)
	// TestLocationConfig tests backup location and credentials.
	TestLocationConfig(context.Context, *TestLocationConfigRequest) (*TestLocationConfigResponse, error)
	mustEmbedUnimplementedLocationsServer()
}

// UnimplementedLocationsServer must be embedded to have forward compatible implementations.
type UnimplementedLocationsServer struct{}

func (UnimplementedLocationsServer) ListLocations(context.Context, *ListLocationsRequest) (*ListLocationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListLocations not implemented")
}

func (UnimplementedLocationsServer) AddLocation(context.Context, *AddLocationRequest) (*AddLocationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddLocation not implemented")
}

func (UnimplementedLocationsServer) ChangeLocation(context.Context, *ChangeLocationRequest) (*ChangeLocationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeLocation not implemented")
}

func (UnimplementedLocationsServer) RemoveLocation(context.Context, *RemoveLocationRequest) (*RemoveLocationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveLocation not implemented")
}

func (UnimplementedLocationsServer) TestLocationConfig(context.Context, *TestLocationConfigRequest) (*TestLocationConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TestLocationConfig not implemented")
}
func (UnimplementedLocationsServer) mustEmbedUnimplementedLocationsServer() {}

// UnsafeLocationsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LocationsServer will
// result in compilation errors.
type UnsafeLocationsServer interface {
	mustEmbedUnimplementedLocationsServer()
}

func RegisterLocationsServer(s grpc.ServiceRegistrar, srv LocationsServer) {
	s.RegisterService(&Locations_ServiceDesc, srv)
}

func _Locations_ListLocations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListLocationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LocationsServer).ListLocations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Locations_ListLocations_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LocationsServer).ListLocations(ctx, req.(*ListLocationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Locations_AddLocation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddLocationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LocationsServer).AddLocation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Locations_AddLocation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LocationsServer).AddLocation(ctx, req.(*AddLocationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Locations_ChangeLocation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangeLocationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LocationsServer).ChangeLocation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Locations_ChangeLocation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LocationsServer).ChangeLocation(ctx, req.(*ChangeLocationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Locations_RemoveLocation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveLocationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LocationsServer).RemoveLocation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Locations_RemoveLocation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LocationsServer).RemoveLocation(ctx, req.(*RemoveLocationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Locations_TestLocationConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TestLocationConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LocationsServer).TestLocationConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Locations_TestLocationConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LocationsServer).TestLocationConfig(ctx, req.(*TestLocationConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Locations_ServiceDesc is the grpc.ServiceDesc for Locations service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Locations_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "backup.v1.Locations",
	HandlerType: (*LocationsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListLocations",
			Handler:    _Locations_ListLocations_Handler,
		},
		{
			MethodName: "AddLocation",
			Handler:    _Locations_AddLocation_Handler,
		},
		{
			MethodName: "ChangeLocation",
			Handler:    _Locations_ChangeLocation_Handler,
		},
		{
			MethodName: "RemoveLocation",
			Handler:    _Locations_RemoveLocation_Handler,
		},
		{
			MethodName: "TestLocationConfig",
			Handler:    _Locations_TestLocationConfig_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "managementpb/backup/locations.proto",
}
