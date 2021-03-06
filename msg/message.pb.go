// Code generated by protoc-gen-go.
// source: message.proto
// DO NOT EDIT!

/*
Package msg is a generated protocol buffer package.

It is generated from these files:
	message.proto

It has these top-level messages:
	Message
*/
package msg

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type MessageType int32

const (
	MessageType_SINGLE MessageType = 0
	MessageType_BOUNCE MessageType = 1
	MessageType_ERROR  MessageType = 2
)

var MessageType_name = map[int32]string{
	0: "SINGLE",
	1: "BOUNCE",
	2: "ERROR",
}
var MessageType_value = map[string]int32{
	"SINGLE": 0,
	"BOUNCE": 1,
	"ERROR":  2,
}

func (x MessageType) String() string {
	return proto.EnumName(MessageType_name, int32(x))
}
func (MessageType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Message struct {
	Type    MessageType `protobuf:"varint,1,opt,name=type,enum=msg.MessageType" json:"type,omitempty"`
	Sender  string      `protobuf:"bytes,2,opt,name=sender" json:"sender,omitempty"`
	Message string      `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
}

func (m *Message) Reset()                    { *m = Message{} }
func (m *Message) String() string            { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()               {}
func (*Message) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Message) GetType() MessageType {
	if m != nil {
		return m.Type
	}
	return MessageType_SINGLE
}

func (m *Message) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *Message) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*Message)(nil), "msg.Message")
	proto.RegisterEnum("msg.MessageType", MessageType_name, MessageType_value)
}

func init() { proto.RegisterFile("message.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 155 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcd, 0x4d, 0x2d, 0x2e,
	0x4e, 0x4c, 0x4f, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xce, 0x2d, 0x4e, 0x57, 0x4a,
	0xe4, 0x62, 0xf7, 0x85, 0x88, 0x0a, 0xa9, 0x70, 0xb1, 0x94, 0x54, 0x16, 0xa4, 0x4a, 0x30, 0x2a,
	0x30, 0x6a, 0xf0, 0x19, 0x09, 0xe8, 0xe5, 0x16, 0xa7, 0xeb, 0x41, 0xe5, 0x42, 0x2a, 0x0b, 0x52,
	0x83, 0xc0, 0xb2, 0x42, 0x62, 0x5c, 0x6c, 0xc5, 0xa9, 0x79, 0x29, 0xa9, 0x45, 0x12, 0x4c, 0x0a,
	0x8c, 0x1a, 0x9c, 0x41, 0x50, 0x9e, 0x90, 0x04, 0x17, 0x3b, 0xd4, 0x78, 0x09, 0x66, 0xb0, 0x04,
	0x8c, 0xab, 0x65, 0xc0, 0xc5, 0x8d, 0x64, 0x8c, 0x10, 0x17, 0x17, 0x5b, 0xb0, 0xa7, 0x9f, 0xbb,
	0x8f, 0xab, 0x00, 0x03, 0x88, 0xed, 0xe4, 0x1f, 0xea, 0xe7, 0xec, 0x2a, 0xc0, 0x28, 0xc4, 0xc9,
	0xc5, 0xea, 0x1a, 0x14, 0xe4, 0x1f, 0x24, 0xc0, 0x94, 0xc4, 0x06, 0x76, 0xa0, 0x31, 0x20, 0x00,
	0x00, 0xff, 0xff, 0x1a, 0x4f, 0xb5, 0x8b, 0xb1, 0x00, 0x00, 0x00,
}
