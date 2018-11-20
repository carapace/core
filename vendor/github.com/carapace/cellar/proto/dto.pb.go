// Code generated by protoc-gen-go.
// source: dto.proto
// DO NOT EDIT!

package proto

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

type ChunkDto struct {
	UncompressedByteSize int64  `protobuf:"varint,1,opt,name=uncompressedByteSize" json:"uncompressedByteSize,omitempty"`
	CompressedDiskSize   int64  `protobuf:"varint,2,opt,name=compressedDiskSize" json:"compressedDiskSize,omitempty"`
	Records              int64  `protobuf:"varint,3,opt,name=records" json:"records,omitempty"`
	FileName             string `protobuf:"bytes,4,opt,name=fileName" json:"fileName,omitempty"`
	StartPos             int64  `protobuf:"varint,5,opt,name=startPos" json:"startPos,omitempty"`
}

func (m *ChunkDto) Reset()                    { *m = ChunkDto{} }
func (m *ChunkDto) String() string            { return proto.CompactTextString(m) }
func (*ChunkDto) ProtoMessage()               {}
func (*ChunkDto) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type BufferDto struct {
	StartPos int64  `protobuf:"varint,1,opt,name=startPos" json:"startPos,omitempty"`
	MaxBytes int64  `protobuf:"varint,2,opt,name=maxBytes" json:"maxBytes,omitempty"`
	Records  int64  `protobuf:"varint,3,opt,name=records" json:"records,omitempty"`
	Pos      int64  `protobuf:"varint,4,opt,name=pos" json:"pos,omitempty"`
	FileName string `protobuf:"bytes,5,opt,name=fileName" json:"fileName,omitempty"`
}

func (m *BufferDto) Reset()                    { *m = BufferDto{} }
func (m *BufferDto) String() string            { return proto.CompactTextString(m) }
func (*BufferDto) ProtoMessage()               {}
func (*BufferDto) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type MetaDto struct {
	MaxKeySize int64 `protobuf:"varint,1,opt,name=maxKeySize" json:"maxKeySize,omitempty"`
	MaxValSize int64 `protobuf:"varint,2,opt,name=maxValSize" json:"maxValSize,omitempty"`
}

func (m *MetaDto) Reset()                    { *m = MetaDto{} }
func (m *MetaDto) String() string            { return proto.CompactTextString(m) }
func (*MetaDto) ProtoMessage()               {}
func (*MetaDto) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func init() {
	proto.RegisterType((*ChunkDto)(nil), "cellar.ChunkDto")
	proto.RegisterType((*BufferDto)(nil), "cellar.BufferDto")
	proto.RegisterType((*MetaDto)(nil), "cellar.MetaDto")
}

func init() { proto.RegisterFile("dto.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 246 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x7c, 0x91, 0xcd, 0x4a, 0xc3, 0x40,
	0x10, 0xc7, 0x59, 0x63, 0x3f, 0x32, 0x27, 0x59, 0x3c, 0x2c, 0x1e, 0xa4, 0xe4, 0xd4, 0x53, 0x0f,
	0xfa, 0x06, 0xb5, 0x17, 0x11, 0x45, 0x22, 0x78, 0x5f, 0x93, 0x09, 0x86, 0xee, 0x76, 0xc2, 0xce,
	0x06, 0x5a, 0x5f, 0xc1, 0x97, 0xf2, 0xd1, 0x64, 0x97, 0x1a, 0x37, 0x50, 0x7a, 0xcb, 0xff, 0x0b,
	0x7e, 0x99, 0x85, 0xbc, 0xf6, 0xb4, 0xea, 0x1c, 0x79, 0x92, 0xd3, 0x0a, 0x8d, 0xd1, 0xae, 0xf8,
	0x11, 0x30, 0x7f, 0xf8, 0xec, 0x77, 0xdb, 0x8d, 0x27, 0x79, 0x07, 0xd7, 0xfd, 0xae, 0x22, 0xdb,
	0x39, 0x64, 0xc6, 0x7a, 0x7d, 0xf0, 0xf8, 0xd6, 0x7e, 0xa1, 0x12, 0x0b, 0xb1, 0xcc, 0xca, 0x93,
	0x99, 0x5c, 0x81, 0xfc, 0x77, 0x37, 0x2d, 0x6f, 0xe3, 0xe2, 0x22, 0x2e, 0x4e, 0x24, 0x52, 0xc1,
	0xcc, 0x61, 0x45, 0xae, 0x66, 0x95, 0xc5, 0xd2, 0x9f, 0x94, 0x37, 0x30, 0x6f, 0x5a, 0x83, 0x2f,
	0xda, 0xa2, 0xba, 0x5c, 0x88, 0x65, 0x5e, 0x0e, 0x3a, 0x64, 0xec, 0xb5, 0xf3, 0xaf, 0xc4, 0x6a,
	0x12, 0x67, 0x83, 0x2e, 0xbe, 0x05, 0xe4, 0xeb, 0xbe, 0x69, 0xd0, 0x85, 0x7f, 0x48, 0x9b, 0x62,
	0xdc, 0x0c, 0x99, 0xd5, 0xfb, 0x80, 0xce, 0x47, 0xc2, 0x41, 0x9f, 0xe1, 0xba, 0x82, 0xac, 0x23,
	0x8e, 0x48, 0x59, 0x19, 0x3e, 0x47, 0xa4, 0x93, 0x31, 0x69, 0xf1, 0x08, 0xb3, 0x67, 0xf4, 0x3a,
	0xa0, 0xdc, 0x02, 0x58, 0xbd, 0x7f, 0xc2, 0x43, 0x72, 0xc4, 0xc4, 0x39, 0xe6, 0xef, 0xda, 0x24,
	0x27, 0x4b, 0x9c, 0x8f, 0x69, 0x7c, 0xaa, 0xfb, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x8a, 0x58,
	0x12, 0xdc, 0xb7, 0x01, 0x00, 0x00,
}