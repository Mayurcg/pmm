// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: qan/v1/filters.proto

package qanv1

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// GetFilteredMetricsNamesRequest contains period for which we need filters.
type GetFilteredMetricsNamesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PeriodStartFrom *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=period_start_from,json=periodStartFrom,proto3" json:"period_start_from,omitempty"`
	PeriodStartTo   *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=period_start_to,json=periodStartTo,proto3" json:"period_start_to,omitempty"`
	MainMetricName  string                 `protobuf:"bytes,3,opt,name=main_metric_name,json=mainMetricName,proto3" json:"main_metric_name,omitempty"`
	Labels          []*MapFieldEntry       `protobuf:"bytes,4,rep,name=labels,proto3" json:"labels,omitempty"`
}

func (x *GetFilteredMetricsNamesRequest) Reset() {
	*x = GetFilteredMetricsNamesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_qan_v1_filters_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFilteredMetricsNamesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFilteredMetricsNamesRequest) ProtoMessage() {}

func (x *GetFilteredMetricsNamesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_qan_v1_filters_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFilteredMetricsNamesRequest.ProtoReflect.Descriptor instead.
func (*GetFilteredMetricsNamesRequest) Descriptor() ([]byte, []int) {
	return file_qan_v1_filters_proto_rawDescGZIP(), []int{0}
}

func (x *GetFilteredMetricsNamesRequest) GetPeriodStartFrom() *timestamppb.Timestamp {
	if x != nil {
		return x.PeriodStartFrom
	}
	return nil
}

func (x *GetFilteredMetricsNamesRequest) GetPeriodStartTo() *timestamppb.Timestamp {
	if x != nil {
		return x.PeriodStartTo
	}
	return nil
}

func (x *GetFilteredMetricsNamesRequest) GetMainMetricName() string {
	if x != nil {
		return x.MainMetricName
	}
	return ""
}

func (x *GetFilteredMetricsNamesRequest) GetLabels() []*MapFieldEntry {
	if x != nil {
		return x.Labels
	}
	return nil
}

// GetFilteredMetricsNamesResponse is map of labels for given period by key.
// Key is label's name and value is label's value and how many times it occur.
type GetFilteredMetricsNamesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Labels map[string]*ListLabels `protobuf:"bytes,1,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *GetFilteredMetricsNamesResponse) Reset() {
	*x = GetFilteredMetricsNamesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_qan_v1_filters_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFilteredMetricsNamesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFilteredMetricsNamesResponse) ProtoMessage() {}

func (x *GetFilteredMetricsNamesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_qan_v1_filters_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFilteredMetricsNamesResponse.ProtoReflect.Descriptor instead.
func (*GetFilteredMetricsNamesResponse) Descriptor() ([]byte, []int) {
	return file_qan_v1_filters_proto_rawDescGZIP(), []int{1}
}

func (x *GetFilteredMetricsNamesResponse) GetLabels() map[string]*ListLabels {
	if x != nil {
		return x.Labels
	}
	return nil
}

// ListLabels is list of label's values: duplicates are impossible.
type ListLabels struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name []*Values `protobuf:"bytes,1,rep,name=name,proto3" json:"name,omitempty"`
}

func (x *ListLabels) Reset() {
	*x = ListLabels{}
	if protoimpl.UnsafeEnabled {
		mi := &file_qan_v1_filters_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListLabels) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListLabels) ProtoMessage() {}

func (x *ListLabels) ProtoReflect() protoreflect.Message {
	mi := &file_qan_v1_filters_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListLabels.ProtoReflect.Descriptor instead.
func (*ListLabels) Descriptor() ([]byte, []int) {
	return file_qan_v1_filters_proto_rawDescGZIP(), []int{2}
}

func (x *ListLabels) GetName() []*Values {
	if x != nil {
		return x.Name
	}
	return nil
}

// Values is label values and main metric percent and per second.
type Values struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value             string  `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	MainMetricPercent float32 `protobuf:"fixed32,2,opt,name=main_metric_percent,json=mainMetricPercent,proto3" json:"main_metric_percent,omitempty"`
	MainMetricPerSec  float32 `protobuf:"fixed32,3,opt,name=main_metric_per_sec,json=mainMetricPerSec,proto3" json:"main_metric_per_sec,omitempty"`
}

func (x *Values) Reset() {
	*x = Values{}
	if protoimpl.UnsafeEnabled {
		mi := &file_qan_v1_filters_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Values) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Values) ProtoMessage() {}

func (x *Values) ProtoReflect() protoreflect.Message {
	mi := &file_qan_v1_filters_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Values.ProtoReflect.Descriptor instead.
func (*Values) Descriptor() ([]byte, []int) {
	return file_qan_v1_filters_proto_rawDescGZIP(), []int{3}
}

func (x *Values) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *Values) GetMainMetricPercent() float32 {
	if x != nil {
		return x.MainMetricPercent
	}
	return 0
}

func (x *Values) GetMainMetricPerSec() float32 {
	if x != nil {
		return x.MainMetricPerSec
	}
	return 0
}

var File_qan_v1_filters_proto protoreflect.FileDescriptor

var file_qan_v1_filters_proto_rawDesc = []byte{
	0x0a, 0x14, 0x71, 0x61, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x71, 0x61, 0x6e, 0x2e, 0x76, 0x31, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x10, 0x71, 0x61, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x71, 0x61, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x85, 0x02, 0x0a, 0x1e, 0x47, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x65,
	0x64, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x46, 0x0a, 0x11, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x5f, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x5f, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0f, 0x70, 0x65, 0x72,
	0x69, 0x6f, 0x64, 0x53, 0x74, 0x61, 0x72, 0x74, 0x46, 0x72, 0x6f, 0x6d, 0x12, 0x42, 0x0a, 0x0f,
	0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x6f, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x0d, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x53, 0x74, 0x61, 0x72, 0x74, 0x54, 0x6f,
	0x12, 0x28, 0x0a, 0x10, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6d, 0x61, 0x69, 0x6e,
	0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2d, 0x0a, 0x06, 0x6c, 0x61,
	0x62, 0x65, 0x6c, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x71, 0x61, 0x6e,
	0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x61, 0x70, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x22, 0xbd, 0x01, 0x0a, 0x1f, 0x47, 0x65,
	0x74, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x65, 0x64, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73,
	0x4e, 0x61, 0x6d, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4b, 0x0a,
	0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x33, 0x2e,
	0x71, 0x61, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x65, 0x64, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x1a, 0x4d, 0x0a, 0x0b, 0x4c, 0x61,
	0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x28, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x71, 0x61, 0x6e,
	0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x30, 0x0a, 0x0a, 0x4c, 0x69, 0x73,
	0x74, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x22, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x71, 0x61, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x73, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x7d, 0x0a, 0x06, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x2e, 0x0a, 0x13, 0x6d,
	0x61, 0x69, 0x6e, 0x5f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x5f, 0x70, 0x65, 0x72, 0x63, 0x65,
	0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x11, 0x6d, 0x61, 0x69, 0x6e, 0x4d, 0x65,
	0x74, 0x72, 0x69, 0x63, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x12, 0x2d, 0x0a, 0x13, 0x6d,
	0x61, 0x69, 0x6e, 0x5f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x5f, 0x70, 0x65, 0x72, 0x5f, 0x73,
	0x65, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x10, 0x6d, 0x61, 0x69, 0x6e, 0x4d, 0x65,
	0x74, 0x72, 0x69, 0x63, 0x50, 0x65, 0x72, 0x53, 0x65, 0x63, 0x42, 0x7c, 0x0a, 0x0a, 0x63, 0x6f,
	0x6d, 0x2e, 0x71, 0x61, 0x6e, 0x2e, 0x76, 0x31, 0x42, 0x0c, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x65, 0x72, 0x63, 0x6f, 0x6e, 0x61, 0x2f, 0x70, 0x6d, 0x6d,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x71, 0x61, 0x6e, 0x2f, 0x76, 0x31, 0x3b, 0x71, 0x61, 0x6e, 0x76,
	0x31, 0xa2, 0x02, 0x03, 0x51, 0x58, 0x58, 0xaa, 0x02, 0x06, 0x51, 0x61, 0x6e, 0x2e, 0x56, 0x31,
	0xca, 0x02, 0x06, 0x51, 0x61, 0x6e, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x12, 0x51, 0x61, 0x6e, 0x5c,
	0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02,
	0x07, 0x51, 0x61, 0x6e, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_qan_v1_filters_proto_rawDescOnce sync.Once
	file_qan_v1_filters_proto_rawDescData = file_qan_v1_filters_proto_rawDesc
)

func file_qan_v1_filters_proto_rawDescGZIP() []byte {
	file_qan_v1_filters_proto_rawDescOnce.Do(func() {
		file_qan_v1_filters_proto_rawDescData = protoimpl.X.CompressGZIP(file_qan_v1_filters_proto_rawDescData)
	})
	return file_qan_v1_filters_proto_rawDescData
}

var (
	file_qan_v1_filters_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
	file_qan_v1_filters_proto_goTypes  = []any{
		(*GetFilteredMetricsNamesRequest)(nil),  // 0: qan.v1.GetFilteredMetricsNamesRequest
		(*GetFilteredMetricsNamesResponse)(nil), // 1: qan.v1.GetFilteredMetricsNamesResponse
		(*ListLabels)(nil),                      // 2: qan.v1.ListLabels
		(*Values)(nil),                          // 3: qan.v1.Values
		nil,                                     // 4: qan.v1.GetFilteredMetricsNamesResponse.LabelsEntry
		(*timestamppb.Timestamp)(nil),           // 5: google.protobuf.Timestamp
		(*MapFieldEntry)(nil),                   // 6: qan.v1.MapFieldEntry
	}
)

var file_qan_v1_filters_proto_depIdxs = []int32{
	5, // 0: qan.v1.GetFilteredMetricsNamesRequest.period_start_from:type_name -> google.protobuf.Timestamp
	5, // 1: qan.v1.GetFilteredMetricsNamesRequest.period_start_to:type_name -> google.protobuf.Timestamp
	6, // 2: qan.v1.GetFilteredMetricsNamesRequest.labels:type_name -> qan.v1.MapFieldEntry
	4, // 3: qan.v1.GetFilteredMetricsNamesResponse.labels:type_name -> qan.v1.GetFilteredMetricsNamesResponse.LabelsEntry
	3, // 4: qan.v1.ListLabels.name:type_name -> qan.v1.Values
	2, // 5: qan.v1.GetFilteredMetricsNamesResponse.LabelsEntry.value:type_name -> qan.v1.ListLabels
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_qan_v1_filters_proto_init() }
func file_qan_v1_filters_proto_init() {
	if File_qan_v1_filters_proto != nil {
		return
	}
	file_qan_v1_qan_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_qan_v1_filters_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*GetFilteredMetricsNamesRequest); i {
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
		file_qan_v1_filters_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*GetFilteredMetricsNamesResponse); i {
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
		file_qan_v1_filters_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*ListLabels); i {
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
		file_qan_v1_filters_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*Values); i {
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
			RawDescriptor: file_qan_v1_filters_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_qan_v1_filters_proto_goTypes,
		DependencyIndexes: file_qan_v1_filters_proto_depIdxs,
		MessageInfos:      file_qan_v1_filters_proto_msgTypes,
	}.Build()
	File_qan_v1_filters_proto = out.File
	file_qan_v1_filters_proto_rawDesc = nil
	file_qan_v1_filters_proto_goTypes = nil
	file_qan_v1_filters_proto_depIdxs = nil
}
