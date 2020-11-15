// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// OrdenServiceClient is the client API for OrdenService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OrdenServiceClient interface {
	UploadBook(ctx context.Context, opts ...grpc.CallOption) (OrdenService_UploadBookClient, error)
	GetAddressChunks(ctx context.Context, in *BookName, opts ...grpc.CallOption) (OrdenService_GetAddressChunksClient, error)
	DownloadChunk(ctx context.Context, in *ChunkId, opts ...grpc.CallOption) (*SendChunk, error)
}

type ordenServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOrdenServiceClient(cc grpc.ClientConnInterface) OrdenServiceClient {
	return &ordenServiceClient{cc}
}

func (c *ordenServiceClient) UploadBook(ctx context.Context, opts ...grpc.CallOption) (OrdenService_UploadBookClient, error) {
	stream, err := c.cc.NewStream(ctx, &_OrdenService_serviceDesc.Streams[0], "/proto.OrdenService/uploadBook", opts...)
	if err != nil {
		return nil, err
	}
	x := &ordenServiceUploadBookClient{stream}
	return x, nil
}

type OrdenService_UploadBookClient interface {
	Send(*SendChunk) error
	CloseAndRecv() (*ReplyEmpty, error)
	grpc.ClientStream
}

type ordenServiceUploadBookClient struct {
	grpc.ClientStream
}

func (x *ordenServiceUploadBookClient) Send(m *SendChunk) error {
	return x.ClientStream.SendMsg(m)
}

func (x *ordenServiceUploadBookClient) CloseAndRecv() (*ReplyEmpty, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(ReplyEmpty)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *ordenServiceClient) GetAddressChunks(ctx context.Context, in *BookName, opts ...grpc.CallOption) (OrdenService_GetAddressChunksClient, error) {
	stream, err := c.cc.NewStream(ctx, &_OrdenService_serviceDesc.Streams[1], "/proto.OrdenService/getAddressChunks", opts...)
	if err != nil {
		return nil, err
	}
	x := &ordenServiceGetAddressChunksClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type OrdenService_GetAddressChunksClient interface {
	Recv() (*SendUbicacion, error)
	grpc.ClientStream
}

type ordenServiceGetAddressChunksClient struct {
	grpc.ClientStream
}

func (x *ordenServiceGetAddressChunksClient) Recv() (*SendUbicacion, error) {
	m := new(SendUbicacion)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *ordenServiceClient) DownloadChunk(ctx context.Context, in *ChunkId, opts ...grpc.CallOption) (*SendChunk, error) {
	out := new(SendChunk)
	err := c.cc.Invoke(ctx, "/proto.OrdenService/downloadChunk", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrdenServiceServer is the server API for OrdenService service.
// All implementations must embed UnimplementedOrdenServiceServer
// for forward compatibility
type OrdenServiceServer interface {
	UploadBook(OrdenService_UploadBookServer) error
	GetAddressChunks(*BookName, OrdenService_GetAddressChunksServer) error
	DownloadChunk(context.Context, *ChunkId) (*SendChunk, error)
	mustEmbedUnimplementedOrdenServiceServer()
}

// UnimplementedOrdenServiceServer must be embedded to have forward compatible implementations.
type UnimplementedOrdenServiceServer struct {
}

func (UnimplementedOrdenServiceServer) UploadBook(OrdenService_UploadBookServer) error {
	return status.Errorf(codes.Unimplemented, "method UploadBook not implemented")
}
func (UnimplementedOrdenServiceServer) GetAddressChunks(*BookName, OrdenService_GetAddressChunksServer) error {
	return status.Errorf(codes.Unimplemented, "method GetAddressChunks not implemented")
}
func (UnimplementedOrdenServiceServer) DownloadChunk(context.Context, *ChunkId) (*SendChunk, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DownloadChunk not implemented")
}
func (UnimplementedOrdenServiceServer) mustEmbedUnimplementedOrdenServiceServer() {}

// UnsafeOrdenServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrdenServiceServer will
// result in compilation errors.
type UnsafeOrdenServiceServer interface {
	mustEmbedUnimplementedOrdenServiceServer()
}

func RegisterOrdenServiceServer(s *grpc.Server, srv OrdenServiceServer) {
	s.RegisterService(&_OrdenService_serviceDesc, srv)
}

func _OrdenService_UploadBook_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(OrdenServiceServer).UploadBook(&ordenServiceUploadBookServer{stream})
}

type OrdenService_UploadBookServer interface {
	SendAndClose(*ReplyEmpty) error
	Recv() (*SendChunk, error)
	grpc.ServerStream
}

type ordenServiceUploadBookServer struct {
	grpc.ServerStream
}

func (x *ordenServiceUploadBookServer) SendAndClose(m *ReplyEmpty) error {
	return x.ServerStream.SendMsg(m)
}

func (x *ordenServiceUploadBookServer) Recv() (*SendChunk, error) {
	m := new(SendChunk)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _OrdenService_GetAddressChunks_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(BookName)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(OrdenServiceServer).GetAddressChunks(m, &ordenServiceGetAddressChunksServer{stream})
}

type OrdenService_GetAddressChunksServer interface {
	Send(*SendUbicacion) error
	grpc.ServerStream
}

type ordenServiceGetAddressChunksServer struct {
	grpc.ServerStream
}

func (x *ordenServiceGetAddressChunksServer) Send(m *SendUbicacion) error {
	return x.ServerStream.SendMsg(m)
}

func _OrdenService_DownloadChunk_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChunkId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrdenServiceServer).DownloadChunk(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.OrdenService/DownloadChunk",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrdenServiceServer).DownloadChunk(ctx, req.(*ChunkId))
	}
	return interceptor(ctx, in, info, handler)
}

var _OrdenService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.OrdenService",
	HandlerType: (*OrdenServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "downloadChunk",
			Handler:    _OrdenService_DownloadChunk_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "uploadBook",
			Handler:       _OrdenService_UploadBook_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "getAddressChunks",
			Handler:       _OrdenService_GetAddressChunks_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/proto.proto",
}
