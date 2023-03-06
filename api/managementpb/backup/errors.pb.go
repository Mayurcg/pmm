// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1-devel
// 	protoc        (unknown)
// source: managementpb/backup/errors.proto

package backupv1

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// ErrorCode is a set of specific errors that are not present in the standard set of errors
// and returned in the details field of the response.
type ErrorCode int32

const (
	ErrorCode_ERROR_CODE_INVALID ErrorCode = 0
	// ERROR_CODE_XTRABACKUP_NOT_INSTALLED is returned if some xtrabackup component is missing.
	ErrorCode_ERROR_CODE_XTRABACKUP_NOT_INSTALLED ErrorCode = 1
	// ERROR_CODE_INVALID_XTRABACKUP is returned if xtrabackup components have different version.
	ErrorCode_ERROR_CODE_INVALID_XTRABACKUP ErrorCode = 2
	// ERROR_CODE_INCOMPATIBLE_XTRABACKUP is returned if xtrabackup is not compatible with the MySQL.
	ErrorCode_ERROR_CODE_INCOMPATIBLE_XTRABACKUP ErrorCode = 3
	// ERROR_CODE_INCOMPATIBLE_TARGET_MYSQL is returned if target version of MySQL is not compatible for restoring selected artifact.
	ErrorCode_ERROR_CODE_INCOMPATIBLE_TARGET_MYSQL ErrorCode = 4
	// ERROR_CODE_INCOMPATIBLE_TARGET_MONGODB is returned if target version of MongoDB is not compatible for restoring selected artifact.
	ErrorCode_ERROR_CODE_INCOMPATIBLE_TARGET_MONGODB ErrorCode = 5
)

// Enum value maps for ErrorCode.
var (
	ErrorCode_name = map[int32]string{
		0: "ERROR_CODE_INVALID",
		1: "ERROR_CODE_XTRABACKUP_NOT_INSTALLED",
		2: "ERROR_CODE_INVALID_XTRABACKUP",
		3: "ERROR_CODE_INCOMPATIBLE_XTRABACKUP",
		4: "ERROR_CODE_INCOMPATIBLE_TARGET_MYSQL",
		5: "ERROR_CODE_INCOMPATIBLE_TARGET_MONGODB",
	}
	ErrorCode_value = map[string]int32{
		"ERROR_CODE_INVALID":                     0,
		"ERROR_CODE_XTRABACKUP_NOT_INSTALLED":    1,
		"ERROR_CODE_INVALID_XTRABACKUP":          2,
		"ERROR_CODE_INCOMPATIBLE_XTRABACKUP":     3,
		"ERROR_CODE_INCOMPATIBLE_TARGET_MYSQL":   4,
		"ERROR_CODE_INCOMPATIBLE_TARGET_MONGODB": 5,
	}
)

func (x ErrorCode) Enum() *ErrorCode {
	p := new(ErrorCode)
	*p = x
	return p
}

func (x ErrorCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ErrorCode) Descriptor() protoreflect.EnumDescriptor {
	return file_managementpb_backup_errors_proto_enumTypes[0].Descriptor()
}

func (ErrorCode) Type() protoreflect.EnumType {
	return &file_managementpb_backup_errors_proto_enumTypes[0]
}

func (x ErrorCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ErrorCode.Descriptor instead.
func (ErrorCode) EnumDescriptor() ([]byte, []int) {
	return file_managementpb_backup_errors_proto_rawDescGZIP(), []int{0}
}

// Error represents error message returned in the details field in the response.
type Error struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code ErrorCode `protobuf:"varint,1,opt,name=code,proto3,enum=backup.v1.ErrorCode" json:"code,omitempty"`
}

func (x *Error) Reset() {
	*x = Error{}
	if protoimpl.UnsafeEnabled {
		mi := &file_managementpb_backup_errors_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Error) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Error) ProtoMessage() {}

func (x *Error) ProtoReflect() protoreflect.Message {
	mi := &file_managementpb_backup_errors_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Error.ProtoReflect.Descriptor instead.
func (*Error) Descriptor() ([]byte, []int) {
	return file_managementpb_backup_errors_proto_rawDescGZIP(), []int{0}
}

func (x *Error) GetCode() ErrorCode {
	if x != nil {
		return x.Code
	}
	return ErrorCode_ERROR_CODE_INVALID
}

var File_managementpb_backup_errors_proto protoreflect.FileDescriptor

var file_managementpb_backup_errors_proto_rawDesc = []byte{
	0x0a, 0x20, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x2f, 0x62,
	0x61, 0x63, 0x6b, 0x75, 0x70, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x09, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x2e, 0x76, 0x31, 0x22, 0x31, 0x0a,
	0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x28, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x2e, 0x76, 0x31,
	0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x2a, 0xed, 0x01, 0x0a, 0x09, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x16,
	0x0a, 0x12, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x49, 0x4e, 0x56,
	0x41, 0x4c, 0x49, 0x44, 0x10, 0x00, 0x12, 0x27, 0x0a, 0x23, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f,
	0x43, 0x4f, 0x44, 0x45, 0x5f, 0x58, 0x54, 0x52, 0x41, 0x42, 0x41, 0x43, 0x4b, 0x55, 0x50, 0x5f,
	0x4e, 0x4f, 0x54, 0x5f, 0x49, 0x4e, 0x53, 0x54, 0x41, 0x4c, 0x4c, 0x45, 0x44, 0x10, 0x01, 0x12,
	0x21, 0x0a, 0x1d, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x49, 0x4e,
	0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x58, 0x54, 0x52, 0x41, 0x42, 0x41, 0x43, 0x4b, 0x55, 0x50,
	0x10, 0x02, 0x12, 0x26, 0x0a, 0x22, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x43, 0x4f, 0x44, 0x45,
	0x5f, 0x49, 0x4e, 0x43, 0x4f, 0x4d, 0x50, 0x41, 0x54, 0x49, 0x42, 0x4c, 0x45, 0x5f, 0x58, 0x54,
	0x52, 0x41, 0x42, 0x41, 0x43, 0x4b, 0x55, 0x50, 0x10, 0x03, 0x12, 0x28, 0x0a, 0x24, 0x45, 0x52,
	0x52, 0x4f, 0x52, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x49, 0x4e, 0x43, 0x4f, 0x4d, 0x50, 0x41,
	0x54, 0x49, 0x42, 0x4c, 0x45, 0x5f, 0x54, 0x41, 0x52, 0x47, 0x45, 0x54, 0x5f, 0x4d, 0x59, 0x53,
	0x51, 0x4c, 0x10, 0x04, 0x12, 0x2a, 0x0a, 0x26, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x43, 0x4f,
	0x44, 0x45, 0x5f, 0x49, 0x4e, 0x43, 0x4f, 0x4d, 0x50, 0x41, 0x54, 0x49, 0x42, 0x4c, 0x45, 0x5f,
	0x54, 0x41, 0x52, 0x47, 0x45, 0x54, 0x5f, 0x4d, 0x4f, 0x4e, 0x47, 0x4f, 0x44, 0x42, 0x10, 0x05,
	0x42, 0x9a, 0x01, 0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x2e,
	0x76, 0x31, 0x42, 0x0b, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x37, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x65,
	0x72, 0x63, 0x6f, 0x6e, 0x61, 0x2f, 0x70, 0x6d, 0x6d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x75,
	0x70, 0x3b, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x42, 0x58, 0x58,
	0xaa, 0x02, 0x09, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x09, 0x42,
	0x61, 0x63, 0x6b, 0x75, 0x70, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x15, 0x42, 0x61, 0x63, 0x6b, 0x75,
	0x70, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0xea, 0x02, 0x0a, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_managementpb_backup_errors_proto_rawDescOnce sync.Once
	file_managementpb_backup_errors_proto_rawDescData = file_managementpb_backup_errors_proto_rawDesc
)

func file_managementpb_backup_errors_proto_rawDescGZIP() []byte {
	file_managementpb_backup_errors_proto_rawDescOnce.Do(func() {
		file_managementpb_backup_errors_proto_rawDescData = protoimpl.X.CompressGZIP(file_managementpb_backup_errors_proto_rawDescData)
	})
	return file_managementpb_backup_errors_proto_rawDescData
}

var (
	file_managementpb_backup_errors_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
	file_managementpb_backup_errors_proto_msgTypes  = make([]protoimpl.MessageInfo, 1)
	file_managementpb_backup_errors_proto_goTypes   = []interface{}{
		(ErrorCode)(0), // 0: backup.v1.ErrorCode
		(*Error)(nil),  // 1: backup.v1.Error
	}
)

var file_managementpb_backup_errors_proto_depIdxs = []int32{
	0, // 0: backup.v1.Error.code:type_name -> backup.v1.ErrorCode
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_managementpb_backup_errors_proto_init() }
func file_managementpb_backup_errors_proto_init() {
	if File_managementpb_backup_errors_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_managementpb_backup_errors_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Error); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_managementpb_backup_errors_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_managementpb_backup_errors_proto_goTypes,
		DependencyIndexes: file_managementpb_backup_errors_proto_depIdxs,
		EnumInfos:         file_managementpb_backup_errors_proto_enumTypes,
		MessageInfos:      file_managementpb_backup_errors_proto_msgTypes,
	}.Build()
	File_managementpb_backup_errors_proto = out.File
	file_managementpb_backup_errors_proto_rawDesc = nil
	file_managementpb_backup_errors_proto_goTypes = nil
	file_managementpb_backup_errors_proto_depIdxs = nil
}
