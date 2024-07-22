// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             (unknown)
// source: managementpb/dbaas/kubernetes.proto

package dbaasv1beta1

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
	Kubernetes_ListKubernetesClusters_FullMethodName      = "/dbaas.v1beta1.Kubernetes/ListKubernetesClusters"
	Kubernetes_RegisterKubernetesCluster_FullMethodName   = "/dbaas.v1beta1.Kubernetes/RegisterKubernetesCluster"
	Kubernetes_UnregisterKubernetesCluster_FullMethodName = "/dbaas.v1beta1.Kubernetes/UnregisterKubernetesCluster"
	Kubernetes_GetKubernetesCluster_FullMethodName        = "/dbaas.v1beta1.Kubernetes/GetKubernetesCluster"
	Kubernetes_GetResources_FullMethodName                = "/dbaas.v1beta1.Kubernetes/GetResources"
	Kubernetes_ListStorageClasses_FullMethodName          = "/dbaas.v1beta1.Kubernetes/ListStorageClasses"
)

// KubernetesClient is the client API for Kubernetes service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Kubernetes service provides public methods for managing Kubernetes clusters.
type KubernetesClient interface {
	// ListKubernetesClusters returns a list of all registered Kubernetes clusters.
	ListKubernetesClusters(ctx context.Context, in *ListKubernetesClustersRequest, opts ...grpc.CallOption) (*ListKubernetesClustersResponse, error)
	// RegisterKubernetesCluster registers an existing Kubernetes cluster in PMM.
	RegisterKubernetesCluster(ctx context.Context, in *RegisterKubernetesClusterRequest, opts ...grpc.CallOption) (*RegisterKubernetesClusterResponse, error)
	// UnregisterKubernetesCluster removes a registered Kubernetes cluster from PMM.
	UnregisterKubernetesCluster(ctx context.Context, in *UnregisterKubernetesClusterRequest, opts ...grpc.CallOption) (*UnregisterKubernetesClusterResponse, error)
	// GetKubernetesCluster return KubeAuth with Kubernetes config.
	GetKubernetesCluster(ctx context.Context, in *GetKubernetesClusterRequest, opts ...grpc.CallOption) (*GetKubernetesClusterResponse, error)
	// GetResources returns all and available resources of a Kubernetes cluster.
	// NOTE: The user defined in kubeconfig for the cluster has to have rights to
	//
	//	list and get Pods from all Namespaces. Also getting and listing Nodes
	//	has to be allowed.
	GetResources(ctx context.Context, in *GetResourcesRequest, opts ...grpc.CallOption) (*GetResourcesResponse, error)
	// ListStorageClasses returns the names of all storage classes available in a Kubernetes cluster.
	ListStorageClasses(ctx context.Context, in *ListStorageClassesRequest, opts ...grpc.CallOption) (*ListStorageClassesResponse, error)
}

type kubernetesClient struct {
	cc grpc.ClientConnInterface
}

func NewKubernetesClient(cc grpc.ClientConnInterface) KubernetesClient {
	return &kubernetesClient{cc}
}

func (c *kubernetesClient) ListKubernetesClusters(ctx context.Context, in *ListKubernetesClustersRequest, opts ...grpc.CallOption) (*ListKubernetesClustersResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListKubernetesClustersResponse)
	err := c.cc.Invoke(ctx, Kubernetes_ListKubernetesClusters_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kubernetesClient) RegisterKubernetesCluster(ctx context.Context, in *RegisterKubernetesClusterRequest, opts ...grpc.CallOption) (*RegisterKubernetesClusterResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RegisterKubernetesClusterResponse)
	err := c.cc.Invoke(ctx, Kubernetes_RegisterKubernetesCluster_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kubernetesClient) UnregisterKubernetesCluster(ctx context.Context, in *UnregisterKubernetesClusterRequest, opts ...grpc.CallOption) (*UnregisterKubernetesClusterResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UnregisterKubernetesClusterResponse)
	err := c.cc.Invoke(ctx, Kubernetes_UnregisterKubernetesCluster_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kubernetesClient) GetKubernetesCluster(ctx context.Context, in *GetKubernetesClusterRequest, opts ...grpc.CallOption) (*GetKubernetesClusterResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetKubernetesClusterResponse)
	err := c.cc.Invoke(ctx, Kubernetes_GetKubernetesCluster_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kubernetesClient) GetResources(ctx context.Context, in *GetResourcesRequest, opts ...grpc.CallOption) (*GetResourcesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetResourcesResponse)
	err := c.cc.Invoke(ctx, Kubernetes_GetResources_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kubernetesClient) ListStorageClasses(ctx context.Context, in *ListStorageClassesRequest, opts ...grpc.CallOption) (*ListStorageClassesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListStorageClassesResponse)
	err := c.cc.Invoke(ctx, Kubernetes_ListStorageClasses_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KubernetesServer is the server API for Kubernetes service.
// All implementations must embed UnimplementedKubernetesServer
// for forward compatibility
//
// Kubernetes service provides public methods for managing Kubernetes clusters.
type KubernetesServer interface {
	// ListKubernetesClusters returns a list of all registered Kubernetes clusters.
	ListKubernetesClusters(context.Context, *ListKubernetesClustersRequest) (*ListKubernetesClustersResponse, error)
	// RegisterKubernetesCluster registers an existing Kubernetes cluster in PMM.
	RegisterKubernetesCluster(context.Context, *RegisterKubernetesClusterRequest) (*RegisterKubernetesClusterResponse, error)
	// UnregisterKubernetesCluster removes a registered Kubernetes cluster from PMM.
	UnregisterKubernetesCluster(context.Context, *UnregisterKubernetesClusterRequest) (*UnregisterKubernetesClusterResponse, error)
	// GetKubernetesCluster return KubeAuth with Kubernetes config.
	GetKubernetesCluster(context.Context, *GetKubernetesClusterRequest) (*GetKubernetesClusterResponse, error)
	// GetResources returns all and available resources of a Kubernetes cluster.
	// NOTE: The user defined in kubeconfig for the cluster has to have rights to
	//
	//	list and get Pods from all Namespaces. Also getting and listing Nodes
	//	has to be allowed.
	GetResources(context.Context, *GetResourcesRequest) (*GetResourcesResponse, error)
	// ListStorageClasses returns the names of all storage classes available in a Kubernetes cluster.
	ListStorageClasses(context.Context, *ListStorageClassesRequest) (*ListStorageClassesResponse, error)
	mustEmbedUnimplementedKubernetesServer()
}

// UnimplementedKubernetesServer must be embedded to have forward compatible implementations.
type UnimplementedKubernetesServer struct{}

func (UnimplementedKubernetesServer) ListKubernetesClusters(context.Context, *ListKubernetesClustersRequest) (*ListKubernetesClustersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListKubernetesClusters not implemented")
}

func (UnimplementedKubernetesServer) RegisterKubernetesCluster(context.Context, *RegisterKubernetesClusterRequest) (*RegisterKubernetesClusterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterKubernetesCluster not implemented")
}

func (UnimplementedKubernetesServer) UnregisterKubernetesCluster(context.Context, *UnregisterKubernetesClusterRequest) (*UnregisterKubernetesClusterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnregisterKubernetesCluster not implemented")
}

func (UnimplementedKubernetesServer) GetKubernetesCluster(context.Context, *GetKubernetesClusterRequest) (*GetKubernetesClusterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetKubernetesCluster not implemented")
}

func (UnimplementedKubernetesServer) GetResources(context.Context, *GetResourcesRequest) (*GetResourcesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetResources not implemented")
}

func (UnimplementedKubernetesServer) ListStorageClasses(context.Context, *ListStorageClassesRequest) (*ListStorageClassesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListStorageClasses not implemented")
}
func (UnimplementedKubernetesServer) mustEmbedUnimplementedKubernetesServer() {}

// UnsafeKubernetesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to KubernetesServer will
// result in compilation errors.
type UnsafeKubernetesServer interface {
	mustEmbedUnimplementedKubernetesServer()
}

func RegisterKubernetesServer(s grpc.ServiceRegistrar, srv KubernetesServer) {
	s.RegisterService(&Kubernetes_ServiceDesc, srv)
}

func _Kubernetes_ListKubernetesClusters_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListKubernetesClustersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KubernetesServer).ListKubernetesClusters(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Kubernetes_ListKubernetesClusters_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KubernetesServer).ListKubernetesClusters(ctx, req.(*ListKubernetesClustersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Kubernetes_RegisterKubernetesCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterKubernetesClusterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KubernetesServer).RegisterKubernetesCluster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Kubernetes_RegisterKubernetesCluster_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KubernetesServer).RegisterKubernetesCluster(ctx, req.(*RegisterKubernetesClusterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Kubernetes_UnregisterKubernetesCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnregisterKubernetesClusterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KubernetesServer).UnregisterKubernetesCluster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Kubernetes_UnregisterKubernetesCluster_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KubernetesServer).UnregisterKubernetesCluster(ctx, req.(*UnregisterKubernetesClusterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Kubernetes_GetKubernetesCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetKubernetesClusterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KubernetesServer).GetKubernetesCluster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Kubernetes_GetKubernetesCluster_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KubernetesServer).GetKubernetesCluster(ctx, req.(*GetKubernetesClusterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Kubernetes_GetResources_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetResourcesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KubernetesServer).GetResources(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Kubernetes_GetResources_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KubernetesServer).GetResources(ctx, req.(*GetResourcesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Kubernetes_ListStorageClasses_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListStorageClassesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KubernetesServer).ListStorageClasses(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Kubernetes_ListStorageClasses_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KubernetesServer).ListStorageClasses(ctx, req.(*ListStorageClassesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Kubernetes_ServiceDesc is the grpc.ServiceDesc for Kubernetes service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Kubernetes_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "dbaas.v1beta1.Kubernetes",
	HandlerType: (*KubernetesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListKubernetesClusters",
			Handler:    _Kubernetes_ListKubernetesClusters_Handler,
		},
		{
			MethodName: "RegisterKubernetesCluster",
			Handler:    _Kubernetes_RegisterKubernetesCluster_Handler,
		},
		{
			MethodName: "UnregisterKubernetesCluster",
			Handler:    _Kubernetes_UnregisterKubernetesCluster_Handler,
		},
		{
			MethodName: "GetKubernetesCluster",
			Handler:    _Kubernetes_GetKubernetesCluster_Handler,
		},
		{
			MethodName: "GetResources",
			Handler:    _Kubernetes_GetResources_Handler,
		},
		{
			MethodName: "ListStorageClasses",
			Handler:    _Kubernetes_ListStorageClasses_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "managementpb/dbaas/kubernetes.proto",
}
