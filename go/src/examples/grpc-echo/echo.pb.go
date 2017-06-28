// Code generated by protoc-gen-go. DO NOT EDIT.
// source: examples/grpc-echo/echo.proto

/*
Package echo is a generated protocol buffer package.

It is generated from these files:
	examples/grpc-echo/echo.proto

It has these top-level messages:
	PingReq
	PingRes
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

type PingReq struct {
	Ping string `protobuf:"bytes,1,opt,name=ping" json:"ping,omitempty"`
}

func (m *PingReq) Reset()                    { *m = PingReq{} }
func (m *PingReq) String() string            { return proto.CompactTextString(m) }
func (*PingReq) ProtoMessage()               {}
func (*PingReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *PingReq) GetPing() string {
	if m != nil {
		return m.Ping
	}
	return ""
}

type PingRes struct {
	Pong string `protobuf:"bytes,1,opt,name=pong" json:"pong,omitempty"`
}

func (m *PingRes) Reset()                    { *m = PingRes{} }
func (m *PingRes) String() string            { return proto.CompactTextString(m) }
func (*PingRes) ProtoMessage()               {}
func (*PingRes) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *PingRes) GetPong() string {
	if m != nil {
		return m.Pong
	}
	return ""
}

func init() {
	proto.RegisterType((*PingReq)(nil), "echo.PingReq")
	proto.RegisterType((*PingRes)(nil), "echo.PingRes")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Echo service

type EchoClient interface {
	Ping(ctx context.Context, in *PingReq, opts ...grpc.CallOption) (*PingRes, error)
}

type echoClient struct {
	cc *grpc.ClientConn
}

func NewEchoClient(cc *grpc.ClientConn) EchoClient {
	return &echoClient{cc}
}

func (c *echoClient) Ping(ctx context.Context, in *PingReq, opts ...grpc.CallOption) (*PingRes, error) {
	out := new(PingRes)
	err := grpc.Invoke(ctx, "/echo.Echo/Ping", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Echo service

type EchoServer interface {
	Ping(context.Context, *PingReq) (*PingRes, error)
}

func RegisterEchoServer(s *grpc.Server, srv EchoServer) {
	s.RegisterService(&_Echo_serviceDesc, srv)
}

func _Echo_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EchoServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/echo.Echo/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EchoServer).Ping(ctx, req.(*PingReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Echo_serviceDesc = grpc.ServiceDesc{
	ServiceName: "echo.Echo",
	HandlerType: (*EchoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _Echo_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "examples/grpc-echo/echo.proto",
}

func init() { proto.RegisterFile("examples/grpc-echo/echo.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 127 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x92, 0x4d, 0xad, 0x48, 0xcc,
	0x2d, 0xc8, 0x49, 0x2d, 0xd6, 0x4f, 0x2f, 0x2a, 0x48, 0xd6, 0x4d, 0x4d, 0xce, 0xc8, 0xd7, 0x07,
	0x11, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0x2c, 0x20, 0xb6, 0x92, 0x2c, 0x17, 0x7b, 0x40,
	0x66, 0x5e, 0x7a, 0x50, 0x6a, 0xa1, 0x90, 0x10, 0x17, 0x4b, 0x41, 0x66, 0x5e, 0xba, 0x04, 0xa3,
	0x02, 0xa3, 0x06, 0x67, 0x10, 0x98, 0x8d, 0x90, 0x2e, 0x06, 0x4b, 0xe7, 0x23, 0x49, 0xe7, 0xe7,
	0xa5, 0x1b, 0xe9, 0x71, 0xb1, 0xb8, 0x26, 0x67, 0xe4, 0x0b, 0xa9, 0x71, 0xb1, 0x80, 0x94, 0x09,
	0xf1, 0xea, 0x81, 0x2d, 0x80, 0x9a, 0x28, 0x85, 0xc2, 0x2d, 0x56, 0x62, 0x48, 0x62, 0x03, 0x5b,
	0x6d, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x71, 0xab, 0x65, 0x96, 0x9b, 0x00, 0x00, 0x00,
}
