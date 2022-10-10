// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cleaning_code.proto

package cleaningcode

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

type CleaningCode struct {
	// @gotags: dynamo:"CleaningCodeID"
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" dynamo:"CleaningCodeID"`
	// @gotags: dynamo:"CleaningCodeDescription"
	Description          string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty" dynamo:"CleaningCodeDescription"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CleaningCode) Reset()         { *m = CleaningCode{} }
func (m *CleaningCode) String() string { return proto.CompactTextString(m) }
func (*CleaningCode) ProtoMessage()    {}
func (*CleaningCode) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b3143c046fea01e, []int{0}
}

func (m *CleaningCode) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CleaningCode.Unmarshal(m, b)
}
func (m *CleaningCode) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CleaningCode.Marshal(b, m, deterministic)
}
func (m *CleaningCode) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CleaningCode.Merge(m, src)
}
func (m *CleaningCode) XXX_Size() int {
	return xxx_messageInfo_CleaningCode.Size(m)
}
func (m *CleaningCode) XXX_DiscardUnknown() {
	xxx_messageInfo_CleaningCode.DiscardUnknown(m)
}

var xxx_messageInfo_CleaningCode proto.InternalMessageInfo

func (m *CleaningCode) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *CleaningCode) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func init() {
	proto.RegisterType((*CleaningCode)(nil), "magnoliafabrics.CleaningCode")
}

func init() {
	proto.RegisterFile("cleaning_code.proto", fileDescriptor_9b3143c046fea01e)
}

var fileDescriptor_9b3143c046fea01e = []byte{
	// 164 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4e, 0xce, 0x49, 0x4d,
	0xcc, 0xcb, 0xcc, 0x4b, 0x8f, 0x4f, 0xce, 0x4f, 0x49, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0xe2, 0xcf, 0x4d, 0x4c, 0xcf, 0xcb, 0xcf, 0xc9, 0x4c, 0x4c, 0x4b, 0x4c, 0x2a, 0xca, 0x4c, 0x2e,
	0x56, 0x72, 0xe0, 0xe2, 0x71, 0x86, 0xaa, 0x73, 0xce, 0x4f, 0x49, 0x15, 0xe2, 0xe3, 0x62, 0xca,
	0x4c, 0x91, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x62, 0xca, 0x4c, 0x11, 0x52, 0xe0, 0xe2, 0x4e,
	0x49, 0x2d, 0x4e, 0x2e, 0xca, 0x2c, 0x28, 0xc9, 0xcc, 0xcf, 0x93, 0x60, 0x02, 0x4b, 0x20, 0x0b,
	0x39, 0x39, 0x46, 0xd9, 0xa7, 0x67, 0x96, 0x64, 0x94, 0x26, 0xe9, 0x25, 0xe7, 0xe7, 0xea, 0x87,
	0x65, 0x96, 0x24, 0xe6, 0x64, 0x96, 0x54, 0xea, 0x06, 0xe7, 0x97, 0x96, 0x64, 0xe8, 0xc3, 0xac,
	0xd3, 0x85, 0xda, 0xa7, 0x9b, 0x58, 0x90, 0xa9, 0x5f, 0x90, 0x9d, 0xae, 0x0f, 0x73, 0x1c, 0xc8,
	0x6d, 0x49, 0x6c, 0x60, 0xc7, 0x19, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0xba, 0xbb, 0xef, 0x5b,
	0xb3, 0x00, 0x00, 0x00,
}
