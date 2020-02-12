// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/messaging/messaging.proto

package proto_messaging

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type MessagingListType int32

const (
	MessagingListType_Public    MessagingListType = 0
	MessagingListType_Private   MessagingListType = 1
	MessagingListType_Protected MessagingListType = 2
)

var MessagingListType_name = map[int32]string{
	0: "Public",
	1: "Private",
	2: "Protected",
}

var MessagingListType_value = map[string]int32{
	"Public":    0,
	"Private":   1,
	"Protected": 2,
}

func (x MessagingListType) String() string {
	return proto.EnumName(MessagingListType_name, int32(x))
}

func (MessagingListType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_2694837308dc07a6, []int{0}
}

type Message struct {
	Say                  string   `protobuf:"bytes,1,opt,name=say,proto3" json:"say,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_2694837308dc07a6, []int{0}
}

func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (m *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetSay() string {
	if m != nil {
		return m.Say
	}
	return ""
}

type MessagingRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MessagingRequest) Reset()         { *m = MessagingRequest{} }
func (m *MessagingRequest) String() string { return proto.CompactTextString(m) }
func (*MessagingRequest) ProtoMessage()    {}
func (*MessagingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2694837308dc07a6, []int{1}
}

func (m *MessagingRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MessagingRequest.Unmarshal(m, b)
}
func (m *MessagingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MessagingRequest.Marshal(b, m, deterministic)
}
func (m *MessagingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MessagingRequest.Merge(m, src)
}
func (m *MessagingRequest) XXX_Size() int {
	return xxx_messageInfo_MessagingRequest.Size(m)
}
func (m *MessagingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MessagingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MessagingRequest proto.InternalMessageInfo

func (m *MessagingRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type MessagingResponse struct {
	Msg                  string   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MessagingResponse) Reset()         { *m = MessagingResponse{} }
func (m *MessagingResponse) String() string { return proto.CompactTextString(m) }
func (*MessagingResponse) ProtoMessage()    {}
func (*MessagingResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2694837308dc07a6, []int{2}
}

func (m *MessagingResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MessagingResponse.Unmarshal(m, b)
}
func (m *MessagingResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MessagingResponse.Marshal(b, m, deterministic)
}
func (m *MessagingResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MessagingResponse.Merge(m, src)
}
func (m *MessagingResponse) XXX_Size() int {
	return xxx_messageInfo_MessagingResponse.Size(m)
}
func (m *MessagingResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MessagingResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MessagingResponse proto.InternalMessageInfo

func (m *MessagingResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type MessagingListRequest struct {
	MessageType          MessagingListType `protobuf:"varint,1,opt,name=message_type,json=messageType,proto3,enum=proto.messaging.MessagingListType" json:"message_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *MessagingListRequest) Reset()         { *m = MessagingListRequest{} }
func (m *MessagingListRequest) String() string { return proto.CompactTextString(m) }
func (*MessagingListRequest) ProtoMessage()    {}
func (*MessagingListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2694837308dc07a6, []int{3}
}

func (m *MessagingListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MessagingListRequest.Unmarshal(m, b)
}
func (m *MessagingListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MessagingListRequest.Marshal(b, m, deterministic)
}
func (m *MessagingListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MessagingListRequest.Merge(m, src)
}
func (m *MessagingListRequest) XXX_Size() int {
	return xxx_messageInfo_MessagingListRequest.Size(m)
}
func (m *MessagingListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MessagingListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MessagingListRequest proto.InternalMessageInfo

func (m *MessagingListRequest) GetMessageType() MessagingListType {
	if m != nil {
		return m.MessageType
	}
	return MessagingListType_Public
}

type MessagingListResponse struct {
	Messages             string   `protobuf:"bytes,1,opt,name=messages,proto3" json:"messages,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MessagingListResponse) Reset()         { *m = MessagingListResponse{} }
func (m *MessagingListResponse) String() string { return proto.CompactTextString(m) }
func (*MessagingListResponse) ProtoMessage()    {}
func (*MessagingListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2694837308dc07a6, []int{4}
}

func (m *MessagingListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MessagingListResponse.Unmarshal(m, b)
}
func (m *MessagingListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MessagingListResponse.Marshal(b, m, deterministic)
}
func (m *MessagingListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MessagingListResponse.Merge(m, src)
}
func (m *MessagingListResponse) XXX_Size() int {
	return xxx_messageInfo_MessagingListResponse.Size(m)
}
func (m *MessagingListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MessagingListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MessagingListResponse proto.InternalMessageInfo

func (m *MessagingListResponse) GetMessages() string {
	if m != nil {
		return m.Messages
	}
	return ""
}

type StreamingRequest struct {
	Count                int64    `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamingRequest) Reset()         { *m = StreamingRequest{} }
func (m *StreamingRequest) String() string { return proto.CompactTextString(m) }
func (*StreamingRequest) ProtoMessage()    {}
func (*StreamingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2694837308dc07a6, []int{5}
}

func (m *StreamingRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamingRequest.Unmarshal(m, b)
}
func (m *StreamingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamingRequest.Marshal(b, m, deterministic)
}
func (m *StreamingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamingRequest.Merge(m, src)
}
func (m *StreamingRequest) XXX_Size() int {
	return xxx_messageInfo_StreamingRequest.Size(m)
}
func (m *StreamingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StreamingRequest proto.InternalMessageInfo

func (m *StreamingRequest) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

type StreamingResponse struct {
	Count                int64    `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamingResponse) Reset()         { *m = StreamingResponse{} }
func (m *StreamingResponse) String() string { return proto.CompactTextString(m) }
func (*StreamingResponse) ProtoMessage()    {}
func (*StreamingResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2694837308dc07a6, []int{6}
}

func (m *StreamingResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamingResponse.Unmarshal(m, b)
}
func (m *StreamingResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamingResponse.Marshal(b, m, deterministic)
}
func (m *StreamingResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamingResponse.Merge(m, src)
}
func (m *StreamingResponse) XXX_Size() int {
	return xxx_messageInfo_StreamingResponse.Size(m)
}
func (m *StreamingResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamingResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StreamingResponse proto.InternalMessageInfo

func (m *StreamingResponse) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

type Ping struct {
	Stroke               int64    `protobuf:"varint,1,opt,name=stroke,proto3" json:"stroke,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ping) Reset()         { *m = Ping{} }
func (m *Ping) String() string { return proto.CompactTextString(m) }
func (*Ping) ProtoMessage()    {}
func (*Ping) Descriptor() ([]byte, []int) {
	return fileDescriptor_2694837308dc07a6, []int{7}
}

func (m *Ping) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ping.Unmarshal(m, b)
}
func (m *Ping) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ping.Marshal(b, m, deterministic)
}
func (m *Ping) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ping.Merge(m, src)
}
func (m *Ping) XXX_Size() int {
	return xxx_messageInfo_Ping.Size(m)
}
func (m *Ping) XXX_DiscardUnknown() {
	xxx_messageInfo_Ping.DiscardUnknown(m)
}

var xxx_messageInfo_Ping proto.InternalMessageInfo

func (m *Ping) GetStroke() int64 {
	if m != nil {
		return m.Stroke
	}
	return 0
}

type Pong struct {
	Stroke               int64    `protobuf:"varint,1,opt,name=stroke,proto3" json:"stroke,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Pong) Reset()         { *m = Pong{} }
func (m *Pong) String() string { return proto.CompactTextString(m) }
func (*Pong) ProtoMessage()    {}
func (*Pong) Descriptor() ([]byte, []int) {
	return fileDescriptor_2694837308dc07a6, []int{8}
}

func (m *Pong) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Pong.Unmarshal(m, b)
}
func (m *Pong) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Pong.Marshal(b, m, deterministic)
}
func (m *Pong) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pong.Merge(m, src)
}
func (m *Pong) XXX_Size() int {
	return xxx_messageInfo_Pong.Size(m)
}
func (m *Pong) XXX_DiscardUnknown() {
	xxx_messageInfo_Pong.DiscardUnknown(m)
}

var xxx_messageInfo_Pong proto.InternalMessageInfo

func (m *Pong) GetStroke() int64 {
	if m != nil {
		return m.Stroke
	}
	return 0
}

func init() {
	proto.RegisterEnum("proto.messaging.MessagingListType", MessagingListType_name, MessagingListType_value)
	proto.RegisterType((*Message)(nil), "proto.messaging.Message")
	proto.RegisterType((*MessagingRequest)(nil), "proto.messaging.MessagingRequest")
	proto.RegisterType((*MessagingResponse)(nil), "proto.messaging.MessagingResponse")
	proto.RegisterType((*MessagingListRequest)(nil), "proto.messaging.MessagingListRequest")
	proto.RegisterType((*MessagingListResponse)(nil), "proto.messaging.MessagingListResponse")
	proto.RegisterType((*StreamingRequest)(nil), "proto.messaging.StreamingRequest")
	proto.RegisterType((*StreamingResponse)(nil), "proto.messaging.StreamingResponse")
	proto.RegisterType((*Ping)(nil), "proto.messaging.Ping")
	proto.RegisterType((*Pong)(nil), "proto.messaging.Pong")
}

func init() { proto.RegisterFile("proto/messaging/messaging.proto", fileDescriptor_2694837308dc07a6) }

var fileDescriptor_2694837308dc07a6 = []byte{
	// 377 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x51, 0x4f, 0x4f, 0xfa, 0x40,
	0x10, 0x6d, 0xa1, 0xbf, 0x02, 0xc3, 0x4f, 0x2d, 0x1b, 0x30, 0x06, 0x13, 0xff, 0x6c, 0x02, 0x41,
	0x0f, 0x95, 0xc0, 0xd1, 0xc4, 0x8b, 0xf1, 0x26, 0xb1, 0x29, 0x26, 0x9e, 0x8c, 0x29, 0x75, 0xd3,
	0x34, 0xd2, 0x6e, 0xed, 0x2e, 0x24, 0x7c, 0x4c, 0xbf, 0x91, 0xe9, 0x76, 0x5b, 0x6b, 0x29, 0x78,
	0xda, 0x99, 0x7d, 0x6f, 0x5e, 0xde, 0xcc, 0x83, 0xf3, 0x28, 0xa6, 0x9c, 0xde, 0x04, 0x84, 0x31,
	0xc7, 0xf3, 0x43, 0xef, 0xa7, 0x32, 0x05, 0x82, 0x8e, 0xc4, 0x63, 0xe6, 0xdf, 0xf8, 0x14, 0x1a,
	0x33, 0xd1, 0x10, 0x64, 0x40, 0x9d, 0x39, 0x9b, 0x13, 0xf5, 0x42, 0x1d, 0xb5, 0xec, 0xa4, 0xc4,
	0x43, 0x30, 0x66, 0x19, 0xd3, 0x26, 0x9f, 0x2b, 0xc2, 0x38, 0x42, 0xa0, 0x85, 0x4e, 0x40, 0x24,
	0x4d, 0xd4, 0x78, 0x00, 0x9d, 0x02, 0x8f, 0x45, 0x34, 0x64, 0x42, 0x2e, 0x60, 0x5e, 0x26, 0x17,
	0x30, 0x0f, 0xbf, 0x42, 0x37, 0xa7, 0x3d, 0xfa, 0x8c, 0x67, 0x92, 0x0f, 0xf0, 0x3f, 0x35, 0x44,
	0xde, 0xf8, 0x26, 0x4a, 0xa5, 0x0f, 0x27, 0xd8, 0x2c, 0x79, 0x35, 0x7f, 0x0d, 0x3f, 0x6f, 0x22,
	0x62, 0xb7, 0xe5, 0x5c, 0xd2, 0xe0, 0x29, 0xf4, 0x4a, 0xf2, 0xd2, 0x49, 0x1f, 0x9a, 0x92, 0xc7,
	0xa4, 0x9d, 0xbc, 0xc7, 0x23, 0x30, 0xe6, 0x3c, 0x26, 0x4e, 0x50, 0x58, 0xb1, 0x0b, 0xff, 0x5c,
	0xba, 0x0a, 0xb9, 0x20, 0xd7, 0xed, 0xb4, 0xc1, 0x57, 0xd0, 0x29, 0x30, 0xa5, 0x74, 0x35, 0xf5,
	0x0c, 0x34, 0xcb, 0x0f, 0x3d, 0x74, 0x0c, 0x3a, 0xe3, 0x31, 0xfd, 0x20, 0x12, 0x96, 0x9d, 0xc0,
	0xe9, 0x6e, 0xfc, 0xfa, 0xb6, 0x70, 0xcf, 0x6c, 0x57, 0x04, 0xa0, 0x5b, 0xab, 0xc5, 0xd2, 0x77,
	0x0d, 0x05, 0xb5, 0xa1, 0x61, 0xc5, 0xfe, 0xda, 0xe1, 0xc4, 0x50, 0xd1, 0x01, 0xb4, 0xac, 0x98,
	0x72, 0xe2, 0x72, 0xf2, 0x6e, 0xd4, 0x26, 0x5f, 0xb5, 0x42, 0x6a, 0x73, 0x12, 0xaf, 0x7d, 0x97,
	0xa0, 0x27, 0xd0, 0xee, 0x9d, 0xe5, 0x12, 0x5d, 0xee, 0x3e, 0xaa, 0xdc, 0xbe, 0x8f, 0xf7, 0x51,
	0xd2, 0xb5, 0xb1, 0x82, 0x5e, 0x40, 0x4b, 0x9c, 0xa1, 0xc1, 0xfe, 0x94, 0x32, 0xd1, 0xe1, 0x5f,
	0xb4, 0x5c, 0x78, 0x0e, 0x7a, 0x7a, 0xe6, 0x0a, 0xaf, 0xe5, 0xa4, 0x2a, 0xbc, 0x6e, 0x45, 0x84,
	0x95, 0xb1, 0x8a, 0xee, 0xa0, 0x99, 0x04, 0x22, 0x8e, 0xde, 0xdb, 0x9a, 0x49, 0xa0, 0x7e, 0xc5,
	0x37, 0x0d, 0x3d, 0xac, 0x8c, 0xd4, 0xb1, 0xba, 0xd0, 0x05, 0x36, 0xfd, 0x0e, 0x00, 0x00, 0xff,
	0xff, 0x4f, 0x37, 0xc8, 0xa9, 0x60, 0x03, 0x00, 0x00,
}
