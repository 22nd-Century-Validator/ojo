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

// MsgRelay defines the Relay message type.
type MsgRelayPrice struct {
	// authority is the address that signs the message.
	Relayer string `protobuf:"bytes,1,opt,name=relayer,proto3" json:"relayer,omitempty"`
	// destination_chain defines the chain which this will be relayed to.
	DestinationChain string `protobuf:"bytes,2,opt,name=destination_chain,json=destinationChain,proto3" json:"destination_chain,omitempty"`
	// ojo_contract_address defines the ojo contract that GMP is calling.
	OjoContractAddress string `protobuf:"bytes,3,opt,name=ojo_contract_address,json=ojoContractAddress,proto3" json:"ojo_contract_address,omitempty"`
	// client_contract_address defines the client contract that Ojo is calling.
	ClientContractAddress string `protobuf:"bytes,4,opt,name=client_contract_address,json=clientContractAddress,proto3" json:"client_contract_address,omitempty"`
	// denoms defines the denoms that the user wants to relay via GMP.
	Denoms []string `protobuf:"bytes,5,rep,name=denoms,proto3" json:"denoms,omitempty"`
	// token determines the IBC token that the user wants to relay via GMP.
	Token types.Coin `protobuf:"bytes,6,opt,name=token,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"token"`
	// command_selector defines the command to call.
	CommandSelector []byte `protobuf:"bytes,7,opt,name=command_selector,json=commandSelector,proto3" json:"command_selector,omitempty"`
	// command_params defines the command parameters to call.
	CommandParams []byte `protobuf:"bytes,8,opt,name=command_params,json=commandParams,proto3" json:"command_params,omitempty"`
	// timestamp defines the timestamp of the message in terms of the source evm chain block time.
	Timestamp int64 `protobuf:"varint,9,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
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
	// 593 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x53, 0xb1, 0x6f, 0xd3, 0x4e,
	0x18, 0xb5, 0x7f, 0xf9, 0x35, 0xc5, 0x47, 0x0b, 0xe5, 0x48, 0x5b, 0x37, 0x42, 0x8e, 0x89, 0x04,
	0x32, 0x45, 0xb5, 0x9b, 0x22, 0x75, 0xe8, 0x04, 0xed, 0x80, 0x84, 0x54, 0xa9, 0x72, 0x37, 0x96,
	0xe8, 0x62, 0x9f, 0x5c, 0x27, 0xb9, 0xfb, 0x2c, 0xdf, 0x35, 0x34, 0x2b, 0x13, 0x62, 0x62, 0x66,
	0xea, 0xcc, 0xc4, 0xc0, 0x1f, 0xd1, 0xb1, 0x62, 0x42, 0x42, 0x02, 0x94, 0x0c, 0xf0, 0x67, 0x20,
	0xdb, 0x67, 0xd9, 0x14, 0x55, 0x4c, 0xbe, 0xef, 0xbd, 0xf7, 0xbd, 0xbb, 0xf3, 0xfb, 0x0e, 0xdd,
	0x85, 0x21, 0x78, 0x11, 0x4b, 0xbc, 0x49, 0xcf, 0x93, 0x67, 0x6e, 0x92, 0x82, 0x04, 0x8c, 0x60,
	0x08, 0x6e, 0xc4, 0x12, 0x77, 0xd2, 0x6b, 0xb7, 0x22, 0x88, 0x20, 0x87, 0xbd, 0x6c, 0x55, 0x28,
	0xda, 0xad, 0x5a, 0x5b, 0x26, 0x2c, 0xd0, 0x8d, 0x00, 0x04, 0x03, 0xd1, 0x2f, 0xe4, 0x45, 0xa1,
	0xa8, 0xf5, 0xa2, 0xf2, 0x98, 0x88, 0xb2, 0x1e, 0x26, 0x22, 0x45, 0x58, 0x8a, 0x18, 0x10, 0x41,
	0xbd, 0x49, 0x6f, 0x40, 0x25, 0xe9, 0x79, 0x01, 0xc4, 0xbc, 0xe0, 0xbb, 0x6f, 0x75, 0xb4, 0x74,
	0x28, 0xa2, 0x63, 0x2a, 0x8f, 0x48, 0x4a, 0x98, 0xc0, 0xbb, 0xc8, 0x20, 0xa7, 0xf2, 0x04, 0xd2,
	0x58, 0x4e, 0x4d, 0xdd, 0xd6, 0x1d, 0x63, 0xdf, 0xfc, 0xfc, 0x69, 0xab, 0xa5, 0xb6, 0x7b, 0x16,
	0x86, 0x29, 0x15, 0xe2, 0x58, 0xa6, 0x31, 0x8f, 0xfc, 0x4a, 0x8a, 0x37, 0x51, 0x33, 0xc9, 0x1d,
	0xcc, 0xff, 0x6c, 0xdd, 0xb9, 0xb9, 0x83, 0xdd, 0xea, 0x96, 0x6e, 0xe1, 0xed, 0x2b, 0xc5, 0xde,
	0xda, 0x9b, 0xf3, 0x8e, 0xf6, 0xeb, 0xbc, 0xa3, 0xbd, 0xfe, 0xf9, 0x71, 0xb3, 0xf2, 0xe8, 0xae,
	0xa1, 0x56, 0xfd, 0x2c, 0x3e, 0x15, 0x09, 0x70, 0x41, 0xbb, 0x5f, 0x1b, 0x68, 0xf9, 0x50, 0x44,
	0x3e, 0x1d, 0x93, 0xe9, 0x51, 0x1a, 0x07, 0x14, 0xef, 0xa0, 0xc5, 0x34, 0xab, 0x68, 0xfa, 0xcf,
	0x33, 0x96, 0x42, 0xfc, 0x18, 0xdd, 0x09, 0xa9, 0x90, 0x31, 0x27, 0x32, 0x06, 0xde, 0x0f, 0x4e,
	0x48, 0xcc, 0xf3, 0xc3, 0x1a, 0xfe, 0x4a, 0x8d, 0x38, 0xc8, 0x70, 0xbc, 0x8d, 0xb2, 0x0c, 0xfa,
	0x01, 0x70, 0x99, 0x92, 0x40, 0xf6, 0x49, 0xe1, 0x69, 0x36, 0x72, 0x3d, 0x86, 0x21, 0x1c, 0x28,
	0x4a, 0xed, 0x86, 0x77, 0xd1, 0x7a, 0x30, 0x8e, 0x29, 0x97, 0x7f, 0x37, 0xfd, 0x9f, 0x37, 0xad,
	0x16, 0xf4, 0xd5, 0xbe, 0x35, 0xd4, 0x0c, 0x29, 0x07, 0x26, 0xcc, 0x05, 0xbb, 0xe1, 0x18, 0xbe,
	0xaa, 0x30, 0x41, 0x0b, 0x12, 0x46, 0x94, 0x9b, 0xcd, 0xfc, 0x7f, 0x6e, 0xb8, 0xea, 0x76, 0x59,
	0x92, 0xae, 0x4a, 0xd2, 0x3d, 0x80, 0x98, 0xef, 0x6f, 0x5f, 0x7c, 0xeb, 0x68, 0x1f, 0xbe, 0x77,
	0x9c, 0x28, 0x96, 0x27, 0xa7, 0x03, 0x37, 0x00, 0xa6, 0xa6, 0x43, 0x7d, 0xb6, 0x44, 0x38, 0xf2,
	0xe4, 0x34, 0xa1, 0x22, 0x6f, 0x10, 0x7e, 0xe1, 0x8c, 0x1f, 0xa1, 0x95, 0x00, 0x18, 0x23, 0x3c,
	0xec, 0x0b, 0x3a, 0xa6, 0x81, 0x84, 0xd4, 0x5c, 0xb4, 0x75, 0x67, 0xc9, 0xbf, 0xad, 0xf0, 0x63,
	0x05, 0xe3, 0x07, 0xe8, 0x56, 0x29, 0x55, 0x31, 0xdf, 0xc8, 0x85, 0xcb, 0x0a, 0x55, 0xd3, 0x73,
	0x0f, 0x19, 0x32, 0x66, 0x54, 0x48, 0xc2, 0x12, 0xd3, 0xb0, 0x75, 0xa7, 0xe1, 0x57, 0xc0, 0x5e,
	0xab, 0x9e, 0x7b, 0x99, 0x4b, 0x77, 0x1d, 0xad, 0xfe, 0x11, 0x6e, 0x19, 0xfb, 0xce, 0x7b, 0x1d,
	0x35, 0x0e, 0x45, 0x84, 0x9f, 0x23, 0xa3, 0x9a, 0x4f, 0xb3, 0x3e, 0x57, 0xf5, 0x69, 0x69, 0xdb,
	0xd7, 0x31, 0xa5, 0x21, 0x7e, 0x81, 0x50, 0x6d, 0x86, 0x36, 0xae, 0xe8, 0x2b, 0xaa, 0x7d, 0xff,
	0x5a, 0xaa, 0xf4, 0xda, 0x7f, 0x7a, 0x31, 0xb3, 0xf4, 0xcb, 0x99, 0xa5, 0xff, 0x98, 0x59, 0xfa,
	0xbb, 0xb9, 0xa5, 0x5d, 0xce, 0x2d, 0xed, 0xcb, 0xdc, 0xd2, 0x5e, 0x3e, 0xac, 0xc5, 0x00, 0x43,
	0xd8, 0xe2, 0x54, 0xbe, 0x82, 0x74, 0x94, 0xad, 0xbd, 0xb3, 0xfc, 0x55, 0xe7, 0x51, 0x0c, 0x9a,
	0xf9, 0x0b, 0x7c, 0xf2, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x20, 0x74, 0xcc, 0x19, 0x24, 0x04, 0x00,
	0x00,
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
	if m.Timestamp != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.Timestamp))
		i--
		dAtA[i] = 0x48
	}
	if len(m.CommandParams) > 0 {
		i -= len(m.CommandParams)
		copy(dAtA[i:], m.CommandParams)
		i = encodeVarintTx(dAtA, i, uint64(len(m.CommandParams)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.CommandSelector) > 0 {
		i -= len(m.CommandSelector)
		copy(dAtA[i:], m.CommandSelector)
		i = encodeVarintTx(dAtA, i, uint64(len(m.CommandSelector)))
		i--
		dAtA[i] = 0x3a
	}
	{
		size, err := m.Token.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	if len(m.Denoms) > 0 {
		for iNdEx := len(m.Denoms) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Denoms[iNdEx])
			copy(dAtA[i:], m.Denoms[iNdEx])
			i = encodeVarintTx(dAtA, i, uint64(len(m.Denoms[iNdEx])))
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.ClientContractAddress) > 0 {
		i -= len(m.ClientContractAddress)
		copy(dAtA[i:], m.ClientContractAddress)
		i = encodeVarintTx(dAtA, i, uint64(len(m.ClientContractAddress)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.OjoContractAddress) > 0 {
		i -= len(m.OjoContractAddress)
		copy(dAtA[i:], m.OjoContractAddress)
		i = encodeVarintTx(dAtA, i, uint64(len(m.OjoContractAddress)))
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
	l = len(m.OjoContractAddress)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.ClientContractAddress)
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
	l = len(m.CommandSelector)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.CommandParams)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.Timestamp != 0 {
		n += 1 + sovTx(uint64(m.Timestamp))
	}
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
				return fmt.Errorf("proto: wrong wireType = %d for field OjoContractAddress", wireType)
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
			m.OjoContractAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClientContractAddress", wireType)
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
			m.ClientContractAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
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
		case 6:
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
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CommandSelector", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CommandSelector = append(m.CommandSelector[:0], dAtA[iNdEx:postIndex]...)
			if m.CommandSelector == nil {
				m.CommandSelector = []byte{}
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CommandParams", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CommandParams = append(m.CommandParams[:0], dAtA[iNdEx:postIndex]...)
			if m.CommandParams == nil {
				m.CommandParams = []byte{}
			}
			iNdEx = postIndex
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			m.Timestamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Timestamp |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
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
