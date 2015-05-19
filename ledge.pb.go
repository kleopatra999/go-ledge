// Code generated by protoc-gen-go.
// source: ledge.proto
// DO NOT EDIT!

/*
Package ledge is a generated protocol buffer package.

It is generated from these files:
	ledge.proto

It has these top-level messages:
	ProtoEntry
*/
package ledge

import proto "github.com/golang/protobuf/proto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal

type ProtoLevel int32

const (
	ProtoLevel_NONE  ProtoLevel = 0
	ProtoLevel_DEBUG ProtoLevel = 1
	ProtoLevel_INFO  ProtoLevel = 2
	ProtoLevel_WARN  ProtoLevel = 3
	ProtoLevel_ERROR ProtoLevel = 4
	ProtoLevel_FATAL ProtoLevel = 5
	ProtoLevel_PANIC ProtoLevel = 6
)

var ProtoLevel_name = map[int32]string{
	0: "NONE",
	1: "DEBUG",
	2: "INFO",
	3: "WARN",
	4: "ERROR",
	5: "FATAL",
	6: "PANIC",
}
var ProtoLevel_value = map[string]int32{
	"NONE":  0,
	"DEBUG": 1,
	"INFO":  2,
	"WARN":  3,
	"ERROR": 4,
	"FATAL": 5,
	"PANIC": 6,
}

func (x ProtoLevel) String() string {
	return proto.EnumName(ProtoLevel_name, int32(x))
}

type ProtoEntry struct {
	Id                   string            `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	TimeUnixNsec         uint64            `protobuf:"varint,2,opt,name=time_unix_nsec" json:"time_unix_nsec,omitempty"`
	ContextTypeToContext map[string][]byte `protobuf:"bytes,3,rep,name=context_type_to_context" json:"context_type_to_context,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value,proto3"`
	EventType            string            `protobuf:"bytes,4,opt,name=event_type" json:"event_type,omitempty"`
	Event                []byte            `protobuf:"bytes,5,opt,name=event,proto3" json:"event,omitempty"`
	WriterOutput         []byte            `protobuf:"bytes,6,opt,name=writer_output,proto3" json:"writer_output,omitempty"`
}

func (m *ProtoEntry) Reset()         { *m = ProtoEntry{} }
func (m *ProtoEntry) String() string { return proto.CompactTextString(m) }
func (*ProtoEntry) ProtoMessage()    {}

func (m *ProtoEntry) GetContextTypeToContext() map[string][]byte {
	if m != nil {
		return m.ContextTypeToContext
	}
	return nil
}

func init() {
	proto.RegisterEnum("ledge.ProtoLevel", ProtoLevel_name, ProtoLevel_value)
}
