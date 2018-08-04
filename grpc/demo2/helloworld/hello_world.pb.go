// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/razeencheng/demo-go/grpc/demo2/helloworld/hello_world.proto

package helloworld // import "github.com/razeencheng/demo-go/grpc/demo2/helloworld"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import any "github.com/golang/protobuf/ptypes/any"

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

type HelloWorldRequest struct {
	Greeting             string            `protobuf:"bytes,1,opt,name=greeting,proto3" json:"greeting,omitempty"`
	Infos                map[string]string `protobuf:"bytes,2,rep,name=infos,proto3" json:"infos,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *HelloWorldRequest) Reset()         { *m = HelloWorldRequest{} }
func (m *HelloWorldRequest) String() string { return proto.CompactTextString(m) }
func (*HelloWorldRequest) ProtoMessage()    {}
func (*HelloWorldRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_hello_world_e1e511201326df80, []int{0}
}
func (m *HelloWorldRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloWorldRequest.Unmarshal(m, b)
}
func (m *HelloWorldRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloWorldRequest.Marshal(b, m, deterministic)
}
func (dst *HelloWorldRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloWorldRequest.Merge(dst, src)
}
func (m *HelloWorldRequest) XXX_Size() int {
	return xxx_messageInfo_HelloWorldRequest.Size(m)
}
func (m *HelloWorldRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloWorldRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HelloWorldRequest proto.InternalMessageInfo

func (m *HelloWorldRequest) GetGreeting() string {
	if m != nil {
		return m.Greeting
	}
	return ""
}

func (m *HelloWorldRequest) GetInfos() map[string]string {
	if m != nil {
		return m.Infos
	}
	return nil
}

type HelloWorldResponse struct {
	Reply                string     `protobuf:"bytes,1,opt,name=reply,proto3" json:"reply,omitempty"`
	Details              []*any.Any `protobuf:"bytes,2,rep,name=details,proto3" json:"details,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *HelloWorldResponse) Reset()         { *m = HelloWorldResponse{} }
func (m *HelloWorldResponse) String() string { return proto.CompactTextString(m) }
func (*HelloWorldResponse) ProtoMessage()    {}
func (*HelloWorldResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_hello_world_e1e511201326df80, []int{1}
}
func (m *HelloWorldResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloWorldResponse.Unmarshal(m, b)
}
func (m *HelloWorldResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloWorldResponse.Marshal(b, m, deterministic)
}
func (dst *HelloWorldResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloWorldResponse.Merge(dst, src)
}
func (m *HelloWorldResponse) XXX_Size() int {
	return xxx_messageInfo_HelloWorldResponse.Size(m)
}
func (m *HelloWorldResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloWorldResponse.DiscardUnknown(m)
}

var xxx_messageInfo_HelloWorldResponse proto.InternalMessageInfo

func (m *HelloWorldResponse) GetReply() string {
	if m != nil {
		return m.Reply
	}
	return ""
}

func (m *HelloWorldResponse) GetDetails() []*any.Any {
	if m != nil {
		return m.Details
	}
	return nil
}

type HelloWorld struct {
	Msg                  string   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloWorld) Reset()         { *m = HelloWorld{} }
func (m *HelloWorld) String() string { return proto.CompactTextString(m) }
func (*HelloWorld) ProtoMessage()    {}
func (*HelloWorld) Descriptor() ([]byte, []int) {
	return fileDescriptor_hello_world_e1e511201326df80, []int{2}
}
func (m *HelloWorld) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloWorld.Unmarshal(m, b)
}
func (m *HelloWorld) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloWorld.Marshal(b, m, deterministic)
}
func (dst *HelloWorld) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloWorld.Merge(dst, src)
}
func (m *HelloWorld) XXX_Size() int {
	return xxx_messageInfo_HelloWorld.Size(m)
}
func (m *HelloWorld) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloWorld.DiscardUnknown(m)
}

var xxx_messageInfo_HelloWorld proto.InternalMessageInfo

func (m *HelloWorld) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type Error struct {
	Msg                  []string `protobuf:"bytes,1,rep,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Error) Reset()         { *m = Error{} }
func (m *Error) String() string { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()    {}
func (*Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_hello_world_e1e511201326df80, []int{3}
}
func (m *Error) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Error.Unmarshal(m, b)
}
func (m *Error) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Error.Marshal(b, m, deterministic)
}
func (dst *Error) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Error.Merge(dst, src)
}
func (m *Error) XXX_Size() int {
	return xxx_messageInfo_Error.Size(m)
}
func (m *Error) XXX_DiscardUnknown() {
	xxx_messageInfo_Error.DiscardUnknown(m)
}

var xxx_messageInfo_Error proto.InternalMessageInfo

func (m *Error) GetMsg() []string {
	if m != nil {
		return m.Msg
	}
	return nil
}

func init() {
	proto.RegisterType((*HelloWorldRequest)(nil), "helloworld.HelloWorldRequest")
	proto.RegisterMapType((map[string]string)(nil), "helloworld.HelloWorldRequest.InfosEntry")
	proto.RegisterType((*HelloWorldResponse)(nil), "helloworld.HelloWorldResponse")
	proto.RegisterType((*HelloWorld)(nil), "helloworld.HelloWorld")
	proto.RegisterType((*Error)(nil), "helloworld.Error")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// HelloWorldServiceClient is the client API for HelloWorldService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HelloWorldServiceClient interface {
	SayHelloWorld(ctx context.Context, in *HelloWorldRequest, opts ...grpc.CallOption) (*HelloWorldResponse, error)
}

type helloWorldServiceClient struct {
	cc *grpc.ClientConn
}

func NewHelloWorldServiceClient(cc *grpc.ClientConn) HelloWorldServiceClient {
	return &helloWorldServiceClient{cc}
}

func (c *helloWorldServiceClient) SayHelloWorld(ctx context.Context, in *HelloWorldRequest, opts ...grpc.CallOption) (*HelloWorldResponse, error) {
	out := new(HelloWorldResponse)
	err := c.cc.Invoke(ctx, "/helloworld.HelloWorldService/SayHelloWorld", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HelloWorldServiceServer is the server API for HelloWorldService service.
type HelloWorldServiceServer interface {
	SayHelloWorld(context.Context, *HelloWorldRequest) (*HelloWorldResponse, error)
}

func RegisterHelloWorldServiceServer(s *grpc.Server, srv HelloWorldServiceServer) {
	s.RegisterService(&_HelloWorldService_serviceDesc, srv)
}

func _HelloWorldService_SayHelloWorld_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloWorldRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloWorldServiceServer).SayHelloWorld(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/helloworld.HelloWorldService/SayHelloWorld",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloWorldServiceServer).SayHelloWorld(ctx, req.(*HelloWorldRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _HelloWorldService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "helloworld.HelloWorldService",
	HandlerType: (*HelloWorldServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHelloWorld",
			Handler:    _HelloWorldService_SayHelloWorld_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/razeencheng/demo-go/grpc/demo2/helloworld/hello_world.proto",
}

func init() {
	proto.RegisterFile("github.com/razeencheng/demo-go/grpc/demo2/helloworld/hello_world.proto", fileDescriptor_hello_world_e1e511201326df80)
}

var fileDescriptor_hello_world_e1e511201326df80 = []byte{
	// 336 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x52, 0x4b, 0x4b, 0xc3, 0x40,
	0x10, 0x36, 0x2d, 0xf5, 0x31, 0x22, 0xe8, 0xd2, 0x43, 0x0d, 0x58, 0x4a, 0x4e, 0xbd, 0xb8, 0x0b,
	0x55, 0xa4, 0x78, 0x10, 0x14, 0x2a, 0x7a, 0x93, 0xf4, 0x20, 0xf4, 0x22, 0x69, 0x3a, 0xdd, 0x06,
	0xb7, 0x3b, 0x71, 0x93, 0x54, 0xe2, 0x3f, 0xf2, 0x5f, 0xca, 0x26, 0x8d, 0x09, 0x88, 0x1e, 0x3c,
	0x04, 0xe6, 0xdb, 0xef, 0x31, 0x0f, 0x02, 0xf7, 0x32, 0x4a, 0x57, 0xd9, 0x9c, 0x87, 0xb4, 0x16,
	0x26, 0xf8, 0x40, 0xd4, 0xe1, 0x0a, 0xb5, 0x14, 0x0b, 0x5c, 0xd3, 0xb9, 0x24, 0x21, 0x4d, 0x1c,
	0x16, 0x60, 0x24, 0x56, 0xa8, 0x14, 0xbd, 0x93, 0x51, 0x8b, 0xb2, 0x7c, 0x29, 0x6a, 0x1e, 0x1b,
	0x4a, 0x89, 0x41, 0xcd, 0xba, 0xa2, 0x91, 0x29, 0x49, 0x05, 0x5a, 0x8a, 0x42, 0x34, 0xcf, 0x96,
	0x22, 0x4e, 0xf3, 0x18, 0x13, 0x11, 0xe8, 0xdc, 0x7e, 0xa5, 0xd9, 0xfb, 0x74, 0xe0, 0xe4, 0xc1,
	0xfa, 0x9f, 0xad, 0xdf, 0xc7, 0xb7, 0x0c, 0x93, 0x94, 0xb9, 0xb0, 0x2f, 0x0d, 0x62, 0x1a, 0x69,
	0xd9, 0x73, 0x06, 0xce, 0xf0, 0xc0, 0xff, 0xc6, 0xec, 0x06, 0x3a, 0x91, 0x5e, 0x52, 0xd2, 0x6b,
	0x0d, 0xda, 0xc3, 0xc3, 0xd1, 0x90, 0xd7, 0xed, 0xf9, 0x8f, 0x24, 0xfe, 0x68, 0xa5, 0x13, 0x9d,
	0x9a, 0xdc, 0x2f, 0x6d, 0xee, 0x18, 0xa0, 0x7e, 0x64, 0xc7, 0xd0, 0x7e, 0xc5, 0x7c, 0xdb, 0xc4,
	0x96, 0xac, 0x0b, 0x9d, 0x4d, 0xa0, 0x32, 0xec, 0xb5, 0x8a, 0xb7, 0x12, 0x5c, 0xb7, 0xc6, 0x8e,
	0x37, 0x03, 0xd6, 0x6c, 0x90, 0xc4, 0xa4, 0x13, 0xb4, 0x7a, 0x83, 0xb1, 0xaa, 0x32, 0x4a, 0xc0,
	0x38, 0xec, 0x2d, 0x30, 0x0d, 0x22, 0x55, 0xcd, 0xd9, 0xe5, 0x92, 0x48, 0x2a, 0xe4, 0xd5, 0x3d,
	0xf8, 0xad, 0xce, 0xfd, 0x4a, 0xe4, 0xf5, 0x01, 0xea, 0x6c, 0x3b, 0xd5, 0x3a, 0xa9, 0x56, 0xb7,
	0xa5, 0x77, 0x0a, 0x9d, 0x89, 0x31, 0x64, 0x6a, 0xaa, 0xbd, 0xa5, 0x46, 0xd8, 0xbc, 0xe0, 0x14,
	0xcd, 0x26, 0x0a, 0x91, 0x3d, 0xc1, 0xd1, 0x34, 0xc8, 0x1b, 0x91, 0x67, 0x7f, 0xde, 0xc9, 0xed,
	0xff, 0x46, 0x97, 0x5b, 0x7a, 0x3b, 0x77, 0x57, 0xb3, 0xcb, 0xff, 0xfc, 0x30, 0xf3, 0xdd, 0x62,
	0xe1, 0x8b, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xcf, 0x1c, 0x94, 0xd7, 0x6f, 0x02, 0x00, 0x00,
}
