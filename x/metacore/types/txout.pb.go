// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: metacore/txout.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

type Txout struct {
	Creator           string   `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Id                uint64   `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	TxinHash          string   `protobuf:"bytes,3,opt,name=txinHash,proto3" json:"txinHash,omitempty"`
	SourceAsset       string   `protobuf:"bytes,4,opt,name=sourceAsset,proto3" json:"sourceAsset,omitempty"`
	SourceAmount      uint64   `protobuf:"varint,5,opt,name=sourceAmount,proto3" json:"sourceAmount,omitempty"`
	MBurnt            uint64   `protobuf:"varint,6,opt,name=mBurnt,proto3" json:"mBurnt,omitempty"`
	MMint             uint64   `protobuf:"varint,7,opt,name=mMint,proto3" json:"mMint,omitempty"`
	DestinationAsset  string   `protobuf:"bytes,8,opt,name=destinationAsset,proto3" json:"destinationAsset,omitempty"`
	DestinationAmount uint64   `protobuf:"varint,9,opt,name=destinationAmount,proto3" json:"destinationAmount,omitempty"`
	FromAddress       string   `protobuf:"bytes,10,opt,name=fromAddress,proto3" json:"fromAddress,omitempty"`
	ToAddress         string   `protobuf:"bytes,11,opt,name=toAddress,proto3" json:"toAddress,omitempty"`
	BlockHeight       uint64   `protobuf:"varint,12,opt,name=blockHeight,proto3" json:"blockHeight,omitempty"`
	Signers           []string `protobuf:"bytes,13,rep,name=signers,proto3" json:"signers,omitempty"`
	FinalizedHeight   uint64   `protobuf:"varint,14,opt,name=finalizedHeight,proto3" json:"finalizedHeight,omitempty"`
}

func (m *Txout) Reset()         { *m = Txout{} }
func (m *Txout) String() string { return proto.CompactTextString(m) }
func (*Txout) ProtoMessage()    {}
func (*Txout) Descriptor() ([]byte, []int) {
	return fileDescriptor_f61438ebc14ac1ae, []int{0}
}
func (m *Txout) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Txout) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Txout.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Txout) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Txout.Merge(m, src)
}
func (m *Txout) XXX_Size() int {
	return m.Size()
}
func (m *Txout) XXX_DiscardUnknown() {
	xxx_messageInfo_Txout.DiscardUnknown(m)
}

var xxx_messageInfo_Txout proto.InternalMessageInfo

func (m *Txout) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *Txout) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Txout) GetTxinHash() string {
	if m != nil {
		return m.TxinHash
	}
	return ""
}

func (m *Txout) GetSourceAsset() string {
	if m != nil {
		return m.SourceAsset
	}
	return ""
}

func (m *Txout) GetSourceAmount() uint64 {
	if m != nil {
		return m.SourceAmount
	}
	return 0
}

func (m *Txout) GetMBurnt() uint64 {
	if m != nil {
		return m.MBurnt
	}
	return 0
}

func (m *Txout) GetMMint() uint64 {
	if m != nil {
		return m.MMint
	}
	return 0
}

func (m *Txout) GetDestinationAsset() string {
	if m != nil {
		return m.DestinationAsset
	}
	return ""
}

func (m *Txout) GetDestinationAmount() uint64 {
	if m != nil {
		return m.DestinationAmount
	}
	return 0
}

func (m *Txout) GetFromAddress() string {
	if m != nil {
		return m.FromAddress
	}
	return ""
}

func (m *Txout) GetToAddress() string {
	if m != nil {
		return m.ToAddress
	}
	return ""
}

func (m *Txout) GetBlockHeight() uint64 {
	if m != nil {
		return m.BlockHeight
	}
	return 0
}

func (m *Txout) GetSigners() []string {
	if m != nil {
		return m.Signers
	}
	return nil
}

func (m *Txout) GetFinalizedHeight() uint64 {
	if m != nil {
		return m.FinalizedHeight
	}
	return 0
}

func init() {
	proto.RegisterType((*Txout)(nil), "MetaProtocol.metacore.metacore.Txout")
}

func init() { proto.RegisterFile("metacore/txout.proto", fileDescriptor_f61438ebc14ac1ae) }

var fileDescriptor_f61438ebc14ac1ae = []byte{
	// 378 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x92, 0xcf, 0x6e, 0xe2, 0x30,
	0x10, 0xc6, 0x13, 0xfe, 0xc7, 0xb0, 0xec, 0xae, 0x85, 0x56, 0x16, 0x5a, 0x79, 0x23, 0x4e, 0xd1,
	0x6a, 0x97, 0x48, 0xed, 0x13, 0xc0, 0x89, 0x43, 0x91, 0x2a, 0xd4, 0x53, 0x6f, 0x21, 0x31, 0xc1,
	0x2a, 0xb1, 0x91, 0x3d, 0x91, 0x68, 0x9f, 0xa2, 0x6f, 0x55, 0x8e, 0x1c, 0x7b, 0xaa, 0x2a, 0x78,
	0x91, 0x2a, 0x0e, 0x81, 0xb4, 0xdc, 0xe6, 0xfb, 0x7d, 0xe3, 0x6f, 0x3c, 0xd2, 0xa0, 0x5e, 0xc2,
	0x20, 0x08, 0xa5, 0x62, 0x3e, 0x6c, 0x64, 0x0a, 0xc3, 0xb5, 0x92, 0x20, 0x31, 0x9d, 0x32, 0x08,
	0x6e, 0xb3, 0x32, 0x94, 0xab, 0x61, 0xd1, 0x72, 0x2a, 0xfa, 0xbd, 0x58, 0xc6, 0xd2, 0xb4, 0xfa,
	0x59, 0x95, 0xbf, 0x1a, 0xbc, 0x54, 0x51, 0xfd, 0x2e, 0x4b, 0xc1, 0x04, 0x35, 0x43, 0xc5, 0x02,
	0x90, 0x8a, 0xd8, 0xae, 0xed, 0x39, 0xb3, 0x42, 0xe2, 0x2e, 0xaa, 0xf0, 0x88, 0x54, 0x5c, 0xdb,
	0xab, 0xcd, 0x2a, 0x3c, 0xc2, 0x7d, 0xd4, 0x82, 0x0d, 0x17, 0x93, 0x40, 0x2f, 0x49, 0xd5, 0xb4,
	0x9e, 0x34, 0x76, 0x51, 0x5b, 0xcb, 0x54, 0x85, 0x6c, 0xa4, 0x35, 0x03, 0x52, 0x33, 0x76, 0x19,
	0xe1, 0x01, 0xea, 0x1c, 0x65, 0x22, 0x53, 0x01, 0xa4, 0x6e, 0x72, 0x3f, 0x31, 0xfc, 0x0b, 0x35,
	0x92, 0x71, 0xaa, 0x04, 0x90, 0x86, 0x71, 0x8f, 0x0a, 0xf7, 0x50, 0x3d, 0x99, 0x72, 0x01, 0xa4,
	0x69, 0x70, 0x2e, 0xf0, 0x5f, 0xf4, 0x23, 0x62, 0x1a, 0xb8, 0x08, 0x80, 0x4b, 0x91, 0x0f, 0x6e,
	0x99, 0xc1, 0x17, 0x1c, 0xff, 0x43, 0x3f, 0xcb, 0x2c, 0xff, 0x82, 0x63, 0xd2, 0x2e, 0x8d, 0x6c,
	0x9b, 0x85, 0x92, 0xc9, 0x28, 0x8a, 0x14, 0xd3, 0x9a, 0xa0, 0x7c, 0x9b, 0x12, 0xc2, 0xbf, 0x91,
	0x03, 0xb2, 0xf0, 0xdb, 0xc6, 0x3f, 0x83, 0xec, 0xfd, 0x7c, 0x25, 0xc3, 0x87, 0x09, 0xe3, 0xf1,
	0x12, 0x48, 0xc7, 0xcc, 0x29, 0x23, 0x4c, 0x51, 0x53, 0xf3, 0x58, 0x30, 0xa5, 0xc9, 0x37, 0xb7,
	0xea, 0x39, 0xe3, 0xda, 0xf6, 0xed, 0x8f, 0x35, 0x2b, 0x20, 0xf6, 0xd0, 0xf7, 0x05, 0x17, 0xc1,
	0x8a, 0x3f, 0xb1, 0xe8, 0x98, 0xd2, 0x35, 0x29, 0x5f, 0xf1, 0xf8, 0x66, 0xbb, 0xa7, 0xf6, 0x6e,
	0x4f, 0xed, 0xf7, 0x3d, 0xb5, 0x9f, 0x0f, 0xd4, 0xda, 0x1d, 0xa8, 0xf5, 0x7a, 0xa0, 0xd6, 0xfd,
	0x55, 0xcc, 0x61, 0x99, 0xce, 0x87, 0xa1, 0x4c, 0xfc, 0xec, 0x48, 0xfe, 0x17, 0x57, 0xe2, 0x9f,
	0x0e, 0x69, 0x73, 0x2e, 0xe1, 0x71, 0xcd, 0xf4, 0xbc, 0x61, 0xce, 0xe3, 0xfa, 0x23, 0x00, 0x00,
	0xff, 0xff, 0x94, 0x14, 0xe4, 0x97, 0x6c, 0x02, 0x00, 0x00,
}

func (m *Txout) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Txout) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Txout) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.FinalizedHeight != 0 {
		i = encodeVarintTxout(dAtA, i, uint64(m.FinalizedHeight))
		i--
		dAtA[i] = 0x70
	}
	if len(m.Signers) > 0 {
		for iNdEx := len(m.Signers) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Signers[iNdEx])
			copy(dAtA[i:], m.Signers[iNdEx])
			i = encodeVarintTxout(dAtA, i, uint64(len(m.Signers[iNdEx])))
			i--
			dAtA[i] = 0x6a
		}
	}
	if m.BlockHeight != 0 {
		i = encodeVarintTxout(dAtA, i, uint64(m.BlockHeight))
		i--
		dAtA[i] = 0x60
	}
	if len(m.ToAddress) > 0 {
		i -= len(m.ToAddress)
		copy(dAtA[i:], m.ToAddress)
		i = encodeVarintTxout(dAtA, i, uint64(len(m.ToAddress)))
		i--
		dAtA[i] = 0x5a
	}
	if len(m.FromAddress) > 0 {
		i -= len(m.FromAddress)
		copy(dAtA[i:], m.FromAddress)
		i = encodeVarintTxout(dAtA, i, uint64(len(m.FromAddress)))
		i--
		dAtA[i] = 0x52
	}
	if m.DestinationAmount != 0 {
		i = encodeVarintTxout(dAtA, i, uint64(m.DestinationAmount))
		i--
		dAtA[i] = 0x48
	}
	if len(m.DestinationAsset) > 0 {
		i -= len(m.DestinationAsset)
		copy(dAtA[i:], m.DestinationAsset)
		i = encodeVarintTxout(dAtA, i, uint64(len(m.DestinationAsset)))
		i--
		dAtA[i] = 0x42
	}
	if m.MMint != 0 {
		i = encodeVarintTxout(dAtA, i, uint64(m.MMint))
		i--
		dAtA[i] = 0x38
	}
	if m.MBurnt != 0 {
		i = encodeVarintTxout(dAtA, i, uint64(m.MBurnt))
		i--
		dAtA[i] = 0x30
	}
	if m.SourceAmount != 0 {
		i = encodeVarintTxout(dAtA, i, uint64(m.SourceAmount))
		i--
		dAtA[i] = 0x28
	}
	if len(m.SourceAsset) > 0 {
		i -= len(m.SourceAsset)
		copy(dAtA[i:], m.SourceAsset)
		i = encodeVarintTxout(dAtA, i, uint64(len(m.SourceAsset)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.TxinHash) > 0 {
		i -= len(m.TxinHash)
		copy(dAtA[i:], m.TxinHash)
		i = encodeVarintTxout(dAtA, i, uint64(len(m.TxinHash)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Id != 0 {
		i = encodeVarintTxout(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTxout(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTxout(dAtA []byte, offset int, v uint64) int {
	offset -= sovTxout(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Txout) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTxout(uint64(l))
	}
	if m.Id != 0 {
		n += 1 + sovTxout(uint64(m.Id))
	}
	l = len(m.TxinHash)
	if l > 0 {
		n += 1 + l + sovTxout(uint64(l))
	}
	l = len(m.SourceAsset)
	if l > 0 {
		n += 1 + l + sovTxout(uint64(l))
	}
	if m.SourceAmount != 0 {
		n += 1 + sovTxout(uint64(m.SourceAmount))
	}
	if m.MBurnt != 0 {
		n += 1 + sovTxout(uint64(m.MBurnt))
	}
	if m.MMint != 0 {
		n += 1 + sovTxout(uint64(m.MMint))
	}
	l = len(m.DestinationAsset)
	if l > 0 {
		n += 1 + l + sovTxout(uint64(l))
	}
	if m.DestinationAmount != 0 {
		n += 1 + sovTxout(uint64(m.DestinationAmount))
	}
	l = len(m.FromAddress)
	if l > 0 {
		n += 1 + l + sovTxout(uint64(l))
	}
	l = len(m.ToAddress)
	if l > 0 {
		n += 1 + l + sovTxout(uint64(l))
	}
	if m.BlockHeight != 0 {
		n += 1 + sovTxout(uint64(m.BlockHeight))
	}
	if len(m.Signers) > 0 {
		for _, s := range m.Signers {
			l = len(s)
			n += 1 + l + sovTxout(uint64(l))
		}
	}
	if m.FinalizedHeight != 0 {
		n += 1 + sovTxout(uint64(m.FinalizedHeight))
	}
	return n
}

func sovTxout(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTxout(x uint64) (n int) {
	return sovTxout(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Txout) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTxout
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
			return fmt.Errorf("proto: Txout: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Txout: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTxout
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
				return ErrInvalidLengthTxout
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTxout
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTxout
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TxinHash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTxout
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
				return ErrInvalidLengthTxout
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTxout
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TxinHash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SourceAsset", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTxout
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
				return ErrInvalidLengthTxout
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTxout
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SourceAsset = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SourceAmount", wireType)
			}
			m.SourceAmount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTxout
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SourceAmount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MBurnt", wireType)
			}
			m.MBurnt = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTxout
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MBurnt |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MMint", wireType)
			}
			m.MMint = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTxout
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MMint |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DestinationAsset", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTxout
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
				return ErrInvalidLengthTxout
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTxout
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DestinationAsset = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DestinationAmount", wireType)
			}
			m.DestinationAmount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTxout
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DestinationAmount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FromAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTxout
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
				return ErrInvalidLengthTxout
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTxout
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FromAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ToAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTxout
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
				return ErrInvalidLengthTxout
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTxout
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ToAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 12:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockHeight", wireType)
			}
			m.BlockHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTxout
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BlockHeight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 13:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signers", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTxout
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
				return ErrInvalidLengthTxout
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTxout
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signers = append(m.Signers, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 14:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FinalizedHeight", wireType)
			}
			m.FinalizedHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTxout
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.FinalizedHeight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTxout(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTxout
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
func skipTxout(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTxout
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
					return 0, ErrIntOverflowTxout
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
					return 0, ErrIntOverflowTxout
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
				return 0, ErrInvalidLengthTxout
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTxout
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTxout
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTxout        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTxout          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTxout = fmt.Errorf("proto: unexpected end of group")
)