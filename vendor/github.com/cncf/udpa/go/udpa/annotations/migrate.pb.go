// Code generated by protoc-gen-go. DO NOT EDIT.
// source: udpa/annotations/migrate.proto

package udpa_annotations

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	descriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"
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

type MigrateAnnotation struct {
	Rename               string   `protobuf:"bytes,1,opt,name=rename,proto3" json:"rename,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MigrateAnnotation) Reset()         { *m = MigrateAnnotation{} }
func (m *MigrateAnnotation) String() string { return proto.CompactTextString(m) }
func (*MigrateAnnotation) ProtoMessage()    {}
func (*MigrateAnnotation) Descriptor() ([]byte, []int) {
	return fileDescriptor_ba8191732d0e246d, []int{0}
}

func (m *MigrateAnnotation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MigrateAnnotation.Unmarshal(m, b)
}
func (m *MigrateAnnotation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MigrateAnnotation.Marshal(b, m, deterministic)
}
func (m *MigrateAnnotation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MigrateAnnotation.Merge(m, src)
}
func (m *MigrateAnnotation) XXX_Size() int {
	return xxx_messageInfo_MigrateAnnotation.Size(m)
}
func (m *MigrateAnnotation) XXX_DiscardUnknown() {
	xxx_messageInfo_MigrateAnnotation.DiscardUnknown(m)
}

var xxx_messageInfo_MigrateAnnotation proto.InternalMessageInfo

func (m *MigrateAnnotation) GetRename() string {
	if m != nil {
		return m.Rename
	}
	return ""
}

var E_MessageMigrate = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*MigrateAnnotation)(nil),
	Field:         171962766,
	Name:          "udpa.annotations.message_migrate",
	Tag:           "bytes,171962766,opt,name=message_migrate",
	Filename:      "udpa/annotations/migrate.proto",
}

var E_FieldMigrate = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*MigrateAnnotation)(nil),
	Field:         171962766,
	Name:          "udpa.annotations.field_migrate",
	Tag:           "bytes,171962766,opt,name=field_migrate",
	Filename:      "udpa/annotations/migrate.proto",
}

var E_EnumMigrate = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.EnumOptions)(nil),
	ExtensionType: (*MigrateAnnotation)(nil),
	Field:         171962766,
	Name:          "udpa.annotations.enum_migrate",
	Tag:           "bytes,171962766,opt,name=enum_migrate",
	Filename:      "udpa/annotations/migrate.proto",
}

var E_EnumValueMigrate = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.EnumValueOptions)(nil),
	ExtensionType: (*MigrateAnnotation)(nil),
	Field:         171962766,
	Name:          "udpa.annotations.enum_value_migrate",
	Tag:           "bytes,171962766,opt,name=enum_value_migrate",
	Filename:      "udpa/annotations/migrate.proto",
}

func init() {
	proto.RegisterType((*MigrateAnnotation)(nil), "udpa.annotations.MigrateAnnotation")
	proto.RegisterExtension(E_MessageMigrate)
	proto.RegisterExtension(E_FieldMigrate)
	proto.RegisterExtension(E_EnumMigrate)
	proto.RegisterExtension(E_EnumValueMigrate)
}

func init() { proto.RegisterFile("udpa/annotations/migrate.proto", fileDescriptor_ba8191732d0e246d) }

var fileDescriptor_ba8191732d0e246d = []byte{
	// 253 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0xd0, 0xbb, 0x4a, 0x03, 0x41,
	0x14, 0x06, 0x60, 0xb6, 0x09, 0x38, 0x89, 0x1a, 0xb7, 0x90, 0x20, 0x5e, 0x56, 0x6d, 0x02, 0xc2,
	0x2c, 0x68, 0x97, 0xce, 0x42, 0xbb, 0x20, 0xa6, 0xb0, 0x95, 0x89, 0x7b, 0x76, 0x18, 0xd9, 0xb9,
	0x30, 0x17, 0x7d, 0x0b, 0x5f, 0xd2, 0x07, 0x51, 0xe6, 0xb6, 0x8a, 0x6b, 0xb5, 0xe5, 0xf0, 0x9f,
	0xf3, 0x7f, 0x87, 0x41, 0xa7, 0xae, 0x51, 0xa4, 0x26, 0x42, 0x48, 0x4b, 0x2c, 0x93, 0xc2, 0xd4,
	0x9c, 0x51, 0x4d, 0x2c, 0x60, 0xa5, 0xa5, 0x95, 0xe5, 0xdc, 0xe7, 0xf8, 0x57, 0x7e, 0x54, 0x51,
	0x29, 0x69, 0x07, 0x75, 0xc8, 0xb7, 0xae, 0xad, 0x1b, 0x30, 0x2f, 0x9a, 0x29, 0x2b, 0x75, 0xdc,
	0xb9, 0xb8, 0x42, 0x07, 0xeb, 0x58, 0x72, 0xdb, 0xef, 0x95, 0x87, 0x68, 0xa2, 0x41, 0x10, 0x0e,
	0x8b, 0xa2, 0x2a, 0x96, 0x3b, 0x9b, 0xf4, 0x5a, 0x29, 0xb4, 0xcf, 0xc1, 0x18, 0x42, 0xe1, 0x39,
	0xc9, 0xe5, 0x19, 0x8e, 0x04, 0xce, 0x04, 0x5e, 0xc7, 0x89, 0x07, 0x15, 0x4e, 0x58, 0x7c, 0x7c,
	0x7e, 0x3d, 0x56, 0xc5, 0x72, 0x7a, 0x7d, 0x89, 0xff, 0x9e, 0x87, 0x07, 0xf2, 0x66, 0x2f, 0xf5,
	0xa7, 0x64, 0xf5, 0x8a, 0x76, 0x5b, 0x06, 0x5d, 0xd3, 0x7b, 0x27, 0x03, 0xef, 0xde, 0xe7, 0xe3,
	0xb4, 0x59, 0xe8, 0xce, 0x16, 0x45, 0x33, 0x10, 0x8e, 0xf7, 0xd4, 0xf1, 0x80, 0xba, 0x13, 0x8e,
	0x8f, 0x93, 0xa6, 0xbe, 0x39, 0x43, 0xef, 0xa8, 0x0c, 0xd0, 0x1b, 0xe9, 0xdc, 0xcf, 0x4f, 0x9e,
	0xff, 0xcb, 0x3d, 0xf9, 0x99, 0x71, 0xe6, 0x1c, 0xf2, 0x7e, 0xca, 0xb6, 0x93, 0x50, 0x7d, 0xf3,
	0x1d, 0x00, 0x00, 0xff, 0xff, 0xe7, 0xb5, 0x81, 0x66, 0x49, 0x02, 0x00, 0x00,
}
