// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cosmos/accounts/interfaces/account_abstraction/v1/interface.proto

package v1

import (
	fmt "fmt"
	tx "github.com/depinnetwork/depin-sdk/types/tx"
	proto "github.com/cosmos/gogoproto/proto"
	any "github.com/cosmos/gogoproto/types/any"
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

// MsgAuthenticate is a message that an x/account account abstraction implementer
// must handle to authenticate a transaction. Always ensure the caller is the Accounts module.
type MsgAuthenticate struct {
	// bundler defines the address of the bundler that sent the operation.
	// NOTE: in case the operation was sent directly by the user, this field will reflect
	// the user address.
	Bundler string `protobuf:"bytes,1,opt,name=bundler,proto3" json:"bundler,omitempty"`
	// raw_tx defines the raw version of the tx, this is useful to compute the signature quickly.
	RawTx *tx.TxRaw `protobuf:"bytes,2,opt,name=raw_tx,json=rawTx,proto3" json:"raw_tx,omitempty"`
	// tx defines the decoded version of the tx, coming from raw_tx.
	Tx *tx.Tx `protobuf:"bytes,3,opt,name=tx,proto3" json:"tx,omitempty"`
	// signer_index defines the index of the signer in the tx.
	// Specifically this can be used to extract the signature at the correct
	// index.
	SignerIndex uint32 `protobuf:"varint,4,opt,name=signer_index,json=signerIndex,proto3" json:"signer_index,omitempty"`
}

func (m *MsgAuthenticate) Reset()         { *m = MsgAuthenticate{} }
func (m *MsgAuthenticate) String() string { return proto.CompactTextString(m) }
func (*MsgAuthenticate) ProtoMessage()    {}
func (*MsgAuthenticate) Descriptor() ([]byte, []int) {
	return fileDescriptor_56b360422260e9d1, []int{0}
}
func (m *MsgAuthenticate) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgAuthenticate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgAuthenticate.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgAuthenticate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgAuthenticate.Merge(m, src)
}
func (m *MsgAuthenticate) XXX_Size() int {
	return m.Size()
}
func (m *MsgAuthenticate) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgAuthenticate.DiscardUnknown(m)
}

var xxx_messageInfo_MsgAuthenticate proto.InternalMessageInfo

func (m *MsgAuthenticate) GetBundler() string {
	if m != nil {
		return m.Bundler
	}
	return ""
}

func (m *MsgAuthenticate) GetRawTx() *tx.TxRaw {
	if m != nil {
		return m.RawTx
	}
	return nil
}

func (m *MsgAuthenticate) GetTx() *tx.Tx {
	if m != nil {
		return m.Tx
	}
	return nil
}

func (m *MsgAuthenticate) GetSignerIndex() uint32 {
	if m != nil {
		return m.SignerIndex
	}
	return 0
}

// MsgAuthenticateResponse is the response to MsgAuthenticate.
// The authentication either fails or succeeds, this is why
// there are no auxiliary fields to the response.
type MsgAuthenticateResponse struct {
}

func (m *MsgAuthenticateResponse) Reset()         { *m = MsgAuthenticateResponse{} }
func (m *MsgAuthenticateResponse) String() string { return proto.CompactTextString(m) }
func (*MsgAuthenticateResponse) ProtoMessage()    {}
func (*MsgAuthenticateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_56b360422260e9d1, []int{1}
}
func (m *MsgAuthenticateResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgAuthenticateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgAuthenticateResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgAuthenticateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgAuthenticateResponse.Merge(m, src)
}
func (m *MsgAuthenticateResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgAuthenticateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgAuthenticateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgAuthenticateResponse proto.InternalMessageInfo

// QueryAuthenticationMethods is a query that an x/account account abstraction implementer
// must handle to return the authentication methods that the account supports.
type QueryAuthenticationMethods struct {
}

func (m *QueryAuthenticationMethods) Reset()         { *m = QueryAuthenticationMethods{} }
func (m *QueryAuthenticationMethods) String() string { return proto.CompactTextString(m) }
func (*QueryAuthenticationMethods) ProtoMessage()    {}
func (*QueryAuthenticationMethods) Descriptor() ([]byte, []int) {
	return fileDescriptor_56b360422260e9d1, []int{2}
}
func (m *QueryAuthenticationMethods) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAuthenticationMethods) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAuthenticationMethods.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAuthenticationMethods) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAuthenticationMethods.Merge(m, src)
}
func (m *QueryAuthenticationMethods) XXX_Size() int {
	return m.Size()
}
func (m *QueryAuthenticationMethods) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAuthenticationMethods.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAuthenticationMethods proto.InternalMessageInfo

// QueryAuthenticationMethodsResponse is the response to QueryAuthenticationMethods.
type QueryAuthenticationMethodsResponse struct {
	// authentication_methods are the authentication methods that the account supports.
	AuthenticationMethods []string `protobuf:"bytes,1,rep,name=authentication_methods,json=authenticationMethods,proto3" json:"authentication_methods,omitempty"`
}

func (m *QueryAuthenticationMethodsResponse) Reset()         { *m = QueryAuthenticationMethodsResponse{} }
func (m *QueryAuthenticationMethodsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryAuthenticationMethodsResponse) ProtoMessage()    {}
func (*QueryAuthenticationMethodsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_56b360422260e9d1, []int{3}
}
func (m *QueryAuthenticationMethodsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAuthenticationMethodsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAuthenticationMethodsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAuthenticationMethodsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAuthenticationMethodsResponse.Merge(m, src)
}
func (m *QueryAuthenticationMethodsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryAuthenticationMethodsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAuthenticationMethodsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAuthenticationMethodsResponse proto.InternalMessageInfo

func (m *QueryAuthenticationMethodsResponse) GetAuthenticationMethods() []string {
	if m != nil {
		return m.AuthenticationMethods
	}
	return nil
}

// TxExtension is the extension option that AA's add to txs when they're bundled.
type TxExtension struct {
	// authentication_gas_limit expresses the gas limit to be used for the authentication part of the
	// bundled tx.
	AuthenticationGasLimit uint64 `protobuf:"varint,1,opt,name=authentication_gas_limit,json=authenticationGasLimit,proto3" json:"authentication_gas_limit,omitempty"`
	// bundler_payment_messages expresses a list of messages that the account
	// executes to pay the bundler for submitting the bundled tx.
	// It can be empty if the bundler does not need any form of payment,
	// the handshake for submitting the UserOperation might have happened off-chain.
	// Bundlers and accounts are free to use any form of payment, in fact the payment can
	// either be empty or be expressed as:
	// - NFT payment
	// - IBC Token payment.
	// - Payment through delegations.
	BundlerPaymentMessages []*any.Any `protobuf:"bytes,2,rep,name=bundler_payment_messages,json=bundlerPaymentMessages,proto3" json:"bundler_payment_messages,omitempty"`
	// bundler_payment_gas_limit defines the gas limit to be used for the bundler payment.
	// This ensures that, since the bundler executes a list of bundled tx and there needs to
	// be minimal trust between bundler and the tx sender, the sender cannot consume
	// the whole bundle gas.
	BundlerPaymentGasLimit uint64 `protobuf:"varint,3,opt,name=bundler_payment_gas_limit,json=bundlerPaymentGasLimit,proto3" json:"bundler_payment_gas_limit,omitempty"`
	// execution_gas_limit defines the gas limit to be used for the execution of the UserOperation's
	// execution messages.
	ExecutionGasLimit uint64 `protobuf:"varint,4,opt,name=execution_gas_limit,json=executionGasLimit,proto3" json:"execution_gas_limit,omitempty"`
}

func (m *TxExtension) Reset()         { *m = TxExtension{} }
func (m *TxExtension) String() string { return proto.CompactTextString(m) }
func (*TxExtension) ProtoMessage()    {}
func (*TxExtension) Descriptor() ([]byte, []int) {
	return fileDescriptor_56b360422260e9d1, []int{4}
}
func (m *TxExtension) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TxExtension) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TxExtension.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TxExtension) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TxExtension.Merge(m, src)
}
func (m *TxExtension) XXX_Size() int {
	return m.Size()
}
func (m *TxExtension) XXX_DiscardUnknown() {
	xxx_messageInfo_TxExtension.DiscardUnknown(m)
}

var xxx_messageInfo_TxExtension proto.InternalMessageInfo

func (m *TxExtension) GetAuthenticationGasLimit() uint64 {
	if m != nil {
		return m.AuthenticationGasLimit
	}
	return 0
}

func (m *TxExtension) GetBundlerPaymentMessages() []*any.Any {
	if m != nil {
		return m.BundlerPaymentMessages
	}
	return nil
}

func (m *TxExtension) GetBundlerPaymentGasLimit() uint64 {
	if m != nil {
		return m.BundlerPaymentGasLimit
	}
	return 0
}

func (m *TxExtension) GetExecutionGasLimit() uint64 {
	if m != nil {
		return m.ExecutionGasLimit
	}
	return 0
}

func init() {
	proto.RegisterType((*MsgAuthenticate)(nil), "cosmos.accounts.interfaces.account_abstraction.v1.MsgAuthenticate")
	proto.RegisterType((*MsgAuthenticateResponse)(nil), "cosmos.accounts.interfaces.account_abstraction.v1.MsgAuthenticateResponse")
	proto.RegisterType((*QueryAuthenticationMethods)(nil), "cosmos.accounts.interfaces.account_abstraction.v1.QueryAuthenticationMethods")
	proto.RegisterType((*QueryAuthenticationMethodsResponse)(nil), "cosmos.accounts.interfaces.account_abstraction.v1.QueryAuthenticationMethodsResponse")
	proto.RegisterType((*TxExtension)(nil), "cosmos.accounts.interfaces.account_abstraction.v1.TxExtension")
}

func init() {
	proto.RegisterFile("cosmos/accounts/interfaces/account_abstraction/v1/interface.proto", fileDescriptor_56b360422260e9d1)
}

var fileDescriptor_56b360422260e9d1 = []byte{
	// 467 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0x4f, 0x8b, 0xd3, 0x40,
	0x18, 0xc6, 0x9b, 0xb6, 0xae, 0xec, 0x54, 0x11, 0xa3, 0xbb, 0x4e, 0x8b, 0x84, 0x5a, 0x10, 0x72,
	0x9a, 0x90, 0x15, 0xc1, 0x3d, 0x56, 0x10, 0x11, 0xac, 0x68, 0xec, 0x49, 0x0f, 0x61, 0x92, 0xbe,
	0x9b, 0x1d, 0x6c, 0x66, 0x4a, 0xe6, 0x4d, 0x3b, 0xbd, 0xfa, 0x09, 0xfc, 0x14, 0x7e, 0x16, 0x8f,
	0x7b, 0xf4, 0x28, 0xed, 0x17, 0x91, 0xcd, 0x9f, 0xd6, 0x0d, 0xf5, 0xb0, 0xc7, 0xbc, 0xcf, 0xf3,
	0xfc, 0xf2, 0xe4, 0x7d, 0x09, 0x19, 0xc7, 0x4a, 0xa7, 0x4a, 0x7b, 0x3c, 0x8e, 0x55, 0x2e, 0x51,
	0x7b, 0x42, 0x22, 0x64, 0x17, 0x3c, 0x86, 0xdd, 0x2c, 0xe4, 0x91, 0xc6, 0x8c, 0xc7, 0x28, 0x94,
	0xf4, 0x96, 0xfe, 0xde, 0xc1, 0x16, 0x99, 0x42, 0x65, 0xfb, 0x25, 0x82, 0xd5, 0x08, 0xb6, 0x47,
	0xb0, 0x03, 0x08, 0xb6, 0xf4, 0x07, 0xfd, 0x44, 0xa9, 0x64, 0x0e, 0x5e, 0x01, 0x88, 0xf2, 0x0b,
	0x8f, 0xcb, 0x75, 0x49, 0x1b, 0x0c, 0xaa, 0x42, 0x68, 0xbc, 0xa5, 0x1f, 0x01, 0x72, 0xdf, 0x43,
	0x53, 0x6a, 0xa3, 0x9f, 0x16, 0x79, 0x30, 0xd1, 0xc9, 0x38, 0xc7, 0x4b, 0x90, 0x28, 0x62, 0x8e,
	0x60, 0x53, 0x72, 0x37, 0xca, 0xe5, 0x6c, 0x0e, 0x19, 0xb5, 0x86, 0x96, 0x7b, 0x1c, 0xd4, 0x8f,
	0xb6, 0x47, 0x8e, 0x32, 0xbe, 0x0a, 0xd1, 0xd0, 0xf6, 0xd0, 0x72, 0x7b, 0x67, 0x94, 0x55, 0x45,
	0xd1, 0xb0, 0x0a, 0xcd, 0xa6, 0x26, 0xe0, 0xab, 0xe0, 0x4e, 0xc6, 0x57, 0x53, 0x63, 0x3f, 0x27,
	0x6d, 0x34, 0xb4, 0x53, 0x98, 0x4f, 0x0e, 0x9b, 0xdb, 0x68, 0xec, 0x67, 0xe4, 0x9e, 0x16, 0x89,
	0x84, 0x2c, 0x14, 0x72, 0x06, 0x86, 0x76, 0x87, 0x96, 0x7b, 0x3f, 0xe8, 0x95, 0xb3, 0x77, 0xd7,
	0xa3, 0x51, 0x9f, 0x3c, 0x69, 0xf4, 0x0c, 0x40, 0x2f, 0x94, 0xd4, 0x30, 0x7a, 0x4a, 0x06, 0x9f,
	0x72, 0xc8, 0xd6, 0xff, 0x88, 0x42, 0xc9, 0x09, 0xe0, 0xa5, 0x9a, 0xe9, 0xd1, 0x57, 0x32, 0xfa,
	0xbf, 0x5a, 0x33, 0xec, 0x97, 0xe4, 0x94, 0xdf, 0x30, 0x84, 0x69, 0xe9, 0xa0, 0xd6, 0xb0, 0xe3,
	0x1e, 0x07, 0x27, 0xfc, 0x20, 0xfc, 0x7b, 0x9b, 0xf4, 0xa6, 0xe6, 0x8d, 0x41, 0x90, 0x5a, 0x28,
	0x69, 0xbf, 0x22, 0xb4, 0x81, 0x49, 0xb8, 0x0e, 0xe7, 0x22, 0x15, 0x58, 0xec, 0xb2, 0x1b, 0x34,
	0x5e, 0xf3, 0x96, 0xeb, 0xf7, 0xd7, 0xaa, 0xfd, 0x81, 0xd0, 0x6a, 0xcb, 0xe1, 0x82, 0xaf, 0x53,
	0x90, 0x18, 0xa6, 0xa0, 0x35, 0x4f, 0x40, 0xd3, 0xf6, 0xb0, 0xe3, 0xf6, 0xce, 0x1e, 0xb3, 0xf2,
	0xc4, 0xac, 0x3e, 0x31, 0x1b, 0xcb, 0x75, 0x70, 0x5a, 0xa5, 0x3e, 0x96, 0xa1, 0x49, 0x95, 0xb1,
	0xcf, 0x49, 0xbf, 0xc9, 0xdb, 0x57, 0xe9, 0x94, 0x55, 0x6e, 0x46, 0x77, 0x55, 0x18, 0x79, 0x04,
	0x06, 0xe2, 0xbc, 0xd1, 0xbf, 0x5b, 0x84, 0x1e, 0xee, 0xa4, 0xda, 0xff, 0xfa, 0xf3, 0xaf, 0x8d,
	0x63, 0x5d, 0x6d, 0x1c, 0xeb, 0xcf, 0xc6, 0xb1, 0x7e, 0x6c, 0x9d, 0xd6, 0xd5, 0xd6, 0x69, 0xfd,
	0xde, 0x3a, 0xad, 0x2f, 0xe7, 0xe5, 0xc5, 0xf5, 0xec, 0x1b, 0x13, 0xca, 0x33, 0xb7, 0xf8, 0x25,
	0xa2, 0xa3, 0xe2, 0x2b, 0x5f, 0xfc, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x24, 0x94, 0x3c, 0x65, 0x4e,
	0x03, 0x00, 0x00,
}

func (m *MsgAuthenticate) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgAuthenticate) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgAuthenticate) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.SignerIndex != 0 {
		i = encodeVarintInterface(dAtA, i, uint64(m.SignerIndex))
		i--
		dAtA[i] = 0x20
	}
	if m.Tx != nil {
		{
			size, err := m.Tx.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintInterface(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.RawTx != nil {
		{
			size, err := m.RawTx.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintInterface(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Bundler) > 0 {
		i -= len(m.Bundler)
		copy(dAtA[i:], m.Bundler)
		i = encodeVarintInterface(dAtA, i, uint64(len(m.Bundler)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgAuthenticateResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgAuthenticateResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgAuthenticateResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryAuthenticationMethods) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAuthenticationMethods) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAuthenticationMethods) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryAuthenticationMethodsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAuthenticationMethodsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAuthenticationMethodsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.AuthenticationMethods) > 0 {
		for iNdEx := len(m.AuthenticationMethods) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.AuthenticationMethods[iNdEx])
			copy(dAtA[i:], m.AuthenticationMethods[iNdEx])
			i = encodeVarintInterface(dAtA, i, uint64(len(m.AuthenticationMethods[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *TxExtension) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TxExtension) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TxExtension) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ExecutionGasLimit != 0 {
		i = encodeVarintInterface(dAtA, i, uint64(m.ExecutionGasLimit))
		i--
		dAtA[i] = 0x20
	}
	if m.BundlerPaymentGasLimit != 0 {
		i = encodeVarintInterface(dAtA, i, uint64(m.BundlerPaymentGasLimit))
		i--
		dAtA[i] = 0x18
	}
	if len(m.BundlerPaymentMessages) > 0 {
		for iNdEx := len(m.BundlerPaymentMessages) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.BundlerPaymentMessages[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintInterface(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if m.AuthenticationGasLimit != 0 {
		i = encodeVarintInterface(dAtA, i, uint64(m.AuthenticationGasLimit))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintInterface(dAtA []byte, offset int, v uint64) int {
	offset -= sovInterface(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgAuthenticate) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Bundler)
	if l > 0 {
		n += 1 + l + sovInterface(uint64(l))
	}
	if m.RawTx != nil {
		l = m.RawTx.Size()
		n += 1 + l + sovInterface(uint64(l))
	}
	if m.Tx != nil {
		l = m.Tx.Size()
		n += 1 + l + sovInterface(uint64(l))
	}
	if m.SignerIndex != 0 {
		n += 1 + sovInterface(uint64(m.SignerIndex))
	}
	return n
}

func (m *MsgAuthenticateResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryAuthenticationMethods) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryAuthenticationMethodsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.AuthenticationMethods) > 0 {
		for _, s := range m.AuthenticationMethods {
			l = len(s)
			n += 1 + l + sovInterface(uint64(l))
		}
	}
	return n
}

func (m *TxExtension) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.AuthenticationGasLimit != 0 {
		n += 1 + sovInterface(uint64(m.AuthenticationGasLimit))
	}
	if len(m.BundlerPaymentMessages) > 0 {
		for _, e := range m.BundlerPaymentMessages {
			l = e.Size()
			n += 1 + l + sovInterface(uint64(l))
		}
	}
	if m.BundlerPaymentGasLimit != 0 {
		n += 1 + sovInterface(uint64(m.BundlerPaymentGasLimit))
	}
	if m.ExecutionGasLimit != 0 {
		n += 1 + sovInterface(uint64(m.ExecutionGasLimit))
	}
	return n
}

func sovInterface(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozInterface(x uint64) (n int) {
	return sovInterface(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgAuthenticate) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowInterface
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
			return fmt.Errorf("proto: MsgAuthenticate: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgAuthenticate: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Bundler", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInterface
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
				return ErrInvalidLengthInterface
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthInterface
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Bundler = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RawTx", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInterface
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
				return ErrInvalidLengthInterface
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthInterface
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.RawTx == nil {
				m.RawTx = &tx.TxRaw{}
			}
			if err := m.RawTx.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tx", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInterface
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
				return ErrInvalidLengthInterface
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthInterface
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Tx == nil {
				m.Tx = &tx.Tx{}
			}
			if err := m.Tx.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SignerIndex", wireType)
			}
			m.SignerIndex = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInterface
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SignerIndex |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipInterface(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthInterface
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
func (m *MsgAuthenticateResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowInterface
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
			return fmt.Errorf("proto: MsgAuthenticateResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgAuthenticateResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipInterface(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthInterface
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
func (m *QueryAuthenticationMethods) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowInterface
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
			return fmt.Errorf("proto: QueryAuthenticationMethods: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAuthenticationMethods: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipInterface(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthInterface
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
func (m *QueryAuthenticationMethodsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowInterface
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
			return fmt.Errorf("proto: QueryAuthenticationMethodsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAuthenticationMethodsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AuthenticationMethods", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInterface
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
				return ErrInvalidLengthInterface
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthInterface
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AuthenticationMethods = append(m.AuthenticationMethods, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipInterface(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthInterface
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
func (m *TxExtension) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowInterface
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
			return fmt.Errorf("proto: TxExtension: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TxExtension: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AuthenticationGasLimit", wireType)
			}
			m.AuthenticationGasLimit = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInterface
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AuthenticationGasLimit |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BundlerPaymentMessages", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInterface
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
				return ErrInvalidLengthInterface
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthInterface
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BundlerPaymentMessages = append(m.BundlerPaymentMessages, &any.Any{})
			if err := m.BundlerPaymentMessages[len(m.BundlerPaymentMessages)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BundlerPaymentGasLimit", wireType)
			}
			m.BundlerPaymentGasLimit = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInterface
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BundlerPaymentGasLimit |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExecutionGasLimit", wireType)
			}
			m.ExecutionGasLimit = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInterface
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ExecutionGasLimit |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipInterface(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthInterface
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
func skipInterface(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowInterface
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
					return 0, ErrIntOverflowInterface
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
					return 0, ErrIntOverflowInterface
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
				return 0, ErrInvalidLengthInterface
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupInterface
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthInterface
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthInterface        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowInterface          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupInterface = fmt.Errorf("proto: unexpected end of group")
)
