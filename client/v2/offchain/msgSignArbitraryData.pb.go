// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: offchain/msgSignArbitraryData.proto

package offchain

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/depinnetwork/depin-sdk/types/msgservice"
	_ "github.com/depinnetwork/depin-sdk/types/tx/amino"
	proto "github.com/cosmos/gogoproto/proto"
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

// MsgSignArbitraryData defines an arbitrary, general-purpose, off-chain message
type MsgSignArbitraryData struct {
	// AppDomain is the application requesting off-chain message signing
	AppDomain string `protobuf:"bytes,1,opt,name=app_domain,json=appDomain,proto3" json:"app_domain,omitempty"`
	// Signer is the sdk.AccAddress of the message signer
	Signer string `protobuf:"bytes,2,opt,name=signer,proto3" json:"signer,omitempty"`
	// Data represents the raw bytes of the content that is signed (text, json, etc)
	Data string `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *MsgSignArbitraryData) Reset()         { *m = MsgSignArbitraryData{} }
func (m *MsgSignArbitraryData) String() string { return proto.CompactTextString(m) }
func (*MsgSignArbitraryData) ProtoMessage()    {}
func (*MsgSignArbitraryData) Descriptor() ([]byte, []int) {
	return fileDescriptor_f3e1b1b538b29252, []int{0}
}
func (m *MsgSignArbitraryData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSignArbitraryData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSignArbitraryData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSignArbitraryData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSignArbitraryData.Merge(m, src)
}
func (m *MsgSignArbitraryData) XXX_Size() int {
	return m.Size()
}
func (m *MsgSignArbitraryData) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSignArbitraryData.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSignArbitraryData proto.InternalMessageInfo

func (m *MsgSignArbitraryData) GetAppDomain() string {
	if m != nil {
		return m.AppDomain
	}
	return ""
}

func (m *MsgSignArbitraryData) GetSigner() string {
	if m != nil {
		return m.Signer
	}
	return ""
}

func (m *MsgSignArbitraryData) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func init() {
	proto.RegisterType((*MsgSignArbitraryData)(nil), "offchain.MsgSignArbitraryData")
}

func init() {
	proto.RegisterFile("offchain/msgSignArbitraryData.proto", fileDescriptor_f3e1b1b538b29252)
}

var fileDescriptor_f3e1b1b538b29252 = []byte{
	// 267 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0xce, 0x4f, 0x4b, 0x4b,
	0xce, 0x48, 0xcc, 0xcc, 0xd3, 0xcf, 0x2d, 0x4e, 0x0f, 0xce, 0x4c, 0xcf, 0x73, 0x2c, 0x4a, 0xca,
	0x2c, 0x29, 0x4a, 0x2c, 0xaa, 0x74, 0x49, 0x2c, 0x49, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0xe2, 0x80, 0x29, 0x92, 0x92, 0x4c, 0xce, 0x2f, 0xce, 0xcd, 0x2f, 0x8e, 0x07, 0x8b, 0xeb, 0x43,
	0x38, 0x10, 0x45, 0x52, 0xe2, 0x10, 0x1e, 0xc8, 0x1c, 0xfd, 0x32, 0x43, 0x10, 0x05, 0x95, 0x10,
	0x4c, 0xcc, 0xcd, 0xcc, 0xcb, 0xd7, 0x07, 0x93, 0x10, 0x21, 0xa5, 0x55, 0x8c, 0x5c, 0x22, 0xbe,
	0x58, 0xec, 0x13, 0x92, 0xe5, 0xe2, 0x4a, 0x2c, 0x28, 0x88, 0x4f, 0xc9, 0xcf, 0x4d, 0xcc, 0xcc,
	0x93, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0xe2, 0x4c, 0x2c, 0x28, 0x70, 0x01, 0x0b, 0x08, 0x19,
	0x70, 0xb1, 0x15, 0x67, 0xa6, 0xe7, 0xa5, 0x16, 0x49, 0x30, 0x81, 0xa4, 0x9c, 0x24, 0x2e, 0x6d,
	0xd1, 0x15, 0x81, 0xba, 0xc2, 0x31, 0x25, 0xa5, 0x28, 0xb5, 0xb8, 0x38, 0xb8, 0xa4, 0x28, 0x33,
	0x2f, 0x3d, 0x08, 0xaa, 0x4e, 0x48, 0x88, 0x8b, 0x25, 0x25, 0xb1, 0x24, 0x51, 0x82, 0x19, 0x6c,
	0x14, 0x98, 0x6d, 0xa5, 0xdb, 0xf4, 0x7c, 0x83, 0x16, 0x54, 0x41, 0xd7, 0xf3, 0x0d, 0x5a, 0xb2,
	0xf0, 0x30, 0xc0, 0xe6, 0x26, 0x27, 0xcb, 0x13, 0x8f, 0xe4, 0x18, 0x2f, 0x3c, 0x92, 0x63, 0x7c,
	0xf0, 0x48, 0x8e, 0x71, 0xc2, 0x63, 0x39, 0x86, 0x0b, 0x8f, 0xe5, 0x18, 0x6e, 0x3c, 0x96, 0x63,
	0x88, 0x92, 0x87, 0x58, 0x5d, 0x9c, 0x92, 0xad, 0x97, 0x99, 0xaf, 0x9f, 0x9c, 0x93, 0x99, 0x9a,
	0x57, 0xa2, 0x5f, 0x66, 0xa4, 0x0f, 0x33, 0x2f, 0x89, 0x0d, 0xec, 0x5d, 0x63, 0x40, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x78, 0x04, 0xe8, 0x80, 0x66, 0x01, 0x00, 0x00,
}

func (m *MsgSignArbitraryData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSignArbitraryData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSignArbitraryData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Data) > 0 {
		i -= len(m.Data)
		copy(dAtA[i:], m.Data)
		i = encodeVarintMsgSignArbitraryData(dAtA, i, uint64(len(m.Data)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Signer) > 0 {
		i -= len(m.Signer)
		copy(dAtA[i:], m.Signer)
		i = encodeVarintMsgSignArbitraryData(dAtA, i, uint64(len(m.Signer)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.AppDomain) > 0 {
		i -= len(m.AppDomain)
		copy(dAtA[i:], m.AppDomain)
		i = encodeVarintMsgSignArbitraryData(dAtA, i, uint64(len(m.AppDomain)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintMsgSignArbitraryData(dAtA []byte, offset int, v uint64) int {
	offset -= sovMsgSignArbitraryData(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgSignArbitraryData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.AppDomain)
	if l > 0 {
		n += 1 + l + sovMsgSignArbitraryData(uint64(l))
	}
	l = len(m.Signer)
	if l > 0 {
		n += 1 + l + sovMsgSignArbitraryData(uint64(l))
	}
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovMsgSignArbitraryData(uint64(l))
	}
	return n
}

func sovMsgSignArbitraryData(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMsgSignArbitraryData(x uint64) (n int) {
	return sovMsgSignArbitraryData(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgSignArbitraryData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsgSignArbitraryData
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
			return fmt.Errorf("proto: MsgSignArbitraryData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSignArbitraryData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AppDomain", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgSignArbitraryData
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
				return ErrInvalidLengthMsgSignArbitraryData
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsgSignArbitraryData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AppDomain = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgSignArbitraryData
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
				return ErrInvalidLengthMsgSignArbitraryData
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsgSignArbitraryData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgSignArbitraryData
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
				return ErrInvalidLengthMsgSignArbitraryData
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsgSignArbitraryData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMsgSignArbitraryData(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMsgSignArbitraryData
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
func skipMsgSignArbitraryData(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMsgSignArbitraryData
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
					return 0, ErrIntOverflowMsgSignArbitraryData
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
					return 0, ErrIntOverflowMsgSignArbitraryData
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
				return 0, ErrInvalidLengthMsgSignArbitraryData
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMsgSignArbitraryData
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMsgSignArbitraryData
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMsgSignArbitraryData        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMsgSignArbitraryData          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMsgSignArbitraryData = fmt.Errorf("proto: unexpected end of group")
)
