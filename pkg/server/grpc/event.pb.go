// Code generated by protoc-gen-go. DO NOT EDIT.
// source: event.proto

package grpc

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"

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

// Ignoring public import of Error from common.proto

// Ignoring public import of TimestampRange from common.proto

type EventEntry struct {
	Id                   uint32               `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	NamespaceId          string               `protobuf:"bytes,2,opt,name=namespace_id,json=namespaceId,proto3" json:"namespace_id,omitempty"`
	LinkId               string               `protobuf:"bytes,3,opt,name=link_id,json=linkId,proto3" json:"link_id,omitempty"`
	AliasId              string               `protobuf:"bytes,4,opt,name=alias_id,json=aliasId,proto3" json:"alias_id,omitempty"`
	TargetId             string               `protobuf:"bytes,5,opt,name=target_id,json=targetId,proto3" json:"target_id,omitempty"`
	Identifier           string               `protobuf:"bytes,6,opt,name=identifier,proto3" json:"identifier,omitempty"`
	Context              []byte               `protobuf:"bytes,7,opt,name=context,proto3" json:"context,omitempty"`
	Code                 int32                `protobuf:"varint,8,opt,name=code,proto3" json:"code,omitempty"`
	Url                  string               `protobuf:"bytes,9,opt,name=url,proto3" json:"url,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,10,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *EventEntry) Reset()         { *m = EventEntry{} }
func (m *EventEntry) String() string { return proto.CompactTextString(m) }
func (*EventEntry) ProtoMessage()    {}
func (*EventEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d17a9d3f0ddf27e, []int{0}
}
func (m *EventEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventEntry.Unmarshal(m, b)
}
func (m *EventEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventEntry.Marshal(b, m, deterministic)
}
func (m *EventEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventEntry.Merge(m, src)
}
func (m *EventEntry) XXX_Size() int {
	return xxx_messageInfo_EventEntry.Size(m)
}
func (m *EventEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_EventEntry.DiscardUnknown(m)
}

var xxx_messageInfo_EventEntry proto.InternalMessageInfo

func (m *EventEntry) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *EventEntry) GetNamespaceId() string {
	if m != nil {
		return m.NamespaceId
	}
	return ""
}

func (m *EventEntry) GetLinkId() string {
	if m != nil {
		return m.LinkId
	}
	return ""
}

func (m *EventEntry) GetAliasId() string {
	if m != nil {
		return m.AliasId
	}
	return ""
}

func (m *EventEntry) GetTargetId() string {
	if m != nil {
		return m.TargetId
	}
	return ""
}

func (m *EventEntry) GetIdentifier() string {
	if m != nil {
		return m.Identifier
	}
	return ""
}

func (m *EventEntry) GetContext() []byte {
	if m != nil {
		return m.Context
	}
	return nil
}

func (m *EventEntry) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *EventEntry) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *EventEntry) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

type EventFilter struct {
	NamespaceId          string          `protobuf:"bytes,1,opt,name=namespace_id,json=namespaceId,proto3" json:"namespace_id,omitempty"`
	LinkId               string          `protobuf:"bytes,2,opt,name=link_id,json=linkId,proto3" json:"link_id,omitempty"`
	AliasId              string          `protobuf:"bytes,3,opt,name=alias_id,json=aliasId,proto3" json:"alias_id,omitempty"`
	TargetId             string          `protobuf:"bytes,4,opt,name=target_id,json=targetId,proto3" json:"target_id,omitempty"`
	Identifier           string          `protobuf:"bytes,5,opt,name=identifier,proto3" json:"identifier,omitempty"`
	Code                 int32           `protobuf:"varint,6,opt,name=code,proto3" json:"code,omitempty"`
	Url                  string          `protobuf:"bytes,7,opt,name=url,proto3" json:"url,omitempty"`
	CreatedAt            *TimestampRange `protobuf:"bytes,8,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	Limit                uint32          `protobuf:"varint,9,opt,name=limit,proto3" json:"limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *EventFilter) Reset()         { *m = EventFilter{} }
func (m *EventFilter) String() string { return proto.CompactTextString(m) }
func (*EventFilter) ProtoMessage()    {}
func (*EventFilter) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d17a9d3f0ddf27e, []int{1}
}
func (m *EventFilter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventFilter.Unmarshal(m, b)
}
func (m *EventFilter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventFilter.Marshal(b, m, deterministic)
}
func (m *EventFilter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventFilter.Merge(m, src)
}
func (m *EventFilter) XXX_Size() int {
	return xxx_messageInfo_EventFilter.Size(m)
}
func (m *EventFilter) XXX_DiscardUnknown() {
	xxx_messageInfo_EventFilter.DiscardUnknown(m)
}

var xxx_messageInfo_EventFilter proto.InternalMessageInfo

func (m *EventFilter) GetNamespaceId() string {
	if m != nil {
		return m.NamespaceId
	}
	return ""
}

func (m *EventFilter) GetLinkId() string {
	if m != nil {
		return m.LinkId
	}
	return ""
}

func (m *EventFilter) GetAliasId() string {
	if m != nil {
		return m.AliasId
	}
	return ""
}

func (m *EventFilter) GetTargetId() string {
	if m != nil {
		return m.TargetId
	}
	return ""
}

func (m *EventFilter) GetIdentifier() string {
	if m != nil {
		return m.Identifier
	}
	return ""
}

func (m *EventFilter) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *EventFilter) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *EventFilter) GetCreatedAt() *TimestampRange {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *EventFilter) GetLimit() uint32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

type ReadEventsRequest struct {
	// Types that are valid to be assigned to Filter:
	//	*ReadEventsRequest_Id
	//	*ReadEventsRequest_Condition
	Filter               isReadEventsRequest_Filter `protobuf_oneof:"filter"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *ReadEventsRequest) Reset()         { *m = ReadEventsRequest{} }
func (m *ReadEventsRequest) String() string { return proto.CompactTextString(m) }
func (*ReadEventsRequest) ProtoMessage()    {}
func (*ReadEventsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d17a9d3f0ddf27e, []int{2}
}
func (m *ReadEventsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadEventsRequest.Unmarshal(m, b)
}
func (m *ReadEventsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadEventsRequest.Marshal(b, m, deterministic)
}
func (m *ReadEventsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadEventsRequest.Merge(m, src)
}
func (m *ReadEventsRequest) XXX_Size() int {
	return xxx_messageInfo_ReadEventsRequest.Size(m)
}
func (m *ReadEventsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadEventsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReadEventsRequest proto.InternalMessageInfo

type isReadEventsRequest_Filter interface {
	isReadEventsRequest_Filter()
}

type ReadEventsRequest_Id struct {
	Id uint32 `protobuf:"varint,1,opt,name=id,proto3,oneof"`
}

type ReadEventsRequest_Condition struct {
	Condition *EventFilter `protobuf:"bytes,2,opt,name=condition,proto3,oneof"`
}

func (*ReadEventsRequest_Id) isReadEventsRequest_Filter() {}

func (*ReadEventsRequest_Condition) isReadEventsRequest_Filter() {}

func (m *ReadEventsRequest) GetFilter() isReadEventsRequest_Filter {
	if m != nil {
		return m.Filter
	}
	return nil
}

func (m *ReadEventsRequest) GetId() uint32 {
	if x, ok := m.GetFilter().(*ReadEventsRequest_Id); ok {
		return x.Id
	}
	return 0
}

func (m *ReadEventsRequest) GetCondition() *EventFilter {
	if x, ok := m.GetFilter().(*ReadEventsRequest_Condition); ok {
		return x.Condition
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*ReadEventsRequest) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _ReadEventsRequest_OneofMarshaler, _ReadEventsRequest_OneofUnmarshaler, _ReadEventsRequest_OneofSizer, []interface{}{
		(*ReadEventsRequest_Id)(nil),
		(*ReadEventsRequest_Condition)(nil),
	}
}

func _ReadEventsRequest_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*ReadEventsRequest)
	// filter
	switch x := m.Filter.(type) {
	case *ReadEventsRequest_Id:
		b.EncodeVarint(1<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.Id))
	case *ReadEventsRequest_Condition:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Condition); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("ReadEventsRequest.Filter has unexpected type %T", x)
	}
	return nil
}

func _ReadEventsRequest_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*ReadEventsRequest)
	switch tag {
	case 1: // filter.id
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Filter = &ReadEventsRequest_Id{uint32(x)}
		return true, err
	case 2: // filter.condition
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(EventFilter)
		err := b.DecodeMessage(msg)
		m.Filter = &ReadEventsRequest_Condition{msg}
		return true, err
	default:
		return false, nil
	}
}

func _ReadEventsRequest_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*ReadEventsRequest)
	// filter
	switch x := m.Filter.(type) {
	case *ReadEventsRequest_Id:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(x.Id))
	case *ReadEventsRequest_Condition:
		s := proto.Size(x.Condition)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type ReadEventsResponse struct {
	Events               []*EventEntry `protobuf:"bytes,1,rep,name=events,proto3" json:"events,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ReadEventsResponse) Reset()         { *m = ReadEventsResponse{} }
func (m *ReadEventsResponse) String() string { return proto.CompactTextString(m) }
func (*ReadEventsResponse) ProtoMessage()    {}
func (*ReadEventsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d17a9d3f0ddf27e, []int{3}
}
func (m *ReadEventsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadEventsResponse.Unmarshal(m, b)
}
func (m *ReadEventsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadEventsResponse.Marshal(b, m, deterministic)
}
func (m *ReadEventsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadEventsResponse.Merge(m, src)
}
func (m *ReadEventsResponse) XXX_Size() int {
	return xxx_messageInfo_ReadEventsResponse.Size(m)
}
func (m *ReadEventsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadEventsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ReadEventsResponse proto.InternalMessageInfo

func (m *ReadEventsResponse) GetEvents() []*EventEntry {
	if m != nil {
		return m.Events
	}
	return nil
}

type ListenEventsRequest struct {
	Filter               *EventFilter `protobuf:"bytes,1,opt,name=filter,proto3" json:"filter,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ListenEventsRequest) Reset()         { *m = ListenEventsRequest{} }
func (m *ListenEventsRequest) String() string { return proto.CompactTextString(m) }
func (*ListenEventsRequest) ProtoMessage()    {}
func (*ListenEventsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d17a9d3f0ddf27e, []int{4}
}
func (m *ListenEventsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListenEventsRequest.Unmarshal(m, b)
}
func (m *ListenEventsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListenEventsRequest.Marshal(b, m, deterministic)
}
func (m *ListenEventsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListenEventsRequest.Merge(m, src)
}
func (m *ListenEventsRequest) XXX_Size() int {
	return xxx_messageInfo_ListenEventsRequest.Size(m)
}
func (m *ListenEventsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListenEventsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListenEventsRequest proto.InternalMessageInfo

func (m *ListenEventsRequest) GetFilter() *EventFilter {
	if m != nil {
		return m.Filter
	}
	return nil
}

func init() {
	proto.RegisterType((*EventEntry)(nil), "grpc.EventEntry")
	proto.RegisterType((*EventFilter)(nil), "grpc.EventFilter")
	proto.RegisterType((*ReadEventsRequest)(nil), "grpc.ReadEventsRequest")
	proto.RegisterType((*ReadEventsResponse)(nil), "grpc.ReadEventsResponse")
	proto.RegisterType((*ListenEventsRequest)(nil), "grpc.ListenEventsRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// EventClient is the client API for Event service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EventClient interface {
	Read(ctx context.Context, in *ReadEventsRequest, opts ...grpc.CallOption) (*ReadEventsResponse, error)
	Listen(ctx context.Context, in *ListenEventsRequest, opts ...grpc.CallOption) (Event_ListenClient, error)
}

type eventClient struct {
	cc *grpc.ClientConn
}

func NewEventClient(cc *grpc.ClientConn) EventClient {
	return &eventClient{cc}
}

func (c *eventClient) Read(ctx context.Context, in *ReadEventsRequest, opts ...grpc.CallOption) (*ReadEventsResponse, error) {
	out := new(ReadEventsResponse)
	err := c.cc.Invoke(ctx, "/grpc.Event/Read", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventClient) Listen(ctx context.Context, in *ListenEventsRequest, opts ...grpc.CallOption) (Event_ListenClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Event_serviceDesc.Streams[0], "/grpc.Event/Listen", opts...)
	if err != nil {
		return nil, err
	}
	x := &eventListenClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Event_ListenClient interface {
	Recv() (*EventEntry, error)
	grpc.ClientStream
}

type eventListenClient struct {
	grpc.ClientStream
}

func (x *eventListenClient) Recv() (*EventEntry, error) {
	m := new(EventEntry)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// EventServer is the server API for Event service.
type EventServer interface {
	Read(context.Context, *ReadEventsRequest) (*ReadEventsResponse, error)
	Listen(*ListenEventsRequest, Event_ListenServer) error
}

func RegisterEventServer(s *grpc.Server, srv EventServer) {
	s.RegisterService(&_Event_serviceDesc, srv)
}

func _Event_Read_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadEventsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventServer).Read(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.Event/Read",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventServer).Read(ctx, req.(*ReadEventsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Event_Listen_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListenEventsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(EventServer).Listen(m, &eventListenServer{stream})
}

type Event_ListenServer interface {
	Send(*EventEntry) error
	grpc.ServerStream
}

type eventListenServer struct {
	grpc.ServerStream
}

func (x *eventListenServer) Send(m *EventEntry) error {
	return x.ServerStream.SendMsg(m)
}

var _Event_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.Event",
	HandlerType: (*EventServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Read",
			Handler:    _Event_Read_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Listen",
			Handler:       _Event_Listen_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "event.proto",
}

func init() { proto.RegisterFile("event.proto", fileDescriptor_2d17a9d3f0ddf27e) }

var fileDescriptor_2d17a9d3f0ddf27e = []byte{
	// 498 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x93, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x86, 0x63, 0x27, 0x71, 0x92, 0x71, 0x8a, 0xd2, 0xa1, 0x52, 0xdd, 0x20, 0x81, 0xf1, 0xc9,
	0x5c, 0x52, 0x48, 0x4f, 0x15, 0x12, 0x02, 0xa4, 0xa2, 0x46, 0xe2, 0x80, 0x56, 0xdc, 0x2b, 0xc7,
	0x3b, 0x89, 0x56, 0xd8, 0xbb, 0xc1, 0xde, 0x20, 0x38, 0xf1, 0x0e, 0xbc, 0x06, 0x2f, 0x89, 0x3c,
	0x4e, 0x88, 0x2b, 0xb7, 0xb9, 0x79, 0xe6, 0xff, 0xb5, 0xf3, 0xcf, 0xa7, 0x31, 0xf8, 0xf4, 0x83,
	0xb4, 0x9d, 0x6d, 0x0a, 0x63, 0x0d, 0xf6, 0xd6, 0xc5, 0x26, 0x9d, 0x8e, 0x53, 0x93, 0xe7, 0x46,
	0xd7, 0xbd, 0xe9, 0x8b, 0xb5, 0x31, 0xeb, 0x8c, 0x2e, 0xb9, 0x5a, 0x6e, 0x57, 0x97, 0x56, 0xe5,
	0x54, 0xda, 0x24, 0xdf, 0xd4, 0x86, 0xe8, 0xaf, 0x0b, 0x70, 0x53, 0x3d, 0x72, 0xa3, 0x6d, 0xf1,
	0x0b, 0x9f, 0x80, 0xab, 0x64, 0xe0, 0x84, 0x4e, 0x7c, 0x22, 0x5c, 0x25, 0xf1, 0x25, 0x8c, 0x75,
	0x92, 0x53, 0xb9, 0x49, 0x52, 0xba, 0x53, 0x32, 0x70, 0x43, 0x27, 0x1e, 0x09, 0xff, 0x7f, 0x6f,
	0x21, 0xf1, 0x1c, 0x06, 0x99, 0xd2, 0xdf, 0x2a, 0xb5, 0xcb, 0xaa, 0x57, 0x95, 0x0b, 0x89, 0x17,
	0x30, 0x4c, 0x32, 0x95, 0x94, 0x95, 0xd2, 0x63, 0x65, 0xc0, 0xf5, 0x42, 0xe2, 0x33, 0x18, 0xd9,
	0xa4, 0x58, 0x93, 0xad, 0xb4, 0x3e, 0x6b, 0xc3, 0xba, 0xb1, 0x90, 0xf8, 0x1c, 0x40, 0x49, 0xd2,
	0x56, 0xad, 0x14, 0x15, 0x81, 0xc7, 0x6a, 0xa3, 0x83, 0x01, 0x0c, 0x52, 0xa3, 0x2d, 0xfd, 0xb4,
	0xc1, 0x20, 0x74, 0xe2, 0xb1, 0xd8, 0x97, 0x88, 0xd0, 0x4b, 0x8d, 0xa4, 0x60, 0x18, 0x3a, 0x71,
	0x5f, 0xf0, 0x37, 0x4e, 0xa0, 0xbb, 0x2d, 0xb2, 0x60, 0xc4, 0xcf, 0x54, 0x9f, 0x78, 0x0d, 0x90,
	0x16, 0x94, 0x58, 0x92, 0x77, 0x89, 0x0d, 0x20, 0x74, 0x62, 0x7f, 0x3e, 0x9d, 0xd5, 0xa0, 0x66,
	0x7b, 0x50, 0xb3, 0xaf, 0x7b, 0x50, 0x62, 0xb4, 0x73, 0x7f, 0xb0, 0xd1, 0x1f, 0x17, 0x7c, 0xa6,
	0xf5, 0x49, 0x65, 0x96, 0x8a, 0x16, 0x1e, 0xe7, 0x28, 0x1e, 0xf7, 0x51, 0x3c, 0xdd, 0x23, 0x78,
	0x7a, 0x47, 0xf1, 0xf4, 0x5b, 0x78, 0xf6, 0x10, 0xbc, 0x36, 0x84, 0xc1, 0x01, 0xc2, 0xd5, 0x3d,
	0x08, 0x43, 0x86, 0x70, 0x36, 0xab, 0x2e, 0xa8, 0xb1, 0x79, 0xa2, 0xd7, 0xd4, 0x58, 0x1f, 0xcf,
	0xa0, 0x9f, 0xa9, 0x5c, 0x59, 0xa6, 0x79, 0x22, 0xea, 0x22, 0x5a, 0xc2, 0xa9, 0xa0, 0x44, 0x32,
	0x97, 0x52, 0xd0, 0xf7, 0x2d, 0x95, 0x16, 0x27, 0x87, 0x43, 0xba, 0xed, 0xf0, 0x29, 0xbd, 0x81,
	0x51, 0x6a, 0xb4, 0x54, 0x56, 0x19, 0xcd, 0x28, 0xfc, 0xf9, 0x69, 0x3d, 0xb0, 0x41, 0xf4, 0xb6,
	0x23, 0x0e, 0xae, 0x8f, 0x43, 0xf0, 0x56, 0xdc, 0x8e, 0xde, 0x01, 0x36, 0x67, 0x94, 0x1b, 0xa3,
	0x4b, 0xc2, 0x18, 0x3c, 0xfe, 0x01, 0xca, 0xc0, 0x09, 0xbb, 0xb1, 0x3f, 0x9f, 0x34, 0xde, 0xe3,
	0x7b, 0x16, 0x3b, 0x3d, 0x7a, 0x0f, 0x4f, 0x3f, 0xab, 0xd2, 0x92, 0xbe, 0x9f, 0xf2, 0xd5, 0x7e,
	0x00, 0x27, 0x7d, 0x28, 0x90, 0xd8, 0x19, 0xe6, 0xbf, 0xa1, 0xcf, 0x6d, 0x7c, 0x0b, 0xbd, 0x2a,
	0x0a, 0x9e, 0xd7, 0xde, 0xd6, 0xea, 0xd3, 0xa0, 0x2d, 0xd4, 0x79, 0xa3, 0x0e, 0x5e, 0x83, 0x57,
	0xe7, 0xc0, 0x8b, 0xda, 0xf5, 0x40, 0xaa, 0x69, 0x6b, 0x8d, 0xa8, 0xf3, 0xda, 0xf9, 0xd2, 0x59,
	0x7a, 0x7c, 0x9c, 0x57, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0x2c, 0x52, 0x58, 0xd0, 0xf6, 0x03,
	0x00, 0x00,
}
