// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: storage/compliance_integration.proto

package storage

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

// Next Tag: 7
type ComplianceIntegration struct {
	Id                  string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" sql:"pk,type(uuid)"`                                // @gotags: sql:"pk,type(uuid)"
	Version             string `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty" search:"Compliance Operator Version,hidden,store"`                      // @gotags: search:"Compliance Operator Version,hidden,store"
	ClusterId           string `protobuf:"bytes,3,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty" search:"Cluster ID,hidden,store" sql:"fk(Cluster:id),no-fk-constraint,type(uuid),index=category:unique;name:compliance_unique_indicator"` // @gotags: search:"Cluster ID,hidden,store" sql:"fk(Cluster:id),no-fk-constraint,type(uuid),index=category:unique;name:compliance_unique_indicator"
	ComplianceNamespace string `protobuf:"bytes,4,opt,name=compliance_namespace,json=complianceNamespace,proto3" json:"compliance_namespace,omitempty"`
	// Collection of errors that occurred while trying to obtain compliance health info.
	StatusErrors         []string `protobuf:"bytes,5,rep,name=status_errors,json=statusErrors,proto3" json:"status_errors,omitempty"`
	OperatorInstalled    bool     `protobuf:"varint,6,opt,name=operator_installed,json=operatorInstalled,proto3" json:"operator_installed,omitempty" search:"Compliance Operator Installed,hidden"` // @gotags: search:"Compliance Operator Installed,hidden"
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ComplianceIntegration) Reset()         { *m = ComplianceIntegration{} }
func (m *ComplianceIntegration) String() string { return proto.CompactTextString(m) }
func (*ComplianceIntegration) ProtoMessage()    {}
func (*ComplianceIntegration) Descriptor() ([]byte, []int) {
	return fileDescriptor_14e26a23cbdbee2c, []int{0}
}
func (m *ComplianceIntegration) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ComplianceIntegration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ComplianceIntegration.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ComplianceIntegration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ComplianceIntegration.Merge(m, src)
}
func (m *ComplianceIntegration) XXX_Size() int {
	return m.Size()
}
func (m *ComplianceIntegration) XXX_DiscardUnknown() {
	xxx_messageInfo_ComplianceIntegration.DiscardUnknown(m)
}

var xxx_messageInfo_ComplianceIntegration proto.InternalMessageInfo

func (m *ComplianceIntegration) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ComplianceIntegration) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *ComplianceIntegration) GetClusterId() string {
	if m != nil {
		return m.ClusterId
	}
	return ""
}

func (m *ComplianceIntegration) GetComplianceNamespace() string {
	if m != nil {
		return m.ComplianceNamespace
	}
	return ""
}

func (m *ComplianceIntegration) GetStatusErrors() []string {
	if m != nil {
		return m.StatusErrors
	}
	return nil
}

func (m *ComplianceIntegration) GetOperatorInstalled() bool {
	if m != nil {
		return m.OperatorInstalled
	}
	return false
}

func (m *ComplianceIntegration) MessageClone() proto.Message {
	return m.Clone()
}
func (m *ComplianceIntegration) Clone() *ComplianceIntegration {
	if m == nil {
		return nil
	}
	cloned := new(ComplianceIntegration)
	*cloned = *m

	if m.StatusErrors != nil {
		cloned.StatusErrors = make([]string, len(m.StatusErrors))
		copy(cloned.StatusErrors, m.StatusErrors)
	}
	return cloned
}

func init() {
	proto.RegisterType((*ComplianceIntegration)(nil), "storage.ComplianceIntegration")
}

func init() {
	proto.RegisterFile("storage/compliance_integration.proto", fileDescriptor_14e26a23cbdbee2c)
}

var fileDescriptor_14e26a23cbdbee2c = []byte{
	// 259 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0xb1, 0x4a, 0xc4, 0x40,
	0x10, 0x86, 0xdd, 0x9c, 0xde, 0x99, 0x45, 0x05, 0x57, 0x85, 0xb5, 0x30, 0x04, 0xb5, 0x48, 0x63,
	0x44, 0xf4, 0x09, 0x14, 0x8b, 0x34, 0x16, 0x29, 0x6d, 0xc2, 0x9a, 0x0c, 0xc7, 0x62, 0x6e, 0x27,
	0xcc, 0xce, 0x89, 0x8f, 0xe2, 0x23, 0x59, 0xfa, 0x08, 0x12, 0x0b, 0x5f, 0x43, 0xc8, 0x65, 0xbd,
	0x2b, 0xe7, 0xfb, 0xbe, 0x62, 0xf8, 0xe5, 0xa5, 0x67, 0x24, 0x33, 0x87, 0xeb, 0x1a, 0x17, 0x5d,
	0x6b, 0x8d, 0xab, 0xa1, 0xb2, 0x8e, 0x61, 0x4e, 0x86, 0x2d, 0xba, 0xbc, 0x23, 0x64, 0x54, 0xb3,
	0xb1, 0x3a, 0xff, 0x15, 0xf2, 0xe4, 0xe1, 0xbf, 0x2c, 0xd6, 0xa1, 0x3a, 0x90, 0x91, 0x6d, 0xb4,
	0x48, 0x45, 0x16, 0x97, 0x91, 0x6d, 0x94, 0x96, 0xb3, 0x37, 0x20, 0x6f, 0xd1, 0xe9, 0x68, 0x80,
	0xe1, 0x54, 0x67, 0x52, 0xd6, 0xed, 0xd2, 0x33, 0x50, 0x65, 0x1b, 0x3d, 0x19, 0x64, 0x3c, 0x92,
	0xa2, 0x51, 0x37, 0xf2, 0x78, 0xe3, 0x17, 0x67, 0x16, 0xe0, 0x3b, 0x53, 0x83, 0xde, 0x1e, 0xc2,
	0xa3, 0xb5, 0x7b, 0x0a, 0x4a, 0x5d, 0xc8, 0x7d, 0xcf, 0x86, 0x97, 0xbe, 0x02, 0x22, 0x24, 0xaf,
	0x77, 0xd2, 0x49, 0x16, 0x97, 0x7b, 0x2b, 0xf8, 0x38, 0x30, 0x75, 0x25, 0x15, 0x76, 0x40, 0x86,
	0x91, 0x2a, 0xeb, 0x3c, 0x9b, 0xb6, 0x85, 0x46, 0x4f, 0x53, 0x91, 0xed, 0x96, 0x87, 0xc1, 0x14,
	0x41, 0xdc, 0xdf, 0x7d, 0xf6, 0x89, 0xf8, 0xea, 0x13, 0xf1, 0xdd, 0x27, 0xe2, 0xe3, 0x27, 0xd9,
	0x92, 0xa7, 0x16, 0x73, 0xcf, 0xa6, 0x7e, 0x25, 0x7c, 0x5f, 0xed, 0x92, 0x8f, 0xb3, 0x3c, 0x87,
	0x7d, 0x5e, 0xa6, 0x03, 0xbf, 0xfd, 0x0b, 0x00, 0x00, 0xff, 0xff, 0x0e, 0x04, 0x7b, 0x31, 0x57,
	0x01, 0x00, 0x00,
}

func (m *ComplianceIntegration) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ComplianceIntegration) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ComplianceIntegration) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.OperatorInstalled {
		i--
		if m.OperatorInstalled {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x30
	}
	if len(m.StatusErrors) > 0 {
		for iNdEx := len(m.StatusErrors) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.StatusErrors[iNdEx])
			copy(dAtA[i:], m.StatusErrors[iNdEx])
			i = encodeVarintComplianceIntegration(dAtA, i, uint64(len(m.StatusErrors[iNdEx])))
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.ComplianceNamespace) > 0 {
		i -= len(m.ComplianceNamespace)
		copy(dAtA[i:], m.ComplianceNamespace)
		i = encodeVarintComplianceIntegration(dAtA, i, uint64(len(m.ComplianceNamespace)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.ClusterId) > 0 {
		i -= len(m.ClusterId)
		copy(dAtA[i:], m.ClusterId)
		i = encodeVarintComplianceIntegration(dAtA, i, uint64(len(m.ClusterId)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Version) > 0 {
		i -= len(m.Version)
		copy(dAtA[i:], m.Version)
		i = encodeVarintComplianceIntegration(dAtA, i, uint64(len(m.Version)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintComplianceIntegration(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintComplianceIntegration(dAtA []byte, offset int, v uint64) int {
	offset -= sovComplianceIntegration(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ComplianceIntegration) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovComplianceIntegration(uint64(l))
	}
	l = len(m.Version)
	if l > 0 {
		n += 1 + l + sovComplianceIntegration(uint64(l))
	}
	l = len(m.ClusterId)
	if l > 0 {
		n += 1 + l + sovComplianceIntegration(uint64(l))
	}
	l = len(m.ComplianceNamespace)
	if l > 0 {
		n += 1 + l + sovComplianceIntegration(uint64(l))
	}
	if len(m.StatusErrors) > 0 {
		for _, s := range m.StatusErrors {
			l = len(s)
			n += 1 + l + sovComplianceIntegration(uint64(l))
		}
	}
	if m.OperatorInstalled {
		n += 2
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovComplianceIntegration(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozComplianceIntegration(x uint64) (n int) {
	return sovComplianceIntegration(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ComplianceIntegration) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowComplianceIntegration
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
			return fmt.Errorf("proto: ComplianceIntegration: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ComplianceIntegration: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowComplianceIntegration
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
				return ErrInvalidLengthComplianceIntegration
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthComplianceIntegration
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowComplianceIntegration
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
				return ErrInvalidLengthComplianceIntegration
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthComplianceIntegration
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Version = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClusterId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowComplianceIntegration
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
				return ErrInvalidLengthComplianceIntegration
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthComplianceIntegration
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClusterId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ComplianceNamespace", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowComplianceIntegration
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
				return ErrInvalidLengthComplianceIntegration
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthComplianceIntegration
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ComplianceNamespace = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StatusErrors", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowComplianceIntegration
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
				return ErrInvalidLengthComplianceIntegration
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthComplianceIntegration
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StatusErrors = append(m.StatusErrors, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field OperatorInstalled", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowComplianceIntegration
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.OperatorInstalled = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipComplianceIntegration(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthComplianceIntegration
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
func skipComplianceIntegration(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowComplianceIntegration
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
					return 0, ErrIntOverflowComplianceIntegration
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
					return 0, ErrIntOverflowComplianceIntegration
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
				return 0, ErrInvalidLengthComplianceIntegration
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupComplianceIntegration
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthComplianceIntegration
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthComplianceIntegration        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowComplianceIntegration          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupComplianceIntegration = fmt.Errorf("proto: unexpected end of group")
)
