// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: testpb/test.proto

package testpb

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types/tx/amino"
	_ "github.com/cosmos/gogoproto/gogoproto"
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

type Streng struct {
	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *Streng) Reset()         { *m = Streng{} }
func (m *Streng) String() string { return proto.CompactTextString(m) }
func (*Streng) ProtoMessage()    {}
func (*Streng) Descriptor() ([]byte, []int) {
	return fileDescriptor_41c67e33ca9d1f26, []int{0}
}
func (m *Streng) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Streng) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Streng.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Streng) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Streng.Merge(m, src)
}
func (m *Streng) XXX_Size() int {
	return m.Size()
}
func (m *Streng) XXX_DiscardUnknown() {
	xxx_messageInfo_Streng.DiscardUnknown(m)
}

var xxx_messageInfo_Streng proto.InternalMessageInfo

func (m *Streng) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type TestRepeatedFields struct {
	NullableOmitempty        []*Streng `protobuf:"bytes,1,rep,name=nullable_omitempty,json=nullableOmitempty,proto3" json:"nullable_omitempty,omitempty"`
	NullableDontOmitempty    []*Streng `protobuf:"bytes,2,rep,name=nullable_dont_omitempty,json=nullableDontOmitempty,proto3" json:"nullable_dont_omitempty,omitempty"`
	NonNullableOmitempty     []Streng  `protobuf:"bytes,3,rep,name=non_nullable_omitempty,json=nonNullableOmitempty,proto3" json:"non_nullable_omitempty"`
	NonNullableDontOmitempty []Streng  `protobuf:"bytes,4,rep,name=non_nullable_dont_omitempty,json=nonNullableDontOmitempty,proto3" json:"non_nullable_dont_omitempty"`
}

func (m *TestRepeatedFields) Reset()         { *m = TestRepeatedFields{} }
func (m *TestRepeatedFields) String() string { return proto.CompactTextString(m) }
func (*TestRepeatedFields) ProtoMessage()    {}
func (*TestRepeatedFields) Descriptor() ([]byte, []int) {
	return fileDescriptor_41c67e33ca9d1f26, []int{1}
}
func (m *TestRepeatedFields) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TestRepeatedFields) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TestRepeatedFields.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TestRepeatedFields) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestRepeatedFields.Merge(m, src)
}
func (m *TestRepeatedFields) XXX_Size() int {
	return m.Size()
}
func (m *TestRepeatedFields) XXX_DiscardUnknown() {
	xxx_messageInfo_TestRepeatedFields.DiscardUnknown(m)
}

var xxx_messageInfo_TestRepeatedFields proto.InternalMessageInfo

func (m *TestRepeatedFields) GetNullableOmitempty() []*Streng {
	if m != nil {
		return m.NullableOmitempty
	}
	return nil
}

func (m *TestRepeatedFields) GetNullableDontOmitempty() []*Streng {
	if m != nil {
		return m.NullableDontOmitempty
	}
	return nil
}

func (m *TestRepeatedFields) GetNonNullableOmitempty() []Streng {
	if m != nil {
		return m.NonNullableOmitempty
	}
	return nil
}

func (m *TestRepeatedFields) GetNonNullableDontOmitempty() []Streng {
	if m != nil {
		return m.NonNullableDontOmitempty
	}
	return nil
}

func init() {
	proto.RegisterType((*Streng)(nil), "testpb.streng")
	proto.RegisterType((*TestRepeatedFields)(nil), "testpb.TestRepeatedFields")
}

func init() { proto.RegisterFile("testpb/test.proto", fileDescriptor_41c67e33ca9d1f26) }

var fileDescriptor_41c67e33ca9d1f26 = []byte{
	// 277 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2c, 0x49, 0x2d, 0x2e,
	0x29, 0x48, 0xd2, 0x07, 0x51, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0x6c, 0x10, 0x21, 0x29,
	0x91, 0xf4, 0xfc, 0xf4, 0x7c, 0xb0, 0x90, 0x3e, 0x88, 0x05, 0x91, 0x95, 0x12, 0x4c, 0xcc, 0xcd,
	0xcc, 0xcb, 0xd7, 0x07, 0x93, 0x10, 0x21, 0x25, 0x39, 0x2e, 0xb6, 0xe2, 0x92, 0xa2, 0xd4, 0xbc,
	0x74, 0x21, 0x11, 0x2e, 0xd6, 0xb2, 0xc4, 0x9c, 0xd2, 0x54, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce,
	0x20, 0x08, 0x47, 0xe9, 0x04, 0x13, 0x97, 0x50, 0x48, 0x6a, 0x71, 0x49, 0x50, 0x6a, 0x41, 0x6a,
	0x62, 0x49, 0x6a, 0x8a, 0x5b, 0x66, 0x6a, 0x4e, 0x4a, 0xb1, 0x90, 0x2d, 0x97, 0x50, 0x5e, 0x69,
	0x4e, 0x4e, 0x62, 0x52, 0x4e, 0x6a, 0x7c, 0x7e, 0x6e, 0x66, 0x49, 0x6a, 0x6e, 0x41, 0x49, 0xa5,
	0x04, 0xa3, 0x02, 0xb3, 0x06, 0xb7, 0x11, 0x9f, 0x1e, 0xc4, 0x11, 0x7a, 0x10, 0x83, 0x83, 0x04,
	0x61, 0x2a, 0xfd, 0x61, 0x0a, 0x85, 0x7c, 0xb9, 0xc4, 0xe1, 0xda, 0x53, 0xf2, 0xf3, 0x4a, 0x90,
	0xcc, 0x60, 0xc2, 0x66, 0x86, 0x13, 0xeb, 0x8a, 0xe7, 0x1b, 0xb4, 0x18, 0x83, 0x44, 0x61, 0xba,
	0x5c, 0xf2, 0xf3, 0x4a, 0x10, 0xc6, 0x79, 0x71, 0x89, 0xe5, 0xe5, 0xe7, 0xc5, 0x63, 0x71, 0x11,
	0x33, 0x56, 0xd3, 0x58, 0x4e, 0xdc, 0x93, 0x67, 0x08, 0x12, 0xc9, 0xcb, 0xcf, 0xf3, 0xc3, 0x70,
	0x5a, 0x04, 0x97, 0x34, 0x8a, 0x59, 0x68, 0xce, 0x63, 0xc1, 0x6a, 0x20, 0x27, 0xc8, 0x40, 0x88,
	0x13, 0x25, 0x90, 0x4c, 0x45, 0x71, 0xa5, 0x93, 0xc4, 0x89, 0x47, 0x72, 0x8c, 0x17, 0x1e, 0xc9,
	0x31, 0x3e, 0x78, 0x24, 0xc7, 0x38, 0xe1, 0xb1, 0x1c, 0xc3, 0x85, 0xc7, 0x72, 0x0c, 0x37, 0x1e,
	0xcb, 0x31, 0x24, 0xb1, 0x81, 0xe3, 0xc2, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xa3, 0x6f, 0x4e,
	0x5e, 0xd1, 0x01, 0x00, 0x00,
}

func (m *Streng) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Streng) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Streng) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Value) > 0 {
		i -= len(m.Value)
		copy(dAtA[i:], m.Value)
		i = encodeVarintTest(dAtA, i, uint64(len(m.Value)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *TestRepeatedFields) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TestRepeatedFields) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TestRepeatedFields) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.NonNullableDontOmitempty) > 0 {
		for iNdEx := len(m.NonNullableDontOmitempty) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.NonNullableDontOmitempty[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTest(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.NonNullableOmitempty) > 0 {
		for iNdEx := len(m.NonNullableOmitempty) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.NonNullableOmitempty[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTest(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.NullableDontOmitempty) > 0 {
		for iNdEx := len(m.NullableDontOmitempty) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.NullableDontOmitempty[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTest(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.NullableOmitempty) > 0 {
		for iNdEx := len(m.NullableOmitempty) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.NullableOmitempty[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTest(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintTest(dAtA []byte, offset int, v uint64) int {
	offset -= sovTest(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Streng) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovTest(uint64(l))
	}
	return n
}

func (m *TestRepeatedFields) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.NullableOmitempty) > 0 {
		for _, e := range m.NullableOmitempty {
			l = e.Size()
			n += 1 + l + sovTest(uint64(l))
		}
	}
	if len(m.NullableDontOmitempty) > 0 {
		for _, e := range m.NullableDontOmitempty {
			l = e.Size()
			n += 1 + l + sovTest(uint64(l))
		}
	}
	if len(m.NonNullableOmitempty) > 0 {
		for _, e := range m.NonNullableOmitempty {
			l = e.Size()
			n += 1 + l + sovTest(uint64(l))
		}
	}
	if len(m.NonNullableDontOmitempty) > 0 {
		for _, e := range m.NonNullableDontOmitempty {
			l = e.Size()
			n += 1 + l + sovTest(uint64(l))
		}
	}
	return n
}

func sovTest(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTest(x uint64) (n int) {
	return sovTest(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Streng) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTest
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
			return fmt.Errorf("proto: streng: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: streng: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTest
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
				return ErrInvalidLengthTest
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTest(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTest
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
func (m *TestRepeatedFields) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTest
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
			return fmt.Errorf("proto: TestRepeatedFields: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TestRepeatedFields: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NullableOmitempty", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTest
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
				return ErrInvalidLengthTest
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NullableOmitempty = append(m.NullableOmitempty, &Streng{})
			if err := m.NullableOmitempty[len(m.NullableOmitempty)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NullableDontOmitempty", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTest
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
				return ErrInvalidLengthTest
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NullableDontOmitempty = append(m.NullableDontOmitempty, &Streng{})
			if err := m.NullableDontOmitempty[len(m.NullableDontOmitempty)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NonNullableOmitempty", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTest
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
				return ErrInvalidLengthTest
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NonNullableOmitempty = append(m.NonNullableOmitempty, Streng{})
			if err := m.NonNullableOmitempty[len(m.NonNullableOmitempty)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NonNullableDontOmitempty", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTest
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
				return ErrInvalidLengthTest
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NonNullableDontOmitempty = append(m.NonNullableDontOmitempty, Streng{})
			if err := m.NonNullableDontOmitempty[len(m.NonNullableDontOmitempty)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTest(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTest
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
func skipTest(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTest
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
					return 0, ErrIntOverflowTest
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
					return 0, ErrIntOverflowTest
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
				return 0, ErrInvalidLengthTest
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTest
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTest
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTest        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTest          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTest = fmt.Errorf("proto: unexpected end of group")
)
