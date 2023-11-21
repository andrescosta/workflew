// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.0
// source: control.proto

package types

import (
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

type GetQueuesDefRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	QueueId    *QueueId    `protobuf:"bytes,1,opt,name=queueId,proto3,oneof" json:"queueId,omitempty"`
	MerchantId *MerchantId `protobuf:"bytes,2,opt,name=merchantId,proto3,oneof" json:"merchantId,omitempty"`
}

func (x *GetQueuesDefRequest) Reset() {
	*x = GetQueuesDefRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_control_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetQueuesDefRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetQueuesDefRequest) ProtoMessage() {}

func (x *GetQueuesDefRequest) ProtoReflect() protoreflect.Message {
	mi := &file_control_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetQueuesDefRequest.ProtoReflect.Descriptor instead.
func (*GetQueuesDefRequest) Descriptor() ([]byte, []int) {
	return file_control_proto_rawDescGZIP(), []int{0}
}

func (x *GetQueuesDefRequest) GetQueueId() *QueueId {
	if x != nil {
		return x.QueueId
	}
	return nil
}

func (x *GetQueuesDefRequest) GetMerchantId() *MerchantId {
	if x != nil {
		return x.MerchantId
	}
	return nil
}

type GetQueuesDefReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Queues []*QueueDef `protobuf:"bytes,1,rep,name=queues,proto3" json:"queues,omitempty"`
}

func (x *GetQueuesDefReply) Reset() {
	*x = GetQueuesDefReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_control_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetQueuesDefReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetQueuesDefReply) ProtoMessage() {}

func (x *GetQueuesDefReply) ProtoReflect() protoreflect.Message {
	mi := &file_control_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetQueuesDefReply.ProtoReflect.Descriptor instead.
func (*GetQueuesDefReply) Descriptor() ([]byte, []int) {
	return file_control_proto_rawDescGZIP(), []int{1}
}

func (x *GetQueuesDefReply) GetQueues() []*QueueDef {
	if x != nil {
		return x.Queues
	}
	return nil
}

type AddQueueDefRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Queues *QueueDef `protobuf:"bytes,1,opt,name=queues,proto3" json:"queues,omitempty"`
}

func (x *AddQueueDefRequest) Reset() {
	*x = AddQueueDefRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_control_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddQueueDefRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddQueueDefRequest) ProtoMessage() {}

func (x *AddQueueDefRequest) ProtoReflect() protoreflect.Message {
	mi := &file_control_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddQueueDefRequest.ProtoReflect.Descriptor instead.
func (*AddQueueDefRequest) Descriptor() ([]byte, []int) {
	return file_control_proto_rawDescGZIP(), []int{2}
}

func (x *AddQueueDefRequest) GetQueues() *QueueDef {
	if x != nil {
		return x.Queues
	}
	return nil
}

type AddQueueDefReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Queues *QueueDef `protobuf:"bytes,1,opt,name=queues,proto3" json:"queues,omitempty"`
}

func (x *AddQueueDefReply) Reset() {
	*x = AddQueueDefReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_control_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddQueueDefReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddQueueDefReply) ProtoMessage() {}

func (x *AddQueueDefReply) ProtoReflect() protoreflect.Message {
	mi := &file_control_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddQueueDefReply.ProtoReflect.Descriptor instead.
func (*AddQueueDefReply) Descriptor() ([]byte, []int) {
	return file_control_proto_rawDescGZIP(), []int{3}
}

func (x *AddQueueDefReply) GetQueues() *QueueDef {
	if x != nil {
		return x.Queues
	}
	return nil
}

type QueueDef struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID         string      `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	QueueId    *QueueId    `protobuf:"bytes,2,opt,name=queueId,proto3" json:"queueId,omitempty"`
	MerchantId *MerchantId `protobuf:"bytes,3,opt,name=merchantId,proto3" json:"merchantId,omitempty"`
	Executor   *Executor   `protobuf:"bytes,4,opt,name=executor,proto3" json:"executor,omitempty"`
}

func (x *QueueDef) Reset() {
	*x = QueueDef{}
	if protoimpl.UnsafeEnabled {
		mi := &file_control_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueueDef) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueueDef) ProtoMessage() {}

func (x *QueueDef) ProtoReflect() protoreflect.Message {
	mi := &file_control_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueueDef.ProtoReflect.Descriptor instead.
func (*QueueDef) Descriptor() ([]byte, []int) {
	return file_control_proto_rawDescGZIP(), []int{4}
}

func (x *QueueDef) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *QueueDef) GetQueueId() *QueueId {
	if x != nil {
		return x.QueueId
	}
	return nil
}

func (x *QueueDef) GetMerchantId() *MerchantId {
	if x != nil {
		return x.MerchantId
	}
	return nil
}

func (x *QueueDef) GetExecutor() *Executor {
	if x != nil {
		return x.Executor
	}
	return nil
}

type Executor struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Type   string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Script string `protobuf:"bytes,3,opt,name=script,proto3" json:"script,omitempty"`
}

func (x *Executor) Reset() {
	*x = Executor{}
	if protoimpl.UnsafeEnabled {
		mi := &file_control_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Executor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Executor) ProtoMessage() {}

func (x *Executor) ProtoReflect() protoreflect.Message {
	mi := &file_control_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Executor.ProtoReflect.Descriptor instead.
func (*Executor) Descriptor() ([]byte, []int) {
	return file_control_proto_rawDescGZIP(), []int{5}
}

func (x *Executor) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Executor) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Executor) GetScript() string {
	if x != nil {
		return x.Script
	}
	return ""
}

type Event struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Schema string `protobuf:"bytes,2,opt,name=schema,proto3" json:"schema,omitempty"`
}

func (x *Event) Reset() {
	*x = Event{}
	if protoimpl.UnsafeEnabled {
		mi := &file_control_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event) ProtoMessage() {}

func (x *Event) ProtoReflect() protoreflect.Message {
	mi := &file_control_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Event.ProtoReflect.Descriptor instead.
func (*Event) Descriptor() ([]byte, []int) {
	return file_control_proto_rawDescGZIP(), []int{6}
}

func (x *Event) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Event) GetSchema() string {
	if x != nil {
		return x.Schema
	}
	return ""
}

type ExecutorDefinition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Event   *Event `protobuf:"bytes,1,opt,name=event,proto3" json:"event,omitempty"`
	Package string `protobuf:"bytes,2,opt,name=package,proto3" json:"package,omitempty"`
}

func (x *ExecutorDefinition) Reset() {
	*x = ExecutorDefinition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_control_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExecutorDefinition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExecutorDefinition) ProtoMessage() {}

func (x *ExecutorDefinition) ProtoReflect() protoreflect.Message {
	mi := &file_control_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExecutorDefinition.ProtoReflect.Descriptor instead.
func (*ExecutorDefinition) Descriptor() ([]byte, []int) {
	return file_control_proto_rawDescGZIP(), []int{7}
}

func (x *ExecutorDefinition) GetEvent() *Event {
	if x != nil {
		return x.Event
	}
	return nil
}

func (x *ExecutorDefinition) GetPackage() string {
	if x != nil {
		return x.Package
	}
	return ""
}

type GetEventRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Events []*Event `protobuf:"bytes,1,rep,name=events,proto3" json:"events,omitempty"`
}

func (x *GetEventRequest) Reset() {
	*x = GetEventRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_control_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEventRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEventRequest) ProtoMessage() {}

func (x *GetEventRequest) ProtoReflect() protoreflect.Message {
	mi := &file_control_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEventRequest.ProtoReflect.Descriptor instead.
func (*GetEventRequest) Descriptor() ([]byte, []int) {
	return file_control_proto_rawDescGZIP(), []int{8}
}

func (x *GetEventRequest) GetEvents() []*Event {
	if x != nil {
		return x.Events
	}
	return nil
}

type GetEventReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Events []*Event `protobuf:"bytes,1,rep,name=events,proto3" json:"events,omitempty"`
}

func (x *GetEventReply) Reset() {
	*x = GetEventReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_control_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEventReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEventReply) ProtoMessage() {}

func (x *GetEventReply) ProtoReflect() protoreflect.Message {
	mi := &file_control_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEventReply.ProtoReflect.Descriptor instead.
func (*GetEventReply) Descriptor() ([]byte, []int) {
	return file_control_proto_rawDescGZIP(), []int{9}
}

func (x *GetEventReply) GetEvents() []*Event {
	if x != nil {
		return x.Events
	}
	return nil
}

type AddEventRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Event []*Event `protobuf:"bytes,1,rep,name=event,proto3" json:"event,omitempty"`
}

func (x *AddEventRequest) Reset() {
	*x = AddEventRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_control_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddEventRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddEventRequest) ProtoMessage() {}

func (x *AddEventRequest) ProtoReflect() protoreflect.Message {
	mi := &file_control_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddEventRequest.ProtoReflect.Descriptor instead.
func (*AddEventRequest) Descriptor() ([]byte, []int) {
	return file_control_proto_rawDescGZIP(), []int{10}
}

func (x *AddEventRequest) GetEvent() []*Event {
	if x != nil {
		return x.Event
	}
	return nil
}

type AddEventReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AddEventReply) Reset() {
	*x = AddEventReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_control_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddEventReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddEventReply) ProtoMessage() {}

func (x *AddEventReply) ProtoReflect() protoreflect.Message {
	mi := &file_control_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddEventReply.ProtoReflect.Descriptor instead.
func (*AddEventReply) Descriptor() ([]byte, []int) {
	return file_control_proto_rawDescGZIP(), []int{11}
}

type AddExecutorRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Definition []*ExecutorDefinition `protobuf:"bytes,1,rep,name=definition,proto3" json:"definition,omitempty"`
}

func (x *AddExecutorRequest) Reset() {
	*x = AddExecutorRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_control_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddExecutorRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddExecutorRequest) ProtoMessage() {}

func (x *AddExecutorRequest) ProtoReflect() protoreflect.Message {
	mi := &file_control_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddExecutorRequest.ProtoReflect.Descriptor instead.
func (*AddExecutorRequest) Descriptor() ([]byte, []int) {
	return file_control_proto_rawDescGZIP(), []int{12}
}

func (x *AddExecutorRequest) GetDefinition() []*ExecutorDefinition {
	if x != nil {
		return x.Definition
	}
	return nil
}

type AddExecutorReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AddExecutorReply) Reset() {
	*x = AddExecutorReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_control_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddExecutorReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddExecutorReply) ProtoMessage() {}

func (x *AddExecutorReply) ProtoReflect() protoreflect.Message {
	mi := &file_control_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddExecutorReply.ProtoReflect.Descriptor instead.
func (*AddExecutorReply) Descriptor() ([]byte, []int) {
	return file_control_proto_rawDescGZIP(), []int{13}
}

type GetExecutorsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Definitions []*ExecutorDefinition `protobuf:"bytes,1,rep,name=definitions,proto3" json:"definitions,omitempty"`
}

func (x *GetExecutorsRequest) Reset() {
	*x = GetExecutorsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_control_proto_msgTypes[14]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetExecutorsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetExecutorsRequest) ProtoMessage() {}

func (x *GetExecutorsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_control_proto_msgTypes[14]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetExecutorsRequest.ProtoReflect.Descriptor instead.
func (*GetExecutorsRequest) Descriptor() ([]byte, []int) {
	return file_control_proto_rawDescGZIP(), []int{14}
}

func (x *GetExecutorsRequest) GetDefinitions() []*ExecutorDefinition {
	if x != nil {
		return x.Definitions
	}
	return nil
}

type GetExecutorsReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Definitions []*ExecutorDefinition `protobuf:"bytes,1,rep,name=definitions,proto3" json:"definitions,omitempty"`
}

func (x *GetExecutorsReply) Reset() {
	*x = GetExecutorsReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_control_proto_msgTypes[15]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetExecutorsReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetExecutorsReply) ProtoMessage() {}

func (x *GetExecutorsReply) ProtoReflect() protoreflect.Message {
	mi := &file_control_proto_msgTypes[15]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetExecutorsReply.ProtoReflect.Descriptor instead.
func (*GetExecutorsReply) Descriptor() ([]byte, []int) {
	return file_control_proto_rawDescGZIP(), []int{15}
}

func (x *GetExecutorsReply) GetDefinitions() []*ExecutorDefinition {
	if x != nil {
		return x.Definitions
	}
	return nil
}

var File_control_proto protoreflect.FileDescriptor

var file_control_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8b, 0x01,
	0x0a, 0x13, 0x47, 0x65, 0x74, 0x51, 0x75, 0x65, 0x75, 0x65, 0x73, 0x44, 0x65, 0x66, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x27, 0x0a, 0x07, 0x71, 0x75, 0x65, 0x75, 0x65, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x51, 0x75, 0x65, 0x75, 0x65, 0x49, 0x64,
	0x48, 0x00, 0x52, 0x07, 0x71, 0x75, 0x65, 0x75, 0x65, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x30,
	0x0a, 0x0a, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x4d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x48,
	0x01, 0x52, 0x0a, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x88, 0x01, 0x01,
	0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x71, 0x75, 0x65, 0x75, 0x65, 0x49, 0x64, 0x42, 0x0d, 0x0a, 0x0b,
	0x5f, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x36, 0x0a, 0x11, 0x47,
	0x65, 0x74, 0x51, 0x75, 0x65, 0x75, 0x65, 0x73, 0x44, 0x65, 0x66, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x12, 0x21, 0x0a, 0x06, 0x71, 0x75, 0x65, 0x75, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x09, 0x2e, 0x51, 0x75, 0x65, 0x75, 0x65, 0x44, 0x65, 0x66, 0x52, 0x06, 0x71, 0x75, 0x65,
	0x75, 0x65, 0x73, 0x22, 0x37, 0x0a, 0x12, 0x41, 0x64, 0x64, 0x51, 0x75, 0x65, 0x75, 0x65, 0x44,
	0x65, 0x66, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x06, 0x71, 0x75, 0x65,
	0x75, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x51, 0x75, 0x65, 0x75,
	0x65, 0x44, 0x65, 0x66, 0x52, 0x06, 0x71, 0x75, 0x65, 0x75, 0x65, 0x73, 0x22, 0x35, 0x0a, 0x10,
	0x41, 0x64, 0x64, 0x51, 0x75, 0x65, 0x75, 0x65, 0x44, 0x65, 0x66, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x12, 0x21, 0x0a, 0x06, 0x71, 0x75, 0x65, 0x75, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x09, 0x2e, 0x51, 0x75, 0x65, 0x75, 0x65, 0x44, 0x65, 0x66, 0x52, 0x06, 0x71, 0x75, 0x65,
	0x75, 0x65, 0x73, 0x22, 0x92, 0x01, 0x0a, 0x08, 0x51, 0x75, 0x65, 0x75, 0x65, 0x44, 0x65, 0x66,
	0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44,
	0x12, 0x22, 0x0a, 0x07, 0x71, 0x75, 0x65, 0x75, 0x65, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x08, 0x2e, 0x51, 0x75, 0x65, 0x75, 0x65, 0x49, 0x64, 0x52, 0x07, 0x71, 0x75, 0x65,
	0x75, 0x65, 0x49, 0x64, 0x12, 0x2b, 0x0a, 0x0a, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74,
	0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x4d, 0x65, 0x72, 0x63, 0x68,
	0x61, 0x6e, 0x74, 0x49, 0x64, 0x52, 0x0a, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x49,
	0x64, 0x12, 0x25, 0x0a, 0x08, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x52, 0x08,
	0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x22, 0x46, 0x0a, 0x08, 0x45, 0x78, 0x65, 0x63,
	0x75, 0x74, 0x6f, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x22, 0x2f, 0x0a, 0x05, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x63, 0x68,
	0x65, 0x6d, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x63, 0x68, 0x65, 0x6d,
	0x61, 0x22, 0x4c, 0x0a, 0x12, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x44, 0x65, 0x66,
	0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x06, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x05,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x22,
	0x31, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1e, 0x0a, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x06, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x06, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x73, 0x22, 0x2f, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x12, 0x1e, 0x0a, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x06, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x06, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x73, 0x22, 0x2f, 0x0a, 0x0f, 0x41, 0x64, 0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x06, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x05, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x22, 0x0f, 0x0a, 0x0d, 0x41, 0x64, 0x64, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x49, 0x0a, 0x12, 0x41, 0x64, 0x64, 0x45, 0x78, 0x65, 0x63,
	0x75, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x33, 0x0a, 0x0a, 0x64,
	0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x13, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0x12, 0x0a, 0x10, 0x41, 0x64, 0x64, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x22, 0x4c, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x45, 0x78, 0x65, 0x63, 0x75,
	0x74, 0x6f, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x35, 0x0a, 0x0b, 0x64,
	0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x13, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x44, 0x65, 0x66, 0x69, 0x6e,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x22, 0x4a, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f,
	0x72, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x35, 0x0a, 0x0b, 0x64, 0x65, 0x66, 0x69, 0x6e,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x45,
	0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x0b, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x32, 0xd0,
	0x02, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x12, 0x2f, 0x0a, 0x09, 0x47, 0x65,
	0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x10, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x47, 0x65, 0x74, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x2f, 0x0a, 0x09, 0x41,
	0x64, 0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x10, 0x2e, 0x41, 0x64, 0x64, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x41, 0x64, 0x64,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x0b,
	0x41, 0x64, 0x64, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x12, 0x13, 0x2e, 0x41, 0x64,
	0x64, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x11, 0x2e, 0x41, 0x64, 0x64, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x3a, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x45, 0x78, 0x65, 0x63,
	0x75, 0x74, 0x6f, 0x72, 0x73, 0x12, 0x14, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x78, 0x65, 0x63, 0x75,
	0x74, 0x6f, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x47, 0x65,
	0x74, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0x00, 0x12, 0x37, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x51, 0x75, 0x65, 0x75, 0x65, 0x73, 0x12, 0x14,
	0x2e, 0x47, 0x65, 0x74, 0x51, 0x75, 0x65, 0x75, 0x65, 0x73, 0x44, 0x65, 0x66, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x47, 0x65, 0x74, 0x51, 0x75, 0x65, 0x75, 0x65, 0x73,
	0x44, 0x65, 0x66, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x35, 0x0a, 0x09, 0x41, 0x64,
	0x64, 0x51, 0x75, 0x65, 0x75, 0x65, 0x73, 0x12, 0x13, 0x2e, 0x41, 0x64, 0x64, 0x51, 0x75, 0x65,
	0x75, 0x65, 0x44, 0x65, 0x66, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x41,
	0x64, 0x64, 0x51, 0x75, 0x65, 0x75, 0x65, 0x44, 0x65, 0x66, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0x00, 0x42, 0x08, 0x5a, 0x06, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_control_proto_rawDescOnce sync.Once
	file_control_proto_rawDescData = file_control_proto_rawDesc
)

func file_control_proto_rawDescGZIP() []byte {
	file_control_proto_rawDescOnce.Do(func() {
		file_control_proto_rawDescData = protoimpl.X.CompressGZIP(file_control_proto_rawDescData)
	})
	return file_control_proto_rawDescData
}

var file_control_proto_msgTypes = make([]protoimpl.MessageInfo, 16)
var file_control_proto_goTypes = []interface{}{
	(*GetQueuesDefRequest)(nil), // 0: GetQueuesDefRequest
	(*GetQueuesDefReply)(nil),   // 1: GetQueuesDefReply
	(*AddQueueDefRequest)(nil),  // 2: AddQueueDefRequest
	(*AddQueueDefReply)(nil),    // 3: AddQueueDefReply
	(*QueueDef)(nil),            // 4: QueueDef
	(*Executor)(nil),            // 5: Executor
	(*Event)(nil),               // 6: Event
	(*ExecutorDefinition)(nil),  // 7: ExecutorDefinition
	(*GetEventRequest)(nil),     // 8: GetEventRequest
	(*GetEventReply)(nil),       // 9: GetEventReply
	(*AddEventRequest)(nil),     // 10: AddEventRequest
	(*AddEventReply)(nil),       // 11: AddEventReply
	(*AddExecutorRequest)(nil),  // 12: AddExecutorRequest
	(*AddExecutorReply)(nil),    // 13: AddExecutorReply
	(*GetExecutorsRequest)(nil), // 14: GetExecutorsRequest
	(*GetExecutorsReply)(nil),   // 15: GetExecutorsReply
	(*QueueId)(nil),             // 16: QueueId
	(*MerchantId)(nil),          // 17: MerchantId
}
var file_control_proto_depIdxs = []int32{
	16, // 0: GetQueuesDefRequest.queueId:type_name -> QueueId
	17, // 1: GetQueuesDefRequest.merchantId:type_name -> MerchantId
	4,  // 2: GetQueuesDefReply.queues:type_name -> QueueDef
	4,  // 3: AddQueueDefRequest.queues:type_name -> QueueDef
	4,  // 4: AddQueueDefReply.queues:type_name -> QueueDef
	16, // 5: QueueDef.queueId:type_name -> QueueId
	17, // 6: QueueDef.merchantId:type_name -> MerchantId
	5,  // 7: QueueDef.executor:type_name -> Executor
	6,  // 8: ExecutorDefinition.event:type_name -> Event
	6,  // 9: GetEventRequest.events:type_name -> Event
	6,  // 10: GetEventReply.events:type_name -> Event
	6,  // 11: AddEventRequest.event:type_name -> Event
	7,  // 12: AddExecutorRequest.definition:type_name -> ExecutorDefinition
	7,  // 13: GetExecutorsRequest.definitions:type_name -> ExecutorDefinition
	7,  // 14: GetExecutorsReply.definitions:type_name -> ExecutorDefinition
	8,  // 15: Control.GetEvents:input_type -> GetEventRequest
	10, // 16: Control.AddEvents:input_type -> AddEventRequest
	12, // 17: Control.AddExecutor:input_type -> AddExecutorRequest
	14, // 18: Control.GetExecutors:input_type -> GetExecutorsRequest
	0,  // 19: Control.GetQueues:input_type -> GetQueuesDefRequest
	2,  // 20: Control.AddQueues:input_type -> AddQueueDefRequest
	9,  // 21: Control.GetEvents:output_type -> GetEventReply
	11, // 22: Control.AddEvents:output_type -> AddEventReply
	13, // 23: Control.AddExecutor:output_type -> AddExecutorReply
	15, // 24: Control.GetExecutors:output_type -> GetExecutorsReply
	1,  // 25: Control.GetQueues:output_type -> GetQueuesDefReply
	3,  // 26: Control.AddQueues:output_type -> AddQueueDefReply
	21, // [21:27] is the sub-list for method output_type
	15, // [15:21] is the sub-list for method input_type
	15, // [15:15] is the sub-list for extension type_name
	15, // [15:15] is the sub-list for extension extendee
	0,  // [0:15] is the sub-list for field type_name
}

func init() { file_control_proto_init() }
func file_control_proto_init() {
	if File_control_proto != nil {
		return
	}
	file_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_control_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetQueuesDefRequest); i {
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
		file_control_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetQueuesDefReply); i {
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
		file_control_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddQueueDefRequest); i {
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
		file_control_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddQueueDefReply); i {
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
		file_control_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueueDef); i {
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
		file_control_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Executor); i {
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
		file_control_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Event); i {
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
		file_control_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExecutorDefinition); i {
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
		file_control_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEventRequest); i {
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
		file_control_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEventReply); i {
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
		file_control_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddEventRequest); i {
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
		file_control_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddEventReply); i {
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
		file_control_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddExecutorRequest); i {
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
		file_control_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddExecutorReply); i {
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
		file_control_proto_msgTypes[14].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetExecutorsRequest); i {
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
		file_control_proto_msgTypes[15].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetExecutorsReply); i {
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
	file_control_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_control_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   16,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_control_proto_goTypes,
		DependencyIndexes: file_control_proto_depIdxs,
		MessageInfos:      file_control_proto_msgTypes,
	}.Build()
	File_control_proto = out.File
	file_control_proto_rawDesc = nil
	file_control_proto_goTypes = nil
	file_control_proto_depIdxs = nil
}
