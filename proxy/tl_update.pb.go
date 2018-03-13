// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tl_update.proto

/*
Package proxy is a generated protocol buffer package.

It is generated from these files:
	tl_update.proto

It has these top-level messages:
	Update
	ListenRequest
*/
package proxy

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import core "github.com/cjongseok/mtproto/core"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Update struct {
	// Types that are valid to be assigned to Value:
	//	*Update_Updates
	//	*Update_UpdatesDifference
	Value isUpdate_Value `protobuf_oneof:"Value"`
}

func (m *Update) Reset()                    { *m = Update{} }
func (m *Update) String() string            { return proto.CompactTextString(m) }
func (*Update) ProtoMessage()               {}
func (*Update) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type isUpdate_Value interface {
	isUpdate_Value()
}

type Update_Updates struct {
	Updates *core.PredUpdates `protobuf:"bytes,1,opt,name=updates,oneof"`
}
type Update_UpdatesDifference struct {
	UpdatesDifference *core.PredUpdatesDifference `protobuf:"bytes,2,opt,name=updatesDifference,oneof"`
}

func (*Update_Updates) isUpdate_Value()           {}
func (*Update_UpdatesDifference) isUpdate_Value() {}

func (m *Update) GetValue() isUpdate_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *Update) GetUpdates() *core.PredUpdates {
	if x, ok := m.GetValue().(*Update_Updates); ok {
		return x.Updates
	}
	return nil
}

func (m *Update) GetUpdatesDifference() *core.PredUpdatesDifference {
	if x, ok := m.GetValue().(*Update_UpdatesDifference); ok {
		return x.UpdatesDifference
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Update) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Update_OneofMarshaler, _Update_OneofUnmarshaler, _Update_OneofSizer, []interface{}{
		(*Update_Updates)(nil),
		(*Update_UpdatesDifference)(nil),
	}
}

func _Update_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Update)
	// Value
	switch x := m.Value.(type) {
	case *Update_Updates:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Updates); err != nil {
			return err
		}
	case *Update_UpdatesDifference:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.UpdatesDifference); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Update.Value has unexpected type %T", x)
	}
	return nil
}

func _Update_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Update)
	switch tag {
	case 1: // Value.updates
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(core.PredUpdates)
		err := b.DecodeMessage(msg)
		m.Value = &Update_Updates{msg}
		return true, err
	case 2: // Value.updatesDifference
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(core.PredUpdatesDifference)
		err := b.DecodeMessage(msg)
		m.Value = &Update_UpdatesDifference{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Update_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Update)
	// Value
	switch x := m.Value.(type) {
	case *Update_Updates:
		s := proto.Size(x.Updates)
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Update_UpdatesDifference:
		s := proto.Size(x.UpdatesDifference)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type ListenRequest struct {
}

func (m *ListenRequest) Reset()                    { *m = ListenRequest{} }
func (m *ListenRequest) String() string            { return proto.CompactTextString(m) }
func (*ListenRequest) ProtoMessage()               {}
func (*ListenRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func init() {
	proto.RegisterType((*Update)(nil), "proxy.Update")
	proto.RegisterType((*ListenRequest)(nil), "proxy.ListenRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for UpdateStreamer service

type UpdateStreamerClient interface {
	ListenOnUpdates(ctx context.Context, in *ListenRequest, opts ...grpc.CallOption) (UpdateStreamer_ListenOnUpdatesClient, error)
}

type updateStreamerClient struct {
	cc *grpc.ClientConn
}

func NewUpdateStreamerClient(cc *grpc.ClientConn) UpdateStreamerClient {
	return &updateStreamerClient{cc}
}

func (c *updateStreamerClient) ListenOnUpdates(ctx context.Context, in *ListenRequest, opts ...grpc.CallOption) (UpdateStreamer_ListenOnUpdatesClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_UpdateStreamer_serviceDesc.Streams[0], c.cc, "/proxy.UpdateStreamer/ListenOnUpdates", opts...)
	if err != nil {
		return nil, err
	}
	x := &updateStreamerListenOnUpdatesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type UpdateStreamer_ListenOnUpdatesClient interface {
	Recv() (*Update, error)
	grpc.ClientStream
}

type updateStreamerListenOnUpdatesClient struct {
	grpc.ClientStream
}

func (x *updateStreamerListenOnUpdatesClient) Recv() (*Update, error) {
	m := new(Update)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for UpdateStreamer service

type UpdateStreamerServer interface {
	ListenOnUpdates(*ListenRequest, UpdateStreamer_ListenOnUpdatesServer) error
}

func RegisterUpdateStreamerServer(s *grpc.Server, srv UpdateStreamerServer) {
	s.RegisterService(&_UpdateStreamer_serviceDesc, srv)
}

func _UpdateStreamer_ListenOnUpdates_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListenRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(UpdateStreamerServer).ListenOnUpdates(m, &updateStreamerListenOnUpdatesServer{stream})
}

type UpdateStreamer_ListenOnUpdatesServer interface {
	Send(*Update) error
	grpc.ServerStream
}

type updateStreamerListenOnUpdatesServer struct {
	grpc.ServerStream
}

func (x *updateStreamerListenOnUpdatesServer) Send(m *Update) error {
	return x.ServerStream.SendMsg(m)
}

var _UpdateStreamer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proxy.UpdateStreamer",
	HandlerType: (*UpdateStreamerServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListenOnUpdates",
			Handler:       _UpdateStreamer_ListenOnUpdates_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "tl_update.proto",
}

func init() { proto.RegisterFile("tl_update.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 232 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x8f, 0xc1, 0x4b, 0xc3, 0x30,
	0x14, 0x87, 0x1b, 0x61, 0x1b, 0x3c, 0x99, 0x65, 0xc1, 0xc3, 0x98, 0x17, 0xe9, 0xc9, 0x8b, 0xc9,
	0x98, 0x37, 0x8f, 0xe2, 0x61, 0xe0, 0x40, 0xa9, 0xe8, 0x55, 0xba, 0xec, 0x6d, 0x56, 0xdb, 0x24,
	0x26, 0x2f, 0x60, 0xff, 0x08, 0xff, 0x67, 0x31, 0xa9, 0x88, 0xf4, 0xf6, 0x78, 0x7c, 0xdf, 0x07,
	0x3f, 0xc8, 0xa9, 0x79, 0x09, 0x76, 0x57, 0x11, 0x0a, 0xeb, 0x0c, 0x19, 0x3e, 0xb2, 0xce, 0x7c,
	0x76, 0x8b, 0xe5, 0xa1, 0xa6, 0xd7, 0xb0, 0x15, 0xca, 0xb4, 0x52, 0xbd, 0x19, 0x7d, 0xf0, 0x68,
	0xde, 0x65, 0x4b, 0x11, 0x92, 0xca, 0x38, 0x94, 0xd4, 0x59, 0xf4, 0x82, 0x9a, 0x24, 0x16, 0x5f,
	0x0c, 0xc6, 0x4f, 0xb1, 0xc4, 0x2f, 0x61, 0x92, 0x9a, 0x7e, 0xce, 0xce, 0xd9, 0xc5, 0xf1, 0x6a,
	0x26, 0x7e, 0x0c, 0xf1, 0xe0, 0x70, 0x97, 0x10, 0xbf, 0xce, 0xca, 0x5f, 0x86, 0xdf, 0xc1, 0xac,
	0x3f, 0x6f, 0xeb, 0xfd, 0x1e, 0x1d, 0x6a, 0x85, 0xf3, 0xa3, 0x28, 0x9e, 0x0d, 0xc4, 0x3f, 0x64,
	0x9d, 0x95, 0x43, 0xef, 0x66, 0x02, 0xa3, 0xe7, 0xaa, 0x09, 0x58, 0xe4, 0x30, 0xdd, 0xd4, 0x9e,
	0x50, 0x97, 0xf8, 0x11, 0xd0, 0xd3, 0x6a, 0x03, 0x27, 0xa9, 0xf1, 0x48, 0x0e, 0xab, 0x16, 0x1d,
	0xbf, 0x86, 0x3c, 0x21, 0xf7, 0xba, 0xaf, 0xf3, 0x53, 0x11, 0xf7, 0x8b, 0x7f, 0xea, 0x62, 0xda,
	0x7f, 0x13, 0x55, 0x64, 0x4b, 0xb6, 0x1d, 0xc7, 0xd5, 0x57, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff,
	0x15, 0xd3, 0xea, 0xe3, 0x41, 0x01, 0x00, 0x00,
}