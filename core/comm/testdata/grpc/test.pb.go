// Code generated by protoc-gen-go. DO NOT EDIT.
// source: test.proto

/*
Package grpc is a generated protocol buffer package.

It is generated from these files:
	test.proto

It has these top-level messages:
	Empty
	Echo
*/
package grpc

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc1 "google.golang.org/grpc"
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

type Empty struct {
}

func (m *Empty) Reset()                    { *m = Empty{} }
func (m *Empty) String() string            { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()               {}
func (*Empty) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Echo struct {
	Payload []byte `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (m *Echo) Reset()                    { *m = Echo{} }
func (m *Echo) String() string            { return proto.CompactTextString(m) }
func (*Echo) ProtoMessage()               {}
func (*Echo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Echo) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func init() {
	proto.RegisterType((*Empty)(nil), "Empty")
	proto.RegisterType((*Echo)(nil), "Echo")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc1.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc1.SupportPackageIsVersion4

// Client API for TestService service

type TestServiceClient interface {
	EmptyCall(ctx context.Context, in *Empty, opts ...grpc1.CallOption) (*Empty, error)
}

type testServiceClient struct {
	cc *grpc1.ClientConn
}

func NewTestServiceClient(cc *grpc1.ClientConn) TestServiceClient {
	return &testServiceClient{cc}
}

func (c *testServiceClient) EmptyCall(ctx context.Context, in *Empty, opts ...grpc1.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc1.Invoke(ctx, "/TestService/EmptyCall", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for TestService service

type TestServiceServer interface {
	EmptyCall(context.Context, *Empty) (*Empty, error)
}

func RegisterTestServiceServer(s *grpc1.Server, srv TestServiceServer) {
	s.RegisterService(&_TestService_serviceDesc, srv)
}

func _TestService_EmptyCall_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc1.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServiceServer).EmptyCall(ctx, in)
	}
	info := &grpc1.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/TestService/EmptyCall",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServiceServer).EmptyCall(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _TestService_serviceDesc = grpc1.ServiceDesc{
	ServiceName: "TestService",
	HandlerType: (*TestServiceServer)(nil),
	Methods: []grpc1.MethodDesc{
		{
			MethodName: "EmptyCall",
			Handler:    _TestService_EmptyCall_Handler,
		},
	},
	Streams:  []grpc1.StreamDesc{},
	Metadata: "test.proto",
}

// Client API for EchoService service

type EchoServiceClient interface {
	EchoCall(ctx context.Context, in *Echo, opts ...grpc1.CallOption) (*Echo, error)
}

type echoServiceClient struct {
	cc *grpc1.ClientConn
}

func NewEchoServiceClient(cc *grpc1.ClientConn) EchoServiceClient {
	return &echoServiceClient{cc}
}

func (c *echoServiceClient) EchoCall(ctx context.Context, in *Echo, opts ...grpc1.CallOption) (*Echo, error) {
	out := new(Echo)
	err := grpc1.Invoke(ctx, "/EchoService/EchoCall", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for EchoService service

type EchoServiceServer interface {
	EchoCall(context.Context, *Echo) (*Echo, error)
}

func RegisterEchoServiceServer(s *grpc1.Server, srv EchoServiceServer) {
	s.RegisterService(&_EchoService_serviceDesc, srv)
}

func _EchoService_EchoCall_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc1.UnaryServerInterceptor) (interface{}, error) {
	in := new(Echo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EchoServiceServer).EchoCall(ctx, in)
	}
	info := &grpc1.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/EchoService/EchoCall",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EchoServiceServer).EchoCall(ctx, req.(*Echo))
	}
	return interceptor(ctx, in, info, handler)
}

var _EchoService_serviceDesc = grpc1.ServiceDesc{
	ServiceName: "EchoService",
	HandlerType: (*EchoServiceServer)(nil),
	Methods: []grpc1.MethodDesc{
		{
			MethodName: "EchoCall",
			Handler:    _EchoService_EchoCall_Handler,
		},
	},
	Streams:  []grpc1.StreamDesc{},
	Metadata: "test.proto",
}

func init() { proto.RegisterFile("test.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 186 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x34, 0x8e, 0x41, 0x8b, 0x83, 0x30,
	0x10, 0x85, 0x11, 0x56, 0xdd, 0x1d, 0xf7, 0xe4, 0x49, 0xdc, 0x8b, 0x78, 0xd9, 0xd2, 0x43, 0x02,
	0x96, 0xd2, 0x7b, 0x8b, 0x7f, 0xa0, 0xed, 0xa9, 0xb7, 0x18, 0xa7, 0x2a, 0x24, 0x24, 0xc4, 0x69,
	0xc1, 0x7f, 0x5f, 0x92, 0xd6, 0xd3, 0x7b, 0x0f, 0xbe, 0x19, 0x3e, 0x00, 0xc2, 0x99, 0x98, 0x75,
	0x86, 0x4c, 0x9d, 0x42, 0xdc, 0x6a, 0x4b, 0x4b, 0x5d, 0xc1, 0x57, 0x2b, 0x47, 0x93, 0x17, 0x90,
	0x5a, 0xb1, 0x28, 0x23, 0xfa, 0x22, 0xaa, 0xa2, 0xcd, 0xef, 0x79, 0x9d, 0xcd, 0x16, 0xb2, 0x2b,
	0xce, 0x74, 0x41, 0xf7, 0x9c, 0x24, 0xe6, 0x7f, 0xf0, 0x13, 0x2e, 0x4f, 0x42, 0xa9, 0x3c, 0x61,
	0xa1, 0x97, 0x9f, 0x6c, 0xfe, 0x21, 0xf3, 0xdf, 0x56, 0xb6, 0x80, 0x6f, 0x3f, 0x03, 0x1a, 0x33,
	0x5f, 0xcb, 0x77, 0x1c, 0x0f, 0xb7, 0xfd, 0x30, 0xd1, 0xf8, 0xe8, 0x98, 0x34, 0x9a, 0x8f, 0x8b,
	0x45, 0xa7, 0xb0, 0x1f, 0xd0, 0xf1, 0xbb, 0xe8, 0xdc, 0x24, 0xb9, 0x34, 0x0e, 0xb9, 0x34, 0x5a,
	0x73, 0x6f, 0xdd, 0x0b, 0x12, 0x7c, 0x70, 0x56, 0x76, 0x49, 0xf0, 0xdf, 0xbd, 0x02, 0x00, 0x00,
	0xff, 0xff, 0xa7, 0xf6, 0xb7, 0xb5, 0xcd, 0x00, 0x00, 0x00,
}
