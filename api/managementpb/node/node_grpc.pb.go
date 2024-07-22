// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             (unknown)
// source: managementpb/node/node.proto

package nodev1beta1

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
	MgmtNode_ListNodes_FullMethodName = "/node.v1beta1.MgmtNode/ListNodes"
	MgmtNode_GetNode_FullMethodName   = "/node.v1beta1.MgmtNode/GetNode"
)

// MgmtNodeClient is the client API for MgmtNode service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// MgmtNode service provides public Management API methods for Nodes.
type MgmtNodeClient interface {
	// ListNode returns a list of nodes.
	ListNodes(ctx context.Context, in *ListNodeRequest, opts ...grpc.CallOption) (*ListNodeResponse, error)
	// GetNode returns a single Node by ID.
	GetNode(ctx context.Context, in *GetNodeRequest, opts ...grpc.CallOption) (*GetNodeResponse, error)
}

type mgmtNodeClient struct {
	cc grpc.ClientConnInterface
}

func NewMgmtNodeClient(cc grpc.ClientConnInterface) MgmtNodeClient {
	return &mgmtNodeClient{cc}
}

func (c *mgmtNodeClient) ListNodes(ctx context.Context, in *ListNodeRequest, opts ...grpc.CallOption) (*ListNodeResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListNodeResponse)
	err := c.cc.Invoke(ctx, MgmtNode_ListNodes_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mgmtNodeClient) GetNode(ctx context.Context, in *GetNodeRequest, opts ...grpc.CallOption) (*GetNodeResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetNodeResponse)
	err := c.cc.Invoke(ctx, MgmtNode_GetNode_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MgmtNodeServer is the server API for MgmtNode service.
// All implementations must embed UnimplementedMgmtNodeServer
// for forward compatibility
//
// MgmtNode service provides public Management API methods for Nodes.
type MgmtNodeServer interface {
	// ListNode returns a list of nodes.
	ListNodes(context.Context, *ListNodeRequest) (*ListNodeResponse, error)
	// GetNode returns a single Node by ID.
	GetNode(context.Context, *GetNodeRequest) (*GetNodeResponse, error)
	mustEmbedUnimplementedMgmtNodeServer()
}

// UnimplementedMgmtNodeServer must be embedded to have forward compatible implementations.
type UnimplementedMgmtNodeServer struct{}

func (UnimplementedMgmtNodeServer) ListNodes(context.Context, *ListNodeRequest) (*ListNodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListNodes not implemented")
}

func (UnimplementedMgmtNodeServer) GetNode(context.Context, *GetNodeRequest) (*GetNodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNode not implemented")
}
func (UnimplementedMgmtNodeServer) mustEmbedUnimplementedMgmtNodeServer() {}

// UnsafeMgmtNodeServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MgmtNodeServer will
// result in compilation errors.
type UnsafeMgmtNodeServer interface {
	mustEmbedUnimplementedMgmtNodeServer()
}

func RegisterMgmtNodeServer(s grpc.ServiceRegistrar, srv MgmtNodeServer) {
	s.RegisterService(&MgmtNode_ServiceDesc, srv)
}

func _MgmtNode_ListNodes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListNodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MgmtNodeServer).ListNodes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MgmtNode_ListNodes_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MgmtNodeServer).ListNodes(ctx, req.(*ListNodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MgmtNode_GetNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MgmtNodeServer).GetNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MgmtNode_GetNode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MgmtNodeServer).GetNode(ctx, req.(*GetNodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MgmtNode_ServiceDesc is the grpc.ServiceDesc for MgmtNode service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MgmtNode_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "node.v1beta1.MgmtNode",
	HandlerType: (*MgmtNodeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListNodes",
			Handler:    _MgmtNode_ListNodes_Handler,
		},
		{
			MethodName: "GetNode",
			Handler:    _MgmtNode_GetNode_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "managementpb/node/node.proto",
}
