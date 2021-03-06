// Code generated by protoc-gen-go. DO NOT EDIT.
// source: orchestration.proto

package logcache_v1

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

type Range struct {
	// start is the first hash within the given range. [start..end]
	Start uint64 `protobuf:"varint,1,opt,name=start,proto3" json:"start,omitempty"`
	// end is the last hash within the given range. [start..end]
	End                  uint64   `protobuf:"varint,2,opt,name=end,proto3" json:"end,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Range) Reset()         { *m = Range{} }
func (m *Range) String() string { return proto.CompactTextString(m) }
func (*Range) ProtoMessage()    {}
func (*Range) Descriptor() ([]byte, []int) {
	return fileDescriptor_orchestration_08848b63d402f2dc, []int{0}
}
func (m *Range) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Range.Unmarshal(m, b)
}
func (m *Range) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Range.Marshal(b, m, deterministic)
}
func (dst *Range) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Range.Merge(dst, src)
}
func (m *Range) XXX_Size() int {
	return xxx_messageInfo_Range.Size(m)
}
func (m *Range) XXX_DiscardUnknown() {
	xxx_messageInfo_Range.DiscardUnknown(m)
}

var xxx_messageInfo_Range proto.InternalMessageInfo

func (m *Range) GetStart() uint64 {
	if m != nil {
		return m.Start
	}
	return 0
}

func (m *Range) GetEnd() uint64 {
	if m != nil {
		return m.End
	}
	return 0
}

type Ranges struct {
	Ranges               []*Range `protobuf:"bytes,1,rep,name=ranges,proto3" json:"ranges,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ranges) Reset()         { *m = Ranges{} }
func (m *Ranges) String() string { return proto.CompactTextString(m) }
func (*Ranges) ProtoMessage()    {}
func (*Ranges) Descriptor() ([]byte, []int) {
	return fileDescriptor_orchestration_08848b63d402f2dc, []int{1}
}
func (m *Ranges) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ranges.Unmarshal(m, b)
}
func (m *Ranges) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ranges.Marshal(b, m, deterministic)
}
func (dst *Ranges) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ranges.Merge(dst, src)
}
func (m *Ranges) XXX_Size() int {
	return xxx_messageInfo_Ranges.Size(m)
}
func (m *Ranges) XXX_DiscardUnknown() {
	xxx_messageInfo_Ranges.DiscardUnknown(m)
}

var xxx_messageInfo_Ranges proto.InternalMessageInfo

func (m *Ranges) GetRanges() []*Range {
	if m != nil {
		return m.Ranges
	}
	return nil
}

type AddRangeRequest struct {
	Range                *Range   `protobuf:"bytes,1,opt,name=range,proto3" json:"range,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddRangeRequest) Reset()         { *m = AddRangeRequest{} }
func (m *AddRangeRequest) String() string { return proto.CompactTextString(m) }
func (*AddRangeRequest) ProtoMessage()    {}
func (*AddRangeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_orchestration_08848b63d402f2dc, []int{2}
}
func (m *AddRangeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddRangeRequest.Unmarshal(m, b)
}
func (m *AddRangeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddRangeRequest.Marshal(b, m, deterministic)
}
func (dst *AddRangeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddRangeRequest.Merge(dst, src)
}
func (m *AddRangeRequest) XXX_Size() int {
	return xxx_messageInfo_AddRangeRequest.Size(m)
}
func (m *AddRangeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddRangeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddRangeRequest proto.InternalMessageInfo

func (m *AddRangeRequest) GetRange() *Range {
	if m != nil {
		return m.Range
	}
	return nil
}

type AddRangeResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddRangeResponse) Reset()         { *m = AddRangeResponse{} }
func (m *AddRangeResponse) String() string { return proto.CompactTextString(m) }
func (*AddRangeResponse) ProtoMessage()    {}
func (*AddRangeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_orchestration_08848b63d402f2dc, []int{3}
}
func (m *AddRangeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddRangeResponse.Unmarshal(m, b)
}
func (m *AddRangeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddRangeResponse.Marshal(b, m, deterministic)
}
func (dst *AddRangeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddRangeResponse.Merge(dst, src)
}
func (m *AddRangeResponse) XXX_Size() int {
	return xxx_messageInfo_AddRangeResponse.Size(m)
}
func (m *AddRangeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AddRangeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AddRangeResponse proto.InternalMessageInfo

type RemoveRangeRequest struct {
	Range                *Range   `protobuf:"bytes,1,opt,name=range,proto3" json:"range,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RemoveRangeRequest) Reset()         { *m = RemoveRangeRequest{} }
func (m *RemoveRangeRequest) String() string { return proto.CompactTextString(m) }
func (*RemoveRangeRequest) ProtoMessage()    {}
func (*RemoveRangeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_orchestration_08848b63d402f2dc, []int{4}
}
func (m *RemoveRangeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoveRangeRequest.Unmarshal(m, b)
}
func (m *RemoveRangeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoveRangeRequest.Marshal(b, m, deterministic)
}
func (dst *RemoveRangeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoveRangeRequest.Merge(dst, src)
}
func (m *RemoveRangeRequest) XXX_Size() int {
	return xxx_messageInfo_RemoveRangeRequest.Size(m)
}
func (m *RemoveRangeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoveRangeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RemoveRangeRequest proto.InternalMessageInfo

func (m *RemoveRangeRequest) GetRange() *Range {
	if m != nil {
		return m.Range
	}
	return nil
}

type RemoveRangeResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RemoveRangeResponse) Reset()         { *m = RemoveRangeResponse{} }
func (m *RemoveRangeResponse) String() string { return proto.CompactTextString(m) }
func (*RemoveRangeResponse) ProtoMessage()    {}
func (*RemoveRangeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_orchestration_08848b63d402f2dc, []int{5}
}
func (m *RemoveRangeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoveRangeResponse.Unmarshal(m, b)
}
func (m *RemoveRangeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoveRangeResponse.Marshal(b, m, deterministic)
}
func (dst *RemoveRangeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoveRangeResponse.Merge(dst, src)
}
func (m *RemoveRangeResponse) XXX_Size() int {
	return xxx_messageInfo_RemoveRangeResponse.Size(m)
}
func (m *RemoveRangeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoveRangeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RemoveRangeResponse proto.InternalMessageInfo

type ListRangesRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListRangesRequest) Reset()         { *m = ListRangesRequest{} }
func (m *ListRangesRequest) String() string { return proto.CompactTextString(m) }
func (*ListRangesRequest) ProtoMessage()    {}
func (*ListRangesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_orchestration_08848b63d402f2dc, []int{6}
}
func (m *ListRangesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListRangesRequest.Unmarshal(m, b)
}
func (m *ListRangesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListRangesRequest.Marshal(b, m, deterministic)
}
func (dst *ListRangesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListRangesRequest.Merge(dst, src)
}
func (m *ListRangesRequest) XXX_Size() int {
	return xxx_messageInfo_ListRangesRequest.Size(m)
}
func (m *ListRangesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListRangesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListRangesRequest proto.InternalMessageInfo

type ListRangesResponse struct {
	Ranges               []*Range `protobuf:"bytes,1,rep,name=ranges,proto3" json:"ranges,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListRangesResponse) Reset()         { *m = ListRangesResponse{} }
func (m *ListRangesResponse) String() string { return proto.CompactTextString(m) }
func (*ListRangesResponse) ProtoMessage()    {}
func (*ListRangesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_orchestration_08848b63d402f2dc, []int{7}
}
func (m *ListRangesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListRangesResponse.Unmarshal(m, b)
}
func (m *ListRangesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListRangesResponse.Marshal(b, m, deterministic)
}
func (dst *ListRangesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListRangesResponse.Merge(dst, src)
}
func (m *ListRangesResponse) XXX_Size() int {
	return xxx_messageInfo_ListRangesResponse.Size(m)
}
func (m *ListRangesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListRangesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListRangesResponse proto.InternalMessageInfo

func (m *ListRangesResponse) GetRanges() []*Range {
	if m != nil {
		return m.Ranges
	}
	return nil
}

type SetRangesRequest struct {
	// The key is the address of the Log Cache node.
	Ranges               map[string]*Ranges `protobuf:"bytes,1,rep,name=ranges,proto3" json:"ranges,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *SetRangesRequest) Reset()         { *m = SetRangesRequest{} }
func (m *SetRangesRequest) String() string { return proto.CompactTextString(m) }
func (*SetRangesRequest) ProtoMessage()    {}
func (*SetRangesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_orchestration_08848b63d402f2dc, []int{8}
}
func (m *SetRangesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetRangesRequest.Unmarshal(m, b)
}
func (m *SetRangesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetRangesRequest.Marshal(b, m, deterministic)
}
func (dst *SetRangesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetRangesRequest.Merge(dst, src)
}
func (m *SetRangesRequest) XXX_Size() int {
	return xxx_messageInfo_SetRangesRequest.Size(m)
}
func (m *SetRangesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SetRangesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SetRangesRequest proto.InternalMessageInfo

func (m *SetRangesRequest) GetRanges() map[string]*Ranges {
	if m != nil {
		return m.Ranges
	}
	return nil
}

type SetRangesResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetRangesResponse) Reset()         { *m = SetRangesResponse{} }
func (m *SetRangesResponse) String() string { return proto.CompactTextString(m) }
func (*SetRangesResponse) ProtoMessage()    {}
func (*SetRangesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_orchestration_08848b63d402f2dc, []int{9}
}
func (m *SetRangesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetRangesResponse.Unmarshal(m, b)
}
func (m *SetRangesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetRangesResponse.Marshal(b, m, deterministic)
}
func (dst *SetRangesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetRangesResponse.Merge(dst, src)
}
func (m *SetRangesResponse) XXX_Size() int {
	return xxx_messageInfo_SetRangesResponse.Size(m)
}
func (m *SetRangesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SetRangesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SetRangesResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Range)(nil), "logcache.v1.Range")
	proto.RegisterType((*Ranges)(nil), "logcache.v1.Ranges")
	proto.RegisterType((*AddRangeRequest)(nil), "logcache.v1.AddRangeRequest")
	proto.RegisterType((*AddRangeResponse)(nil), "logcache.v1.AddRangeResponse")
	proto.RegisterType((*RemoveRangeRequest)(nil), "logcache.v1.RemoveRangeRequest")
	proto.RegisterType((*RemoveRangeResponse)(nil), "logcache.v1.RemoveRangeResponse")
	proto.RegisterType((*ListRangesRequest)(nil), "logcache.v1.ListRangesRequest")
	proto.RegisterType((*ListRangesResponse)(nil), "logcache.v1.ListRangesResponse")
	proto.RegisterType((*SetRangesRequest)(nil), "logcache.v1.SetRangesRequest")
	proto.RegisterMapType((map[string]*Ranges)(nil), "logcache.v1.SetRangesRequest.RangesEntry")
	proto.RegisterType((*SetRangesResponse)(nil), "logcache.v1.SetRangesResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// OrchestrationClient is the client API for Orchestration service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type OrchestrationClient interface {
	AddRange(ctx context.Context, in *AddRangeRequest, opts ...grpc.CallOption) (*AddRangeResponse, error)
	RemoveRange(ctx context.Context, in *RemoveRangeRequest, opts ...grpc.CallOption) (*RemoveRangeResponse, error)
	ListRanges(ctx context.Context, in *ListRangesRequest, opts ...grpc.CallOption) (*ListRangesResponse, error)
	SetRanges(ctx context.Context, in *SetRangesRequest, opts ...grpc.CallOption) (*SetRangesResponse, error)
}

type orchestrationClient struct {
	cc *grpc.ClientConn
}

func NewOrchestrationClient(cc *grpc.ClientConn) OrchestrationClient {
	return &orchestrationClient{cc}
}

func (c *orchestrationClient) AddRange(ctx context.Context, in *AddRangeRequest, opts ...grpc.CallOption) (*AddRangeResponse, error) {
	out := new(AddRangeResponse)
	err := c.cc.Invoke(ctx, "/logcache.v1.Orchestration/AddRange", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orchestrationClient) RemoveRange(ctx context.Context, in *RemoveRangeRequest, opts ...grpc.CallOption) (*RemoveRangeResponse, error) {
	out := new(RemoveRangeResponse)
	err := c.cc.Invoke(ctx, "/logcache.v1.Orchestration/RemoveRange", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orchestrationClient) ListRanges(ctx context.Context, in *ListRangesRequest, opts ...grpc.CallOption) (*ListRangesResponse, error) {
	out := new(ListRangesResponse)
	err := c.cc.Invoke(ctx, "/logcache.v1.Orchestration/ListRanges", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orchestrationClient) SetRanges(ctx context.Context, in *SetRangesRequest, opts ...grpc.CallOption) (*SetRangesResponse, error) {
	out := new(SetRangesResponse)
	err := c.cc.Invoke(ctx, "/logcache.v1.Orchestration/SetRanges", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrchestrationServer is the server API for Orchestration service.
type OrchestrationServer interface {
	AddRange(context.Context, *AddRangeRequest) (*AddRangeResponse, error)
	RemoveRange(context.Context, *RemoveRangeRequest) (*RemoveRangeResponse, error)
	ListRanges(context.Context, *ListRangesRequest) (*ListRangesResponse, error)
	SetRanges(context.Context, *SetRangesRequest) (*SetRangesResponse, error)
}

func RegisterOrchestrationServer(s *grpc.Server, srv OrchestrationServer) {
	s.RegisterService(&_Orchestration_serviceDesc, srv)
}

func _Orchestration_AddRange_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddRangeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrchestrationServer).AddRange(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/logcache.v1.Orchestration/AddRange",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrchestrationServer).AddRange(ctx, req.(*AddRangeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Orchestration_RemoveRange_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveRangeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrchestrationServer).RemoveRange(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/logcache.v1.Orchestration/RemoveRange",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrchestrationServer).RemoveRange(ctx, req.(*RemoveRangeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Orchestration_ListRanges_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRangesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrchestrationServer).ListRanges(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/logcache.v1.Orchestration/ListRanges",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrchestrationServer).ListRanges(ctx, req.(*ListRangesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Orchestration_SetRanges_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetRangesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrchestrationServer).SetRanges(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/logcache.v1.Orchestration/SetRanges",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrchestrationServer).SetRanges(ctx, req.(*SetRangesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Orchestration_serviceDesc = grpc.ServiceDesc{
	ServiceName: "logcache.v1.Orchestration",
	HandlerType: (*OrchestrationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddRange",
			Handler:    _Orchestration_AddRange_Handler,
		},
		{
			MethodName: "RemoveRange",
			Handler:    _Orchestration_RemoveRange_Handler,
		},
		{
			MethodName: "ListRanges",
			Handler:    _Orchestration_ListRanges_Handler,
		},
		{
			MethodName: "SetRanges",
			Handler:    _Orchestration_SetRanges_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "orchestration.proto",
}

func init() { proto.RegisterFile("orchestration.proto", fileDescriptor_orchestration_08848b63d402f2dc) }

var fileDescriptor_orchestration_08848b63d402f2dc = []byte{
	// 362 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0xcf, 0x4e, 0xc2, 0x40,
	0x10, 0xc6, 0x2d, 0x58, 0x22, 0xd3, 0x18, 0x71, 0xaa, 0x09, 0x69, 0xe4, 0x4f, 0x7a, 0x02, 0x0f,
	0x35, 0xa2, 0x07, 0xa3, 0x89, 0x91, 0x83, 0x07, 0x13, 0x22, 0xc9, 0xfa, 0x04, 0x15, 0x26, 0x40,
	0xc4, 0x2e, 0x76, 0x17, 0x12, 0x1e, 0xc8, 0xc7, 0xf1, 0x9d, 0x4c, 0x77, 0x0b, 0x74, 0xc1, 0x9a,
	0xe8, 0x6d, 0x98, 0xf9, 0xe6, 0xb7, 0xfb, 0xed, 0x47, 0xc1, 0xe5, 0xf1, 0x60, 0x4c, 0x42, 0xc6,
	0xa1, 0x9c, 0xf0, 0x28, 0x98, 0xc5, 0x5c, 0x72, 0x74, 0xa6, 0x7c, 0x34, 0x08, 0x07, 0x63, 0x0a,
	0x16, 0x97, 0xfe, 0x05, 0xd8, 0x2c, 0x8c, 0x46, 0x84, 0x27, 0x60, 0x0b, 0x19, 0xc6, 0xb2, 0x6a,
	0x35, 0xad, 0xd6, 0x3e, 0xd3, 0x3f, 0xb0, 0x02, 0x45, 0x8a, 0x86, 0xd5, 0x82, 0xea, 0x25, 0xa5,
	0x7f, 0x0d, 0x25, 0xb5, 0x20, 0xf0, 0x1c, 0x4a, 0xb1, 0xaa, 0xaa, 0x56, 0xb3, 0xd8, 0x72, 0x3a,
	0x18, 0x64, 0xc0, 0x81, 0x12, 0xb1, 0x54, 0xe1, 0xdf, 0xc1, 0x51, 0x77, 0x38, 0xd4, 0x3d, 0xfa,
	0x98, 0x93, 0x90, 0xd8, 0x02, 0x5b, 0x0d, 0xd5, 0x81, 0x3f, 0x6f, 0x6b, 0x81, 0x8f, 0x50, 0xd9,
	0x2c, 0x8b, 0x19, 0x8f, 0x04, 0xf9, 0xf7, 0x80, 0x8c, 0xde, 0xf9, 0x82, 0xfe, 0xc9, 0x3c, 0x05,
	0xd7, 0xd8, 0x4f, 0xb1, 0x2e, 0x1c, 0xf7, 0x26, 0x42, 0x6a, 0x87, 0x29, 0xd5, 0x7f, 0x00, 0xcc,
	0x36, 0xb5, 0xf4, 0x4f, 0xf6, 0x3f, 0x2d, 0xa8, 0xbc, 0x90, 0x89, 0xc5, 0xee, 0x16, 0xa0, 0x6d,
	0x00, 0xb6, 0xe5, 0x9a, 0x28, 0x1e, 0x23, 0x19, 0x2f, 0x57, 0x5c, 0xef, 0x19, 0x9c, 0x4c, 0x3b,
	0x49, 0xeb, 0x8d, 0x96, 0xca, 0x7c, 0x99, 0x25, 0x25, 0xb6, 0xc1, 0x5e, 0x84, 0xd3, 0x39, 0xa9,
	0x04, 0x9d, 0x8e, 0xbb, 0x7b, 0x47, 0xc1, 0xb4, 0xe2, 0xb6, 0x70, 0x63, 0x25, 0xf6, 0x33, 0xe7,
	0x6a, 0xa3, 0x9d, 0xaf, 0x02, 0x1c, 0xf6, 0xb3, 0xff, 0x23, 0x7c, 0x82, 0x83, 0x55, 0x20, 0x78,
	0x66, 0x20, 0xb7, 0x42, 0xf6, 0x6a, 0x39, 0xd3, 0xf4, 0xb9, 0xf7, 0x90, 0x81, 0x93, 0xc9, 0x01,
	0x1b, 0xe6, 0x05, 0x77, 0x12, 0xf6, 0x9a, 0xf9, 0x82, 0x35, 0xb3, 0x0f, 0xb0, 0xc9, 0x0b, 0xeb,
	0xc6, 0xc6, 0x4e, 0xba, 0x5e, 0x23, 0x77, 0xbe, 0x06, 0xf6, 0xa0, 0xbc, 0x7e, 0x16, 0xac, 0xfd,
	0x1a, 0x93, 0x57, 0xcf, 0x1b, 0xaf, 0x68, 0xaf, 0x25, 0xf5, 0x19, 0x5e, 0x7d, 0x07, 0x00, 0x00,
	0xff, 0xff, 0x68, 0x15, 0x04, 0x7c, 0x9d, 0x03, 0x00, 0x00,
}
