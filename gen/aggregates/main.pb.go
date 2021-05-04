// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: aggregates/main.proto

package aggregates_v1

import (
	fmt "fmt"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Aggreate codes
type AggregateType int32

const (
	AggregateType_AGGREGATE_TYPE_INVALID AggregateType = 0
	AggregateType_USER                   AggregateType = 2
	AggregateType_ORDER                  AggregateType = 3
)

var AggregateType_name = map[int32]string{
	0: "AGGREGATE_TYPE_INVALID",
	2: "USER",
	3: "ORDER",
}

var AggregateType_value = map[string]int32{
	"AGGREGATE_TYPE_INVALID": 0,
	"USER":                   2,
	"ORDER":                  3,
}

func (x AggregateType) String() string {
	return proto.EnumName(AggregateType_name, int32(x))
}

func (AggregateType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a69aea3c9c7c4073, []int{0}
}

func init() {
	proto.RegisterEnum("aggregates_v1.AggregateType", AggregateType_name, AggregateType_value)
}

func init() { proto.RegisterFile("aggregates/main.proto", fileDescriptor_a69aea3c9c7c4073) }

var fileDescriptor_a69aea3c9c7c4073 = []byte{
	// 119 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4d, 0x4c, 0x4f, 0x2f,
	0x4a, 0x4d, 0x4f, 0x2c, 0x49, 0x2d, 0xd6, 0xcf, 0x4d, 0xcc, 0xcc, 0xd3, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x17, 0xe2, 0x45, 0x08, 0xc7, 0x97, 0x19, 0x6a, 0x39, 0x70, 0xf1, 0x3a, 0xc2, 0x04, 0x42,
	0x2a, 0x0b, 0x52, 0x85, 0xa4, 0xb8, 0xc4, 0x1c, 0xdd, 0xdd, 0x83, 0x5c, 0xdd, 0x1d, 0x43, 0x5c,
	0xe3, 0x43, 0x22, 0x03, 0x5c, 0xe3, 0x3d, 0xfd, 0xc2, 0x1c, 0x7d, 0x3c, 0x5d, 0x04, 0x18, 0x84,
	0x38, 0xb8, 0x58, 0x42, 0x83, 0x5d, 0x83, 0x04, 0x98, 0x84, 0x38, 0xb9, 0x58, 0xfd, 0x83, 0x5c,
	0x5c, 0x83, 0x04, 0x98, 0x93, 0xd8, 0xc0, 0xe6, 0x1a, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x48,
	0xb8, 0x2b, 0x54, 0x70, 0x00, 0x00, 0x00,
}
