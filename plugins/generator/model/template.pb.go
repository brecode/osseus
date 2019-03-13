// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: template.proto

package model

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// A plugin resource
type Template struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Result               string   `protobuf:"bytes,2,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Template) Reset()         { *m = Template{} }
func (m *Template) String() string { return proto.CompactTextString(m) }
func (*Template) ProtoMessage()    {}
func (*Template) Descriptor() ([]byte, []int) {
	return fileDescriptor_b1b68e1b5f001c74, []int{0}
}
func (m *Template) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Template.Unmarshal(m, b)
}
func (m *Template) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Template.Marshal(b, m, deterministic)
}
func (m *Template) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Template.Merge(m, src)
}
func (m *Template) XXX_Size() int {
	return xxx_messageInfo_Template.Size(m)
}
func (m *Template) XXX_DiscardUnknown() {
	xxx_messageInfo_Template.DiscardUnknown(m)
}

var xxx_messageInfo_Template proto.InternalMessageInfo

func (m *Template) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Template) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

func (*Template) XXX_MessageName() string {
	return "model.Template"
}
func init() {
	proto.RegisterType((*Template)(nil), "model.Template")
}

func init() { proto.RegisterFile("template.proto", fileDescriptor_b1b68e1b5f001c74) }

var fileDescriptor_b1b68e1b5f001c74 = []byte{
	// 128 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2b, 0x49, 0xcd, 0x2d,
	0xc8, 0x49, 0x2c, 0x49, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xcd, 0xcd, 0x4f, 0x49,
	0xcd, 0x91, 0xd2, 0x4d, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x4f, 0xcf,
	0x4f, 0xcf, 0xd7, 0x07, 0xcb, 0x26, 0x95, 0xa6, 0x81, 0x79, 0x60, 0x0e, 0x98, 0x05, 0xd1, 0xa5,
	0x64, 0xc6, 0xc5, 0x11, 0x02, 0x35, 0x47, 0x48, 0x88, 0x8b, 0x25, 0x2f, 0x31, 0x37, 0x55, 0x82,
	0x51, 0x81, 0x51, 0x83, 0x33, 0x08, 0xcc, 0x16, 0x12, 0xe3, 0x62, 0x2b, 0x4a, 0x2d, 0x2e, 0xcd,
	0x29, 0x91, 0x60, 0x02, 0x8b, 0x42, 0x79, 0x4e, 0x2c, 0x27, 0x1e, 0xcb, 0x31, 0x26, 0xb1, 0x81,
	0x0d, 0x31, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x23, 0x2c, 0x90, 0xf9, 0x8c, 0x00, 0x00, 0x00,
}
