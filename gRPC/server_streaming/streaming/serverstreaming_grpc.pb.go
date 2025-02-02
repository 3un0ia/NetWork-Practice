// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.20.3
// source: serverstreaming.proto

package streaming

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	ServerStreaming_GetServerResponse_FullMethodName = "/serverstreaming.ServerStreaming/GetServerResponse"
)

// ServerStreamingClient is the client API for ServerStreaming service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServerStreamingClient interface {
	// A Client streaming RPC.
	GetServerResponse(ctx context.Context, in *Number, opts ...grpc.CallOption) (grpc.ServerStreamingClient[Message], error)
}

type serverStreamingClient struct {
	cc grpc.ClientConnInterface
}

func NewServerStreamingClient(cc grpc.ClientConnInterface) ServerStreamingClient {
	return &serverStreamingClient{cc}
}

func (c *serverStreamingClient) GetServerResponse(ctx context.Context, in *Number, opts ...grpc.CallOption) (grpc.ServerStreamingClient[Message], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &ServerStreaming_ServiceDesc.Streams[0], ServerStreaming_GetServerResponse_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[Number, Message]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type ServerStreaming_GetServerResponseClient = grpc.ServerStreamingClient[Message]

// ServerStreamingServer is the server API for ServerStreaming service.
// All implementations must embed UnimplementedServerStreamingServer
// for forward compatibility.
type ServerStreamingServer interface {
	// A Client streaming RPC.
	GetServerResponse(*Number, grpc.ServerStreamingServer[Message]) error
	mustEmbedUnimplementedServerStreamingServer()
}

// UnimplementedServerStreamingServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedServerStreamingServer struct{}

func (UnimplementedServerStreamingServer) GetServerResponse(*Number, grpc.ServerStreamingServer[Message]) error {
	return status.Errorf(codes.Unimplemented, "method GetServerResponse not implemented")
}
func (UnimplementedServerStreamingServer) mustEmbedUnimplementedServerStreamingServer() {}
func (UnimplementedServerStreamingServer) testEmbeddedByValue()                         {}

// UnsafeServerStreamingServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServerStreamingServer will
// result in compilation errors.
type UnsafeServerStreamingServer interface {
	mustEmbedUnimplementedServerStreamingServer()
}

func RegisterServerStreamingServer(s grpc.ServiceRegistrar, srv ServerStreamingServer) {
	// If the following call pancis, it indicates UnimplementedServerStreamingServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ServerStreaming_ServiceDesc, srv)
}

func _ServerStreaming_GetServerResponse_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Number)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ServerStreamingServer).GetServerResponse(m, &grpc.GenericServerStream[Number, Message]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type ServerStreaming_GetServerResponseServer = grpc.ServerStreamingServer[Message]

// ServerStreaming_ServiceDesc is the grpc.ServiceDesc for ServerStreaming service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ServerStreaming_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "serverstreaming.ServerStreaming",
	HandlerType: (*ServerStreamingServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetServerResponse",
			Handler:       _ServerStreaming_GetServerResponse_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "serverstreaming.proto",
}
