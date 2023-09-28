// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ojo/oracle/v1/events.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	proto "github.com/cosmos/gogoproto/proto"
	_ "github.com/gogo/protobuf/gogoproto"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// EventDelegateFeedConsent is emitted on Msg/DelegateFeedConsent
type EventDelegateFeedConsent struct {
	// Operator bech32 address who delegates his feed consent
	Operator string `protobuf:"bytes,1,opt,name=operator,proto3" json:"operator,omitempty"`
	// Delegate bech32 address
	Delegate string `protobuf:"bytes,2,opt,name=delegate,proto3" json:"delegate,omitempty"`
}

func (m *EventDelegateFeedConsent) Reset()         { *m = EventDelegateFeedConsent{} }
func (m *EventDelegateFeedConsent) String() string { return proto.CompactTextString(m) }
func (*EventDelegateFeedConsent) ProtoMessage()    {}
func (*EventDelegateFeedConsent) Descriptor() ([]byte, []int) {
	return fileDescriptor_6b8914220a86f2fd, []int{0}
}
func (m *EventDelegateFeedConsent) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventDelegateFeedConsent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventDelegateFeedConsent.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventDelegateFeedConsent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventDelegateFeedConsent.Merge(m, src)
}
func (m *EventDelegateFeedConsent) XXX_Size() int {
	return m.Size()
}
func (m *EventDelegateFeedConsent) XXX_DiscardUnknown() {
	xxx_messageInfo_EventDelegateFeedConsent.DiscardUnknown(m)
}

var xxx_messageInfo_EventDelegateFeedConsent proto.InternalMessageInfo

// EventSetFxRate is emitted on exchange rate update
type EventSetFxRate struct {
	// uToken denom
	Denom string `protobuf:"bytes,1,opt,name=denom,proto3" json:"denom,omitempty"`
	// Exchange rate (based to USD)
	Rate github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,2,opt,name=rate,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"rate"`
	// block num on which exchange rate was updated
	BlockNum uint64 `protobuf:"varint,3,opt,name=block_num,json=blockNum,proto3" json:"block_num,omitempty"`
}

func (m *EventSetFxRate) Reset()         { *m = EventSetFxRate{} }
func (m *EventSetFxRate) String() string { return proto.CompactTextString(m) }
func (*EventSetFxRate) ProtoMessage()    {}
func (*EventSetFxRate) Descriptor() ([]byte, []int) {
	return fileDescriptor_6b8914220a86f2fd, []int{1}
}
func (m *EventSetFxRate) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventSetFxRate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventSetFxRate.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventSetFxRate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventSetFxRate.Merge(m, src)
}
func (m *EventSetFxRate) XXX_Size() int {
	return m.Size()
}
func (m *EventSetFxRate) XXX_DiscardUnknown() {
	xxx_messageInfo_EventSetFxRate.DiscardUnknown(m)
}

var xxx_messageInfo_EventSetFxRate proto.InternalMessageInfo

func init() {
	proto.RegisterType((*EventDelegateFeedConsent)(nil), "ojo.oracle.v1.EventDelegateFeedConsent")
	proto.RegisterType((*EventSetFxRate)(nil), "ojo.oracle.v1.EventSetFxRate")
}

func init() { proto.RegisterFile("ojo/oracle/v1/events.proto", fileDescriptor_6b8914220a86f2fd) }

var fileDescriptor_6b8914220a86f2fd = []byte{
	// 340 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0x3f, 0x4f, 0xfa, 0x40,
	0x18, 0xc7, 0x7b, 0xbf, 0x1f, 0x1a, 0xb8, 0x44, 0x87, 0x86, 0xa1, 0x62, 0x72, 0x10, 0x06, 0x83,
	0x43, 0xdb, 0x10, 0x1d, 0x5d, 0x44, 0xc4, 0xcd, 0x98, 0xb2, 0xb9, 0x90, 0xd2, 0x3e, 0xa9, 0xfc,
	0xe9, 0x3d, 0xe4, 0xee, 0x40, 0x7c, 0x03, 0xce, 0x2e, 0xbe, 0x13, 0x5e, 0x04, 0x23, 0x61, 0x32,
	0x0e, 0x44, 0xe9, 0x1b, 0x31, 0xed, 0x15, 0x75, 0x73, 0xba, 0xe7, 0xcf, 0xf7, 0xf9, 0x7c, 0x9f,
	0xbb, 0xa3, 0x15, 0x1c, 0xa2, 0x8b, 0xc2, 0x0f, 0xc6, 0xe0, 0xce, 0x9a, 0x2e, 0xcc, 0x80, 0x2b,
	0xe9, 0x4c, 0x04, 0x2a, 0x34, 0x0f, 0x70, 0x88, 0x8e, 0xee, 0x39, 0xb3, 0x66, 0xe5, 0x28, 0x40,
	0x19, 0xa3, 0xec, 0x65, 0x4d, 0x57, 0x27, 0x5a, 0x59, 0x29, 0x47, 0x18, 0xa1, 0xae, 0xa7, 0x91,
	0xae, 0xd6, 0x9f, 0x09, 0xb5, 0xae, 0x53, 0x60, 0x1b, 0xc6, 0x10, 0xf9, 0x0a, 0x3a, 0x00, 0xe1,
	0x15, 0x72, 0x09, 0x5c, 0x99, 0xe7, 0xb4, 0x88, 0x13, 0x10, 0xbe, 0x42, 0x61, 0x91, 0x1a, 0x69,
	0x94, 0x5a, 0xd6, 0x7a, 0x61, 0x97, 0x73, 0xec, 0x65, 0x18, 0x0a, 0x90, 0xb2, 0xab, 0xc4, 0x80,
	0x47, 0xde, 0xb7, 0x32, 0x9d, 0x0a, 0x73, 0x98, 0xf5, 0xef, 0xaf, 0xa9, 0x9d, 0xb2, 0xfe, 0x4a,
	0xe8, 0x61, 0xb6, 0x48, 0x17, 0x54, 0x67, 0xee, 0xf9, 0x0a, 0xcc, 0x32, 0xdd, 0x0b, 0x81, 0x63,
	0xac, 0xbd, 0x3d, 0x9d, 0x98, 0x77, 0xb4, 0x20, 0x7e, 0xd0, 0x17, 0xcb, 0x4d, 0xd5, 0x78, 0xdf,
	0x54, 0x4f, 0xa2, 0x81, 0x7a, 0x98, 0xf6, 0x9d, 0x00, 0xe3, 0xfc, 0xda, 0xf9, 0x61, 0xcb, 0x70,
	0xe4, 0xaa, 0xa7, 0x09, 0x48, 0xa7, 0x0d, 0xc1, 0x7a, 0x61, 0xd3, 0x7c, 0x91, 0x36, 0x04, 0x5e,
	0x46, 0x32, 0x8f, 0x69, 0xa9, 0x3f, 0xc6, 0x60, 0xd4, 0xe3, 0xd3, 0xd8, 0xfa, 0x5f, 0x23, 0x8d,
	0x82, 0x57, 0xcc, 0x0a, 0xb7, 0xd3, 0xb8, 0x75, 0xb3, 0xfc, 0x64, 0xc6, 0x72, 0xcb, 0xc8, 0x6a,
	0xcb, 0xc8, 0xc7, 0x96, 0x91, 0x97, 0x84, 0x19, 0xab, 0x84, 0x19, 0x6f, 0x09, 0x33, 0xee, 0x4f,
	0x7f, 0xd9, 0xe2, 0x10, 0x6d, 0x0e, 0xea, 0x11, 0xc5, 0x28, 0x8d, 0xdd, 0xf9, 0xee, 0xcf, 0x32,
	0xf7, 0xfe, 0x7e, 0xf6, 0xe0, 0x67, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x1c, 0x70, 0xa6, 0x63,
	0xce, 0x01, 0x00, 0x00,
}

func (m *EventDelegateFeedConsent) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventDelegateFeedConsent) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventDelegateFeedConsent) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Delegate) > 0 {
		i -= len(m.Delegate)
		copy(dAtA[i:], m.Delegate)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Delegate)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Operator) > 0 {
		i -= len(m.Operator)
		copy(dAtA[i:], m.Operator)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Operator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EventSetFxRate) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventSetFxRate) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventSetFxRate) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.BlockNum != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.BlockNum))
		i--
		dAtA[i] = 0x18
	}
	{
		size := m.Rate.Size()
		i -= size
		if _, err := m.Rate.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintEvents(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintEvents(dAtA []byte, offset int, v uint64) int {
	offset -= sovEvents(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *EventDelegateFeedConsent) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Operator)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.Delegate)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	return n
}

func (m *EventSetFxRate) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = m.Rate.Size()
	n += 1 + l + sovEvents(uint64(l))
	if m.BlockNum != 0 {
		n += 1 + sovEvents(uint64(m.BlockNum))
	}
	return n
}

func sovEvents(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEvents(x uint64) (n int) {
	return sovEvents(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *EventDelegateFeedConsent) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: EventDelegateFeedConsent: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventDelegateFeedConsent: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Operator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Operator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Delegate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Delegate = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *EventSetFxRate) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: EventSetFxRate: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventSetFxRate: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Rate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Rate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockNum", wireType)
			}
			m.BlockNum = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BlockNum |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipEvents(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEvents
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
					return 0, ErrIntOverflowEvents
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
					return 0, ErrIntOverflowEvents
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
				return 0, ErrInvalidLengthEvents
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEvents
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEvents
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEvents        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEvents          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEvents = fmt.Errorf("proto: unexpected end of group")
)
