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

type Transaction struct {
	Sender               string      `protobuf:"bytes,1,opt,name=Sender,proto3" json:"Sender,omitempty"`
	Recipient            string      `protobuf:"bytes,2,opt,name=Recipient,proto3" json:"Recipient,omitempty"`
	Amount               float64     `protobuf:"fixed64,3,opt,name=Amount,proto3" json:"Amount,omitempty"`
	Fee                  float64     `protobuf:"fixed64,6,opt,name=Fee,proto3" json:"Fee,omitempty"`
	Meta                 []byte      `protobuf:"bytes,5,opt,name=Meta,proto3" json:"Meta,omitempty"`
	Signatures           *Signatures `protobuf:"bytes,4,opt,name=Signatures,proto3" json:"Signatures,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Transaction) Reset()         { *m = Transaction{} }
func (m *Transaction) String() string { return proto.CompactTextString(m) }
func (*Transaction) ProtoMessage()    {}
func (*Transaction) Descriptor() ([]byte, []int) {
	return fileDescriptor_transaction_6896bcc40d697704, []int{0}
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

func (m *Transaction) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
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

type TransactionResponse struct {
	Payload              []byte    `protobuf:"bytes,1,opt,name=Payload,proto3" json:"Payload,omitempty"`
	Response             *Response `protobuf:"bytes,2,opt,name=Response,proto3" json:"Response,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *TransactionResponse) Reset()         { *m = TransactionResponse{} }
func (m *TransactionResponse) String() string { return proto.CompactTextString(m) }
func (*TransactionResponse) ProtoMessage()    {}
func (*TransactionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_transaction_6896bcc40d697704, []int{1}
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

func init() {
	proto.RegisterType((*Transaction)(nil), "transaction.Transaction")
	proto.RegisterType((*TransactionResponse)(nil), "transaction.TransactionResponse")
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
	err := c.cc.Invoke(ctx, "/transaction.TransactionService/Create", in, out, opts...)
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
		FullMethod: "/transaction.TransactionService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServiceServer).Create(ctx, req.(*Transaction))
	}
	return interceptor(ctx, in, info, handler)
}

var _TransactionService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "transaction.TransactionService",
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
	proto.RegisterFile("definitions/transaction.proto", fileDescriptor_transaction_6896bcc40d697704)
}

var fileDescriptor_transaction_6896bcc40d697704 = []byte{
	// 284 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0x51, 0x4b, 0xc3, 0x30,
	0x14, 0x85, 0xc9, 0x36, 0xab, 0xbb, 0x1d, 0x28, 0x57, 0x90, 0x30, 0x14, 0xca, 0x9e, 0xfa, 0x94,
	0xe1, 0xfc, 0x05, 0x53, 0x11, 0x7c, 0x18, 0x48, 0xe6, 0xd3, 0x5e, 0x24, 0xb6, 0x57, 0x09, 0x68,
	0x52, 0xd2, 0x6c, 0xb0, 0xbf, 0xe5, 0x2f, 0x94, 0x66, 0xed, 0x1a, 0x1f, 0x7c, 0xbb, 0xe7, 0x9c,
	0xaf, 0xb7, 0x27, 0x09, 0xdc, 0x94, 0xf4, 0xa1, 0x8d, 0xf6, 0xda, 0x9a, 0x7a, 0xee, 0x9d, 0x32,
	0xb5, 0x2a, 0x1a, 0x21, 0x2a, 0x67, 0xbd, 0xc5, 0x34, 0xb2, 0xa6, 0xd3, 0x98, 0xd5, 0x25, 0x19,
	0xaf, 0xfd, 0xfe, 0x00, 0xfe, 0xcd, 0x1c, 0xd5, 0x95, 0x35, 0x35, 0x1d, 0xb2, 0xd9, 0x0f, 0x83,
	0xf4, 0xb5, 0xdf, 0x83, 0x57, 0x90, 0xac, 0xc9, 0x94, 0xe4, 0x38, 0xcb, 0x58, 0x3e, 0x96, 0xad,
	0xc2, 0x6b, 0x18, 0x4b, 0x2a, 0x74, 0xa5, 0xc9, 0x78, 0x3e, 0x08, 0x51, 0x6f, 0x34, 0x5f, 0x2d,
	0xbf, 0xed, 0xd6, 0x78, 0x3e, 0xcc, 0x58, 0xce, 0x64, 0xab, 0xf0, 0x02, 0x86, 0x4f, 0x44, 0x3c,
	0x09, 0x66, 0x33, 0x22, 0xc2, 0x68, 0x45, 0x5e, 0xf1, 0x93, 0x8c, 0xe5, 0x13, 0x19, 0x66, 0x9c,
	0x03, 0xac, 0xf5, 0xa7, 0x51, 0x7e, 0xeb, 0xa8, 0xe6, 0xa3, 0x8c, 0xe5, 0xe9, 0xe2, 0x5c, 0x3c,
	0x2f, 0x57, 0xa2, 0xb7, 0x65, 0x84, 0xcc, 0xde, 0xe0, 0x32, 0xea, 0x2c, 0xdb, 0x13, 0x21, 0x87,
	0xd3, 0x17, 0xb5, 0xff, 0xb2, 0xaa, 0x0c, 0xe5, 0x27, 0xb2, 0x93, 0x28, 0xe0, 0xac, 0xa3, 0x42,
	0xf9, 0x74, 0x81, 0xa2, 0x33, 0x8e, 0x83, 0x3c, 0x32, 0x8b, 0x0d, 0x60, 0xf4, 0x83, 0x35, 0xb9,
	0x9d, 0x2e, 0x08, 0x1f, 0x21, 0x79, 0x70, 0xa4, 0x3c, 0x21, 0x17, 0xf1, 0x73, 0x44, 0xe8, 0x34,
	0xfb, 0x2f, 0xe9, 0x76, 0xdf, 0x8f, 0x36, 0x83, 0xdd, 0xed, 0x7b, 0x12, 0xae, 0xff, 0xee, 0x37,
	0x00, 0x00, 0xff, 0xff, 0xa8, 0xaf, 0x78, 0x47, 0xe4, 0x01, 0x00, 0x00,
}