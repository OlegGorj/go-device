// Code generated by protoc-gen-go.
// source: device.proto
// DO NOT EDIT!

/*
Package device is a generated protocol buffer package.

It is generated from these files:
	device.proto

It has these top-level messages:
	DeviceName
	Device
*/
package device

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "golang.org/x/net/context"
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

type DeviceName struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *DeviceName) Reset()                    { *m = DeviceName{} }
func (m *DeviceName) String() string            { return proto.CompactTextString(m) }
func (*DeviceName) ProtoMessage()               {}
func (*DeviceName) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *DeviceName) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Device struct {
	Name   string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Active bool   `protobuf:"varint,2,opt,name=active" json:"active,omitempty"`
}

func (m *Device) Reset()                    { *m = Device{} }
func (m *Device) String() string            { return proto.CompactTextString(m) }
func (*Device) ProtoMessage()               {}
func (*Device) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Device) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Device) GetActive() bool {
	if m != nil {
		return m.Active
	}
	return false
}

func init() {
	proto.RegisterType((*DeviceName)(nil), "DeviceName")
	proto.RegisterType((*Device)(nil), "Device")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Publisher API

type Publisher interface {
	Publish(ctx context.Context, msg interface{}, opts ...client.PublishOption) error
}

type publisher struct {
	c     client.Client
	topic string
}

func (p *publisher) Publish(ctx context.Context, msg interface{}, opts ...client.PublishOption) error {
	return p.c.Publish(ctx, p.c.NewPublication(p.topic, msg), opts...)
}

func NewPublisher(topic string, c client.Client) Publisher {
	if c == nil {
		c = client.NewClient()
	}
	return &publisher{c, topic}
}

// Subscriber API

func RegisterSubscriber(topic string, s server.Server, h interface{}, opts ...server.SubscriberOption) error {
	return s.Subscribe(s.NewSubscriber(topic, h, opts...))
}

// Client API for DevSvc service

type DevSvcClient interface {
	Get(ctx context.Context, in *DeviceName, opts ...client.CallOption) (*Device, error)
}

type devSvcClient struct {
	c           client.Client
	serviceName string
}

func NewDevSvcClient(serviceName string, c client.Client) DevSvcClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "devsvc"
	}
	return &devSvcClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *devSvcClient) Get(ctx context.Context, in *DeviceName, opts ...client.CallOption) (*Device, error) {
	req := c.c.NewRequest(c.serviceName, "DevSvc.Get", in)
	out := new(Device)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for DevSvc service

type DevSvcHandler interface {
	Get(context.Context, *DeviceName, *Device) error
}

func RegisterDevSvcHandler(s server.Server, hdlr DevSvcHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&DevSvc{hdlr}, opts...))
}

type DevSvc struct {
	DevSvcHandler
}

func (h *DevSvc) Get(ctx context.Context, in *DeviceName, out *Device) error {
	return h.DevSvcHandler.Get(ctx, in, out)
}

func init() { proto.RegisterFile("device.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 123 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x49, 0x49, 0x2d, 0xcb,
	0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0x52, 0xe0, 0xe2, 0x72, 0x01, 0xf3, 0xfd,
	0x12, 0x73, 0x53, 0x85, 0x84, 0xb8, 0x58, 0xf2, 0x12, 0x73, 0x53, 0x25, 0x18, 0x15, 0x18, 0x35,
	0x38, 0x83, 0xc0, 0x6c, 0x25, 0x13, 0x2e, 0x36, 0x88, 0x0a, 0x6c, 0xb2, 0x42, 0x62, 0x5c, 0x6c,
	0x89, 0xc9, 0x25, 0x99, 0x65, 0xa9, 0x12, 0x4c, 0x0a, 0x8c, 0x1a, 0x1c, 0x41, 0x50, 0x9e, 0x91,
	0x3a, 0x58, 0x57, 0x70, 0x59, 0xb2, 0x90, 0x2c, 0x17, 0xb3, 0x7b, 0x6a, 0x89, 0x10, 0xb7, 0x1e,
	0xc2, 0x1e, 0x29, 0x76, 0x28, 0x47, 0x89, 0x21, 0x89, 0x0d, 0xec, 0x0e, 0x63, 0x40, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x46, 0x18, 0x0e, 0x50, 0x97, 0x00, 0x00, 0x00,
}
