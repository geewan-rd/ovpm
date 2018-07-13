// Code generated by protoc-gen-go. DO NOT EDIT.
// source: auth.proto

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

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type AuthStatusRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthStatusRequest) Reset()         { *m = AuthStatusRequest{} }
func (m *AuthStatusRequest) String() string { return proto.CompactTextString(m) }
func (*AuthStatusRequest) ProtoMessage()    {}
func (*AuthStatusRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_auth_79ec97a74d2b7cf3, []int{0}
}
func (m *AuthStatusRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthStatusRequest.Unmarshal(m, b)
}
func (m *AuthStatusRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthStatusRequest.Marshal(b, m, deterministic)
}
func (dst *AuthStatusRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthStatusRequest.Merge(dst, src)
}
func (m *AuthStatusRequest) XXX_Size() int {
	return xxx_messageInfo_AuthStatusRequest.Size(m)
}
func (m *AuthStatusRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthStatusRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AuthStatusRequest proto.InternalMessageInfo

type AuthAuthenticateRequest struct {
	Username             string   `protobuf:"bytes,1,opt,name=username" json:"username,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthAuthenticateRequest) Reset()         { *m = AuthAuthenticateRequest{} }
func (m *AuthAuthenticateRequest) String() string { return proto.CompactTextString(m) }
func (*AuthAuthenticateRequest) ProtoMessage()    {}
func (*AuthAuthenticateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_auth_79ec97a74d2b7cf3, []int{1}
}
func (m *AuthAuthenticateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthAuthenticateRequest.Unmarshal(m, b)
}
func (m *AuthAuthenticateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthAuthenticateRequest.Marshal(b, m, deterministic)
}
func (dst *AuthAuthenticateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthAuthenticateRequest.Merge(dst, src)
}
func (m *AuthAuthenticateRequest) XXX_Size() int {
	return xxx_messageInfo_AuthAuthenticateRequest.Size(m)
}
func (m *AuthAuthenticateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthAuthenticateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AuthAuthenticateRequest proto.InternalMessageInfo

func (m *AuthAuthenticateRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *AuthAuthenticateRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type AuthStatusResponse struct {
	User                 *UserResponse_User `protobuf:"bytes,1,opt,name=user" json:"user,omitempty"`
	IsRoot               bool               `protobuf:"varint,2,opt,name=is_root,json=isRoot" json:"is_root,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *AuthStatusResponse) Reset()         { *m = AuthStatusResponse{} }
func (m *AuthStatusResponse) String() string { return proto.CompactTextString(m) }
func (*AuthStatusResponse) ProtoMessage()    {}
func (*AuthStatusResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_auth_79ec97a74d2b7cf3, []int{2}
}
func (m *AuthStatusResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthStatusResponse.Unmarshal(m, b)
}
func (m *AuthStatusResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthStatusResponse.Marshal(b, m, deterministic)
}
func (dst *AuthStatusResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthStatusResponse.Merge(dst, src)
}
func (m *AuthStatusResponse) XXX_Size() int {
	return xxx_messageInfo_AuthStatusResponse.Size(m)
}
func (m *AuthStatusResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthStatusResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AuthStatusResponse proto.InternalMessageInfo

func (m *AuthStatusResponse) GetUser() *UserResponse_User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *AuthStatusResponse) GetIsRoot() bool {
	if m != nil {
		return m.IsRoot
	}
	return false
}

type AuthAuthenticateResponse struct {
	Token                string   `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthAuthenticateResponse) Reset()         { *m = AuthAuthenticateResponse{} }
func (m *AuthAuthenticateResponse) String() string { return proto.CompactTextString(m) }
func (*AuthAuthenticateResponse) ProtoMessage()    {}
func (*AuthAuthenticateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_auth_79ec97a74d2b7cf3, []int{3}
}
func (m *AuthAuthenticateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthAuthenticateResponse.Unmarshal(m, b)
}
func (m *AuthAuthenticateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthAuthenticateResponse.Marshal(b, m, deterministic)
}
func (dst *AuthAuthenticateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthAuthenticateResponse.Merge(dst, src)
}
func (m *AuthAuthenticateResponse) XXX_Size() int {
	return xxx_messageInfo_AuthAuthenticateResponse.Size(m)
}
func (m *AuthAuthenticateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthAuthenticateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AuthAuthenticateResponse proto.InternalMessageInfo

func (m *AuthAuthenticateResponse) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func init() {
	proto.RegisterType((*AuthStatusRequest)(nil), "pb.AuthStatusRequest")
	proto.RegisterType((*AuthAuthenticateRequest)(nil), "pb.AuthAuthenticateRequest")
	proto.RegisterType((*AuthStatusResponse)(nil), "pb.AuthStatusResponse")
	proto.RegisterType((*AuthAuthenticateResponse)(nil), "pb.AuthAuthenticateResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AuthServiceClient is the client API for AuthService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuthServiceClient interface {
	Status(ctx context.Context, in *AuthStatusRequest, opts ...grpc.CallOption) (*AuthStatusResponse, error)
	Authenticate(ctx context.Context, in *AuthAuthenticateRequest, opts ...grpc.CallOption) (*AuthAuthenticateResponse, error)
}

type authServiceClient struct {
	cc *grpc.ClientConn
}

func NewAuthServiceClient(cc *grpc.ClientConn) AuthServiceClient {
	return &authServiceClient{cc}
}

func (c *authServiceClient) Status(ctx context.Context, in *AuthStatusRequest, opts ...grpc.CallOption) (*AuthStatusResponse, error) {
	out := new(AuthStatusResponse)
	err := c.cc.Invoke(ctx, "/pb.AuthService/Status", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) Authenticate(ctx context.Context, in *AuthAuthenticateRequest, opts ...grpc.CallOption) (*AuthAuthenticateResponse, error) {
	out := new(AuthAuthenticateResponse)
	err := c.cc.Invoke(ctx, "/pb.AuthService/Authenticate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServiceServer is the server API for AuthService service.
type AuthServiceServer interface {
	Status(context.Context, *AuthStatusRequest) (*AuthStatusResponse, error)
	Authenticate(context.Context, *AuthAuthenticateRequest) (*AuthAuthenticateResponse, error)
}

func RegisterAuthServiceServer(s *grpc.Server, srv AuthServiceServer) {
	s.RegisterService(&_AuthService_serviceDesc, srv)
}

func _AuthService_Status_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).Status(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.AuthService/Status",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).Status(ctx, req.(*AuthStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_Authenticate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthAuthenticateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).Authenticate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.AuthService/Authenticate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).Authenticate(ctx, req.(*AuthAuthenticateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AuthService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.AuthService",
	HandlerType: (*AuthServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Status",
			Handler:    _AuthService_Status_Handler,
		},
		{
			MethodName: "Authenticate",
			Handler:    _AuthService_Authenticate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth.proto",
}

func init() { proto.RegisterFile("auth.proto", fileDescriptor_auth_79ec97a74d2b7cf3) }

var fileDescriptor_auth_79ec97a74d2b7cf3 = []byte{
	// 306 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x51, 0x41, 0x4e, 0x2a, 0x41,
	0x10, 0xcd, 0x90, 0xff, 0x11, 0x0b, 0x37, 0x36, 0x22, 0x38, 0x10, 0x63, 0x26, 0x2e, 0xd4, 0x05,
	0xa3, 0xb8, 0x73, 0xe7, 0x11, 0x1c, 0x35, 0x71, 0x67, 0x1a, 0xac, 0xc0, 0x44, 0xed, 0x1a, 0xbb,
	0x6a, 0x70, 0xef, 0x15, 0x3c, 0x9a, 0x1e, 0xc1, 0x83, 0x98, 0xee, 0x1e, 0xcc, 0x28, 0x2c, 0x3a,
	0xe9, 0x57, 0xaf, 0xea, 0xbd, 0x97, 0x2a, 0x00, 0x5d, 0xca, 0x7c, 0x54, 0x58, 0x12, 0x52, 0x8d,
	0x62, 0x12, 0x0f, 0x67, 0x44, 0xb3, 0x27, 0x4c, 0x75, 0x91, 0xa7, 0xda, 0x18, 0x12, 0x2d, 0x39,
	0x19, 0x0e, 0x1d, 0x31, 0x94, 0x8c, 0x36, 0xfc, 0x93, 0x0e, 0x6c, 0x5f, 0x96, 0x32, 0xbf, 0x16,
	0x2d, 0x25, 0x67, 0xf8, 0x52, 0x22, 0x4b, 0x72, 0x05, 0x3d, 0x57, 0x74, 0x0f, 0x8d, 0xe4, 0x53,
	0x2d, 0x58, 0x51, 0x2a, 0x86, 0x96, 0x9b, 0x36, 0xfa, 0x19, 0xfb, 0xd1, 0x41, 0x74, 0xb4, 0x99,
	0xfd, 0x60, 0xc7, 0x15, 0x9a, 0xf9, 0x95, 0xec, 0x43, 0xbf, 0x11, 0xb8, 0x25, 0x4e, 0xee, 0x40,
	0xd5, 0x7d, 0xb8, 0x20, 0xc3, 0xa8, 0x8e, 0xe1, 0x9f, 0x9b, 0xf6, 0x4a, 0xed, 0x71, 0x77, 0x54,
	0x4c, 0x46, 0xb7, 0x8c, 0x76, 0xc9, 0x07, 0xe0, 0x5b, 0x54, 0x0f, 0x36, 0x72, 0xbe, 0xb7, 0x44,
	0xe2, 0xb5, 0x5b, 0x59, 0x33, 0xe7, 0x8c, 0x48, 0x92, 0x53, 0xe8, 0xaf, 0x86, 0xad, 0xf4, 0x77,
	0xe0, 0xbf, 0xd0, 0x23, 0x9a, 0x2a, 0x6a, 0x00, 0xe3, 0xcf, 0x08, 0xda, 0x3e, 0x0c, 0xda, 0x45,
	0x3e, 0x45, 0x75, 0x03, 0xcd, 0x90, 0x4b, 0xf9, 0x04, 0x2b, 0xfb, 0x88, 0x77, 0xff, 0x96, 0x83,
	0x7c, 0x32, 0x78, 0xfb, 0xf8, 0x7a, 0x6f, 0x74, 0x55, 0xc7, 0x2f, 0x7a, 0x71, 0x96, 0xba, 0x33,
	0xa4, 0x1c, 0xb4, 0x08, 0xb6, 0xea, 0x99, 0xd4, 0x60, 0x29, 0xb2, 0x66, 0xad, 0xf1, 0x70, 0x3d,
	0x59, 0xf9, 0x1c, 0x7a, 0x9f, 0xfd, 0x64, 0xef, 0x97, 0x8f, 0xae, 0xb5, 0x5e, 0x44, 0x27, 0x93,
	0xa6, 0xbf, 0xe8, 0xf9, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0x87, 0xd8, 0x13, 0x0a, 0x0d, 0x02,
	0x00, 0x00,
}
