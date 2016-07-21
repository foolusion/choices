// Code generated by protoc-gen-go.
// source: elwin.proto
// DO NOT EDIT!

/*
Package elwin is a generated protocol buffer package.

It is generated from these files:
	elwin.proto

It has these top-level messages:
	Identifier
	Experiments
	Experiment
	Param
*/
package elwin

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

type Identifier struct {
	TeamID string `protobuf:"bytes,1,opt,name=teamID" json:"teamID,omitempty"`
	UserID string `protobuf:"bytes,2,opt,name=userID" json:"userID,omitempty"`
}

func (m *Identifier) Reset()                    { *m = Identifier{} }
func (m *Identifier) String() string            { return proto.CompactTextString(m) }
func (*Identifier) ProtoMessage()               {}
func (*Identifier) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Experiments struct {
	Experiments map[string]*Experiment `protobuf:"bytes,1,rep,name=experiments" json:"experiments,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *Experiments) Reset()                    { *m = Experiments{} }
func (m *Experiments) String() string            { return proto.CompactTextString(m) }
func (*Experiments) ProtoMessage()               {}
func (*Experiments) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Experiments) GetExperiments() map[string]*Experiment {
	if m != nil {
		return m.Experiments
	}
	return nil
}

type Experiment struct {
	Name   string   `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Params []*Param `protobuf:"bytes,2,rep,name=params" json:"params,omitempty"`
}

func (m *Experiment) Reset()                    { *m = Experiment{} }
func (m *Experiment) String() string            { return proto.CompactTextString(m) }
func (*Experiment) ProtoMessage()               {}
func (*Experiment) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Experiment) GetParams() []*Param {
	if m != nil {
		return m.Params
	}
	return nil
}

type Param struct {
	Key   string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
}

func (m *Param) Reset()                    { *m = Param{} }
func (m *Param) String() string            { return proto.CompactTextString(m) }
func (*Param) ProtoMessage()               {}
func (*Param) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func init() {
	proto.RegisterType((*Identifier)(nil), "elwin.Identifier")
	proto.RegisterType((*Experiments)(nil), "elwin.Experiments")
	proto.RegisterType((*Experiment)(nil), "elwin.Experiment")
	proto.RegisterType((*Param)(nil), "elwin.Param")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for Elwin service

type ElwinClient interface {
	GetNamespaces(ctx context.Context, in *Identifier, opts ...grpc.CallOption) (*Experiments, error)
}

type elwinClient struct {
	cc *grpc.ClientConn
}

func NewElwinClient(cc *grpc.ClientConn) ElwinClient {
	return &elwinClient{cc}
}

func (c *elwinClient) GetNamespaces(ctx context.Context, in *Identifier, opts ...grpc.CallOption) (*Experiments, error) {
	out := new(Experiments)
	err := grpc.Invoke(ctx, "/elwin.Elwin/GetNamespaces", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Elwin service

type ElwinServer interface {
	GetNamespaces(context.Context, *Identifier) (*Experiments, error)
}

func RegisterElwinServer(s *grpc.Server, srv ElwinServer) {
	s.RegisterService(&_Elwin_serviceDesc, srv)
}

func _Elwin_GetNamespaces_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Identifier)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ElwinServer).GetNamespaces(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/elwin.Elwin/GetNamespaces",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ElwinServer).GetNamespaces(ctx, req.(*Identifier))
	}
	return interceptor(ctx, in, info, handler)
}

var _Elwin_serviceDesc = grpc.ServiceDesc{
	ServiceName: "elwin.Elwin",
	HandlerType: (*ElwinServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetNamespaces",
			Handler:    _Elwin_GetNamespaces_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor0,
}

func init() { proto.RegisterFile("elwin.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 262 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x51, 0x4d, 0x4b, 0x03, 0x31,
	0x10, 0x75, 0x5b, 0x77, 0xa1, 0xb3, 0x0a, 0x75, 0x10, 0x91, 0x9e, 0x24, 0x0a, 0xf6, 0x54, 0x61,
	0xbd, 0x14, 0xf1, 0x22, 0xb8, 0x4a, 0x2f, 0xa2, 0xfb, 0x0f, 0xa2, 0x8e, 0xb0, 0xd8, 0x4d, 0x43,
	0x92, 0xaa, 0xfd, 0x45, 0xfe, 0x4d, 0xf3, 0x25, 0x1b, 0xac, 0xb7, 0x79, 0xef, 0x4d, 0xde, 0x7b,
	0x4c, 0xa0, 0xa4, 0xe5, 0x67, 0x2b, 0x66, 0x52, 0xad, 0xcc, 0x0a, 0x73, 0x0f, 0xd8, 0x35, 0xc0,
	0xe2, 0x95, 0x84, 0x69, 0xdf, 0x5a, 0x52, 0x78, 0x04, 0x85, 0x21, 0xde, 0x2d, 0x6e, 0x8f, 0xb3,
	0x93, 0x6c, 0x3a, 0x6a, 0x22, 0x72, 0xfc, 0x5a, 0x93, 0xb2, 0xfc, 0x20, 0xf0, 0x01, 0xb1, 0xef,
	0x0c, 0xca, 0xfa, 0x4b, 0x92, 0x6a, 0x3b, 0xeb, 0xa1, 0xb1, 0xb6, 0x19, 0x3d, 0xb4, 0x26, 0xc3,
	0x69, 0x59, 0x9d, 0xce, 0x42, 0x6e, 0xb2, 0x98, 0xce, 0xb5, 0x30, 0x6a, 0xd3, 0xa4, 0xef, 0x26,
	0x4f, 0x30, 0xfe, 0xbb, 0x80, 0x63, 0x18, 0xbe, 0xd3, 0x26, 0xf6, 0x72, 0x23, 0x9e, 0x43, 0xfe,
	0xc1, 0x97, 0x6b, 0xf2, 0x9d, 0xca, 0xea, 0x60, 0x2b, 0xa6, 0x09, 0xfa, 0xd5, 0x60, 0x9e, 0xb1,
	0x3b, 0x80, 0x5e, 0x40, 0x84, 0x5d, 0xc1, 0x3b, 0x8a, 0x6e, 0x7e, 0xc6, 0x33, 0x28, 0x24, 0x57,
	0xbc, 0xd3, 0xd6, 0xcf, 0xd5, 0xde, 0x8b, 0x7e, 0x8f, 0x8e, 0x6c, 0xa2, 0xc6, 0x2e, 0x20, 0xf7,
	0xc4, 0x3f, 0x7d, 0x0e, 0xd3, 0x3e, 0xa3, 0x18, 0x5e, 0xdd, 0x40, 0x5e, 0x3b, 0x1f, 0x9c, 0xc3,
	0xfe, 0x3d, 0x99, 0x07, 0x1b, 0xa5, 0x25, 0x7f, 0x21, 0x8d, 0xbf, 0x85, 0xfb, 0xfb, 0x4f, 0x70,
	0xfb, 0x54, 0x6c, 0xe7, 0xb9, 0xf0, 0x3f, 0x76, 0xf9, 0x13, 0x00, 0x00, 0xff, 0xff, 0x4e, 0x95,
	0xe3, 0xfa, 0xc0, 0x01, 0x00, 0x00,
}
