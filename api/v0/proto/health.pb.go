// Code generated by protoc-gen-go. DO NOT EDIT.
// source: health.proto

package v0

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

type HealthCheckResponse_ServingStatus int32

const (
	HealthCheckResponse_UNKNOWN     HealthCheckResponse_ServingStatus = 0
	HealthCheckResponse_SERVING     HealthCheckResponse_ServingStatus = 1
	HealthCheckResponse_NOT_SERVING HealthCheckResponse_ServingStatus = 2
)

var HealthCheckResponse_ServingStatus_name = map[int32]string{
	0: "UNKNOWN",
	1: "SERVING",
	2: "NOT_SERVING",
}
var HealthCheckResponse_ServingStatus_value = map[string]int32{
	"UNKNOWN":     0,
	"SERVING":     1,
	"NOT_SERVING": 2,
}

func (x HealthCheckResponse_ServingStatus) String() string {
	return proto.EnumName(HealthCheckResponse_ServingStatus_name, int32(x))
}
func (HealthCheckResponse_ServingStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_health_ec8b97e0b4e40950, []int{1, 0}
}

type HealthCheckRequest struct {
	Service              string   `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HealthCheckRequest) Reset()         { *m = HealthCheckRequest{} }
func (m *HealthCheckRequest) String() string { return proto.CompactTextString(m) }
func (*HealthCheckRequest) ProtoMessage()    {}
func (*HealthCheckRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_health_ec8b97e0b4e40950, []int{0}
}
func (m *HealthCheckRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthCheckRequest.Unmarshal(m, b)
}
func (m *HealthCheckRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthCheckRequest.Marshal(b, m, deterministic)
}
func (dst *HealthCheckRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthCheckRequest.Merge(dst, src)
}
func (m *HealthCheckRequest) XXX_Size() int {
	return xxx_messageInfo_HealthCheckRequest.Size(m)
}
func (m *HealthCheckRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthCheckRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HealthCheckRequest proto.InternalMessageInfo

func (m *HealthCheckRequest) GetService() string {
	if m != nil {
		return m.Service
	}
	return ""
}

type HealthCheckResponse struct {
	Status               HealthCheckResponse_ServingStatus `protobuf:"varint,1,opt,name=status,proto3,enum=v0.HealthCheckResponse_ServingStatus" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                          `json:"-"`
	XXX_unrecognized     []byte                            `json:"-"`
	XXX_sizecache        int32                             `json:"-"`
}

func (m *HealthCheckResponse) Reset()         { *m = HealthCheckResponse{} }
func (m *HealthCheckResponse) String() string { return proto.CompactTextString(m) }
func (*HealthCheckResponse) ProtoMessage()    {}
func (*HealthCheckResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_health_ec8b97e0b4e40950, []int{1}
}
func (m *HealthCheckResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthCheckResponse.Unmarshal(m, b)
}
func (m *HealthCheckResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthCheckResponse.Marshal(b, m, deterministic)
}
func (dst *HealthCheckResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthCheckResponse.Merge(dst, src)
}
func (m *HealthCheckResponse) XXX_Size() int {
	return xxx_messageInfo_HealthCheckResponse.Size(m)
}
func (m *HealthCheckResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthCheckResponse.DiscardUnknown(m)
}

var xxx_messageInfo_HealthCheckResponse proto.InternalMessageInfo

func (m *HealthCheckResponse) GetStatus() HealthCheckResponse_ServingStatus {
	if m != nil {
		return m.Status
	}
	return HealthCheckResponse_UNKNOWN
}

func init() {
	proto.RegisterType((*HealthCheckRequest)(nil), "v0.HealthCheckRequest")
	proto.RegisterType((*HealthCheckResponse)(nil), "v0.HealthCheckResponse")
	proto.RegisterEnum("v0.HealthCheckResponse_ServingStatus", HealthCheckResponse_ServingStatus_name, HealthCheckResponse_ServingStatus_value)
}

func init() { proto.RegisterFile("health.proto", fileDescriptor_health_ec8b97e0b4e40950) }

var fileDescriptor_health_ec8b97e0b4e40950 = []byte{
	// 179 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xc9, 0x48, 0x4d, 0xcc,
	0x29, 0xc9, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x33, 0x50, 0xd2, 0xe3, 0x12,
	0xf2, 0x00, 0x8b, 0x39, 0x67, 0xa4, 0x26, 0x67, 0x07, 0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97, 0x08,
	0x49, 0x70, 0xb1, 0x17, 0xa7, 0x16, 0x95, 0x65, 0x26, 0xa7, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0x70,
	0x06, 0xc1, 0xb8, 0x4a, 0x13, 0x18, 0xb9, 0x84, 0x51, 0x34, 0x14, 0x17, 0xe4, 0xe7, 0x15, 0xa7,
	0x0a, 0xd9, 0x72, 0xb1, 0x15, 0x97, 0x24, 0x96, 0x94, 0x16, 0x83, 0x35, 0xf0, 0x19, 0xa9, 0xea,
	0x95, 0x19, 0xe8, 0x61, 0x51, 0xa8, 0x17, 0x0c, 0x32, 0x28, 0x2f, 0x3d, 0x18, 0xac, 0x38, 0x08,
	0xaa, 0x49, 0xc9, 0x8a, 0x8b, 0x17, 0x45, 0x42, 0x88, 0x9b, 0x8b, 0x3d, 0xd4, 0xcf, 0xdb, 0xcf,
	0x3f, 0xdc, 0x4f, 0x80, 0x01, 0xc4, 0x09, 0x76, 0x0d, 0x0a, 0xf3, 0xf4, 0x73, 0x17, 0x60, 0x14,
	0xe2, 0xe7, 0xe2, 0xf6, 0xf3, 0x0f, 0x89, 0x87, 0x09, 0x30, 0x39, 0xb1, 0x44, 0x31, 0x95, 0x19,
	0x24, 0xb1, 0x81, 0xfd, 0x64, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x28, 0xd1, 0xde, 0x45, 0xe3,
	0x00, 0x00, 0x00,
}