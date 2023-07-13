// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.3
// source: controller/v1/controller.proto

package controller

import (
	v1 "github.com/nlnwa/veidemann-api/go/config/v1"
	v11 "github.com/nlnwa/veidemann-api/go/frontier/v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type RunStatus int32

const (
	RunStatus_RUNNING         RunStatus = 0
	RunStatus_PAUSED          RunStatus = 1
	RunStatus_PAUSE_REQUESTED RunStatus = 2
)

// Enum value maps for RunStatus.
var (
	RunStatus_name = map[int32]string{
		0: "RUNNING",
		1: "PAUSED",
		2: "PAUSE_REQUESTED",
	}
	RunStatus_value = map[string]int32{
		"RUNNING":         0,
		"PAUSED":          1,
		"PAUSE_REQUESTED": 2,
	}
)

func (x RunStatus) Enum() *RunStatus {
	p := new(RunStatus)
	*p = x
	return p
}

func (x RunStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RunStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_controller_v1_controller_proto_enumTypes[0].Descriptor()
}

func (RunStatus) Type() protoreflect.EnumType {
	return &file_controller_v1_controller_proto_enumTypes[0]
}

func (x RunStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RunStatus.Descriptor instead.
func (RunStatus) EnumDescriptor() ([]byte, []int) {
	return file_controller_v1_controller_proto_rawDescGZIP(), []int{0}
}

// Kick of a crawl job immediately
// If a job is already running for this job_id, then new seeds are added to the job instead of starting a new one.
type RunCrawlRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JobId string `protobuf:"bytes,5,opt,name=job_id,json=jobId,proto3" json:"job_id,omitempty"`
	// If seed id is submitted, only this seed will be harvested.
	// If empty, all seeds configured with the submitted job id will be harvested.
	SeedId string `protobuf:"bytes,6,opt,name=seed_id,json=seedId,proto3" json:"seed_id,omitempty"`
}

func (x *RunCrawlRequest) Reset() {
	*x = RunCrawlRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controller_v1_controller_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RunCrawlRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RunCrawlRequest) ProtoMessage() {}

func (x *RunCrawlRequest) ProtoReflect() protoreflect.Message {
	mi := &file_controller_v1_controller_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RunCrawlRequest.ProtoReflect.Descriptor instead.
func (*RunCrawlRequest) Descriptor() ([]byte, []int) {
	return file_controller_v1_controller_proto_rawDescGZIP(), []int{0}
}

func (x *RunCrawlRequest) GetJobId() string {
	if x != nil {
		return x.JobId
	}
	return ""
}

func (x *RunCrawlRequest) GetSeedId() string {
	if x != nil {
		return x.SeedId
	}
	return ""
}

type RunCrawlReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JobExecutionId string `protobuf:"bytes,1,opt,name=job_execution_id,json=jobExecutionId,proto3" json:"job_execution_id,omitempty"`
}

func (x *RunCrawlReply) Reset() {
	*x = RunCrawlReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controller_v1_controller_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RunCrawlReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RunCrawlReply) ProtoMessage() {}

func (x *RunCrawlReply) ProtoReflect() protoreflect.Message {
	mi := &file_controller_v1_controller_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RunCrawlReply.ProtoReflect.Descriptor instead.
func (*RunCrawlReply) Descriptor() ([]byte, []int) {
	return file_controller_v1_controller_proto_rawDescGZIP(), []int{1}
}

func (x *RunCrawlReply) GetJobExecutionId() string {
	if x != nil {
		return x.JobExecutionId
	}
	return ""
}

type RoleList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Role []v1.Role `protobuf:"varint,1,rep,packed,name=role,proto3,enum=veidemann.api.config.v1.Role" json:"role,omitempty"`
}

func (x *RoleList) Reset() {
	*x = RoleList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controller_v1_controller_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoleList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoleList) ProtoMessage() {}

func (x *RoleList) ProtoReflect() protoreflect.Message {
	mi := &file_controller_v1_controller_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoleList.ProtoReflect.Descriptor instead.
func (*RoleList) Descriptor() ([]byte, []int) {
	return file_controller_v1_controller_proto_rawDescGZIP(), []int{2}
}

func (x *RoleList) GetRole() []v1.Role {
	if x != nil {
		return x.Role
	}
	return nil
}

type OpenIdConnectIssuerReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OpenIdConnectIssuer string `protobuf:"bytes,1,opt,name=open_id_connect_issuer,json=openIdConnectIssuer,proto3" json:"open_id_connect_issuer,omitempty"`
}

func (x *OpenIdConnectIssuerReply) Reset() {
	*x = OpenIdConnectIssuerReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controller_v1_controller_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OpenIdConnectIssuerReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OpenIdConnectIssuerReply) ProtoMessage() {}

func (x *OpenIdConnectIssuerReply) ProtoReflect() protoreflect.Message {
	mi := &file_controller_v1_controller_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OpenIdConnectIssuerReply.ProtoReflect.Descriptor instead.
func (*OpenIdConnectIssuerReply) Descriptor() ([]byte, []int) {
	return file_controller_v1_controller_proto_rawDescGZIP(), []int{3}
}

func (x *OpenIdConnectIssuerReply) GetOpenIdConnectIssuer() string {
	if x != nil {
		return x.OpenIdConnectIssuer
	}
	return ""
}

type CrawlerStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RunStatus RunStatus `protobuf:"varint,1,opt,name=runStatus,proto3,enum=veidemann.api.controller.v1.RunStatus" json:"runStatus,omitempty"`
	// The number of busy CrawlHostGroups which essentially is the number of web pages currently downloading
	BusyCrawlHostGroupCount int64 `protobuf:"varint,2,opt,name=busyCrawlHostGroupCount,proto3" json:"busyCrawlHostGroupCount,omitempty"`
	// Total number of queued URI's
	QueueSize int64 `protobuf:"varint,3,opt,name=queueSize,proto3" json:"queueSize,omitempty"`
}

func (x *CrawlerStatus) Reset() {
	*x = CrawlerStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controller_v1_controller_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CrawlerStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CrawlerStatus) ProtoMessage() {}

func (x *CrawlerStatus) ProtoReflect() protoreflect.Message {
	mi := &file_controller_v1_controller_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CrawlerStatus.ProtoReflect.Descriptor instead.
func (*CrawlerStatus) Descriptor() ([]byte, []int) {
	return file_controller_v1_controller_proto_rawDescGZIP(), []int{4}
}

func (x *CrawlerStatus) GetRunStatus() RunStatus {
	if x != nil {
		return x.RunStatus
	}
	return RunStatus_RUNNING
}

func (x *CrawlerStatus) GetBusyCrawlHostGroupCount() int64 {
	if x != nil {
		return x.BusyCrawlHostGroupCount
	}
	return 0
}

func (x *CrawlerStatus) GetQueueSize() int64 {
	if x != nil {
		return x.QueueSize
	}
	return 0
}

var File_controller_v1_controller_proto protoreflect.FileDescriptor

var file_controller_v1_controller_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x1b, 0x76, 0x65, 0x69, 0x64, 0x65, 0x6d, 0x61, 0x6e, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x1d, 0x63,
	0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x66, 0x72,
	0x6f, 0x6e, 0x74, 0x69, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x69,
	0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x69,
	0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x76, 0x31,
	0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x41, 0x0a,
	0x0f, 0x52, 0x75, 0x6e, 0x43, 0x72, 0x61, 0x77, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x15, 0x0a, 0x06, 0x6a, 0x6f, 0x62, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x6a, 0x6f, 0x62, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x73, 0x65, 0x65, 0x64, 0x5f,
	0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x65, 0x64, 0x49, 0x64,
	0x22, 0x39, 0x0a, 0x0d, 0x52, 0x75, 0x6e, 0x43, 0x72, 0x61, 0x77, 0x6c, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x28, 0x0a, 0x10, 0x6a, 0x6f, 0x62, 0x5f, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6a, 0x6f, 0x62,
	0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x3d, 0x0a, 0x08, 0x52,
	0x6f, 0x6c, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x31, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x1d, 0x2e, 0x76, 0x65, 0x69, 0x64, 0x65, 0x6d, 0x61, 0x6e,
	0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x2e,
	0x52, 0x6f, 0x6c, 0x65, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x22, 0x4f, 0x0a, 0x18, 0x4f, 0x70,
	0x65, 0x6e, 0x49, 0x64, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x49, 0x73, 0x73, 0x75, 0x65,
	0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x33, 0x0a, 0x16, 0x6f, 0x70, 0x65, 0x6e, 0x5f, 0x69,
	0x64, 0x5f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x73, 0x73, 0x75, 0x65, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x6f, 0x70, 0x65, 0x6e, 0x49, 0x64, 0x43, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x49, 0x73, 0x73, 0x75, 0x65, 0x72, 0x22, 0xad, 0x01, 0x0a, 0x0d,
	0x43, 0x72, 0x61, 0x77, 0x6c, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x44, 0x0a,
	0x09, 0x72, 0x75, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x26, 0x2e, 0x76, 0x65, 0x69, 0x64, 0x65, 0x6d, 0x61, 0x6e, 0x6e, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x52,
	0x75, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x09, 0x72, 0x75, 0x6e, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x38, 0x0a, 0x17, 0x62, 0x75, 0x73, 0x79, 0x43, 0x72, 0x61, 0x77, 0x6c,
	0x48, 0x6f, 0x73, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x17, 0x62, 0x75, 0x73, 0x79, 0x43, 0x72, 0x61, 0x77, 0x6c, 0x48,
	0x6f, 0x73, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1c, 0x0a,
	0x09, 0x71, 0x75, 0x65, 0x75, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x09, 0x71, 0x75, 0x65, 0x75, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x2a, 0x39, 0x0a, 0x09, 0x52,
	0x75, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0b, 0x0a, 0x07, 0x52, 0x55, 0x4e, 0x4e,
	0x49, 0x4e, 0x47, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x50, 0x41, 0x55, 0x53, 0x45, 0x44, 0x10,
	0x01, 0x12, 0x13, 0x0a, 0x0f, 0x50, 0x41, 0x55, 0x53, 0x45, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x45,
	0x53, 0x54, 0x45, 0x44, 0x10, 0x02, 0x32, 0xe1, 0x07, 0x0a, 0x0a, 0x43, 0x6f, 0x6e, 0x74, 0x72,
	0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x12, 0x58, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x6c, 0x65,
	0x73, 0x46, 0x6f, 0x72, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x25, 0x2e, 0x76, 0x65, 0x69, 0x64, 0x65, 0x6d, 0x61,
	0x6e, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x00, 0x12,
	0x66, 0x0a, 0x08, 0x52, 0x75, 0x6e, 0x43, 0x72, 0x61, 0x77, 0x6c, 0x12, 0x2c, 0x2e, 0x76, 0x65,
	0x69, 0x64, 0x65, 0x6d, 0x61, 0x6e, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x75, 0x6e, 0x43, 0x72, 0x61,
	0x77, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x76, 0x65, 0x69, 0x64,
	0x65, 0x6d, 0x61, 0x6e, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f,
	0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x75, 0x6e, 0x43, 0x72, 0x61, 0x77, 0x6c,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x72, 0x0a, 0x13, 0x41, 0x62, 0x6f, 0x72, 0x74,
	0x43, 0x72, 0x61, 0x77, 0x6c, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x28,
	0x2e, 0x76, 0x65, 0x69, 0x64, 0x65, 0x6d, 0x61, 0x6e, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63,
	0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x65,
	0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x1a, 0x2f, 0x2e, 0x76, 0x65, 0x69, 0x64, 0x65,
	0x6d, 0x61, 0x6e, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x69, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x61, 0x77, 0x6c, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74,
	0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x00, 0x12, 0x6e, 0x0a, 0x11, 0x41,
	0x62, 0x6f, 0x72, 0x74, 0x4a, 0x6f, 0x62, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x28, 0x2e, 0x76, 0x65, 0x69, 0x64, 0x65, 0x6d, 0x61, 0x6e, 0x6e, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x45,
	0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x1a, 0x2d, 0x2e, 0x76, 0x65, 0x69,
	0x64, 0x65, 0x6d, 0x61, 0x6e, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x66, 0x72, 0x6f, 0x6e, 0x74,
	0x69, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4a, 0x6f, 0x62, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74,
	0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x00, 0x12, 0x69, 0x0a, 0x16, 0x47,
	0x65, 0x74, 0x4f, 0x70, 0x65, 0x6e, 0x49, 0x64, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x49,
	0x73, 0x73, 0x75, 0x65, 0x72, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x35, 0x2e,
	0x76, 0x65, 0x69, 0x64, 0x65, 0x6d, 0x61, 0x6e, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f,
	0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x70, 0x65, 0x6e,
	0x49, 0x64, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x49, 0x73, 0x73, 0x75, 0x65, 0x72, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x40, 0x0a, 0x0c, 0x50, 0x61, 0x75, 0x73, 0x65, 0x43,
	0x72, 0x61, 0x77, 0x6c, 0x65, 0x72, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x42, 0x0a, 0x0e, 0x55, 0x6e, 0x50, 0x61,
	0x75, 0x73, 0x65, 0x43, 0x72, 0x61, 0x77, 0x6c, 0x65, 0x72, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x4e, 0x0a, 0x06,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x2a,
	0x2e, 0x76, 0x65, 0x69, 0x64, 0x65, 0x6d, 0x61, 0x6e, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63,
	0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x61,
	0x77, 0x6c, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x00, 0x12, 0x76, 0x0a, 0x1b,
	0x51, 0x75, 0x65, 0x75, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x46, 0x6f, 0x72, 0x43, 0x72, 0x61,
	0x77, 0x6c, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2b, 0x2e, 0x76, 0x65,
	0x69, 0x64, 0x65, 0x6d, 0x61, 0x6e, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x66, 0x72, 0x6f, 0x6e,
	0x74, 0x69, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x61, 0x77, 0x6c, 0x45, 0x78, 0x65,
	0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x1a, 0x28, 0x2e, 0x76, 0x65, 0x69, 0x64, 0x65,
	0x6d, 0x61, 0x6e, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x69, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x74, 0x0a, 0x1b, 0x51, 0x75, 0x65, 0x75, 0x65, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x46, 0x6f, 0x72, 0x43, 0x72, 0x61, 0x77, 0x6c, 0x48, 0x6f, 0x73, 0x74, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x12, 0x29, 0x2e, 0x76, 0x65, 0x69, 0x64, 0x65, 0x6d, 0x61, 0x6e, 0x6e, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x69, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x72, 0x61, 0x77, 0x6c, 0x48, 0x6f, 0x73, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x1a, 0x28,
	0x2e, 0x76, 0x65, 0x69, 0x64, 0x65, 0x6d, 0x61, 0x6e, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x66,
	0x72, 0x6f, 0x6e, 0x74, 0x69, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x78, 0x0a, 0x25, 0x6e, 0x6f,
	0x2e, 0x6e, 0x62, 0x2e, 0x6e, 0x6e, 0x61, 0x2e, 0x76, 0x65, 0x69, 0x64, 0x65, 0x6d, 0x61, 0x6e,
	0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x42, 0x11, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x01, 0x5a, 0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x6c, 0x6e, 0x77, 0x61, 0x2f, 0x76, 0x65, 0x69, 0x64, 0x65,
	0x6d, 0x61, 0x6e, 0x6e, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6f, 0x2f, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x3b, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f,
	0x6c, 0x6c, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_controller_v1_controller_proto_rawDescOnce sync.Once
	file_controller_v1_controller_proto_rawDescData = file_controller_v1_controller_proto_rawDesc
)

func file_controller_v1_controller_proto_rawDescGZIP() []byte {
	file_controller_v1_controller_proto_rawDescOnce.Do(func() {
		file_controller_v1_controller_proto_rawDescData = protoimpl.X.CompressGZIP(file_controller_v1_controller_proto_rawDescData)
	})
	return file_controller_v1_controller_proto_rawDescData
}

var file_controller_v1_controller_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_controller_v1_controller_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_controller_v1_controller_proto_goTypes = []interface{}{
	(RunStatus)(0),                   // 0: veidemann.api.controller.v1.RunStatus
	(*RunCrawlRequest)(nil),          // 1: veidemann.api.controller.v1.RunCrawlRequest
	(*RunCrawlReply)(nil),            // 2: veidemann.api.controller.v1.RunCrawlReply
	(*RoleList)(nil),                 // 3: veidemann.api.controller.v1.RoleList
	(*OpenIdConnectIssuerReply)(nil), // 4: veidemann.api.controller.v1.OpenIdConnectIssuerReply
	(*CrawlerStatus)(nil),            // 5: veidemann.api.controller.v1.CrawlerStatus
	(v1.Role)(0),                     // 6: veidemann.api.config.v1.Role
	(*emptypb.Empty)(nil),            // 7: google.protobuf.Empty
	(*ExecutionId)(nil),              // 8: veidemann.api.controller.v1.ExecutionId
	(*v11.CrawlExecutionId)(nil),     // 9: veidemann.api.frontier.v1.CrawlExecutionId
	(*v11.CrawlHostGroup)(nil),       // 10: veidemann.api.frontier.v1.CrawlHostGroup
	(*v11.CrawlExecutionStatus)(nil), // 11: veidemann.api.frontier.v1.CrawlExecutionStatus
	(*v11.JobExecutionStatus)(nil),   // 12: veidemann.api.frontier.v1.JobExecutionStatus
	(*v11.CountResponse)(nil),        // 13: veidemann.api.frontier.v1.CountResponse
}
var file_controller_v1_controller_proto_depIdxs = []int32{
	6,  // 0: veidemann.api.controller.v1.RoleList.role:type_name -> veidemann.api.config.v1.Role
	0,  // 1: veidemann.api.controller.v1.CrawlerStatus.runStatus:type_name -> veidemann.api.controller.v1.RunStatus
	7,  // 2: veidemann.api.controller.v1.Controller.GetRolesForActiveUser:input_type -> google.protobuf.Empty
	1,  // 3: veidemann.api.controller.v1.Controller.RunCrawl:input_type -> veidemann.api.controller.v1.RunCrawlRequest
	8,  // 4: veidemann.api.controller.v1.Controller.AbortCrawlExecution:input_type -> veidemann.api.controller.v1.ExecutionId
	8,  // 5: veidemann.api.controller.v1.Controller.AbortJobExecution:input_type -> veidemann.api.controller.v1.ExecutionId
	7,  // 6: veidemann.api.controller.v1.Controller.GetOpenIdConnectIssuer:input_type -> google.protobuf.Empty
	7,  // 7: veidemann.api.controller.v1.Controller.PauseCrawler:input_type -> google.protobuf.Empty
	7,  // 8: veidemann.api.controller.v1.Controller.UnPauseCrawler:input_type -> google.protobuf.Empty
	7,  // 9: veidemann.api.controller.v1.Controller.Status:input_type -> google.protobuf.Empty
	9,  // 10: veidemann.api.controller.v1.Controller.QueueCountForCrawlExecution:input_type -> veidemann.api.frontier.v1.CrawlExecutionId
	10, // 11: veidemann.api.controller.v1.Controller.QueueCountForCrawlHostGroup:input_type -> veidemann.api.frontier.v1.CrawlHostGroup
	3,  // 12: veidemann.api.controller.v1.Controller.GetRolesForActiveUser:output_type -> veidemann.api.controller.v1.RoleList
	2,  // 13: veidemann.api.controller.v1.Controller.RunCrawl:output_type -> veidemann.api.controller.v1.RunCrawlReply
	11, // 14: veidemann.api.controller.v1.Controller.AbortCrawlExecution:output_type -> veidemann.api.frontier.v1.CrawlExecutionStatus
	12, // 15: veidemann.api.controller.v1.Controller.AbortJobExecution:output_type -> veidemann.api.frontier.v1.JobExecutionStatus
	4,  // 16: veidemann.api.controller.v1.Controller.GetOpenIdConnectIssuer:output_type -> veidemann.api.controller.v1.OpenIdConnectIssuerReply
	7,  // 17: veidemann.api.controller.v1.Controller.PauseCrawler:output_type -> google.protobuf.Empty
	7,  // 18: veidemann.api.controller.v1.Controller.UnPauseCrawler:output_type -> google.protobuf.Empty
	5,  // 19: veidemann.api.controller.v1.Controller.Status:output_type -> veidemann.api.controller.v1.CrawlerStatus
	13, // 20: veidemann.api.controller.v1.Controller.QueueCountForCrawlExecution:output_type -> veidemann.api.frontier.v1.CountResponse
	13, // 21: veidemann.api.controller.v1.Controller.QueueCountForCrawlHostGroup:output_type -> veidemann.api.frontier.v1.CountResponse
	12, // [12:22] is the sub-list for method output_type
	2,  // [2:12] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_controller_v1_controller_proto_init() }
func file_controller_v1_controller_proto_init() {
	if File_controller_v1_controller_proto != nil {
		return
	}
	file_controller_v1_resources_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_controller_v1_controller_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RunCrawlRequest); i {
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
		file_controller_v1_controller_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RunCrawlReply); i {
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
		file_controller_v1_controller_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoleList); i {
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
		file_controller_v1_controller_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OpenIdConnectIssuerReply); i {
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
		file_controller_v1_controller_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CrawlerStatus); i {
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
			RawDescriptor: file_controller_v1_controller_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_controller_v1_controller_proto_goTypes,
		DependencyIndexes: file_controller_v1_controller_proto_depIdxs,
		EnumInfos:         file_controller_v1_controller_proto_enumTypes,
		MessageInfos:      file_controller_v1_controller_proto_msgTypes,
	}.Build()
	File_controller_v1_controller_proto = out.File
	file_controller_v1_controller_proto_rawDesc = nil
	file_controller_v1_controller_proto_goTypes = nil
	file_controller_v1_controller_proto_depIdxs = nil
}
