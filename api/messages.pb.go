// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v4.25.2
// source: proto/messages.proto

package api

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type PublishMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Time  *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=time,proto3" json:"time,omitempty"`
	Key   []byte                 `protobuf:"bytes,3,opt,name=key,proto3" json:"key,omitempty"`
	Value []byte                 `protobuf:"bytes,4,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *PublishMessage) Reset() {
	*x = PublishMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_messages_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublishMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublishMessage) ProtoMessage() {}

func (x *PublishMessage) ProtoReflect() protoreflect.Message {
	mi := &file_proto_messages_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublishMessage.ProtoReflect.Descriptor instead.
func (*PublishMessage) Descriptor() ([]byte, []int) {
	return file_proto_messages_proto_rawDescGZIP(), []int{0}
}

func (x *PublishMessage) GetTime() *timestamppb.Timestamp {
	if x != nil {
		return x.Time
	}
	return nil
}

func (x *PublishMessage) GetKey() []byte {
	if x != nil {
		return x.Key
	}
	return nil
}

func (x *PublishMessage) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

type ConsumeMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Offset int64                  `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`
	Time   *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=time,proto3" json:"time,omitempty"`
	Key    []byte                 `protobuf:"bytes,3,opt,name=key,proto3" json:"key,omitempty"`
	Value  []byte                 `protobuf:"bytes,4,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *ConsumeMessage) Reset() {
	*x = ConsumeMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_messages_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConsumeMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConsumeMessage) ProtoMessage() {}

func (x *ConsumeMessage) ProtoReflect() protoreflect.Message {
	mi := &file_proto_messages_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConsumeMessage.ProtoReflect.Descriptor instead.
func (*ConsumeMessage) Descriptor() ([]byte, []int) {
	return file_proto_messages_proto_rawDescGZIP(), []int{1}
}

func (x *ConsumeMessage) GetOffset() int64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *ConsumeMessage) GetTime() *timestamppb.Timestamp {
	if x != nil {
		return x.Time
	}
	return nil
}

func (x *ConsumeMessage) GetKey() []byte {
	if x != nil {
		return x.Key
	}
	return nil
}

func (x *ConsumeMessage) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

type PublishRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string            `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Messages []*PublishMessage `protobuf:"bytes,2,rep,name=messages,proto3" json:"messages,omitempty"`
}

func (x *PublishRequest) Reset() {
	*x = PublishRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_messages_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublishRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublishRequest) ProtoMessage() {}

func (x *PublishRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_messages_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublishRequest.ProtoReflect.Descriptor instead.
func (*PublishRequest) Descriptor() ([]byte, []int) {
	return file_proto_messages_proto_rawDescGZIP(), []int{2}
}

func (x *PublishRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PublishRequest) GetMessages() []*PublishMessage {
	if x != nil {
		return x.Messages
	}
	return nil
}

type PublishResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NextOffset int64 `protobuf:"varint,1,opt,name=next_offset,json=nextOffset,proto3" json:"next_offset,omitempty"`
}

func (x *PublishResponse) Reset() {
	*x = PublishResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_messages_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublishResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublishResponse) ProtoMessage() {}

func (x *PublishResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_messages_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
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
	return file_proto_messages_proto_rawDescGZIP(), []int{3}
}

func (x *PublishResponse) GetNextOffset() int64 {
	if x != nil {
		return x.NextOffset
	}
	return 0
}

type ConsumeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Offset   int64  `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	MaxCount int64  `protobuf:"varint,3,opt,name=max_count,json=maxCount,proto3" json:"max_count,omitempty"`
}

func (x *ConsumeRequest) Reset() {
	*x = ConsumeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_messages_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConsumeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConsumeRequest) ProtoMessage() {}

func (x *ConsumeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_messages_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConsumeRequest.ProtoReflect.Descriptor instead.
func (*ConsumeRequest) Descriptor() ([]byte, []int) {
	return file_proto_messages_proto_rawDescGZIP(), []int{4}
}

func (x *ConsumeRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ConsumeRequest) GetOffset() int64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *ConsumeRequest) GetMaxCount() int64 {
	if x != nil {
		return x.MaxCount
	}
	return 0
}

type ConsumeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NextOffset int64             `protobuf:"varint,1,opt,name=next_offset,json=nextOffset,proto3" json:"next_offset,omitempty"`
	Messages   []*ConsumeMessage `protobuf:"bytes,2,rep,name=messages,proto3" json:"messages,omitempty"`
}

func (x *ConsumeResponse) Reset() {
	*x = ConsumeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_messages_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConsumeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConsumeResponse) ProtoMessage() {}

func (x *ConsumeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_messages_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConsumeResponse.ProtoReflect.Descriptor instead.
func (*ConsumeResponse) Descriptor() ([]byte, []int) {
	return file_proto_messages_proto_rawDescGZIP(), []int{5}
}

func (x *ConsumeResponse) GetNextOffset() int64 {
	if x != nil {
		return x.NextOffset
	}
	return 0
}

func (x *ConsumeResponse) GetMessages() []*ConsumeMessage {
	if x != nil {
		return x.Messages
	}
	return nil
}

var File_proto_messages_proto protoreflect.FileDescriptor

var file_proto_messages_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x61, 0x70, 0x69, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x68, 0x0a, 0x0e,
	0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x2e,
	0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x80, 0x01, 0x0a, 0x0e, 0x43, 0x6f, 0x6e, 0x73, 0x75,
	0x6d, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66,
	0x73, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65,
	0x74, 0x12, 0x2e, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x74, 0x69, 0x6d,
	0x65, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x55, 0x0a, 0x0e, 0x50, 0x75, 0x62,
	0x6c, 0x69, 0x73, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x2f, 0x0a, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x13, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73,
	0x22, 0x32, 0x0a, 0x0f, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x6f, 0x66, 0x66, 0x73,
	0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x6e, 0x65, 0x78, 0x74, 0x4f, 0x66,
	0x66, 0x73, 0x65, 0x74, 0x22, 0x59, 0x0a, 0x0e, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66,
	0x66, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73,
	0x65, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x61, 0x78, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x6d, 0x61, 0x78, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22,
	0x63, 0x0a, 0x0f, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x6f, 0x66, 0x66, 0x73, 0x65,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x6e, 0x65, 0x78, 0x74, 0x4f, 0x66, 0x66,
	0x73, 0x65, 0x74, 0x12, 0x2f, 0x0a, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x6f, 0x6e, 0x73,
	0x75, 0x6d, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x08, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x73, 0x32, 0x7a, 0x0a, 0x08, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73,
	0x12, 0x36, 0x0a, 0x07, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x12, 0x13, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x14, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x36, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x73,
	0x75, 0x6d, 0x65, 0x12, 0x13, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43,
	0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x42, 0x17, 0x5a, 0x15, 0x67, 0x6f, 0x2e, 0x6b, 0x6c, 0x65, 0x76, 0x2e, 0x64, 0x65, 0x76, 0x2f,
	0x6b, 0x6c, 0x65, 0x76, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_proto_messages_proto_rawDescOnce sync.Once
	file_proto_messages_proto_rawDescData = file_proto_messages_proto_rawDesc
)

func file_proto_messages_proto_rawDescGZIP() []byte {
	file_proto_messages_proto_rawDescOnce.Do(func() {
		file_proto_messages_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_messages_proto_rawDescData)
	})
	return file_proto_messages_proto_rawDescData
}

var file_proto_messages_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_proto_messages_proto_goTypes = []interface{}{
	(*PublishMessage)(nil),        // 0: api.PublishMessage
	(*ConsumeMessage)(nil),        // 1: api.ConsumeMessage
	(*PublishRequest)(nil),        // 2: api.PublishRequest
	(*PublishResponse)(nil),       // 3: api.PublishResponse
	(*ConsumeRequest)(nil),        // 4: api.ConsumeRequest
	(*ConsumeResponse)(nil),       // 5: api.ConsumeResponse
	(*timestamppb.Timestamp)(nil), // 6: google.protobuf.Timestamp
}
var file_proto_messages_proto_depIdxs = []int32{
	6, // 0: api.PublishMessage.time:type_name -> google.protobuf.Timestamp
	6, // 1: api.ConsumeMessage.time:type_name -> google.protobuf.Timestamp
	0, // 2: api.PublishRequest.messages:type_name -> api.PublishMessage
	1, // 3: api.ConsumeResponse.messages:type_name -> api.ConsumeMessage
	2, // 4: api.Messages.Publish:input_type -> api.PublishRequest
	4, // 5: api.Messages.Consume:input_type -> api.ConsumeRequest
	3, // 6: api.Messages.Publish:output_type -> api.PublishResponse
	5, // 7: api.Messages.Consume:output_type -> api.ConsumeResponse
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_proto_messages_proto_init() }
func file_proto_messages_proto_init() {
	if File_proto_messages_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_messages_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublishMessage); i {
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
		file_proto_messages_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConsumeMessage); i {
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
		file_proto_messages_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublishRequest); i {
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
		file_proto_messages_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublishResponse); i {
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
		file_proto_messages_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConsumeRequest); i {
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
		file_proto_messages_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConsumeResponse); i {
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
			RawDescriptor: file_proto_messages_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_messages_proto_goTypes,
		DependencyIndexes: file_proto_messages_proto_depIdxs,
		MessageInfos:      file_proto_messages_proto_msgTypes,
	}.Build()
	File_proto_messages_proto = out.File
	file_proto_messages_proto_rawDesc = nil
	file_proto_messages_proto_goTypes = nil
	file_proto_messages_proto_depIdxs = nil
}
