// Code generated by protoc-gen-go. DO NOT EDIT.
// source: echo.proto

/*
Package echo is a generated protocol buffer package.

It is generated from these files:
	echo.proto

It has these top-level messages:
	GreetReq
	GreetRes
*/
package echo

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

type GreetReq struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *GreetReq) Reset()                    { *m = GreetReq{} }
func (m *GreetReq) String() string            { return proto.CompactTextString(m) }
func (*GreetReq) ProtoMessage()               {}
func (*GreetReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *GreetReq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type GreetRes struct {
	Greeting string `protobuf:"bytes,1,opt,name=greeting" json:"greeting,omitempty"`
}

func (m *GreetRes) Reset()                    { *m = GreetRes{} }
func (m *GreetRes) String() string            { return proto.CompactTextString(m) }
func (*GreetRes) ProtoMessage()               {}
func (*GreetRes) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *GreetRes) GetGreeting() string {
	if m != nil {
		return m.Greeting
	}
	return ""
}

func init() {
	proto.RegisterType((*GreetReq)(nil), "echo.GreetReq")
	proto.RegisterType((*GreetRes)(nil), "echo.GreetRes")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Echo service

type EchoClient interface {
	Greet(ctx context.Context, in *GreetReq, opts ...grpc.CallOption) (*GreetRes, error)
}

type echoClient struct {
	cc *grpc.ClientConn
}

func NewEchoClient(cc *grpc.ClientConn) EchoClient {
	return &echoClient{cc}
}

func (c *echoClient) Greet(ctx context.Context, in *GreetReq, opts ...grpc.CallOption) (*GreetRes, error) {
	out := new(GreetRes)
	err := grpc.Invoke(ctx, "/echo.Echo/Greet", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Echo service

type EchoServer interface {
	Greet(context.Context, *GreetReq) (*GreetRes, error)
}

func RegisterEchoServer(s *grpc.Server, srv EchoServer) {
	s.RegisterService(&_Echo_serviceDesc, srv)
}

func _Echo_Greet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GreetReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EchoServer).Greet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/echo.Echo/Greet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EchoServer).Greet(ctx, req.(*GreetReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Echo_serviceDesc = grpc.ServiceDesc{
	ServiceName: "echo.Echo",
	HandlerType: (*EchoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Greet",
			Handler:    _Echo_Greet_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "echo.proto",
}

func init() { proto.RegisterFile("echo.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 119 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0x4d, 0xce, 0xc8,
	0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x01, 0xb1, 0x95, 0xe4, 0xb8, 0x38, 0xdc, 0x8b,
	0x52, 0x53, 0x4b, 0x82, 0x52, 0x0b, 0x85, 0x84, 0xb8, 0x58, 0xf2, 0x12, 0x73, 0x53, 0x25, 0x18,
	0x15, 0x18, 0x35, 0x38, 0x83, 0xc0, 0x6c, 0x25, 0x35, 0xb8, 0x7c, 0xb1, 0x90, 0x14, 0x17, 0x47,
	0x3a, 0x88, 0x9d, 0x99, 0x97, 0x0e, 0x55, 0x03, 0xe7, 0x1b, 0x19, 0x72, 0xb1, 0xb8, 0x26, 0x67,
	0xe4, 0x0b, 0x69, 0x72, 0xb1, 0x82, 0xd5, 0x0b, 0xf1, 0xe9, 0x81, 0xed, 0x82, 0x19, 0x2e, 0x85,
	0xca, 0x2f, 0x56, 0x62, 0x48, 0x62, 0x03, 0xbb, 0xc3, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xb1,
	0xbd, 0xca, 0x48, 0x95, 0x00, 0x00, 0x00,
}