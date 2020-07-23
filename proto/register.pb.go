// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v4.0.0
// source: register.proto

package registerproto

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type AppExporterInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Team       string   `protobuf:"bytes,2,opt,name=team,proto3" json:"team,omitempty"`
	Tags       []string `protobuf:"bytes,3,rep,name=tags,proto3" json:"tags,omitempty"`
	Scheme     string   `protobuf:"bytes,4,opt,name=scheme,proto3" json:"scheme,omitempty"`
	MetricPath string   `protobuf:"bytes,5,opt,name=metricPath,proto3" json:"metricPath,omitempty"`
}

func (x *AppExporterInfo) Reset() {
	*x = AppExporterInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_register_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppExporterInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppExporterInfo) ProtoMessage() {}

func (x *AppExporterInfo) ProtoReflect() protoreflect.Message {
	mi := &file_register_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppExporterInfo.ProtoReflect.Descriptor instead.
func (*AppExporterInfo) Descriptor() ([]byte, []int) {
	return file_register_proto_rawDescGZIP(), []int{0}
}

func (x *AppExporterInfo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *AppExporterInfo) GetTeam() string {
	if x != nil {
		return x.Team
	}
	return ""
}

func (x *AppExporterInfo) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *AppExporterInfo) GetScheme() string {
	if x != nil {
		return x.Scheme
	}
	return ""
}

func (x *AppExporterInfo) GetMetricPath() string {
	if x != nil {
		return x.MetricPath
	}
	return ""
}

type ListReqMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Team string `protobuf:"bytes,1,opt,name=team,proto3" json:"team,omitempty"`
}

func (x *ListReqMsg) Reset() {
	*x = ListReqMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_register_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListReqMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListReqMsg) ProtoMessage() {}

func (x *ListReqMsg) ProtoReflect() protoreflect.Message {
	mi := &file_register_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListReqMsg.ProtoReflect.Descriptor instead.
func (*ListReqMsg) Descriptor() ([]byte, []int) {
	return file_register_proto_rawDescGZIP(), []int{1}
}

func (x *ListReqMsg) GetTeam() string {
	if x != nil {
		return x.Team
	}
	return ""
}

type ListRespMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppInfos []*AppExporterInfo `protobuf:"bytes,1,rep,name=appInfos,proto3" json:"appInfos,omitempty"`
}

func (x *ListRespMsg) Reset() {
	*x = ListRespMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_register_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRespMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRespMsg) ProtoMessage() {}

func (x *ListRespMsg) ProtoReflect() protoreflect.Message {
	mi := &file_register_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRespMsg.ProtoReflect.Descriptor instead.
func (*ListRespMsg) Descriptor() ([]byte, []int) {
	return file_register_proto_rawDescGZIP(), []int{2}
}

func (x *ListRespMsg) GetAppInfos() []*AppExporterInfo {
	if x != nil {
		return x.AppInfos
	}
	return nil
}

type AddReqMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Team     string             `protobuf:"bytes,1,opt,name=team,proto3" json:"team,omitempty"`
	AppInfos []*AppExporterInfo `protobuf:"bytes,2,rep,name=appInfos,proto3" json:"appInfos,omitempty"`
}

func (x *AddReqMsg) Reset() {
	*x = AddReqMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_register_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddReqMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddReqMsg) ProtoMessage() {}

func (x *AddReqMsg) ProtoReflect() protoreflect.Message {
	mi := &file_register_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddReqMsg.ProtoReflect.Descriptor instead.
func (*AddReqMsg) Descriptor() ([]byte, []int) {
	return file_register_proto_rawDescGZIP(), []int{3}
}

func (x *AddReqMsg) GetTeam() string {
	if x != nil {
		return x.Team
	}
	return ""
}

func (x *AddReqMsg) GetAppInfos() []*AppExporterInfo {
	if x != nil {
		return x.AppInfos
	}
	return nil
}

type UpdateReqMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Team     string             `protobuf:"bytes,1,opt,name=team,proto3" json:"team,omitempty"`
	AppInfos []*AppExporterInfo `protobuf:"bytes,2,rep,name=appInfos,proto3" json:"appInfos,omitempty"`
}

func (x *UpdateReqMsg) Reset() {
	*x = UpdateReqMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_register_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateReqMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateReqMsg) ProtoMessage() {}

func (x *UpdateReqMsg) ProtoReflect() protoreflect.Message {
	mi := &file_register_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateReqMsg.ProtoReflect.Descriptor instead.
func (*UpdateReqMsg) Descriptor() ([]byte, []int) {
	return file_register_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateReqMsg) GetTeam() string {
	if x != nil {
		return x.Team
	}
	return ""
}

func (x *UpdateReqMsg) GetAppInfos() []*AppExporterInfo {
	if x != nil {
		return x.AppInfos
	}
	return nil
}

type RemoveReqMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Team      string   `protobuf:"bytes,1,opt,name=team,proto3" json:"team,omitempty"`
	ServiceId []string `protobuf:"bytes,2,rep,name=serviceId,proto3" json:"serviceId,omitempty"`
}

func (x *RemoveReqMsg) Reset() {
	*x = RemoveReqMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_register_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveReqMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveReqMsg) ProtoMessage() {}

func (x *RemoveReqMsg) ProtoReflect() protoreflect.Message {
	mi := &file_register_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveReqMsg.ProtoReflect.Descriptor instead.
func (*RemoveReqMsg) Descriptor() ([]byte, []int) {
	return file_register_proto_rawDescGZIP(), []int{5}
}

func (x *RemoveReqMsg) GetTeam() string {
	if x != nil {
		return x.Team
	}
	return ""
}

func (x *RemoveReqMsg) GetServiceId() []string {
	if x != nil {
		return x.ServiceId
	}
	return nil
}

type RespResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ret int32 `protobuf:"varint,1,opt,name=ret,proto3" json:"ret,omitempty"`
}

func (x *RespResult) Reset() {
	*x = RespResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_register_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RespResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RespResult) ProtoMessage() {}

func (x *RespResult) ProtoReflect() protoreflect.Message {
	mi := &file_register_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RespResult.ProtoReflect.Descriptor instead.
func (*RespResult) Descriptor() ([]byte, []int) {
	return file_register_proto_rawDescGZIP(), []int{6}
}

func (x *RespResult) GetRet() int32 {
	if x != nil {
		return x.Ret
	}
	return 0
}

var File_register_proto protoreflect.FileDescriptor

var file_register_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0d, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x81, 0x01, 0x0a, 0x0f, 0x41, 0x70, 0x70, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x61, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x74, 0x65, 0x61, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x63, 0x68, 0x65, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x63, 0x68,
	0x65, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x50, 0x61, 0x74,
	0x68, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x50,
	0x61, 0x74, 0x68, 0x22, 0x20, 0x0a, 0x0a, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x4d, 0x73,
	0x67, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x61, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x74, 0x65, 0x61, 0x6d, 0x22, 0x49, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x4d, 0x73, 0x67, 0x12, 0x3a, 0x0a, 0x08, 0x61, 0x70, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65,
	0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x70, 0x70, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74,
	0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x61, 0x70, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x73,
	0x22, 0x5b, 0x0a, 0x09, 0x41, 0x64, 0x64, 0x52, 0x65, 0x71, 0x4d, 0x73, 0x67, 0x12, 0x12, 0x0a,
	0x04, 0x74, 0x65, 0x61, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x61,
	0x6d, 0x12, 0x3a, 0x0a, 0x08, 0x61, 0x70, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x70, 0x70, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x08, 0x61, 0x70, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x22, 0x5e, 0x0a,
	0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x4d, 0x73, 0x67, 0x12, 0x12, 0x0a,
	0x04, 0x74, 0x65, 0x61, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x61,
	0x6d, 0x12, 0x3a, 0x0a, 0x08, 0x61, 0x70, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x70, 0x70, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x08, 0x61, 0x70, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x22, 0x40, 0x0a,
	0x0c, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x52, 0x65, 0x71, 0x4d, 0x73, 0x67, 0x12, 0x12, 0x0a,
	0x04, 0x74, 0x65, 0x61, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x61,
	0x6d, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x22,
	0x1e, 0x0a, 0x0a, 0x52, 0x65, 0x73, 0x70, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x10, 0x0a,
	0x03, 0x72, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x72, 0x65, 0x74, 0x32,
	0xac, 0x02, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x46, 0x0a, 0x0b, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x2e, 0x72, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x71, 0x4d, 0x73, 0x67, 0x1a, 0x1a, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x4d, 0x73,
	0x67, 0x22, 0x00, 0x12, 0x43, 0x0a, 0x0a, 0x41, 0x64, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x18, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x41, 0x64, 0x64, 0x52, 0x65, 0x71, 0x4d, 0x73, 0x67, 0x1a, 0x19, 0x2e, 0x72, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x70,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x12, 0x49, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x2e, 0x72, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x4d, 0x73, 0x67, 0x1a, 0x19, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65,
	0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x22, 0x00, 0x12, 0x49, 0x0a, 0x0d, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x52, 0x65, 0x71, 0x4d, 0x73,
	0x67, 0x1a, 0x19, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_register_proto_rawDescOnce sync.Once
	file_register_proto_rawDescData = file_register_proto_rawDesc
)

func file_register_proto_rawDescGZIP() []byte {
	file_register_proto_rawDescOnce.Do(func() {
		file_register_proto_rawDescData = protoimpl.X.CompressGZIP(file_register_proto_rawDescData)
	})
	return file_register_proto_rawDescData
}

var file_register_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_register_proto_goTypes = []interface{}{
	(*AppExporterInfo)(nil), // 0: registerproto.AppExporterInfo
	(*ListReqMsg)(nil),      // 1: registerproto.ListReqMsg
	(*ListRespMsg)(nil),     // 2: registerproto.ListRespMsg
	(*AddReqMsg)(nil),       // 3: registerproto.AddReqMsg
	(*UpdateReqMsg)(nil),    // 4: registerproto.UpdateReqMsg
	(*RemoveReqMsg)(nil),    // 5: registerproto.RemoveReqMsg
	(*RespResult)(nil),      // 6: registerproto.RespResult
}
var file_register_proto_depIdxs = []int32{
	0, // 0: registerproto.ListRespMsg.appInfos:type_name -> registerproto.AppExporterInfo
	0, // 1: registerproto.AddReqMsg.appInfos:type_name -> registerproto.AppExporterInfo
	0, // 2: registerproto.UpdateReqMsg.appInfos:type_name -> registerproto.AppExporterInfo
	1, // 3: registerproto.Request.ListRequest:input_type -> registerproto.ListReqMsg
	3, // 4: registerproto.Request.AddRequest:input_type -> registerproto.AddReqMsg
	4, // 5: registerproto.Request.UpdateRequest:input_type -> registerproto.UpdateReqMsg
	5, // 6: registerproto.Request.RemoveRequest:input_type -> registerproto.RemoveReqMsg
	2, // 7: registerproto.Request.ListRequest:output_type -> registerproto.ListRespMsg
	6, // 8: registerproto.Request.AddRequest:output_type -> registerproto.RespResult
	6, // 9: registerproto.Request.UpdateRequest:output_type -> registerproto.RespResult
	6, // 10: registerproto.Request.RemoveRequest:output_type -> registerproto.RespResult
	7, // [7:11] is the sub-list for method output_type
	3, // [3:7] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_register_proto_init() }
func file_register_proto_init() {
	if File_register_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_register_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppExporterInfo); i {
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
		file_register_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListReqMsg); i {
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
		file_register_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRespMsg); i {
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
		file_register_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddReqMsg); i {
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
		file_register_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateReqMsg); i {
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
		file_register_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveReqMsg); i {
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
		file_register_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RespResult); i {
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
			RawDescriptor: file_register_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_register_proto_goTypes,
		DependencyIndexes: file_register_proto_depIdxs,
		MessageInfos:      file_register_proto_msgTypes,
	}.Build()
	File_register_proto = out.File
	file_register_proto_rawDesc = nil
	file_register_proto_goTypes = nil
	file_register_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// RequestClient is the client API for Request service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RequestClient interface {
	ListRequest(ctx context.Context, in *ListReqMsg, opts ...grpc.CallOption) (*ListRespMsg, error)
	AddRequest(ctx context.Context, in *AddReqMsg, opts ...grpc.CallOption) (*RespResult, error)
	UpdateRequest(ctx context.Context, in *UpdateReqMsg, opts ...grpc.CallOption) (*RespResult, error)
	RemoveRequest(ctx context.Context, in *RemoveReqMsg, opts ...grpc.CallOption) (*RespResult, error)
}

type requestClient struct {
	cc grpc.ClientConnInterface
}

func NewRequestClient(cc grpc.ClientConnInterface) RequestClient {
	return &requestClient{cc}
}

func (c *requestClient) ListRequest(ctx context.Context, in *ListReqMsg, opts ...grpc.CallOption) (*ListRespMsg, error) {
	out := new(ListRespMsg)
	err := c.cc.Invoke(ctx, "/registerproto.Request/ListRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *requestClient) AddRequest(ctx context.Context, in *AddReqMsg, opts ...grpc.CallOption) (*RespResult, error) {
	out := new(RespResult)
	err := c.cc.Invoke(ctx, "/registerproto.Request/AddRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *requestClient) UpdateRequest(ctx context.Context, in *UpdateReqMsg, opts ...grpc.CallOption) (*RespResult, error) {
	out := new(RespResult)
	err := c.cc.Invoke(ctx, "/registerproto.Request/UpdateRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *requestClient) RemoveRequest(ctx context.Context, in *RemoveReqMsg, opts ...grpc.CallOption) (*RespResult, error) {
	out := new(RespResult)
	err := c.cc.Invoke(ctx, "/registerproto.Request/RemoveRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RequestServer is the server API for Request service.
type RequestServer interface {
	ListRequest(context.Context, *ListReqMsg) (*ListRespMsg, error)
	AddRequest(context.Context, *AddReqMsg) (*RespResult, error)
	UpdateRequest(context.Context, *UpdateReqMsg) (*RespResult, error)
	RemoveRequest(context.Context, *RemoveReqMsg) (*RespResult, error)
}

// UnimplementedRequestServer can be embedded to have forward compatible implementations.
type UnimplementedRequestServer struct {
}

func (*UnimplementedRequestServer) ListRequest(context.Context, *ListReqMsg) (*ListRespMsg, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListRequest not implemented")
}
func (*UnimplementedRequestServer) AddRequest(context.Context, *AddReqMsg) (*RespResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddRequest not implemented")
}
func (*UnimplementedRequestServer) UpdateRequest(context.Context, *UpdateReqMsg) (*RespResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRequest not implemented")
}
func (*UnimplementedRequestServer) RemoveRequest(context.Context, *RemoveReqMsg) (*RespResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveRequest not implemented")
}

func RegisterRequestServer(s *grpc.Server, srv RequestServer) {
	s.RegisterService(&_Request_serviceDesc, srv)
}

func _Request_ListRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListReqMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RequestServer).ListRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/registerproto.Request/ListRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RequestServer).ListRequest(ctx, req.(*ListReqMsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _Request_AddRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddReqMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RequestServer).AddRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/registerproto.Request/AddRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RequestServer).AddRequest(ctx, req.(*AddReqMsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _Request_UpdateRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateReqMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RequestServer).UpdateRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/registerproto.Request/UpdateRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RequestServer).UpdateRequest(ctx, req.(*UpdateReqMsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _Request_RemoveRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveReqMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RequestServer).RemoveRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/registerproto.Request/RemoveRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RequestServer).RemoveRequest(ctx, req.(*RemoveReqMsg))
	}
	return interceptor(ctx, in, info, handler)
}

var _Request_serviceDesc = grpc.ServiceDesc{
	ServiceName: "registerproto.Request",
	HandlerType: (*RequestServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListRequest",
			Handler:    _Request_ListRequest_Handler,
		},
		{
			MethodName: "AddRequest",
			Handler:    _Request_AddRequest_Handler,
		},
		{
			MethodName: "UpdateRequest",
			Handler:    _Request_UpdateRequest_Handler,
		},
		{
			MethodName: "RemoveRequest",
			Handler:    _Request_RemoveRequest_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "register.proto",
}