// Code generated by protoc-gen-go. DO NOT EDIT.
// source: definitions/transaction.proto

package v1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type TransactionType int32

const (
	TransactionType_TransactionTypeUnknown TransactionType = 0
	TransactionType_Transfer               TransactionType = 1
	TransactionType_SCCreation             TransactionType = 2
	TransactionType_SCInvocation           TransactionType = 3
	TransactionType_Vote                   TransactionType = 4
)

var TransactionType_name = map[int32]string{
	0: "TransactionTypeUnknown",
	1: "Transfer",
	2: "SCCreation",
	3: "SCInvocation",
	4: "Vote",
}
var TransactionType_value = map[string]int32{
	"TransactionTypeUnknown": 0,
	"Transfer":               1,
	"SCCreation":             2,
	"SCInvocation":           3,
	"Vote":                   4,
}

func (x TransactionType) String() string {
	return proto.EnumName(TransactionType_name, int32(x))
}
func (TransactionType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_transaction_593d5ce024b495f1, []int{0}
}

type Transaction struct {
	Type                 TransactionType `protobuf:"varint,1,opt,name=Type,proto3,enum=v1.transaction.TransactionType" json:"Type,omitempty"`
	Recipient            string          `protobuf:"bytes,2,opt,name=Recipient,proto3" json:"Recipient,omitempty"`
	Amount               float64         `protobuf:"fixed64,3,opt,name=Amount,proto3" json:"Amount,omitempty"`
	Fee                  float64         `protobuf:"fixed64,4,opt,name=Fee,proto3" json:"Fee,omitempty"`
	Meta                 []byte          `protobuf:"bytes,5,opt,name=Meta,proto3" json:"Meta,omitempty"`
	Signatures           *Signatures     `protobuf:"bytes,6,opt,name=Signatures,proto3" json:"Signatures,omitempty"`
	Namespace            *Namespace      `protobuf:"bytes,7,opt,name=Namespace,proto3" json:"Namespace,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Transaction) Reset()         { *m = Transaction{} }
func (m *Transaction) String() string { return proto.CompactTextString(m) }
func (*Transaction) ProtoMessage()    {}
func (*Transaction) Descriptor() ([]byte, []int) {
	return fileDescriptor_transaction_593d5ce024b495f1, []int{0}
}
func (m *Transaction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Transaction.Unmarshal(m, b)
}
func (m *Transaction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Transaction.Marshal(b, m, deterministic)
}
func (dst *Transaction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Transaction.Merge(dst, src)
}
func (m *Transaction) XXX_Size() int {
	return xxx_messageInfo_Transaction.Size(m)
}
func (m *Transaction) XXX_DiscardUnknown() {
	xxx_messageInfo_Transaction.DiscardUnknown(m)
}

var xxx_messageInfo_Transaction proto.InternalMessageInfo

func (m *Transaction) GetType() TransactionType {
	if m != nil {
		return m.Type
	}
	return TransactionType_TransactionTypeUnknown
}

func (m *Transaction) GetRecipient() string {
	if m != nil {
		return m.Recipient
	}
	return ""
}

func (m *Transaction) GetAmount() float64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *Transaction) GetFee() float64 {
	if m != nil {
		return m.Fee
	}
	return 0
}

func (m *Transaction) GetMeta() []byte {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *Transaction) GetSignatures() *Signatures {
	if m != nil {
		return m.Signatures
	}
	return nil
}

func (m *Transaction) GetNamespace() *Namespace {
	if m != nil {
		return m.Namespace
	}
	return nil
}

type TransactionResponse struct {
	Payload              []byte     `protobuf:"bytes,1,opt,name=Payload,proto3" json:"Payload,omitempty"`
	Response             *Response  `protobuf:"bytes,2,opt,name=Response,proto3" json:"Response,omitempty"`
	Namespace            *Namespace `protobuf:"bytes,3,opt,name=Namespace,proto3" json:"Namespace,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *TransactionResponse) Reset()         { *m = TransactionResponse{} }
func (m *TransactionResponse) String() string { return proto.CompactTextString(m) }
func (*TransactionResponse) ProtoMessage()    {}
func (*TransactionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_transaction_593d5ce024b495f1, []int{1}
}
func (m *TransactionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransactionResponse.Unmarshal(m, b)
}
func (m *TransactionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransactionResponse.Marshal(b, m, deterministic)
}
func (dst *TransactionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransactionResponse.Merge(dst, src)
}
func (m *TransactionResponse) XXX_Size() int {
	return xxx_messageInfo_TransactionResponse.Size(m)
}
func (m *TransactionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TransactionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TransactionResponse proto.InternalMessageInfo

func (m *TransactionResponse) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *TransactionResponse) GetResponse() *Response {
	if m != nil {
		return m.Response
	}
	return nil
}

func (m *TransactionResponse) GetNamespace() *Namespace {
	if m != nil {
		return m.Namespace
	}
	return nil
}

func init() {
	proto.RegisterType((*Transaction)(nil), "v1.transaction.Transaction")
	proto.RegisterType((*TransactionResponse)(nil), "v1.transaction.TransactionResponse")
	proto.RegisterEnum("v1.transaction.TransactionType", TransactionType_name, TransactionType_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TransactionServiceClient is the client API for TransactionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TransactionServiceClient interface {
	Create(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*TransactionResponse, error)
}

type transactionServiceClient struct {
	cc *grpc.ClientConn
}

func NewTransactionServiceClient(cc *grpc.ClientConn) TransactionServiceClient {
	return &transactionServiceClient{cc}
}

func (c *transactionServiceClient) Create(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*TransactionResponse, error) {
	out := new(TransactionResponse)
	err := c.cc.Invoke(ctx, "/v1.transaction.TransactionService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TransactionServiceServer is the server API for TransactionService service.
type TransactionServiceServer interface {
	Create(context.Context, *Transaction) (*TransactionResponse, error)
}

func RegisterTransactionServiceServer(s *grpc.Server, srv TransactionServiceServer) {
	s.RegisterService(&_TransactionService_serviceDesc, srv)
}

func _TransactionService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Transaction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.transaction.TransactionService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServiceServer).Create(ctx, req.(*Transaction))
	}
	return interceptor(ctx, in, info, handler)
}

var _TransactionService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "v1.transaction.TransactionService",
	HandlerType: (*TransactionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _TransactionService_Create_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "definitions/transaction.proto",
}

func init() {
	proto.RegisterFile("definitions/transaction.proto", fileDescriptor_transaction_593d5ce024b495f1)
}

var fileDescriptor_transaction_593d5ce024b495f1 = []byte{
	// 400 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x51, 0x6f, 0xd3, 0x30,
	0x10, 0x80, 0x71, 0x1b, 0xb2, 0xf6, 0x1a, 0x95, 0xe8, 0x10, 0xc3, 0xca, 0x40, 0x44, 0xe3, 0x25,
	0xe2, 0x21, 0xa8, 0x99, 0xf8, 0x01, 0xa3, 0x12, 0xd2, 0x90, 0x86, 0x90, 0x33, 0x78, 0xe0, 0x09,
	0x93, 0xde, 0x26, 0x0b, 0x66, 0x47, 0x89, 0x17, 0xd4, 0x9f, 0xc0, 0x4f, 0xe0, 0xdf, 0xa2, 0x78,
	0x49, 0x93, 0xee, 0xa1, 0xe2, 0xcd, 0x77, 0xdf, 0x77, 0xf1, 0xdd, 0xc5, 0xf0, 0x72, 0x43, 0xd7,
	0x4a, 0x2b, 0xab, 0x8c, 0xae, 0xdf, 0xda, 0x4a, 0xea, 0x5a, 0x16, 0x6d, 0x90, 0x96, 0x95, 0xb1,
	0x06, 0x97, 0xcd, 0x2a, 0x1d, 0x65, 0xa3, 0x68, 0xac, 0xab, 0x0d, 0x69, 0xab, 0xec, 0xf6, 0xde,
	0xdd, 0x67, 0x15, 0xd5, 0xa5, 0xd1, 0x35, 0x75, 0xec, 0x64, 0xcc, 0xb4, 0xbc, 0xa5, 0xba, 0x94,
	0x45, 0x07, 0x4f, 0xff, 0x4c, 0x60, 0x71, 0x35, 0x5c, 0x82, 0x67, 0xe0, 0x5d, 0x6d, 0x4b, 0xe2,
	0x2c, 0x66, 0xc9, 0x32, 0x7b, 0x95, 0xee, 0xf7, 0x90, 0x8e, 0xd4, 0x56, 0x13, 0x4e, 0xc6, 0x17,
	0x30, 0x17, 0x54, 0xa8, 0x52, 0x91, 0xb6, 0x7c, 0x12, 0xb3, 0x64, 0x2e, 0x86, 0x04, 0x1e, 0x83,
	0x7f, 0x7e, 0x6b, 0xee, 0xb4, 0xe5, 0xd3, 0x98, 0x25, 0x4c, 0x74, 0x11, 0x86, 0x30, 0xfd, 0x40,
	0xc4, 0x3d, 0x97, 0x6c, 0x8f, 0x88, 0xe0, 0x5d, 0x92, 0x95, 0xfc, 0x71, 0xcc, 0x92, 0x40, 0xb8,
	0x33, 0x66, 0x00, 0xb9, 0xba, 0xd1, 0xd2, 0xde, 0x55, 0x54, 0x73, 0x3f, 0x66, 0xc9, 0x22, 0xc3,
	0xb6, 0xad, 0x8b, 0xf3, 0xcb, 0x74, 0x20, 0x62, 0x64, 0xe1, 0x3b, 0x98, 0x7f, 0xea, 0xe7, 0xe4,
	0x47, 0xae, 0xe4, 0x79, 0x5b, 0x32, 0x0c, 0xbf, 0xc3, 0x62, 0x30, 0x4f, 0xff, 0x32, 0x78, 0x3a,
	0x1a, 0x50, 0x74, 0x6b, 0x44, 0x0e, 0x47, 0x9f, 0xe5, 0xf6, 0x97, 0x91, 0x1b, 0xb7, 0x96, 0x40,
	0xf4, 0x21, 0xae, 0x60, 0xd6, 0x5b, 0x6e, 0xee, 0x45, 0xf6, 0xac, 0xbd, 0x67, 0xf7, 0x03, 0x7a,
	0x28, 0x76, 0xda, 0x7e, 0x6f, 0xd3, 0xff, 0xed, 0xed, 0xcd, 0x0d, 0x3c, 0x79, 0xb0, 0x7b, 0x8c,
	0xe0, 0xf8, 0x41, 0xea, 0x8b, 0xfe, 0xa9, 0xcd, 0x6f, 0x1d, 0x3e, 0xc2, 0x00, 0x66, 0x8e, 0x5d,
	0x53, 0x15, 0x32, 0x5c, 0x02, 0xe4, 0xeb, 0x75, 0x45, 0xb2, 0x15, 0xc3, 0x09, 0x86, 0x10, 0xe4,
	0xeb, 0x0b, 0xdd, 0x98, 0xe2, 0x3e, 0x33, 0xc5, 0x19, 0x78, 0x5f, 0x8d, 0xa5, 0xd0, 0xcb, 0xbe,
	0x03, 0x8e, 0xbe, 0x9a, 0x53, 0xd5, 0xa8, 0x82, 0xf0, 0x23, 0xf8, 0xae, 0x9e, 0xf0, 0xe4, 0xc0,
	0x93, 0x88, 0x5e, 0x1f, 0x80, 0xfd, 0x06, 0xde, 0x7b, 0xdf, 0x26, 0xcd, 0xea, 0x87, 0xef, 0xde,
	0xdf, 0xd9, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa0, 0x9b, 0x8d, 0xba, 0x05, 0x03, 0x00, 0x00,
}
