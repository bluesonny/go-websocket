// Code generated by protoc-gen-go. DO NOT EDIT.
// source: servers/grpc.proto

package pb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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

type Send2ClientReq struct {
	SystemId             string      `protobuf:"bytes,1,opt,name=systemId,proto3" json:"systemId,omitempty"`
	MessageId            string      `protobuf:"bytes,2,opt,name=messageId,proto3" json:"messageId,omitempty"`
	SendUserId           string      `protobuf:"bytes,3,opt,name=sendUserId,proto3" json:"sendUserId,omitempty"`
	ClientId             string      `protobuf:"bytes,4,opt,name=clientId,proto3" json:"clientId,omitempty"`
	Code                 int32       `protobuf:"varint,5,opt,name=code,proto3" json:"code,omitempty"`
	Message              string      `protobuf:"bytes,6,opt,name=message,proto3" json:"message,omitempty"`
	Data                 interface{} `protobuf:"bytes,7,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Send2ClientReq) Reset()         { *m = Send2ClientReq{} }
func (m *Send2ClientReq) String() string { return proto.CompactTextString(m) }
func (*Send2ClientReq) ProtoMessage()    {}
func (*Send2ClientReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_30edbedfc0173ed4, []int{0}
}

func (m *Send2ClientReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Send2ClientReq.Unmarshal(m, b)
}
func (m *Send2ClientReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Send2ClientReq.Marshal(b, m, deterministic)
}
func (m *Send2ClientReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Send2ClientReq.Merge(m, src)
}
func (m *Send2ClientReq) XXX_Size() int {
	return xxx_messageInfo_Send2ClientReq.Size(m)
}
func (m *Send2ClientReq) XXX_DiscardUnknown() {
	xxx_messageInfo_Send2ClientReq.DiscardUnknown(m)
}

var xxx_messageInfo_Send2ClientReq proto.InternalMessageInfo

func (m *Send2ClientReq) GetSystemId() string {
	if m != nil {
		return m.SystemId
	}
	return ""
}

func (m *Send2ClientReq) GetMessageId() string {
	if m != nil {
		return m.MessageId
	}
	return ""
}

func (m *Send2ClientReq) GetSendUserId() string {
	if m != nil {
		return m.SendUserId
	}
	return ""
}

func (m *Send2ClientReq) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

func (m *Send2ClientReq) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Send2ClientReq) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *Send2ClientReq) GetData() interface{} {
	if m != nil {
		return m.Data
	}
	return ""
}

type CloseClientReq struct {
	SystemId             string   `protobuf:"bytes,1,opt,name=systemId,proto3" json:"systemId,omitempty"`
	ClientId             string   `protobuf:"bytes,2,opt,name=clientId,proto3" json:"clientId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CloseClientReq) Reset()         { *m = CloseClientReq{} }
func (m *CloseClientReq) String() string { return proto.CompactTextString(m) }
func (*CloseClientReq) ProtoMessage()    {}
func (*CloseClientReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_30edbedfc0173ed4, []int{1}
}

func (m *CloseClientReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CloseClientReq.Unmarshal(m, b)
}
func (m *CloseClientReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CloseClientReq.Marshal(b, m, deterministic)
}
func (m *CloseClientReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CloseClientReq.Merge(m, src)
}
func (m *CloseClientReq) XXX_Size() int {
	return xxx_messageInfo_CloseClientReq.Size(m)
}
func (m *CloseClientReq) XXX_DiscardUnknown() {
	xxx_messageInfo_CloseClientReq.DiscardUnknown(m)
}

var xxx_messageInfo_CloseClientReq proto.InternalMessageInfo

func (m *CloseClientReq) GetSystemId() string {
	if m != nil {
		return m.SystemId
	}
	return ""
}

func (m *CloseClientReq) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

type BindGroupReq struct {
	SystemId             string   `protobuf:"bytes,1,opt,name=systemId,proto3" json:"systemId,omitempty"`
	GroupName            string   `protobuf:"bytes,2,opt,name=groupName,proto3" json:"groupName,omitempty"`
	ClientId             string   `protobuf:"bytes,3,opt,name=clientId,proto3" json:"clientId,omitempty"`
	UserId               string   `protobuf:"bytes,4,opt,name=userId,proto3" json:"userId,omitempty"`
	Extend               string   `protobuf:"bytes,5,opt,name=extend,proto3" json:"extend,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BindGroupReq) Reset()         { *m = BindGroupReq{} }
func (m *BindGroupReq) String() string { return proto.CompactTextString(m) }
func (*BindGroupReq) ProtoMessage()    {}
func (*BindGroupReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_30edbedfc0173ed4, []int{2}
}

func (m *BindGroupReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BindGroupReq.Unmarshal(m, b)
}
func (m *BindGroupReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BindGroupReq.Marshal(b, m, deterministic)
}
func (m *BindGroupReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BindGroupReq.Merge(m, src)
}
func (m *BindGroupReq) XXX_Size() int {
	return xxx_messageInfo_BindGroupReq.Size(m)
}
func (m *BindGroupReq) XXX_DiscardUnknown() {
	xxx_messageInfo_BindGroupReq.DiscardUnknown(m)
}

var xxx_messageInfo_BindGroupReq proto.InternalMessageInfo

func (m *BindGroupReq) GetSystemId() string {
	if m != nil {
		return m.SystemId
	}
	return ""
}

func (m *BindGroupReq) GetGroupName() string {
	if m != nil {
		return m.GroupName
	}
	return ""
}

func (m *BindGroupReq) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

func (m *BindGroupReq) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *BindGroupReq) GetExtend() string {
	if m != nil {
		return m.Extend
	}
	return ""
}

type Send2GroupReq struct {
	SystemId             string      `protobuf:"bytes,1,opt,name=systemId,proto3" json:"systemId,omitempty"`
	MessageId            string      `protobuf:"bytes,2,opt,name=messageId,proto3" json:"messageId,omitempty"`
	SendUserId           string      `protobuf:"bytes,3,opt,name=sendUserId,proto3" json:"sendUserId,omitempty"`
	GroupName            string      `protobuf:"bytes,4,opt,name=groupName,proto3" json:"groupName,omitempty"`
	Code                 int32       `protobuf:"varint,5,opt,name=code,proto3" json:"code,omitempty"`
	Message              string      `protobuf:"bytes,6,opt,name=message,proto3" json:"message,omitempty"`
	Data                 interface{} `protobuf:"bytes,7,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Send2GroupReq) Reset()         { *m = Send2GroupReq{} }
func (m *Send2GroupReq) String() string { return proto.CompactTextString(m) }
func (*Send2GroupReq) ProtoMessage()    {}
func (*Send2GroupReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_30edbedfc0173ed4, []int{3}
}

func (m *Send2GroupReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Send2GroupReq.Unmarshal(m, b)
}
func (m *Send2GroupReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Send2GroupReq.Marshal(b, m, deterministic)
}
func (m *Send2GroupReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Send2GroupReq.Merge(m, src)
}
func (m *Send2GroupReq) XXX_Size() int {
	return xxx_messageInfo_Send2GroupReq.Size(m)
}
func (m *Send2GroupReq) XXX_DiscardUnknown() {
	xxx_messageInfo_Send2GroupReq.DiscardUnknown(m)
}

var xxx_messageInfo_Send2GroupReq proto.InternalMessageInfo

func (m *Send2GroupReq) GetSystemId() string {
	if m != nil {
		return m.SystemId
	}
	return ""
}

func (m *Send2GroupReq) GetMessageId() string {
	if m != nil {
		return m.MessageId
	}
	return ""
}

func (m *Send2GroupReq) GetSendUserId() string {
	if m != nil {
		return m.SendUserId
	}
	return ""
}

func (m *Send2GroupReq) GetGroupName() string {
	if m != nil {
		return m.GroupName
	}
	return ""
}

func (m *Send2GroupReq) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Send2GroupReq) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *Send2GroupReq) GetData() interface{} {
	if m != nil {
		return m.Data
	}
	return ""
}

type Send2SystemReq struct {
	SystemId             string   `protobuf:"bytes,1,opt,name=systemId,proto3" json:"systemId,omitempty"`
	MessageId            string   `protobuf:"bytes,2,opt,name=messageId,proto3" json:"messageId,omitempty"`
	SendUserId           string   `protobuf:"bytes,3,opt,name=sendUserId,proto3" json:"sendUserId,omitempty"`
	Code                 int32    `protobuf:"varint,4,opt,name=code,proto3" json:"code,omitempty"`
	Message              string   `protobuf:"bytes,5,opt,name=message,proto3" json:"message,omitempty"`
	Data                 string   `protobuf:"bytes,6,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Send2SystemReq) Reset()         { *m = Send2SystemReq{} }
func (m *Send2SystemReq) String() string { return proto.CompactTextString(m) }
func (*Send2SystemReq) ProtoMessage()    {}
func (*Send2SystemReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_30edbedfc0173ed4, []int{4}
}

func (m *Send2SystemReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Send2SystemReq.Unmarshal(m, b)
}
func (m *Send2SystemReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Send2SystemReq.Marshal(b, m, deterministic)
}
func (m *Send2SystemReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Send2SystemReq.Merge(m, src)
}
func (m *Send2SystemReq) XXX_Size() int {
	return xxx_messageInfo_Send2SystemReq.Size(m)
}
func (m *Send2SystemReq) XXX_DiscardUnknown() {
	xxx_messageInfo_Send2SystemReq.DiscardUnknown(m)
}

var xxx_messageInfo_Send2SystemReq proto.InternalMessageInfo

func (m *Send2SystemReq) GetSystemId() string {
	if m != nil {
		return m.SystemId
	}
	return ""
}

func (m *Send2SystemReq) GetMessageId() string {
	if m != nil {
		return m.MessageId
	}
	return ""
}

func (m *Send2SystemReq) GetSendUserId() string {
	if m != nil {
		return m.SendUserId
	}
	return ""
}

func (m *Send2SystemReq) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Send2SystemReq) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *Send2SystemReq) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

type GetGroupClientsReq struct {
	SystemId             string   `protobuf:"bytes,1,opt,name=systemId,proto3" json:"systemId,omitempty"`
	GroupName            string   `protobuf:"bytes,2,opt,name=groupName,proto3" json:"groupName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetGroupClientsReq) Reset()         { *m = GetGroupClientsReq{} }
func (m *GetGroupClientsReq) String() string { return proto.CompactTextString(m) }
func (*GetGroupClientsReq) ProtoMessage()    {}
func (*GetGroupClientsReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_30edbedfc0173ed4, []int{5}
}

func (m *GetGroupClientsReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetGroupClientsReq.Unmarshal(m, b)
}
func (m *GetGroupClientsReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetGroupClientsReq.Marshal(b, m, deterministic)
}
func (m *GetGroupClientsReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetGroupClientsReq.Merge(m, src)
}
func (m *GetGroupClientsReq) XXX_Size() int {
	return xxx_messageInfo_GetGroupClientsReq.Size(m)
}
func (m *GetGroupClientsReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetGroupClientsReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetGroupClientsReq proto.InternalMessageInfo

func (m *GetGroupClientsReq) GetSystemId() string {
	if m != nil {
		return m.SystemId
	}
	return ""
}

func (m *GetGroupClientsReq) GetGroupName() string {
	if m != nil {
		return m.GroupName
	}
	return ""
}

type Send2ClientReply struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Send2ClientReply) Reset()         { *m = Send2ClientReply{} }
func (m *Send2ClientReply) String() string { return proto.CompactTextString(m) }
func (*Send2ClientReply) ProtoMessage()    {}
func (*Send2ClientReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_30edbedfc0173ed4, []int{6}
}

func (m *Send2ClientReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Send2ClientReply.Unmarshal(m, b)
}
func (m *Send2ClientReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Send2ClientReply.Marshal(b, m, deterministic)
}
func (m *Send2ClientReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Send2ClientReply.Merge(m, src)
}
func (m *Send2ClientReply) XXX_Size() int {
	return xxx_messageInfo_Send2ClientReply.Size(m)
}
func (m *Send2ClientReply) XXX_DiscardUnknown() {
	xxx_messageInfo_Send2ClientReply.DiscardUnknown(m)
}

var xxx_messageInfo_Send2ClientReply proto.InternalMessageInfo

type CloseClientReply struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CloseClientReply) Reset()         { *m = CloseClientReply{} }
func (m *CloseClientReply) String() string { return proto.CompactTextString(m) }
func (*CloseClientReply) ProtoMessage()    {}
func (*CloseClientReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_30edbedfc0173ed4, []int{7}
}

func (m *CloseClientReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CloseClientReply.Unmarshal(m, b)
}
func (m *CloseClientReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CloseClientReply.Marshal(b, m, deterministic)
}
func (m *CloseClientReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CloseClientReply.Merge(m, src)
}
func (m *CloseClientReply) XXX_Size() int {
	return xxx_messageInfo_CloseClientReply.Size(m)
}
func (m *CloseClientReply) XXX_DiscardUnknown() {
	xxx_messageInfo_CloseClientReply.DiscardUnknown(m)
}

var xxx_messageInfo_CloseClientReply proto.InternalMessageInfo

type BindGroupReply struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BindGroupReply) Reset()         { *m = BindGroupReply{} }
func (m *BindGroupReply) String() string { return proto.CompactTextString(m) }
func (*BindGroupReply) ProtoMessage()    {}
func (*BindGroupReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_30edbedfc0173ed4, []int{8}
}

func (m *BindGroupReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BindGroupReply.Unmarshal(m, b)
}
func (m *BindGroupReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BindGroupReply.Marshal(b, m, deterministic)
}
func (m *BindGroupReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BindGroupReply.Merge(m, src)
}
func (m *BindGroupReply) XXX_Size() int {
	return xxx_messageInfo_BindGroupReply.Size(m)
}
func (m *BindGroupReply) XXX_DiscardUnknown() {
	xxx_messageInfo_BindGroupReply.DiscardUnknown(m)
}

var xxx_messageInfo_BindGroupReply proto.InternalMessageInfo

type Send2GroupReply struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Send2GroupReply) Reset()         { *m = Send2GroupReply{} }
func (m *Send2GroupReply) String() string { return proto.CompactTextString(m) }
func (*Send2GroupReply) ProtoMessage()    {}
func (*Send2GroupReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_30edbedfc0173ed4, []int{9}
}

func (m *Send2GroupReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Send2GroupReply.Unmarshal(m, b)
}
func (m *Send2GroupReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Send2GroupReply.Marshal(b, m, deterministic)
}
func (m *Send2GroupReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Send2GroupReply.Merge(m, src)
}
func (m *Send2GroupReply) XXX_Size() int {
	return xxx_messageInfo_Send2GroupReply.Size(m)
}
func (m *Send2GroupReply) XXX_DiscardUnknown() {
	xxx_messageInfo_Send2GroupReply.DiscardUnknown(m)
}

var xxx_messageInfo_Send2GroupReply proto.InternalMessageInfo

type Send2SystemReply struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Send2SystemReply) Reset()         { *m = Send2SystemReply{} }
func (m *Send2SystemReply) String() string { return proto.CompactTextString(m) }
func (*Send2SystemReply) ProtoMessage()    {}
func (*Send2SystemReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_30edbedfc0173ed4, []int{10}
}

func (m *Send2SystemReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Send2SystemReply.Unmarshal(m, b)
}
func (m *Send2SystemReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Send2SystemReply.Marshal(b, m, deterministic)
}
func (m *Send2SystemReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Send2SystemReply.Merge(m, src)
}
func (m *Send2SystemReply) XXX_Size() int {
	return xxx_messageInfo_Send2SystemReply.Size(m)
}
func (m *Send2SystemReply) XXX_DiscardUnknown() {
	xxx_messageInfo_Send2SystemReply.DiscardUnknown(m)
}

var xxx_messageInfo_Send2SystemReply proto.InternalMessageInfo

type GetGroupClientsReply struct {
	List                 []string `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetGroupClientsReply) Reset()         { *m = GetGroupClientsReply{} }
func (m *GetGroupClientsReply) String() string { return proto.CompactTextString(m) }
func (*GetGroupClientsReply) ProtoMessage()    {}
func (*GetGroupClientsReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_30edbedfc0173ed4, []int{11}
}

func (m *GetGroupClientsReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetGroupClientsReply.Unmarshal(m, b)
}
func (m *GetGroupClientsReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetGroupClientsReply.Marshal(b, m, deterministic)
}
func (m *GetGroupClientsReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetGroupClientsReply.Merge(m, src)
}
func (m *GetGroupClientsReply) XXX_Size() int {
	return xxx_messageInfo_GetGroupClientsReply.Size(m)
}
func (m *GetGroupClientsReply) XXX_DiscardUnknown() {
	xxx_messageInfo_GetGroupClientsReply.DiscardUnknown(m)
}

var xxx_messageInfo_GetGroupClientsReply proto.InternalMessageInfo

func (m *GetGroupClientsReply) GetList() []string {
	if m != nil {
		return m.List
	}
	return nil
}

func init() {
	proto.RegisterType((*Send2ClientReq)(nil), "Send2ClientReq")
	proto.RegisterType((*CloseClientReq)(nil), "CloseClientReq")
	proto.RegisterType((*BindGroupReq)(nil), "BindGroupReq")
	proto.RegisterType((*Send2GroupReq)(nil), "Send2GroupReq")
	proto.RegisterType((*Send2SystemReq)(nil), "Send2SystemReq")
	proto.RegisterType((*GetGroupClientsReq)(nil), "GetGroupClientsReq")
	proto.RegisterType((*Send2ClientReply)(nil), "Send2ClientReply")
	proto.RegisterType((*CloseClientReply)(nil), "CloseClientReply")
	proto.RegisterType((*BindGroupReply)(nil), "BindGroupReply")
	proto.RegisterType((*Send2GroupReply)(nil), "Send2GroupReply")
	proto.RegisterType((*Send2SystemReply)(nil), "Send2SystemReply")
	proto.RegisterType((*GetGroupClientsReply)(nil), "GetGroupClientsReply")
}

func init() { proto.RegisterFile("servers/grpc.proto", fileDescriptor_30edbedfc0173ed4) }

var fileDescriptor_30edbedfc0173ed4 = []byte{
	// 477 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x54, 0xcb, 0x6e, 0x1a, 0x31,
	0x14, 0x9d, 0xe1, 0x95, 0x72, 0x1b, 0x18, 0xb8, 0x7d, 0xc8, 0x1a, 0x55, 0x15, 0x9a, 0x55, 0x54,
	0xa9, 0x6e, 0x95, 0x7c, 0x40, 0xa5, 0xb0, 0x48, 0xd9, 0x64, 0x01, 0xea, 0xa6, 0x3b, 0x82, 0xaf,
	0x10, 0xd2, 0xbc, 0x3a, 0x9e, 0x44, 0xe5, 0x3f, 0xfa, 0x19, 0xfd, 0x90, 0x76, 0xd7, 0x4f, 0xaa,
	0xec, 0xc1, 0x83, 0x3d, 0xb0, 0x40, 0x42, 0xd9, 0xf9, 0x1e, 0xbf, 0xce, 0x39, 0x3e, 0xd7, 0x80,
	0x92, 0x8a, 0x27, 0x2a, 0xe4, 0xa7, 0x75, 0x91, 0xaf, 0x78, 0x5e, 0x64, 0x65, 0x16, 0xfd, 0xf1,
	0x61, 0xb8, 0xa0, 0x54, 0x5c, 0x4f, 0xe3, 0x0d, 0xa5, 0xe5, 0x9c, 0x7e, 0x60, 0x08, 0x2f, 0xe4,
	0x56, 0x96, 0x94, 0xcc, 0x04, 0xf3, 0x27, 0xfe, 0x55, 0x7f, 0x5e, 0xd7, 0xf8, 0x0e, 0xfa, 0x09,
	0x49, 0xb9, 0x5c, 0xd3, 0x4c, 0xb0, 0x96, 0x9e, 0xdc, 0x03, 0xf8, 0x1e, 0x40, 0x52, 0x2a, 0xbe,
	0x49, 0x2a, 0x66, 0x82, 0xb5, 0xf5, 0xb4, 0x85, 0xa8, 0x93, 0x57, 0xfa, 0x9a, 0x99, 0x60, 0x9d,
	0xea, 0x64, 0x53, 0x23, 0x42, 0x67, 0x95, 0x09, 0x62, 0xdd, 0x89, 0x7f, 0xd5, 0x9d, 0xeb, 0x31,
	0x32, 0xb8, 0xd8, 0x1d, 0xce, 0x7a, 0x7a, 0xb9, 0x29, 0xd5, 0x6a, 0xb1, 0x2c, 0x97, 0xec, 0x42,
	0xc3, 0x7a, 0x1c, 0x7d, 0x85, 0xe1, 0x34, 0xce, 0x24, 0x9d, 0xa6, 0xc4, 0xe6, 0xd2, 0x72, 0xb9,
	0x44, 0xbf, 0x7c, 0xb8, 0xbc, 0xdd, 0xa4, 0xe2, 0xae, 0xc8, 0x1e, 0xf3, 0x13, 0x2c, 0x59, 0xab,
	0x75, 0xf7, 0xcb, 0x84, 0x8c, 0x25, 0x35, 0xe0, 0x5c, 0xd3, 0x6e, 0x48, 0x7e, 0x0b, 0xbd, 0xc7,
	0xca, 0xaa, 0xca, 0x8c, 0x5d, 0xa5, 0x70, 0xfa, 0x59, 0x52, 0x2a, 0xb4, 0x19, 0xfd, 0xf9, 0xae,
	0x8a, 0xfe, 0xfa, 0x30, 0xd0, 0x6f, 0x75, 0x2a, 0xaf, 0x33, 0x9e, 0xca, 0x51, 0xd5, 0x69, 0xaa,
	0x3a, 0xff, 0xb1, 0x7e, 0x9b, 0xdc, 0x2d, 0x34, 0xdf, 0xe7, 0x15, 0x63, 0xe8, 0x76, 0x8e, 0xd3,
	0xed, 0x1e, 0xa7, 0xdb, 0xb3, 0xe8, 0xde, 0x03, 0xde, 0x51, 0xa9, 0x7d, 0xaf, 0xe2, 0x25, 0xcf,
	0x8a, 0x45, 0x84, 0x30, 0x72, 0xba, 0x2e, 0x8f, 0xb7, 0x0a, 0x73, 0xf2, 0xab, 0xb0, 0x11, 0x0c,
	0xad, 0x20, 0x2a, 0x64, 0x0c, 0x81, 0x9d, 0x81, 0xdd, 0x46, 0xc7, 0x4a, 0x85, 0x7d, 0x80, 0xd7,
	0x07, 0x84, 0xf3, 0x78, 0xab, 0xc4, 0xc5, 0x1b, 0x59, 0x32, 0x7f, 0xd2, 0x56, 0xe2, 0xd4, 0xf8,
	0xfa, 0x5f, 0x0b, 0x06, 0xd3, 0x2c, 0x49, 0xb2, 0x74, 0x41, 0xc5, 0xd3, 0x66, 0x45, 0x78, 0x03,
	0x2f, 0x2d, 0x7a, 0x18, 0x70, 0xf7, 0x8b, 0x08, 0xc7, 0xfc, 0x80, 0xbd, 0xa7, 0x36, 0x59, 0xfc,
	0x31, 0xe0, 0x6e, 0x37, 0x86, 0x63, 0x7e, 0x20, 0xcf, 0xc3, 0x8f, 0xd0, 0xaf, 0x05, 0xe2, 0x80,
	0xdb, 0x5d, 0x17, 0x06, 0xbc, 0xa1, 0xdd, 0xc3, 0xcf, 0x00, 0x7b, 0xf5, 0x38, 0xe4, 0x4e, 0x3b,
	0x84, 0x23, 0xde, 0xb4, 0xc6, 0xab, 0xa5, 0x54, 0xe6, 0x18, 0x29, 0x75, 0xea, 0x8c, 0x14, 0xdb,
	0x3b, 0x0f, 0xbf, 0x40, 0xd0, 0x70, 0x0f, 0x5f, 0xf1, 0xc3, 0x00, 0x84, 0x6f, 0xf8, 0x31, 0x93,
	0x23, 0xef, 0xf6, 0xf2, 0x3b, 0x98, 0xcf, 0x36, 0x7f, 0x78, 0xe8, 0xe9, 0xbf, 0xf6, 0xe6, 0x7f,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x91, 0x68, 0x55, 0x86, 0x81, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CommonServiceClient is the client API for CommonService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CommonServiceClient interface {
	Send2Client(ctx context.Context, in *Send2ClientReq, opts ...grpc.CallOption) (*Send2ClientReply, error)
	CloseClient(ctx context.Context, in *CloseClientReq, opts ...grpc.CallOption) (*CloseClientReply, error)
	BindGroup(ctx context.Context, in *BindGroupReq, opts ...grpc.CallOption) (*BindGroupReply, error)
	Send2Group(ctx context.Context, in *Send2GroupReq, opts ...grpc.CallOption) (*Send2GroupReply, error)
	Send2System(ctx context.Context, in *Send2SystemReq, opts ...grpc.CallOption) (*Send2SystemReply, error)
	GetGroupClients(ctx context.Context, in *GetGroupClientsReq, opts ...grpc.CallOption) (*GetGroupClientsReply, error)
}

type commonServiceClient struct {
	cc *grpc.ClientConn
}

func NewCommonServiceClient(cc *grpc.ClientConn) CommonServiceClient {
	return &commonServiceClient{cc}
}

func (c *commonServiceClient) Send2Client(ctx context.Context, in *Send2ClientReq, opts ...grpc.CallOption) (*Send2ClientReply, error) {
	out := new(Send2ClientReply)
	err := c.cc.Invoke(ctx, "/CommonService/Send2Client", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commonServiceClient) CloseClient(ctx context.Context, in *CloseClientReq, opts ...grpc.CallOption) (*CloseClientReply, error) {
	out := new(CloseClientReply)
	err := c.cc.Invoke(ctx, "/CommonService/CloseClient", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commonServiceClient) BindGroup(ctx context.Context, in *BindGroupReq, opts ...grpc.CallOption) (*BindGroupReply, error) {
	out := new(BindGroupReply)
	err := c.cc.Invoke(ctx, "/CommonService/BindGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commonServiceClient) Send2Group(ctx context.Context, in *Send2GroupReq, opts ...grpc.CallOption) (*Send2GroupReply, error) {
	out := new(Send2GroupReply)
	err := c.cc.Invoke(ctx, "/CommonService/Send2Group", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commonServiceClient) Send2System(ctx context.Context, in *Send2SystemReq, opts ...grpc.CallOption) (*Send2SystemReply, error) {
	out := new(Send2SystemReply)
	err := c.cc.Invoke(ctx, "/CommonService/Send2System", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commonServiceClient) GetGroupClients(ctx context.Context, in *GetGroupClientsReq, opts ...grpc.CallOption) (*GetGroupClientsReply, error) {
	out := new(GetGroupClientsReply)
	err := c.cc.Invoke(ctx, "/CommonService/GetGroupClients", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CommonServiceServer is the server API for CommonService service.
type CommonServiceServer interface {
	Send2Client(context.Context, *Send2ClientReq) (*Send2ClientReply, error)
	CloseClient(context.Context, *CloseClientReq) (*CloseClientReply, error)
	BindGroup(context.Context, *BindGroupReq) (*BindGroupReply, error)
	Send2Group(context.Context, *Send2GroupReq) (*Send2GroupReply, error)
	Send2System(context.Context, *Send2SystemReq) (*Send2SystemReply, error)
	GetGroupClients(context.Context, *GetGroupClientsReq) (*GetGroupClientsReply, error)
}

func RegisterCommonServiceServer(s *grpc.Server, srv CommonServiceServer) {
	s.RegisterService(&_CommonService_serviceDesc, srv)
}

func _CommonService_Send2Client_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Send2ClientReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommonServiceServer).Send2Client(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CommonService/Send2Client",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommonServiceServer).Send2Client(ctx, req.(*Send2ClientReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommonService_CloseClient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CloseClientReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommonServiceServer).CloseClient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CommonService/CloseClient",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommonServiceServer).CloseClient(ctx, req.(*CloseClientReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommonService_BindGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BindGroupReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommonServiceServer).BindGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CommonService/BindGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommonServiceServer).BindGroup(ctx, req.(*BindGroupReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommonService_Send2Group_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Send2GroupReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommonServiceServer).Send2Group(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CommonService/Send2Group",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommonServiceServer).Send2Group(ctx, req.(*Send2GroupReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommonService_Send2System_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Send2SystemReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommonServiceServer).Send2System(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CommonService/Send2System",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommonServiceServer).Send2System(ctx, req.(*Send2SystemReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommonService_GetGroupClients_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGroupClientsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommonServiceServer).GetGroupClients(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CommonService/GetGroupClients",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommonServiceServer).GetGroupClients(ctx, req.(*GetGroupClientsReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _CommonService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "CommonService",
	HandlerType: (*CommonServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Send2Client",
			Handler:    _CommonService_Send2Client_Handler,
		},
		{
			MethodName: "CloseClient",
			Handler:    _CommonService_CloseClient_Handler,
		},
		{
			MethodName: "BindGroup",
			Handler:    _CommonService_BindGroup_Handler,
		},
		{
			MethodName: "Send2Group",
			Handler:    _CommonService_Send2Group_Handler,
		},
		{
			MethodName: "Send2System",
			Handler:    _CommonService_Send2System_Handler,
		},
		{
			MethodName: "GetGroupClients",
			Handler:    _CommonService_GetGroupClients_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "servers/grpc.proto",
}
