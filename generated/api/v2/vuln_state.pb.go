// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: api/v2/vuln_state.proto

package v2

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

// VulnerabilityState are the possible applicable to CVE. By default all vulnerabilities are in observed state.
type VulnerabilityState int32

const (
	// This is the default state and indicates that the CVE is not excluded from policy evaluation and risk evaluation.
	VulnerabilityState_OBSERVED VulnerabilityState = 0
	// Indicates that the vulnerability is deferred. A deferred CVE is excluded from policy evaluation and risk evaluation.
	VulnerabilityState_DEFERRED VulnerabilityState = 1
	// Indicates that the vulnerability is a false-positive. A false-positive CVE is excluded from policy evaluation and risk evaluation.
	VulnerabilityState_FALSE_POSITIVE VulnerabilityState = 2
)

var VulnerabilityState_name = map[int32]string{
	0: "OBSERVED",
	1: "DEFERRED",
	2: "FALSE_POSITIVE",
}

var VulnerabilityState_value = map[string]int32{
	"OBSERVED":       0,
	"DEFERRED":       1,
	"FALSE_POSITIVE": 2,
}

func (x VulnerabilityState) String() string {
	return proto.EnumName(VulnerabilityState_name, int32(x))
}

func (VulnerabilityState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_3d97ed81f434f268, []int{0}
}

func init() {
	proto.RegisterEnum("v2.VulnerabilityState", VulnerabilityState_name, VulnerabilityState_value)
}

func init() { proto.RegisterFile("api/v2/vuln_state.proto", fileDescriptor_3d97ed81f434f268) }

var fileDescriptor_3d97ed81f434f268 = []byte{
	// 169 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4f, 0x2c, 0xc8, 0xd4,
	0x2f, 0x33, 0xd2, 0x2f, 0x2b, 0xcd, 0xc9, 0x8b, 0x2f, 0x2e, 0x49, 0x2c, 0x49, 0xd5, 0x2b, 0x28,
	0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x33, 0xd2, 0x72, 0xe1, 0x12, 0x0a, 0x2b, 0xcd, 0xc9, 0x4b,
	0x2d, 0x4a, 0x4c, 0xca, 0xcc, 0xc9, 0x2c, 0xa9, 0x0c, 0x06, 0xc9, 0x0b, 0xf1, 0x70, 0x71, 0xf8,
	0x3b, 0x05, 0xbb, 0x06, 0x85, 0xb9, 0xba, 0x08, 0x30, 0x80, 0x78, 0x2e, 0xae, 0x6e, 0xae, 0x41,
	0x41, 0xae, 0x2e, 0x02, 0x8c, 0x42, 0x42, 0x5c, 0x7c, 0x6e, 0x8e, 0x3e, 0xc1, 0xae, 0xf1, 0x01,
	0xfe, 0xc1, 0x9e, 0x21, 0x9e, 0x61, 0xae, 0x02, 0x4c, 0x4e, 0x7a, 0x27, 0x1e, 0xc9, 0x31, 0x5e,
	0x78, 0x24, 0xc7, 0xf8, 0xe0, 0x91, 0x1c, 0xe3, 0x8c, 0xc7, 0x72, 0x0c, 0x5c, 0x12, 0x99, 0xf9,
	0x7a, 0xc5, 0x25, 0x89, 0xc9, 0xd9, 0x45, 0xf9, 0x15, 0x10, 0xeb, 0xf4, 0x12, 0x0b, 0x32, 0xf5,
	0xca, 0x8c, 0xa2, 0x98, 0xca, 0x8c, 0x92, 0xd8, 0xc0, 0x22, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x6e, 0x15, 0x82, 0xfb, 0x9b, 0x00, 0x00, 0x00,
}
