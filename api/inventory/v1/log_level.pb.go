// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        (unknown)
// source: inventory/v1/log_level.proto

package inventoryv1

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

// Log level for exporters
type LogLevel int32

const (
	// Auto
	LogLevel_LOG_LEVEL_UNSPECIFIED LogLevel = 0
	LogLevel_LOG_LEVEL_FATAL       LogLevel = 1
	LogLevel_LOG_LEVEL_ERROR       LogLevel = 2
	LogLevel_LOG_LEVEL_WARN        LogLevel = 3
	LogLevel_LOG_LEVEL_INFO        LogLevel = 4
	LogLevel_LOG_LEVEL_DEBUG       LogLevel = 5
)

// Enum value maps for LogLevel.
var (
	LogLevel_name = map[int32]string{
		0: "LOG_LEVEL_UNSPECIFIED",
		1: "LOG_LEVEL_FATAL",
		2: "LOG_LEVEL_ERROR",
		3: "LOG_LEVEL_WARN",
		4: "LOG_LEVEL_INFO",
		5: "LOG_LEVEL_DEBUG",
	}
	LogLevel_value = map[string]int32{
		"LOG_LEVEL_UNSPECIFIED": 0,
		"LOG_LEVEL_FATAL":       1,
		"LOG_LEVEL_ERROR":       2,
		"LOG_LEVEL_WARN":        3,
		"LOG_LEVEL_INFO":        4,
		"LOG_LEVEL_DEBUG":       5,
	}
)

func (x LogLevel) Enum() *LogLevel {
	p := new(LogLevel)
	*p = x
	return p
}

func (x LogLevel) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LogLevel) Descriptor() protoreflect.EnumDescriptor {
	return file_inventory_v1_log_level_proto_enumTypes[0].Descriptor()
}

func (LogLevel) Type() protoreflect.EnumType {
	return &file_inventory_v1_log_level_proto_enumTypes[0]
}

func (x LogLevel) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LogLevel.Descriptor instead.
func (LogLevel) EnumDescriptor() ([]byte, []int) {
	return file_inventory_v1_log_level_proto_rawDescGZIP(), []int{0}
}

var File_inventory_v1_log_level_proto protoreflect.FileDescriptor

var file_inventory_v1_log_level_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x6c,
	0x6f, 0x67, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c,
	0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2a, 0x8c, 0x01, 0x0a,
	0x08, 0x4c, 0x6f, 0x67, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x19, 0x0a, 0x15, 0x4c, 0x4f, 0x47,
	0x5f, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49,
	0x45, 0x44, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x4c, 0x4f, 0x47, 0x5f, 0x4c, 0x45, 0x56, 0x45,
	0x4c, 0x5f, 0x46, 0x41, 0x54, 0x41, 0x4c, 0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f, 0x4c, 0x4f, 0x47,
	0x5f, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x02, 0x12, 0x12,
	0x0a, 0x0e, 0x4c, 0x4f, 0x47, 0x5f, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f, 0x57, 0x41, 0x52, 0x4e,
	0x10, 0x03, 0x12, 0x12, 0x0a, 0x0e, 0x4c, 0x4f, 0x47, 0x5f, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f,
	0x49, 0x4e, 0x46, 0x4f, 0x10, 0x04, 0x12, 0x13, 0x0a, 0x0f, 0x4c, 0x4f, 0x47, 0x5f, 0x4c, 0x45,
	0x56, 0x45, 0x4c, 0x5f, 0x44, 0x45, 0x42, 0x55, 0x47, 0x10, 0x05, 0x42, 0xa7, 0x01, 0x0a, 0x10,
	0x63, 0x6f, 0x6d, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x76, 0x31,
	0x42, 0x0d, 0x4c, 0x6f, 0x67, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x65,
	0x72, 0x63, 0x6f, 0x6e, 0x61, 0x2f, 0x70, 0x6d, 0x6d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x69, 0x6e,
	0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2f, 0x76, 0x31, 0x3b, 0x69, 0x6e, 0x76, 0x65, 0x6e,
	0x74, 0x6f, 0x72, 0x79, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x49, 0x58, 0x58, 0xaa, 0x02, 0x0c, 0x49,
	0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0c, 0x49, 0x6e,
	0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x18, 0x49, 0x6e, 0x76,
	0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0d, 0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72,
	0x79, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_inventory_v1_log_level_proto_rawDescOnce sync.Once
	file_inventory_v1_log_level_proto_rawDescData = file_inventory_v1_log_level_proto_rawDesc
)

func file_inventory_v1_log_level_proto_rawDescGZIP() []byte {
	file_inventory_v1_log_level_proto_rawDescOnce.Do(func() {
		file_inventory_v1_log_level_proto_rawDescData = protoimpl.X.CompressGZIP(file_inventory_v1_log_level_proto_rawDescData)
	})
	return file_inventory_v1_log_level_proto_rawDescData
}

var (
	file_inventory_v1_log_level_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
	file_inventory_v1_log_level_proto_goTypes   = []any{
		(LogLevel)(0), // 0: inventory.v1.LogLevel
	}
)

var file_inventory_v1_log_level_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_inventory_v1_log_level_proto_init() }
func file_inventory_v1_log_level_proto_init() {
	if File_inventory_v1_log_level_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_inventory_v1_log_level_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_inventory_v1_log_level_proto_goTypes,
		DependencyIndexes: file_inventory_v1_log_level_proto_depIdxs,
		EnumInfos:         file_inventory_v1_log_level_proto_enumTypes,
	}.Build()
	File_inventory_v1_log_level_proto = out.File
	file_inventory_v1_log_level_proto_rawDesc = nil
	file_inventory_v1_log_level_proto_goTypes = nil
	file_inventory_v1_log_level_proto_depIdxs = nil
}
