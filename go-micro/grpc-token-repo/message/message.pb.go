// Code generated by protoc-gen-go. DO NOT EDIT.
// source: message.proto

package message

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

type ReqNil struct {
	Req                  int32    `protobuf:"varint,1,opt,name=req,proto3" json:"req,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqNil) Reset()         { *m = ReqNil{} }
func (m *ReqNil) String() string { return proto.CompactTextString(m) }
func (*ReqNil) ProtoMessage()    {}
func (*ReqNil) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{0}
}

func (m *ReqNil) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqNil.Unmarshal(m, b)
}
func (m *ReqNil) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqNil.Marshal(b, m, deterministic)
}
func (m *ReqNil) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqNil.Merge(m, src)
}
func (m *ReqNil) XXX_Size() int {
	return xxx_messageInfo_ReqNil.Size(m)
}
func (m *ReqNil) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqNil.DiscardUnknown(m)
}

var xxx_messageInfo_ReqNil proto.InternalMessageInfo

func (m *ReqNil) GetReq() int32 {
	if m != nil {
		return m.Req
	}
	return 0
}

type RespMsg struct {
	// code, data, msg
	Code                 string   `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Data                 string   `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	Msg                  string   `protobuf:"bytes,3,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RespMsg) Reset()         { *m = RespMsg{} }
func (m *RespMsg) String() string { return proto.CompactTextString(m) }
func (*RespMsg) ProtoMessage()    {}
func (*RespMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{1}
}

func (m *RespMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RespMsg.Unmarshal(m, b)
}
func (m *RespMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RespMsg.Marshal(b, m, deterministic)
}
func (m *RespMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RespMsg.Merge(m, src)
}
func (m *RespMsg) XXX_Size() int {
	return xxx_messageInfo_RespMsg.Size(m)
}
func (m *RespMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_RespMsg.DiscardUnknown(m)
}

var xxx_messageInfo_RespMsg proto.InternalMessageInfo

func (m *RespMsg) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *RespMsg) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func (m *RespMsg) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func init() {
	proto.RegisterType((*ReqNil)(nil), "message.ReqNil")
	proto.RegisterType((*RespMsg)(nil), "message.RespMsg")
}

func init() {
	proto.RegisterFile("message.proto", fileDescriptor_33c57e4bae7b9afd)
}

var fileDescriptor_33c57e4bae7b9afd = []byte{
	// 175 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcd, 0x4d, 0x2d, 0x2e,
	0x4e, 0x4c, 0x4f, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x87, 0x72, 0x95, 0xa4, 0xb8,
	0xd8, 0x82, 0x52, 0x0b, 0xfd, 0x32, 0x73, 0x84, 0x04, 0xb8, 0x98, 0x8b, 0x52, 0x0b, 0x25, 0x18,
	0x15, 0x18, 0x35, 0x58, 0x83, 0x40, 0x4c, 0x25, 0x67, 0x2e, 0xf6, 0xa0, 0xd4, 0xe2, 0x02, 0xdf,
	0xe2, 0x74, 0x21, 0x21, 0x2e, 0x96, 0xe4, 0xfc, 0x94, 0x54, 0xb0, 0x2c, 0x67, 0x10, 0x98, 0x0d,
	0x12, 0x4b, 0x49, 0x2c, 0x49, 0x94, 0x60, 0x82, 0x88, 0x81, 0xd8, 0x20, 0x43, 0x72, 0x8b, 0xd3,
	0x25, 0x98, 0xc1, 0x42, 0x20, 0xa6, 0x51, 0x36, 0x17, 0xb7, 0x63, 0x69, 0x49, 0x46, 0x70, 0x6a,
	0x51, 0x59, 0x66, 0x72, 0xaa, 0x90, 0x1e, 0x17, 0xa7, 0x73, 0x46, 0xb6, 0x47, 0x6a, 0x62, 0x4e,
	0x49, 0x86, 0x10, 0xbf, 0x1e, 0xcc, 0x55, 0x10, 0x37, 0x48, 0x09, 0x20, 0x09, 0x40, 0x2c, 0xd6,
	0xe1, 0x62, 0x77, 0x4f, 0x2d, 0x09, 0xce, 0xcf, 0x4d, 0x25, 0x42, 0x75, 0x12, 0x1b, 0xd8, 0x77,
	0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe2, 0x9a, 0x80, 0xe3, 0xee, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AuthServiceClient is the client API for AuthService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuthServiceClient interface {
	ChkHealth(ctx context.Context, in *ReqNil, opts ...grpc.CallOption) (*RespMsg, error)
	GetSome(ctx context.Context, in *ReqNil, opts ...grpc.CallOption) (*RespMsg, error)
}

type authServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthServiceClient(cc grpc.ClientConnInterface) AuthServiceClient {
	return &authServiceClient{cc}
}

func (c *authServiceClient) ChkHealth(ctx context.Context, in *ReqNil, opts ...grpc.CallOption) (*RespMsg, error) {
	out := new(RespMsg)
	err := c.cc.Invoke(ctx, "/message.AuthService/ChkHealth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) GetSome(ctx context.Context, in *ReqNil, opts ...grpc.CallOption) (*RespMsg, error) {
	out := new(RespMsg)
	err := c.cc.Invoke(ctx, "/message.AuthService/GetSome", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServiceServer is the server API for AuthService service.
type AuthServiceServer interface {
	ChkHealth(context.Context, *ReqNil) (*RespMsg, error)
	GetSome(context.Context, *ReqNil) (*RespMsg, error)
}

// UnimplementedAuthServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAuthServiceServer struct {
}

func (*UnimplementedAuthServiceServer) ChkHealth(ctx context.Context, req *ReqNil) (*RespMsg, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChkHealth not implemented")
}
func (*UnimplementedAuthServiceServer) GetSome(ctx context.Context, req *ReqNil) (*RespMsg, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSome not implemented")
}

func RegisterAuthServiceServer(s *grpc.Server, srv AuthServiceServer) {
	s.RegisterService(&_AuthService_serviceDesc, srv)
}

func _AuthService_ChkHealth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqNil)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).ChkHealth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/message.AuthService/ChkHealth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).ChkHealth(ctx, req.(*ReqNil))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_GetSome_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqNil)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).GetSome(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/message.AuthService/GetSome",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).GetSome(ctx, req.(*ReqNil))
	}
	return interceptor(ctx, in, info, handler)
}

var _AuthService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "message.AuthService",
	HandlerType: (*AuthServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ChkHealth",
			Handler:    _AuthService_ChkHealth_Handler,
		},
		{
			MethodName: "GetSome",
			Handler:    _AuthService_GetSome_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "message.proto",
}