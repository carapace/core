// Code generated by protoc-gen-go. DO NOT EDIT.
// source: signatures.proto

package v0

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
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

//
// Witness contains the signatures by approving users. ECDSA is used to sign the entire message; excluding the signature
// field, which should be set to its null value.
//
// The signing protocol goes as follows:
//
// 1. A keypair is generated from a user provided seed.
// 2. The config's witness field is set to it's nill state (using the Reset() method of Witness)
// 3. The config is serialized to a byte array using the protobuf generated serialization.
// 4. ECDSA is used to sign the message.
// 5. The public key, signature.R and signature.S are appended to the witness field (possibly with previous signatures)
//
// A RecoveryPublicKey may be used by certain sets, mainly for account resetting purposes.
type Witness struct {
	Signatures           []*Signature `protobuf:"bytes,1,rep,name=Signatures,proto3" json:"Signatures,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Witness) Reset()         { *m = Witness{} }
func (m *Witness) String() string { return proto.CompactTextString(m) }
func (*Witness) ProtoMessage()    {}
func (*Witness) Descriptor() ([]byte, []int) {
	return fileDescriptor_signatures_4ab464274f9c3597, []int{0}
}
func (m *Witness) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Witness.Unmarshal(m, b)
}
func (m *Witness) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Witness.Marshal(b, m, deterministic)
}
func (dst *Witness) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Witness.Merge(dst, src)
}
func (m *Witness) XXX_Size() int {
	return xxx_messageInfo_Witness.Size(m)
}
func (m *Witness) XXX_DiscardUnknown() {
	xxx_messageInfo_Witness.DiscardUnknown(m)
}

var xxx_messageInfo_Witness proto.InternalMessageInfo

func (m *Witness) GetSignatures() []*Signature {
	if m != nil {
		return m.Signatures
	}
	return nil
}

//
// A signature consists of:
// 1. either the PrimaryPublicKey (used for regular operations)
// or RecoveryPublicKey (used for account recovery)
// 2. The R and S of the ECDSA signature as byte arrays.
type Signature struct {
	// Types that are valid to be assigned to Key:
	//	*Signature_PrimaryPublicKey
	//	*Signature_RecoveryPublicKey
	Key                  isSignature_Key `protobuf_oneof:"Key"`
	R                    []byte          `protobuf:"bytes,3,opt,name=R,proto3" json:"R,omitempty"`
	S                    []byte          `protobuf:"bytes,4,opt,name=S,proto3" json:"S,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Signature) Reset()         { *m = Signature{} }
func (m *Signature) String() string { return proto.CompactTextString(m) }
func (*Signature) ProtoMessage()    {}
func (*Signature) Descriptor() ([]byte, []int) {
	return fileDescriptor_signatures_4ab464274f9c3597, []int{1}
}
func (m *Signature) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Signature.Unmarshal(m, b)
}
func (m *Signature) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Signature.Marshal(b, m, deterministic)
}
func (dst *Signature) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Signature.Merge(dst, src)
}
func (m *Signature) XXX_Size() int {
	return xxx_messageInfo_Signature.Size(m)
}
func (m *Signature) XXX_DiscardUnknown() {
	xxx_messageInfo_Signature.DiscardUnknown(m)
}

var xxx_messageInfo_Signature proto.InternalMessageInfo

type isSignature_Key interface {
	isSignature_Key()
}

type Signature_PrimaryPublicKey struct {
	PrimaryPublicKey []byte `protobuf:"bytes,1,opt,name=PrimaryPublicKey,proto3,oneof"`
}
type Signature_RecoveryPublicKey struct {
	RecoveryPublicKey []byte `protobuf:"bytes,2,opt,name=RecoveryPublicKey,proto3,oneof"`
}

func (*Signature_PrimaryPublicKey) isSignature_Key()  {}
func (*Signature_RecoveryPublicKey) isSignature_Key() {}

func (m *Signature) GetKey() isSignature_Key {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *Signature) GetPrimaryPublicKey() []byte {
	if x, ok := m.GetKey().(*Signature_PrimaryPublicKey); ok {
		return x.PrimaryPublicKey
	}
	return nil
}

func (m *Signature) GetRecoveryPublicKey() []byte {
	if x, ok := m.GetKey().(*Signature_RecoveryPublicKey); ok {
		return x.RecoveryPublicKey
	}
	return nil
}

func (m *Signature) GetR() []byte {
	if m != nil {
		return m.R
	}
	return nil
}

func (m *Signature) GetS() []byte {
	if m != nil {
		return m.S
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Signature) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Signature_OneofMarshaler, _Signature_OneofUnmarshaler, _Signature_OneofSizer, []interface{}{
		(*Signature_PrimaryPublicKey)(nil),
		(*Signature_RecoveryPublicKey)(nil),
	}
}

func _Signature_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Signature)
	// Key
	switch x := m.Key.(type) {
	case *Signature_PrimaryPublicKey:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		b.EncodeRawBytes(x.PrimaryPublicKey)
	case *Signature_RecoveryPublicKey:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		b.EncodeRawBytes(x.RecoveryPublicKey)
	case nil:
	default:
		return fmt.Errorf("Signature.Key has unexpected type %T", x)
	}
	return nil
}

func _Signature_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Signature)
	switch tag {
	case 1: // Key.PrimaryPublicKey
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeRawBytes(true)
		m.Key = &Signature_PrimaryPublicKey{x}
		return true, err
	case 2: // Key.RecoveryPublicKey
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeRawBytes(true)
		m.Key = &Signature_RecoveryPublicKey{x}
		return true, err
	default:
		return false, nil
	}
}

func _Signature_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Signature)
	// Key
	switch x := m.Key.(type) {
	case *Signature_PrimaryPublicKey:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.PrimaryPublicKey)))
		n += len(x.PrimaryPublicKey)
	case *Signature_RecoveryPublicKey:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.RecoveryPublicKey)))
		n += len(x.RecoveryPublicKey)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*Witness)(nil), "v0.Witness")
	proto.RegisterType((*Signature)(nil), "v0.Signature")
}

func init() { proto.RegisterFile("signatures.proto", fileDescriptor_signatures_4ab464274f9c3597) }

var fileDescriptor_signatures_4ab464274f9c3597 = []byte{
	// 216 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x28, 0xce, 0x4c, 0xcf,
	0x4b, 0x2c, 0x29, 0x2d, 0x4a, 0x2d, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x33,
	0x90, 0x12, 0x2f, 0x4b, 0xcc, 0xc9, 0x4c, 0x49, 0x2c, 0x49, 0xd5, 0x87, 0x31, 0x20, 0x92, 0x4a,
	0x2e, 0x5c, 0xec, 0xe1, 0x99, 0x25, 0x79, 0xa9, 0xc5, 0xc5, 0x42, 0x96, 0x5c, 0x5c, 0xc1, 0x70,
	0xbd, 0x12, 0x8c, 0x0a, 0xcc, 0x1a, 0xdc, 0x46, 0xbc, 0x7a, 0x65, 0x06, 0x7a, 0x70, 0x51, 0x27,
	0xae, 0x5d, 0x2f, 0x0f, 0x30, 0xb3, 0x4e, 0x62, 0x64, 0xe2, 0x60, 0x0c, 0x42, 0x52, 0xac, 0xb4,
	0x92, 0x91, 0x8b, 0x13, 0xce, 0x15, 0xd2, 0xe1, 0x12, 0x08, 0x28, 0xca, 0xcc, 0x4d, 0x2c, 0xaa,
	0x0c, 0x28, 0x4d, 0xca, 0xc9, 0x4c, 0xf6, 0x4e, 0xad, 0x94, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0xf1,
	0x60, 0x08, 0xc2, 0x90, 0x11, 0xd2, 0xe3, 0x12, 0x0c, 0x4a, 0x4d, 0xce, 0x2f, 0x4b, 0x45, 0x56,
	0xce, 0x04, 0x55, 0x8e, 0x29, 0x25, 0x24, 0xce, 0xc5, 0x18, 0x24, 0xc1, 0x0c, 0x92, 0x77, 0xe2,
	0x04, 0x39, 0x87, 0xa5, 0x8a, 0x49, 0x80, 0x31, 0x88, 0x31, 0x08, 0x24, 0x11, 0x2c, 0xc1, 0x82,
	0x21, 0x11, 0xec, 0xc4, 0xc3, 0xc5, 0x0c, 0xd2, 0xc8, 0xba, 0xe3, 0xe5, 0x01, 0x66, 0x46, 0x27,
	0x96, 0x28, 0xa6, 0x32, 0x83, 0x24, 0x36, 0xb0, 0xf7, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff,
	0x44, 0xea, 0x5e, 0x62, 0x2f, 0x01, 0x00, 0x00,
}