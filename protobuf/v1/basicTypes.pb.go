// Code generated by protoc-gen-go. DO NOT EDIT.
// source: v1/basicTypes.proto

/*
Package org_apache_geode_internal_protocol_protobuf_v1 is a generated protocol buffer package.

It is generated from these files:
	v1/basicTypes.proto
	v1/clientProtocol.proto
	v1/connection_API.proto
	v1/locator_API.proto
	v1/region_API.proto

It has these top-level messages:
	Entry
	EncodedValue
	CustomEncodedValue
	Region
	Server
	Error
	KeyedError
	Message
	Request
	Response
	ErrorResponse
	AuthenticationRequest
	AuthenticationResponse
	GetAvailableServersRequest
	GetAvailableServersResponse
	PutRequest
	PutResponse
	GetRequest
	GetResponse
	PutAllRequest
	PutAllResponse
	GetAllRequest
	GetAllResponse
	RemoveRequest
	RemoveResponse
	GetRegionNamesRequest
	GetRegionNamesResponse
	GetRegionRequest
	GetRegionResponse
*/
package org_apache_geode_internal_protocol_protobuf_v1

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

type EncodingType int32

const (
	EncodingType_INVALID EncodingType = 0
	EncodingType_JSON    EncodingType = 10
)

var EncodingType_name = map[int32]string{
	0:  "INVALID",
	10: "JSON",
}
var EncodingType_value = map[string]int32{
	"INVALID": 0,
	"JSON":    10,
}

func (x EncodingType) String() string {
	return proto.EnumName(EncodingType_name, int32(x))
}
func (EncodingType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Entry struct {
	Key   *EncodedValue `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Value *EncodedValue `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
}

func (m *Entry) Reset()                    { *m = Entry{} }
func (m *Entry) String() string            { return proto.CompactTextString(m) }
func (*Entry) ProtoMessage()               {}
func (*Entry) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Entry) GetKey() *EncodedValue {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *Entry) GetValue() *EncodedValue {
	if m != nil {
		return m.Value
	}
	return nil
}

type EncodedValue struct {
	// Types that are valid to be assigned to Value:
	//	*EncodedValue_IntResult
	//	*EncodedValue_LongResult
	//	*EncodedValue_ShortResult
	//	*EncodedValue_ByteResult
	//	*EncodedValue_BooleanResult
	//	*EncodedValue_DoubleResult
	//	*EncodedValue_FloatResult
	//	*EncodedValue_BinaryResult
	//	*EncodedValue_StringResult
	//	*EncodedValue_CustomEncodedValue
	Value isEncodedValue_Value `protobuf_oneof:"value"`
}

func (m *EncodedValue) Reset()                    { *m = EncodedValue{} }
func (m *EncodedValue) String() string            { return proto.CompactTextString(m) }
func (*EncodedValue) ProtoMessage()               {}
func (*EncodedValue) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type isEncodedValue_Value interface {
	isEncodedValue_Value()
}

type EncodedValue_IntResult struct {
	IntResult int32 `protobuf:"varint,1,opt,name=intResult,oneof"`
}
type EncodedValue_LongResult struct {
	LongResult int64 `protobuf:"varint,2,opt,name=longResult,oneof"`
}
type EncodedValue_ShortResult struct {
	ShortResult int32 `protobuf:"varint,3,opt,name=shortResult,oneof"`
}
type EncodedValue_ByteResult struct {
	ByteResult int32 `protobuf:"varint,4,opt,name=byteResult,oneof"`
}
type EncodedValue_BooleanResult struct {
	BooleanResult bool `protobuf:"varint,5,opt,name=booleanResult,oneof"`
}
type EncodedValue_DoubleResult struct {
	DoubleResult float64 `protobuf:"fixed64,6,opt,name=doubleResult,oneof"`
}
type EncodedValue_FloatResult struct {
	FloatResult float32 `protobuf:"fixed32,7,opt,name=floatResult,oneof"`
}
type EncodedValue_BinaryResult struct {
	BinaryResult []byte `protobuf:"bytes,8,opt,name=binaryResult,proto3,oneof"`
}
type EncodedValue_StringResult struct {
	StringResult string `protobuf:"bytes,9,opt,name=stringResult,oneof"`
}
type EncodedValue_CustomEncodedValue struct {
	CustomEncodedValue *CustomEncodedValue `protobuf:"bytes,50,opt,name=customEncodedValue,oneof"`
}

func (*EncodedValue_IntResult) isEncodedValue_Value()          {}
func (*EncodedValue_LongResult) isEncodedValue_Value()         {}
func (*EncodedValue_ShortResult) isEncodedValue_Value()        {}
func (*EncodedValue_ByteResult) isEncodedValue_Value()         {}
func (*EncodedValue_BooleanResult) isEncodedValue_Value()      {}
func (*EncodedValue_DoubleResult) isEncodedValue_Value()       {}
func (*EncodedValue_FloatResult) isEncodedValue_Value()        {}
func (*EncodedValue_BinaryResult) isEncodedValue_Value()       {}
func (*EncodedValue_StringResult) isEncodedValue_Value()       {}
func (*EncodedValue_CustomEncodedValue) isEncodedValue_Value() {}

func (m *EncodedValue) GetValue() isEncodedValue_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *EncodedValue) GetIntResult() int32 {
	if x, ok := m.GetValue().(*EncodedValue_IntResult); ok {
		return x.IntResult
	}
	return 0
}

func (m *EncodedValue) GetLongResult() int64 {
	if x, ok := m.GetValue().(*EncodedValue_LongResult); ok {
		return x.LongResult
	}
	return 0
}

func (m *EncodedValue) GetShortResult() int32 {
	if x, ok := m.GetValue().(*EncodedValue_ShortResult); ok {
		return x.ShortResult
	}
	return 0
}

func (m *EncodedValue) GetByteResult() int32 {
	if x, ok := m.GetValue().(*EncodedValue_ByteResult); ok {
		return x.ByteResult
	}
	return 0
}

func (m *EncodedValue) GetBooleanResult() bool {
	if x, ok := m.GetValue().(*EncodedValue_BooleanResult); ok {
		return x.BooleanResult
	}
	return false
}

func (m *EncodedValue) GetDoubleResult() float64 {
	if x, ok := m.GetValue().(*EncodedValue_DoubleResult); ok {
		return x.DoubleResult
	}
	return 0
}

func (m *EncodedValue) GetFloatResult() float32 {
	if x, ok := m.GetValue().(*EncodedValue_FloatResult); ok {
		return x.FloatResult
	}
	return 0
}

func (m *EncodedValue) GetBinaryResult() []byte {
	if x, ok := m.GetValue().(*EncodedValue_BinaryResult); ok {
		return x.BinaryResult
	}
	return nil
}

func (m *EncodedValue) GetStringResult() string {
	if x, ok := m.GetValue().(*EncodedValue_StringResult); ok {
		return x.StringResult
	}
	return ""
}

func (m *EncodedValue) GetCustomEncodedValue() *CustomEncodedValue {
	if x, ok := m.GetValue().(*EncodedValue_CustomEncodedValue); ok {
		return x.CustomEncodedValue
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*EncodedValue) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _EncodedValue_OneofMarshaler, _EncodedValue_OneofUnmarshaler, _EncodedValue_OneofSizer, []interface{}{
		(*EncodedValue_IntResult)(nil),
		(*EncodedValue_LongResult)(nil),
		(*EncodedValue_ShortResult)(nil),
		(*EncodedValue_ByteResult)(nil),
		(*EncodedValue_BooleanResult)(nil),
		(*EncodedValue_DoubleResult)(nil),
		(*EncodedValue_FloatResult)(nil),
		(*EncodedValue_BinaryResult)(nil),
		(*EncodedValue_StringResult)(nil),
		(*EncodedValue_CustomEncodedValue)(nil),
	}
}

func _EncodedValue_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*EncodedValue)
	// value
	switch x := m.Value.(type) {
	case *EncodedValue_IntResult:
		b.EncodeVarint(1<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.IntResult))
	case *EncodedValue_LongResult:
		b.EncodeVarint(2<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.LongResult))
	case *EncodedValue_ShortResult:
		b.EncodeVarint(3<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.ShortResult))
	case *EncodedValue_ByteResult:
		b.EncodeVarint(4<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.ByteResult))
	case *EncodedValue_BooleanResult:
		t := uint64(0)
		if x.BooleanResult {
			t = 1
		}
		b.EncodeVarint(5<<3 | proto.WireVarint)
		b.EncodeVarint(t)
	case *EncodedValue_DoubleResult:
		b.EncodeVarint(6<<3 | proto.WireFixed64)
		b.EncodeFixed64(math.Float64bits(x.DoubleResult))
	case *EncodedValue_FloatResult:
		b.EncodeVarint(7<<3 | proto.WireFixed32)
		b.EncodeFixed32(uint64(math.Float32bits(x.FloatResult)))
	case *EncodedValue_BinaryResult:
		b.EncodeVarint(8<<3 | proto.WireBytes)
		b.EncodeRawBytes(x.BinaryResult)
	case *EncodedValue_StringResult:
		b.EncodeVarint(9<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.StringResult)
	case *EncodedValue_CustomEncodedValue:
		b.EncodeVarint(50<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.CustomEncodedValue); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("EncodedValue.Value has unexpected type %T", x)
	}
	return nil
}

func _EncodedValue_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*EncodedValue)
	switch tag {
	case 1: // value.intResult
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Value = &EncodedValue_IntResult{int32(x)}
		return true, err
	case 2: // value.longResult
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Value = &EncodedValue_LongResult{int64(x)}
		return true, err
	case 3: // value.shortResult
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Value = &EncodedValue_ShortResult{int32(x)}
		return true, err
	case 4: // value.byteResult
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Value = &EncodedValue_ByteResult{int32(x)}
		return true, err
	case 5: // value.booleanResult
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Value = &EncodedValue_BooleanResult{x != 0}
		return true, err
	case 6: // value.doubleResult
		if wire != proto.WireFixed64 {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeFixed64()
		m.Value = &EncodedValue_DoubleResult{math.Float64frombits(x)}
		return true, err
	case 7: // value.floatResult
		if wire != proto.WireFixed32 {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeFixed32()
		m.Value = &EncodedValue_FloatResult{math.Float32frombits(uint32(x))}
		return true, err
	case 8: // value.binaryResult
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeRawBytes(true)
		m.Value = &EncodedValue_BinaryResult{x}
		return true, err
	case 9: // value.stringResult
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Value = &EncodedValue_StringResult{x}
		return true, err
	case 50: // value.customEncodedValue
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(CustomEncodedValue)
		err := b.DecodeMessage(msg)
		m.Value = &EncodedValue_CustomEncodedValue{msg}
		return true, err
	default:
		return false, nil
	}
}

func _EncodedValue_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*EncodedValue)
	// value
	switch x := m.Value.(type) {
	case *EncodedValue_IntResult:
		n += proto.SizeVarint(1<<3 | proto.WireVarint)
		n += proto.SizeVarint(uint64(x.IntResult))
	case *EncodedValue_LongResult:
		n += proto.SizeVarint(2<<3 | proto.WireVarint)
		n += proto.SizeVarint(uint64(x.LongResult))
	case *EncodedValue_ShortResult:
		n += proto.SizeVarint(3<<3 | proto.WireVarint)
		n += proto.SizeVarint(uint64(x.ShortResult))
	case *EncodedValue_ByteResult:
		n += proto.SizeVarint(4<<3 | proto.WireVarint)
		n += proto.SizeVarint(uint64(x.ByteResult))
	case *EncodedValue_BooleanResult:
		n += proto.SizeVarint(5<<3 | proto.WireVarint)
		n += 1
	case *EncodedValue_DoubleResult:
		n += proto.SizeVarint(6<<3 | proto.WireFixed64)
		n += 8
	case *EncodedValue_FloatResult:
		n += proto.SizeVarint(7<<3 | proto.WireFixed32)
		n += 4
	case *EncodedValue_BinaryResult:
		n += proto.SizeVarint(8<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.BinaryResult)))
		n += len(x.BinaryResult)
	case *EncodedValue_StringResult:
		n += proto.SizeVarint(9<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.StringResult)))
		n += len(x.StringResult)
	case *EncodedValue_CustomEncodedValue:
		s := proto.Size(x.CustomEncodedValue)
		n += proto.SizeVarint(50<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type CustomEncodedValue struct {
	EncodingType EncodingType `protobuf:"varint,1,opt,name=encodingType,enum=org.apache.geode.internal.protocol.protobuf.v1.EncodingType" json:"encodingType,omitempty"`
	Value        []byte       `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *CustomEncodedValue) Reset()                    { *m = CustomEncodedValue{} }
func (m *CustomEncodedValue) String() string            { return proto.CompactTextString(m) }
func (*CustomEncodedValue) ProtoMessage()               {}
func (*CustomEncodedValue) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *CustomEncodedValue) GetEncodingType() EncodingType {
	if m != nil {
		return m.EncodingType
	}
	return EncodingType_INVALID
}

func (m *CustomEncodedValue) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

type Region struct {
	Name            string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	DataPolicy      string `protobuf:"bytes,2,opt,name=dataPolicy" json:"dataPolicy,omitempty"`
	Scope           string `protobuf:"bytes,3,opt,name=scope" json:"scope,omitempty"`
	KeyConstraint   string `protobuf:"bytes,4,opt,name=keyConstraint" json:"keyConstraint,omitempty"`
	ValueConstraint string `protobuf:"bytes,5,opt,name=valueConstraint" json:"valueConstraint,omitempty"`
	Persisted       bool   `protobuf:"varint,6,opt,name=persisted" json:"persisted,omitempty"`
	Size            int64  `protobuf:"varint,7,opt,name=size" json:"size,omitempty"`
}

func (m *Region) Reset()                    { *m = Region{} }
func (m *Region) String() string            { return proto.CompactTextString(m) }
func (*Region) ProtoMessage()               {}
func (*Region) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Region) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Region) GetDataPolicy() string {
	if m != nil {
		return m.DataPolicy
	}
	return ""
}

func (m *Region) GetScope() string {
	if m != nil {
		return m.Scope
	}
	return ""
}

func (m *Region) GetKeyConstraint() string {
	if m != nil {
		return m.KeyConstraint
	}
	return ""
}

func (m *Region) GetValueConstraint() string {
	if m != nil {
		return m.ValueConstraint
	}
	return ""
}

func (m *Region) GetPersisted() bool {
	if m != nil {
		return m.Persisted
	}
	return false
}

func (m *Region) GetSize() int64 {
	if m != nil {
		return m.Size
	}
	return 0
}

type Server struct {
	Hostname string `protobuf:"bytes,1,opt,name=hostname" json:"hostname,omitempty"`
	Port     int32  `protobuf:"varint,2,opt,name=port" json:"port,omitempty"`
}

func (m *Server) Reset()                    { *m = Server{} }
func (m *Server) String() string            { return proto.CompactTextString(m) }
func (*Server) ProtoMessage()               {}
func (*Server) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Server) GetHostname() string {
	if m != nil {
		return m.Hostname
	}
	return ""
}

func (m *Server) GetPort() int32 {
	if m != nil {
		return m.Port
	}
	return 0
}

type Error struct {
	ErrorCode int32  `protobuf:"varint,1,opt,name=errorCode" json:"errorCode,omitempty"`
	Message   string `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
}

func (m *Error) Reset()                    { *m = Error{} }
func (m *Error) String() string            { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()               {}
func (*Error) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *Error) GetErrorCode() int32 {
	if m != nil {
		return m.ErrorCode
	}
	return 0
}

func (m *Error) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type KeyedError struct {
	Key   *EncodedValue `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Error *Error        `protobuf:"bytes,2,opt,name=error" json:"error,omitempty"`
}

func (m *KeyedError) Reset()                    { *m = KeyedError{} }
func (m *KeyedError) String() string            { return proto.CompactTextString(m) }
func (*KeyedError) ProtoMessage()               {}
func (*KeyedError) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *KeyedError) GetKey() *EncodedValue {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *KeyedError) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func init() {
	proto.RegisterType((*Entry)(nil), "org.apache.geode.internal.protocol.protobuf.v1.Entry")
	proto.RegisterType((*EncodedValue)(nil), "org.apache.geode.internal.protocol.protobuf.v1.EncodedValue")
	proto.RegisterType((*CustomEncodedValue)(nil), "org.apache.geode.internal.protocol.protobuf.v1.CustomEncodedValue")
	proto.RegisterType((*Region)(nil), "org.apache.geode.internal.protocol.protobuf.v1.Region")
	proto.RegisterType((*Server)(nil), "org.apache.geode.internal.protocol.protobuf.v1.Server")
	proto.RegisterType((*Error)(nil), "org.apache.geode.internal.protocol.protobuf.v1.Error")
	proto.RegisterType((*KeyedError)(nil), "org.apache.geode.internal.protocol.protobuf.v1.KeyedError")
	proto.RegisterEnum("org.apache.geode.internal.protocol.protobuf.v1.EncodingType", EncodingType_name, EncodingType_value)
}

func init() { proto.RegisterFile("v1/basicTypes.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 584 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x94, 0xd1, 0x6e, 0xd3, 0x3e,
	0x14, 0xc6, 0x9b, 0x75, 0x69, 0x9b, 0xb3, 0xee, 0xff, 0x9f, 0x0c, 0x17, 0x15, 0x42, 0x53, 0x14,
	0x0d, 0x14, 0x71, 0x11, 0xb4, 0x21, 0x24, 0x2e, 0x90, 0x10, 0x1b, 0x93, 0x3a, 0x86, 0x06, 0xf2,
	0xd0, 0xae, 0x71, 0x92, 0xb3, 0x2e, 0x5a, 0x66, 0x57, 0xb6, 0x5b, 0x29, 0x3c, 0x03, 0x2f, 0x82,
	0x78, 0x1c, 0x78, 0x20, 0x64, 0xc7, 0xd9, 0x5c, 0xc6, 0xcd, 0xd0, 0xee, 0xce, 0xf9, 0xf4, 0x9d,
	0x5f, 0x4e, 0xed, 0xaf, 0x86, 0x07, 0xcb, 0xdd, 0xe7, 0x39, 0x53, 0x55, 0xf1, 0xb9, 0x99, 0xa3,
	0xca, 0xe6, 0x52, 0x68, 0x41, 0x32, 0x21, 0x67, 0x19, 0x9b, 0xb3, 0xe2, 0x02, 0xb3, 0x19, 0x8a,
	0x12, 0xb3, 0x8a, 0x6b, 0x94, 0x9c, 0xd5, 0xad, 0xa1, 0x10, 0xae, 0xc8, 0x17, 0xe7, 0xd9, 0x72,
	0x37, 0xf9, 0x11, 0x40, 0x78, 0xc8, 0xb5, 0x6c, 0xc8, 0x09, 0xf4, 0x2f, 0xb1, 0x99, 0x04, 0x71,
	0x90, 0x6e, 0xec, 0xbd, 0xbe, 0x23, 0x27, 0x3b, 0xe4, 0x85, 0x28, 0xb1, 0x3c, 0x63, 0xf5, 0x02,
	0xa9, 0x01, 0x11, 0x0a, 0xe1, 0xd2, 0x74, 0x93, 0xb5, 0x7b, 0x20, 0xb6, 0xa8, 0xe4, 0x67, 0x1f,
	0xc6, 0xbe, 0x4e, 0xb6, 0x21, 0xaa, 0xb8, 0xa6, 0xa8, 0x16, 0xb5, 0xb6, 0xab, 0x87, 0xd3, 0x1e,
	0xbd, 0x91, 0x48, 0x0c, 0x50, 0x0b, 0x3e, 0x73, 0x06, 0xb3, 0x49, 0x7f, 0xda, 0xa3, 0x9e, 0x46,
	0x12, 0xd8, 0x50, 0x17, 0x42, 0x76, 0x8c, 0xbe, 0x63, 0xf8, 0xa2, 0xa1, 0xe4, 0x8d, 0x46, 0x67,
	0x59, 0x77, 0x16, 0x4f, 0x23, 0x4f, 0x61, 0x33, 0x17, 0xa2, 0x46, 0xc6, 0x9d, 0x29, 0x8c, 0x83,
	0x74, 0x34, 0xed, 0xd1, 0x55, 0x99, 0xec, 0xc0, 0xb8, 0x14, 0x8b, 0xbc, 0xee, 0x58, 0x83, 0x38,
	0x48, 0x83, 0x69, 0x8f, 0xae, 0xa8, 0x66, 0xa7, 0xf3, 0x5a, 0xb0, 0x6e, 0xa7, 0x61, 0x1c, 0xa4,
	0x6b, 0x66, 0x27, 0x4f, 0x34, 0xa4, 0xbc, 0xe2, 0x4c, 0x36, 0xce, 0x34, 0x8a, 0x83, 0x74, 0x6c,
	0x48, 0xbe, 0x6a, 0x5c, 0x4a, 0xcb, 0xea, 0xfa, 0x04, 0xa2, 0x38, 0x48, 0x23, 0xe3, 0xf2, 0x55,
	0xa2, 0x81, 0x14, 0x0b, 0xa5, 0xc5, 0x95, 0x7f, 0xb6, 0x93, 0x3d, 0x7b, 0x6f, 0xfb, 0x77, 0xbd,
	0xb7, 0x83, 0x5b, 0xa4, 0x69, 0x8f, 0xfe, 0x85, 0xbf, 0x3f, 0x74, 0x01, 0x49, 0xbe, 0x05, 0x40,
	0x6e, 0x4f, 0x91, 0x2f, 0x30, 0x46, 0xd3, 0x57, 0x7c, 0x66, 0x12, 0x6e, 0xaf, 0xf7, 0xbf, 0x7f,
	0xcc, 0x91, 0x63, 0xd0, 0x15, 0x22, 0x79, 0xe8, 0x47, 0x74, 0xdc, 0x85, 0xec, 0x57, 0x00, 0x03,
	0x8a, 0xb3, 0x4a, 0x70, 0x42, 0x60, 0x9d, 0xb3, 0xab, 0xf6, 0xd3, 0x11, 0xb5, 0x35, 0xd9, 0x06,
	0x28, 0x99, 0x66, 0x9f, 0x44, 0x5d, 0x15, 0x8d, 0x9d, 0x8c, 0xa8, 0xa7, 0x18, 0xa8, 0x2a, 0xc4,
	0x1c, 0x6d, 0x94, 0x22, 0xda, 0x36, 0x64, 0x07, 0x36, 0x2f, 0xb1, 0x39, 0x10, 0x5c, 0x69, 0xc9,
	0x2a, 0xde, 0xa6, 0x28, 0xa2, 0xab, 0x22, 0x49, 0xe1, 0x7f, 0xbb, 0x83, 0xe7, 0x0b, 0xad, 0xef,
	0x4f, 0x99, 0x3c, 0x86, 0x68, 0x8e, 0x52, 0x55, 0x4a, 0x63, 0x69, 0x53, 0x34, 0xa2, 0x37, 0x82,
	0xd9, 0x5b, 0x55, 0x5f, 0xd1, 0x26, 0xa7, 0x4f, 0x6d, 0x9d, 0xbc, 0x82, 0xc1, 0x29, 0xca, 0x25,
	0x4a, 0xf2, 0x08, 0x46, 0x17, 0x42, 0x69, 0xef, 0x97, 0x5d, 0xf7, 0x66, 0x72, 0x2e, 0x64, 0xfb,
	0x57, 0x09, 0xa9, 0xad, 0x93, 0x37, 0x10, 0x1e, 0x4a, 0x29, 0xa4, 0xf9, 0x28, 0x9a, 0xe2, 0x40,
	0x94, 0xed, 0x64, 0x48, 0x6f, 0x04, 0x32, 0x81, 0xe1, 0x15, 0x2a, 0xc5, 0x66, 0xe8, 0x4e, 0xa5,
	0x6b, 0x93, 0xef, 0x01, 0xc0, 0x31, 0x36, 0x58, 0xb6, 0x98, 0xfb, 0x7e, 0x69, 0x8e, 0x21, 0xb4,
	0x5b, 0xb8, 0x97, 0xe6, 0xe5, 0x9d, 0x89, 0x66, 0x98, 0xb6, 0x8c, 0x67, 0x4f, 0xdc, 0x0b, 0xd3,
	0x65, 0x64, 0x03, 0x86, 0x47, 0x27, 0x67, 0x6f, 0x3f, 0x1c, 0xbd, 0xdb, 0xea, 0x91, 0x11, 0xac,
	0xbf, 0x3f, 0xfd, 0x78, 0xb2, 0x05, 0xf9, 0xc0, 0x02, 0x5e, 0xfc, 0x0e, 0x00, 0x00, 0xff, 0xff,
	0xcd, 0x3c, 0x4e, 0xca, 0x85, 0x05, 0x00, 0x00,
}
