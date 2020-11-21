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

// LibroServiceClient is the client API for LibroService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LibroServiceClient interface {
	UploadBook(ctx context.Context, opts ...grpc.CallOption) (LibroService_UploadBookClient, error)
	GetAddressChunks(ctx context.Context, in *BookName, opts ...grpc.CallOption) (LibroService_GetAddressChunksClient, error)
	DownloadChunk(ctx context.Context, in *ChunkId, opts ...grpc.CallOption) (*SendChunk, error)
	SendPropuesta(ctx context.Context, in *Propuesta, opts ...grpc.CallOption) (*Propuesta, error)
	OrdenarChunk(ctx context.Context, in *SendChunk, opts ...grpc.CallOption) (*ReplyEmpty, error)
	VerStatus(ctx context.Context, in *Status, opts ...grpc.CallOption) (*Status, error)
}

type libroServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLibroServiceClient(cc grpc.ClientConnInterface) LibroServiceClient {
	return &libroServiceClient{cc}
}

func (c *libroServiceClient) UploadBook(ctx context.Context, opts ...grpc.CallOption) (LibroService_UploadBookClient, error) {
	stream, err := c.cc.NewStream(ctx, &_LibroService_serviceDesc.Streams[0], "/proto.LibroService/uploadBook", opts...)
	if err != nil {
		return nil, err
	}
	x := &libroServiceUploadBookClient{stream}
	return x, nil
}

type LibroService_UploadBookClient interface {
	Send(*SendChunk) error
	CloseAndRecv() (*ReplyEmpty, error)
	grpc.ClientStream
}

type libroServiceUploadBookClient struct {
	grpc.ClientStream
}

func (x *libroServiceUploadBookClient) Send(m *SendChunk) error {
	return x.ClientStream.SendMsg(m)
}

func (x *libroServiceUploadBookClient) CloseAndRecv() (*ReplyEmpty, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(ReplyEmpty)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *libroServiceClient) GetAddressChunks(ctx context.Context, in *BookName, opts ...grpc.CallOption) (LibroService_GetAddressChunksClient, error) {
	stream, err := c.cc.NewStream(ctx, &_LibroService_serviceDesc.Streams[1], "/proto.LibroService/getAddressChunks", opts...)
	if err != nil {
		return nil, err
	}
	x := &libroServiceGetAddressChunksClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type LibroService_GetAddressChunksClient interface {
	Recv() (*SendUbicacion, error)
	grpc.ClientStream
}

type libroServiceGetAddressChunksClient struct {
	grpc.ClientStream
}

func (x *libroServiceGetAddressChunksClient) Recv() (*SendUbicacion, error) {
	m := new(SendUbicacion)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *libroServiceClient) DownloadChunk(ctx context.Context, in *ChunkId, opts ...grpc.CallOption) (*SendChunk, error) {
	out := new(SendChunk)
	err := c.cc.Invoke(ctx, "/proto.LibroService/downloadChunk", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *libroServiceClient) SendPropuesta(ctx context.Context, in *Propuesta, opts ...grpc.CallOption) (*Propuesta, error) {
	out := new(Propuesta)
	err := c.cc.Invoke(ctx, "/proto.LibroService/sendPropuesta", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *libroServiceClient) OrdenarChunk(ctx context.Context, in *SendChunk, opts ...grpc.CallOption) (*ReplyEmpty, error) {
	out := new(ReplyEmpty)
	err := c.cc.Invoke(ctx, "/proto.LibroService/OrdenarChunk", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *libroServiceClient) VerStatus(ctx context.Context, in *Status, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/proto.LibroService/verStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LibroServiceServer is the server API for LibroService service.
// All implementations must embed UnimplementedLibroServiceServer
// for forward compatibility
type LibroServiceServer interface {
	UploadBook(LibroService_UploadBookServer) error
	GetAddressChunks(*BookName, LibroService_GetAddressChunksServer) error
	DownloadChunk(context.Context, *ChunkId) (*SendChunk, error)
	SendPropuesta(context.Context, *Propuesta) (*Propuesta, error)
	OrdenarChunk(context.Context, *SendChunk) (*ReplyEmpty, error)
	VerStatus(context.Context, *Status) (*Status, error)
	mustEmbedUnimplementedLibroServiceServer()
}

// UnimplementedLibroServiceServer must be embedded to have forward compatible implementations.
type UnimplementedLibroServiceServer struct {
}

func (UnimplementedLibroServiceServer) UploadBook(LibroService_UploadBookServer) error {
	return status.Errorf(codes.Unimplemented, "method UploadBook not implemented")
}
func (UnimplementedLibroServiceServer) GetAddressChunks(*BookName, LibroService_GetAddressChunksServer) error {
	return status.Errorf(codes.Unimplemented, "method GetAddressChunks not implemented")
}
func (UnimplementedLibroServiceServer) DownloadChunk(context.Context, *ChunkId) (*SendChunk, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DownloadChunk not implemented")
}
func (UnimplementedLibroServiceServer) SendPropuesta(context.Context, *Propuesta) (*Propuesta, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendPropuesta not implemented")
}
func (UnimplementedLibroServiceServer) OrdenarChunk(context.Context, *SendChunk) (*ReplyEmpty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OrdenarChunk not implemented")
}
func (UnimplementedLibroServiceServer) VerStatus(context.Context, *Status) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerStatus not implemented")
}
func (UnimplementedLibroServiceServer) mustEmbedUnimplementedLibroServiceServer() {}

// UnsafeLibroServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LibroServiceServer will
// result in compilation errors.
type UnsafeLibroServiceServer interface {
	mustEmbedUnimplementedLibroServiceServer()
}

func RegisterLibroServiceServer(s *grpc.Server, srv LibroServiceServer) {
	s.RegisterService(&_LibroService_serviceDesc, srv)
}

func _LibroService_UploadBook_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(LibroServiceServer).UploadBook(&libroServiceUploadBookServer{stream})
}

type LibroService_UploadBookServer interface {
	SendAndClose(*ReplyEmpty) error
	Recv() (*SendChunk, error)
	grpc.ServerStream
}

type libroServiceUploadBookServer struct {
	grpc.ServerStream
}

func (x *libroServiceUploadBookServer) SendAndClose(m *ReplyEmpty) error {
	return x.ServerStream.SendMsg(m)
}

func (x *libroServiceUploadBookServer) Recv() (*SendChunk, error) {
	m := new(SendChunk)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _LibroService_GetAddressChunks_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(BookName)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(LibroServiceServer).GetAddressChunks(m, &libroServiceGetAddressChunksServer{stream})
}

type LibroService_GetAddressChunksServer interface {
	Send(*SendUbicacion) error
	grpc.ServerStream
}

type libroServiceGetAddressChunksServer struct {
	grpc.ServerStream
}

func (x *libroServiceGetAddressChunksServer) Send(m *SendUbicacion) error {
	return x.ServerStream.SendMsg(m)
}

func _LibroService_DownloadChunk_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChunkId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LibroServiceServer).DownloadChunk(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.LibroService/DownloadChunk",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LibroServiceServer).DownloadChunk(ctx, req.(*ChunkId))
	}
	return interceptor(ctx, in, info, handler)
}

func _LibroService_SendPropuesta_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Propuesta)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LibroServiceServer).SendPropuesta(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.LibroService/SendPropuesta",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LibroServiceServer).SendPropuesta(ctx, req.(*Propuesta))
	}
	return interceptor(ctx, in, info, handler)
}

func _LibroService_OrdenarChunk_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendChunk)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LibroServiceServer).OrdenarChunk(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.LibroService/OrdenarChunk",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LibroServiceServer).OrdenarChunk(ctx, req.(*SendChunk))
	}
	return interceptor(ctx, in, info, handler)
}

func _LibroService_VerStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Status)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LibroServiceServer).VerStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.LibroService/VerStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LibroServiceServer).VerStatus(ctx, req.(*Status))
	}
	return interceptor(ctx, in, info, handler)
}

var _LibroService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.LibroService",
	HandlerType: (*LibroServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "downloadChunk",
			Handler:    _LibroService_DownloadChunk_Handler,
		},
		{
			MethodName: "sendPropuesta",
			Handler:    _LibroService_SendPropuesta_Handler,
		},
		{
			MethodName: "OrdenarChunk",
			Handler:    _LibroService_OrdenarChunk_Handler,
		},
		{
			MethodName: "verStatus",
			Handler:    _LibroService_VerStatus_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "uploadBook",
			Handler:       _LibroService_UploadBook_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "getAddressChunks",
			Handler:       _LibroService_GetAddressChunks_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/proto.proto",
}
