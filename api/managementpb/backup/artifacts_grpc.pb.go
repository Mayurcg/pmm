// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: managementpb/backup/artifacts.proto

package backupv1

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
	Artifacts_ListArtifacts_FullMethodName      = "/backup.v1.Artifacts/ListArtifacts"
	Artifacts_DeleteArtifact_FullMethodName     = "/backup.v1.Artifacts/DeleteArtifact"
	Artifacts_ListPitrTimeranges_FullMethodName = "/backup.v1.Artifacts/ListPitrTimeranges"
)

// ArtifactsClient is the client API for Artifacts service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Artifacts service provides public methods for managing backup artifacts.
type ArtifactsClient interface {
	// ListArtifacts returns a list of all backup artifacts.
	ListArtifacts(ctx context.Context, in *ListArtifactsRequest, opts ...grpc.CallOption) (*ListArtifactsResponse, error)
	// DeleteArtifact deletes specified artifact.
	DeleteArtifact(ctx context.Context, in *DeleteArtifactRequest, opts ...grpc.CallOption) (*DeleteArtifactResponse, error)
	// ListPitrTimeranges list the available MongoDB PITR timeranges in a given backup location
	ListPitrTimeranges(ctx context.Context, in *ListPitrTimerangesRequest, opts ...grpc.CallOption) (*ListPitrTimerangesResponse, error)
}

type artifactsClient struct {
	cc grpc.ClientConnInterface
}

func NewArtifactsClient(cc grpc.ClientConnInterface) ArtifactsClient {
	return &artifactsClient{cc}
}

func (c *artifactsClient) ListArtifacts(ctx context.Context, in *ListArtifactsRequest, opts ...grpc.CallOption) (*ListArtifactsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListArtifactsResponse)
	err := c.cc.Invoke(ctx, Artifacts_ListArtifacts_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *artifactsClient) DeleteArtifact(ctx context.Context, in *DeleteArtifactRequest, opts ...grpc.CallOption) (*DeleteArtifactResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteArtifactResponse)
	err := c.cc.Invoke(ctx, Artifacts_DeleteArtifact_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *artifactsClient) ListPitrTimeranges(ctx context.Context, in *ListPitrTimerangesRequest, opts ...grpc.CallOption) (*ListPitrTimerangesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListPitrTimerangesResponse)
	err := c.cc.Invoke(ctx, Artifacts_ListPitrTimeranges_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ArtifactsServer is the server API for Artifacts service.
// All implementations must embed UnimplementedArtifactsServer
// for forward compatibility.
//
// Artifacts service provides public methods for managing backup artifacts.
type ArtifactsServer interface {
	// ListArtifacts returns a list of all backup artifacts.
	ListArtifacts(context.Context, *ListArtifactsRequest) (*ListArtifactsResponse, error)
	// DeleteArtifact deletes specified artifact.
	DeleteArtifact(context.Context, *DeleteArtifactRequest) (*DeleteArtifactResponse, error)
	// ListPitrTimeranges list the available MongoDB PITR timeranges in a given backup location
	ListPitrTimeranges(context.Context, *ListPitrTimerangesRequest) (*ListPitrTimerangesResponse, error)
	mustEmbedUnimplementedArtifactsServer()
}

// UnimplementedArtifactsServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedArtifactsServer struct{}

func (UnimplementedArtifactsServer) ListArtifacts(context.Context, *ListArtifactsRequest) (*ListArtifactsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListArtifacts not implemented")
}

func (UnimplementedArtifactsServer) DeleteArtifact(context.Context, *DeleteArtifactRequest) (*DeleteArtifactResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteArtifact not implemented")
}

func (UnimplementedArtifactsServer) ListPitrTimeranges(context.Context, *ListPitrTimerangesRequest) (*ListPitrTimerangesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPitrTimeranges not implemented")
}
func (UnimplementedArtifactsServer) mustEmbedUnimplementedArtifactsServer() {}
func (UnimplementedArtifactsServer) testEmbeddedByValue()                   {}

// UnsafeArtifactsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ArtifactsServer will
// result in compilation errors.
type UnsafeArtifactsServer interface {
	mustEmbedUnimplementedArtifactsServer()
}

func RegisterArtifactsServer(s grpc.ServiceRegistrar, srv ArtifactsServer) {
	// If the following call pancis, it indicates UnimplementedArtifactsServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Artifacts_ServiceDesc, srv)
}

func _Artifacts_ListArtifacts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListArtifactsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArtifactsServer).ListArtifacts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Artifacts_ListArtifacts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArtifactsServer).ListArtifacts(ctx, req.(*ListArtifactsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Artifacts_DeleteArtifact_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteArtifactRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArtifactsServer).DeleteArtifact(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Artifacts_DeleteArtifact_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArtifactsServer).DeleteArtifact(ctx, req.(*DeleteArtifactRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Artifacts_ListPitrTimeranges_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPitrTimerangesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArtifactsServer).ListPitrTimeranges(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Artifacts_ListPitrTimeranges_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArtifactsServer).ListPitrTimeranges(ctx, req.(*ListPitrTimerangesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Artifacts_ServiceDesc is the grpc.ServiceDesc for Artifacts service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Artifacts_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "backup.v1.Artifacts",
	HandlerType: (*ArtifactsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListArtifacts",
			Handler:    _Artifacts_ListArtifacts_Handler,
		},
		{
			MethodName: "DeleteArtifact",
			Handler:    _Artifacts_DeleteArtifact_Handler,
		},
		{
			MethodName: "ListPitrTimeranges",
			Handler:    _Artifacts_ListPitrTimeranges_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "managementpb/backup/artifacts.proto",
}
