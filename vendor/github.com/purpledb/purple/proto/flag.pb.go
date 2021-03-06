// Code generated by protoc-gen-go. DO NOT EDIT.
// source: flag.proto

package proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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

type FlagGetRequest struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FlagGetRequest) Reset()         { *m = FlagGetRequest{} }
func (m *FlagGetRequest) String() string { return proto.CompactTextString(m) }
func (*FlagGetRequest) ProtoMessage()    {}
func (*FlagGetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_01fdf51d06af45bb, []int{0}
}

func (m *FlagGetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FlagGetRequest.Unmarshal(m, b)
}
func (m *FlagGetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FlagGetRequest.Marshal(b, m, deterministic)
}
func (m *FlagGetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FlagGetRequest.Merge(m, src)
}
func (m *FlagGetRequest) XXX_Size() int {
	return xxx_messageInfo_FlagGetRequest.Size(m)
}
func (m *FlagGetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FlagGetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FlagGetRequest proto.InternalMessageInfo

func (m *FlagGetRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type FlagSetRequest struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value                bool     `protobuf:"varint,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FlagSetRequest) Reset()         { *m = FlagSetRequest{} }
func (m *FlagSetRequest) String() string { return proto.CompactTextString(m) }
func (*FlagSetRequest) ProtoMessage()    {}
func (*FlagSetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_01fdf51d06af45bb, []int{1}
}

func (m *FlagSetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FlagSetRequest.Unmarshal(m, b)
}
func (m *FlagSetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FlagSetRequest.Marshal(b, m, deterministic)
}
func (m *FlagSetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FlagSetRequest.Merge(m, src)
}
func (m *FlagSetRequest) XXX_Size() int {
	return xxx_messageInfo_FlagSetRequest.Size(m)
}
func (m *FlagSetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FlagSetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FlagSetRequest proto.InternalMessageInfo

func (m *FlagSetRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *FlagSetRequest) GetValue() bool {
	if m != nil {
		return m.Value
	}
	return false
}

type FlagResponse struct {
	Value                bool     `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FlagResponse) Reset()         { *m = FlagResponse{} }
func (m *FlagResponse) String() string { return proto.CompactTextString(m) }
func (*FlagResponse) ProtoMessage()    {}
func (*FlagResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_01fdf51d06af45bb, []int{2}
}

func (m *FlagResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FlagResponse.Unmarshal(m, b)
}
func (m *FlagResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FlagResponse.Marshal(b, m, deterministic)
}
func (m *FlagResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FlagResponse.Merge(m, src)
}
func (m *FlagResponse) XXX_Size() int {
	return xxx_messageInfo_FlagResponse.Size(m)
}
func (m *FlagResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FlagResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FlagResponse proto.InternalMessageInfo

func (m *FlagResponse) GetValue() bool {
	if m != nil {
		return m.Value
	}
	return false
}

func init() {
	proto.RegisterType((*FlagGetRequest)(nil), "proto.FlagGetRequest")
	proto.RegisterType((*FlagSetRequest)(nil), "proto.FlagSetRequest")
	proto.RegisterType((*FlagResponse)(nil), "proto.FlagResponse")
}

func init() { proto.RegisterFile("flag.proto", fileDescriptor_01fdf51d06af45bb) }

var fileDescriptor_01fdf51d06af45bb = []byte{
	// 175 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0xcb, 0x49, 0x4c,
	0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x53, 0x52, 0x3c, 0xc9, 0xf9, 0xb9, 0xb9,
	0xf9, 0x79, 0x10, 0x41, 0x25, 0x25, 0x2e, 0x3e, 0xb7, 0x9c, 0xc4, 0x74, 0xf7, 0xd4, 0x92, 0xa0,
	0xd4, 0xc2, 0xd2, 0xd4, 0xe2, 0x12, 0x21, 0x01, 0x2e, 0xe6, 0xec, 0xd4, 0x4a, 0x09, 0x46, 0x05,
	0x46, 0x0d, 0xce, 0x20, 0x10, 0x53, 0xc9, 0x02, 0xa2, 0x26, 0x18, 0x8f, 0x1a, 0x21, 0x11, 0x2e,
	0xd6, 0xb2, 0xc4, 0x9c, 0xd2, 0x54, 0x09, 0x26, 0x05, 0x46, 0x0d, 0x8e, 0x20, 0x08, 0x47, 0x49,
	0x85, 0x8b, 0x07, 0xa4, 0x33, 0x28, 0xb5, 0xb8, 0x20, 0x3f, 0xaf, 0x38, 0x15, 0xa1, 0x8a, 0x11,
	0x49, 0x95, 0x51, 0x2e, 0x17, 0x0b, 0x48, 0x95, 0x90, 0x29, 0x17, 0x3b, 0xd4, 0x2d, 0x42, 0xa2,
	0x10, 0xe7, 0xe9, 0xa1, 0xba, 0x4d, 0x4a, 0x18, 0x49, 0x18, 0x6e, 0xa8, 0x1e, 0x44, 0x5b, 0x30,
	0x9a, 0x36, 0x84, 0x73, 0xa5, 0x78, 0xa0, 0xc2, 0xae, 0xb9, 0x05, 0x25, 0x95, 0x49, 0x6c, 0x60,
	0x8e, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xb1, 0xfe, 0xaa, 0xb1, 0x1c, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// FlagClient is the client API for Flag service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FlagClient interface {
	FlagGet(ctx context.Context, in *FlagGetRequest, opts ...grpc.CallOption) (*FlagResponse, error)
	FlagSet(ctx context.Context, in *FlagSetRequest, opts ...grpc.CallOption) (*Empty, error)
}

type flagClient struct {
	cc *grpc.ClientConn
}

func NewFlagClient(cc *grpc.ClientConn) FlagClient {
	return &flagClient{cc}
}

func (c *flagClient) FlagGet(ctx context.Context, in *FlagGetRequest, opts ...grpc.CallOption) (*FlagResponse, error) {
	out := new(FlagResponse)
	err := c.cc.Invoke(ctx, "/proto.Flag/FlagGet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *flagClient) FlagSet(ctx context.Context, in *FlagSetRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/proto.Flag/FlagSet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FlagServer is the server API for Flag service.
type FlagServer interface {
	FlagGet(context.Context, *FlagGetRequest) (*FlagResponse, error)
	FlagSet(context.Context, *FlagSetRequest) (*Empty, error)
}

func RegisterFlagServer(s *grpc.Server, srv FlagServer) {
	s.RegisterService(&_Flag_serviceDesc, srv)
}

func _Flag_FlagGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FlagGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlagServer).FlagGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Flag/FlagGet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlagServer).FlagGet(ctx, req.(*FlagGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Flag_FlagSet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FlagSetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlagServer).FlagSet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Flag/FlagSet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlagServer).FlagSet(ctx, req.(*FlagSetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Flag_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Flag",
	HandlerType: (*FlagServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FlagGet",
			Handler:    _Flag_FlagGet_Handler,
		},
		{
			MethodName: "FlagSet",
			Handler:    _Flag_FlagSet_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "flag.proto",
}
