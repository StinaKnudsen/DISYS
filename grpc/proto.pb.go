// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.2
// source: proto.proto

package grpc

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

type ChatMessageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ParticipantId    string `protobuf:"bytes,1,opt,name=participant_id,json=participantId,proto3" json:"participant_id,omitempty"`
	Message          string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	LogicalTimestamp int64  `protobuf:"varint,3,opt,name=logical_timestamp,json=logicalTimestamp,proto3" json:"logical_timestamp,omitempty"`
}

func (x *ChatMessageRequest) Reset() {
	*x = ChatMessageRequest{}
	mi := &file_proto_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ChatMessageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatMessageRequest) ProtoMessage() {}

func (x *ChatMessageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatMessageRequest.ProtoReflect.Descriptor instead.
func (*ChatMessageRequest) Descriptor() ([]byte, []int) {
	return file_proto_proto_rawDescGZIP(), []int{0}
}

func (x *ChatMessageRequest) GetParticipantId() string {
	if x != nil {
		return x.ParticipantId
	}
	return ""
}

func (x *ChatMessageRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *ChatMessageRequest) GetLogicalTimestamp() int64 {
	if x != nil {
		return x.LogicalTimestamp
	}
	return 0
}

type PublishResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *PublishResponse) Reset() {
	*x = PublishResponse{}
	mi := &file_proto_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PublishResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublishResponse) ProtoMessage() {}

func (x *PublishResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublishResponse.ProtoReflect.Descriptor instead.
func (*PublishResponse) Descriptor() ([]byte, []int) {
	return file_proto_proto_rawDescGZIP(), []int{1}
}

func (x *PublishResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type BroadcastMessageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ParticipantId    string `protobuf:"bytes,1,opt,name=participant_id,json=participantId,proto3" json:"participant_id,omitempty"`
	Message          string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	LogicalTimestamp int64  `protobuf:"varint,3,opt,name=logical_timestamp,json=logicalTimestamp,proto3" json:"logical_timestamp,omitempty"`
}

func (x *BroadcastMessageRequest) Reset() {
	*x = BroadcastMessageRequest{}
	mi := &file_proto_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BroadcastMessageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BroadcastMessageRequest) ProtoMessage() {}

func (x *BroadcastMessageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BroadcastMessageRequest.ProtoReflect.Descriptor instead.
func (*BroadcastMessageRequest) Descriptor() ([]byte, []int) {
	return file_proto_proto_rawDescGZIP(), []int{2}
}

func (x *BroadcastMessageRequest) GetParticipantId() string {
	if x != nil {
		return x.ParticipantId
	}
	return ""
}

func (x *BroadcastMessageRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *BroadcastMessageRequest) GetLogicalTimestamp() int64 {
	if x != nil {
		return x.LogicalTimestamp
	}
	return 0
}

type BroadcastResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *BroadcastResponse) Reset() {
	*x = BroadcastResponse{}
	mi := &file_proto_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BroadcastResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BroadcastResponse) ProtoMessage() {}

func (x *BroadcastResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BroadcastResponse.ProtoReflect.Descriptor instead.
func (*BroadcastResponse) Descriptor() ([]byte, []int) {
	return file_proto_proto_rawDescGZIP(), []int{3}
}

func (x *BroadcastResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type JoinRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ParticipantId    string `protobuf:"bytes,1,opt,name=participant_id,json=participantId,proto3" json:"participant_id,omitempty"`
	LogicalTimestamp int64  `protobuf:"varint,2,opt,name=logical_timestamp,json=logicalTimestamp,proto3" json:"logical_timestamp,omitempty"`
}

func (x *JoinRequest) Reset() {
	*x = JoinRequest{}
	mi := &file_proto_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *JoinRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinRequest) ProtoMessage() {}

func (x *JoinRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoinRequest.ProtoReflect.Descriptor instead.
func (*JoinRequest) Descriptor() ([]byte, []int) {
	return file_proto_proto_rawDescGZIP(), []int{4}
}

func (x *JoinRequest) GetParticipantId() string {
	if x != nil {
		return x.ParticipantId
	}
	return ""
}

func (x *JoinRequest) GetLogicalTimestamp() int64 {
	if x != nil {
		return x.LogicalTimestamp
	}
	return 0
}

type JoinResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WelcomeMessage   string `protobuf:"bytes,1,opt,name=welcome_message,json=welcomeMessage,proto3" json:"welcome_message,omitempty"`
	LogicalTimestamp int64  `protobuf:"varint,2,opt,name=logical_timestamp,json=logicalTimestamp,proto3" json:"logical_timestamp,omitempty"`
}

func (x *JoinResponse) Reset() {
	*x = JoinResponse{}
	mi := &file_proto_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *JoinResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinResponse) ProtoMessage() {}

func (x *JoinResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoinResponse.ProtoReflect.Descriptor instead.
func (*JoinResponse) Descriptor() ([]byte, []int) {
	return file_proto_proto_rawDescGZIP(), []int{5}
}

func (x *JoinResponse) GetWelcomeMessage() string {
	if x != nil {
		return x.WelcomeMessage
	}
	return ""
}

func (x *JoinResponse) GetLogicalTimestamp() int64 {
	if x != nil {
		return x.LogicalTimestamp
	}
	return 0
}

type LeaveRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ParticipantId    string `protobuf:"bytes,1,opt,name=participant_id,json=participantId,proto3" json:"participant_id,omitempty"`
	LogicalTimestamp int64  `protobuf:"varint,2,opt,name=logical_timestamp,json=logicalTimestamp,proto3" json:"logical_timestamp,omitempty"`
}

func (x *LeaveRequest) Reset() {
	*x = LeaveRequest{}
	mi := &file_proto_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LeaveRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LeaveRequest) ProtoMessage() {}

func (x *LeaveRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LeaveRequest.ProtoReflect.Descriptor instead.
func (*LeaveRequest) Descriptor() ([]byte, []int) {
	return file_proto_proto_rawDescGZIP(), []int{6}
}

func (x *LeaveRequest) GetParticipantId() string {
	if x != nil {
		return x.ParticipantId
	}
	return ""
}

func (x *LeaveRequest) GetLogicalTimestamp() int64 {
	if x != nil {
		return x.LogicalTimestamp
	}
	return 0
}

type LeaveResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GoodbyeMessage   string `protobuf:"bytes,1,opt,name=goodbye_message,json=goodbyeMessage,proto3" json:"goodbye_message,omitempty"`
	LogicalTimestamp int64  `protobuf:"varint,2,opt,name=logical_timestamp,json=logicalTimestamp,proto3" json:"logical_timestamp,omitempty"`
}

func (x *LeaveResponse) Reset() {
	*x = LeaveResponse{}
	mi := &file_proto_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LeaveResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LeaveResponse) ProtoMessage() {}

func (x *LeaveResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LeaveResponse.ProtoReflect.Descriptor instead.
func (*LeaveResponse) Descriptor() ([]byte, []int) {
	return file_proto_proto_rawDescGZIP(), []int{7}
}

func (x *LeaveResponse) GetGoodbyeMessage() string {
	if x != nil {
		return x.GoodbyeMessage
	}
	return ""
}

func (x *LeaveResponse) GetLogicalTimestamp() int64 {
	if x != nil {
		return x.LogicalTimestamp
	}
	return 0
}

type ListenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ParticipantId string `protobuf:"bytes,1,opt,name=participant_id,json=participantId,proto3" json:"participant_id,omitempty"`
}

func (x *ListenRequest) Reset() {
	*x = ListenRequest{}
	mi := &file_proto_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListenRequest) ProtoMessage() {}

func (x *ListenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListenRequest.ProtoReflect.Descriptor instead.
func (*ListenRequest) Descriptor() ([]byte, []int) {
	return file_proto_proto_rawDescGZIP(), []int{8}
}

func (x *ListenRequest) GetParticipantId() string {
	if x != nil {
		return x.ParticipantId
	}
	return ""
}

var File_proto_proto protoreflect.FileDescriptor

var file_proto_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x43,
	0x68, 0x69, 0x74, 0x74, 0x79, 0x43, 0x68, 0x61, 0x74, 0x74, 0x79, 0x22, 0x82, 0x01, 0x0a, 0x12,
	0x43, 0x68, 0x61, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x70, 0x61, 0x72, 0x74,
	0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x2b, 0x0a, 0x11, 0x6c, 0x6f, 0x67, 0x69, 0x63, 0x61, 0x6c, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x10,
	0x6c, 0x6f, 0x67, 0x69, 0x63, 0x61, 0x6c, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x22, 0x2b, 0x0a, 0x0f, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x87, 0x01,
	0x0a, 0x17, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x70, 0x61, 0x72,
	0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x49, 0x64,
	0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x2b, 0x0a, 0x11, 0x6c, 0x6f,
	0x67, 0x69, 0x63, 0x61, 0x6c, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x10, 0x6c, 0x6f, 0x67, 0x69, 0x63, 0x61, 0x6c, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0x2d, 0x0a, 0x11, 0x42, 0x72, 0x6f, 0x61, 0x64,
	0x63, 0x61, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x61, 0x0a, 0x0b, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69,
	0x70, 0x61, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x70,
	0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x2b, 0x0a, 0x11,
	0x6c, 0x6f, 0x67, 0x69, 0x63, 0x61, 0x6c, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x10, 0x6c, 0x6f, 0x67, 0x69, 0x63, 0x61, 0x6c,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0x64, 0x0a, 0x0c, 0x4a, 0x6f, 0x69,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x77, 0x65, 0x6c,
	0x63, 0x6f, 0x6d, 0x65, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0e, 0x77, 0x65, 0x6c, 0x63, 0x6f, 0x6d, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x2b, 0x0a, 0x11, 0x6c, 0x6f, 0x67, 0x69, 0x63, 0x61, 0x6c, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x10, 0x6c,
	0x6f, 0x67, 0x69, 0x63, 0x61, 0x6c, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22,
	0x62, 0x0a, 0x0c, 0x4c, 0x65, 0x61, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x25, 0x0a, 0x0e, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69,
	0x70, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x2b, 0x0a, 0x11, 0x6c, 0x6f, 0x67, 0x69, 0x63, 0x61,
	0x6c, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x10, 0x6c, 0x6f, 0x67, 0x69, 0x63, 0x61, 0x6c, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x22, 0x65, 0x0a, 0x0d, 0x4c, 0x65, 0x61, 0x76, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x67, 0x6f, 0x6f, 0x64, 0x62, 0x79, 0x65, 0x5f,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x67,
	0x6f, 0x6f, 0x64, 0x62, 0x79, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x2b, 0x0a,
	0x11, 0x6c, 0x6f, 0x67, 0x69, 0x63, 0x61, 0x6c, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x10, 0x6c, 0x6f, 0x67, 0x69, 0x63, 0x61,
	0x6c, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0x36, 0x0a, 0x0d, 0x4c, 0x69,
	0x73, 0x74, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x70,
	0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0d, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74,
	0x49, 0x64, 0x32, 0x9f, 0x03, 0x0a, 0x13, 0x43, 0x68, 0x69, 0x74, 0x74, 0x79, 0x43, 0x68, 0x61,
	0x74, 0x74, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x51, 0x0a, 0x0e, 0x50, 0x75,
	0x62, 0x6c, 0x69, 0x73, 0x68, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x20, 0x2e, 0x43,
	0x68, 0x69, 0x74, 0x74, 0x79, 0x43, 0x68, 0x61, 0x74, 0x74, 0x79, 0x2e, 0x43, 0x68, 0x61, 0x74,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d,
	0x2e, 0x43, 0x68, 0x69, 0x74, 0x74, 0x79, 0x43, 0x68, 0x61, 0x74, 0x74, 0x79, 0x2e, 0x50, 0x75,
	0x62, 0x6c, 0x69, 0x73, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5a, 0x0a,
	0x10, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x25, 0x2e, 0x43, 0x68, 0x69, 0x74, 0x74, 0x79, 0x43, 0x68, 0x61, 0x74, 0x74, 0x79,
	0x2e, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x43, 0x68, 0x69, 0x74, 0x74,
	0x79, 0x43, 0x68, 0x61, 0x74, 0x74, 0x79, 0x2e, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a, 0x04, 0x4a, 0x6f, 0x69,
	0x6e, 0x12, 0x19, 0x2e, 0x43, 0x68, 0x69, 0x74, 0x74, 0x79, 0x43, 0x68, 0x61, 0x74, 0x74, 0x79,
	0x2e, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x43,
	0x68, 0x69, 0x74, 0x74, 0x79, 0x43, 0x68, 0x61, 0x74, 0x74, 0x79, 0x2e, 0x4a, 0x6f, 0x69, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x40, 0x0a, 0x05, 0x4c, 0x65, 0x61, 0x76,
	0x65, 0x12, 0x1a, 0x2e, 0x43, 0x68, 0x69, 0x74, 0x74, 0x79, 0x43, 0x68, 0x61, 0x74, 0x74, 0x79,
	0x2e, 0x4c, 0x65, 0x61, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e,
	0x43, 0x68, 0x69, 0x74, 0x74, 0x79, 0x43, 0x68, 0x61, 0x74, 0x74, 0x79, 0x2e, 0x4c, 0x65, 0x61,
	0x76, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x58, 0x0a, 0x10, 0x4c, 0x69,
	0x73, 0x74, 0x65, 0x6e, 0x54, 0x6f, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x12, 0x1b,
	0x2e, 0x43, 0x68, 0x69, 0x74, 0x74, 0x79, 0x43, 0x68, 0x61, 0x74, 0x74, 0x79, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x43, 0x68,
	0x69, 0x74, 0x74, 0x79, 0x43, 0x68, 0x61, 0x74, 0x74, 0x79, 0x2e, 0x42, 0x72, 0x6f, 0x61, 0x64,
	0x63, 0x61, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x30, 0x01, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2e, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_proto_rawDescOnce sync.Once
	file_proto_proto_rawDescData = file_proto_proto_rawDesc
)

func file_proto_proto_rawDescGZIP() []byte {
	file_proto_proto_rawDescOnce.Do(func() {
		file_proto_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_proto_rawDescData)
	})
	return file_proto_proto_rawDescData
}

var file_proto_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_proto_proto_goTypes = []any{
	(*ChatMessageRequest)(nil),      // 0: ChittyChatty.ChatMessageRequest
	(*PublishResponse)(nil),         // 1: ChittyChatty.PublishResponse
	(*BroadcastMessageRequest)(nil), // 2: ChittyChatty.BroadcastMessageRequest
	(*BroadcastResponse)(nil),       // 3: ChittyChatty.BroadcastResponse
	(*JoinRequest)(nil),             // 4: ChittyChatty.JoinRequest
	(*JoinResponse)(nil),            // 5: ChittyChatty.JoinResponse
	(*LeaveRequest)(nil),            // 6: ChittyChatty.LeaveRequest
	(*LeaveResponse)(nil),           // 7: ChittyChatty.LeaveResponse
	(*ListenRequest)(nil),           // 8: ChittyChatty.ListenRequest
}
var file_proto_proto_depIdxs = []int32{
	0, // 0: ChittyChatty.ChittyChattyService.PublishMessage:input_type -> ChittyChatty.ChatMessageRequest
	2, // 1: ChittyChatty.ChittyChattyService.BroadcastMessage:input_type -> ChittyChatty.BroadcastMessageRequest
	4, // 2: ChittyChatty.ChittyChattyService.Join:input_type -> ChittyChatty.JoinRequest
	6, // 3: ChittyChatty.ChittyChattyService.Leave:input_type -> ChittyChatty.LeaveRequest
	8, // 4: ChittyChatty.ChittyChattyService.ListenToMessages:input_type -> ChittyChatty.ListenRequest
	1, // 5: ChittyChatty.ChittyChattyService.PublishMessage:output_type -> ChittyChatty.PublishResponse
	3, // 6: ChittyChatty.ChittyChattyService.BroadcastMessage:output_type -> ChittyChatty.BroadcastResponse
	5, // 7: ChittyChatty.ChittyChattyService.Join:output_type -> ChittyChatty.JoinResponse
	7, // 8: ChittyChatty.ChittyChattyService.Leave:output_type -> ChittyChatty.LeaveResponse
	2, // 9: ChittyChatty.ChittyChattyService.ListenToMessages:output_type -> ChittyChatty.BroadcastMessageRequest
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_proto_init() }
func file_proto_proto_init() {
	if File_proto_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_proto_goTypes,
		DependencyIndexes: file_proto_proto_depIdxs,
		MessageInfos:      file_proto_proto_msgTypes,
	}.Build()
	File_proto_proto = out.File
	file_proto_proto_rawDesc = nil
	file_proto_proto_goTypes = nil
	file_proto_proto_depIdxs = nil
}
