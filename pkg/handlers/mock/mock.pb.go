// Code generated by protoc-gen-go. DO NOT EDIT.
// source: mock.proto

package mock

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

type Mock struct {
	Mock                 string   `protobuf:"bytes,1,opt,name=Mock,proto3" json:"Mock,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Mock) Reset()         { *m = Mock{} }
func (m *Mock) String() string { return proto.CompactTextString(m) }
func (*Mock) ProtoMessage()    {}
func (*Mock) Descriptor() ([]byte, []int) {
	return fileDescriptor_mock_7fcf3b0ff111bc44, []int{0}
}
func (m *Mock) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Mock.Unmarshal(m, b)
}
func (m *Mock) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Mock.Marshal(b, m, deterministic)
}
func (dst *Mock) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Mock.Merge(dst, src)
}
func (m *Mock) XXX_Size() int {
	return xxx_messageInfo_Mock.Size(m)
}
func (m *Mock) XXX_DiscardUnknown() {
	xxx_messageInfo_Mock.DiscardUnknown(m)
}

var xxx_messageInfo_Mock proto.InternalMessageInfo

func (m *Mock) GetMock() string {
	if m != nil {
		return m.Mock
	}
	return ""
}

func init() {
	proto.RegisterType((*Mock)(nil), "mock.Mock")
}

func init() { proto.RegisterFile("mock.proto", fileDescriptor_mock_7fcf3b0ff111bc44) }

var fileDescriptor_mock_7fcf3b0ff111bc44 = []byte{
	// 70 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xca, 0xcd, 0x4f, 0xce,
	0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x01, 0xb1, 0x95, 0xa4, 0xb8, 0x58, 0x7c, 0xf3,
	0x93, 0xb3, 0x85, 0x84, 0x20, 0xb4, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x98, 0xed, 0xc4,
	0x16, 0x05, 0x56, 0x93, 0xc4, 0x06, 0xd6, 0x60, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x78, 0xf4,
	0xc8, 0x8d, 0x3e, 0x00, 0x00, 0x00,
}
