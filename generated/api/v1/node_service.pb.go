// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: api/v1/node_service.proto

package v1

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	storage "github.com/stackrox/rox/generated/storage"
	io "io"
	math "math"
	math_bits "math/bits"
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

type ListNodesRequest struct {
	ClusterId            string   `protobuf:"bytes,1,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListNodesRequest) Reset()         { *m = ListNodesRequest{} }
func (m *ListNodesRequest) String() string { return proto.CompactTextString(m) }
func (*ListNodesRequest) ProtoMessage()    {}
func (*ListNodesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_bc3747bf289f5ecf, []int{0}
}
func (m *ListNodesRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ListNodesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ListNodesRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ListNodesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListNodesRequest.Merge(m, src)
}
func (m *ListNodesRequest) XXX_Size() int {
	return m.Size()
}
func (m *ListNodesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListNodesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListNodesRequest proto.InternalMessageInfo

func (m *ListNodesRequest) GetClusterId() string {
	if m != nil {
		return m.ClusterId
	}
	return ""
}

func (m *ListNodesRequest) MessageClone() proto.Message {
	return m.Clone()
}
func (m *ListNodesRequest) Clone() *ListNodesRequest {
	if m == nil {
		return nil
	}
	cloned := new(ListNodesRequest)
	*cloned = *m

	return cloned
}

type ListNodesResponse struct {
	Nodes                []*storage.Node `protobuf:"bytes,1,rep,name=nodes,proto3" json:"nodes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *ListNodesResponse) Reset()         { *m = ListNodesResponse{} }
func (m *ListNodesResponse) String() string { return proto.CompactTextString(m) }
func (*ListNodesResponse) ProtoMessage()    {}
func (*ListNodesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_bc3747bf289f5ecf, []int{1}
}
func (m *ListNodesResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ListNodesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ListNodesResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ListNodesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListNodesResponse.Merge(m, src)
}
func (m *ListNodesResponse) XXX_Size() int {
	return m.Size()
}
func (m *ListNodesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListNodesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListNodesResponse proto.InternalMessageInfo

func (m *ListNodesResponse) GetNodes() []*storage.Node {
	if m != nil {
		return m.Nodes
	}
	return nil
}

func (m *ListNodesResponse) MessageClone() proto.Message {
	return m.Clone()
}
func (m *ListNodesResponse) Clone() *ListNodesResponse {
	if m == nil {
		return nil
	}
	cloned := new(ListNodesResponse)
	*cloned = *m

	if m.Nodes != nil {
		cloned.Nodes = make([]*storage.Node, len(m.Nodes))
		for idx, v := range m.Nodes {
			cloned.Nodes[idx] = v.Clone()
		}
	}
	return cloned
}

type GetNodeRequest struct {
	ClusterId            string   `protobuf:"bytes,1,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
	NodeId               string   `protobuf:"bytes,2,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetNodeRequest) Reset()         { *m = GetNodeRequest{} }
func (m *GetNodeRequest) String() string { return proto.CompactTextString(m) }
func (*GetNodeRequest) ProtoMessage()    {}
func (*GetNodeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_bc3747bf289f5ecf, []int{2}
}
func (m *GetNodeRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetNodeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetNodeRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetNodeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetNodeRequest.Merge(m, src)
}
func (m *GetNodeRequest) XXX_Size() int {
	return m.Size()
}
func (m *GetNodeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetNodeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetNodeRequest proto.InternalMessageInfo

func (m *GetNodeRequest) GetClusterId() string {
	if m != nil {
		return m.ClusterId
	}
	return ""
}

func (m *GetNodeRequest) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

func (m *GetNodeRequest) MessageClone() proto.Message {
	return m.Clone()
}
func (m *GetNodeRequest) Clone() *GetNodeRequest {
	if m == nil {
		return nil
	}
	cloned := new(GetNodeRequest)
	*cloned = *m

	return cloned
}

type ExportNodeRequest struct {
	Timeout              int32    `protobuf:"varint,1,opt,name=timeout,proto3" json:"timeout,omitempty"`
	Query                string   `protobuf:"bytes,2,opt,name=query,proto3" json:"query,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExportNodeRequest) Reset()         { *m = ExportNodeRequest{} }
func (m *ExportNodeRequest) String() string { return proto.CompactTextString(m) }
func (*ExportNodeRequest) ProtoMessage()    {}
func (*ExportNodeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_bc3747bf289f5ecf, []int{3}
}
func (m *ExportNodeRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ExportNodeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ExportNodeRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ExportNodeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExportNodeRequest.Merge(m, src)
}
func (m *ExportNodeRequest) XXX_Size() int {
	return m.Size()
}
func (m *ExportNodeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ExportNodeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ExportNodeRequest proto.InternalMessageInfo

func (m *ExportNodeRequest) GetTimeout() int32 {
	if m != nil {
		return m.Timeout
	}
	return 0
}

func (m *ExportNodeRequest) GetQuery() string {
	if m != nil {
		return m.Query
	}
	return ""
}

func (m *ExportNodeRequest) MessageClone() proto.Message {
	return m.Clone()
}
func (m *ExportNodeRequest) Clone() *ExportNodeRequest {
	if m == nil {
		return nil
	}
	cloned := new(ExportNodeRequest)
	*cloned = *m

	return cloned
}

type ExportNodeResponse struct {
	Node                 *storage.Node `protobuf:"bytes,1,opt,name=node,proto3" json:"node,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ExportNodeResponse) Reset()         { *m = ExportNodeResponse{} }
func (m *ExportNodeResponse) String() string { return proto.CompactTextString(m) }
func (*ExportNodeResponse) ProtoMessage()    {}
func (*ExportNodeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_bc3747bf289f5ecf, []int{4}
}
func (m *ExportNodeResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ExportNodeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ExportNodeResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ExportNodeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExportNodeResponse.Merge(m, src)
}
func (m *ExportNodeResponse) XXX_Size() int {
	return m.Size()
}
func (m *ExportNodeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ExportNodeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ExportNodeResponse proto.InternalMessageInfo

func (m *ExportNodeResponse) GetNode() *storage.Node {
	if m != nil {
		return m.Node
	}
	return nil
}

func (m *ExportNodeResponse) MessageClone() proto.Message {
	return m.Clone()
}
func (m *ExportNodeResponse) Clone() *ExportNodeResponse {
	if m == nil {
		return nil
	}
	cloned := new(ExportNodeResponse)
	*cloned = *m

	cloned.Node = m.Node.Clone()
	return cloned
}

func init() {
	proto.RegisterType((*ListNodesRequest)(nil), "v1.ListNodesRequest")
	proto.RegisterType((*ListNodesResponse)(nil), "v1.ListNodesResponse")
	proto.RegisterType((*GetNodeRequest)(nil), "v1.GetNodeRequest")
	proto.RegisterType((*ExportNodeRequest)(nil), "v1.ExportNodeRequest")
	proto.RegisterType((*ExportNodeResponse)(nil), "v1.ExportNodeResponse")
}

func init() { proto.RegisterFile("api/v1/node_service.proto", fileDescriptor_bc3747bf289f5ecf) }

var fileDescriptor_bc3747bf289f5ecf = []byte{
	// 413 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xcd, 0xaa, 0xd3, 0x40,
	0x14, 0xc7, 0x6f, 0xa2, 0xbd, 0xa5, 0xa7, 0x28, 0xed, 0xa1, 0xd6, 0x18, 0x34, 0xd4, 0x71, 0xd3,
	0xd5, 0xc4, 0xd4, 0x85, 0xae, 0x15, 0xd1, 0x82, 0xb8, 0x88, 0x20, 0xc5, 0x4d, 0x89, 0xcd, 0x50,
	0x06, 0x6b, 0x26, 0xcd, 0x4c, 0x42, 0xa5, 0x74, 0xe3, 0x2b, 0xb8, 0xd1, 0x37, 0x72, 0x29, 0xf8,
	0x02, 0x52, 0x7d, 0x10, 0x99, 0x99, 0xd4, 0x7e, 0xb9, 0xb8, 0xcb, 0xf3, 0xf5, 0xfb, 0x9f, 0xf3,
	0x9f, 0x81, 0x3b, 0x49, 0xce, 0xc3, 0x2a, 0x0a, 0x33, 0x91, 0xb2, 0xa9, 0x64, 0x45, 0xc5, 0x67,
	0x8c, 0xe6, 0x85, 0x50, 0x02, 0xdd, 0x2a, 0xf2, 0xef, 0xce, 0x85, 0x98, 0x2f, 0x58, 0xa8, 0xbb,
	0x92, 0x2c, 0x13, 0x2a, 0x51, 0x5c, 0x64, 0xd2, 0x76, 0xf8, 0x28, 0x95, 0x28, 0x92, 0x39, 0x33,
	0xd3, 0x36, 0x47, 0x22, 0xe8, 0xbc, 0xe2, 0x52, 0xbd, 0x16, 0x29, 0x93, 0x31, 0x5b, 0x96, 0x4c,
	0x2a, 0xbc, 0x07, 0x30, 0x5b, 0x94, 0x52, 0xb1, 0x62, 0xca, 0x53, 0xcf, 0x19, 0x38, 0xc3, 0x56,
	0xdc, 0xaa, 0x33, 0xe3, 0x94, 0x3c, 0x81, 0xee, 0xc1, 0x88, 0xcc, 0x45, 0x26, 0x19, 0x3e, 0x80,
	0x86, 0xa6, 0x4a, 0xcf, 0x19, 0x5c, 0x1b, 0xb6, 0x47, 0x37, 0x68, 0xad, 0x45, 0x75, 0x5b, 0x6c,
	0x6b, 0xe4, 0x25, 0xdc, 0x7c, 0xc1, 0xcc, 0xe0, 0xd5, 0xa4, 0xf0, 0x36, 0x34, 0xcd, 0xa5, 0x3c,
	0xf5, 0x5c, 0x53, 0xbb, 0xd4, 0xe1, 0x38, 0x25, 0xcf, 0xa0, 0xfb, 0x7c, 0x95, 0x8b, 0xe2, 0x08,
	0xe6, 0x41, 0x53, 0xf1, 0x8f, 0x4c, 0x94, 0xca, 0x90, 0x1a, 0xf1, 0x2e, 0xc4, 0x1e, 0x34, 0x96,
	0x25, 0x2b, 0x3e, 0xd5, 0x14, 0x1b, 0x90, 0xc7, 0x80, 0x87, 0x90, 0xfa, 0x92, 0xfb, 0x70, 0x5d,
	0x8b, 0x18, 0xc4, 0xd9, 0x21, 0xa6, 0x34, 0xfa, 0xe6, 0x42, 0x5b, 0x87, 0x6f, 0xec, 0x03, 0xe0,
	0x04, 0x5a, 0xff, 0x1c, 0xc1, 0x1e, 0xad, 0x22, 0x7a, 0xea, 0xa9, 0x7f, 0xeb, 0x24, 0x6b, 0xc5,
	0x48, 0xf0, 0xf9, 0xe7, 0x9f, 0x2f, 0xae, 0x87, 0xfd, 0xdd, 0xa3, 0xca, 0x70, 0xbd, 0x37, 0x64,
	0x83, 0x6f, 0xa1, 0x59, 0x3b, 0x86, 0xa8, 0x09, 0xc7, 0xf6, 0xf9, 0xc7, 0xdb, 0x91, 0xa1, 0xa1,
	0x11, 0x1c, 0xfc, 0x9f, 0x16, 0xae, 0x6b, 0x33, 0x37, 0x38, 0x81, 0xf6, 0xfe, 0x74, 0x89, 0x66,
	0xbb, 0x33, 0x43, 0xfd, 0xfe, 0x69, 0xba, 0xde, 0xda, 0x33, 0x3a, 0x88, 0x1d, 0xad, 0xc3, 0x4c,
	0xdd, 0xca, 0x3d, 0x74, 0x9e, 0xd2, 0xef, 0xdb, 0xc0, 0xf9, 0xb1, 0x0d, 0x9c, 0x5f, 0xdb, 0xc0,
	0xf9, 0xfa, 0x3b, 0xb8, 0x00, 0x8f, 0x0b, 0x2a, 0x55, 0x32, 0xfb, 0x50, 0x88, 0x95, 0xfd, 0x75,
	0x34, 0xc9, 0x39, 0xad, 0xa2, 0x77, 0x6e, 0x15, 0x4d, 0x2e, 0xde, 0x5f, 0x9a, 0xdc, 0xa3, 0xbf,
	0x01, 0x00, 0x00, 0xff, 0xff, 0x02, 0x83, 0xd1, 0xf2, 0xdc, 0x02, 0x00, 0x00,
}

func (m *ListNodesRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ListNodesRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ListNodesRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.ClusterId) > 0 {
		i -= len(m.ClusterId)
		copy(dAtA[i:], m.ClusterId)
		i = encodeVarintNodeService(dAtA, i, uint64(len(m.ClusterId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ListNodesResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ListNodesResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ListNodesResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Nodes) > 0 {
		for iNdEx := len(m.Nodes) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Nodes[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintNodeService(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *GetNodeRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetNodeRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetNodeRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.NodeId) > 0 {
		i -= len(m.NodeId)
		copy(dAtA[i:], m.NodeId)
		i = encodeVarintNodeService(dAtA, i, uint64(len(m.NodeId)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.ClusterId) > 0 {
		i -= len(m.ClusterId)
		copy(dAtA[i:], m.ClusterId)
		i = encodeVarintNodeService(dAtA, i, uint64(len(m.ClusterId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ExportNodeRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ExportNodeRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ExportNodeRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Query) > 0 {
		i -= len(m.Query)
		copy(dAtA[i:], m.Query)
		i = encodeVarintNodeService(dAtA, i, uint64(len(m.Query)))
		i--
		dAtA[i] = 0x12
	}
	if m.Timeout != 0 {
		i = encodeVarintNodeService(dAtA, i, uint64(m.Timeout))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *ExportNodeResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ExportNodeResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ExportNodeResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Node != nil {
		{
			size, err := m.Node.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintNodeService(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintNodeService(dAtA []byte, offset int, v uint64) int {
	offset -= sovNodeService(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ListNodesRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ClusterId)
	if l > 0 {
		n += 1 + l + sovNodeService(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ListNodesResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Nodes) > 0 {
		for _, e := range m.Nodes {
			l = e.Size()
			n += 1 + l + sovNodeService(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *GetNodeRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ClusterId)
	if l > 0 {
		n += 1 + l + sovNodeService(uint64(l))
	}
	l = len(m.NodeId)
	if l > 0 {
		n += 1 + l + sovNodeService(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ExportNodeRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Timeout != 0 {
		n += 1 + sovNodeService(uint64(m.Timeout))
	}
	l = len(m.Query)
	if l > 0 {
		n += 1 + l + sovNodeService(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ExportNodeResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Node != nil {
		l = m.Node.Size()
		n += 1 + l + sovNodeService(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovNodeService(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozNodeService(x uint64) (n int) {
	return sovNodeService(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ListNodesRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNodeService
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ListNodesRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ListNodesRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClusterId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNodeService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthNodeService
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNodeService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClusterId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipNodeService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthNodeService
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ListNodesResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNodeService
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ListNodesResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ListNodesResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Nodes", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNodeService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthNodeService
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthNodeService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Nodes = append(m.Nodes, &storage.Node{})
			if err := m.Nodes[len(m.Nodes)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipNodeService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthNodeService
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *GetNodeRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNodeService
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GetNodeRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetNodeRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClusterId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNodeService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthNodeService
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNodeService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClusterId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NodeId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNodeService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthNodeService
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNodeService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NodeId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipNodeService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthNodeService
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ExportNodeRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNodeService
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ExportNodeRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ExportNodeRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timeout", wireType)
			}
			m.Timeout = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNodeService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Timeout |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Query", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNodeService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthNodeService
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNodeService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Query = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipNodeService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthNodeService
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ExportNodeResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNodeService
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ExportNodeResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ExportNodeResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Node", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNodeService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthNodeService
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthNodeService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Node == nil {
				m.Node = &storage.Node{}
			}
			if err := m.Node.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipNodeService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthNodeService
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipNodeService(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowNodeService
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowNodeService
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowNodeService
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthNodeService
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupNodeService
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthNodeService
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthNodeService        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowNodeService          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupNodeService = fmt.Errorf("proto: unexpected end of group")
)
