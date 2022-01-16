// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cns/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
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

type MsgCreateGRecord struct {
	Creator string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Key     string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	Value   string `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *MsgCreateGRecord) Reset()         { *m = MsgCreateGRecord{} }
func (m *MsgCreateGRecord) String() string { return proto.CompactTextString(m) }
func (*MsgCreateGRecord) ProtoMessage()    {}
func (*MsgCreateGRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_6219fe2c381986e6, []int{0}
}
func (m *MsgCreateGRecord) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateGRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateGRecord.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateGRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateGRecord.Merge(m, src)
}
func (m *MsgCreateGRecord) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateGRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateGRecord.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateGRecord proto.InternalMessageInfo

func (m *MsgCreateGRecord) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgCreateGRecord) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *MsgCreateGRecord) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type MsgCreateGRecordResponse struct {
}

func (m *MsgCreateGRecordResponse) Reset()         { *m = MsgCreateGRecordResponse{} }
func (m *MsgCreateGRecordResponse) String() string { return proto.CompactTextString(m) }
func (*MsgCreateGRecordResponse) ProtoMessage()    {}
func (*MsgCreateGRecordResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6219fe2c381986e6, []int{1}
}
func (m *MsgCreateGRecordResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateGRecordResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateGRecordResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateGRecordResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateGRecordResponse.Merge(m, src)
}
func (m *MsgCreateGRecordResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateGRecordResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateGRecordResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateGRecordResponse proto.InternalMessageInfo

type MsgUpdateGRecord struct {
	Creator string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Key     string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	Value   string `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *MsgUpdateGRecord) Reset()         { *m = MsgUpdateGRecord{} }
func (m *MsgUpdateGRecord) String() string { return proto.CompactTextString(m) }
func (*MsgUpdateGRecord) ProtoMessage()    {}
func (*MsgUpdateGRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_6219fe2c381986e6, []int{2}
}
func (m *MsgUpdateGRecord) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUpdateGRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUpdateGRecord.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUpdateGRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUpdateGRecord.Merge(m, src)
}
func (m *MsgUpdateGRecord) XXX_Size() int {
	return m.Size()
}
func (m *MsgUpdateGRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUpdateGRecord.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUpdateGRecord proto.InternalMessageInfo

func (m *MsgUpdateGRecord) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgUpdateGRecord) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *MsgUpdateGRecord) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type MsgUpdateGRecordResponse struct {
}

func (m *MsgUpdateGRecordResponse) Reset()         { *m = MsgUpdateGRecordResponse{} }
func (m *MsgUpdateGRecordResponse) String() string { return proto.CompactTextString(m) }
func (*MsgUpdateGRecordResponse) ProtoMessage()    {}
func (*MsgUpdateGRecordResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6219fe2c381986e6, []int{3}
}
func (m *MsgUpdateGRecordResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUpdateGRecordResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUpdateGRecordResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUpdateGRecordResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUpdateGRecordResponse.Merge(m, src)
}
func (m *MsgUpdateGRecordResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgUpdateGRecordResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUpdateGRecordResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUpdateGRecordResponse proto.InternalMessageInfo

type MsgDeleteGRecord struct {
	Creator string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Key     string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
}

func (m *MsgDeleteGRecord) Reset()         { *m = MsgDeleteGRecord{} }
func (m *MsgDeleteGRecord) String() string { return proto.CompactTextString(m) }
func (*MsgDeleteGRecord) ProtoMessage()    {}
func (*MsgDeleteGRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_6219fe2c381986e6, []int{4}
}
func (m *MsgDeleteGRecord) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgDeleteGRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgDeleteGRecord.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgDeleteGRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgDeleteGRecord.Merge(m, src)
}
func (m *MsgDeleteGRecord) XXX_Size() int {
	return m.Size()
}
func (m *MsgDeleteGRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgDeleteGRecord.DiscardUnknown(m)
}

var xxx_messageInfo_MsgDeleteGRecord proto.InternalMessageInfo

func (m *MsgDeleteGRecord) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgDeleteGRecord) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type MsgDeleteGRecordResponse struct {
}

func (m *MsgDeleteGRecordResponse) Reset()         { *m = MsgDeleteGRecordResponse{} }
func (m *MsgDeleteGRecordResponse) String() string { return proto.CompactTextString(m) }
func (*MsgDeleteGRecordResponse) ProtoMessage()    {}
func (*MsgDeleteGRecordResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6219fe2c381986e6, []int{5}
}
func (m *MsgDeleteGRecordResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgDeleteGRecordResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgDeleteGRecordResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgDeleteGRecordResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgDeleteGRecordResponse.Merge(m, src)
}
func (m *MsgDeleteGRecordResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgDeleteGRecordResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgDeleteGRecordResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgDeleteGRecordResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgCreateGRecord)(nil), "jtremback.cns.cns.MsgCreateGRecord")
	proto.RegisterType((*MsgCreateGRecordResponse)(nil), "jtremback.cns.cns.MsgCreateGRecordResponse")
	proto.RegisterType((*MsgUpdateGRecord)(nil), "jtremback.cns.cns.MsgUpdateGRecord")
	proto.RegisterType((*MsgUpdateGRecordResponse)(nil), "jtremback.cns.cns.MsgUpdateGRecordResponse")
	proto.RegisterType((*MsgDeleteGRecord)(nil), "jtremback.cns.cns.MsgDeleteGRecord")
	proto.RegisterType((*MsgDeleteGRecordResponse)(nil), "jtremback.cns.cns.MsgDeleteGRecordResponse")
}

func init() { proto.RegisterFile("cns/tx.proto", fileDescriptor_6219fe2c381986e6) }

var fileDescriptor_6219fe2c381986e6 = []byte{
	// 289 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x49, 0xce, 0x2b, 0xd6,
	0x2f, 0xa9, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0xcc, 0x2a, 0x29, 0x4a, 0xcd, 0x4d,
	0x4a, 0x4c, 0xce, 0xd6, 0x4b, 0xce, 0x2b, 0x06, 0x61, 0x29, 0x21, 0x90, 0x82, 0xf4, 0xf8, 0xa2,
	0xd4, 0xe4, 0xfc, 0xa2, 0x14, 0x88, 0x32, 0xa5, 0x10, 0x2e, 0x01, 0xdf, 0xe2, 0x74, 0xe7, 0xa2,
	0xd4, 0xc4, 0x92, 0x54, 0xf7, 0x20, 0xb0, 0x8c, 0x90, 0x04, 0x17, 0x7b, 0x32, 0x48, 0x20, 0xbf,
	0x48, 0x82, 0x51, 0x81, 0x51, 0x83, 0x33, 0x08, 0xc6, 0x15, 0x12, 0xe0, 0x62, 0xce, 0x4e, 0xad,
	0x94, 0x60, 0x02, 0x8b, 0x82, 0x98, 0x42, 0x22, 0x5c, 0xac, 0x65, 0x89, 0x39, 0xa5, 0xa9, 0x12,
	0xcc, 0x60, 0x31, 0x08, 0x47, 0x49, 0x8a, 0x4b, 0x02, 0xdd, 0xd4, 0xa0, 0xd4, 0xe2, 0x82, 0xfc,
	0xbc, 0xe2, 0x54, 0xa8, 0x8d, 0xa1, 0x05, 0x29, 0x34, 0xb0, 0x11, 0xc5, 0x54, 0xb8, 0x8d, 0x76,
	0x60, 0x1b, 0x5d, 0x52, 0x73, 0x52, 0xc9, 0xb2, 0x11, 0x6a, 0x36, 0x8a, 0x7e, 0x98, 0xd9, 0x46,
	0xeb, 0x98, 0xb8, 0x98, 0x7d, 0x8b, 0xd3, 0x85, 0x12, 0xb9, 0x78, 0x51, 0x03, 0x51, 0x59, 0x0f,
	0x23, 0x02, 0xf4, 0xd0, 0xc3, 0x44, 0x4a, 0x9b, 0x08, 0x45, 0x30, 0xab, 0x40, 0x56, 0xa0, 0x86,
	0x1a, 0x0e, 0x2b, 0x50, 0x14, 0xe1, 0xb2, 0x02, 0x6b, 0x48, 0x81, 0xac, 0x40, 0x0d, 0x26, 0x1c,
	0x56, 0xa0, 0x28, 0xc2, 0x65, 0x05, 0xd6, 0x00, 0x73, 0xb2, 0x3b, 0xf1, 0x48, 0x8e, 0xf1, 0xc2,
	0x23, 0x39, 0xc6, 0x07, 0x8f, 0xe4, 0x18, 0x27, 0x3c, 0x96, 0x63, 0xb8, 0xf0, 0x58, 0x8e, 0xe1,
	0xc6, 0x63, 0x39, 0x86, 0x28, 0x95, 0xf4, 0xcc, 0x92, 0x8c, 0xd2, 0x24, 0xbd, 0xe4, 0xfc, 0x5c,
	0x7d, 0xb8, 0x81, 0xfa, 0xa0, 0x34, 0x5b, 0x01, 0x26, 0x4b, 0x2a, 0x0b, 0x52, 0x8b, 0x93, 0xd8,
	0xc0, 0xe9, 0xd6, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xeb, 0xb8, 0xd7, 0x77, 0xee, 0x02, 0x00,
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
	CreateGRecord(ctx context.Context, in *MsgCreateGRecord, opts ...grpc.CallOption) (*MsgCreateGRecordResponse, error)
	UpdateGRecord(ctx context.Context, in *MsgUpdateGRecord, opts ...grpc.CallOption) (*MsgUpdateGRecordResponse, error)
	DeleteGRecord(ctx context.Context, in *MsgDeleteGRecord, opts ...grpc.CallOption) (*MsgDeleteGRecordResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) CreateGRecord(ctx context.Context, in *MsgCreateGRecord, opts ...grpc.CallOption) (*MsgCreateGRecordResponse, error) {
	out := new(MsgCreateGRecordResponse)
	err := c.cc.Invoke(ctx, "/jtremback.cns.cns.Msg/CreateGRecord", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) UpdateGRecord(ctx context.Context, in *MsgUpdateGRecord, opts ...grpc.CallOption) (*MsgUpdateGRecordResponse, error) {
	out := new(MsgUpdateGRecordResponse)
	err := c.cc.Invoke(ctx, "/jtremback.cns.cns.Msg/UpdateGRecord", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) DeleteGRecord(ctx context.Context, in *MsgDeleteGRecord, opts ...grpc.CallOption) (*MsgDeleteGRecordResponse, error) {
	out := new(MsgDeleteGRecordResponse)
	err := c.cc.Invoke(ctx, "/jtremback.cns.cns.Msg/DeleteGRecord", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	CreateGRecord(context.Context, *MsgCreateGRecord) (*MsgCreateGRecordResponse, error)
	UpdateGRecord(context.Context, *MsgUpdateGRecord) (*MsgUpdateGRecordResponse, error)
	DeleteGRecord(context.Context, *MsgDeleteGRecord) (*MsgDeleteGRecordResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) CreateGRecord(ctx context.Context, req *MsgCreateGRecord) (*MsgCreateGRecordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateGRecord not implemented")
}
func (*UnimplementedMsgServer) UpdateGRecord(ctx context.Context, req *MsgUpdateGRecord) (*MsgUpdateGRecordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateGRecord not implemented")
}
func (*UnimplementedMsgServer) DeleteGRecord(ctx context.Context, req *MsgDeleteGRecord) (*MsgDeleteGRecordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteGRecord not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_CreateGRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreateGRecord)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreateGRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/jtremback.cns.cns.Msg/CreateGRecord",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreateGRecord(ctx, req.(*MsgCreateGRecord))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_UpdateGRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateGRecord)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateGRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/jtremback.cns.cns.Msg/UpdateGRecord",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateGRecord(ctx, req.(*MsgUpdateGRecord))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_DeleteGRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgDeleteGRecord)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).DeleteGRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/jtremback.cns.cns.Msg/DeleteGRecord",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).DeleteGRecord(ctx, req.(*MsgDeleteGRecord))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "jtremback.cns.cns.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateGRecord",
			Handler:    _Msg_CreateGRecord_Handler,
		},
		{
			MethodName: "UpdateGRecord",
			Handler:    _Msg_UpdateGRecord_Handler,
		},
		{
			MethodName: "DeleteGRecord",
			Handler:    _Msg_DeleteGRecord_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cns/tx.proto",
}

func (m *MsgCreateGRecord) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateGRecord) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateGRecord) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Value) > 0 {
		i -= len(m.Value)
		copy(dAtA[i:], m.Value)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Value)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Key) > 0 {
		i -= len(m.Key)
		copy(dAtA[i:], m.Key)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Key)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgCreateGRecordResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateGRecordResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateGRecordResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgUpdateGRecord) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUpdateGRecord) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUpdateGRecord) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Value) > 0 {
		i -= len(m.Value)
		copy(dAtA[i:], m.Value)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Value)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Key) > 0 {
		i -= len(m.Key)
		copy(dAtA[i:], m.Key)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Key)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgUpdateGRecordResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUpdateGRecordResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUpdateGRecordResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgDeleteGRecord) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgDeleteGRecord) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgDeleteGRecord) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Key) > 0 {
		i -= len(m.Key)
		copy(dAtA[i:], m.Key)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Key)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgDeleteGRecordResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgDeleteGRecordResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgDeleteGRecordResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
func (m *MsgCreateGRecord) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgCreateGRecordResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgUpdateGRecord) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgUpdateGRecordResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgDeleteGRecord) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgDeleteGRecordResponse) Size() (n int) {
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
func (m *MsgCreateGRecord) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgCreateGRecord: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateGRecord: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
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
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
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
			m.Key = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
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
			m.Value = string(dAtA[iNdEx:postIndex])
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
func (m *MsgCreateGRecordResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgCreateGRecordResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateGRecordResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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
func (m *MsgUpdateGRecord) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgUpdateGRecord: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUpdateGRecord: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
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
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
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
			m.Key = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
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
			m.Value = string(dAtA[iNdEx:postIndex])
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
func (m *MsgUpdateGRecordResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgUpdateGRecordResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUpdateGRecordResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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
func (m *MsgDeleteGRecord) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgDeleteGRecord: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgDeleteGRecord: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
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
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
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
			m.Key = string(dAtA[iNdEx:postIndex])
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
func (m *MsgDeleteGRecordResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgDeleteGRecordResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgDeleteGRecordResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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