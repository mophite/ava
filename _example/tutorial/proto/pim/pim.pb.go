// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proto/pim/pim.proto

package pim

import (
	fmt "fmt"
	ava "github.com/go-ava/ava"
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

type SayReq struct {
	Message string `protobuf:"bytes,1,opt,name=Message,proto3" json:"Message,omitempty"`
	Name    string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (m *SayReq) Reset()         { *m = SayReq{} }
func (m *SayReq) String() string { return proto.CompactTextString(m) }
func (*SayReq) ProtoMessage()    {}
func (*SayReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_4a457e6a329e2485, []int{0}
}
func (m *SayReq) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SayReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SayReq.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SayReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SayReq.Merge(m, src)
}
func (m *SayReq) XXX_Size() int {
	return m.Size()
}
func (m *SayReq) XXX_DiscardUnknown() {
	xxx_messageInfo_SayReq.DiscardUnknown(m)
}

var xxx_messageInfo_SayReq proto.InternalMessageInfo

func (m *SayReq) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *SayReq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type SayRsp struct {
	Message string `protobuf:"bytes,1,opt,name=Message,proto3" json:"Message,omitempty"`
}

func (m *SayRsp) Reset()         { *m = SayRsp{} }
func (m *SayRsp) String() string { return proto.CompactTextString(m) }
func (*SayRsp) ProtoMessage()    {}
func (*SayRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_4a457e6a329e2485, []int{1}
}
func (m *SayRsp) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SayRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SayRsp.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SayRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SayRsp.Merge(m, src)
}
func (m *SayRsp) XXX_Size() int {
	return m.Size()
}
func (m *SayRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_SayRsp.DiscardUnknown(m)
}

var xxx_messageInfo_SayRsp proto.InternalMessageInfo

func (m *SayRsp) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

// HandshakeReq to establish a connection
type HandshakeReq struct {
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (m *HandshakeReq) Reset()         { *m = HandshakeReq{} }
func (m *HandshakeReq) String() string { return proto.CompactTextString(m) }
func (*HandshakeReq) ProtoMessage()    {}
func (*HandshakeReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_4a457e6a329e2485, []int{2}
}
func (m *HandshakeReq) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *HandshakeReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_HandshakeReq.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *HandshakeReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HandshakeReq.Merge(m, src)
}
func (m *HandshakeReq) XXX_Size() int {
	return m.Size()
}
func (m *HandshakeReq) XXX_DiscardUnknown() {
	xxx_messageInfo_HandshakeReq.DiscardUnknown(m)
}

var xxx_messageInfo_HandshakeReq proto.InternalMessageInfo

func (m *HandshakeReq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// response content.
// BroadcastMessage usually use for broadcast
type BroadcastMessage struct {
	// you can add a business serial number for client
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (m *BroadcastMessage) Reset()         { *m = BroadcastMessage{} }
func (m *BroadcastMessage) String() string { return proto.CompactTextString(m) }
func (*BroadcastMessage) ProtoMessage()    {}
func (*BroadcastMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_4a457e6a329e2485, []int{3}
}
func (m *BroadcastMessage) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BroadcastMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BroadcastMessage.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BroadcastMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BroadcastMessage.Merge(m, src)
}
func (m *BroadcastMessage) XXX_Size() int {
	return m.Size()
}
func (m *BroadcastMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_BroadcastMessage.DiscardUnknown(m)
}

var xxx_messageInfo_BroadcastMessage proto.InternalMessageInfo

func (m *BroadcastMessage) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*SayReq)(nil), "pim.SayReq")
	proto.RegisterType((*SayRsp)(nil), "pim.SayRsp")
	proto.RegisterType((*HandshakeReq)(nil), "pim.HandshakeReq")
	proto.RegisterType((*BroadcastMessage)(nil), "pim.BroadcastMessage")
}

func init() { proto.RegisterFile("proto/pim/pim.proto", fileDescriptor_4a457e6a329e2485) }

var fileDescriptor_4a457e6a329e2485 = []byte{
	// 237 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2e, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x2f, 0xc8, 0xcc, 0x05, 0x61, 0x3d, 0x30, 0x4f, 0x88, 0xb9, 0x20, 0x33, 0x57, 0xc9,
	0x8c, 0x8b, 0x2d, 0x38, 0xb1, 0x32, 0x28, 0xb5, 0x50, 0x48, 0x82, 0x8b, 0xdd, 0x37, 0xb5, 0xb8,
	0x38, 0x31, 0x3d, 0x55, 0x82, 0x51, 0x81, 0x51, 0x83, 0x33, 0x08, 0xc6, 0x15, 0x12, 0xe2, 0x62,
	0xc9, 0x4b, 0xcc, 0x4d, 0x95, 0x60, 0x02, 0x0b, 0x83, 0xd9, 0x4a, 0x4a, 0x10, 0x7d, 0xc5, 0x05,
	0xb8, 0xf5, 0x29, 0x29, 0x71, 0xf1, 0x78, 0x24, 0xe6, 0xa5, 0x14, 0x67, 0x24, 0x66, 0xa7, 0x82,
	0x6c, 0x80, 0x99, 0xc3, 0x88, 0x64, 0x8e, 0x0e, 0x97, 0x80, 0x53, 0x51, 0x7e, 0x62, 0x4a, 0x72,
	0x62, 0x71, 0x09, 0xcc, 0x3e, 0x09, 0x2e, 0xf6, 0x5c, 0xa8, 0x89, 0x10, 0x2b, 0x61, 0x5c, 0xa3,
	0x05, 0x8c, 0x5c, 0x4c, 0x9e, 0xb9, 0x42, 0x4a, 0x5c, 0x2c, 0xc1, 0xa9, 0x79, 0x29, 0x42, 0xdc,
	0x7a, 0x20, 0xdf, 0x40, 0xdc, 0x2f, 0x85, 0xe0, 0x14, 0x17, 0x28, 0x31, 0x08, 0x59, 0x72, 0x71,
	0xc2, 0x2d, 0x17, 0x12, 0x04, 0xcb, 0x21, 0x3b, 0x46, 0x4a, 0x14, 0x2c, 0x84, 0x6e, 0xb7, 0x12,
	0x83, 0x01, 0xa3, 0x90, 0x35, 0x17, 0x27, 0x5c, 0x5c, 0x08, 0xbb, 0x3a, 0x9c, 0xda, 0x9d, 0x24,
	0x4e, 0x3c, 0x92, 0x63, 0xbc, 0xf0, 0x48, 0x8e, 0xf1, 0xc1, 0x23, 0x39, 0xc6, 0x09, 0x8f, 0xe5,
	0x18, 0x2e, 0x3c, 0x96, 0x63, 0xb8, 0xf1, 0x58, 0x8e, 0x21, 0x89, 0x0d, 0x1c, 0xec, 0xc6, 0x80,
	0x00, 0x00, 0x00, 0xff, 0xff, 0xca, 0x42, 0xc5, 0x53, 0x8d, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ ava.Service

// This is a compile-time assertion to ensure that this generated file
// is compatible with the ava package it is being compiled against.
const _ = ava.SupportPackageIsVersion1

type ImClient interface {
	Send(c *ava.Context, req *SayReq, opts ...ava.InvokeOptions) (*SayRsp, error)
	Handshake(c *ava.Context, req *HandshakeReq, opts ...ava.InvokeOptions) chan *BroadcastMessage
	Broadcast(c *ava.Context, req *BroadcastMessage, opts ...ava.InvokeOptions) (*BroadcastMessage, error)
}

type imClient struct {
	c *ava.Client
}

func NewImClient(c *ava.Client) ImClient {
	return &imClient{c}
}

func (cc *imClient) Send(c *ava.Context, req *SayReq, opts ...ava.InvokeOptions) (*SayRsp, error) {
	rsp := &SayRsp{}
	err := cc.c.InvokeRR(c, "/im/send", req, rsp, opts...)
	return rsp, err
}

func (cc *imClient) Handshake(c *ava.Context, req *HandshakeReq, opts ...ava.InvokeOptions) chan *BroadcastMessage {
	data := cc.c.InvokeRS(c, "/im/handshake", req, opts...)
	if data == nil {
		return nil
	}

	var rsp = make(chan *BroadcastMessage, cap(data))
	go func() {
		for b := range data {
			v := &BroadcastMessage{}
			err := c.Codec().Decode(b, v)
			if err != nil {
				c.Errorf("client decode pakcet err=%v |method=%s |data=%s", err, c.Metadata.Method(), req.String())
				continue
			}
			rsp <- v
		}
		close(rsp)
	}()
	return rsp
}

func (cc *imClient) Broadcast(c *ava.Context, req *BroadcastMessage, opts ...ava.InvokeOptions) (*BroadcastMessage, error) {
	rsp := &BroadcastMessage{}
	err := cc.c.InvokeRR(c, "/im/broadcast", req, rsp, opts...)
	return rsp, err
}

// ImServer is the server API for Im ava.
type ImServer interface {
	Send(c *ava.Context, req *SayReq, rsp *SayRsp)
	Handshake(c *ava.Context, req *HandshakeReq, exit chan struct{}) chan *BroadcastMessage
	Broadcast(c *ava.Context, req *BroadcastMessage, rsp *BroadcastMessage)
}

func RegisterImServer(s *ava.Server, h ImServer) {
	var r = &imHandler{h: h, s: s}
	s.RegisterHandler("/"+s.Name()+"/im/send", r.Send)
	s.RegisterStreamHandler("/"+s.Name()+"/im/handshake", r.Handshake)
	s.RegisterHandler("/"+s.Name()+"/im/broadcast", r.Broadcast)
}

type imHandler struct {
	h ImServer
	s *ava.Server
}

func (r *imHandler) Send(c *ava.Context, req *ava.Packet, interrupt ava.Interceptor) (rsp proto.Message, err error) {
	var in SayReq
	err = c.Codec().Decode(req.Bytes(), &in)
	if err != nil {
		c.Errorf("server decode packet err=%v |method=%s |data=%s", err, c.Metadata.Method(), req.String())
		return nil, err
	}
	var out = SayRsp{}
	if interrupt == nil {
		r.h.Send(c, &in, &out)
		return &out, err
	}
	f := func(c *ava.Context, req proto.Message) proto.Message {
		r.h.Send(c, req.(*SayReq), &out)
		return &out
	}
	return interrupt(c, &in, f)
}

func (r *imHandler) Handshake(c *ava.Context, req *ava.Packet, exit chan struct{}) chan proto.Message {
	var in HandshakeReq
	err := c.Codec().Decode(req.Bytes(), &in)
	if err != nil {
		c.Errorf("server decode packet err=%v |method=%s |data=%s", err, c.Metadata.Method(), req.String())
		return nil
	}

	out := r.h.Handshake(c, &in, exit)
	if out == nil {
		return nil
	}

	var rsp = make(chan proto.Message, cap(out))

	go func() {
		for d := range out {
			rsp <- d
		}
		close(rsp)
	}()
	return rsp
}

func (r *imHandler) Broadcast(c *ava.Context, req *ava.Packet, interrupt ava.Interceptor) (rsp proto.Message, err error) {
	var in BroadcastMessage
	err = c.Codec().Decode(req.Bytes(), &in)
	if err != nil {
		c.Errorf("server decode packet err=%v |method=%s |data=%s", err, c.Metadata.Method(), req.String())
		return nil, err
	}
	var out = BroadcastMessage{}
	if interrupt == nil {
		r.h.Broadcast(c, &in, &out)
		return &out, err
	}
	f := func(c *ava.Context, req proto.Message) proto.Message {
		r.h.Broadcast(c, req.(*BroadcastMessage), &out)
		return &out
	}
	return interrupt(c, &in, f)
}

func (m *SayReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SayReq) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SayReq) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintPim(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Message) > 0 {
		i -= len(m.Message)
		copy(dAtA[i:], m.Message)
		i = encodeVarintPim(dAtA, i, uint64(len(m.Message)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *SayRsp) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SayRsp) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SayRsp) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Message) > 0 {
		i -= len(m.Message)
		copy(dAtA[i:], m.Message)
		i = encodeVarintPim(dAtA, i, uint64(len(m.Message)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *HandshakeReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HandshakeReq) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *HandshakeReq) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintPim(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *BroadcastMessage) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BroadcastMessage) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BroadcastMessage) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Message) > 0 {
		i -= len(m.Message)
		copy(dAtA[i:], m.Message)
		i = encodeVarintPim(dAtA, i, uint64(len(m.Message)))
		i--
		dAtA[i] = 0x12
	}
	return len(dAtA) - i, nil
}

func encodeVarintPim(dAtA []byte, offset int, v uint64) int {
	offset -= sovPim(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *SayReq) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Message)
	if l > 0 {
		n += 1 + l + sovPim(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovPim(uint64(l))
	}
	return n
}

func (m *SayRsp) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Message)
	if l > 0 {
		n += 1 + l + sovPim(uint64(l))
	}
	return n
}

func (m *HandshakeReq) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovPim(uint64(l))
	}
	return n
}

func (m *BroadcastMessage) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Message)
	if l > 0 {
		n += 1 + l + sovPim(uint64(l))
	}
	return n
}

func sovPim(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPim(x uint64) (n int) {
	return sovPim(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *SayReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPim
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
			return fmt.Errorf("proto: SayReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SayReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Message", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPim
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
				return ErrInvalidLengthPim
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPim
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Message = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPim
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
				return ErrInvalidLengthPim
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPim
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPim(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPim
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
func (m *SayRsp) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPim
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
			return fmt.Errorf("proto: SayRsp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SayRsp: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Message", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPim
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
				return ErrInvalidLengthPim
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPim
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Message = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPim(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPim
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
func (m *HandshakeReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPim
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
			return fmt.Errorf("proto: HandshakeReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HandshakeReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPim
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
				return ErrInvalidLengthPim
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPim
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPim(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPim
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
func (m *BroadcastMessage) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPim
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
			return fmt.Errorf("proto: BroadcastMessage: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BroadcastMessage: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Message", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPim
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
				return ErrInvalidLengthPim
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPim
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Message = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPim(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPim
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
func skipPim(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPim
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
					return 0, ErrIntOverflowPim
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
					return 0, ErrIntOverflowPim
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
				return 0, ErrInvalidLengthPim
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPim
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPim
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPim        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPim          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPim = fmt.Errorf("proto: unexpected end of group")
)