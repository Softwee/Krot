// Code generated by protoc-gen-go.
// source: rpc.proto
// DO NOT EDIT!

/*
Package rpc is a generated protocol buffer package.

It is generated from these files:
	rpc.proto

It has these top-level messages:
	Subscription
	SubscriptionId
	Response
*/
package rpc

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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

type Subscription struct {
	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	Type   uint32 `protobuf:"varint,2,opt,name=type" json:"type,omitempty"`
	Url    string `protobuf:"bytes,3,opt,name=url" json:"url,omitempty"`
	Tag    string `protobuf:"bytes,4,opt,name=tag" json:"tag,omitempty"`
	PollMs uint32 `protobuf:"varint,5,opt,name=poll_ms,json=pollMs" json:"poll_ms,omitempty"`
	Status uint32 `protobuf:"varint,6,opt,name=status" json:"status,omitempty"`
}

func (m *Subscription) Reset()                    { *m = Subscription{} }
func (m *Subscription) String() string            { return proto.CompactTextString(m) }
func (*Subscription) ProtoMessage()               {}
func (*Subscription) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type SubscriptionId struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *SubscriptionId) Reset()                    { *m = SubscriptionId{} }
func (m *SubscriptionId) String() string            { return proto.CompactTextString(m) }
func (*SubscriptionId) ProtoMessage()               {}
func (*SubscriptionId) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type Response struct {
	Success bool   `protobuf:"varint,1,opt,name=success" json:"success,omitempty"`
	Error   string `protobuf:"bytes,2,opt,name=error" json:"error,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func init() {
	proto.RegisterType((*Subscription)(nil), "Subscription")
	proto.RegisterType((*SubscriptionId)(nil), "SubscriptionId")
	proto.RegisterType((*Response)(nil), "Response")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for SubscriptionService service

type SubscriptionServiceClient interface {
	Subscribe(ctx context.Context, in *Subscription, opts ...grpc.CallOption) (*Response, error)
	ResumeSubscription(ctx context.Context, in *SubscriptionId, opts ...grpc.CallOption) (*Response, error)
	StopSubscription(ctx context.Context, in *SubscriptionId, opts ...grpc.CallOption) (*Response, error)
	Unsubscribe(ctx context.Context, in *SubscriptionId, opts ...grpc.CallOption) (*Response, error)
}

type subscriptionServiceClient struct {
	cc *grpc.ClientConn
}

func NewSubscriptionServiceClient(cc *grpc.ClientConn) SubscriptionServiceClient {
	return &subscriptionServiceClient{cc}
}

func (c *subscriptionServiceClient) Subscribe(ctx context.Context, in *Subscription, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/SubscriptionService/Subscribe", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subscriptionServiceClient) ResumeSubscription(ctx context.Context, in *SubscriptionId, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/SubscriptionService/ResumeSubscription", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subscriptionServiceClient) StopSubscription(ctx context.Context, in *SubscriptionId, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/SubscriptionService/StopSubscription", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subscriptionServiceClient) Unsubscribe(ctx context.Context, in *SubscriptionId, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/SubscriptionService/Unsubscribe", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for SubscriptionService service

type SubscriptionServiceServer interface {
	Subscribe(context.Context, *Subscription) (*Response, error)
	ResumeSubscription(context.Context, *SubscriptionId) (*Response, error)
	StopSubscription(context.Context, *SubscriptionId) (*Response, error)
	Unsubscribe(context.Context, *SubscriptionId) (*Response, error)
}

func RegisterSubscriptionServiceServer(s *grpc.Server, srv SubscriptionServiceServer) {
	s.RegisterService(&_SubscriptionService_serviceDesc, srv)
}

func _SubscriptionService_Subscribe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Subscription)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubscriptionServiceServer).Subscribe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SubscriptionService/Subscribe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubscriptionServiceServer).Subscribe(ctx, req.(*Subscription))
	}
	return interceptor(ctx, in, info, handler)
}

func _SubscriptionService_ResumeSubscription_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubscriptionId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubscriptionServiceServer).ResumeSubscription(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SubscriptionService/ResumeSubscription",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubscriptionServiceServer).ResumeSubscription(ctx, req.(*SubscriptionId))
	}
	return interceptor(ctx, in, info, handler)
}

func _SubscriptionService_StopSubscription_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubscriptionId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubscriptionServiceServer).StopSubscription(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SubscriptionService/StopSubscription",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubscriptionServiceServer).StopSubscription(ctx, req.(*SubscriptionId))
	}
	return interceptor(ctx, in, info, handler)
}

func _SubscriptionService_Unsubscribe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubscriptionId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubscriptionServiceServer).Unsubscribe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SubscriptionService/Unsubscribe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubscriptionServiceServer).Unsubscribe(ctx, req.(*SubscriptionId))
	}
	return interceptor(ctx, in, info, handler)
}

var _SubscriptionService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "SubscriptionService",
	HandlerType: (*SubscriptionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Subscribe",
			Handler:    _SubscriptionService_Subscribe_Handler,
		},
		{
			MethodName: "ResumeSubscription",
			Handler:    _SubscriptionService_ResumeSubscription_Handler,
		},
		{
			MethodName: "StopSubscription",
			Handler:    _SubscriptionService_StopSubscription_Handler,
		},
		{
			MethodName: "Unsubscribe",
			Handler:    _SubscriptionService_Unsubscribe_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor0,
}

func init() { proto.RegisterFile("rpc.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 277 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x91, 0xcf, 0x4a, 0xc3, 0x40,
	0x10, 0xc6, 0x4d, 0xff, 0xa4, 0xcd, 0x68, 0x6b, 0x19, 0x45, 0x17, 0x4f, 0x25, 0x17, 0x0b, 0x42,
	0x90, 0x7a, 0xf3, 0x0d, 0x7a, 0xf0, 0xb2, 0xc1, 0x73, 0xc9, 0x9f, 0xa1, 0x04, 0xd2, 0xec, 0xb2,
	0xb3, 0x2b, 0xf8, 0x16, 0xbe, 0x9a, 0x6f, 0x24, 0xd9, 0xd8, 0x92, 0x5c, 0x04, 0x6f, 0xf3, 0x7d,
	0x7c, 0xdf, 0xec, 0x8f, 0x1d, 0x88, 0x8c, 0x2e, 0x12, 0x6d, 0x94, 0x55, 0xf1, 0x57, 0x00, 0x57,
	0xa9, 0xcb, 0xb9, 0x30, 0x95, 0xb6, 0x95, 0x6a, 0xf0, 0x1e, 0x66, 0x8e, 0xc9, 0xec, 0xab, 0x52,
	0x04, 0xeb, 0x60, 0x13, 0xc9, 0xb0, 0x95, 0xbb, 0x12, 0x11, 0x26, 0xf6, 0x53, 0x93, 0x18, 0xad,
	0x83, 0xcd, 0x42, 0xfa, 0x19, 0x57, 0x30, 0x76, 0xa6, 0x16, 0x63, 0x1f, 0x6c, 0xc7, 0xd6, 0xb1,
	0xd9, 0x41, 0x4c, 0x3a, 0xc7, 0x66, 0x87, 0x76, 0xa1, 0x56, 0x75, 0xbd, 0x3f, 0xb2, 0x98, 0xfa,
	0x6a, 0xd8, 0xca, 0x37, 0xc6, 0x3b, 0x08, 0xd9, 0x66, 0xd6, 0xb1, 0x08, 0x3b, 0xbf, 0x53, 0xf1,
	0x1a, 0x96, 0x7d, 0xa2, 0x5d, 0x89, 0x4b, 0x18, 0x9d, 0x71, 0x46, 0x55, 0x19, 0xbf, 0xc2, 0x5c,
	0x12, 0x6b, 0xd5, 0x30, 0xa1, 0x80, 0x19, 0xbb, 0xa2, 0x20, 0x66, 0x1f, 0x98, 0xcb, 0x93, 0xc4,
	0x5b, 0x98, 0x92, 0x31, 0xca, 0x78, 0xe2, 0x48, 0x76, 0x62, 0xfb, 0x1d, 0xc0, 0x4d, 0x7f, 0x7d,
	0x4a, 0xe6, 0xa3, 0x2a, 0x08, 0x1f, 0x21, 0xfa, 0xb5, 0x73, 0xc2, 0x45, 0xd2, 0x8f, 0x3c, 0x44,
	0xc9, 0xe9, 0xb9, 0xf8, 0x02, 0xb7, 0x80, 0x92, 0xd8, 0x1d, 0x69, 0xf0, 0x6d, 0xd7, 0xc9, 0x90,
	0x79, 0xd8, 0x79, 0x86, 0x55, 0x6a, 0x95, 0xfe, 0x47, 0xe3, 0x09, 0x2e, 0xdf, 0x1b, 0x3e, 0x03,
	0xfd, 0x19, 0xce, 0x43, 0x7f, 0xcb, 0x97, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x8f, 0xf8, 0xc2,
	0x00, 0xd8, 0x01, 0x00, 0x00,
}
