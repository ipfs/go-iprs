// Code generated by protoc-gen-go.
// source: record.proto
// DO NOT EDIT!

/*
Package ipfs_iprs is a generated protocol buffer package.

It is generated from these files:
	record.proto

It has these top-level messages:
	Record
*/
package ipfs_iprs

import proto "code.google.com/p/goprotobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type Record struct {
	// optional bytes Type = 1;     // link to the record type object (already in links)
	Version          *uint64 `protobuf:"varint,2,opt" json:"Version,omitempty"`
	Validity         []byte  `protobuf:"bytes,3,opt" json:"Validity,omitempty"`
	Value            []byte  `protobuf:"bytes,4,opt" json:"Value,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Record) Reset()         { *m = Record{} }
func (m *Record) String() string { return proto.CompactTextString(m) }
func (*Record) ProtoMessage()    {}

func (m *Record) GetVersion() uint64 {
	if m != nil && m.Version != nil {
		return *m.Version
	}
	return 0
}

func (m *Record) GetValidity() []byte {
	if m != nil {
		return m.Validity
	}
	return nil
}

func (m *Record) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func init() {
}
