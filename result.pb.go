// Code generated by protoc-gen-go.
// source: github.com/marcushines/openresult/proto/result.proto
// DO NOT EDIT!

/*
Package openresult is a generated protocol buffer package.

It is generated from these files:
	github.com/marcushines/openresult/proto/result.proto

It has these top-level messages:
	Value
	Path
	Report
	Suite
	Case
	Result
*/
package openresult

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/any"
import openresult1 "github.com/marcushines/openresult/proto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Operation int32

const (
	Operation_UNKNOWN            Operation = 0
	Operation_EQUAL              Operation = 1
	Operation_NOT_EQUAL          Operation = 2
	Operation_GREATER_THAN       Operation = 3
	Operation_LESS_THAN          Operation = 4
	Operation_GREATER_THAN_EQUAL Operation = 5
	Operation_LESS_THAN_EQUAL    Operation = 6
	Operation_IN                 Operation = 7
	Operation_NOT_IN             Operation = 8
)

var Operation_name = map[int32]string{
	0: "UNKNOWN",
	1: "EQUAL",
	2: "NOT_EQUAL",
	3: "GREATER_THAN",
	4: "LESS_THAN",
	5: "GREATER_THAN_EQUAL",
	6: "LESS_THAN_EQUAL",
	7: "IN",
	8: "NOT_IN",
}
var Operation_value = map[string]int32{
	"UNKNOWN":            0,
	"EQUAL":              1,
	"NOT_EQUAL":          2,
	"GREATER_THAN":       3,
	"LESS_THAN":          4,
	"GREATER_THAN_EQUAL": 5,
	"LESS_THAN_EQUAL":    6,
	"IN":                 7,
	"NOT_IN":             8,
}

func (x Operation) String() string {
	return proto.EnumName(Operation_name, int32(x))
}
func (Operation) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// Value stores the scalar value of different understood types by the Report.
type Value struct {
	// Types that are valid to be assigned to Value:
	//	*Value_IntValue
	//	*Value_DoubleValue
	//	*Value_StringValue
	//	*Value_BoolValue
	//	*Value_ByteValue
	//	*Value_AnyValue
	//	*Value_JsonValue
	Value isValue_Value `protobuf_oneof:"value"`
}

func (m *Value) Reset()                    { *m = Value{} }
func (m *Value) String() string            { return proto.CompactTextString(m) }
func (*Value) ProtoMessage()               {}
func (*Value) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type isValue_Value interface {
	isValue_Value()
}

type Value_IntValue struct {
	IntValue int64 `protobuf:"varint,1,opt,name=int_value,json=intValue,oneof"`
}
type Value_DoubleValue struct {
	DoubleValue float64 `protobuf:"fixed64,2,opt,name=double_value,json=doubleValue,oneof"`
}
type Value_StringValue struct {
	StringValue string `protobuf:"bytes,3,opt,name=string_value,json=stringValue,oneof"`
}
type Value_BoolValue struct {
	BoolValue bool `protobuf:"varint,4,opt,name=bool_value,json=boolValue,oneof"`
}
type Value_ByteValue struct {
	ByteValue []byte `protobuf:"bytes,5,opt,name=byte_value,json=byteValue,proto3,oneof"`
}
type Value_AnyValue struct {
	AnyValue *google_protobuf.Any `protobuf:"bytes,6,opt,name=any_value,json=anyValue,oneof"`
}
type Value_JsonValue struct {
	JsonValue []byte `protobuf:"bytes,7,opt,name=json_value,json=jsonValue,proto3,oneof"`
}

func (*Value_IntValue) isValue_Value()    {}
func (*Value_DoubleValue) isValue_Value() {}
func (*Value_StringValue) isValue_Value() {}
func (*Value_BoolValue) isValue_Value()   {}
func (*Value_ByteValue) isValue_Value()   {}
func (*Value_AnyValue) isValue_Value()    {}
func (*Value_JsonValue) isValue_Value()   {}

func (m *Value) GetValue() isValue_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *Value) GetIntValue() int64 {
	if x, ok := m.GetValue().(*Value_IntValue); ok {
		return x.IntValue
	}
	return 0
}

func (m *Value) GetDoubleValue() float64 {
	if x, ok := m.GetValue().(*Value_DoubleValue); ok {
		return x.DoubleValue
	}
	return 0
}

func (m *Value) GetStringValue() string {
	if x, ok := m.GetValue().(*Value_StringValue); ok {
		return x.StringValue
	}
	return ""
}

func (m *Value) GetBoolValue() bool {
	if x, ok := m.GetValue().(*Value_BoolValue); ok {
		return x.BoolValue
	}
	return false
}

func (m *Value) GetByteValue() []byte {
	if x, ok := m.GetValue().(*Value_ByteValue); ok {
		return x.ByteValue
	}
	return nil
}

func (m *Value) GetAnyValue() *google_protobuf.Any {
	if x, ok := m.GetValue().(*Value_AnyValue); ok {
		return x.AnyValue
	}
	return nil
}

func (m *Value) GetJsonValue() []byte {
	if x, ok := m.GetValue().(*Value_JsonValue); ok {
		return x.JsonValue
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Value) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Value_OneofMarshaler, _Value_OneofUnmarshaler, _Value_OneofSizer, []interface{}{
		(*Value_IntValue)(nil),
		(*Value_DoubleValue)(nil),
		(*Value_StringValue)(nil),
		(*Value_BoolValue)(nil),
		(*Value_ByteValue)(nil),
		(*Value_AnyValue)(nil),
		(*Value_JsonValue)(nil),
	}
}

func _Value_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Value)
	// value
	switch x := m.Value.(type) {
	case *Value_IntValue:
		b.EncodeVarint(1<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.IntValue))
	case *Value_DoubleValue:
		b.EncodeVarint(2<<3 | proto.WireFixed64)
		b.EncodeFixed64(math.Float64bits(x.DoubleValue))
	case *Value_StringValue:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.StringValue)
	case *Value_BoolValue:
		t := uint64(0)
		if x.BoolValue {
			t = 1
		}
		b.EncodeVarint(4<<3 | proto.WireVarint)
		b.EncodeVarint(t)
	case *Value_ByteValue:
		b.EncodeVarint(5<<3 | proto.WireBytes)
		b.EncodeRawBytes(x.ByteValue)
	case *Value_AnyValue:
		b.EncodeVarint(6<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.AnyValue); err != nil {
			return err
		}
	case *Value_JsonValue:
		b.EncodeVarint(7<<3 | proto.WireBytes)
		b.EncodeRawBytes(x.JsonValue)
	case nil:
	default:
		return fmt.Errorf("Value.Value has unexpected type %T", x)
	}
	return nil
}

func _Value_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Value)
	switch tag {
	case 1: // value.int_value
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Value = &Value_IntValue{int64(x)}
		return true, err
	case 2: // value.double_value
		if wire != proto.WireFixed64 {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeFixed64()
		m.Value = &Value_DoubleValue{math.Float64frombits(x)}
		return true, err
	case 3: // value.string_value
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Value = &Value_StringValue{x}
		return true, err
	case 4: // value.bool_value
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Value = &Value_BoolValue{x != 0}
		return true, err
	case 5: // value.byte_value
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeRawBytes(true)
		m.Value = &Value_ByteValue{x}
		return true, err
	case 6: // value.any_value
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(google_protobuf.Any)
		err := b.DecodeMessage(msg)
		m.Value = &Value_AnyValue{msg}
		return true, err
	case 7: // value.json_value
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeRawBytes(true)
		m.Value = &Value_JsonValue{x}
		return true, err
	default:
		return false, nil
	}
}

func _Value_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Value)
	// value
	switch x := m.Value.(type) {
	case *Value_IntValue:
		n += proto.SizeVarint(1<<3 | proto.WireVarint)
		n += proto.SizeVarint(uint64(x.IntValue))
	case *Value_DoubleValue:
		n += proto.SizeVarint(2<<3 | proto.WireFixed64)
		n += 8
	case *Value_StringValue:
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.StringValue)))
		n += len(x.StringValue)
	case *Value_BoolValue:
		n += proto.SizeVarint(4<<3 | proto.WireVarint)
		n += 1
	case *Value_ByteValue:
		n += proto.SizeVarint(5<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.ByteValue)))
		n += len(x.ByteValue)
	case *Value_AnyValue:
		s := proto.Size(x.AnyValue)
		n += proto.SizeVarint(6<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Value_JsonValue:
		n += proto.SizeVarint(7<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.JsonValue)))
		n += len(x.JsonValue)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// Path contains a list of nodes that make the unique path
// to an value in the tree.
type Path struct {
	Node []string `protobuf:"bytes,1,rep,name=node" json:"node,omitempty"`
}

func (m *Path) Reset()                    { *m = Path{} }
func (m *Path) String() string            { return proto.CompactTextString(m) }
func (*Path) ProtoMessage()               {}
func (*Path) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type Report struct {
	Id         string   `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	TestSuites []*Suite `protobuf:"bytes,2,rep,name=test_suites,json=testSuites" json:"test_suites,omitempty"`
	Timestamp  string   `protobuf:"bytes,3,opt,name=timestamp" json:"timestamp,omitempty"`
}

func (m *Report) Reset()                    { *m = Report{} }
func (m *Report) String() string            { return proto.CompactTextString(m) }
func (*Report) ProtoMessage()               {}
func (*Report) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Report) GetTestSuites() []*Suite {
	if m != nil {
		return m.TestSuites
	}
	return nil
}

type Suite struct {
	Id        string               `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	TestCases []*Case              `protobuf:"bytes,2,rep,name=test_cases,json=testCases" json:"test_cases,omitempty"`
	Testbed   *openresult1.Testbed `protobuf:"bytes,3,opt,name=testbed" json:"testbed,omitempty"`
	Meta      map[string]*Value    `protobuf:"bytes,4,rep,name=meta" json:"meta,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *Suite) Reset()                    { *m = Suite{} }
func (m *Suite) String() string            { return proto.CompactTextString(m) }
func (*Suite) ProtoMessage()               {}
func (*Suite) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Suite) GetTestCases() []*Case {
	if m != nil {
		return m.TestCases
	}
	return nil
}

func (m *Suite) GetTestbed() *openresult1.Testbed {
	if m != nil {
		return m.Testbed
	}
	return nil
}

func (m *Suite) GetMeta() map[string]*Value {
	if m != nil {
		return m.Meta
	}
	return nil
}

type Case struct {
	Id      string            `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Text    string            `protobuf:"bytes,2,opt,name=text" json:"text,omitempty"`
	Results []*Result         `protobuf:"bytes,3,rep,name=results" json:"results,omitempty"`
	Meta    map[string]*Value `protobuf:"bytes,4,rep,name=meta" json:"meta,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *Case) Reset()                    { *m = Case{} }
func (m *Case) String() string            { return proto.CompactTextString(m) }
func (*Case) ProtoMessage()               {}
func (*Case) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Case) GetResults() []*Result {
	if m != nil {
		return m.Results
	}
	return nil
}

func (m *Case) GetMeta() map[string]*Value {
	if m != nil {
		return m.Meta
	}
	return nil
}

type Result struct {
	// Path identifier of the result. The path only makes
	// the result unique within it's own tree.
	Path *Path `protobuf:"bytes,1,opt,name=path" json:"path,omitempty"`
	// Globally unique identifier of the result. Every
	// result with a globally unique result should be
	// indexable.
	Uuid string `protobuf:"bytes,2,opt,name=uuid" json:"uuid,omitempty"`
	// Timestamp in nanos since unix epoch.
	Timestamp uint64 `protobuf:"varint,3,opt,name=timestamp" json:"timestamp,omitempty"`
	// Duration of result in nano seconds.
	Duration uint64 `protobuf:"varint,4,opt,name=duration" json:"duration,omitempty"`
	// text describes the result
	Text string `protobuf:"bytes,5,opt,name=text" json:"text,omitempty"`
	// All operations are evaluated in the following order got op want.
	// If op is not set or is UNKNOWN no comparison is done.
	// If want is not set then no comparison is done.
	Got  *Value    `protobuf:"bytes,6,opt,name=got" json:"got,omitempty"`
	Want *Value    `protobuf:"bytes,7,opt,name=want" json:"want,omitempty"`
	Op   Operation `protobuf:"varint,8,opt,name=op,enum=openresult.Operation" json:"op,omitempty"`
	// Children results of the parent.
	Results []*Result `protobuf:"bytes,9,rep,name=results" json:"results,omitempty"`
	// Arbitrary properties of the result.
	Meta map[string]*google_protobuf.Any `protobuf:"bytes,10,rep,name=meta" json:"meta,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *Result) Reset()                    { *m = Result{} }
func (m *Result) String() string            { return proto.CompactTextString(m) }
func (*Result) ProtoMessage()               {}
func (*Result) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *Result) GetPath() *Path {
	if m != nil {
		return m.Path
	}
	return nil
}

func (m *Result) GetGot() *Value {
	if m != nil {
		return m.Got
	}
	return nil
}

func (m *Result) GetWant() *Value {
	if m != nil {
		return m.Want
	}
	return nil
}

func (m *Result) GetResults() []*Result {
	if m != nil {
		return m.Results
	}
	return nil
}

func (m *Result) GetMeta() map[string]*google_protobuf.Any {
	if m != nil {
		return m.Meta
	}
	return nil
}

func init() {
	proto.RegisterType((*Value)(nil), "openresult.Value")
	proto.RegisterType((*Path)(nil), "openresult.Path")
	proto.RegisterType((*Report)(nil), "openresult.Report")
	proto.RegisterType((*Suite)(nil), "openresult.Suite")
	proto.RegisterType((*Case)(nil), "openresult.Case")
	proto.RegisterType((*Result)(nil), "openresult.Result")
	proto.RegisterEnum("openresult.Operation", Operation_name, Operation_value)
}

func init() {
	proto.RegisterFile("github.com/marcushines/openresult/proto/result.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 708 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xb4, 0x94, 0xdd, 0x6e, 0x12, 0x41,
	0x14, 0xc7, 0xbb, 0x5f, 0xc0, 0x1e, 0x6a, 0xc5, 0xa9, 0x9a, 0x8a, 0x35, 0x1a, 0xaa, 0xb1, 0x69,
	0x74, 0x31, 0x54, 0x13, 0xe3, 0x1d, 0x1a, 0x62, 0xd5, 0x96, 0xea, 0x94, 0xea, 0x25, 0x59, 0xca,
	0x4a, 0xb7, 0xc2, 0xce, 0x66, 0x77, 0x56, 0xe5, 0x31, 0x4c, 0x7c, 0x12, 0x9f, 0xc5, 0x47, 0xf1,
	0x01, 0x3c, 0xf3, 0xc1, 0xb2, 0x2d, 0x44, 0xbd, 0xf1, 0x8a, 0x33, 0xe7, 0xff, 0xdb, 0x33, 0x73,
	0xfe, 0x67, 0x06, 0x78, 0x3c, 0x0a, 0xf9, 0x69, 0x36, 0xf0, 0x4e, 0xd8, 0xa4, 0x39, 0xf1, 0x93,
	0x93, 0x2c, 0x3d, 0x0d, 0xa3, 0x20, 0x6d, 0xb2, 0x38, 0x88, 0x92, 0x20, 0xcd, 0xc6, 0xbc, 0x19,
	0x27, 0x8c, 0xb3, 0xa6, 0x5a, 0x78, 0x72, 0x41, 0x60, 0x2e, 0xd7, 0x6f, 0x8c, 0x18, 0x1b, 0x8d,
	0x03, 0x85, 0x0d, 0xb2, 0x8f, 0x4d, 0x3f, 0x9a, 0x2a, 0xac, 0xfe, 0xe4, 0x5f, 0x8b, 0xf3, 0x20,
	0xe5, 0x83, 0x60, 0xa8, 0x3e, 0x6b, 0x7c, 0x33, 0xc1, 0x79, 0xef, 0x8f, 0xb3, 0x80, 0xdc, 0x02,
	0x37, 0x8c, 0x78, 0xff, 0xb3, 0x58, 0x6c, 0x18, 0x77, 0x8c, 0x6d, 0x6b, 0x6f, 0x85, 0x56, 0x30,
	0xa5, 0xe4, 0x2d, 0x58, 0x1d, 0xb2, 0x6c, 0x30, 0x0e, 0x34, 0x61, 0x22, 0x61, 0x20, 0x51, 0x55,
	0xd9, 0x1c, 0x4a, 0x79, 0x12, 0x46, 0x23, 0x0d, 0x59, 0x08, 0xb9, 0x02, 0x52, 0x59, 0x05, 0xdd,
	0x06, 0x18, 0x30, 0x36, 0xd6, 0x88, 0x8d, 0x48, 0x05, 0x11, 0x57, 0xe4, 0xe6, 0xc0, 0x94, 0xcf,
	0x36, 0x72, 0x10, 0x58, 0x95, 0x00, 0xe6, 0x14, 0xb0, 0x0b, 0x2e, 0x36, 0xae, 0xf5, 0x12, 0xea,
	0xd5, 0xd6, 0x55, 0x4f, 0x59, 0xe3, 0xcd, 0xac, 0xf1, 0xda, 0xd1, 0x54, 0x34, 0x80, 0x60, 0x5e,
	0xf5, 0x2c, 0x65, 0x91, 0xfe, 0xaa, 0x3c, 0xab, 0x2a, 0x72, 0x12, 0x78, 0x5e, 0x06, 0x47, 0x6a,
	0x8d, 0x3a, 0xd8, 0x6f, 0x7d, 0x7e, 0x4a, 0x08, 0xd8, 0x11, 0x1b, 0x0a, 0x33, 0xac, 0x6d, 0x97,
	0xca, 0xb8, 0x71, 0x06, 0x25, 0x1a, 0xc4, 0x2c, 0xe1, 0x64, 0x0d, 0xcc, 0x70, 0x28, 0x8d, 0x72,
	0x29, 0x46, 0xa4, 0x05, 0x55, 0x61, 0x6d, 0x3f, 0xcd, 0x42, 0xfc, 0x45, 0x7f, 0x2c, 0x3c, 0xd6,
	0x15, 0x6f, 0xee, 0xbf, 0x77, 0x24, 0x14, 0x0a, 0x82, 0x92, 0x61, 0x4a, 0x36, 0xc1, 0xe5, 0xe1,
	0x04, 0xd7, 0xfe, 0x24, 0x56, 0x66, 0xd1, 0x79, 0xa2, 0xf1, 0xcb, 0x00, 0x47, 0x82, 0x0b, 0x7b,
	0x35, 0x41, 0x56, 0xe9, 0x9f, 0xf8, 0x69, 0xbe, 0x55, 0xad, 0xb8, 0xd5, 0x0b, 0x14, 0xb0, 0x14,
	0x32, 0x22, 0x4a, 0xc9, 0x43, 0x28, 0xeb, 0xb9, 0xcb, 0x6d, 0xaa, 0xad, 0xf5, 0x22, 0xdd, 0x53,
	0x12, 0x9d, 0x31, 0x58, 0xdf, 0x9e, 0x04, 0xdc, 0xc7, 0xe1, 0x88, 0xca, 0x37, 0x17, 0x9a, 0xf0,
	0x0e, 0x50, 0xed, 0x44, 0x3c, 0x99, 0x52, 0x09, 0xd6, 0x5f, 0x83, 0x9b, 0xa7, 0x48, 0x0d, 0xac,
	0x4f, 0xc1, 0x54, 0x1f, 0x57, 0x84, 0xe4, 0xbe, 0xb6, 0x56, 0xde, 0x9a, 0x0b, 0xae, 0x48, 0xf3,
	0xa9, 0xd2, 0x9f, 0x99, 0x4f, 0x8d, 0xc6, 0x4f, 0x03, 0x6c, 0x71, 0xea, 0x85, 0xae, 0x71, 0x1e,
	0x3c, 0xf8, 0xca, 0x65, 0x11, 0x9c, 0x87, 0x88, 0xc9, 0x03, 0x28, 0xab, 0x3a, 0x29, 0x36, 0x26,
	0x0e, 0x4b, 0x8a, 0xb5, 0xa9, 0xfc, 0xa1, 0x33, 0x84, 0x78, 0xe7, 0xfa, 0xaa, 0x5f, 0x74, 0xec,
	0xbf, 0xb6, 0xf5, 0xc3, 0x12, 0x57, 0x47, 0x68, 0xe4, 0x2e, 0xd8, 0x31, 0x5e, 0x30, 0x59, 0xea,
	0xc2, 0xe0, 0xc4, 0xc5, 0xa3, 0x52, 0x15, 0xed, 0x66, 0x19, 0x1a, 0xa0, 0xdb, 0x15, 0xf1, 0xe2,
	0x85, 0xb1, 0x0b, 0x17, 0x86, 0xd4, 0xa1, 0x32, 0xcc, 0x12, 0x9f, 0x87, 0x2c, 0x92, 0xef, 0xca,
	0xa6, 0xf9, 0x3a, 0x37, 0xcf, 0x29, 0x98, 0xb7, 0x05, 0xd6, 0x88, 0x71, 0xfd, 0x82, 0x96, 0x9c,
	0x5e, 0xa8, 0xe4, 0x1e, 0xd8, 0x5f, 0xfc, 0x88, 0xcb, 0x17, 0xb3, 0x94, 0x92, 0x32, 0x62, 0x26,
	0x8b, 0x37, 0x2a, 0x08, 0xad, 0xb5, 0xae, 0x15, 0xa1, 0xc3, 0x38, 0x50, 0x47, 0xa0, 0x08, 0x14,
	0xe7, 0xe5, 0xfe, 0x7d, 0x5e, 0x8f, 0xf4, 0xbc, 0x40, 0xa2, 0x9b, 0x8b, 0xe8, 0xc2, 0xc4, 0x0e,
	0xfe, 0x3c, 0xb1, 0x9d, 0xf3, 0x13, 0x5b, 0xfa, 0xaf, 0x51, 0x18, 0xda, 0xce, 0x77, 0x03, 0xdc,
	0xbc, 0x01, 0x52, 0x85, 0xf2, 0x71, 0xf7, 0x4d, 0xf7, 0xf0, 0x43, 0xb7, 0xb6, 0x42, 0x5c, 0x70,
	0x3a, 0xef, 0x8e, 0xdb, 0xfb, 0x35, 0x83, 0x5c, 0x02, 0xb7, 0x7b, 0xd8, 0xeb, 0xab, 0xa5, 0x89,
	0xdb, 0xae, 0xbe, 0xa4, 0x9d, 0x76, 0xaf, 0x43, 0xfb, 0xbd, 0xbd, 0x76, 0xb7, 0x66, 0x09, 0x60,
	0xbf, 0x73, 0x74, 0xa4, 0x96, 0x36, 0xb9, 0x0e, 0xa4, 0x08, 0xe8, 0x0f, 0x1d, 0xb2, 0x0e, 0x97,
	0x73, 0x4c, 0x27, 0x4b, 0xa4, 0x04, 0xe6, 0xab, 0x6e, 0xad, 0x4c, 0x00, 0x4a, 0x62, 0x13, 0x8c,
	0x2b, 0x83, 0x92, 0x3c, 0xef, 0xee, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x52, 0x9e, 0x0c, 0x88,
	0x52, 0x06, 0x00, 0x00,
}
