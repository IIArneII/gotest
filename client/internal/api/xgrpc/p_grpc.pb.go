// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.5
// source: internal/api/xgrpc/p.proto

package xgrpc

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// TestServiceClient is the client API for TestService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TestServiceClient interface {
	VerifyClient(ctx context.Context, in *Verify, opts ...grpc.CallOption) (*VerifyAnswer, error)
	SendMessage(ctx context.Context, in *Message, opts ...grpc.CallOption) (*MessageAnswer, error)
	SendMessageToOtherClient(ctx context.Context, in *MessageToOther, opts ...grpc.CallOption) (*MessageToOtherAnswer, error)
}

type testServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTestServiceClient(cc grpc.ClientConnInterface) TestServiceClient {
	return &testServiceClient{cc}
}

func (c *testServiceClient) VerifyClient(ctx context.Context, in *Verify, opts ...grpc.CallOption) (*VerifyAnswer, error) {
	out := new(VerifyAnswer)
	err := c.cc.Invoke(ctx, "/xgrpc.TestService/VerifyClient", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testServiceClient) SendMessage(ctx context.Context, in *Message, opts ...grpc.CallOption) (*MessageAnswer, error) {
	out := new(MessageAnswer)
	err := c.cc.Invoke(ctx, "/xgrpc.TestService/SendMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testServiceClient) SendMessageToOtherClient(ctx context.Context, in *MessageToOther, opts ...grpc.CallOption) (*MessageToOtherAnswer, error) {
	out := new(MessageToOtherAnswer)
	err := c.cc.Invoke(ctx, "/xgrpc.TestService/SendMessageToOtherClient", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TestServiceServer is the server API for TestService service.
// All implementations must embed UnimplementedTestServiceServer
// for forward compatibility
type TestServiceServer interface {
	VerifyClient(context.Context, *Verify) (*VerifyAnswer, error)
	SendMessage(context.Context, *Message) (*MessageAnswer, error)
	SendMessageToOtherClient(context.Context, *MessageToOther) (*MessageToOtherAnswer, error)
	mustEmbedUnimplementedTestServiceServer()
}

// UnimplementedTestServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTestServiceServer struct {
}

func (UnimplementedTestServiceServer) VerifyClient(context.Context, *Verify) (*VerifyAnswer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyClient not implemented")
}
func (UnimplementedTestServiceServer) SendMessage(context.Context, *Message) (*MessageAnswer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMessage not implemented")
}
func (UnimplementedTestServiceServer) SendMessageToOtherClient(context.Context, *MessageToOther) (*MessageToOtherAnswer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMessageToOtherClient not implemented")
}
func (UnimplementedTestServiceServer) mustEmbedUnimplementedTestServiceServer() {}

// UnsafeTestServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TestServiceServer will
// result in compilation errors.
type UnsafeTestServiceServer interface {
	mustEmbedUnimplementedTestServiceServer()
}

func RegisterTestServiceServer(s grpc.ServiceRegistrar, srv TestServiceServer) {
	s.RegisterService(&TestService_ServiceDesc, srv)
}

func _TestService_VerifyClient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Verify)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServiceServer).VerifyClient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/xgrpc.TestService/VerifyClient",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServiceServer).VerifyClient(ctx, req.(*Verify))
	}
	return interceptor(ctx, in, info, handler)
}

func _TestService_SendMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServiceServer).SendMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/xgrpc.TestService/SendMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServiceServer).SendMessage(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

func _TestService_SendMessageToOtherClient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MessageToOther)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServiceServer).SendMessageToOtherClient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/xgrpc.TestService/SendMessageToOtherClient",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServiceServer).SendMessageToOtherClient(ctx, req.(*MessageToOther))
	}
	return interceptor(ctx, in, info, handler)
}

// TestService_ServiceDesc is the grpc.ServiceDesc for TestService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TestService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "xgrpc.TestService",
	HandlerType: (*TestServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "VerifyClient",
			Handler:    _TestService_VerifyClient_Handler,
		},
		{
			MethodName: "SendMessage",
			Handler:    _TestService_SendMessage_Handler,
		},
		{
			MethodName: "SendMessageToOtherClient",
			Handler:    _TestService_SendMessageToOtherClient_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "internal/api/xgrpc/p.proto",
}
