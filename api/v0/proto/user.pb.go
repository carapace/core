// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

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

type User struct {
	Name                 string   `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	Email                string   `protobuf:"bytes,2,opt,name=Email,proto3" json:"Email,omitempty"`
	PrimaryPublicKey     []byte   `protobuf:"bytes,3,opt,name=PrimaryPublicKey,proto3" json:"PrimaryPublicKey,omitempty"`
	RecoveryPublicKey    []byte   `protobuf:"bytes,4,opt,name=RecoveryPublicKey,proto3" json:"RecoveryPublicKey,omitempty"`
	AuthLevel            int32    `protobuf:"varint,5,opt,name=AuthLevel,proto3" json:"AuthLevel,omitempty"`
	SuperUser            bool     `protobuf:"varint,6,opt,name=SuperUser,proto3" json:"SuperUser,omitempty"`
	Weight               int32    `protobuf:"varint,7,opt,name=Weight,proto3" json:"Weight,omitempty"`
	Set                  string   `protobuf:"bytes,8,opt,name=Set,proto3" json:"Set,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_7b4b4056daf2b4d8, []int{0}
}
func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (dst *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(dst, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetPrimaryPublicKey() []byte {
	if m != nil {
		return m.PrimaryPublicKey
	}
	return nil
}

func (m *User) GetRecoveryPublicKey() []byte {
	if m != nil {
		return m.RecoveryPublicKey
	}
	return nil
}

func (m *User) GetAuthLevel() int32 {
	if m != nil {
		return m.AuthLevel
	}
	return 0
}

func (m *User) GetSuperUser() bool {
	if m != nil {
		return m.SuperUser
	}
	return false
}

func (m *User) GetWeight() int32 {
	if m != nil {
		return m.Weight
	}
	return 0
}

func (m *User) GetSet() string {
	if m != nil {
		return m.Set
	}
	return ""
}

type UserSet struct {
	Set                  string   `protobuf:"bytes,1,opt,name=Set,proto3" json:"Set,omitempty"`
	Users                []*User  `protobuf:"bytes,2,rep,name=Users,proto3" json:"Users,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserSet) Reset()         { *m = UserSet{} }
func (m *UserSet) String() string { return proto.CompactTextString(m) }
func (*UserSet) ProtoMessage()    {}
func (*UserSet) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_7b4b4056daf2b4d8, []int{1}
}
func (m *UserSet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserSet.Unmarshal(m, b)
}
func (m *UserSet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserSet.Marshal(b, m, deterministic)
}
func (dst *UserSet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserSet.Merge(dst, src)
}
func (m *UserSet) XXX_Size() int {
	return xxx_messageInfo_UserSet.Size(m)
}
func (m *UserSet) XXX_DiscardUnknown() {
	xxx_messageInfo_UserSet.DiscardUnknown(m)
}

var xxx_messageInfo_UserSet proto.InternalMessageInfo

func (m *UserSet) GetSet() string {
	if m != nil {
		return m.Set
	}
	return ""
}

func (m *UserSet) GetUsers() []*User {
	if m != nil {
		return m.Users
	}
	return nil
}

func init() {
	proto.RegisterType((*User)(nil), "v0.User")
	proto.RegisterType((*UserSet)(nil), "v0.UserSet")
}

func init() { proto.RegisterFile("user.proto", fileDescriptor_user_7b4b4056daf2b4d8) }

var fileDescriptor_user_7b4b4056daf2b4d8 = []byte{
	// 239 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0x4d, 0x4b, 0xc3, 0x40,
	0x10, 0x86, 0xd9, 0x7c, 0x35, 0x1d, 0x3d, 0xd4, 0x41, 0x64, 0x0e, 0x22, 0x4b, 0x4f, 0x8b, 0x48,
	0x08, 0x7a, 0xf4, 0xa4, 0xe0, 0x49, 0x91, 0xb2, 0x41, 0x04, 0x6f, 0x69, 0x19, 0x6c, 0x20, 0x21,
	0x65, 0x93, 0x5d, 0xe8, 0xbf, 0xf6, 0x27, 0xc8, 0x6e, 0xa5, 0x15, 0x72, 0x9b, 0x79, 0x9e, 0xf7,
	0x85, 0xdd, 0x01, 0xb0, 0x03, 0x9b, 0x62, 0x67, 0xfa, 0xb1, 0xc7, 0xc8, 0x95, 0xcb, 0x1f, 0x01,
	0xc9, 0xc7, 0xc0, 0x06, 0x11, 0x92, 0xf7, 0xba, 0x63, 0x12, 0x52, 0xa8, 0xb9, 0x0e, 0x33, 0x5e,
	0x42, 0xfa, 0xd2, 0xd5, 0x4d, 0x4b, 0x51, 0x80, 0x87, 0x05, 0x6f, 0x61, 0xb1, 0x32, 0x4d, 0x57,
	0x9b, 0xfd, 0xca, 0xae, 0xdb, 0x66, 0xf3, 0xca, 0x7b, 0x8a, 0xa5, 0x50, 0xe7, 0x7a, 0xc2, 0xf1,
	0x0e, 0x2e, 0x34, 0x6f, 0x7a, 0xc7, 0xff, 0xc3, 0x49, 0x08, 0x4f, 0x05, 0x5e, 0xc3, 0xfc, 0xc9,
	0x8e, 0xdb, 0x37, 0x76, 0xdc, 0x52, 0x2a, 0x85, 0x4a, 0xf5, 0x09, 0x78, 0x5b, 0xd9, 0x1d, 0x1b,
	0xff, 0x5c, 0xca, 0xa4, 0x50, 0xb9, 0x3e, 0x01, 0xbc, 0x82, 0xec, 0x93, 0x9b, 0xef, 0xed, 0x48,
	0xb3, 0x50, 0xfc, 0xdb, 0x70, 0x01, 0x71, 0xc5, 0x23, 0xe5, 0xe1, 0x07, 0x7e, 0x5c, 0x3e, 0xc2,
	0xcc, 0x37, 0x2a, 0x3e, 0x4a, 0x71, 0x94, 0x78, 0x03, 0xa9, 0x97, 0x03, 0x45, 0x32, 0x56, 0x67,
	0xf7, 0x79, 0xe1, 0xca, 0xc2, 0x03, 0x7d, 0xc0, 0xcf, 0xc9, 0x57, 0xe4, 0xca, 0x75, 0x16, 0x0e,
	0xf8, 0xf0, 0x1b, 0x00, 0x00, 0xff, 0xff, 0xd2, 0xca, 0xe6, 0x6f, 0x4e, 0x01, 0x00, 0x00,
}