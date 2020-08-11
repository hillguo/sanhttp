// Code generated by protoc-gen-go. DO NOT EDIT.
// source: test.proto

package example

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

type TestReq struct {
	Uin                  string   `protobuf:"bytes,1,opt,name=uin,proto3" json:"uin,omitempty"`
	Vid                  uint64   `protobuf:"varint,2,opt,name=vid,proto3" json:"vid,omitempty"`
	Ids                  []uint32 `protobuf:"varint,3,rep,packed,name=ids,proto3" json:"ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TestReq) Reset()         { *m = TestReq{} }
func (m *TestReq) String() string { return proto.CompactTextString(m) }
func (*TestReq) ProtoMessage()    {}
func (*TestReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_c161fcfdc0c3ff1e, []int{0}
}

func (m *TestReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TestReq.Unmarshal(m, b)
}
func (m *TestReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TestReq.Marshal(b, m, deterministic)
}
func (m *TestReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestReq.Merge(m, src)
}
func (m *TestReq) XXX_Size() int {
	return xxx_messageInfo_TestReq.Size(m)
}
func (m *TestReq) XXX_DiscardUnknown() {
	xxx_messageInfo_TestReq.DiscardUnknown(m)
}

var xxx_messageInfo_TestReq proto.InternalMessageInfo

func (m *TestReq) GetUin() string {
	if m != nil {
		return m.Uin
	}
	return ""
}

func (m *TestReq) GetVid() uint64 {
	if m != nil {
		return m.Vid
	}
	return 0
}

func (m *TestReq) GetIds() []uint32 {
	if m != nil {
		return m.Ids
	}
	return nil
}

type TestRsp struct {
	Vid                  uint64   `protobuf:"varint,1,opt,name=vid,proto3" json:"vid,omitempty"`
	Ok                   bool     `protobuf:"varint,2,opt,name=ok,proto3" json:"ok,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TestRsp) Reset()         { *m = TestRsp{} }
func (m *TestRsp) String() string { return proto.CompactTextString(m) }
func (*TestRsp) ProtoMessage()    {}
func (*TestRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_c161fcfdc0c3ff1e, []int{1}
}

func (m *TestRsp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TestRsp.Unmarshal(m, b)
}
func (m *TestRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TestRsp.Marshal(b, m, deterministic)
}
func (m *TestRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestRsp.Merge(m, src)
}
func (m *TestRsp) XXX_Size() int {
	return xxx_messageInfo_TestRsp.Size(m)
}
func (m *TestRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_TestRsp.DiscardUnknown(m)
}

var xxx_messageInfo_TestRsp proto.InternalMessageInfo

func (m *TestRsp) GetVid() uint64 {
	if m != nil {
		return m.Vid
	}
	return 0
}

func (m *TestRsp) GetOk() bool {
	if m != nil {
		return m.Ok
	}
	return false
}

type GetVidByNameReq struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Type                 string   `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetVidByNameReq) Reset()         { *m = GetVidByNameReq{} }
func (m *GetVidByNameReq) String() string { return proto.CompactTextString(m) }
func (*GetVidByNameReq) ProtoMessage()    {}
func (*GetVidByNameReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_c161fcfdc0c3ff1e, []int{2}
}

func (m *GetVidByNameReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetVidByNameReq.Unmarshal(m, b)
}
func (m *GetVidByNameReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetVidByNameReq.Marshal(b, m, deterministic)
}
func (m *GetVidByNameReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetVidByNameReq.Merge(m, src)
}
func (m *GetVidByNameReq) XXX_Size() int {
	return xxx_messageInfo_GetVidByNameReq.Size(m)
}
func (m *GetVidByNameReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetVidByNameReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetVidByNameReq proto.InternalMessageInfo

func (m *GetVidByNameReq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GetVidByNameReq) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

type GetVidByNameRsp struct {
	Vid                  uint64   `protobuf:"varint,1,opt,name=vid,proto3" json:"vid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetVidByNameRsp) Reset()         { *m = GetVidByNameRsp{} }
func (m *GetVidByNameRsp) String() string { return proto.CompactTextString(m) }
func (*GetVidByNameRsp) ProtoMessage()    {}
func (*GetVidByNameRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_c161fcfdc0c3ff1e, []int{3}
}

func (m *GetVidByNameRsp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetVidByNameRsp.Unmarshal(m, b)
}
func (m *GetVidByNameRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetVidByNameRsp.Marshal(b, m, deterministic)
}
func (m *GetVidByNameRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetVidByNameRsp.Merge(m, src)
}
func (m *GetVidByNameRsp) XXX_Size() int {
	return xxx_messageInfo_GetVidByNameRsp.Size(m)
}
func (m *GetVidByNameRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_GetVidByNameRsp.DiscardUnknown(m)
}

var xxx_messageInfo_GetVidByNameRsp proto.InternalMessageInfo

func (m *GetVidByNameRsp) GetVid() uint64 {
	if m != nil {
		return m.Vid
	}
	return 0
}

func init() {
	proto.RegisterType((*TestReq)(nil), "test.TestReq")
	proto.RegisterType((*TestRsp)(nil), "test.TestRsp")
	proto.RegisterType((*GetVidByNameReq)(nil), "test.GetVidByNameReq")
	proto.RegisterType((*GetVidByNameRsp)(nil), "test.GetVidByNameRsp")
}

func init() { proto.RegisterFile("test.proto", fileDescriptor_c161fcfdc0c3ff1e) }

var fileDescriptor_c161fcfdc0c3ff1e = []byte{
	// 218 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0x49, 0x2d, 0x2e,
	0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x01, 0xb1, 0x95, 0xec, 0xb9, 0xd8, 0x43, 0x52,
	0x8b, 0x4b, 0x82, 0x52, 0x0b, 0x85, 0x04, 0xb8, 0x98, 0x4b, 0x33, 0xf3, 0x24, 0x18, 0x15, 0x18,
	0x35, 0x38, 0x83, 0x40, 0x4c, 0x90, 0x48, 0x59, 0x66, 0x8a, 0x04, 0x93, 0x02, 0xa3, 0x06, 0x4b,
	0x10, 0x88, 0x09, 0x12, 0xc9, 0x4c, 0x29, 0x96, 0x60, 0x56, 0x60, 0xd6, 0xe0, 0x0d, 0x02, 0x31,
	0x95, 0xb4, 0xa1, 0x06, 0x14, 0x17, 0xc0, 0x94, 0x33, 0x22, 0x94, 0xf3, 0x71, 0x31, 0xe5, 0x67,
	0x83, 0xf5, 0x73, 0x04, 0x31, 0xe5, 0x67, 0x2b, 0x59, 0x72, 0xf1, 0xbb, 0xa7, 0x96, 0x84, 0x65,
	0xa6, 0x38, 0x55, 0xfa, 0x25, 0xe6, 0xa6, 0x82, 0x6c, 0x15, 0xe2, 0x62, 0xc9, 0x4b, 0xcc, 0x4d,
	0x85, 0x5a, 0x0b, 0x66, 0x83, 0xc4, 0x4a, 0x2a, 0x0b, 0x52, 0xc1, 0x1a, 0x39, 0x83, 0xc0, 0x6c,
	0x25, 0x65, 0x34, 0xad, 0xd8, 0xec, 0x33, 0xca, 0xe3, 0x62, 0x01, 0x39, 0x46, 0x48, 0x0d, 0x4a,
	0xf3, 0xea, 0x81, 0x3d, 0x0c, 0xf5, 0xa1, 0x14, 0x32, 0xb7, 0xb8, 0x40, 0x89, 0x41, 0xc8, 0x8e,
	0x8b, 0x07, 0xd9, 0x50, 0x21, 0x51, 0x88, 0x02, 0x34, 0x37, 0x4a, 0x61, 0x13, 0x06, 0xe9, 0x77,
	0x62, 0x8b, 0x62, 0xc9, 0x4d, 0xcc, 0xcc, 0x4b, 0x62, 0x03, 0x07, 0xa9, 0x31, 0x20, 0x00, 0x00,
	0xff, 0xff, 0xed, 0xc2, 0x59, 0x67, 0x60, 0x01, 0x00, 0x00,
}
