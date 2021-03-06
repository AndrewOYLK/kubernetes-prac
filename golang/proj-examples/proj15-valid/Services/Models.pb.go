// Code generated by protoc-gen-go. DO NOT EDIT.
// source: Models.proto

package Services

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type UserModel struct {
	UserId               int32    `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	UserName             string   `protobuf:"bytes,2,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	UserPwd              string   `protobuf:"bytes,3,opt,name=user_pwd,json=userPwd,proto3" json:"user_pwd,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserModel) Reset()         { *m = UserModel{} }
func (m *UserModel) String() string { return proto.CompactTextString(m) }
func (*UserModel) ProtoMessage()    {}
func (*UserModel) Descriptor() ([]byte, []int) {
	return fileDescriptor_96b05f67b8e9f86a, []int{0}
}

func (m *UserModel) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserModel.Unmarshal(m, b)
}
func (m *UserModel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserModel.Marshal(b, m, deterministic)
}
func (m *UserModel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserModel.Merge(m, src)
}
func (m *UserModel) XXX_Size() int {
	return xxx_messageInfo_UserModel.Size(m)
}
func (m *UserModel) XXX_DiscardUnknown() {
	xxx_messageInfo_UserModel.DiscardUnknown(m)
}

var xxx_messageInfo_UserModel proto.InternalMessageInfo

func (m *UserModel) GetUserId() int32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *UserModel) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *UserModel) GetUserPwd() string {
	if m != nil {
		return m.UserPwd
	}
	return ""
}

func init() {
	proto.RegisterType((*UserModel)(nil), "Services.UserModel")
}

func init() { proto.RegisterFile("Models.proto", fileDescriptor_96b05f67b8e9f86a) }

var fileDescriptor_96b05f67b8e9f86a = []byte{
	// 127 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xf1, 0xcd, 0x4f, 0x49,
	0xcd, 0x29, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x08, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c,
	0x4e, 0x2d, 0x56, 0x8a, 0xe1, 0xe2, 0x0c, 0x2d, 0x4e, 0x2d, 0x02, 0xcb, 0x0a, 0x89, 0x73, 0xb1,
	0x97, 0x16, 0xa7, 0x16, 0xc5, 0x67, 0xa6, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0xb0, 0x06, 0xb1, 0x81,
	0xb8, 0x9e, 0x29, 0x42, 0xd2, 0x5c, 0x9c, 0x60, 0x89, 0xbc, 0xc4, 0xdc, 0x54, 0x09, 0x26, 0x05,
	0x46, 0x0d, 0xce, 0x20, 0x0e, 0x90, 0x80, 0x5f, 0x62, 0x6e, 0xaa, 0x90, 0x24, 0x17, 0x98, 0x1d,
	0x5f, 0x50, 0x9e, 0x22, 0xc1, 0x0c, 0x96, 0x03, 0x9b, 0x12, 0x50, 0x9e, 0x92, 0xc4, 0x06, 0xb6,
	0xce, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xcd, 0x28, 0x9b, 0xf9, 0x7e, 0x00, 0x00, 0x00,
}
