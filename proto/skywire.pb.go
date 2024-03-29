// Code generated by protoc-gen-go. DO NOT EDIT.
// source: skywire.proto

package skywire

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

type Member struct {
	Type                 string   `protobuf:"bytes,10,opt,name=type,proto3" json:"type,omitempty"`
	Name                 string   `protobuf:"bytes,20,opt,name=name,proto3" json:"name,omitempty"`
	Endian               uint32   `protobuf:"varint,30,opt,name=endian,proto3" json:"endian,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Member) Reset()         { *m = Member{} }
func (m *Member) String() string { return proto.CompactTextString(m) }
func (*Member) ProtoMessage()    {}
func (*Member) Descriptor() ([]byte, []int) {
	return fileDescriptor_f3a41043b04a0f40, []int{0}
}

func (m *Member) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Member.Unmarshal(m, b)
}
func (m *Member) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Member.Marshal(b, m, deterministic)
}
func (m *Member) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Member.Merge(m, src)
}
func (m *Member) XXX_Size() int {
	return xxx_messageInfo_Member.Size(m)
}
func (m *Member) XXX_DiscardUnknown() {
	xxx_messageInfo_Member.DiscardUnknown(m)
}

var xxx_messageInfo_Member proto.InternalMessageInfo

func (m *Member) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Member) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Member) GetEndian() uint32 {
	if m != nil {
		return m.Endian
	}
	return 0
}

type Stencil struct {
	StringValues         map[string]string  `protobuf:"bytes,10,rep,name=string_values,json=stringValues,proto3" json:"string_values,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Members              map[uint32]*Member `protobuf:"bytes,40,rep,name=members,proto3" json:"members,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Includes             []string           `protobuf:"bytes,100,rep,name=includes,proto3" json:"includes,omitempty"`
	NativeTypes          []string           `protobuf:"bytes,101,rep,name=native_types,json=nativeTypes,proto3" json:"native_types,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Stencil) Reset()         { *m = Stencil{} }
func (m *Stencil) String() string { return proto.CompactTextString(m) }
func (*Stencil) ProtoMessage()    {}
func (*Stencil) Descriptor() ([]byte, []int) {
	return fileDescriptor_f3a41043b04a0f40, []int{1}
}

func (m *Stencil) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Stencil.Unmarshal(m, b)
}
func (m *Stencil) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Stencil.Marshal(b, m, deterministic)
}
func (m *Stencil) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Stencil.Merge(m, src)
}
func (m *Stencil) XXX_Size() int {
	return xxx_messageInfo_Stencil.Size(m)
}
func (m *Stencil) XXX_DiscardUnknown() {
	xxx_messageInfo_Stencil.DiscardUnknown(m)
}

var xxx_messageInfo_Stencil proto.InternalMessageInfo

func (m *Stencil) GetStringValues() map[string]string {
	if m != nil {
		return m.StringValues
	}
	return nil
}

func (m *Stencil) GetMembers() map[uint32]*Member {
	if m != nil {
		return m.Members
	}
	return nil
}

func (m *Stencil) GetIncludes() []string {
	if m != nil {
		return m.Includes
	}
	return nil
}

func (m *Stencil) GetNativeTypes() []string {
	if m != nil {
		return m.NativeTypes
	}
	return nil
}

type RequestMetadata struct {
	Id                   uint64   `protobuf:"varint,10,opt,name=id,proto3" json:"id,omitempty"`
	GenerateTests        bool     `protobuf:"varint,20,opt,name=generate_tests,json=generateTests,proto3" json:"generate_tests,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestMetadata) Reset()         { *m = RequestMetadata{} }
func (m *RequestMetadata) String() string { return proto.CompactTextString(m) }
func (*RequestMetadata) ProtoMessage()    {}
func (*RequestMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_f3a41043b04a0f40, []int{2}
}

func (m *RequestMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestMetadata.Unmarshal(m, b)
}
func (m *RequestMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestMetadata.Marshal(b, m, deterministic)
}
func (m *RequestMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestMetadata.Merge(m, src)
}
func (m *RequestMetadata) XXX_Size() int {
	return xxx_messageInfo_RequestMetadata.Size(m)
}
func (m *RequestMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_RequestMetadata proto.InternalMessageInfo

func (m *RequestMetadata) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *RequestMetadata) GetGenerateTests() bool {
	if m != nil {
		return m.GenerateTests
	}
	return false
}

type SkywireRequest struct {
	Meta                 *RequestMetadata `protobuf:"bytes,10,opt,name=meta,proto3" json:"meta,omitempty"`
	MsgSpec              *Stencil         `protobuf:"bytes,100,opt,name=msg_spec,json=msgSpec,proto3" json:"msg_spec,omitempty"`
	LangSpec             *Stencil         `protobuf:"bytes,110,opt,name=lang_spec,json=langSpec,proto3" json:"lang_spec,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *SkywireRequest) Reset()         { *m = SkywireRequest{} }
func (m *SkywireRequest) String() string { return proto.CompactTextString(m) }
func (*SkywireRequest) ProtoMessage()    {}
func (*SkywireRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f3a41043b04a0f40, []int{3}
}

func (m *SkywireRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SkywireRequest.Unmarshal(m, b)
}
func (m *SkywireRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SkywireRequest.Marshal(b, m, deterministic)
}
func (m *SkywireRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SkywireRequest.Merge(m, src)
}
func (m *SkywireRequest) XXX_Size() int {
	return xxx_messageInfo_SkywireRequest.Size(m)
}
func (m *SkywireRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SkywireRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SkywireRequest proto.InternalMessageInfo

func (m *SkywireRequest) GetMeta() *RequestMetadata {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *SkywireRequest) GetMsgSpec() *Stencil {
	if m != nil {
		return m.MsgSpec
	}
	return nil
}

func (m *SkywireRequest) GetLangSpec() *Stencil {
	if m != nil {
		return m.LangSpec
	}
	return nil
}

type SkywireResponse struct {
	Meta                 *RequestMetadata `protobuf:"bytes,10,opt,name=meta,proto3" json:"meta,omitempty"`
	Impl                 string           `protobuf:"bytes,100,opt,name=impl,proto3" json:"impl,omitempty"`
	Header               string           `protobuf:"bytes,120,opt,name=header,proto3" json:"header,omitempty"`
	Test                 string           `protobuf:"bytes,130,opt,name=test,proto3" json:"test,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *SkywireResponse) Reset()         { *m = SkywireResponse{} }
func (m *SkywireResponse) String() string { return proto.CompactTextString(m) }
func (*SkywireResponse) ProtoMessage()    {}
func (*SkywireResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f3a41043b04a0f40, []int{4}
}

func (m *SkywireResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SkywireResponse.Unmarshal(m, b)
}
func (m *SkywireResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SkywireResponse.Marshal(b, m, deterministic)
}
func (m *SkywireResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SkywireResponse.Merge(m, src)
}
func (m *SkywireResponse) XXX_Size() int {
	return xxx_messageInfo_SkywireResponse.Size(m)
}
func (m *SkywireResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SkywireResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SkywireResponse proto.InternalMessageInfo

func (m *SkywireResponse) GetMeta() *RequestMetadata {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *SkywireResponse) GetImpl() string {
	if m != nil {
		return m.Impl
	}
	return ""
}

func (m *SkywireResponse) GetHeader() string {
	if m != nil {
		return m.Header
	}
	return ""
}

func (m *SkywireResponse) GetTest() string {
	if m != nil {
		return m.Test
	}
	return ""
}

func init() {
	proto.RegisterType((*Member)(nil), "skywire.Member")
	proto.RegisterType((*Stencil)(nil), "skywire.Stencil")
	proto.RegisterMapType((map[uint32]*Member)(nil), "skywire.Stencil.MembersEntry")
	proto.RegisterMapType((map[string]string)(nil), "skywire.Stencil.StringValuesEntry")
	proto.RegisterType((*RequestMetadata)(nil), "skywire.RequestMetadata")
	proto.RegisterType((*SkywireRequest)(nil), "skywire.SkywireRequest")
	proto.RegisterType((*SkywireResponse)(nil), "skywire.SkywireResponse")
}

func init() { proto.RegisterFile("skywire.proto", fileDescriptor_f3a41043b04a0f40) }

var fileDescriptor_f3a41043b04a0f40 = []byte{
	// 470 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x86, 0xb1, 0x63, 0x62, 0x67, 0x12, 0x27, 0x65, 0xa8, 0x60, 0x15, 0x09, 0x14, 0x2c, 0x55,
	0xb2, 0x04, 0xe4, 0x10, 0x0e, 0x20, 0x2e, 0x3d, 0xa1, 0x56, 0xa0, 0x5e, 0x36, 0x15, 0xd7, 0x68,
	0x1b, 0x8f, 0x82, 0x55, 0x7b, 0x63, 0xbc, 0x9b, 0x42, 0x6e, 0x88, 0xc7, 0xe0, 0xc9, 0x78, 0x1c,
	0xb4, 0xeb, 0x4d, 0x30, 0x89, 0x38, 0xf4, 0x36, 0xf3, 0xfb, 0x9f, 0xd9, 0x6f, 0x67, 0xd6, 0x10,
	0xab, 0xdb, 0xed, 0xb7, 0xbc, 0xa6, 0x69, 0x55, 0xaf, 0xf5, 0x1a, 0x43, 0x97, 0x26, 0x97, 0xd0,
	0xbd, 0xa2, 0xf2, 0x86, 0x6a, 0x44, 0x08, 0xf4, 0xb6, 0x22, 0x06, 0x13, 0x2f, 0xed, 0x71, 0x1b,
	0x1b, 0x4d, 0x8a, 0x92, 0xd8, 0x69, 0xa3, 0x99, 0x18, 0x9f, 0x40, 0x97, 0x64, 0x96, 0x0b, 0xc9,
	0x9e, 0x4f, 0xbc, 0x34, 0xe6, 0x2e, 0x4b, 0x7e, 0xfb, 0x10, 0xce, 0x35, 0xc9, 0x65, 0x5e, 0xe0,
	0x05, 0xc4, 0x4a, 0xd7, 0xb9, 0x5c, 0x2d, 0xee, 0x44, 0xb1, 0x21, 0xc5, 0x60, 0xd2, 0x49, 0xfb,
	0xb3, 0x64, 0xba, 0xa3, 0x70, 0xc6, 0xe9, 0xdc, 0xba, 0x3e, 0x5b, 0xd3, 0x07, 0xa9, 0xeb, 0x2d,
	0x1f, 0xa8, 0x96, 0x84, 0x6f, 0x21, 0x2c, 0x2d, 0x9e, 0x62, 0xa9, 0x6d, 0xf1, 0xec, 0xa8, 0x45,
	0x83, 0xef, 0xaa, 0x77, 0x6e, 0x1c, 0x43, 0x94, 0xcb, 0x65, 0xb1, 0xc9, 0x48, 0xb1, 0x6c, 0xd2,
	0x49, 0x7b, 0x7c, 0x9f, 0xe3, 0x0b, 0x18, 0x48, 0xa1, 0xf3, 0x3b, 0x5a, 0x98, 0x4b, 0x2a, 0x46,
	0xf6, 0x7b, 0xbf, 0xd1, 0xae, 0x8d, 0x34, 0x3e, 0x87, 0x47, 0x47, 0x68, 0x78, 0x02, 0x9d, 0x5b,
	0xda, 0x32, 0xcf, 0x0e, 0xc3, 0x84, 0x78, 0x0a, 0x0f, 0xed, 0x05, 0x99, 0x6f, 0xb5, 0x26, 0x79,
	0xef, 0xbf, 0xf3, 0xc6, 0x9f, 0x60, 0xd0, 0x06, 0x6b, 0xd7, 0xc6, 0x4d, 0xed, 0x59, 0xbb, 0xb6,
	0x3f, 0x1b, 0xed, 0x2f, 0xd6, 0xd4, 0xb5, 0x9a, 0x25, 0x97, 0x30, 0xe2, 0xf4, 0x75, 0x43, 0x4a,
	0x5f, 0x91, 0x16, 0x99, 0xd0, 0x02, 0x87, 0xe0, 0xe7, 0x99, 0xdd, 0x55, 0xc0, 0xfd, 0x3c, 0xc3,
	0x33, 0x18, 0xae, 0x48, 0x52, 0x2d, 0x34, 0x2d, 0x34, 0x29, 0xad, 0xec, 0xce, 0x22, 0x1e, 0xef,
	0xd4, 0x6b, 0x23, 0x26, 0xbf, 0x3c, 0x18, 0xce, 0x9b, 0x73, 0x5c, 0x47, 0x7c, 0x05, 0x41, 0x49,
	0x5a, 0xd8, 0x5e, 0xfd, 0x19, 0xdb, 0x63, 0x1c, 0x9c, 0xc8, 0xad, 0x0b, 0x5f, 0x42, 0x54, 0xaa,
	0xd5, 0x42, 0x55, 0xb4, 0x64, 0x99, 0xad, 0x38, 0x39, 0xdc, 0x08, 0x0f, 0x4b, 0xb5, 0x9a, 0x57,
	0xb4, 0xc4, 0xd7, 0xd0, 0x2b, 0x84, 0x74, 0x6e, 0xf9, 0x1f, 0x77, 0x64, 0x2c, 0xc6, 0x9e, 0xfc,
	0xf0, 0x60, 0xb4, 0x87, 0x53, 0xd5, 0x5a, 0x2a, 0xba, 0x27, 0x1d, 0x42, 0x90, 0x97, 0x55, 0x61,
	0xc9, 0x7a, 0xdc, 0xc6, 0xe6, 0xbd, 0x7e, 0x21, 0x91, 0x51, 0xcd, 0xbe, 0x5b, 0xd5, 0x65, 0xf8,
	0x18, 0x02, 0x33, 0x28, 0xf6, 0xd3, 0x73, 0x0f, 0x9e, 0x94, 0x9e, 0x7d, 0x84, 0xd0, 0x11, 0xe0,
	0x39, 0x44, 0x17, 0x6e, 0x76, 0xf8, 0xf4, 0x2f, 0xf5, 0x3f, 0xc3, 0x1b, 0xb3, 0xe3, 0x0f, 0x0d,
	0x78, 0xf2, 0xe0, 0xa6, 0x6b, 0x7f, 0xb5, 0x37, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x57, 0x1a,
	0xf1, 0xe9, 0x7b, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SkywireClient is the client API for Skywire service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SkywireClient interface {
	Generate(ctx context.Context, in *SkywireRequest, opts ...grpc.CallOption) (*SkywireResponse, error)
}

type skywireClient struct {
	cc *grpc.ClientConn
}

func NewSkywireClient(cc *grpc.ClientConn) SkywireClient {
	return &skywireClient{cc}
}

func (c *skywireClient) Generate(ctx context.Context, in *SkywireRequest, opts ...grpc.CallOption) (*SkywireResponse, error) {
	out := new(SkywireResponse)
	err := c.cc.Invoke(ctx, "/skywire.Skywire/Generate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SkywireServer is the server API for Skywire service.
type SkywireServer interface {
	Generate(context.Context, *SkywireRequest) (*SkywireResponse, error)
}

// UnimplementedSkywireServer can be embedded to have forward compatible implementations.
type UnimplementedSkywireServer struct {
}

func (*UnimplementedSkywireServer) Generate(ctx context.Context, req *SkywireRequest) (*SkywireResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Generate not implemented")
}

func RegisterSkywireServer(s *grpc.Server, srv SkywireServer) {
	s.RegisterService(&_Skywire_serviceDesc, srv)
}

func _Skywire_Generate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SkywireRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SkywireServer).Generate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/skywire.Skywire/Generate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SkywireServer).Generate(ctx, req.(*SkywireRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Skywire_serviceDesc = grpc.ServiceDesc{
	ServiceName: "skywire.Skywire",
	HandlerType: (*SkywireServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Generate",
			Handler:    _Skywire_Generate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "skywire.proto",
}
