// Code generated by protoc-gen-go. DO NOT EDIT.
// source: calendar_server.proto

package grpc

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type EventType int32

const (
	EventType_UNKNOWN_EVENT_TYPE EventType = 0
	EventType_TASK               EventType = 1
	EventType_MEETING            EventType = 2
)

var EventType_name = map[int32]string{
	0: "UNKNOWN_EVENT_TYPE",
	1: "TASK",
	2: "MEETING",
}

var EventType_value = map[string]int32{
	"UNKNOWN_EVENT_TYPE": 0,
	"TASK":               1,
	"MEETING":            2,
}

func (x EventType) String() string {
	return proto.EnumName(EventType_name, int32(x))
}

func (EventType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e3d25d49f056cdb2, []int{0}
}

type CreateCalendarEventRequest struct {
	Name                 string               `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Type                 EventType            `protobuf:"varint,2,opt,name=type,proto3,enum=grpc.EventType" json:"type,omitempty"`
	StartTime            *timestamp.Timestamp `protobuf:"bytes,3,opt,name=startTime,proto3" json:"startTime,omitempty"`
	EndTime              *timestamp.Timestamp `protobuf:"bytes,4,opt,name=endTime,proto3" json:"endTime,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *CreateCalendarEventRequest) Reset()         { *m = CreateCalendarEventRequest{} }
func (m *CreateCalendarEventRequest) String() string { return proto.CompactTextString(m) }
func (*CreateCalendarEventRequest) ProtoMessage()    {}
func (*CreateCalendarEventRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e3d25d49f056cdb2, []int{0}
}

func (m *CreateCalendarEventRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateCalendarEventRequest.Unmarshal(m, b)
}
func (m *CreateCalendarEventRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateCalendarEventRequest.Marshal(b, m, deterministic)
}
func (m *CreateCalendarEventRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateCalendarEventRequest.Merge(m, src)
}
func (m *CreateCalendarEventRequest) XXX_Size() int {
	return xxx_messageInfo_CreateCalendarEventRequest.Size(m)
}
func (m *CreateCalendarEventRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateCalendarEventRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateCalendarEventRequest proto.InternalMessageInfo

func (m *CreateCalendarEventRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateCalendarEventRequest) GetType() EventType {
	if m != nil {
		return m.Type
	}
	return EventType_UNKNOWN_EVENT_TYPE
}

func (m *CreateCalendarEventRequest) GetStartTime() *timestamp.Timestamp {
	if m != nil {
		return m.StartTime
	}
	return nil
}

func (m *CreateCalendarEventRequest) GetEndTime() *timestamp.Timestamp {
	if m != nil {
		return m.EndTime
	}
	return nil
}

type CalendarEventObject struct {
	Id                   int64                `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string               `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Type                 EventType            `protobuf:"varint,3,opt,name=type,proto3,enum=grpc.EventType" json:"type,omitempty"`
	StartTime            *timestamp.Timestamp `protobuf:"bytes,4,opt,name=startTime,proto3" json:"startTime,omitempty"`
	EndTime              *timestamp.Timestamp `protobuf:"bytes,5,opt,name=endTime,proto3" json:"endTime,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *CalendarEventObject) Reset()         { *m = CalendarEventObject{} }
func (m *CalendarEventObject) String() string { return proto.CompactTextString(m) }
func (*CalendarEventObject) ProtoMessage()    {}
func (*CalendarEventObject) Descriptor() ([]byte, []int) {
	return fileDescriptor_e3d25d49f056cdb2, []int{1}
}

func (m *CalendarEventObject) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CalendarEventObject.Unmarshal(m, b)
}
func (m *CalendarEventObject) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CalendarEventObject.Marshal(b, m, deterministic)
}
func (m *CalendarEventObject) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CalendarEventObject.Merge(m, src)
}
func (m *CalendarEventObject) XXX_Size() int {
	return xxx_messageInfo_CalendarEventObject.Size(m)
}
func (m *CalendarEventObject) XXX_DiscardUnknown() {
	xxx_messageInfo_CalendarEventObject.DiscardUnknown(m)
}

var xxx_messageInfo_CalendarEventObject proto.InternalMessageInfo

func (m *CalendarEventObject) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *CalendarEventObject) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CalendarEventObject) GetType() EventType {
	if m != nil {
		return m.Type
	}
	return EventType_UNKNOWN_EVENT_TYPE
}

func (m *CalendarEventObject) GetStartTime() *timestamp.Timestamp {
	if m != nil {
		return m.StartTime
	}
	return nil
}

func (m *CalendarEventObject) GetEndTime() *timestamp.Timestamp {
	if m != nil {
		return m.EndTime
	}
	return nil
}

type ListCalendarEventRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListCalendarEventRequest) Reset()         { *m = ListCalendarEventRequest{} }
func (m *ListCalendarEventRequest) String() string { return proto.CompactTextString(m) }
func (*ListCalendarEventRequest) ProtoMessage()    {}
func (*ListCalendarEventRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e3d25d49f056cdb2, []int{2}
}

func (m *ListCalendarEventRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListCalendarEventRequest.Unmarshal(m, b)
}
func (m *ListCalendarEventRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListCalendarEventRequest.Marshal(b, m, deterministic)
}
func (m *ListCalendarEventRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListCalendarEventRequest.Merge(m, src)
}
func (m *ListCalendarEventRequest) XXX_Size() int {
	return xxx_messageInfo_ListCalendarEventRequest.Size(m)
}
func (m *ListCalendarEventRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListCalendarEventRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListCalendarEventRequest proto.InternalMessageInfo

type ListCalendarEventResponse struct {
	Events               []*CalendarEventObject `protobuf:"bytes,1,rep,name=events,proto3" json:"events,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *ListCalendarEventResponse) Reset()         { *m = ListCalendarEventResponse{} }
func (m *ListCalendarEventResponse) String() string { return proto.CompactTextString(m) }
func (*ListCalendarEventResponse) ProtoMessage()    {}
func (*ListCalendarEventResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e3d25d49f056cdb2, []int{3}
}

func (m *ListCalendarEventResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListCalendarEventResponse.Unmarshal(m, b)
}
func (m *ListCalendarEventResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListCalendarEventResponse.Marshal(b, m, deterministic)
}
func (m *ListCalendarEventResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListCalendarEventResponse.Merge(m, src)
}
func (m *ListCalendarEventResponse) XXX_Size() int {
	return xxx_messageInfo_ListCalendarEventResponse.Size(m)
}
func (m *ListCalendarEventResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListCalendarEventResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListCalendarEventResponse proto.InternalMessageInfo

func (m *ListCalendarEventResponse) GetEvents() []*CalendarEventObject {
	if m != nil {
		return m.Events
	}
	return nil
}

type GetCalendarEventRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetCalendarEventRequest) Reset()         { *m = GetCalendarEventRequest{} }
func (m *GetCalendarEventRequest) String() string { return proto.CompactTextString(m) }
func (*GetCalendarEventRequest) ProtoMessage()    {}
func (*GetCalendarEventRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e3d25d49f056cdb2, []int{4}
}

func (m *GetCalendarEventRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCalendarEventRequest.Unmarshal(m, b)
}
func (m *GetCalendarEventRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCalendarEventRequest.Marshal(b, m, deterministic)
}
func (m *GetCalendarEventRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCalendarEventRequest.Merge(m, src)
}
func (m *GetCalendarEventRequest) XXX_Size() int {
	return xxx_messageInfo_GetCalendarEventRequest.Size(m)
}
func (m *GetCalendarEventRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCalendarEventRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetCalendarEventRequest proto.InternalMessageInfo

func (m *GetCalendarEventRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type DeleteCalendarEventRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteCalendarEventRequest) Reset()         { *m = DeleteCalendarEventRequest{} }
func (m *DeleteCalendarEventRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteCalendarEventRequest) ProtoMessage()    {}
func (*DeleteCalendarEventRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e3d25d49f056cdb2, []int{5}
}

func (m *DeleteCalendarEventRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteCalendarEventRequest.Unmarshal(m, b)
}
func (m *DeleteCalendarEventRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteCalendarEventRequest.Marshal(b, m, deterministic)
}
func (m *DeleteCalendarEventRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteCalendarEventRequest.Merge(m, src)
}
func (m *DeleteCalendarEventRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteCalendarEventRequest.Size(m)
}
func (m *DeleteCalendarEventRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteCalendarEventRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteCalendarEventRequest proto.InternalMessageInfo

func (m *DeleteCalendarEventRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type DeleteCalendarEventResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteCalendarEventResponse) Reset()         { *m = DeleteCalendarEventResponse{} }
func (m *DeleteCalendarEventResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteCalendarEventResponse) ProtoMessage()    {}
func (*DeleteCalendarEventResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e3d25d49f056cdb2, []int{6}
}

func (m *DeleteCalendarEventResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteCalendarEventResponse.Unmarshal(m, b)
}
func (m *DeleteCalendarEventResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteCalendarEventResponse.Marshal(b, m, deterministic)
}
func (m *DeleteCalendarEventResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteCalendarEventResponse.Merge(m, src)
}
func (m *DeleteCalendarEventResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteCalendarEventResponse.Size(m)
}
func (m *DeleteCalendarEventResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteCalendarEventResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteCalendarEventResponse proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("grpc.EventType", EventType_name, EventType_value)
	proto.RegisterType((*CreateCalendarEventRequest)(nil), "grpc.CreateCalendarEventRequest")
	proto.RegisterType((*CalendarEventObject)(nil), "grpc.CalendarEventObject")
	proto.RegisterType((*ListCalendarEventRequest)(nil), "grpc.ListCalendarEventRequest")
	proto.RegisterType((*ListCalendarEventResponse)(nil), "grpc.ListCalendarEventResponse")
	proto.RegisterType((*GetCalendarEventRequest)(nil), "grpc.GetCalendarEventRequest")
	proto.RegisterType((*DeleteCalendarEventRequest)(nil), "grpc.DeleteCalendarEventRequest")
	proto.RegisterType((*DeleteCalendarEventResponse)(nil), "grpc.DeleteCalendarEventResponse")
}

func init() { proto.RegisterFile("calendar_server.proto", fileDescriptor_e3d25d49f056cdb2) }

var fileDescriptor_e3d25d49f056cdb2 = []byte{
	// 438 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x52, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0x65, 0x1d, 0xd3, 0x36, 0x63, 0x29, 0x8d, 0x06, 0x09, 0xdc, 0x45, 0xa5, 0xc6, 0x5c, 0x02,
	0x42, 0xae, 0x08, 0x1c, 0x10, 0x37, 0xd4, 0x5a, 0x11, 0x94, 0xba, 0xc8, 0x35, 0x20, 0x4e, 0x95,
	0x63, 0x0f, 0x91, 0x51, 0xfc, 0x81, 0x77, 0x5b, 0xa9, 0x3f, 0x10, 0x8e, 0xfc, 0x26, 0x94, 0x5d,
	0xbb, 0xa8, 0x95, 0x5d, 0x22, 0x6e, 0xab, 0x99, 0x37, 0xb3, 0xef, 0xbd, 0x79, 0x30, 0x4a, 0xe2,
	0x25, 0x15, 0x69, 0x5c, 0x7b, 0x55, 0x5d, 0xca, 0x12, 0xcd, 0x45, 0x5d, 0x25, 0x7c, 0x6f, 0x51,
	0x96, 0x8b, 0x25, 0xed, 0xab, 0xda, 0xfc, 0xfc, 0xdb, 0xbe, 0xcc, 0x72, 0x12, 0x32, 0xce, 0x2b,
	0x0d, 0x73, 0x7f, 0x32, 0xe0, 0x07, 0x35, 0xc5, 0x92, 0x0e, 0x9a, 0x79, 0xff, 0x82, 0x0a, 0x19,
	0xd2, 0x8f, 0x73, 0x12, 0x12, 0x11, 0xcc, 0x22, 0xce, 0xc9, 0x66, 0x0e, 0x9b, 0x0c, 0x43, 0xf5,
	0xc6, 0x27, 0x60, 0xca, 0xcb, 0x8a, 0x6c, 0xc3, 0x61, 0x93, 0xd1, 0x74, 0xdb, 0x5b, 0x7d, 0xe4,
	0xa9, 0xa9, 0xe8, 0xb2, 0xa2, 0x50, 0x35, 0xf1, 0x35, 0x0c, 0x85, 0x8c, 0x6b, 0x19, 0x65, 0x39,
	0xd9, 0x03, 0x87, 0x4d, 0xac, 0x29, 0xf7, 0x34, 0x19, 0xaf, 0x25, 0xe3, 0x45, 0x2d, 0x99, 0xf0,
	0x2f, 0x18, 0x5f, 0xc1, 0x26, 0x15, 0xa9, 0x9a, 0x33, 0xff, 0x39, 0xd7, 0x42, 0xdd, 0xdf, 0x0c,
	0xee, 0x5d, 0x53, 0x70, 0x32, 0xff, 0x4e, 0x89, 0xc4, 0x11, 0x18, 0x59, 0xaa, 0xe8, 0x0f, 0x42,
	0x23, 0x4b, 0xaf, 0x04, 0x19, 0x1d, 0x82, 0x06, 0x6b, 0x0b, 0x32, 0xff, 0x53, 0xd0, 0xdd, 0xf5,
	0x05, 0x71, 0xb0, 0x3f, 0x64, 0x42, 0x76, 0x5d, 0xc5, 0x0d, 0x60, 0xa7, 0xa3, 0x27, 0xaa, 0xb2,
	0x10, 0x84, 0x2f, 0x60, 0x83, 0x56, 0x05, 0x61, 0x33, 0x67, 0x30, 0xb1, 0xa6, 0x3b, 0x5a, 0x4f,
	0x87, 0x39, 0x61, 0x03, 0x74, 0x9f, 0xc2, 0x83, 0x19, 0x75, 0x7e, 0x75, 0xd3, 0x3f, 0xf7, 0x39,
	0xf0, 0x43, 0x5a, 0x52, 0x4f, 0x5c, 0x6e, 0xa2, 0x77, 0xe1, 0x61, 0x27, 0x5a, 0x53, 0x7d, 0xf6,
	0x06, 0x86, 0x57, 0x36, 0xe3, 0x7d, 0xc0, 0x4f, 0xc1, 0x51, 0x70, 0xf2, 0x25, 0x38, 0xf3, 0x3f,
	0xfb, 0x41, 0x74, 0x16, 0x7d, 0xfd, 0xe8, 0x8f, 0xef, 0xe0, 0x16, 0x98, 0xd1, 0xdb, 0xd3, 0xa3,
	0x31, 0x43, 0x0b, 0x36, 0x8f, 0x7d, 0x3f, 0x7a, 0x17, 0xcc, 0xc6, 0xc6, 0xf4, 0x97, 0x01, 0xdb,
	0xed, 0xd6, 0x53, 0xaa, 0x2f, 0xb2, 0x84, 0xf0, 0x3d, 0x58, 0x3a, 0xcb, 0x6a, 0x2b, 0x3a, 0x8d,
	0xf2, 0xde, 0x78, 0xf3, 0x7e, 0x6f, 0xf0, 0x18, 0x60, 0xe5, 0xb1, 0x2a, 0x09, 0x7c, 0xa4, 0x81,
	0x7d, 0x17, 0xe1, 0x7b, 0xbd, 0xfd, 0xe6, 0x2a, 0x87, 0xb0, 0x35, 0x23, 0xbd, 0x0d, 0x77, 0x35,
	0xb8, 0xc7, 0xf2, 0xdb, 0x48, 0x85, 0x60, 0x69, 0x3f, 0xaf, 0x09, 0xec, 0x3f, 0x08, 0x7f, 0x7c,
	0x0b, 0x42, 0x33, 0x9b, 0x6f, 0xa8, 0x14, 0xbe, 0xfc, 0x13, 0x00, 0x00, 0xff, 0xff, 0x18, 0x0e,
	0x68, 0xc2, 0x41, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CalendarServiceClient is the client API for CalendarService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CalendarServiceClient interface {
	CreateEvent(ctx context.Context, in *CreateCalendarEventRequest, opts ...grpc.CallOption) (*CalendarEventObject, error)
	ListEvents(ctx context.Context, in *ListCalendarEventRequest, opts ...grpc.CallOption) (*ListCalendarEventResponse, error)
	GetEvent(ctx context.Context, in *GetCalendarEventRequest, opts ...grpc.CallOption) (*CalendarEventObject, error)
	DeleteEvent(ctx context.Context, in *DeleteCalendarEventRequest, opts ...grpc.CallOption) (*DeleteCalendarEventResponse, error)
}

type calendarServiceClient struct {
	cc *grpc.ClientConn
}

func NewCalendarServiceClient(cc *grpc.ClientConn) CalendarServiceClient {
	return &calendarServiceClient{cc}
}

func (c *calendarServiceClient) CreateEvent(ctx context.Context, in *CreateCalendarEventRequest, opts ...grpc.CallOption) (*CalendarEventObject, error) {
	out := new(CalendarEventObject)
	err := c.cc.Invoke(ctx, "/grpc.CalendarService/CreateEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calendarServiceClient) ListEvents(ctx context.Context, in *ListCalendarEventRequest, opts ...grpc.CallOption) (*ListCalendarEventResponse, error) {
	out := new(ListCalendarEventResponse)
	err := c.cc.Invoke(ctx, "/grpc.CalendarService/ListEvents", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calendarServiceClient) GetEvent(ctx context.Context, in *GetCalendarEventRequest, opts ...grpc.CallOption) (*CalendarEventObject, error) {
	out := new(CalendarEventObject)
	err := c.cc.Invoke(ctx, "/grpc.CalendarService/GetEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calendarServiceClient) DeleteEvent(ctx context.Context, in *DeleteCalendarEventRequest, opts ...grpc.CallOption) (*DeleteCalendarEventResponse, error) {
	out := new(DeleteCalendarEventResponse)
	err := c.cc.Invoke(ctx, "/grpc.CalendarService/DeleteEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CalendarServiceServer is the server API for CalendarService service.
type CalendarServiceServer interface {
	CreateEvent(context.Context, *CreateCalendarEventRequest) (*CalendarEventObject, error)
	ListEvents(context.Context, *ListCalendarEventRequest) (*ListCalendarEventResponse, error)
	GetEvent(context.Context, *GetCalendarEventRequest) (*CalendarEventObject, error)
	DeleteEvent(context.Context, *DeleteCalendarEventRequest) (*DeleteCalendarEventResponse, error)
}

// UnimplementedCalendarServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCalendarServiceServer struct {
}

func (*UnimplementedCalendarServiceServer) CreateEvent(ctx context.Context, req *CreateCalendarEventRequest) (*CalendarEventObject, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateEvent not implemented")
}
func (*UnimplementedCalendarServiceServer) ListEvents(ctx context.Context, req *ListCalendarEventRequest) (*ListCalendarEventResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListEvents not implemented")
}
func (*UnimplementedCalendarServiceServer) GetEvent(ctx context.Context, req *GetCalendarEventRequest) (*CalendarEventObject, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEvent not implemented")
}
func (*UnimplementedCalendarServiceServer) DeleteEvent(ctx context.Context, req *DeleteCalendarEventRequest) (*DeleteCalendarEventResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteEvent not implemented")
}

func RegisterCalendarServiceServer(s *grpc.Server, srv CalendarServiceServer) {
	s.RegisterService(&_CalendarService_serviceDesc, srv)
}

func _CalendarService_CreateEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCalendarEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalendarServiceServer).CreateEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.CalendarService/CreateEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalendarServiceServer).CreateEvent(ctx, req.(*CreateCalendarEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalendarService_ListEvents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCalendarEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalendarServiceServer).ListEvents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.CalendarService/ListEvents",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalendarServiceServer).ListEvents(ctx, req.(*ListCalendarEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalendarService_GetEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCalendarEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalendarServiceServer).GetEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.CalendarService/GetEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalendarServiceServer).GetEvent(ctx, req.(*GetCalendarEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalendarService_DeleteEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCalendarEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalendarServiceServer).DeleteEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.CalendarService/DeleteEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalendarServiceServer).DeleteEvent(ctx, req.(*DeleteCalendarEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CalendarService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.CalendarService",
	HandlerType: (*CalendarServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateEvent",
			Handler:    _CalendarService_CreateEvent_Handler,
		},
		{
			MethodName: "ListEvents",
			Handler:    _CalendarService_ListEvents_Handler,
		},
		{
			MethodName: "GetEvent",
			Handler:    _CalendarService_GetEvent_Handler,
		},
		{
			MethodName: "DeleteEvent",
			Handler:    _CalendarService_DeleteEvent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "calendar_server.proto",
}
