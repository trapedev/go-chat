// Code generated by protoc-gen-go. DO NOT EDIT.
// source: hellogrpc.proto

package HelloGrpc

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// 空のmessage
type Null struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Null) Reset()         { *m = Null{} }
func (m *Null) String() string { return proto.CompactTextString(m) }
func (*Null) ProtoMessage()    {}
func (*Null) Descriptor() ([]byte, []int) {
	return fileDescriptor_94272a0853831a50, []int{0}
}

func (m *Null) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Null.Unmarshal(m, b)
}
func (m *Null) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Null.Marshal(b, m, deterministic)
}
func (m *Null) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Null.Merge(m, src)
}
func (m *Null) XXX_Size() int {
	return xxx_messageInfo_Null.Size(m)
}
func (m *Null) XXX_DiscardUnknown() {
	xxx_messageInfo_Null.DiscardUnknown(m)
}

var xxx_messageInfo_Null proto.InternalMessageInfo

type GreetRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GreetRequest) Reset()         { *m = GreetRequest{} }
func (m *GreetRequest) String() string { return proto.CompactTextString(m) }
func (*GreetRequest) ProtoMessage()    {}
func (*GreetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_94272a0853831a50, []int{1}
}

func (m *GreetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GreetRequest.Unmarshal(m, b)
}
func (m *GreetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GreetRequest.Marshal(b, m, deterministic)
}
func (m *GreetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GreetRequest.Merge(m, src)
}
func (m *GreetRequest) XXX_Size() int {
	return xxx_messageInfo_GreetRequest.Size(m)
}
func (m *GreetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GreetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GreetRequest proto.InternalMessageInfo

func (m *GreetRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type GreetMessage struct {
	Msg                  string   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GreetMessage) Reset()         { *m = GreetMessage{} }
func (m *GreetMessage) String() string { return proto.CompactTextString(m) }
func (*GreetMessage) ProtoMessage()    {}
func (*GreetMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_94272a0853831a50, []int{2}
}

func (m *GreetMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GreetMessage.Unmarshal(m, b)
}
func (m *GreetMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GreetMessage.Marshal(b, m, deterministic)
}
func (m *GreetMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GreetMessage.Merge(m, src)
}
func (m *GreetMessage) XXX_Size() int {
	return xxx_messageInfo_GreetMessage.Size(m)
}
func (m *GreetMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_GreetMessage.DiscardUnknown(m)
}

var xxx_messageInfo_GreetMessage proto.InternalMessageInfo

func (m *GreetMessage) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type RoomRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RoomRequest) Reset()         { *m = RoomRequest{} }
func (m *RoomRequest) String() string { return proto.CompactTextString(m) }
func (*RoomRequest) ProtoMessage()    {}
func (*RoomRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_94272a0853831a50, []int{3}
}

func (m *RoomRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoomRequest.Unmarshal(m, b)
}
func (m *RoomRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoomRequest.Marshal(b, m, deterministic)
}
func (m *RoomRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoomRequest.Merge(m, src)
}
func (m *RoomRequest) XXX_Size() int {
	return xxx_messageInfo_RoomRequest.Size(m)
}
func (m *RoomRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RoomRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RoomRequest proto.InternalMessageInfo

func (m *RoomRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type RoomInfo struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	MessageCount         int32    `protobuf:"varint,2,opt,name=messageCount,proto3" json:"messageCount,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RoomInfo) Reset()         { *m = RoomInfo{} }
func (m *RoomInfo) String() string { return proto.CompactTextString(m) }
func (*RoomInfo) ProtoMessage()    {}
func (*RoomInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_94272a0853831a50, []int{4}
}

func (m *RoomInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoomInfo.Unmarshal(m, b)
}
func (m *RoomInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoomInfo.Marshal(b, m, deterministic)
}
func (m *RoomInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoomInfo.Merge(m, src)
}
func (m *RoomInfo) XXX_Size() int {
	return xxx_messageInfo_RoomInfo.Size(m)
}
func (m *RoomInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_RoomInfo.DiscardUnknown(m)
}

var xxx_messageInfo_RoomInfo proto.InternalMessageInfo

func (m *RoomInfo) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *RoomInfo) GetMessageCount() int32 {
	if m != nil {
		return m.MessageCount
	}
	return 0
}

type RoomList struct {
	Rooms                []*RoomInfo `protobuf:"bytes,1,rep,name=rooms,proto3" json:"rooms,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *RoomList) Reset()         { *m = RoomList{} }
func (m *RoomList) String() string { return proto.CompactTextString(m) }
func (*RoomList) ProtoMessage()    {}
func (*RoomList) Descriptor() ([]byte, []int) {
	return fileDescriptor_94272a0853831a50, []int{5}
}

func (m *RoomList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoomList.Unmarshal(m, b)
}
func (m *RoomList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoomList.Marshal(b, m, deterministic)
}
func (m *RoomList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoomList.Merge(m, src)
}
func (m *RoomList) XXX_Size() int {
	return xxx_messageInfo_RoomList.Size(m)
}
func (m *RoomList) XXX_DiscardUnknown() {
	xxx_messageInfo_RoomList.DiscardUnknown(m)
}

var xxx_messageInfo_RoomList proto.InternalMessageInfo

func (m *RoomList) GetRooms() []*RoomInfo {
	if m != nil {
		return m.Rooms
	}
	return nil
}

type SendRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Content              string   `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SendRequest) Reset()         { *m = SendRequest{} }
func (m *SendRequest) String() string { return proto.CompactTextString(m) }
func (*SendRequest) ProtoMessage()    {}
func (*SendRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_94272a0853831a50, []int{6}
}

func (m *SendRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendRequest.Unmarshal(m, b)
}
func (m *SendRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendRequest.Marshal(b, m, deterministic)
}
func (m *SendRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendRequest.Merge(m, src)
}
func (m *SendRequest) XXX_Size() int {
	return xxx_messageInfo_SendRequest.Size(m)
}
func (m *SendRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SendRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SendRequest proto.InternalMessageInfo

func (m *SendRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *SendRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SendRequest) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

type SendResult struct {
	Result               bool     `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SendResult) Reset()         { *m = SendResult{} }
func (m *SendResult) String() string { return proto.CompactTextString(m) }
func (*SendResult) ProtoMessage()    {}
func (*SendResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_94272a0853831a50, []int{7}
}

func (m *SendResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendResult.Unmarshal(m, b)
}
func (m *SendResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendResult.Marshal(b, m, deterministic)
}
func (m *SendResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendResult.Merge(m, src)
}
func (m *SendResult) XXX_Size() int {
	return xxx_messageInfo_SendResult.Size(m)
}
func (m *SendResult) XXX_DiscardUnknown() {
	xxx_messageInfo_SendResult.DiscardUnknown(m)
}

var xxx_messageInfo_SendResult proto.InternalMessageInfo

func (m *SendResult) GetResult() bool {
	if m != nil {
		return m.Result
	}
	return false
}

type MessagesRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MessagesRequest) Reset()         { *m = MessagesRequest{} }
func (m *MessagesRequest) String() string { return proto.CompactTextString(m) }
func (*MessagesRequest) ProtoMessage()    {}
func (*MessagesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_94272a0853831a50, []int{8}
}

func (m *MessagesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MessagesRequest.Unmarshal(m, b)
}
func (m *MessagesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MessagesRequest.Marshal(b, m, deterministic)
}
func (m *MessagesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MessagesRequest.Merge(m, src)
}
func (m *MessagesRequest) XXX_Size() int {
	return xxx_messageInfo_MessagesRequest.Size(m)
}
func (m *MessagesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MessagesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MessagesRequest proto.InternalMessageInfo

func (m *MessagesRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type Message struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Content              string   `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_94272a0853831a50, []int{9}
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

func (m *Message) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Message) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Message) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func init() {
	proto.RegisterType((*Null)(nil), "HelloGrpc.Null")
	proto.RegisterType((*GreetRequest)(nil), "HelloGrpc.GreetRequest")
	proto.RegisterType((*GreetMessage)(nil), "HelloGrpc.GreetMessage")
	proto.RegisterType((*RoomRequest)(nil), "HelloGrpc.RoomRequest")
	proto.RegisterType((*RoomInfo)(nil), "HelloGrpc.RoomInfo")
	proto.RegisterType((*RoomList)(nil), "HelloGrpc.RoomList")
	proto.RegisterType((*SendRequest)(nil), "HelloGrpc.SendRequest")
	proto.RegisterType((*SendResult)(nil), "HelloGrpc.SendResult")
	proto.RegisterType((*MessagesRequest)(nil), "HelloGrpc.MessagesRequest")
	proto.RegisterType((*Message)(nil), "HelloGrpc.Message")
}

func init() { proto.RegisterFile("hellogrpc.proto", fileDescriptor_94272a0853831a50) }

var fileDescriptor_94272a0853831a50 = []byte{
	// 406 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x93, 0x4b, 0x6f, 0xd3, 0x40,
	0x10, 0xc7, 0x6d, 0xe7, 0x3d, 0x8e, 0x08, 0x1a, 0x44, 0xb0, 0x2c, 0x81, 0xcc, 0x8a, 0x83, 0x39,
	0x60, 0xa1, 0x20, 0x24, 0x0e, 0x08, 0x91, 0x70, 0x30, 0x88, 0xb6, 0xaa, 0x9c, 0x4f, 0x90, 0xc6,
	0xdb, 0xd4, 0x8a, 0xed, 0x4d, 0x77, 0xd7, 0xfd, 0xfa, 0xad, 0xbc, 0x7e, 0x36, 0x6e, 0x0e, 0x55,
	0xe5, 0xcb, 0x78, 0x5e, 0x3b, 0xf3, 0xff, 0xed, 0xc2, 0xec, 0x86, 0xc6, 0x31, 0xdb, 0xf1, 0xc3,
	0xd6, 0x3b, 0x70, 0x26, 0x19, 0x4e, 0xfe, 0xe6, 0x0e, 0x9f, 0x1f, 0xb6, 0x64, 0x08, 0xfd, 0x8b,
	0x2c, 0x8e, 0x09, 0x81, 0xa9, 0xcf, 0x29, 0x95, 0x01, 0xbd, 0xcd, 0xa8, 0x90, 0x88, 0xd0, 0x4f,
	0x37, 0x09, 0xb5, 0x74, 0x47, 0x77, 0x27, 0x81, 0xb2, 0x89, 0x53, 0xe6, 0x9c, 0x53, 0x21, 0x36,
	0x3b, 0x8a, 0xaf, 0xa1, 0x97, 0x88, 0x5d, 0x99, 0x92, 0x9b, 0xe4, 0x3d, 0x98, 0x01, 0x63, 0x49,
	0xd5, 0xe4, 0x15, 0x18, 0x51, 0x58, 0xc6, 0x8d, 0x28, 0x24, 0xbf, 0x60, 0x9c, 0x87, 0xff, 0xa5,
	0xd7, 0xec, 0x38, 0x86, 0x04, 0xa6, 0x49, 0xd1, 0xf7, 0x0f, 0xcb, 0x52, 0x69, 0x19, 0x8e, 0xee,
	0x0e, 0x82, 0x47, 0x3e, 0xf2, 0xbd, 0xa8, 0x3f, 0x8b, 0x84, 0xc4, 0xcf, 0x30, 0xe0, 0x8c, 0x25,
	0xc2, 0xd2, 0x9d, 0x9e, 0x6b, 0x2e, 0xde, 0x78, 0xf5, 0x4e, 0x5e, 0x75, 0x46, 0x50, 0x64, 0x90,
	0xff, 0x60, 0xae, 0x69, 0x1a, 0x9e, 0x98, 0xaa, 0x5e, 0xd5, 0x68, 0x56, 0x45, 0x0b, 0x46, 0x5b,
	0x96, 0x4a, 0x9a, 0x4a, 0xab, 0xa7, 0xdc, 0xd5, 0x2f, 0xf9, 0x04, 0x50, 0x34, 0x13, 0x59, 0x2c,
	0x71, 0x0e, 0x43, 0xae, 0x2c, 0xd5, 0x6f, 0x1c, 0x94, 0x7f, 0xe4, 0x23, 0xcc, 0x4a, 0x95, 0xc4,
	0x29, 0x31, 0x7c, 0x18, 0x55, 0x42, 0xbe, 0x68, 0xa2, 0xc5, 0xbd, 0x01, 0x0d, 0x50, 0x5c, 0x82,
	0xa9, 0x20, 0xad, 0x29, 0xbf, 0xa3, 0x1c, 0xdf, 0xb5, 0x74, 0x69, 0x03, 0xb6, 0x3b, 0x81, 0x72,
	0x18, 0xa2, 0xe1, 0x0f, 0x18, 0x2d, 0xc3, 0x30, 0x57, 0x11, 0xe7, 0x47, 0xb2, 0x56, 0xd5, 0x4f,
	0xc9, 0x4d, 0x34, 0xfc, 0x09, 0xa6, 0x4f, 0x65, 0xcd, 0xf8, 0x99, 0xd5, 0x0b, 0x18, 0x97, 0xd5,
	0x02, 0x67, 0xad, 0x94, 0xfc, 0x82, 0x76, 0x6a, 0xf2, 0x4b, 0x40, 0x34, 0xfc, 0x5d, 0xb0, 0xad,
	0x94, 0x6c, 0x9f, 0xd8, 0x62, 0x6e, 0xbf, 0xed, 0xf8, 0x15, 0x26, 0xcd, 0xd5, 0x95, 0x60, 0xf5,
	0xf6, 0x02, 0xed, 0x56, 0xe6, 0x11, 0x42, 0x1b, 0xbb, 0x31, 0xa2, 0x7d, 0xd5, 0x57, 0x5f, 0xe0,
	0x43, 0xc4, 0x3c, 0xc9, 0xf6, 0xd1, 0x9e, 0xe5, 0x9f, 0xd7, 0x3c, 0x38, 0xa1, 0x30, 0xac, 0x1a,
	0x40, 0x97, 0xfa, 0xd5, 0x50, 0xbd, 0xc2, 0x6f, 0x0f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xe9, 0x7f,
	0x02, 0xba, 0x98, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// HelloGrpcClient is the client API for HelloGrpc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HelloGrpcClient interface {
	// Sends a greeting
	GreetServer(ctx context.Context, in *GreetRequest, opts ...grpc.CallOption) (*GreetMessage, error)
	// チャットルーム追加
	AddRoom(ctx context.Context, in *RoomRequest, opts ...grpc.CallOption) (*RoomInfo, error)
	// チャットルーム情報取得
	GetRoomInfo(ctx context.Context, in *RoomRequest, opts ...grpc.CallOption) (*RoomInfo, error)
	// チャットルーム一覧取得
	GetRooms(ctx context.Context, in *Null, opts ...grpc.CallOption) (*RoomList, error)
	// Client-side streamingを用いてメッセージを交換する
	SendMessage(ctx context.Context, opts ...grpc.CallOption) (HelloGrpc_SendMessageClient, error)
	// Server-side streamingを用いてメッセージを交換する
	GetMessages(ctx context.Context, in *MessagesRequest, opts ...grpc.CallOption) (HelloGrpc_GetMessagesClient, error)
}

type helloGrpcClient struct {
	cc *grpc.ClientConn
}

func NewHelloGrpcClient(cc *grpc.ClientConn) HelloGrpcClient {
	return &helloGrpcClient{cc}
}

func (c *helloGrpcClient) GreetServer(ctx context.Context, in *GreetRequest, opts ...grpc.CallOption) (*GreetMessage, error) {
	out := new(GreetMessage)
	err := c.cc.Invoke(ctx, "/HelloGrpc.HelloGrpc/GreetServer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *helloGrpcClient) AddRoom(ctx context.Context, in *RoomRequest, opts ...grpc.CallOption) (*RoomInfo, error) {
	out := new(RoomInfo)
	err := c.cc.Invoke(ctx, "/HelloGrpc.HelloGrpc/AddRoom", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *helloGrpcClient) GetRoomInfo(ctx context.Context, in *RoomRequest, opts ...grpc.CallOption) (*RoomInfo, error) {
	out := new(RoomInfo)
	err := c.cc.Invoke(ctx, "/HelloGrpc.HelloGrpc/GetRoomInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *helloGrpcClient) GetRooms(ctx context.Context, in *Null, opts ...grpc.CallOption) (*RoomList, error) {
	out := new(RoomList)
	err := c.cc.Invoke(ctx, "/HelloGrpc.HelloGrpc/GetRooms", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *helloGrpcClient) SendMessage(ctx context.Context, opts ...grpc.CallOption) (HelloGrpc_SendMessageClient, error) {
	stream, err := c.cc.NewStream(ctx, &_HelloGrpc_serviceDesc.Streams[0], "/HelloGrpc.HelloGrpc/SendMessage", opts...)
	if err != nil {
		return nil, err
	}
	x := &helloGrpcSendMessageClient{stream}
	return x, nil
}

type HelloGrpc_SendMessageClient interface {
	Send(*SendRequest) error
	CloseAndRecv() (*SendResult, error)
	grpc.ClientStream
}

type helloGrpcSendMessageClient struct {
	grpc.ClientStream
}

func (x *helloGrpcSendMessageClient) Send(m *SendRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *helloGrpcSendMessageClient) CloseAndRecv() (*SendResult, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(SendResult)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *helloGrpcClient) GetMessages(ctx context.Context, in *MessagesRequest, opts ...grpc.CallOption) (HelloGrpc_GetMessagesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_HelloGrpc_serviceDesc.Streams[1], "/HelloGrpc.HelloGrpc/GetMessages", opts...)
	if err != nil {
		return nil, err
	}
	x := &helloGrpcGetMessagesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type HelloGrpc_GetMessagesClient interface {
	Recv() (*Message, error)
	grpc.ClientStream
}

type helloGrpcGetMessagesClient struct {
	grpc.ClientStream
}

func (x *helloGrpcGetMessagesClient) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// HelloGrpcServer is the server API for HelloGrpc service.
type HelloGrpcServer interface {
	// Sends a greeting
	GreetServer(context.Context, *GreetRequest) (*GreetMessage, error)
	// チャットルーム追加
	AddRoom(context.Context, *RoomRequest) (*RoomInfo, error)
	// チャットルーム情報取得
	GetRoomInfo(context.Context, *RoomRequest) (*RoomInfo, error)
	// チャットルーム一覧取得
	GetRooms(context.Context, *Null) (*RoomList, error)
	// Client-side streamingを用いてメッセージを交換する
	SendMessage(HelloGrpc_SendMessageServer) error
	// Server-side streamingを用いてメッセージを交換する
	GetMessages(*MessagesRequest, HelloGrpc_GetMessagesServer) error
}

// UnimplementedHelloGrpcServer can be embedded to have forward compatible implementations.
type UnimplementedHelloGrpcServer struct {
}

func (*UnimplementedHelloGrpcServer) GreetServer(ctx context.Context, req *GreetRequest) (*GreetMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GreetServer not implemented")
}
func (*UnimplementedHelloGrpcServer) AddRoom(ctx context.Context, req *RoomRequest) (*RoomInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddRoom not implemented")
}
func (*UnimplementedHelloGrpcServer) GetRoomInfo(ctx context.Context, req *RoomRequest) (*RoomInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRoomInfo not implemented")
}
func (*UnimplementedHelloGrpcServer) GetRooms(ctx context.Context, req *Null) (*RoomList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRooms not implemented")
}
func (*UnimplementedHelloGrpcServer) SendMessage(srv HelloGrpc_SendMessageServer) error {
	return status.Errorf(codes.Unimplemented, "method SendMessage not implemented")
}
func (*UnimplementedHelloGrpcServer) GetMessages(req *MessagesRequest, srv HelloGrpc_GetMessagesServer) error {
	return status.Errorf(codes.Unimplemented, "method GetMessages not implemented")
}

func RegisterHelloGrpcServer(s *grpc.Server, srv HelloGrpcServer) {
	s.RegisterService(&_HelloGrpc_serviceDesc, srv)
}

func _HelloGrpc_GreetServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GreetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloGrpcServer).GreetServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/HelloGrpc.HelloGrpc/GreetServer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloGrpcServer).GreetServer(ctx, req.(*GreetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HelloGrpc_AddRoom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoomRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloGrpcServer).AddRoom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/HelloGrpc.HelloGrpc/AddRoom",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloGrpcServer).AddRoom(ctx, req.(*RoomRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HelloGrpc_GetRoomInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoomRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloGrpcServer).GetRoomInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/HelloGrpc.HelloGrpc/GetRoomInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloGrpcServer).GetRoomInfo(ctx, req.(*RoomRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HelloGrpc_GetRooms_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Null)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloGrpcServer).GetRooms(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/HelloGrpc.HelloGrpc/GetRooms",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloGrpcServer).GetRooms(ctx, req.(*Null))
	}
	return interceptor(ctx, in, info, handler)
}

func _HelloGrpc_SendMessage_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(HelloGrpcServer).SendMessage(&helloGrpcSendMessageServer{stream})
}

type HelloGrpc_SendMessageServer interface {
	SendAndClose(*SendResult) error
	Recv() (*SendRequest, error)
	grpc.ServerStream
}

type helloGrpcSendMessageServer struct {
	grpc.ServerStream
}

func (x *helloGrpcSendMessageServer) SendAndClose(m *SendResult) error {
	return x.ServerStream.SendMsg(m)
}

func (x *helloGrpcSendMessageServer) Recv() (*SendRequest, error) {
	m := new(SendRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _HelloGrpc_GetMessages_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(MessagesRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(HelloGrpcServer).GetMessages(m, &helloGrpcGetMessagesServer{stream})
}

type HelloGrpc_GetMessagesServer interface {
	Send(*Message) error
	grpc.ServerStream
}

type helloGrpcGetMessagesServer struct {
	grpc.ServerStream
}

func (x *helloGrpcGetMessagesServer) Send(m *Message) error {
	return x.ServerStream.SendMsg(m)
}

var _HelloGrpc_serviceDesc = grpc.ServiceDesc{
	ServiceName: "HelloGrpc.HelloGrpc",
	HandlerType: (*HelloGrpcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GreetServer",
			Handler:    _HelloGrpc_GreetServer_Handler,
		},
		{
			MethodName: "AddRoom",
			Handler:    _HelloGrpc_AddRoom_Handler,
		},
		{
			MethodName: "GetRoomInfo",
			Handler:    _HelloGrpc_GetRoomInfo_Handler,
		},
		{
			MethodName: "GetRooms",
			Handler:    _HelloGrpc_GetRooms_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SendMessage",
			Handler:       _HelloGrpc_SendMessage_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "GetMessages",
			Handler:       _HelloGrpc_GetMessages_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "hellogrpc.proto",
}