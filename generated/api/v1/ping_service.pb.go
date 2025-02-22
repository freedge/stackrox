// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: api/v1/ping_service.proto

package v1

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

type PongMessage struct {
	Status               string   `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PongMessage) Reset()         { *m = PongMessage{} }
func (m *PongMessage) String() string { return proto.CompactTextString(m) }
func (*PongMessage) ProtoMessage()    {}
func (*PongMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_e11869b18a72c24a, []int{0}
}
func (m *PongMessage) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PongMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PongMessage.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PongMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PongMessage.Merge(m, src)
}
func (m *PongMessage) XXX_Size() int {
	return m.Size()
}
func (m *PongMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_PongMessage.DiscardUnknown(m)
}

var xxx_messageInfo_PongMessage proto.InternalMessageInfo

func (m *PongMessage) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *PongMessage) MessageClone() proto.Message {
	return m.Clone()
}
func (m *PongMessage) Clone() *PongMessage {
	if m == nil {
		return nil
	}
	cloned := new(PongMessage)
	*cloned = *m

	return cloned
}

func init() {
	proto.RegisterType((*PongMessage)(nil), "v1.PongMessage")
}

func init() { proto.RegisterFile("api/v1/ping_service.proto", fileDescriptor_e11869b18a72c24a) }

var fileDescriptor_e11869b18a72c24a = []byte{
	// 220 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4c, 0x2c, 0xc8, 0xd4,
	0x2f, 0x33, 0xd4, 0x2f, 0xc8, 0xcc, 0x4b, 0x8f, 0x2f, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x33, 0x94, 0x12, 0x82, 0x4a, 0xa7, 0xe6, 0x16,
	0x94, 0x54, 0x42, 0xc4, 0xa5, 0x64, 0xd2, 0xf3, 0xf3, 0xd3, 0x73, 0x52, 0xf5, 0x41, 0x52, 0x89,
	0x79, 0x79, 0xf9, 0x25, 0x89, 0x25, 0x99, 0xf9, 0x79, 0xc5, 0x10, 0x59, 0x25, 0x55, 0x2e, 0xee,
	0x80, 0xfc, 0xbc, 0x74, 0xdf, 0xd4, 0xe2, 0xe2, 0xc4, 0xf4, 0x54, 0x21, 0x31, 0x2e, 0xb6, 0xe2,
	0x92, 0xc4, 0x92, 0xd2, 0x62, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x28, 0xcf, 0xc8, 0x99,
	0x8b, 0x3b, 0x20, 0x33, 0x2f, 0x3d, 0x18, 0x62, 0xa3, 0x90, 0x09, 0x17, 0x0b, 0x88, 0x2b, 0xc4,
	0xa9, 0x57, 0x66, 0xa8, 0xe7, 0x0a, 0xb2, 0x4c, 0x8a, 0x1f, 0xc4, 0x44, 0x32, 0x4a, 0x49, 0xa0,
	0xe9, 0xf2, 0x93, 0xc9, 0x4c, 0x5c, 0x42, 0x1c, 0x30, 0xf7, 0x3a, 0xe9, 0x9d, 0x78, 0x24, 0xc7,
	0x78, 0xe1, 0x91, 0x1c, 0xe3, 0x83, 0x47, 0x72, 0x8c, 0x33, 0x1e, 0xcb, 0x31, 0x70, 0x49, 0x64,
	0xe6, 0xeb, 0x15, 0x97, 0x24, 0x26, 0x67, 0x17, 0xe5, 0x57, 0x40, 0x1c, 0xa4, 0x97, 0x58, 0x90,
	0xa9, 0x57, 0x66, 0x18, 0xc5, 0x54, 0x66, 0x18, 0xc1, 0x98, 0xc4, 0x06, 0x16, 0x33, 0x06, 0x04,
	0x00, 0x00, 0xff, 0xff, 0xe9, 0xa7, 0x6a, 0x1a, 0xf7, 0x00, 0x00, 0x00,
}

func (m *PongMessage) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PongMessage) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PongMessage) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Status) > 0 {
		i -= len(m.Status)
		copy(dAtA[i:], m.Status)
		i = encodeVarintPingService(dAtA, i, uint64(len(m.Status)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintPingService(dAtA []byte, offset int, v uint64) int {
	offset -= sovPingService(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *PongMessage) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Status)
	if l > 0 {
		n += 1 + l + sovPingService(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovPingService(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPingService(x uint64) (n int) {
	return sovPingService(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *PongMessage) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPingService
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
			return fmt.Errorf("proto: PongMessage: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PongMessage: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPingService
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
				return ErrInvalidLengthPingService
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPingService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Status = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPingService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPingService
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
func skipPingService(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPingService
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
					return 0, ErrIntOverflowPingService
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
					return 0, ErrIntOverflowPingService
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
				return 0, ErrInvalidLengthPingService
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPingService
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPingService
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPingService        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPingService          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPingService = fmt.Errorf("proto: unexpected end of group")
)
