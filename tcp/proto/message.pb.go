// Code generated by protoc-gen-go. DO NOT EDIT.
// source: message.proto

/*
Package message is a generated protocol buffer package.

It is generated from these files:
	message.proto

It has these top-level messages:
	Message
*/
package message

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

type Message struct {
	Cmd      string `protobuf:"bytes,1,opt,name=cmd" json:"cmd,omitempty"`
	Sender   string `protobuf:"bytes,2,opt,name=sender" json:"sender,omitempty"`
	Receiver string `protobuf:"bytes,3,opt,name=receiver" json:"receiver,omitempty"`
	Body     []byte `protobuf:"bytes,4,opt,name=body,proto3" json:"body,omitempty"`
}

func (m *Message) Reset()                    { *m = Message{} }
func (m *Message) String() string            { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()               {}
func (*Message) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Message) GetCmd() string {
	if m != nil {
		return m.Cmd
	}
	return ""
}

func (m *Message) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *Message) GetReceiver() string {
	if m != nil {
		return m.Receiver
	}
	return ""
}

func (m *Message) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

func init() {
	proto.RegisterType((*Message)(nil), "Message")
}

func init() { proto.RegisterFile("message.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 117 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcd, 0x4d, 0x2d, 0x2e,
	0x4e, 0x4c, 0x4f, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0x4a, 0xe6, 0x62, 0xf7, 0x85, 0x08,
	0x08, 0x09, 0x70, 0x31, 0x27, 0xe7, 0xa6, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x81, 0x98,
	0x42, 0x62, 0x5c, 0x6c, 0xc5, 0xa9, 0x79, 0x29, 0xa9, 0x45, 0x12, 0x4c, 0x60, 0x41, 0x28, 0x4f,
	0x48, 0x8a, 0x8b, 0xa3, 0x28, 0x35, 0x39, 0x35, 0xb3, 0x2c, 0xb5, 0x48, 0x82, 0x19, 0x2c, 0x03,
	0xe7, 0x0b, 0x09, 0x71, 0xb1, 0x24, 0xe5, 0xa7, 0x54, 0x4a, 0xb0, 0x28, 0x30, 0x6a, 0xf0, 0x04,
	0x81, 0xd9, 0x49, 0x6c, 0x60, 0xbb, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xaa, 0x56, 0xd3,
	0x83, 0x7c, 0x00, 0x00, 0x00,
}
