// Code generated by protoc-gen-go. DO NOT EDIT.
// source: helloworld.proto

/*
Package helloworld is a generated protocol buffer package.

It is generated from these files:
	helloworld.proto

It has these top-level messages:
	SayHelloInput
	SayHelloOutput
	CustomList
*/
package helloworld

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

type SayHelloInput struct {
	FirstName           string `protobuf:"bytes,1,opt,name=FirstName" json:"FirstName,omitempty"`
	LastName            string `protobuf:"bytes,2,opt,name=LastName" json:"LastName,omitempty"`
	ForceError          bool   `protobuf:"varint,3,opt,name=ForceError" json:"ForceError,omitempty"`
	DesiredErrorMessage string `protobuf:"bytes,4,opt,name=DesiredErrorMessage" json:"DesiredErrorMessage,omitempty"`
}

func (m *SayHelloInput) Reset()                    { *m = SayHelloInput{} }
func (m *SayHelloInput) String() string            { return proto.CompactTextString(m) }
func (*SayHelloInput) ProtoMessage()               {}
func (*SayHelloInput) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *SayHelloInput) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *SayHelloInput) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *SayHelloInput) GetForceError() bool {
	if m != nil {
		return m.ForceError
	}
	return false
}

func (m *SayHelloInput) GetDesiredErrorMessage() string {
	if m != nil {
		return m.DesiredErrorMessage
	}
	return ""
}

type SayHelloOutput struct {
	Msg string `protobuf:"bytes,1,opt,name=Msg" json:"Msg,omitempty"`
}

func (m *SayHelloOutput) Reset()                    { *m = SayHelloOutput{} }
func (m *SayHelloOutput) String() string            { return proto.CompactTextString(m) }
func (*SayHelloOutput) ProtoMessage()               {}
func (*SayHelloOutput) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *SayHelloOutput) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type CustomList struct {
	List []int32 `protobuf:"varint,1,rep,packed,name=List" json:"List,omitempty"`
}

func (m *CustomList) Reset()                    { *m = CustomList{} }
func (m *CustomList) String() string            { return proto.CompactTextString(m) }
func (*CustomList) ProtoMessage()               {}
func (*CustomList) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *CustomList) GetList() []int32 {
	if m != nil {
		return m.List
	}
	return nil
}

func init() {
	proto.RegisterType((*SayHelloInput)(nil), "SayHelloInput")
	proto.RegisterType((*SayHelloOutput)(nil), "SayHelloOutput")
	proto.RegisterType((*CustomList)(nil), "CustomList")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for HelloWorld service

type HelloWorldClient interface {
	SayHello(ctx context.Context, in *SayHelloInput, opts ...grpc.CallOption) (*SayHelloOutput, error)
	Randomise(ctx context.Context, in *CustomList, opts ...grpc.CallOption) (*CustomList, error)
}

type helloWorldClient struct {
	cc *grpc.ClientConn
}

func NewHelloWorldClient(cc *grpc.ClientConn) HelloWorldClient {
	return &helloWorldClient{cc}
}

func (c *helloWorldClient) SayHello(ctx context.Context, in *SayHelloInput, opts ...grpc.CallOption) (*SayHelloOutput, error) {
	out := new(SayHelloOutput)
	err := grpc.Invoke(ctx, "/HelloWorld/SayHello", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *helloWorldClient) Randomise(ctx context.Context, in *CustomList, opts ...grpc.CallOption) (*CustomList, error) {
	out := new(CustomList)
	err := grpc.Invoke(ctx, "/HelloWorld/Randomise", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for HelloWorld service

type HelloWorldServer interface {
	SayHello(context.Context, *SayHelloInput) (*SayHelloOutput, error)
	Randomise(context.Context, *CustomList) (*CustomList, error)
}

func RegisterHelloWorldServer(s *grpc.Server, srv HelloWorldServer) {
	s.RegisterService(&_HelloWorld_serviceDesc, srv)
}

func _HelloWorld_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SayHelloInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloWorldServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/HelloWorld/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloWorldServer).SayHello(ctx, req.(*SayHelloInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _HelloWorld_Randomise_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CustomList)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloWorldServer).Randomise(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/HelloWorld/Randomise",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloWorldServer).Randomise(ctx, req.(*CustomList))
	}
	return interceptor(ctx, in, info, handler)
}

var _HelloWorld_serviceDesc = grpc.ServiceDesc{
	ServiceName: "HelloWorld",
	HandlerType: (*HelloWorldServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _HelloWorld_SayHello_Handler,
		},
		{
			MethodName: "Randomise",
			Handler:    _HelloWorld_Randomise_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "helloworld.proto",
}

func init() { proto.RegisterFile("helloworld.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 240 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x50, 0xdd, 0x4a, 0xc3, 0x30,
	0x18, 0x25, 0x76, 0x4a, 0x7b, 0xc4, 0x39, 0x3e, 0x6f, 0x4a, 0x11, 0x29, 0x05, 0xa1, 0x20, 0x14,
	0xd1, 0x47, 0x50, 0x87, 0xc2, 0xa6, 0x50, 0x2f, 0xbc, 0x35, 0xda, 0x30, 0x03, 0xed, 0x32, 0xf2,
	0xa5, 0x88, 0xcf, 0xe2, 0xcb, 0x4a, 0xc2, 0xe6, 0xaa, 0xec, 0x2a, 0xdf, 0x39, 0x07, 0xce, 0x4f,
	0x30, 0xf9, 0x50, 0x6d, 0x6b, 0x3e, 0x8d, 0x6d, 0x9b, 0x6a, 0x65, 0x8d, 0x33, 0xc5, 0xb7, 0xc0,
	0xd1, 0xb3, 0xfc, 0xba, 0xf7, 0xfc, 0xc3, 0x72, 0xd5, 0x3b, 0x3a, 0x45, 0x32, 0xd5, 0x96, 0xdd,
	0xa3, 0xec, 0x54, 0x2a, 0x72, 0x51, 0x26, 0xf5, 0x96, 0xa0, 0x0c, 0xf1, 0x4c, 0xae, 0xc5, 0xbd,
	0x20, 0xfe, 0x62, 0x3a, 0x03, 0xa6, 0xc6, 0xbe, 0xab, 0x3b, 0x6b, 0x8d, 0x4d, 0xa3, 0x5c, 0x94,
	0x71, 0x3d, 0x60, 0xe8, 0x12, 0x27, 0xb7, 0x8a, 0xb5, 0x55, 0x4d, 0xc0, 0x73, 0xc5, 0x2c, 0x17,
	0x2a, 0x1d, 0x05, 0x9b, 0x5d, 0x52, 0x51, 0x60, 0xbc, 0x29, 0xf7, 0xd4, 0x3b, 0xdf, 0x6e, 0x82,
	0x68, 0xce, 0x8b, 0x75, 0x2f, 0x7f, 0x16, 0x39, 0x70, 0xd3, 0xb3, 0x33, 0xdd, 0x4c, 0xb3, 0x23,
	0xc2, 0xc8, 0xbf, 0xa9, 0xc8, 0xa3, 0x72, 0xbf, 0x0e, 0xf7, 0xd5, 0x2b, 0x10, 0x2c, 0x5e, 0xfc,
	0x6e, 0xba, 0x40, 0xbc, 0xf1, 0xa4, 0x71, 0xf5, 0x67, 0x7b, 0x76, 0x5c, 0xfd, 0x8b, 0x3b, 0x47,
	0x52, 0xcb, 0x65, 0x63, 0x3a, 0xcd, 0x8a, 0x0e, 0xab, 0x6d, 0x50, 0x36, 0x04, 0x6f, 0x07, 0xe1,
	0x33, 0xaf, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0xb7, 0xf0, 0xa0, 0x56, 0x60, 0x01, 0x00, 0x00,
}
