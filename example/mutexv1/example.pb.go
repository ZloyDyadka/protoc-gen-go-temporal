// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: example.proto

// buf:lint:ignore PACKAGE_DIRECTORY_MATCH

package mutexv1

import (
	_ "github.com/cludden/protoc-gen-go-temporal/gen/temporal/v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
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

// MutexRequest describes the input to a Mutex workflow/activity
type MutexRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Resource string `protobuf:"bytes,1,opt,name=resource,proto3" json:"resource,omitempty"`
}

func (x *MutexRequest) Reset() {
	*x = MutexRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_example_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MutexRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MutexRequest) ProtoMessage() {}

func (x *MutexRequest) ProtoReflect() protoreflect.Message {
	mi := &file_example_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MutexRequest.ProtoReflect.Descriptor instead.
func (*MutexRequest) Descriptor() ([]byte, []int) {
	return file_example_proto_rawDescGZIP(), []int{0}
}

func (x *MutexRequest) GetResource() string {
	if x != nil {
		return x.Resource
	}
	return ""
}

// SampleWorkflowWithMutexRequest describes the input to a SampleWorkflowWithMutex workflow
type SampleWorkflowWithMutexRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Resource string  `protobuf:"bytes,1,opt,name=resource,proto3" json:"resource,omitempty"`
	Dest     string  `protobuf:"bytes,2,opt,name=dest,proto3" json:"dest,omitempty"`
	Amount   float64 `protobuf:"fixed64,3,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *SampleWorkflowWithMutexRequest) Reset() {
	*x = SampleWorkflowWithMutexRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_example_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SampleWorkflowWithMutexRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SampleWorkflowWithMutexRequest) ProtoMessage() {}

func (x *SampleWorkflowWithMutexRequest) ProtoReflect() protoreflect.Message {
	mi := &file_example_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SampleWorkflowWithMutexRequest.ProtoReflect.Descriptor instead.
func (*SampleWorkflowWithMutexRequest) Descriptor() ([]byte, []int) {
	return file_example_proto_rawDescGZIP(), []int{1}
}

func (x *SampleWorkflowWithMutexRequest) GetResource() string {
	if x != nil {
		return x.Resource
	}
	return ""
}

func (x *SampleWorkflowWithMutexRequest) GetDest() string {
	if x != nil {
		return x.Dest
	}
	return ""
}

func (x *SampleWorkflowWithMutexRequest) GetAmount() float64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

// SampleWorkflowWithMutexResponse describes the output from a SampleWorkflowWithMutex workflow
type SampleWorkflowWithMutexResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result string `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *SampleWorkflowWithMutexResponse) Reset() {
	*x = SampleWorkflowWithMutexResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_example_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SampleWorkflowWithMutexResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SampleWorkflowWithMutexResponse) ProtoMessage() {}

func (x *SampleWorkflowWithMutexResponse) ProtoReflect() protoreflect.Message {
	mi := &file_example_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SampleWorkflowWithMutexResponse.ProtoReflect.Descriptor instead.
func (*SampleWorkflowWithMutexResponse) Descriptor() ([]byte, []int) {
	return file_example_proto_rawDescGZIP(), []int{2}
}

func (x *SampleWorkflowWithMutexResponse) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

// AcquireLeaseRequest describes the input to a AcquireLease signal
type AcquireLeaseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WorkflowId string               `protobuf:"bytes,1,opt,name=workflow_id,json=workflowId,proto3" json:"workflow_id,omitempty"`
	Timeout    *durationpb.Duration `protobuf:"bytes,2,opt,name=timeout,proto3" json:"timeout,omitempty"`
}

func (x *AcquireLeaseRequest) Reset() {
	*x = AcquireLeaseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_example_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AcquireLeaseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AcquireLeaseRequest) ProtoMessage() {}

func (x *AcquireLeaseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_example_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AcquireLeaseRequest.ProtoReflect.Descriptor instead.
func (*AcquireLeaseRequest) Descriptor() ([]byte, []int) {
	return file_example_proto_rawDescGZIP(), []int{3}
}

func (x *AcquireLeaseRequest) GetWorkflowId() string {
	if x != nil {
		return x.WorkflowId
	}
	return ""
}

func (x *AcquireLeaseRequest) GetTimeout() *durationpb.Duration {
	if x != nil {
		return x.Timeout
	}
	return nil
}

// LeaseAcquiredRequest describes the input to a LeaseAcquired signal
type LeaseAcquiredRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WorkflowId string `protobuf:"bytes,1,opt,name=workflow_id,json=workflowId,proto3" json:"workflow_id,omitempty"`
	RunId      string `protobuf:"bytes,2,opt,name=run_id,json=runId,proto3" json:"run_id,omitempty"`
	LeaseId    string `protobuf:"bytes,3,opt,name=lease_id,json=leaseId,proto3" json:"lease_id,omitempty"`
}

func (x *LeaseAcquiredRequest) Reset() {
	*x = LeaseAcquiredRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_example_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LeaseAcquiredRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LeaseAcquiredRequest) ProtoMessage() {}

func (x *LeaseAcquiredRequest) ProtoReflect() protoreflect.Message {
	mi := &file_example_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LeaseAcquiredRequest.ProtoReflect.Descriptor instead.
func (*LeaseAcquiredRequest) Descriptor() ([]byte, []int) {
	return file_example_proto_rawDescGZIP(), []int{4}
}

func (x *LeaseAcquiredRequest) GetWorkflowId() string {
	if x != nil {
		return x.WorkflowId
	}
	return ""
}

func (x *LeaseAcquiredRequest) GetRunId() string {
	if x != nil {
		return x.RunId
	}
	return ""
}

func (x *LeaseAcquiredRequest) GetLeaseId() string {
	if x != nil {
		return x.LeaseId
	}
	return ""
}

// RenewLeaseRequest describes the input to a RenewLease signal
type RenewLeaseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LeaseId string               `protobuf:"bytes,1,opt,name=lease_id,json=leaseId,proto3" json:"lease_id,omitempty"`
	Timeout *durationpb.Duration `protobuf:"bytes,2,opt,name=timeout,proto3" json:"timeout,omitempty"`
}

func (x *RenewLeaseRequest) Reset() {
	*x = RenewLeaseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_example_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RenewLeaseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RenewLeaseRequest) ProtoMessage() {}

func (x *RenewLeaseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_example_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RenewLeaseRequest.ProtoReflect.Descriptor instead.
func (*RenewLeaseRequest) Descriptor() ([]byte, []int) {
	return file_example_proto_rawDescGZIP(), []int{5}
}

func (x *RenewLeaseRequest) GetLeaseId() string {
	if x != nil {
		return x.LeaseId
	}
	return ""
}

func (x *RenewLeaseRequest) GetTimeout() *durationpb.Duration {
	if x != nil {
		return x.Timeout
	}
	return nil
}

// RevokeLeaseRequest describes the input to a RevokeLease signal
type RevokeLeaseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LeaseId string `protobuf:"bytes,1,opt,name=lease_id,json=leaseId,proto3" json:"lease_id,omitempty"`
}

func (x *RevokeLeaseRequest) Reset() {
	*x = RevokeLeaseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_example_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RevokeLeaseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RevokeLeaseRequest) ProtoMessage() {}

func (x *RevokeLeaseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_example_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RevokeLeaseRequest.ProtoReflect.Descriptor instead.
func (*RevokeLeaseRequest) Descriptor() ([]byte, []int) {
	return file_example_proto_rawDescGZIP(), []int{6}
}

func (x *RevokeLeaseRequest) GetLeaseId() string {
	if x != nil {
		return x.LeaseId
	}
	return ""
}

var File_example_proto protoreflect.FileDescriptor

var file_example_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x12, 0x6d, 0x79, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2e, 0x6d, 0x75, 0x74, 0x65, 0x78,
	0x2e, 0x76, 0x31, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1a, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x65,
	0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2a, 0x0a, 0x0c,
	0x4d, 0x75, 0x74, 0x65, 0x78, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x22, 0x68, 0x0a, 0x1e, 0x53, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x57, 0x69, 0x74, 0x68, 0x4d, 0x75,
	0x74, 0x65, 0x78, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x65, 0x73, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x22, 0x39, 0x0a, 0x1f, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x57, 0x6f, 0x72, 0x6b,
	0x66, 0x6c, 0x6f, 0x77, 0x57, 0x69, 0x74, 0x68, 0x4d, 0x75, 0x74, 0x65, 0x78, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x6b, 0x0a,
	0x13, 0x41, 0x63, 0x71, 0x75, 0x69, 0x72, 0x65, 0x4c, 0x65, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x77, 0x6f, 0x72, 0x6b, 0x66,
	0x6c, 0x6f, 0x77, 0x49, 0x64, 0x12, 0x33, 0x0a, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x22, 0x69, 0x0a, 0x14, 0x4c, 0x65,
	0x61, 0x73, 0x65, 0x41, 0x63, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f,
	0x77, 0x49, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x72, 0x75, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x72, 0x75, 0x6e, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x6c, 0x65,
	0x61, 0x73, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6c, 0x65,
	0x61, 0x73, 0x65, 0x49, 0x64, 0x22, 0x63, 0x0a, 0x11, 0x52, 0x65, 0x6e, 0x65, 0x77, 0x4c, 0x65,
	0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x6c, 0x65,
	0x61, 0x73, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6c, 0x65,
	0x61, 0x73, 0x65, 0x49, 0x64, 0x12, 0x33, 0x0a, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x22, 0x2f, 0x0a, 0x12, 0x52, 0x65,
	0x76, 0x6f, 0x6b, 0x65, 0x4c, 0x65, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x19, 0x0a, 0x08, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x49, 0x64, 0x32, 0xe5, 0x05, 0x0a, 0x05,
	0x4d, 0x75, 0x74, 0x65, 0x78, 0x12, 0x97, 0x01, 0x0a, 0x05, 0x4d, 0x75, 0x74, 0x65, 0x78, 0x12,
	0x20, 0x2e, 0x6d, 0x79, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2e, 0x6d, 0x75, 0x74, 0x65,
	0x78, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x75, 0x74, 0x65, 0x78, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x54, 0x8a, 0xc4, 0x03, 0x4c, 0x12,
	0x10, 0x0a, 0x0c, 0x41, 0x63, 0x71, 0x75, 0x69, 0x72, 0x65, 0x4c, 0x65, 0x61, 0x73, 0x65, 0x10,
	0x01, 0x12, 0x0c, 0x0a, 0x0a, 0x52, 0x65, 0x6e, 0x65, 0x77, 0x4c, 0x65, 0x61, 0x73, 0x65, 0x12,
	0x0d, 0x0a, 0x0b, 0x52, 0x65, 0x76, 0x6f, 0x6b, 0x65, 0x4c, 0x65, 0x61, 0x73, 0x65, 0x32, 0x1b,
	0x28, 0x01, 0x32, 0x03, 0x08, 0x90, 0x1c, 0x62, 0x12, 0x6d, 0x75, 0x74, 0x65, 0x78, 0x2f, 0x24,
	0x7b, 0x21, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x7d, 0x92, 0xc4, 0x03, 0x00, 0x12,
	0xd9, 0x01, 0x0a, 0x17, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c,
	0x6f, 0x77, 0x57, 0x69, 0x74, 0x68, 0x4d, 0x75, 0x74, 0x65, 0x78, 0x12, 0x32, 0x2e, 0x6d, 0x79,
	0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2e, 0x6d, 0x75, 0x74, 0x65, 0x78, 0x2e, 0x76, 0x31,
	0x2e, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x57,
	0x69, 0x74, 0x68, 0x4d, 0x75, 0x74, 0x65, 0x78, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x33, 0x2e, 0x6d, 0x79, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2e, 0x6d, 0x75, 0x74, 0x65,
	0x78, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x66,
	0x6c, 0x6f, 0x77, 0x57, 0x69, 0x74, 0x68, 0x4d, 0x75, 0x74, 0x65, 0x78, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x55, 0x8a, 0xc4, 0x03, 0x51, 0x12, 0x0f, 0x0a, 0x0d, 0x4c, 0x65,
	0x61, 0x73, 0x65, 0x41, 0x63, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x32, 0x3e, 0x28, 0x02, 0x32,
	0x03, 0x08, 0x90, 0x1c, 0x62, 0x35, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2d, 0x77, 0x6f, 0x72,
	0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x2d, 0x77, 0x69, 0x74, 0x68, 0x2d, 0x6d, 0x75, 0x74, 0x65, 0x78,
	0x2f, 0x24, 0x7b, 0x21, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x7d, 0x2f, 0x24, 0x7b,
	0x21, 0x75, 0x75, 0x69, 0x64, 0x5f, 0x76, 0x34, 0x28, 0x29, 0x7d, 0x12, 0x55, 0x0a, 0x0c, 0x41,
	0x63, 0x71, 0x75, 0x69, 0x72, 0x65, 0x4c, 0x65, 0x61, 0x73, 0x65, 0x12, 0x27, 0x2e, 0x6d, 0x79,
	0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2e, 0x6d, 0x75, 0x74, 0x65, 0x78, 0x2e, 0x76, 0x31,
	0x2e, 0x41, 0x63, 0x71, 0x75, 0x69, 0x72, 0x65, 0x4c, 0x65, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x04, 0xa2, 0xc4,
	0x03, 0x00, 0x12, 0x57, 0x0a, 0x0d, 0x4c, 0x65, 0x61, 0x73, 0x65, 0x41, 0x63, 0x71, 0x75, 0x69,
	0x72, 0x65, 0x64, 0x12, 0x28, 0x2e, 0x6d, 0x79, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2e,
	0x6d, 0x75, 0x74, 0x65, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x65, 0x61, 0x73, 0x65, 0x41, 0x63,
	0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x04, 0xa2, 0xc4, 0x03, 0x00, 0x12, 0x51, 0x0a, 0x0a, 0x52,
	0x65, 0x6e, 0x65, 0x77, 0x4c, 0x65, 0x61, 0x73, 0x65, 0x12, 0x25, 0x2e, 0x6d, 0x79, 0x63, 0x6f,
	0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2e, 0x6d, 0x75, 0x74, 0x65, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x52,
	0x65, 0x6e, 0x65, 0x77, 0x4c, 0x65, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x04, 0xa2, 0xc4, 0x03, 0x00, 0x12, 0x53,
	0x0a, 0x0b, 0x52, 0x65, 0x76, 0x6f, 0x6b, 0x65, 0x4c, 0x65, 0x61, 0x73, 0x65, 0x12, 0x26, 0x2e,
	0x6d, 0x79, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2e, 0x6d, 0x75, 0x74, 0x65, 0x78, 0x2e,
	0x76, 0x31, 0x2e, 0x52, 0x65, 0x76, 0x6f, 0x6b, 0x65, 0x4c, 0x65, 0x61, 0x73, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x04, 0xa2,
	0xc4, 0x03, 0x00, 0x1a, 0x0e, 0x8a, 0xc4, 0x03, 0x0a, 0x0a, 0x08, 0x6d, 0x75, 0x74, 0x65, 0x78,
	0x2d, 0x76, 0x31, 0x42, 0xc7, 0x01, 0x0a, 0x16, 0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x79, 0x63, 0x6f,
	0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2e, 0x6d, 0x75, 0x74, 0x65, 0x78, 0x2e, 0x76, 0x31, 0x42, 0x0c,
	0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x35,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x75, 0x64, 0x64,
	0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f,
	0x2d, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2f, 0x67, 0x65, 0x6e, 0x3b, 0x6d, 0x75,
	0x74, 0x65, 0x78, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x4d, 0x4d, 0x58, 0xaa, 0x02, 0x12, 0x4d, 0x79,
	0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2e, 0x4d, 0x75, 0x74, 0x65, 0x78, 0x2e, 0x56, 0x31,
	0xca, 0x02, 0x12, 0x4d, 0x79, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x5c, 0x4d, 0x75, 0x74,
	0x65, 0x78, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x1e, 0x4d, 0x79, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e,
	0x79, 0x5c, 0x4d, 0x75, 0x74, 0x65, 0x78, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x14, 0x4d, 0x79, 0x63, 0x6f, 0x6d, 0x70, 0x61,
	0x6e, 0x79, 0x3a, 0x3a, 0x4d, 0x75, 0x74, 0x65, 0x78, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_example_proto_rawDescOnce sync.Once
	file_example_proto_rawDescData = file_example_proto_rawDesc
)

func file_example_proto_rawDescGZIP() []byte {
	file_example_proto_rawDescOnce.Do(func() {
		file_example_proto_rawDescData = protoimpl.X.CompressGZIP(file_example_proto_rawDescData)
	})
	return file_example_proto_rawDescData
}

var file_example_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_example_proto_goTypes = []interface{}{
	(*MutexRequest)(nil),                    // 0: mycompany.mutex.v1.MutexRequest
	(*SampleWorkflowWithMutexRequest)(nil),  // 1: mycompany.mutex.v1.SampleWorkflowWithMutexRequest
	(*SampleWorkflowWithMutexResponse)(nil), // 2: mycompany.mutex.v1.SampleWorkflowWithMutexResponse
	(*AcquireLeaseRequest)(nil),             // 3: mycompany.mutex.v1.AcquireLeaseRequest
	(*LeaseAcquiredRequest)(nil),            // 4: mycompany.mutex.v1.LeaseAcquiredRequest
	(*RenewLeaseRequest)(nil),               // 5: mycompany.mutex.v1.RenewLeaseRequest
	(*RevokeLeaseRequest)(nil),              // 6: mycompany.mutex.v1.RevokeLeaseRequest
	(*durationpb.Duration)(nil),             // 7: google.protobuf.Duration
	(*emptypb.Empty)(nil),                   // 8: google.protobuf.Empty
}
var file_example_proto_depIdxs = []int32{
	7, // 0: mycompany.mutex.v1.AcquireLeaseRequest.timeout:type_name -> google.protobuf.Duration
	7, // 1: mycompany.mutex.v1.RenewLeaseRequest.timeout:type_name -> google.protobuf.Duration
	0, // 2: mycompany.mutex.v1.Mutex.Mutex:input_type -> mycompany.mutex.v1.MutexRequest
	1, // 3: mycompany.mutex.v1.Mutex.SampleWorkflowWithMutex:input_type -> mycompany.mutex.v1.SampleWorkflowWithMutexRequest
	3, // 4: mycompany.mutex.v1.Mutex.AcquireLease:input_type -> mycompany.mutex.v1.AcquireLeaseRequest
	4, // 5: mycompany.mutex.v1.Mutex.LeaseAcquired:input_type -> mycompany.mutex.v1.LeaseAcquiredRequest
	5, // 6: mycompany.mutex.v1.Mutex.RenewLease:input_type -> mycompany.mutex.v1.RenewLeaseRequest
	6, // 7: mycompany.mutex.v1.Mutex.RevokeLease:input_type -> mycompany.mutex.v1.RevokeLeaseRequest
	8, // 8: mycompany.mutex.v1.Mutex.Mutex:output_type -> google.protobuf.Empty
	2, // 9: mycompany.mutex.v1.Mutex.SampleWorkflowWithMutex:output_type -> mycompany.mutex.v1.SampleWorkflowWithMutexResponse
	8, // 10: mycompany.mutex.v1.Mutex.AcquireLease:output_type -> google.protobuf.Empty
	8, // 11: mycompany.mutex.v1.Mutex.LeaseAcquired:output_type -> google.protobuf.Empty
	8, // 12: mycompany.mutex.v1.Mutex.RenewLease:output_type -> google.protobuf.Empty
	8, // 13: mycompany.mutex.v1.Mutex.RevokeLease:output_type -> google.protobuf.Empty
	8, // [8:14] is the sub-list for method output_type
	2, // [2:8] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_example_proto_init() }
func file_example_proto_init() {
	if File_example_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_example_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MutexRequest); i {
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
		file_example_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SampleWorkflowWithMutexRequest); i {
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
		file_example_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SampleWorkflowWithMutexResponse); i {
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
		file_example_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AcquireLeaseRequest); i {
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
		file_example_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LeaseAcquiredRequest); i {
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
		file_example_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RenewLeaseRequest); i {
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
		file_example_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RevokeLeaseRequest); i {
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
			RawDescriptor: file_example_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_example_proto_goTypes,
		DependencyIndexes: file_example_proto_depIdxs,
		MessageInfos:      file_example_proto_msgTypes,
	}.Build()
	File_example_proto = out.File
	file_example_proto_rawDesc = nil
	file_example_proto_goTypes = nil
	file_example_proto_depIdxs = nil
}
