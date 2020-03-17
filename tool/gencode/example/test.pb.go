// Code generated by protoc-gen-go. DO NOT EDIT.
// source: example/test.proto

package test

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
	return fileDescriptor_3861ab4853c9fd5c, []int{0}
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
	return fileDescriptor_3861ab4853c9fd5c, []int{1}
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

func init() {
	proto.RegisterType((*TestReq)(nil), "TestReq")
	proto.RegisterType((*TestRsp)(nil), "TestRsp")
}

func init() {
	proto.RegisterFile("example/test.proto", fileDescriptor_3861ab4853c9fd5c)
}

var fileDescriptor_3861ab4853c9fd5c = []byte{
	// 147 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4a, 0xad, 0x48, 0xcc,
	0x2d, 0xc8, 0x49, 0xd5, 0x2f, 0x49, 0x2d, 0x2e, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0xb2,
	0xe7, 0x62, 0x0f, 0x49, 0x2d, 0x2e, 0x09, 0x4a, 0x2d, 0x14, 0x12, 0xe0, 0x62, 0x2e, 0xcd, 0xcc,
	0x93, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x02, 0x31, 0x41, 0x22, 0x65, 0x99, 0x29, 0x12, 0x4c,
	0x0a, 0x8c, 0x1a, 0x2c, 0x41, 0x20, 0x26, 0x48, 0x24, 0x33, 0xa5, 0x58, 0x82, 0x59, 0x81, 0x59,
	0x83, 0x37, 0x08, 0xc4, 0x54, 0xd2, 0x86, 0x1a, 0x50, 0x5c, 0x00, 0x53, 0xce, 0x88, 0x50, 0xce,
	0xc7, 0xc5, 0x94, 0x9f, 0x0d, 0xd6, 0xcf, 0x11, 0xc4, 0x94, 0x9f, 0x6d, 0xa4, 0xc2, 0xc5, 0x02,
	0x52, 0x2c, 0x24, 0x03, 0xa5, 0x39, 0xf4, 0xa0, 0x96, 0x4b, 0x41, 0x59, 0xc5, 0x05, 0x4a, 0x0c,
	0x49, 0x6c, 0x60, 0xa7, 0x19, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x3c, 0xeb, 0xcc, 0xc0, 0xb0,
	0x00, 0x00, 0x00,
}