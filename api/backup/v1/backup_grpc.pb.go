// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: backup/v1/backup.proto

package backupv1

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
	BackupService_StartBackup_FullMethodName                    = "/backup.v1.BackupService/StartBackup"
	BackupService_ListArtifactCompatibleServices_FullMethodName = "/backup.v1.BackupService/ListArtifactCompatibleServices"
	BackupService_ScheduleBackup_FullMethodName                 = "/backup.v1.BackupService/ScheduleBackup"
	BackupService_ListScheduledBackups_FullMethodName           = "/backup.v1.BackupService/ListScheduledBackups"
	BackupService_ChangeScheduledBackup_FullMethodName          = "/backup.v1.BackupService/ChangeScheduledBackup"
	BackupService_RemoveScheduledBackup_FullMethodName          = "/backup.v1.BackupService/RemoveScheduledBackup"
	BackupService_GetLogs_FullMethodName                        = "/backup.v1.BackupService/GetLogs"
	BackupService_ListArtifacts_FullMethodName                  = "/backup.v1.BackupService/ListArtifacts"
	BackupService_DeleteArtifact_FullMethodName                 = "/backup.v1.BackupService/DeleteArtifact"
	BackupService_ListPitrTimeranges_FullMethodName             = "/backup.v1.BackupService/ListPitrTimeranges"
)

// BackupServiceClient is the client API for BackupService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BackupServiceClient interface {
	// StartBackup request backup specified service to location.
	StartBackup(ctx context.Context, in *StartBackupRequest, opts ...grpc.CallOption) (*StartBackupResponse, error)
	// ListArtifactCompatibleServices lists compatible services for restoring a backup.
	ListArtifactCompatibleServices(ctx context.Context, in *ListArtifactCompatibleServicesRequest, opts ...grpc.CallOption) (*ListArtifactCompatibleServicesResponse, error)
	// ScheduleBackup schedules repeated backup.
	ScheduleBackup(ctx context.Context, in *ScheduleBackupRequest, opts ...grpc.CallOption) (*ScheduleBackupResponse, error)
	// ListScheduledBackups returns all scheduled backups.
	ListScheduledBackups(ctx context.Context, in *ListScheduledBackupsRequest, opts ...grpc.CallOption) (*ListScheduledBackupsResponse, error)
	// ChangeScheduledBackup changes existing scheduled backup.
	ChangeScheduledBackup(ctx context.Context, in *ChangeScheduledBackupRequest, opts ...grpc.CallOption) (*ChangeScheduledBackupResponse, error)
	// RemoveScheduledBackup removes existing scheduled backup.
	RemoveScheduledBackup(ctx context.Context, in *RemoveScheduledBackupRequest, opts ...grpc.CallOption) (*RemoveScheduledBackupResponse, error)
	// GetLogs returns logs from the underlying tools for a backup/restore job.
	GetLogs(ctx context.Context, in *GetLogsRequest, opts ...grpc.CallOption) (*GetLogsResponse, error)
	// ListArtifacts returns a list of all backup artifacts.
	ListArtifacts(ctx context.Context, in *ListArtifactsRequest, opts ...grpc.CallOption) (*ListArtifactsResponse, error)
	// DeleteArtifact deletes specified artifact.
	DeleteArtifact(ctx context.Context, in *DeleteArtifactRequest, opts ...grpc.CallOption) (*DeleteArtifactResponse, error)
	// ListPitrTimeranges list the available MongoDB PITR timeranges in a given backup location
	ListPitrTimeranges(ctx context.Context, in *ListPitrTimerangesRequest, opts ...grpc.CallOption) (*ListPitrTimerangesResponse, error)
}

type backupServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBackupServiceClient(cc grpc.ClientConnInterface) BackupServiceClient {
	return &backupServiceClient{cc}
}

func (c *backupServiceClient) StartBackup(ctx context.Context, in *StartBackupRequest, opts ...grpc.CallOption) (*StartBackupResponse, error) {
	out := new(StartBackupResponse)
	err := c.cc.Invoke(ctx, BackupService_StartBackup_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *backupServiceClient) ListArtifactCompatibleServices(ctx context.Context, in *ListArtifactCompatibleServicesRequest, opts ...grpc.CallOption) (*ListArtifactCompatibleServicesResponse, error) {
	out := new(ListArtifactCompatibleServicesResponse)
	err := c.cc.Invoke(ctx, BackupService_ListArtifactCompatibleServices_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *backupServiceClient) ScheduleBackup(ctx context.Context, in *ScheduleBackupRequest, opts ...grpc.CallOption) (*ScheduleBackupResponse, error) {
	out := new(ScheduleBackupResponse)
	err := c.cc.Invoke(ctx, BackupService_ScheduleBackup_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *backupServiceClient) ListScheduledBackups(ctx context.Context, in *ListScheduledBackupsRequest, opts ...grpc.CallOption) (*ListScheduledBackupsResponse, error) {
	out := new(ListScheduledBackupsResponse)
	err := c.cc.Invoke(ctx, BackupService_ListScheduledBackups_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *backupServiceClient) ChangeScheduledBackup(ctx context.Context, in *ChangeScheduledBackupRequest, opts ...grpc.CallOption) (*ChangeScheduledBackupResponse, error) {
	out := new(ChangeScheduledBackupResponse)
	err := c.cc.Invoke(ctx, BackupService_ChangeScheduledBackup_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *backupServiceClient) RemoveScheduledBackup(ctx context.Context, in *RemoveScheduledBackupRequest, opts ...grpc.CallOption) (*RemoveScheduledBackupResponse, error) {
	out := new(RemoveScheduledBackupResponse)
	err := c.cc.Invoke(ctx, BackupService_RemoveScheduledBackup_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *backupServiceClient) GetLogs(ctx context.Context, in *GetLogsRequest, opts ...grpc.CallOption) (*GetLogsResponse, error) {
	out := new(GetLogsResponse)
	err := c.cc.Invoke(ctx, BackupService_GetLogs_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *backupServiceClient) ListArtifacts(ctx context.Context, in *ListArtifactsRequest, opts ...grpc.CallOption) (*ListArtifactsResponse, error) {
	out := new(ListArtifactsResponse)
	err := c.cc.Invoke(ctx, BackupService_ListArtifacts_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *backupServiceClient) DeleteArtifact(ctx context.Context, in *DeleteArtifactRequest, opts ...grpc.CallOption) (*DeleteArtifactResponse, error) {
	out := new(DeleteArtifactResponse)
	err := c.cc.Invoke(ctx, BackupService_DeleteArtifact_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *backupServiceClient) ListPitrTimeranges(ctx context.Context, in *ListPitrTimerangesRequest, opts ...grpc.CallOption) (*ListPitrTimerangesResponse, error) {
	out := new(ListPitrTimerangesResponse)
	err := c.cc.Invoke(ctx, BackupService_ListPitrTimeranges_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BackupServiceServer is the server API for BackupService service.
// All implementations must embed UnimplementedBackupServiceServer
// for forward compatibility
type BackupServiceServer interface {
	// StartBackup request backup specified service to location.
	StartBackup(context.Context, *StartBackupRequest) (*StartBackupResponse, error)
	// ListArtifactCompatibleServices lists compatible services for restoring a backup.
	ListArtifactCompatibleServices(context.Context, *ListArtifactCompatibleServicesRequest) (*ListArtifactCompatibleServicesResponse, error)
	// ScheduleBackup schedules repeated backup.
	ScheduleBackup(context.Context, *ScheduleBackupRequest) (*ScheduleBackupResponse, error)
	// ListScheduledBackups returns all scheduled backups.
	ListScheduledBackups(context.Context, *ListScheduledBackupsRequest) (*ListScheduledBackupsResponse, error)
	// ChangeScheduledBackup changes existing scheduled backup.
	ChangeScheduledBackup(context.Context, *ChangeScheduledBackupRequest) (*ChangeScheduledBackupResponse, error)
	// RemoveScheduledBackup removes existing scheduled backup.
	RemoveScheduledBackup(context.Context, *RemoveScheduledBackupRequest) (*RemoveScheduledBackupResponse, error)
	// GetLogs returns logs from the underlying tools for a backup/restore job.
	GetLogs(context.Context, *GetLogsRequest) (*GetLogsResponse, error)
	// ListArtifacts returns a list of all backup artifacts.
	ListArtifacts(context.Context, *ListArtifactsRequest) (*ListArtifactsResponse, error)
	// DeleteArtifact deletes specified artifact.
	DeleteArtifact(context.Context, *DeleteArtifactRequest) (*DeleteArtifactResponse, error)
	// ListPitrTimeranges list the available MongoDB PITR timeranges in a given backup location
	ListPitrTimeranges(context.Context, *ListPitrTimerangesRequest) (*ListPitrTimerangesResponse, error)
	mustEmbedUnimplementedBackupServiceServer()
}

// UnimplementedBackupServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBackupServiceServer struct{}

func (UnimplementedBackupServiceServer) StartBackup(context.Context, *StartBackupRequest) (*StartBackupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartBackup not implemented")
}

func (UnimplementedBackupServiceServer) ListArtifactCompatibleServices(context.Context, *ListArtifactCompatibleServicesRequest) (*ListArtifactCompatibleServicesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListArtifactCompatibleServices not implemented")
}

func (UnimplementedBackupServiceServer) ScheduleBackup(context.Context, *ScheduleBackupRequest) (*ScheduleBackupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ScheduleBackup not implemented")
}

func (UnimplementedBackupServiceServer) ListScheduledBackups(context.Context, *ListScheduledBackupsRequest) (*ListScheduledBackupsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListScheduledBackups not implemented")
}

func (UnimplementedBackupServiceServer) ChangeScheduledBackup(context.Context, *ChangeScheduledBackupRequest) (*ChangeScheduledBackupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeScheduledBackup not implemented")
}

func (UnimplementedBackupServiceServer) RemoveScheduledBackup(context.Context, *RemoveScheduledBackupRequest) (*RemoveScheduledBackupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveScheduledBackup not implemented")
}

func (UnimplementedBackupServiceServer) GetLogs(context.Context, *GetLogsRequest) (*GetLogsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLogs not implemented")
}

func (UnimplementedBackupServiceServer) ListArtifacts(context.Context, *ListArtifactsRequest) (*ListArtifactsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListArtifacts not implemented")
}

func (UnimplementedBackupServiceServer) DeleteArtifact(context.Context, *DeleteArtifactRequest) (*DeleteArtifactResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteArtifact not implemented")
}

func (UnimplementedBackupServiceServer) ListPitrTimeranges(context.Context, *ListPitrTimerangesRequest) (*ListPitrTimerangesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPitrTimeranges not implemented")
}
func (UnimplementedBackupServiceServer) mustEmbedUnimplementedBackupServiceServer() {}

// UnsafeBackupServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BackupServiceServer will
// result in compilation errors.
type UnsafeBackupServiceServer interface {
	mustEmbedUnimplementedBackupServiceServer()
}

func RegisterBackupServiceServer(s grpc.ServiceRegistrar, srv BackupServiceServer) {
	s.RegisterService(&BackupService_ServiceDesc, srv)
}

func _BackupService_StartBackup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartBackupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackupServiceServer).StartBackup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BackupService_StartBackup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackupServiceServer).StartBackup(ctx, req.(*StartBackupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BackupService_ListArtifactCompatibleServices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListArtifactCompatibleServicesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackupServiceServer).ListArtifactCompatibleServices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BackupService_ListArtifactCompatibleServices_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackupServiceServer).ListArtifactCompatibleServices(ctx, req.(*ListArtifactCompatibleServicesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BackupService_ScheduleBackup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ScheduleBackupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackupServiceServer).ScheduleBackup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BackupService_ScheduleBackup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackupServiceServer).ScheduleBackup(ctx, req.(*ScheduleBackupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BackupService_ListScheduledBackups_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListScheduledBackupsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackupServiceServer).ListScheduledBackups(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BackupService_ListScheduledBackups_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackupServiceServer).ListScheduledBackups(ctx, req.(*ListScheduledBackupsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BackupService_ChangeScheduledBackup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangeScheduledBackupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackupServiceServer).ChangeScheduledBackup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BackupService_ChangeScheduledBackup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackupServiceServer).ChangeScheduledBackup(ctx, req.(*ChangeScheduledBackupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BackupService_RemoveScheduledBackup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveScheduledBackupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackupServiceServer).RemoveScheduledBackup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BackupService_RemoveScheduledBackup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackupServiceServer).RemoveScheduledBackup(ctx, req.(*RemoveScheduledBackupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BackupService_GetLogs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLogsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackupServiceServer).GetLogs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BackupService_GetLogs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackupServiceServer).GetLogs(ctx, req.(*GetLogsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BackupService_ListArtifacts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListArtifactsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackupServiceServer).ListArtifacts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BackupService_ListArtifacts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackupServiceServer).ListArtifacts(ctx, req.(*ListArtifactsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BackupService_DeleteArtifact_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteArtifactRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackupServiceServer).DeleteArtifact(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BackupService_DeleteArtifact_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackupServiceServer).DeleteArtifact(ctx, req.(*DeleteArtifactRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BackupService_ListPitrTimeranges_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPitrTimerangesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackupServiceServer).ListPitrTimeranges(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BackupService_ListPitrTimeranges_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackupServiceServer).ListPitrTimeranges(ctx, req.(*ListPitrTimerangesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BackupService_ServiceDesc is the grpc.ServiceDesc for BackupService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BackupService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "backup.v1.BackupService",
	HandlerType: (*BackupServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StartBackup",
			Handler:    _BackupService_StartBackup_Handler,
		},
		{
			MethodName: "ListArtifactCompatibleServices",
			Handler:    _BackupService_ListArtifactCompatibleServices_Handler,
		},
		{
			MethodName: "ScheduleBackup",
			Handler:    _BackupService_ScheduleBackup_Handler,
		},
		{
			MethodName: "ListScheduledBackups",
			Handler:    _BackupService_ListScheduledBackups_Handler,
		},
		{
			MethodName: "ChangeScheduledBackup",
			Handler:    _BackupService_ChangeScheduledBackup_Handler,
		},
		{
			MethodName: "RemoveScheduledBackup",
			Handler:    _BackupService_RemoveScheduledBackup_Handler,
		},
		{
			MethodName: "GetLogs",
			Handler:    _BackupService_GetLogs_Handler,
		},
		{
			MethodName: "ListArtifacts",
			Handler:    _BackupService_ListArtifacts_Handler,
		},
		{
			MethodName: "DeleteArtifact",
			Handler:    _BackupService_DeleteArtifact_Handler,
		},
		{
			MethodName: "ListPitrTimeranges",
			Handler:    _BackupService_ListPitrTimeranges_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "backup/v1/backup.proto",
}
