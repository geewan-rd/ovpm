// Code generated by protoc-gen-go. DO NOT EDIT.
// source: vpn.proto

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type VPNProto int32

const (
	VPNProto_NOPREF VPNProto = 0
	VPNProto_UDP    VPNProto = 1
	VPNProto_TCP    VPNProto = 2
)

var VPNProto_name = map[int32]string{
	0: "NOPREF",
	1: "UDP",
	2: "TCP",
}
var VPNProto_value = map[string]int32{
	"NOPREF": 0,
	"UDP":    1,
	"TCP":    2,
}

func (x VPNProto) String() string {
	return proto.EnumName(VPNProto_name, int32(x))
}
func (VPNProto) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

type VPNStatusRequest struct {
}

func (m *VPNStatusRequest) Reset()                    { *m = VPNStatusRequest{} }
func (m *VPNStatusRequest) String() string            { return proto.CompactTextString(m) }
func (*VPNStatusRequest) ProtoMessage()               {}
func (*VPNStatusRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

type VPNInitRequest struct {
	Hostname  string   `protobuf:"bytes,1,opt,name=hostname" json:"hostname,omitempty"`
	Port      string   `protobuf:"bytes,2,opt,name=port" json:"port,omitempty"`
	ProtoPref VPNProto `protobuf:"varint,3,opt,name=proto_pref,json=protoPref,enum=pb.VPNProto" json:"proto_pref,omitempty"`
	IpBlock   string   `protobuf:"bytes,4,opt,name=ip_block,json=ipBlock" json:"ip_block,omitempty"`
	Dns       string   `protobuf:"bytes,5,opt,name=dns" json:"dns,omitempty"`
}

func (m *VPNInitRequest) Reset()                    { *m = VPNInitRequest{} }
func (m *VPNInitRequest) String() string            { return proto.CompactTextString(m) }
func (*VPNInitRequest) ProtoMessage()               {}
func (*VPNInitRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *VPNInitRequest) GetHostname() string {
	if m != nil {
		return m.Hostname
	}
	return ""
}

func (m *VPNInitRequest) GetPort() string {
	if m != nil {
		return m.Port
	}
	return ""
}

func (m *VPNInitRequest) GetProtoPref() VPNProto {
	if m != nil {
		return m.ProtoPref
	}
	return VPNProto_NOPREF
}

func (m *VPNInitRequest) GetIpBlock() string {
	if m != nil {
		return m.IpBlock
	}
	return ""
}

func (m *VPNInitRequest) GetDns() string {
	if m != nil {
		return m.Dns
	}
	return ""
}

type VPNUpdateRequest struct {
	IpBlock string `protobuf:"bytes,1,opt,name=ip_block,json=ipBlock" json:"ip_block,omitempty"`
	Dns     string `protobuf:"bytes,2,opt,name=dns" json:"dns,omitempty"`
}

func (m *VPNUpdateRequest) Reset()                    { *m = VPNUpdateRequest{} }
func (m *VPNUpdateRequest) String() string            { return proto.CompactTextString(m) }
func (*VPNUpdateRequest) ProtoMessage()               {}
func (*VPNUpdateRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func (m *VPNUpdateRequest) GetIpBlock() string {
	if m != nil {
		return m.IpBlock
	}
	return ""
}

func (m *VPNUpdateRequest) GetDns() string {
	if m != nil {
		return m.Dns
	}
	return ""
}

type VPNStatusResponse struct {
	Name         string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	SerialNumber string `protobuf:"bytes,2,opt,name=serial_number,json=serialNumber" json:"serial_number,omitempty"`
	Hostname     string `protobuf:"bytes,3,opt,name=hostname" json:"hostname,omitempty"`
	Port         string `protobuf:"bytes,4,opt,name=port" json:"port,omitempty"`
	Cert         string `protobuf:"bytes,5,opt,name=cert" json:"cert,omitempty"`
	CaCert       string `protobuf:"bytes,6,opt,name=ca_cert,json=caCert" json:"ca_cert,omitempty"`
	Net          string `protobuf:"bytes,7,opt,name=net" json:"net,omitempty"`
	Mask         string `protobuf:"bytes,8,opt,name=mask" json:"mask,omitempty"`
	CreatedAt    string `protobuf:"bytes,9,opt,name=created_at,json=createdAt" json:"created_at,omitempty"`
	Proto        string `protobuf:"bytes,10,opt,name=proto" json:"proto,omitempty"`
	Dns          string `protobuf:"bytes,11,opt,name=dns" json:"dns,omitempty"`
}

func (m *VPNStatusResponse) Reset()                    { *m = VPNStatusResponse{} }
func (m *VPNStatusResponse) String() string            { return proto.CompactTextString(m) }
func (*VPNStatusResponse) ProtoMessage()               {}
func (*VPNStatusResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

func (m *VPNStatusResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *VPNStatusResponse) GetSerialNumber() string {
	if m != nil {
		return m.SerialNumber
	}
	return ""
}

func (m *VPNStatusResponse) GetHostname() string {
	if m != nil {
		return m.Hostname
	}
	return ""
}

func (m *VPNStatusResponse) GetPort() string {
	if m != nil {
		return m.Port
	}
	return ""
}

func (m *VPNStatusResponse) GetCert() string {
	if m != nil {
		return m.Cert
	}
	return ""
}

func (m *VPNStatusResponse) GetCaCert() string {
	if m != nil {
		return m.CaCert
	}
	return ""
}

func (m *VPNStatusResponse) GetNet() string {
	if m != nil {
		return m.Net
	}
	return ""
}

func (m *VPNStatusResponse) GetMask() string {
	if m != nil {
		return m.Mask
	}
	return ""
}

func (m *VPNStatusResponse) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *VPNStatusResponse) GetProto() string {
	if m != nil {
		return m.Proto
	}
	return ""
}

func (m *VPNStatusResponse) GetDns() string {
	if m != nil {
		return m.Dns
	}
	return ""
}

type VPNInitResponse struct {
}

func (m *VPNInitResponse) Reset()                    { *m = VPNInitResponse{} }
func (m *VPNInitResponse) String() string            { return proto.CompactTextString(m) }
func (*VPNInitResponse) ProtoMessage()               {}
func (*VPNInitResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{4} }

type VPNUpdateResponse struct {
}

func (m *VPNUpdateResponse) Reset()                    { *m = VPNUpdateResponse{} }
func (m *VPNUpdateResponse) String() string            { return proto.CompactTextString(m) }
func (*VPNUpdateResponse) ProtoMessage()               {}
func (*VPNUpdateResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{5} }

func init() {
	proto.RegisterType((*VPNStatusRequest)(nil), "pb.VPNStatusRequest")
	proto.RegisterType((*VPNInitRequest)(nil), "pb.VPNInitRequest")
	proto.RegisterType((*VPNUpdateRequest)(nil), "pb.VPNUpdateRequest")
	proto.RegisterType((*VPNStatusResponse)(nil), "pb.VPNStatusResponse")
	proto.RegisterType((*VPNInitResponse)(nil), "pb.VPNInitResponse")
	proto.RegisterType((*VPNUpdateResponse)(nil), "pb.VPNUpdateResponse")
	proto.RegisterEnum("pb.VPNProto", VPNProto_name, VPNProto_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for VPNService service

type VPNServiceClient interface {
	Status(ctx context.Context, in *VPNStatusRequest, opts ...grpc.CallOption) (*VPNStatusResponse, error)
	Init(ctx context.Context, in *VPNInitRequest, opts ...grpc.CallOption) (*VPNInitResponse, error)
	Update(ctx context.Context, in *VPNUpdateRequest, opts ...grpc.CallOption) (*VPNUpdateResponse, error)
}

type vPNServiceClient struct {
	cc *grpc.ClientConn
}

func NewVPNServiceClient(cc *grpc.ClientConn) VPNServiceClient {
	return &vPNServiceClient{cc}
}

func (c *vPNServiceClient) Status(ctx context.Context, in *VPNStatusRequest, opts ...grpc.CallOption) (*VPNStatusResponse, error) {
	out := new(VPNStatusResponse)
	err := grpc.Invoke(ctx, "/pb.VPNService/Status", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vPNServiceClient) Init(ctx context.Context, in *VPNInitRequest, opts ...grpc.CallOption) (*VPNInitResponse, error) {
	out := new(VPNInitResponse)
	err := grpc.Invoke(ctx, "/pb.VPNService/Init", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vPNServiceClient) Update(ctx context.Context, in *VPNUpdateRequest, opts ...grpc.CallOption) (*VPNUpdateResponse, error) {
	out := new(VPNUpdateResponse)
	err := grpc.Invoke(ctx, "/pb.VPNService/Update", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for VPNService service

type VPNServiceServer interface {
	Status(context.Context, *VPNStatusRequest) (*VPNStatusResponse, error)
	Init(context.Context, *VPNInitRequest) (*VPNInitResponse, error)
	Update(context.Context, *VPNUpdateRequest) (*VPNUpdateResponse, error)
}

func RegisterVPNServiceServer(s *grpc.Server, srv VPNServiceServer) {
	s.RegisterService(&_VPNService_serviceDesc, srv)
}

func _VPNService_Status_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VPNStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VPNServiceServer).Status(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.VPNService/Status",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VPNServiceServer).Status(ctx, req.(*VPNStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VPNService_Init_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VPNInitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VPNServiceServer).Init(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.VPNService/Init",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VPNServiceServer).Init(ctx, req.(*VPNInitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VPNService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VPNUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VPNServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.VPNService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VPNServiceServer).Update(ctx, req.(*VPNUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _VPNService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.VPNService",
	HandlerType: (*VPNServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Status",
			Handler:    _VPNService_Status_Handler,
		},
		{
			MethodName: "Init",
			Handler:    _VPNService_Init_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _VPNService_Update_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "vpn.proto",
}

func init() { proto.RegisterFile("vpn.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 502 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x86, 0xb1, 0x93, 0x3a, 0xc9, 0x50, 0x8a, 0x3b, 0x4d, 0x55, 0x13, 0xa8, 0x54, 0x99, 0x4b,
	0x14, 0xa4, 0x44, 0x94, 0x5b, 0x2f, 0x08, 0x0a, 0x48, 0x48, 0xc8, 0x98, 0xd0, 0xe6, 0x6a, 0xad,
	0x9d, 0x6d, 0xb1, 0x9a, 0xec, 0x2e, 0xde, 0x4d, 0x1e, 0x00, 0x1e, 0x01, 0x89, 0x03, 0xaf, 0xc5,
	0x2b, 0xf0, 0x20, 0xc8, 0xb3, 0x4e, 0x6a, 0x23, 0x7a, 0x9b, 0xfd, 0x67, 0xf7, 0xd7, 0x3f, 0xdf,
	0x0e, 0xf4, 0xd6, 0x4a, 0x8c, 0x55, 0x21, 0x8d, 0x44, 0x57, 0xa5, 0x83, 0x27, 0xd7, 0x52, 0x5e,
	0x2f, 0xf8, 0x84, 0xa9, 0x7c, 0xc2, 0x84, 0x90, 0x86, 0x99, 0x5c, 0x0a, 0x6d, 0x6f, 0x84, 0x08,
	0xfe, 0x2c, 0x8e, 0x3e, 0x1b, 0x66, 0x56, 0x7a, 0xca, 0xbf, 0xae, 0xb8, 0x36, 0xe1, 0x2f, 0x07,
	0xf6, 0x66, 0x71, 0xf4, 0x5e, 0xe4, 0xa6, 0x92, 0x70, 0x00, 0xdd, 0x2f, 0x52, 0x1b, 0xc1, 0x96,
	0x3c, 0x70, 0x4e, 0x9c, 0x61, 0x6f, 0xba, 0x3d, 0x23, 0x42, 0x5b, 0xc9, 0xc2, 0x04, 0x2e, 0xe9,
	0x54, 0xe3, 0x33, 0x00, 0xf2, 0x4f, 0x54, 0xc1, 0xaf, 0x82, 0xd6, 0x89, 0x33, 0xdc, 0x3b, 0xdd,
	0x1d, 0xab, 0x74, 0x3c, 0x8b, 0xa3, 0xb8, 0x6c, 0x4c, 0x7b, 0xd4, 0x8f, 0x0b, 0x7e, 0x85, 0x8f,
	0xa0, 0x9b, 0xab, 0x24, 0x5d, 0xc8, 0xec, 0x26, 0x68, 0x93, 0x49, 0x27, 0x57, 0xaf, 0xcb, 0x23,
	0xfa, 0xd0, 0x9a, 0x0b, 0x1d, 0xec, 0x90, 0x5a, 0x96, 0xe1, 0x4b, 0x0a, 0x7c, 0xa9, 0xe6, 0xcc,
	0xf0, 0x4d, 0xba, 0xba, 0x81, 0xf3, 0x5f, 0x03, 0xf7, 0xd6, 0xe0, 0xa7, 0x0b, 0xfb, 0xb5, 0x91,
	0xb5, 0x92, 0x42, 0xd3, 0x10, 0xb5, 0xe1, 0xa8, 0xc6, 0xa7, 0xf0, 0x40, 0xf3, 0x22, 0x67, 0x8b,
	0x44, 0xac, 0x96, 0x29, 0x2f, 0x2a, 0x97, 0x5d, 0x2b, 0x46, 0xa4, 0x35, 0xc8, 0xb4, 0xee, 0x20,
	0xd3, 0xae, 0x91, 0x41, 0x68, 0x67, 0xbc, 0x30, 0xd5, 0x48, 0x54, 0xe3, 0x11, 0x74, 0x32, 0x96,
	0x90, 0xec, 0x91, 0xec, 0x65, 0xec, 0xbc, 0x6c, 0xf8, 0xd0, 0x12, 0xdc, 0x04, 0x1d, 0x9b, 0x5e,
	0x70, 0x7a, 0xbe, 0x64, 0xfa, 0x26, 0xe8, 0xda, 0xe7, 0x65, 0x8d, 0xc7, 0x00, 0x59, 0xc1, 0x99,
	0xe1, 0xf3, 0x84, 0x99, 0xa0, 0x47, 0x9d, 0x5e, 0xa5, 0xbc, 0x32, 0xd8, 0x87, 0x1d, 0x62, 0x1d,
	0x00, 0x75, 0xec, 0x61, 0x03, 0xe6, 0xfe, 0x2d, 0x98, 0x7d, 0x78, 0xb8, 0xfd, 0x75, 0x4b, 0x25,
	0x3c, 0x20, 0x54, 0x1b, 0xd8, 0x56, 0x1c, 0x0d, 0xa1, 0xbb, 0xf9, 0x45, 0x04, 0xf0, 0xa2, 0x8f,
	0xf1, 0xf4, 0xed, 0x3b, 0xff, 0x1e, 0x76, 0xa0, 0x75, 0xf9, 0x26, 0xf6, 0x9d, 0xb2, 0xb8, 0x38,
	0x8f, 0x7d, 0xf7, 0xf4, 0xbb, 0x0b, 0x50, 0xa2, 0xe6, 0xc5, 0x3a, 0xcf, 0x38, 0x7e, 0x02, 0xcf,
	0x52, 0xc7, 0x7e, 0xb5, 0x0a, 0x8d, 0xbd, 0x1b, 0x1c, 0xfe, 0xa3, 0x56, 0x21, 0x06, 0xdf, 0x7e,
	0xff, 0xf9, 0xe1, 0xf6, 0x11, 0x69, 0x85, 0xd7, 0xcf, 0x27, 0x6b, 0x25, 0x26, 0xda, 0x1a, 0x7d,
	0x80, 0x76, 0x19, 0x18, 0xb1, 0x7a, 0x5a, 0xdb, 0xd9, 0xc1, 0x41, 0x43, 0xab, 0xcc, 0x1e, 0x93,
	0xd9, 0x61, 0xe8, 0xd7, 0xcd, 0x72, 0x91, 0x9b, 0x33, 0x67, 0x84, 0x17, 0xe0, 0xd9, 0x59, 0xb7,
	0x01, 0x1b, 0x7b, 0xb6, 0x0d, 0xd8, 0x04, 0x12, 0x1e, 0x93, 0xe7, 0x51, 0xd8, 0x08, 0xb8, 0xa2,
	0x3b, 0x67, 0xce, 0x28, 0xf5, 0x08, 0xf8, 0x8b, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xb8, 0x57,
	0x61, 0xfa, 0x98, 0x03, 0x00, 0x00,
}
