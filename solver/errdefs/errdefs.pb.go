// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: errdefs.proto

package errdefs

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	pb "github.com/moby/buildkit/solver/pb"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Vertex struct {
	Digest               string   `protobuf:"bytes,1,opt,name=digest,proto3" json:"digest,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Vertex) Reset()         { *m = Vertex{} }
func (m *Vertex) String() string { return proto.CompactTextString(m) }
func (*Vertex) ProtoMessage()    {}
func (*Vertex) Descriptor() ([]byte, []int) {
	return fileDescriptor_689dc58a5060aff5, []int{0}
}
func (m *Vertex) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Vertex.Unmarshal(m, b)
}
func (m *Vertex) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Vertex.Marshal(b, m, deterministic)
}
func (m *Vertex) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Vertex.Merge(m, src)
}
func (m *Vertex) XXX_Size() int {
	return xxx_messageInfo_Vertex.Size(m)
}
func (m *Vertex) XXX_DiscardUnknown() {
	xxx_messageInfo_Vertex.DiscardUnknown(m)
}

var xxx_messageInfo_Vertex proto.InternalMessageInfo

func (m *Vertex) GetDigest() string {
	if m != nil {
		return m.Digest
	}
	return ""
}

type Source struct {
	Info                 *pb.SourceInfo `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
	Ranges               []*pb.Range    `protobuf:"bytes,2,rep,name=ranges,proto3" json:"ranges,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Source) Reset()         { *m = Source{} }
func (m *Source) String() string { return proto.CompactTextString(m) }
func (*Source) ProtoMessage()    {}
func (*Source) Descriptor() ([]byte, []int) {
	return fileDescriptor_689dc58a5060aff5, []int{1}
}
func (m *Source) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Source.Unmarshal(m, b)
}
func (m *Source) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Source.Marshal(b, m, deterministic)
}
func (m *Source) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Source.Merge(m, src)
}
func (m *Source) XXX_Size() int {
	return xxx_messageInfo_Source.Size(m)
}
func (m *Source) XXX_DiscardUnknown() {
	xxx_messageInfo_Source.DiscardUnknown(m)
}

var xxx_messageInfo_Source proto.InternalMessageInfo

func (m *Source) GetInfo() *pb.SourceInfo {
	if m != nil {
		return m.Info
	}
	return nil
}

func (m *Source) GetRanges() []*pb.Range {
	if m != nil {
		return m.Ranges
	}
	return nil
}

func init() {
	proto.RegisterType((*Vertex)(nil), "errdefs.Vertex")
	proto.RegisterType((*Source)(nil), "errdefs.Source")
}

func init() { proto.RegisterFile("errdefs.proto", fileDescriptor_689dc58a5060aff5) }

var fileDescriptor_689dc58a5060aff5 = []byte{
	// 177 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x2c, 0xcd, 0xc1, 0x8a, 0x83, 0x30,
	0x10, 0x80, 0x61, 0xdc, 0x5d, 0xb2, 0x18, 0xd9, 0x3d, 0xe4, 0x50, 0xa4, 0x27, 0xeb, 0xc9, 0x43,
	0x49, 0xc0, 0x3e, 0x45, 0x4f, 0x85, 0x14, 0x7a, 0x6f, 0x74, 0xb4, 0xa1, 0xea, 0x84, 0x49, 0x2c,
	0xed, 0xdb, 0x17, 0x6d, 0x8e, 0xff, 0x7c, 0x33, 0x0c, 0xff, 0x03, 0xa2, 0x16, 0x3a, 0x2f, 0x1d,
	0x61, 0x40, 0xf1, 0x1b, 0x73, 0xbb, 0xef, 0x6d, 0xb8, 0xcd, 0x46, 0x36, 0x38, 0xaa, 0x11, 0xcd,
	0x4b, 0x99, 0xd9, 0x0e, 0xed, 0xdd, 0x06, 0xe5, 0x71, 0x78, 0x00, 0x29, 0x67, 0x14, 0xba, 0x78,
	0x56, 0x16, 0x9c, 0x5d, 0x80, 0x02, 0x3c, 0xc5, 0x86, 0xb3, 0xd6, 0xf6, 0xe0, 0x43, 0x9e, 0x14,
	0x49, 0x95, 0xea, 0x58, 0xe5, 0x89, 0xb3, 0x33, 0xce, 0xd4, 0x80, 0x28, 0xf9, 0x8f, 0x9d, 0x3a,
	0x5c, 0x3d, 0xab, 0xff, 0xa5, 0x33, 0xf2, 0x23, 0xc7, 0xa9, 0x43, 0xbd, 0x9a, 0xd8, 0x71, 0x46,
	0xd7, 0xa9, 0x07, 0x9f, 0x7f, 0x15, 0xdf, 0x55, 0x56, 0xa7, 0xcb, 0x96, 0x5e, 0x26, 0x3a, 0x82,
	0x61, 0xeb, 0xe7, 0xc3, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x93, 0xb5, 0x8b, 0x2a, 0xc1, 0x00, 0x00,
	0x00,
}