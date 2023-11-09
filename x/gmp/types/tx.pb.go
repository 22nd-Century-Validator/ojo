// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ojo/gmp/v1/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/cosmos-sdk/types/msgservice"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	_ "github.com/gogo/protobuf/gogoproto"
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

// MsgSetParams defines the SetParams message type.
type MsgSetParams struct {
	// authority is the address that controls the module (defaults to x/gov unless overwritten).
	Authority string `protobuf:"bytes,1,opt,name=authority,proto3" json:"authority,omitempty"`
	// params defines the gmp parameters to update.
	Params *Params `protobuf:"bytes,2,opt,name=params,proto3" json:"params,omitempty"`
}

func (m *MsgSetParams) Reset()         { *m = MsgSetParams{} }
func (m *MsgSetParams) String() string { return proto.CompactTextString(m) }
func (*MsgSetParams) ProtoMessage()    {}
func (*MsgSetParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_5f28b599a8e3f91d, []int{0}
}
func (m *MsgSetParams) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSetParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSetParams.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSetParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSetParams.Merge(m, src)
}
func (m *MsgSetParams) XXX_Size() int {
	return m.Size()
}
func (m *MsgSetParams) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSetParams.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSetParams proto.InternalMessageInfo

// MsgSetParamsResponse defines the SetParams response type.
type MsgSetParamsResponse struct {
}

func (m *MsgSetParamsResponse) Reset()         { *m = MsgSetParamsResponse{} }
func (m *MsgSetParamsResponse) String() string { return proto.CompactTextString(m) }
func (*MsgSetParamsResponse) ProtoMessage()    {}
func (*MsgSetParamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5f28b599a8e3f91d, []int{1}
}
func (m *MsgSetParamsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSetParamsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSetParamsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSetParamsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSetParamsResponse.Merge(m, src)
}
func (m *MsgSetParamsResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgSetParamsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSetParamsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSetParamsResponse proto.InternalMessageInfo

// MsgRelayPrice defines the Relay message type.
type MsgRelayPrice struct {
	// authority is the address that signs the message.
	Relayer string `protobuf:"bytes,1,opt,name=relayer,proto3" json:"relayer,omitempty"`
	// destination_chain defines the chain which this will be relayed to.
	DestinationChain string `protobuf:"bytes,2,opt,name=destination_chain,json=destinationChain,proto3" json:"destination_chain,omitempty"`
	// destination_address defines the destination contract address to call.
	DestinationAddress string `protobuf:"bytes,3,opt,name=destination_address,json=destinationAddress,proto3" json:"destination_address,omitempty"`
	// denoms defines the denoms that the user wants to relay via GMP.
	Denoms []string `protobuf:"bytes,4,rep,name=denoms,proto3" json:"denoms,omitempty"`
	// token determines the IBC token that the user wants to relay via GMP.
	Token types.Coin `protobuf:"bytes,5,opt,name=token,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"token"`
}

func (m *MsgRelayPrice) Reset()         { *m = MsgRelayPrice{} }
func (m *MsgRelayPrice) String() string { return proto.CompactTextString(m) }
func (*MsgRelayPrice) ProtoMessage()    {}
func (*MsgRelayPrice) Descriptor() ([]byte, []int) {
	return fileDescriptor_5f28b599a8e3f91d, []int{2}
}
func (m *MsgRelayPrice) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgRelayPrice) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgRelayPrice.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgRelayPrice) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgRelayPrice.Merge(m, src)
}
func (m *MsgRelayPrice) XXX_Size() int {
	return m.Size()
}
func (m *MsgRelayPrice) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgRelayPrice.DiscardUnknown(m)
}

var xxx_messageInfo_MsgRelayPrice proto.InternalMessageInfo

// MsgRelay defines the Relay response type.
type MsgRelayPriceResponse struct {
}

func (m *MsgRelayPriceResponse) Reset()         { *m = MsgRelayPriceResponse{} }
func (m *MsgRelayPriceResponse) String() string { return proto.CompactTextString(m) }
func (*MsgRelayPriceResponse) ProtoMessage()    {}
func (*MsgRelayPriceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5f28b599a8e3f91d, []int{3}
}
func (m *MsgRelayPriceResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgRelayPriceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgRelayPriceResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgRelayPriceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgRelayPriceResponse.Merge(m, src)
}
func (m *MsgRelayPriceResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgRelayPriceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgRelayPriceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgRelayPriceResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgSetParams)(nil), "ojo.gmp.v1.MsgSetParams")
	proto.RegisterType((*MsgSetParamsResponse)(nil), "ojo.gmp.v1.MsgSetParamsResponse")
	proto.RegisterType((*MsgRelayPrice)(nil), "ojo.gmp.v1.MsgRelayPrice")
	proto.RegisterType((*MsgRelayPriceResponse)(nil), "ojo.gmp.v1.MsgRelayPriceResponse")
}

func init() { proto.RegisterFile("ojo/gmp/v1/tx.proto", fileDescriptor_5f28b599a8e3f91d) }

var fileDescriptor_5f28b599a8e3f91d = []byte{
	// 509 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x53, 0xb1, 0x6f, 0xd3, 0x4e,
	0x14, 0xb6, 0x9b, 0x5f, 0xf3, 0x53, 0x0e, 0x90, 0xe0, 0x1a, 0x52, 0x27, 0x83, 0x13, 0x32, 0xa0,
	0xa8, 0x28, 0x3e, 0x12, 0x24, 0x86, 0x4e, 0x90, 0x0e, 0x48, 0x48, 0x91, 0x2a, 0x77, 0x63, 0xa9,
	0x2e, 0xf6, 0xe9, 0xe2, 0x04, 0xdf, 0xb3, 0x7c, 0xd7, 0xd0, 0xac, 0x4c, 0x88, 0x89, 0x99, 0xa9,
	0x33, 0x13, 0x42, 0xfc, 0x11, 0x1d, 0x2b, 0x26, 0x26, 0x40, 0xc9, 0x00, 0x7f, 0x06, 0xb2, 0x7d,
	0x96, 0x0f, 0xa4, 0x8a, 0x29, 0xf7, 0xde, 0xf7, 0xbd, 0xef, 0xde, 0x7d, 0x5f, 0x8c, 0xf6, 0x60,
	0x01, 0x84, 0xc7, 0x09, 0x59, 0x8d, 0x88, 0x3a, 0xf7, 0x92, 0x14, 0x14, 0x60, 0x04, 0x0b, 0xf0,
	0x78, 0x9c, 0x78, 0xab, 0x51, 0xa7, 0xc9, 0x81, 0x43, 0xde, 0x26, 0xd9, 0xa9, 0x60, 0x74, 0x9a,
	0xc6, 0x58, 0x46, 0x2c, 0xba, 0xed, 0x00, 0x64, 0x0c, 0xf2, 0xb4, 0xa0, 0x17, 0x85, 0x86, 0xf6,
	0x8b, 0x8a, 0xc4, 0x92, 0x67, 0x33, 0xb1, 0xe4, 0x1a, 0x70, 0x35, 0x30, 0xa3, 0x92, 0x91, 0xd5,
	0x68, 0xc6, 0x14, 0x1d, 0x91, 0x00, 0x22, 0x51, 0xe0, 0xfd, 0xb7, 0x36, 0xba, 0x39, 0x95, 0xfc,
	0x84, 0xa9, 0x63, 0x9a, 0xd2, 0x58, 0xe2, 0xc7, 0xa8, 0x41, 0xcf, 0xd4, 0x1c, 0xd2, 0x48, 0xad,
	0x1d, 0xbb, 0x67, 0x0f, 0x1a, 0x13, 0xe7, 0xcb, 0xe7, 0x61, 0x53, 0x5f, 0xf7, 0x34, 0x0c, 0x53,
	0x26, 0xe5, 0x89, 0x4a, 0x23, 0xc1, 0xfd, 0x8a, 0x8a, 0x0f, 0x50, 0x3d, 0xc9, 0x15, 0x9c, 0x9d,
	0x9e, 0x3d, 0xb8, 0x31, 0xc6, 0x5e, 0xf5, 0x4a, 0xaf, 0xd0, 0xf6, 0x35, 0xe3, 0xb0, 0xf5, 0xe6,
	0xa2, 0x6b, 0xfd, 0xba, 0xe8, 0x5a, 0xaf, 0x7f, 0x7e, 0x3c, 0xa8, 0x34, 0xfa, 0x2d, 0xd4, 0x34,
	0x77, 0xf1, 0x99, 0x4c, 0x40, 0x48, 0xd6, 0xff, 0xb4, 0x83, 0x6e, 0x4d, 0x25, 0xf7, 0xd9, 0x4b,
	0xba, 0x3e, 0x4e, 0xa3, 0x80, 0xe1, 0x31, 0xfa, 0x3f, 0xcd, 0x2a, 0x96, 0xfe, 0x73, 0xc7, 0x92,
	0x88, 0x1f, 0xa0, 0x3b, 0x21, 0x93, 0x2a, 0x12, 0x54, 0x45, 0x20, 0x4e, 0x83, 0x39, 0x8d, 0x44,
	0xbe, 0x6c, 0xc3, 0xbf, 0x6d, 0x00, 0x47, 0x59, 0x1f, 0x13, 0xb4, 0x67, 0x92, 0x69, 0x21, 0xe9,
	0xd4, 0x72, 0x3a, 0x36, 0x20, 0x7d, 0x19, 0x6e, 0xa1, 0x7a, 0xc8, 0x04, 0xc4, 0xd2, 0xf9, 0xaf,
	0x57, 0x1b, 0x34, 0x7c, 0x5d, 0x61, 0x8a, 0x76, 0x15, 0x2c, 0x99, 0x70, 0x76, 0x73, 0x5b, 0xda,
	0x9e, 0x5e, 0x32, 0x0b, 0xc4, 0xd3, 0x81, 0x78, 0x47, 0x10, 0x89, 0xc9, 0xc3, 0xcb, 0x6f, 0x5d,
	0xeb, 0xc3, 0xf7, 0xee, 0x80, 0x47, 0x6a, 0x7e, 0x36, 0xf3, 0x02, 0x88, 0x75, 0xc8, 0xfa, 0x67,
	0x28, 0xc3, 0x25, 0x51, 0xeb, 0x84, 0xc9, 0x7c, 0x40, 0xfa, 0x85, 0xf2, 0x61, 0xd3, 0xb4, 0xb3,
	0x7c, 0x6e, 0x7f, 0x1f, 0xdd, 0xfd, 0xc3, 0xb3, 0xd2, 0xcd, 0xf1, 0x7b, 0x1b, 0xd5, 0xa6, 0x92,
	0xe3, 0x67, 0xa8, 0x51, 0xc5, 0xee, 0x98, 0x71, 0x99, 0x21, 0x74, 0x7a, 0xd7, 0x21, 0xa5, 0x20,
	0x7e, 0x8e, 0x90, 0x11, 0x4d, 0xfb, 0x2f, 0x7e, 0x05, 0x75, 0xee, 0x5d, 0x0b, 0x95, 0x5a, 0x93,
	0x27, 0x97, 0x1b, 0xd7, 0xbe, 0xda, 0xb8, 0xf6, 0x8f, 0x8d, 0x6b, 0xbf, 0xdb, 0xba, 0xd6, 0xd5,
	0xd6, 0xb5, 0xbe, 0x6e, 0x5d, 0xeb, 0xc5, 0x7d, 0xc3, 0x16, 0x58, 0xc0, 0x50, 0x30, 0xf5, 0x0a,
	0xd2, 0x65, 0x76, 0x26, 0xe7, 0xf9, 0xc7, 0x92, 0x5b, 0x33, 0xab, 0xe7, 0x7f, 0xec, 0x47, 0xbf,
	0x03, 0x00, 0x00, 0xff, 0xff, 0x63, 0x69, 0x51, 0x3b, 0x7b, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	// SetParams sets the parameters for the gmp module.
	SetParams(ctx context.Context, in *MsgSetParams, opts ...grpc.CallOption) (*MsgSetParamsResponse, error)
	// Relay relays Ojo data via GMP.
	RelayPrice(ctx context.Context, in *MsgRelayPrice, opts ...grpc.CallOption) (*MsgRelayPriceResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) SetParams(ctx context.Context, in *MsgSetParams, opts ...grpc.CallOption) (*MsgSetParamsResponse, error) {
	out := new(MsgSetParamsResponse)
	err := c.cc.Invoke(ctx, "/ojo.gmp.v1.Msg/SetParams", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) RelayPrice(ctx context.Context, in *MsgRelayPrice, opts ...grpc.CallOption) (*MsgRelayPriceResponse, error) {
	out := new(MsgRelayPriceResponse)
	err := c.cc.Invoke(ctx, "/ojo.gmp.v1.Msg/RelayPrice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// SetParams sets the parameters for the gmp module.
	SetParams(context.Context, *MsgSetParams) (*MsgSetParamsResponse, error)
	// Relay relays Ojo data via GMP.
	RelayPrice(context.Context, *MsgRelayPrice) (*MsgRelayPriceResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) SetParams(ctx context.Context, req *MsgSetParams) (*MsgSetParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetParams not implemented")
}
func (*UnimplementedMsgServer) RelayPrice(ctx context.Context, req *MsgRelayPrice) (*MsgRelayPriceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RelayPrice not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_SetParams_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgSetParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).SetParams(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ojo.gmp.v1.Msg/SetParams",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).SetParams(ctx, req.(*MsgSetParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_RelayPrice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgRelayPrice)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).RelayPrice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ojo.gmp.v1.Msg/RelayPrice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).RelayPrice(ctx, req.(*MsgRelayPrice))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ojo.gmp.v1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetParams",
			Handler:    _Msg_SetParams_Handler,
		},
		{
			MethodName: "RelayPrice",
			Handler:    _Msg_RelayPrice_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ojo/gmp/v1/tx.proto",
}

func (m *MsgSetParams) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSetParams) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSetParams) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Params != nil {
		{
			size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTx(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Authority) > 0 {
		i -= len(m.Authority)
		copy(dAtA[i:], m.Authority)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Authority)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgSetParamsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSetParamsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSetParamsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgRelayPrice) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgRelayPrice) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgRelayPrice) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Token.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	if len(m.Denoms) > 0 {
		for iNdEx := len(m.Denoms) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Denoms[iNdEx])
			copy(dAtA[i:], m.Denoms[iNdEx])
			i = encodeVarintTx(dAtA, i, uint64(len(m.Denoms[iNdEx])))
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.DestinationAddress) > 0 {
		i -= len(m.DestinationAddress)
		copy(dAtA[i:], m.DestinationAddress)
		i = encodeVarintTx(dAtA, i, uint64(len(m.DestinationAddress)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.DestinationChain) > 0 {
		i -= len(m.DestinationChain)
		copy(dAtA[i:], m.DestinationChain)
		i = encodeVarintTx(dAtA, i, uint64(len(m.DestinationChain)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Relayer) > 0 {
		i -= len(m.Relayer)
		copy(dAtA[i:], m.Relayer)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Relayer)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgRelayPriceResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgRelayPriceResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgRelayPriceResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgSetParams) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Authority)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.Params != nil {
		l = m.Params.Size()
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgSetParamsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgRelayPrice) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Relayer)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.DestinationChain)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.DestinationAddress)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if len(m.Denoms) > 0 {
		for _, s := range m.Denoms {
			l = len(s)
			n += 1 + l + sovTx(uint64(l))
		}
	}
	l = m.Token.Size()
	n += 1 + l + sovTx(uint64(l))
	return n
}

func (m *MsgRelayPriceResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgSetParams) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgSetParams: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSetParams: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Authority", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Authority = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Params == nil {
				m.Params = &Params{}
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgSetParamsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgSetParamsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSetParamsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgRelayPrice) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgRelayPrice: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgRelayPrice: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Relayer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Relayer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DestinationChain", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DestinationChain = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DestinationAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DestinationAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denoms", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denoms = append(m.Denoms, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Token", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Token.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgRelayPriceResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgRelayPriceResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgRelayPriceResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
