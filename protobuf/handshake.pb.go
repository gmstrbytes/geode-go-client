// Code generated by protoc-gen-go. DO NOT EDIT.
// source: handshake.proto

/*
Package org_apache_geode_internal_protocol_protobuf is a generated protocol buffer package.

It is generated from these files:
	handshake.proto

It has these top-level messages:
	NewConnectionHandshake
	HandshakeAcknowledgement
*/
package org_apache_geode_internal_protocol_protobuf

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type MajorVersions int32

const (
	MajorVersions_INVALID_MAJOR_VERSION MajorVersions = 0
	MajorVersions_CURRENT_MAJOR_VERSION MajorVersions = 1
)

var MajorVersions_name = map[int32]string{
	0: "INVALID_MAJOR_VERSION",
	1: "CURRENT_MAJOR_VERSION",
}
var MajorVersions_value = map[string]int32{
	"INVALID_MAJOR_VERSION": 0,
	"CURRENT_MAJOR_VERSION": 1,
}

func (x MajorVersions) String() string {
	return proto.EnumName(MajorVersions_name, int32(x))
}
func (MajorVersions) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type MinorVersions int32

const (
	MinorVersions_INVALID_MINOR_VERSION MinorVersions = 0
	MinorVersions_CURRENT_MINOR_VERSION MinorVersions = 1
)

var MinorVersions_name = map[int32]string{
	0: "INVALID_MINOR_VERSION",
	1: "CURRENT_MINOR_VERSION",
}
var MinorVersions_value = map[string]int32{
	"INVALID_MINOR_VERSION": 0,
	"CURRENT_MINOR_VERSION": 1,
}

func (x MinorVersions) String() string {
	return proto.EnumName(MinorVersions_name, int32(x))
}
func (MinorVersions) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type NewConnectionHandshake struct {
	MajorVersion uint32 `protobuf:"fixed32,1,opt,name=majorVersion" json:"majorVersion,omitempty"`
	MinorVersion uint32 `protobuf:"fixed32,2,opt,name=minorVersion" json:"minorVersion,omitempty"`
}

func (m *NewConnectionHandshake) Reset()                    { *m = NewConnectionHandshake{} }
func (m *NewConnectionHandshake) String() string            { return proto.CompactTextString(m) }
func (*NewConnectionHandshake) ProtoMessage()               {}
func (*NewConnectionHandshake) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *NewConnectionHandshake) GetMajorVersion() uint32 {
	if m != nil {
		return m.MajorVersion
	}
	return 0
}

func (m *NewConnectionHandshake) GetMinorVersion() uint32 {
	if m != nil {
		return m.MinorVersion
	}
	return 0
}

type HandshakeAcknowledgement struct {
	ServerMajorVersion int32 `protobuf:"varint,1,opt,name=serverMajorVersion" json:"serverMajorVersion,omitempty"`
	ServerMinorVersion int32 `protobuf:"varint,2,opt,name=serverMinorVersion" json:"serverMinorVersion,omitempty"`
	HandshakePassed    bool  `protobuf:"varint,3,opt,name=handshakePassed" json:"handshakePassed,omitempty"`
}

func (m *HandshakeAcknowledgement) Reset()                    { *m = HandshakeAcknowledgement{} }
func (m *HandshakeAcknowledgement) String() string            { return proto.CompactTextString(m) }
func (*HandshakeAcknowledgement) ProtoMessage()               {}
func (*HandshakeAcknowledgement) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *HandshakeAcknowledgement) GetServerMajorVersion() int32 {
	if m != nil {
		return m.ServerMajorVersion
	}
	return 0
}

func (m *HandshakeAcknowledgement) GetServerMinorVersion() int32 {
	if m != nil {
		return m.ServerMinorVersion
	}
	return 0
}

func (m *HandshakeAcknowledgement) GetHandshakePassed() bool {
	if m != nil {
		return m.HandshakePassed
	}
	return false
}

func init() {
	proto.RegisterType((*NewConnectionHandshake)(nil), "org.apache.geode.internal.protocol.protobuf.NewConnectionHandshake")
	proto.RegisterType((*HandshakeAcknowledgement)(nil), "org.apache.geode.internal.protocol.protobuf.HandshakeAcknowledgement")
	proto.RegisterEnum("org.apache.geode.internal.protocol.protobuf.MajorVersions", MajorVersions_name, MajorVersions_value)
	proto.RegisterEnum("org.apache.geode.internal.protocol.protobuf.MinorVersions", MinorVersions_name, MinorVersions_value)
}

func init() { proto.RegisterFile("handshake.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 260 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0xdf, 0x4a, 0xc3, 0x30,
	0x14, 0xc6, 0x8d, 0xe2, 0x1f, 0x82, 0xa2, 0x04, 0x94, 0x7a, 0x37, 0x7a, 0x55, 0x14, 0x72, 0xe3,
	0x13, 0x94, 0x59, 0xb0, 0xe2, 0x32, 0x89, 0xda, 0xdb, 0x99, 0xb5, 0xc7, 0xb6, 0x6e, 0x3b, 0x67,
	0x24, 0xd5, 0xbd, 0x90, 0x0f, 0x2a, 0xd6, 0xe9, 0xd2, 0xa2, 0x77, 0xe1, 0xfb, 0x7e, 0xfc, 0xbe,
	0x90, 0xf0, 0xe3, 0xca, 0x60, 0xe1, 0x2a, 0x33, 0x03, 0xb9, 0xb4, 0xd4, 0x90, 0xb8, 0x24, 0x5b,
	0x4a, 0xb3, 0x34, 0x79, 0x05, 0xb2, 0x04, 0x2a, 0x40, 0xd6, 0xd8, 0x80, 0x45, 0x33, 0xff, 0x06,
	0x72, 0x5a, 0x1f, 0xa6, 0x6f, 0x2f, 0xe1, 0x33, 0x3f, 0x53, 0xb0, 0x1a, 0x12, 0x22, 0xe4, 0x4d,
	0x4d, 0x78, 0xf3, 0x23, 0x13, 0x21, 0x3f, 0x5c, 0x98, 0x57, 0xb2, 0x19, 0x58, 0x57, 0x13, 0x06,
	0x6c, 0xc0, 0xa2, 0x7d, 0xdd, 0xc9, 0x5a, 0xa6, 0xc6, 0x0d, 0xb3, 0xbd, 0x66, 0xbc, 0x2c, 0xfc,
	0x60, 0x3c, 0xf8, 0xb5, 0xc6, 0xf9, 0x0c, 0x69, 0x35, 0x87, 0xa2, 0x84, 0x05, 0x60, 0x23, 0x24,
	0x17, 0x0e, 0xec, 0x3b, 0xd8, 0x51, 0x7f, 0x6a, 0x57, 0xff, 0xd1, 0x78, 0x7c, 0x7f, 0x76, 0xc3,
	0x7b, 0x8d, 0x88, 0xbc, 0xe7, 0xb9, 0x37, 0xce, 0x41, 0x11, 0xec, 0x0c, 0x58, 0x74, 0xa0, 0xfb,
	0xf1, 0x45, 0xc2, 0x8f, 0xfc, 0x25, 0x27, 0xce, 0xf9, 0x69, 0xaa, 0xb2, 0xf8, 0x2e, 0xbd, 0x9e,
	0x8c, 0xe2, 0xdb, 0xb1, 0x9e, 0x64, 0x89, 0x7e, 0x48, 0xc7, 0xea, 0x64, 0xeb, 0xab, 0x1a, 0x3e,
	0x69, 0x9d, 0xa8, 0xc7, 0x5e, 0xc5, 0x5a, 0x8d, 0x77, 0x81, 0xae, 0x26, 0x55, 0xff, 0x6a, 0x3a,
	0x15, 0x9b, 0xee, 0xb5, 0x1f, 0x74, 0xf5, 0x19, 0x00, 0x00, 0xff, 0xff, 0x07, 0x77, 0x94, 0xa8,
	0xdd, 0x01, 0x00, 0x00,
}