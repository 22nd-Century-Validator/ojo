// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ojo/airdrop/v1/query.proto

package types

import (
	context "context"
	fmt "fmt"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// QueryAirdropAccountRequest is the request type for the QueryAirdropAccount RPC method.
type QueryAirdropAccountRequest struct {
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
}

func (m *QueryAirdropAccountRequest) Reset()         { *m = QueryAirdropAccountRequest{} }
func (m *QueryAirdropAccountRequest) String() string { return proto.CompactTextString(m) }
func (*QueryAirdropAccountRequest) ProtoMessage()    {}
func (*QueryAirdropAccountRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e174292c6e21a38, []int{0}
}
func (m *QueryAirdropAccountRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAirdropAccountRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAirdropAccountRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAirdropAccountRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAirdropAccountRequest.Merge(m, src)
}
func (m *QueryAirdropAccountRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryAirdropAccountRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAirdropAccountRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAirdropAccountRequest proto.InternalMessageInfo

func (m *QueryAirdropAccountRequest) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

// QueryAirdropAccountResponse is the response type for the QueryAirdropAccount RPC method.
type QueryAirdropAccountResponse struct {
	Eligible     bool   `protobuf:"varint,1,opt,name=eligible,proto3" json:"eligible,omitempty"`
	Claimed      bool   `protobuf:"varint,2,opt,name=claimed,proto3" json:"claimed,omitempty"`
	ClaimAddress string `protobuf:"bytes,3,opt,name=claim_address,json=claimAddress,proto3" json:"claim_address,omitempty"`
}

func (m *QueryAirdropAccountResponse) Reset()         { *m = QueryAirdropAccountResponse{} }
func (m *QueryAirdropAccountResponse) String() string { return proto.CompactTextString(m) }
func (*QueryAirdropAccountResponse) ProtoMessage()    {}
func (*QueryAirdropAccountResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e174292c6e21a38, []int{1}
}
func (m *QueryAirdropAccountResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAirdropAccountResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAirdropAccountResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAirdropAccountResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAirdropAccountResponse.Merge(m, src)
}
func (m *QueryAirdropAccountResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryAirdropAccountResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAirdropAccountResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAirdropAccountResponse proto.InternalMessageInfo

func (m *QueryAirdropAccountResponse) GetEligible() bool {
	if m != nil {
		return m.Eligible
	}
	return false
}

func (m *QueryAirdropAccountResponse) GetClaimed() bool {
	if m != nil {
		return m.Claimed
	}
	return false
}

func (m *QueryAirdropAccountResponse) GetClaimAddress() string {
	if m != nil {
		return m.ClaimAddress
	}
	return ""
}

func init() {
	proto.RegisterType((*QueryAirdropAccountRequest)(nil), "ojo.airdrop.v1.QueryAirdropAccountRequest")
	proto.RegisterType((*QueryAirdropAccountResponse)(nil), "ojo.airdrop.v1.QueryAirdropAccountResponse")
}

func init() { proto.RegisterFile("ojo/airdrop/v1/query.proto", fileDescriptor_8e174292c6e21a38) }

var fileDescriptor_8e174292c6e21a38 = []byte{
	// 309 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xca, 0xcf, 0xca, 0xd7,
	0x4f, 0xcc, 0x2c, 0x4a, 0x29, 0xca, 0x2f, 0xd0, 0x2f, 0x33, 0xd4, 0x2f, 0x2c, 0x4d, 0x2d, 0xaa,
	0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0xcb, 0xcf, 0xca, 0xd7, 0x83, 0xca, 0xe9, 0x95,
	0x19, 0x4a, 0xc9, 0xa4, 0xe7, 0xe7, 0xa7, 0xe7, 0xa4, 0xea, 0x27, 0x16, 0x64, 0xea, 0x27, 0xe6,
	0xe5, 0xe5, 0x97, 0x24, 0x96, 0x64, 0xe6, 0xe7, 0x15, 0x43, 0x54, 0x2b, 0x99, 0x71, 0x49, 0x05,
	0x82, 0x34, 0x3b, 0x42, 0x34, 0x38, 0x26, 0x27, 0xe7, 0x97, 0xe6, 0x95, 0x04, 0xa5, 0x16, 0x96,
	0xa6, 0x16, 0x97, 0x08, 0x49, 0x70, 0xb1, 0x27, 0xa6, 0xa4, 0x14, 0xa5, 0x16, 0x17, 0x4b, 0x30,
	0x2a, 0x30, 0x6a, 0x70, 0x06, 0xc1, 0xb8, 0x4a, 0x15, 0x5c, 0xd2, 0x58, 0xf5, 0x15, 0x17, 0xe4,
	0xe7, 0x15, 0xa7, 0x0a, 0x49, 0x71, 0x71, 0xa4, 0xe6, 0x64, 0xa6, 0x67, 0x26, 0xe5, 0xa4, 0x82,
	0x75, 0x72, 0x04, 0xc1, 0xf9, 0x20, 0x43, 0x93, 0x73, 0x12, 0x33, 0x73, 0x53, 0x53, 0x24, 0x98,
	0xc0, 0x52, 0x30, 0xae, 0x90, 0x32, 0x17, 0x2f, 0x98, 0x19, 0x0f, 0xb3, 0x94, 0x19, 0x6c, 0x29,
	0x0f, 0x58, 0xd0, 0x11, 0x22, 0x66, 0x34, 0x83, 0x91, 0x8b, 0x15, 0x6c, 0xb5, 0x50, 0x1f, 0x23,
	0x97, 0x30, 0x16, 0x47, 0x08, 0x69, 0xe9, 0xa1, 0x06, 0x81, 0x1e, 0x6e, 0x1f, 0x4a, 0x69, 0x13,
	0xa5, 0x16, 0xe2, 0x2b, 0x25, 0xb9, 0xa6, 0xcb, 0x4f, 0x26, 0x33, 0x49, 0x08, 0x89, 0xe9, 0x23,
	0x87, 0x7f, 0x35, 0xd4, 0xb5, 0xb5, 0x4e, 0x2e, 0x27, 0x1e, 0xc9, 0x31, 0x5e, 0x78, 0x24, 0xc7,
	0xf8, 0xe0, 0x91, 0x1c, 0xe3, 0x84, 0xc7, 0x72, 0x0c, 0x17, 0x1e, 0xcb, 0x31, 0xdc, 0x78, 0x2c,
	0xc7, 0x10, 0xa5, 0x95, 0x9e, 0x59, 0x92, 0x51, 0x9a, 0xa4, 0x97, 0x9c, 0x9f, 0x0b, 0xd2, 0xab,
	0x9b, 0x97, 0x5a, 0x52, 0x9e, 0x5f, 0x94, 0x0d, 0x36, 0xa7, 0x02, 0x6e, 0x52, 0x49, 0x65, 0x41,
	0x6a, 0x71, 0x12, 0x1b, 0x38, 0x66, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xd5, 0xdc, 0x86,
	0x33, 0xe5, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// QueryAirdropAccount returns an existing airdrop account, along with whether or not
	// the user is eligible to claim, and whether or not the airdrop has been claimed.
	// If the airdrop has been claimed, the account to which the tokens were sent should be
	//returned as well.
	QueryAirdropAccount(ctx context.Context, in *QueryAirdropAccountRequest, opts ...grpc.CallOption) (*QueryAirdropAccountResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) QueryAirdropAccount(ctx context.Context, in *QueryAirdropAccountRequest, opts ...grpc.CallOption) (*QueryAirdropAccountResponse, error) {
	out := new(QueryAirdropAccountResponse)
	err := c.cc.Invoke(ctx, "/ojo.airdrop.v1.Query/QueryAirdropAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// QueryAirdropAccount returns an existing airdrop account, along with whether or not
	// the user is eligible to claim, and whether or not the airdrop has been claimed.
	// If the airdrop has been claimed, the account to which the tokens were sent should be
	//returned as well.
	QueryAirdropAccount(context.Context, *QueryAirdropAccountRequest) (*QueryAirdropAccountResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) QueryAirdropAccount(ctx context.Context, req *QueryAirdropAccountRequest) (*QueryAirdropAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryAirdropAccount not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_QueryAirdropAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAirdropAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).QueryAirdropAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ojo.airdrop.v1.Query/QueryAirdropAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).QueryAirdropAccount(ctx, req.(*QueryAirdropAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ojo.airdrop.v1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "QueryAirdropAccount",
			Handler:    _Query_QueryAirdropAccount_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ojo/airdrop/v1/query.proto",
}

func (m *QueryAirdropAccountRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAirdropAccountRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAirdropAccountRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryAirdropAccountResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAirdropAccountResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAirdropAccountResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ClaimAddress) > 0 {
		i -= len(m.ClaimAddress)
		copy(dAtA[i:], m.ClaimAddress)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.ClaimAddress)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Claimed {
		i--
		if m.Claimed {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x10
	}
	if m.Eligible {
		i--
		if m.Eligible {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryAirdropAccountRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryAirdropAccountResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Eligible {
		n += 2
	}
	if m.Claimed {
		n += 2
	}
	l = len(m.ClaimAddress)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryAirdropAccountRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryAirdropAccountRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAirdropAccountRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryAirdropAccountResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryAirdropAccountResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAirdropAccountResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Eligible", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
			m.Eligible = bool(v != 0)
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Claimed", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
			m.Claimed = bool(v != 0)
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClaimAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClaimAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
