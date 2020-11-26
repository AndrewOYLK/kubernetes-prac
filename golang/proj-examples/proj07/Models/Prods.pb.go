// Code generated by protoc-gen-go. DO NOT EDIT.
// source: Prods.proto

package Models

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
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

type ProdModel struct {
	// @inject_tag: json:"pid"
	ProdID int32 `protobuf:"varint,1,opt,name=ProdID,proto3" json:"pid"`
	// @inject_tag: json:"pname"
	ProdName             string   `protobuf:"bytes,2,opt,name=ProdName,proto3" json:"pname"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProdModel) Reset()         { *m = ProdModel{} }
func (m *ProdModel) String() string { return proto.CompactTextString(m) }
func (*ProdModel) ProtoMessage()    {}
func (*ProdModel) Descriptor() ([]byte, []int) {
	return fileDescriptor_031bb715af7191d4, []int{0}
}

func (m *ProdModel) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProdModel.Unmarshal(m, b)
}
func (m *ProdModel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProdModel.Marshal(b, m, deterministic)
}
func (m *ProdModel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProdModel.Merge(m, src)
}
func (m *ProdModel) XXX_Size() int {
	return xxx_messageInfo_ProdModel.Size(m)
}
func (m *ProdModel) XXX_DiscardUnknown() {
	xxx_messageInfo_ProdModel.DiscardUnknown(m)
}

var xxx_messageInfo_ProdModel proto.InternalMessageInfo

func (m *ProdModel) GetProdID() int32 {
	if m != nil {
		return m.ProdID
	}
	return 0
}

func (m *ProdModel) GetProdName() string {
	if m != nil {
		return m.ProdName
	}
	return ""
}

type ProdsRequest struct {
	Size                 int32    `protobuf:"varint,1,opt,name=size,proto3" json:"size,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProdsRequest) Reset()         { *m = ProdsRequest{} }
func (m *ProdsRequest) String() string { return proto.CompactTextString(m) }
func (*ProdsRequest) ProtoMessage()    {}
func (*ProdsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_031bb715af7191d4, []int{1}
}

func (m *ProdsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProdsRequest.Unmarshal(m, b)
}
func (m *ProdsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProdsRequest.Marshal(b, m, deterministic)
}
func (m *ProdsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProdsRequest.Merge(m, src)
}
func (m *ProdsRequest) XXX_Size() int {
	return xxx_messageInfo_ProdsRequest.Size(m)
}
func (m *ProdsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ProdsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ProdsRequest proto.InternalMessageInfo

func (m *ProdsRequest) GetSize() int32 {
	if m != nil {
		return m.Size
	}
	return 0
}

type ProdListResponse struct {
	Data                 []*ProdModel `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ProdListResponse) Reset()         { *m = ProdListResponse{} }
func (m *ProdListResponse) String() string { return proto.CompactTextString(m) }
func (*ProdListResponse) ProtoMessage()    {}
func (*ProdListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_031bb715af7191d4, []int{2}
}

func (m *ProdListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProdListResponse.Unmarshal(m, b)
}
func (m *ProdListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProdListResponse.Marshal(b, m, deterministic)
}
func (m *ProdListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProdListResponse.Merge(m, src)
}
func (m *ProdListResponse) XXX_Size() int {
	return xxx_messageInfo_ProdListResponse.Size(m)
}
func (m *ProdListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ProdListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ProdListResponse proto.InternalMessageInfo

func (m *ProdListResponse) GetData() []*ProdModel {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*ProdModel)(nil), "Models.ProdModel")
	proto.RegisterType((*ProdsRequest)(nil), "Models.ProdsRequest")
	proto.RegisterType((*ProdListResponse)(nil), "Models.ProdListResponse")
}

func init() {
	proto.RegisterFile("Prods.proto", fileDescriptor_031bb715af7191d4)
}

var fileDescriptor_031bb715af7191d4 = []byte{
	// 161 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x0e, 0x28, 0xca, 0x4f,
	0x29, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xf3, 0xcd, 0x4f, 0x49, 0xcd, 0x29, 0x56,
	0xb2, 0xe7, 0xe2, 0x04, 0x09, 0x83, 0x79, 0x42, 0x62, 0x5c, 0x6c, 0x20, 0x8e, 0xa7, 0x8b, 0x04,
	0xa3, 0x02, 0xa3, 0x06, 0x6b, 0x10, 0x94, 0x27, 0x24, 0xc5, 0xc5, 0x01, 0x62, 0xf9, 0x25, 0xe6,
	0xa6, 0x4a, 0x30, 0x29, 0x30, 0x6a, 0x70, 0x06, 0xc1, 0xf9, 0x4a, 0x4a, 0x5c, 0x3c, 0x60, 0x73,
	0x83, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x84, 0x84, 0xb8, 0x58, 0x8a, 0x33, 0xab, 0x52, 0xa1,
	0x26, 0x80, 0xd9, 0x4a, 0x96, 0x5c, 0x02, 0x20, 0x35, 0x3e, 0x99, 0xc5, 0x25, 0x41, 0xa9, 0xc5,
	0x05, 0xf9, 0x79, 0xc5, 0xa9, 0x42, 0xaa, 0x5c, 0x2c, 0x29, 0x89, 0x25, 0x89, 0x12, 0x8c, 0x0a,
	0xcc, 0x1a, 0xdc, 0x46, 0x82, 0x7a, 0x10, 0xf7, 0xe8, 0xc1, 0x1d, 0x13, 0x04, 0x96, 0x4e, 0x62,
	0x03, 0x3b, 0xd7, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x87, 0xb4, 0xde, 0x5b, 0xbd, 0x00, 0x00,
	0x00,
}
