// Code generated by protoc-gen-go. DO NOT EDIT.
// source: enabled.proto

package v0

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"
import _ "github.com/lyft/protoc-gen-validate/validate"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Unset-able fields (used in reporting and internal state management.
type Enabled struct {
	From                 *timestamp.Timestamp `protobuf:"bytes,1,opt,name=From,proto3" json:"From,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Enabled) Reset()         { *m = Enabled{} }
func (m *Enabled) String() string { return proto.CompactTextString(m) }
func (*Enabled) ProtoMessage()    {}
func (*Enabled) Descriptor() ([]byte, []int) {
	return fileDescriptor_enabled_155a69b52cdb8b99, []int{0}
}
func (m *Enabled) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Enabled.Unmarshal(m, b)
}
func (m *Enabled) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Enabled.Marshal(b, m, deterministic)
}
func (dst *Enabled) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Enabled.Merge(dst, src)
}
func (m *Enabled) XXX_Size() int {
	return xxx_messageInfo_Enabled.Size(m)
}
func (m *Enabled) XXX_DiscardUnknown() {
	xxx_messageInfo_Enabled.DiscardUnknown(m)
}

var xxx_messageInfo_Enabled proto.InternalMessageInfo

func (m *Enabled) GetFrom() *timestamp.Timestamp {
	if m != nil {
		return m.From
	}
	return nil
}

func init() {
	proto.RegisterType((*Enabled)(nil), "v0.Enabled")
}

func init() { proto.RegisterFile("enabled.proto", fileDescriptor_enabled_155a69b52cdb8b99) }

var fileDescriptor_enabled_155a69b52cdb8b99 = []byte{
	// 144 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4d, 0xcd, 0x4b, 0x4c,
	0xca, 0x49, 0x4d, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x33, 0x90, 0x92, 0x4f,
	0xcf, 0xcf, 0x4f, 0xcf, 0x49, 0xd5, 0x07, 0x8b, 0x24, 0x95, 0xa6, 0xe9, 0x97, 0x64, 0xe6, 0xa6,
	0x16, 0x97, 0x24, 0xe6, 0x16, 0x40, 0x14, 0x49, 0x89, 0x97, 0x25, 0xe6, 0x64, 0xa6, 0x24, 0x96,
	0xa4, 0xea, 0xc3, 0x18, 0x10, 0x09, 0x25, 0x57, 0x2e, 0x76, 0x57, 0x88, 0x71, 0x42, 0x56, 0x5c,
	0x2c, 0x6e, 0x45, 0xf9, 0xb9, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0xdc, 0x46, 0x52, 0x7a, 0x10, 0x33,
	0xf5, 0x60, 0x66, 0xea, 0x85, 0xc0, 0xcc, 0x74, 0xe2, 0xda, 0xf5, 0xf2, 0x00, 0x33, 0xeb, 0x26,
	0x46, 0x26, 0x21, 0x86, 0x20, 0xb0, 0x1e, 0x27, 0x96, 0x28, 0xa6, 0x32, 0x83, 0x24, 0x36, 0xb0,
	0x5a, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x56, 0xbb, 0xa8, 0xb8, 0xa2, 0x00, 0x00, 0x00,
}
