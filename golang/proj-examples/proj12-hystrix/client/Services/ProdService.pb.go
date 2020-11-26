// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ProdService.proto

package Services

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

type ProdsRequest struct {
	// @inject_tag: json:"size",form:"size"
	Size int32 `protobuf:"varint,1,opt,name=size,proto3" json:"size" form:"size"`
	// @inject_tag: uri:"pid"
	ProdId               int32    `protobuf:"varint,2,opt,name=prod_id,json=prodId,proto3" json:"prod_id,omitempty" uri:"pid"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProdsRequest) Reset()         { *m = ProdsRequest{} }
func (m *ProdsRequest) String() string { return proto.CompactTextString(m) }
func (*ProdsRequest) ProtoMessage()    {}
func (*ProdsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_50db98fd6a3e2ab5, []int{0}
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

func (m *ProdsRequest) GetProdId() int32 {
	if m != nil {
		return m.ProdId
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
	return fileDescriptor_50db98fd6a3e2ab5, []int{1}
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

type ProdDetailResponse struct {
	Data                 *ProdModel `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ProdDetailResponse) Reset()         { *m = ProdDetailResponse{} }
func (m *ProdDetailResponse) String() string { return proto.CompactTextString(m) }
func (*ProdDetailResponse) ProtoMessage()    {}
func (*ProdDetailResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_50db98fd6a3e2ab5, []int{2}
}

func (m *ProdDetailResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProdDetailResponse.Unmarshal(m, b)
}
func (m *ProdDetailResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProdDetailResponse.Marshal(b, m, deterministic)
}
func (m *ProdDetailResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProdDetailResponse.Merge(m, src)
}
func (m *ProdDetailResponse) XXX_Size() int {
	return xxx_messageInfo_ProdDetailResponse.Size(m)
}
func (m *ProdDetailResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ProdDetailResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ProdDetailResponse proto.InternalMessageInfo

func (m *ProdDetailResponse) GetData() *ProdModel {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*ProdsRequest)(nil), "Services.ProdsRequest")
	proto.RegisterType((*ProdListResponse)(nil), "Services.ProdListResponse")
	proto.RegisterType((*ProdDetailResponse)(nil), "Services.ProdDetailResponse")
}

func init() {
	proto.RegisterFile("ProdService.proto", fileDescriptor_50db98fd6a3e2ab5)
}

var fileDescriptor_50db98fd6a3e2ab5 = []byte{
	// 223 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x0c, 0x28, 0xca, 0x4f,
	0x09, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x80,
	0x72, 0x8b, 0xa5, 0x78, 0x7c, 0xf3, 0x53, 0x52, 0x73, 0x8a, 0x21, 0xe2, 0x4a, 0xd6, 0x5c, 0x3c,
	0x20, 0xc5, 0xc5, 0x41, 0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25, 0x42, 0x42, 0x5c, 0x2c, 0xc5, 0x99,
	0x55, 0xa9, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0xac, 0x41, 0x60, 0xb6, 0x90, 0x38, 0x17, 0x7b, 0x41,
	0x51, 0x7e, 0x4a, 0x7c, 0x66, 0x8a, 0x04, 0x13, 0x58, 0x98, 0x0d, 0xc4, 0xf5, 0x4c, 0x51, 0xb2,
	0xe6, 0x12, 0x00, 0x69, 0xf6, 0xc9, 0x2c, 0x2e, 0x09, 0x4a, 0x2d, 0x2e, 0xc8, 0xcf, 0x2b, 0x4e,
	0x15, 0x52, 0xe7, 0x62, 0x49, 0x49, 0x2c, 0x49, 0x94, 0x60, 0x54, 0x60, 0xd6, 0xe0, 0x36, 0x12,
	0xd6, 0x83, 0xd9, 0xab, 0x07, 0x52, 0x09, 0xb6, 0x3a, 0x08, 0xac, 0x40, 0xc9, 0x96, 0x4b, 0x08,
	0x24, 0xe4, 0x92, 0x5a, 0x92, 0x98, 0x99, 0x83, 0x45, 0x3b, 0x23, 0x5e, 0xed, 0x46, 0x33, 0x18,
	0xb9, 0xb8, 0x91, 0xbc, 0x29, 0xe4, 0xc4, 0xc5, 0xe3, 0x9e, 0x5a, 0x02, 0xf6, 0x0b, 0xc8, 0x3d,
	0x42, 0x62, 0xa8, 0x5a, 0x61, 0x1e, 0x94, 0x92, 0x42, 0x15, 0x47, 0x71, 0xbb, 0x2b, 0x17, 0x2f,
	0xd4, 0x0c, 0x88, 0xab, 0x70, 0x1a, 0x22, 0x83, 0x2a, 0x8e, 0xea, 0x87, 0x24, 0x36, 0x70, 0xd0,
	0x1a, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0xa7, 0x66, 0x0d, 0xd6, 0x87, 0x01, 0x00, 0x00,
}
