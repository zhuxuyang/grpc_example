// Code generated by protoc-gen-go. DO NOT EDIT.
// source: example.proto

package protos

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type HelloRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	Time                 int64    `protobuf:"varint,2,opt,name=Time,proto3" json:"Time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloRequest) Reset()         { *m = HelloRequest{} }
func (m *HelloRequest) String() string { return proto.CompactTextString(m) }
func (*HelloRequest) ProtoMessage()    {}
func (*HelloRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_15a1dc8d40dadaa6, []int{0}
}

func (m *HelloRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloRequest.Unmarshal(m, b)
}
func (m *HelloRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloRequest.Marshal(b, m, deterministic)
}
func (m *HelloRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloRequest.Merge(m, src)
}
func (m *HelloRequest) XXX_Size() int {
	return xxx_messageInfo_HelloRequest.Size(m)
}
func (m *HelloRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HelloRequest proto.InternalMessageInfo

func (m *HelloRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *HelloRequest) GetTime() int64 {
	if m != nil {
		return m.Time
	}
	return 0
}

type HelloResponse struct {
	Answer               string   `protobuf:"bytes,1,opt,name=Answer,proto3" json:"Answer,omitempty"`
	Time                 int64    `protobuf:"varint,2,opt,name=Time,proto3" json:"Time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloResponse) Reset()         { *m = HelloResponse{} }
func (m *HelloResponse) String() string { return proto.CompactTextString(m) }
func (*HelloResponse) ProtoMessage()    {}
func (*HelloResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_15a1dc8d40dadaa6, []int{1}
}

func (m *HelloResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloResponse.Unmarshal(m, b)
}
func (m *HelloResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloResponse.Marshal(b, m, deterministic)
}
func (m *HelloResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloResponse.Merge(m, src)
}
func (m *HelloResponse) XXX_Size() int {
	return xxx_messageInfo_HelloResponse.Size(m)
}
func (m *HelloResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloResponse.DiscardUnknown(m)
}

var xxx_messageInfo_HelloResponse proto.InternalMessageInfo

func (m *HelloResponse) GetAnswer() string {
	if m != nil {
		return m.Answer
	}
	return ""
}

func (m *HelloResponse) GetTime() int64 {
	if m != nil {
		return m.Time
	}
	return 0
}

func init() {
	proto.RegisterType((*HelloRequest)(nil), "protos.HelloRequest")
	proto.RegisterType((*HelloResponse)(nil), "protos.HelloResponse")
}

func init() { proto.RegisterFile("example.proto", fileDescriptor_15a1dc8d40dadaa6) }

var fileDescriptor_15a1dc8d40dadaa6 = []byte{
	// 153 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4d, 0xad, 0x48, 0xcc,
	0x2d, 0xc8, 0x49, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x03, 0x53, 0xc5, 0x4a, 0x66,
	0x5c, 0x3c, 0x1e, 0xa9, 0x39, 0x39, 0xf9, 0x41, 0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25, 0x42, 0x42,
	0x5c, 0x2c, 0x7e, 0x89, 0xb9, 0xa9, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x60, 0x36, 0x48,
	0x2c, 0x24, 0x33, 0x37, 0x55, 0x82, 0x49, 0x81, 0x51, 0x83, 0x39, 0x08, 0xcc, 0x56, 0xb2, 0xe6,
	0xe2, 0x85, 0xea, 0x2b, 0x2e, 0xc8, 0xcf, 0x2b, 0x4e, 0x15, 0x12, 0xe3, 0x62, 0x73, 0xcc, 0x2b,
	0x2e, 0x4f, 0x2d, 0x82, 0x6a, 0x85, 0xf2, 0xb0, 0x69, 0x36, 0x72, 0xe4, 0x62, 0x77, 0x85, 0xb8,
	0x46, 0xc8, 0x8c, 0x8b, 0x15, 0x6c, 0x8e, 0x90, 0x08, 0xc4, 0x61, 0xc5, 0x7a, 0xc8, 0xce, 0x91,
	0x12, 0x45, 0x13, 0x85, 0x58, 0xa6, 0xc4, 0x90, 0x04, 0x71, 0xbf, 0x31, 0x20, 0x00, 0x00, 0xff,
	0xff, 0x96, 0x48, 0xc9, 0x42, 0xd7, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ExampleClient is the router API for Example service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ExampleClient interface {
	// 只是返回zh_user表
	Hello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error)
}

type exampleClient struct {
	cc *grpc.ClientConn
}

func NewExampleClient(cc *grpc.ClientConn) ExampleClient {
	return &exampleClient{cc}
}

func (c *exampleClient) Hello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error) {
	out := new(HelloResponse)
	err := c.cc.Invoke(ctx, "/protos.Example/Hello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ExampleServer is the server API for Example service.
type ExampleServer interface {
	// 只是返回zh_user表
	Hello(context.Context, *HelloRequest) (*HelloResponse, error)
}

// UnimplementedExampleServer can be embedded to have forward compatible implementations.
type UnimplementedExampleServer struct {
}

func (*UnimplementedExampleServer) Hello(ctx context.Context, req *HelloRequest) (*HelloResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Hello not implemented")
}

func RegisterExampleServer(s *grpc.Server, srv ExampleServer) {
	s.RegisterService(&_Example_serviceDesc, srv)
}

func _Example_Hello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExampleServer).Hello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Example/Hello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExampleServer).Hello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Example_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protos.Example",
	HandlerType: (*ExampleServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Hello",
			Handler:    _Example_Hello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "example.proto",
}
