// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service.proto

package v0

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import empty "github.com/golang/protobuf/ptypes/empty"

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

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CoreServiceClient is the client API for CoreService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CoreServiceClient interface {
	ConfigService(ctx context.Context, in *Config, opts ...grpc.CallOption) (*Response, error)
	InfoService(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*RepeatedInfo, error)
	Check(ctx context.Context, in *HealthCheckRequest, opts ...grpc.CallOption) (*HealthCheckResponse, error)
}

type coreServiceClient struct {
	cc *grpc.ClientConn
}

func NewCoreServiceClient(cc *grpc.ClientConn) CoreServiceClient {
	return &coreServiceClient{cc}
}

func (c *coreServiceClient) ConfigService(ctx context.Context, in *Config, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/v0.CoreService/ConfigService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreServiceClient) InfoService(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*RepeatedInfo, error) {
	out := new(RepeatedInfo)
	err := c.cc.Invoke(ctx, "/v0.CoreService/InfoService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreServiceClient) Check(ctx context.Context, in *HealthCheckRequest, opts ...grpc.CallOption) (*HealthCheckResponse, error) {
	out := new(HealthCheckResponse)
	err := c.cc.Invoke(ctx, "/v0.CoreService/Check", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CoreServiceServer is the server API for CoreService service.
type CoreServiceServer interface {
	ConfigService(context.Context, *Config) (*Response, error)
	InfoService(context.Context, *empty.Empty) (*RepeatedInfo, error)
	Check(context.Context, *HealthCheckRequest) (*HealthCheckResponse, error)
}

func RegisterCoreServiceServer(s *grpc.Server, srv CoreServiceServer) {
	s.RegisterService(&_CoreService_serviceDesc, srv)
}

func _CoreService_ConfigService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Config)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServiceServer).ConfigService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v0.CoreService/ConfigService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServiceServer).ConfigService(ctx, req.(*Config))
	}
	return interceptor(ctx, in, info, handler)
}

func _CoreService_InfoService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServiceServer).InfoService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v0.CoreService/InfoService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServiceServer).InfoService(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _CoreService_Check_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HealthCheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServiceServer).Check(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v0.CoreService/Check",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServiceServer).Check(ctx, req.(*HealthCheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CoreService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "v0.CoreService",
	HandlerType: (*CoreServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ConfigService",
			Handler:    _CoreService_ConfigService_Handler,
		},
		{
			MethodName: "InfoService",
			Handler:    _CoreService_InfoService_Handler,
		},
		{
			MethodName: "Check",
			Handler:    _CoreService_Check_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}

func init() { proto.RegisterFile("service.proto", fileDescriptor_service_4ccc745ef771367a) }

var fileDescriptor_service_4ccc745ef771367a = []byte{
	// 208 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x8e, 0xb1, 0x4e, 0xc5, 0x20,
	0x18, 0x85, 0x63, 0xa3, 0x0e, 0xb4, 0x35, 0x86, 0xa1, 0x26, 0xf8, 0x04, 0x2e, 0xb4, 0xd1, 0x41,
	0x67, 0x1b, 0x13, 0x5d, 0xeb, 0xe6, 0xd6, 0xd6, 0x9f, 0xb6, 0xb1, 0xf2, 0x73, 0x81, 0x92, 0xdc,
	0x67, 0xba, 0x2f, 0x79, 0x03, 0x94, 0xe9, 0x6e, 0x9c, 0xef, 0x70, 0xf8, 0x20, 0xa5, 0x01, 0xed,
	0x96, 0x11, 0xb8, 0xd2, 0x68, 0x91, 0x66, 0xae, 0x61, 0x64, 0x91, 0x02, 0x63, 0x66, 0x77, 0x1a,
	0x8c, 0x42, 0x69, 0xf6, 0x9e, 0x15, 0x23, 0x4a, 0xb1, 0x4c, 0x29, 0xcd, 0xd0, 0xaf, 0x76, 0xde,
	0xd3, 0xe3, 0x84, 0x38, 0xad, 0x50, 0x87, 0x34, 0x6c, 0xa2, 0x86, 0x7f, 0x65, 0x8f, 0xb1, 0x7c,
	0x3e, 0x5d, 0x91, 0xbc, 0x45, 0x0d, 0xdf, 0x51, 0x47, 0x9f, 0x48, 0xd9, 0x86, 0xa7, 0x12, 0x20,
	0xdc, 0x35, 0x3c, 0x22, 0x56, 0xf8, 0x73, 0xb7, 0x9b, 0xe9, 0x2b, 0xc9, 0xbf, 0xa4, 0xc0, 0x74,
	0xb1, 0xe2, 0xd1, 0xc3, 0x93, 0x87, 0x7f, 0x78, 0x0f, 0xbb, 0x8f, 0x23, 0x05, 0xbd, 0x85, 0x5f,
	0x3f, 0xa0, 0x6f, 0xe4, 0xa6, 0x9d, 0x61, 0xfc, 0xa3, 0x95, 0xaf, 0x3e, 0xc3, 0x5f, 0x03, 0xe8,
	0xe0, 0xb0, 0x81, 0xb1, 0xec, 0xe1, 0x82, 0x47, 0xe5, 0xfb, 0xf5, 0x4f, 0xe6, 0x9a, 0xe1, 0x36,
	0x18, 0x5e, 0xce, 0x01, 0x00, 0x00, 0xff, 0xff, 0x83, 0x27, 0x70, 0x6d, 0x24, 0x01, 0x00, 0x00,
}