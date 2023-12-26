// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: proto/demo_grpc.proto

package demo_proto

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

// DemoServiceClient is the client API for DemoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DemoServiceClient interface {
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error)
	StringToChar(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (DemoService_StringToCharClient, error)
	Adder(ctx context.Context, in *AdderRequest, opts ...grpc.CallOption) (*AdderResponse, error)
}

type demoServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDemoServiceClient(cc grpc.ClientConnInterface) DemoServiceClient {
	return &demoServiceClient{cc}
}

func (c *demoServiceClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error) {
	out := new(HelloResponse)
	err := c.cc.Invoke(ctx, "/demo_proto.DemoService/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *demoServiceClient) StringToChar(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (DemoService_StringToCharClient, error) {
	stream, err := c.cc.NewStream(ctx, &DemoService_ServiceDesc.Streams[0], "/demo_proto.DemoService/StringToChar", opts...)
	if err != nil {
		return nil, err
	}
	x := &demoServiceStringToCharClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type DemoService_StringToCharClient interface {
	Recv() (*CharResponse, error)
	grpc.ClientStream
}

type demoServiceStringToCharClient struct {
	grpc.ClientStream
}

func (x *demoServiceStringToCharClient) Recv() (*CharResponse, error) {
	m := new(CharResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *demoServiceClient) Adder(ctx context.Context, in *AdderRequest, opts ...grpc.CallOption) (*AdderResponse, error) {
	out := new(AdderResponse)
	err := c.cc.Invoke(ctx, "/demo_proto.DemoService/Adder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DemoServiceServer is the server API for DemoService service.
// All implementations must embed UnimplementedDemoServiceServer
// for forward compatibility
type DemoServiceServer interface {
	SayHello(context.Context, *HelloRequest) (*HelloResponse, error)
	StringToChar(*HelloRequest, DemoService_StringToCharServer) error
	Adder(context.Context, *AdderRequest) (*AdderResponse, error)
	mustEmbedUnimplementedDemoServiceServer()
}

// UnimplementedDemoServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDemoServiceServer struct {
}

func (UnimplementedDemoServiceServer) SayHello(context.Context, *HelloRequest) (*HelloResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (UnimplementedDemoServiceServer) StringToChar(*HelloRequest, DemoService_StringToCharServer) error {
	return status.Errorf(codes.Unimplemented, "method StringToChar not implemented")
}
func (UnimplementedDemoServiceServer) Adder(context.Context, *AdderRequest) (*AdderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Adder not implemented")
}
func (UnimplementedDemoServiceServer) mustEmbedUnimplementedDemoServiceServer() {}

// UnsafeDemoServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DemoServiceServer will
// result in compilation errors.
type UnsafeDemoServiceServer interface {
	mustEmbedUnimplementedDemoServiceServer()
}

func RegisterDemoServiceServer(s grpc.ServiceRegistrar, srv DemoServiceServer) {
	s.RegisterService(&DemoService_ServiceDesc, srv)
}

func _DemoService_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DemoServiceServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/demo_proto.DemoService/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DemoServiceServer).SayHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DemoService_StringToChar_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(HelloRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DemoServiceServer).StringToChar(m, &demoServiceStringToCharServer{stream})
}

type DemoService_StringToCharServer interface {
	Send(*CharResponse) error
	grpc.ServerStream
}

type demoServiceStringToCharServer struct {
	grpc.ServerStream
}

func (x *demoServiceStringToCharServer) Send(m *CharResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _DemoService_Adder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DemoServiceServer).Adder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/demo_proto.DemoService/Adder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DemoServiceServer).Adder(ctx, req.(*AdderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DemoService_ServiceDesc is the grpc.ServiceDesc for DemoService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DemoService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "demo_proto.DemoService",
	HandlerType: (*DemoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _DemoService_SayHello_Handler,
		},
		{
			MethodName: "Adder",
			Handler:    _DemoService_Adder_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StringToChar",
			Handler:       _DemoService_StringToChar_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/demo_grpc.proto",
}
