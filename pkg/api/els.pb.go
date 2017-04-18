// Code generated by protoc-gen-go.
// source: els.proto
// DO NOT EDIT!

/*
Package api is a generated protocol buffer package.

It is generated from these files:
	els.proto

It has these top-level messages:
	RoutingKeyRequest
	ServiceInstanceReponse
	AddRoutingKeyRequest
*/
package api

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

type RoutingKeyRequest struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *RoutingKeyRequest) Reset()                    { *m = RoutingKeyRequest{} }
func (m *RoutingKeyRequest) String() string            { return proto.CompactTextString(m) }
func (*RoutingKeyRequest) ProtoMessage()               {}
func (*RoutingKeyRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *RoutingKeyRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type ServiceInstanceReponse struct {
	ServiceUri string `protobuf:"bytes,1,opt,name=serviceUri" json:"serviceUri,omitempty"`
	Tags       string `protobuf:"bytes,2,opt,name=tags" json:"tags,omitempty"`
}

func (m *ServiceInstanceReponse) Reset()                    { *m = ServiceInstanceReponse{} }
func (m *ServiceInstanceReponse) String() string            { return proto.CompactTextString(m) }
func (*ServiceInstanceReponse) ProtoMessage()               {}
func (*ServiceInstanceReponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ServiceInstanceReponse) GetServiceUri() string {
	if m != nil {
		return m.ServiceUri
	}
	return ""
}

func (m *ServiceInstanceReponse) GetTags() string {
	if m != nil {
		return m.Tags
	}
	return ""
}

type AddRoutingKeyRequest struct {
	ServiceUri string `protobuf:"bytes,1,opt,name=serviceUri" json:"serviceUri,omitempty"`
	Tags       string `protobuf:"bytes,2,opt,name=tags" json:"tags,omitempty"`
	RoutingKey string `protobuf:"bytes,3,opt,name=routingKey" json:"routingKey,omitempty"`
}

func (m *AddRoutingKeyRequest) Reset()                    { *m = AddRoutingKeyRequest{} }
func (m *AddRoutingKeyRequest) String() string            { return proto.CompactTextString(m) }
func (*AddRoutingKeyRequest) ProtoMessage()               {}
func (*AddRoutingKeyRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *AddRoutingKeyRequest) GetServiceUri() string {
	if m != nil {
		return m.ServiceUri
	}
	return ""
}

func (m *AddRoutingKeyRequest) GetTags() string {
	if m != nil {
		return m.Tags
	}
	return ""
}

func (m *AddRoutingKeyRequest) GetRoutingKey() string {
	if m != nil {
		return m.RoutingKey
	}
	return ""
}

func init() {
	proto.RegisterType((*RoutingKeyRequest)(nil), "api.RoutingKeyRequest")
	proto.RegisterType((*ServiceInstanceReponse)(nil), "api.ServiceInstanceReponse")
	proto.RegisterType((*AddRoutingKeyRequest)(nil), "api.AddRoutingKeyRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Els service

type ElsClient interface {
	// Get a service by routingKey
	GetServiceInstanceByKey(ctx context.Context, in *RoutingKeyRequest, opts ...grpc.CallOption) (*ServiceInstanceReponse, error)
	// Add a routingKey to a service
	AddRoutingKey(ctx context.Context, in *AddRoutingKeyRequest, opts ...grpc.CallOption) (*ServiceInstanceReponse, error)
}

type elsClient struct {
	cc *grpc.ClientConn
}

func NewElsClient(cc *grpc.ClientConn) ElsClient {
	return &elsClient{cc}
}

func (c *elsClient) GetServiceInstanceByKey(ctx context.Context, in *RoutingKeyRequest, opts ...grpc.CallOption) (*ServiceInstanceReponse, error) {
	out := new(ServiceInstanceReponse)
	err := grpc.Invoke(ctx, "/api.Els/GetServiceInstanceByKey", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *elsClient) AddRoutingKey(ctx context.Context, in *AddRoutingKeyRequest, opts ...grpc.CallOption) (*ServiceInstanceReponse, error) {
	out := new(ServiceInstanceReponse)
	err := grpc.Invoke(ctx, "/api.Els/AddRoutingKey", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Els service

type ElsServer interface {
	// Get a service by routingKey
	GetServiceInstanceByKey(context.Context, *RoutingKeyRequest) (*ServiceInstanceReponse, error)
	// Add a routingKey to a service
	AddRoutingKey(context.Context, *AddRoutingKeyRequest) (*ServiceInstanceReponse, error)
}

func RegisterElsServer(s *grpc.Server, srv ElsServer) {
	s.RegisterService(&_Els_serviceDesc, srv)
}

func _Els_GetServiceInstanceByKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoutingKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ElsServer).GetServiceInstanceByKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Els/GetServiceInstanceByKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ElsServer).GetServiceInstanceByKey(ctx, req.(*RoutingKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Els_AddRoutingKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddRoutingKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ElsServer).AddRoutingKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Els/AddRoutingKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ElsServer).AddRoutingKey(ctx, req.(*AddRoutingKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Els_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.Els",
	HandlerType: (*ElsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetServiceInstanceByKey",
			Handler:    _Els_GetServiceInstanceByKey_Handler,
		},
		{
			MethodName: "AddRoutingKey",
			Handler:    _Els_AddRoutingKey_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "els.proto",
}

func init() { proto.RegisterFile("els.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 214 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4c, 0xcd, 0x29, 0xd6,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4e, 0x2c, 0xc8, 0x54, 0x52, 0xe6, 0x12, 0x0c, 0xca,
	0x2f, 0x2d, 0xc9, 0xcc, 0x4b, 0xf7, 0x4e, 0xad, 0x0c, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11,
	0xe2, 0xe3, 0x62, 0xca, 0x4c, 0x91, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x62, 0xca, 0x4c, 0x51,
	0xf2, 0xe1, 0x12, 0x0b, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xf5, 0xcc, 0x2b, 0x2e, 0x49, 0xcc,
	0x4b, 0x4e, 0x0d, 0x4a, 0x2d, 0xc8, 0xcf, 0x2b, 0x4e, 0x15, 0x92, 0xe3, 0xe2, 0x2a, 0x86, 0xc8,
	0x84, 0x16, 0x65, 0x42, 0x75, 0x20, 0x89, 0x08, 0x09, 0x71, 0xb1, 0x94, 0x24, 0xa6, 0x17, 0x4b,
	0x30, 0x81, 0x65, 0xc0, 0x6c, 0xa5, 0x2c, 0x2e, 0x11, 0xc7, 0x94, 0x14, 0x4c, 0x5b, 0xc9, 0x30,
	0x0b, 0xa4, 0xa7, 0x08, 0x6e, 0x90, 0x04, 0x33, 0x44, 0x0f, 0x42, 0xc4, 0x68, 0x11, 0x23, 0x17,
	0xb3, 0x6b, 0x4e, 0xb1, 0x50, 0x00, 0x97, 0xb8, 0x7b, 0x6a, 0x09, 0x9a, 0x27, 0x9c, 0x2a, 0xbd,
	0x53, 0x2b, 0x85, 0xc4, 0xf4, 0x12, 0x0b, 0x32, 0xf5, 0x30, 0x9c, 0x23, 0x25, 0x0d, 0x16, 0xc7,
	0xee, 0x6f, 0x25, 0x06, 0x21, 0x4f, 0x2e, 0x5e, 0x14, 0x5f, 0x08, 0x49, 0x82, 0xd5, 0x63, 0xf3,
	0x19, 0x01, 0xa3, 0x92, 0xd8, 0xc0, 0xf1, 0x61, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x97, 0xd9,
	0x5c, 0xa2, 0x9c, 0x01, 0x00, 0x00,
}
