// Code generated by protoc-gen-go. DO NOT EDIT.
// source: alpha/a.proto

package proto_all_go

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

type AlphaTest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AlphaTest) Reset()         { *m = AlphaTest{} }
func (m *AlphaTest) String() string { return proto.CompactTextString(m) }
func (*AlphaTest) ProtoMessage()    {}
func (*AlphaTest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9e24ddc4fb8838e8, []int{0}
}

func (m *AlphaTest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AlphaTest.Unmarshal(m, b)
}
func (m *AlphaTest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AlphaTest.Marshal(b, m, deterministic)
}
func (m *AlphaTest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AlphaTest.Merge(m, src)
}
func (m *AlphaTest) XXX_Size() int {
	return xxx_messageInfo_AlphaTest.Size(m)
}
func (m *AlphaTest) XXX_DiscardUnknown() {
	xxx_messageInfo_AlphaTest.DiscardUnknown(m)
}

var xxx_messageInfo_AlphaTest proto.InternalMessageInfo

func init() {
	proto.RegisterType((*AlphaTest)(nil), "alpha.AlphaTest")
}

func init() { proto.RegisterFile("alpha/a.proto", fileDescriptor_9e24ddc4fb8838e8) }

var fileDescriptor_9e24ddc4fb8838e8 = []byte{
	// 87 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4d, 0xcc, 0x29, 0xc8,
	0x48, 0xd4, 0x4f, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x73, 0x95, 0xb8, 0xb9,
	0x38, 0x1d, 0x41, 0x8c, 0x90, 0xd4, 0xe2, 0x12, 0x27, 0xb9, 0x28, 0x99, 0xf4, 0xcc, 0x92, 0x8c,
	0xd2, 0x24, 0xbd, 0xe4, 0xfc, 0x5c, 0x7d, 0x93, 0xbc, 0x92, 0x54, 0x7d, 0xb0, 0x62, 0xdd, 0xc4,
	0x9c, 0x1c, 0xdd, 0xf4, 0xfc, 0x24, 0x36, 0x30, 0xcf, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x9d,
	0x50, 0x67, 0x55, 0x4b, 0x00, 0x00, 0x00,
}
