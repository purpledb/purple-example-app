// Code generated by protoc-gen-go. DO NOT EDIT.
// source: counter.proto

package proto

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type IncrementCounterRequest struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Amount               int64    `protobuf:"varint,2,opt,name=amount,proto3" json:"amount,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IncrementCounterRequest) Reset()         { *m = IncrementCounterRequest{} }
func (m *IncrementCounterRequest) String() string { return proto.CompactTextString(m) }
func (*IncrementCounterRequest) ProtoMessage()    {}
func (*IncrementCounterRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_75dcd656fce7132f, []int{0}
}

func (m *IncrementCounterRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IncrementCounterRequest.Unmarshal(m, b)
}
func (m *IncrementCounterRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IncrementCounterRequest.Marshal(b, m, deterministic)
}
func (m *IncrementCounterRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IncrementCounterRequest.Merge(m, src)
}
func (m *IncrementCounterRequest) XXX_Size() int {
	return xxx_messageInfo_IncrementCounterRequest.Size(m)
}
func (m *IncrementCounterRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_IncrementCounterRequest.DiscardUnknown(m)
}

var xxx_messageInfo_IncrementCounterRequest proto.InternalMessageInfo

func (m *IncrementCounterRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *IncrementCounterRequest) GetAmount() int64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

type GetCounterRequest struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetCounterRequest) Reset()         { *m = GetCounterRequest{} }
func (m *GetCounterRequest) String() string { return proto.CompactTextString(m) }
func (*GetCounterRequest) ProtoMessage()    {}
func (*GetCounterRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_75dcd656fce7132f, []int{1}
}

func (m *GetCounterRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCounterRequest.Unmarshal(m, b)
}
func (m *GetCounterRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCounterRequest.Marshal(b, m, deterministic)
}
func (m *GetCounterRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCounterRequest.Merge(m, src)
}
func (m *GetCounterRequest) XXX_Size() int {
	return xxx_messageInfo_GetCounterRequest.Size(m)
}
func (m *GetCounterRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCounterRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetCounterRequest proto.InternalMessageInfo

func (m *GetCounterRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type GetCounterResponse struct {
	Value                int64    `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetCounterResponse) Reset()         { *m = GetCounterResponse{} }
func (m *GetCounterResponse) String() string { return proto.CompactTextString(m) }
func (*GetCounterResponse) ProtoMessage()    {}
func (*GetCounterResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_75dcd656fce7132f, []int{2}
}

func (m *GetCounterResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCounterResponse.Unmarshal(m, b)
}
func (m *GetCounterResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCounterResponse.Marshal(b, m, deterministic)
}
func (m *GetCounterResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCounterResponse.Merge(m, src)
}
func (m *GetCounterResponse) XXX_Size() int {
	return xxx_messageInfo_GetCounterResponse.Size(m)
}
func (m *GetCounterResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCounterResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetCounterResponse proto.InternalMessageInfo

func (m *GetCounterResponse) GetValue() int64 {
	if m != nil {
		return m.Value
	}
	return 0
}

func init() {
	proto.RegisterType((*IncrementCounterRequest)(nil), "proto.IncrementCounterRequest")
	proto.RegisterType((*GetCounterRequest)(nil), "proto.GetCounterRequest")
	proto.RegisterType((*GetCounterResponse)(nil), "proto.GetCounterResponse")
}

func init() { proto.RegisterFile("counter.proto", fileDescriptor_75dcd656fce7132f) }

var fileDescriptor_75dcd656fce7132f = []byte{
	// 197 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4d, 0xce, 0x2f, 0xcd,
	0x2b, 0x49, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x53, 0x52, 0x3c, 0xc9,
	0xf9, 0xb9, 0xb9, 0xf9, 0x79, 0x10, 0x41, 0x25, 0x67, 0x2e, 0x71, 0xcf, 0xbc, 0xe4, 0xa2, 0xd4,
	0xdc, 0xd4, 0xbc, 0x12, 0x67, 0x88, 0xf2, 0xa0, 0xd4, 0xc2, 0xd2, 0xd4, 0xe2, 0x12, 0x21, 0x01,
	0x2e, 0xe6, 0xec, 0xd4, 0x4a, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x10, 0x53, 0x48, 0x8c,
	0x8b, 0x2d, 0x31, 0x17, 0xa4, 0x48, 0x82, 0x49, 0x81, 0x51, 0x83, 0x39, 0x08, 0xca, 0x53, 0x52,
	0xe5, 0x12, 0x74, 0x4f, 0x25, 0xa8, 0x5d, 0x49, 0x8b, 0x4b, 0x08, 0x59, 0x59, 0x71, 0x41, 0x7e,
	0x5e, 0x71, 0xaa, 0x90, 0x08, 0x17, 0x6b, 0x59, 0x62, 0x4e, 0x69, 0x2a, 0x58, 0x25, 0x73, 0x10,
	0x84, 0x63, 0xd4, 0xc7, 0xc8, 0xc5, 0x0e, 0x55, 0x29, 0xe4, 0xc0, 0x25, 0x80, 0xee, 0x46, 0x21,
	0x39, 0x88, 0xfb, 0xf5, 0x70, 0x38, 0x5e, 0x8a, 0x07, 0x2a, 0xef, 0x9a, 0x5b, 0x50, 0x52, 0x29,
	0xe4, 0xc8, 0xc5, 0x85, 0xb0, 0x59, 0x48, 0x02, 0x2a, 0x87, 0xe1, 0x66, 0x29, 0x49, 0x2c, 0x32,
	0x10, 0x67, 0x26, 0xb1, 0x81, 0x65, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x73, 0x49, 0x74,
	0x7d, 0x55, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CounterClient is the client API for Counter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CounterClient interface {
	IncrementCounter(ctx context.Context, in *IncrementCounterRequest, opts ...grpc.CallOption) (*Empty, error)
	GetCounter(ctx context.Context, in *GetCounterRequest, opts ...grpc.CallOption) (*GetCounterResponse, error)
}

type counterClient struct {
	cc *grpc.ClientConn
}

func NewCounterClient(cc *grpc.ClientConn) CounterClient {
	return &counterClient{cc}
}

func (c *counterClient) IncrementCounter(ctx context.Context, in *IncrementCounterRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/proto.Counter/IncrementCounter", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *counterClient) GetCounter(ctx context.Context, in *GetCounterRequest, opts ...grpc.CallOption) (*GetCounterResponse, error) {
	out := new(GetCounterResponse)
	err := c.cc.Invoke(ctx, "/proto.Counter/GetCounter", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CounterServer is the server API for Counter service.
type CounterServer interface {
	IncrementCounter(context.Context, *IncrementCounterRequest) (*Empty, error)
	GetCounter(context.Context, *GetCounterRequest) (*GetCounterResponse, error)
}

func RegisterCounterServer(s *grpc.Server, srv CounterServer) {
	s.RegisterService(&_Counter_serviceDesc, srv)
}

func _Counter_IncrementCounter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IncrementCounterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CounterServer).IncrementCounter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Counter/IncrementCounter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CounterServer).IncrementCounter(ctx, req.(*IncrementCounterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Counter_GetCounter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCounterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CounterServer).GetCounter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Counter/GetCounter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CounterServer).GetCounter(ctx, req.(*GetCounterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Counter_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Counter",
	HandlerType: (*CounterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "IncrementCounter",
			Handler:    _Counter_IncrementCounter_Handler,
		},
		{
			MethodName: "GetCounter",
			Handler:    _Counter_GetCounter_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "counter.proto",
}
