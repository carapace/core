// Code generated by protoc-gen-go. DO NOT EDIT.
// source: definitions/service.proto

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

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CoreClient is the client API for Core service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CoreClient interface {
	Config(ctx context.Context, in *Config, opts ...grpc.CallOption) (*Response, error)
}

type coreClient struct {
	cc *grpc.ClientConn
}

func NewCoreClient(cc *grpc.ClientConn) CoreClient {
	return &coreClient{cc}
}

func (c *coreClient) Config(ctx context.Context, in *Config, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/v1.response.Core/Config", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CoreServer is the server API for Core service.
type CoreServer interface {
	Config(context.Context, *Config) (*Response, error)
}

func RegisterCoreServer(s *grpc.Server, srv CoreServer) {
	s.RegisterService(&_Core_serviceDesc, srv)
}

func _Core_Config_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Config)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServer).Config(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.response.Core/Config",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServer).Config(ctx, req.(*Config))
	}
	return interceptor(ctx, in, info, handler)
}

var _Core_serviceDesc = grpc.ServiceDesc{
	ServiceName: "v1.response.Core",
	HandlerType: (*CoreServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Config",
			Handler:    _Core_Config_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "definitions/service.proto",
}

func init() { proto.RegisterFile("definitions/service.proto", fileDescriptor_service_f0c1fd36b204390a) }

var fileDescriptor_service_f0c1fd36b204390a = []byte{
	// 129 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4c, 0x49, 0x4d, 0xcb,
	0xcc, 0xcb, 0x2c, 0xc9, 0xcc, 0xcf, 0x2b, 0xd6, 0x2f, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x2e, 0x33, 0xd4, 0x2b, 0x4a, 0x2d, 0x2e, 0xc8, 0xcf,
	0x2b, 0x4e, 0x95, 0x92, 0x42, 0x56, 0x07, 0x13, 0x85, 0x28, 0x94, 0x92, 0x40, 0x96, 0x4b, 0xce,
	0xcf, 0x4b, 0xcb, 0x4c, 0x87, 0xc8, 0x18, 0x59, 0x71, 0xb1, 0x38, 0xe7, 0x17, 0xa5, 0x0a, 0x19,
	0x71, 0xb1, 0x39, 0x83, 0xc5, 0x85, 0x04, 0xf5, 0xca, 0x0c, 0xf5, 0xa0, 0x6a, 0x20, 0x42, 0x52,
	0xa2, 0x7a, 0x48, 0x16, 0xe9, 0x05, 0x41, 0x19, 0x4e, 0x2c, 0x51, 0x4c, 0x65, 0x86, 0x49, 0x6c,
	0x60, 0x83, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x5a, 0x95, 0x38, 0x13, 0xa8, 0x00, 0x00,
	0x00,
}