// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             (unknown)
// source: gophkeeper/gophkeeperapi/v1/file_service.proto

package v1

import (
	context "context"

	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	FileServiceV1_UploadFile_FullMethodName   = "/gophkeeper.gophkeeperapi.v1.FileServiceV1/UploadFile"
	FileServiceV1_DownloadFile_FullMethodName = "/gophkeeper.gophkeeperapi.v1.FileServiceV1/DownloadFile"
	FileServiceV1_DeleteFile_FullMethodName   = "/gophkeeper.gophkeeperapi.v1.FileServiceV1/DeleteFile"
	FileServiceV1_GetFiles_FullMethodName     = "/gophkeeper.gophkeeperapi.v1.FileServiceV1/GetFiles"
)

// FileServiceV1Client is the client API for FileServiceV1 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FileServiceV1Client interface {
	UploadFile(ctx context.Context, opts ...grpc.CallOption) (FileServiceV1_UploadFileClient, error)
	DownloadFile(ctx context.Context, in *DownloadFileRequestV1, opts ...grpc.CallOption) (FileServiceV1_DownloadFileClient, error)
	DeleteFile(ctx context.Context, in *DeleteFileRequestV1, opts ...grpc.CallOption) (*DeleteFileResponseV1, error)
	GetFiles(ctx context.Context, in *GetFilesRequestV1, opts ...grpc.CallOption) (*GetFilesResponseV1, error)
}

type fileServiceV1Client struct {
	cc grpc.ClientConnInterface
}

func NewFileServiceV1Client(cc grpc.ClientConnInterface) FileServiceV1Client {
	return &fileServiceV1Client{cc}
}

func (c *fileServiceV1Client) UploadFile(ctx context.Context, opts ...grpc.CallOption) (FileServiceV1_UploadFileClient, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &FileServiceV1_ServiceDesc.Streams[0], FileServiceV1_UploadFile_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &fileServiceV1UploadFileClient{ClientStream: stream}
	return x, nil
}

type FileServiceV1_UploadFileClient interface {
	Send(*UploadFileRequestV1) error
	CloseAndRecv() (*UploadFileResponseV1, error)
	grpc.ClientStream
}

type fileServiceV1UploadFileClient struct {
	grpc.ClientStream
}

func (x *fileServiceV1UploadFileClient) Send(m *UploadFileRequestV1) error {
	return x.ClientStream.SendMsg(m)
}

func (x *fileServiceV1UploadFileClient) CloseAndRecv() (*UploadFileResponseV1, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(UploadFileResponseV1)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *fileServiceV1Client) DownloadFile(ctx context.Context, in *DownloadFileRequestV1, opts ...grpc.CallOption) (FileServiceV1_DownloadFileClient, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &FileServiceV1_ServiceDesc.Streams[1], FileServiceV1_DownloadFile_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &fileServiceV1DownloadFileClient{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type FileServiceV1_DownloadFileClient interface {
	Recv() (*DownloadFileResponseV1, error)
	grpc.ClientStream
}

type fileServiceV1DownloadFileClient struct {
	grpc.ClientStream
}

func (x *fileServiceV1DownloadFileClient) Recv() (*DownloadFileResponseV1, error) {
	m := new(DownloadFileResponseV1)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *fileServiceV1Client) DeleteFile(ctx context.Context, in *DeleteFileRequestV1, opts ...grpc.CallOption) (*DeleteFileResponseV1, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteFileResponseV1)
	err := c.cc.Invoke(ctx, FileServiceV1_DeleteFile_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileServiceV1Client) GetFiles(ctx context.Context, in *GetFilesRequestV1, opts ...grpc.CallOption) (*GetFilesResponseV1, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetFilesResponseV1)
	err := c.cc.Invoke(ctx, FileServiceV1_GetFiles_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FileServiceV1Server is the server API for FileServiceV1 service.
// All implementations must embed UnimplementedFileServiceV1Server
// for forward compatibility
type FileServiceV1Server interface {
	UploadFile(FileServiceV1_UploadFileServer) error
	DownloadFile(*DownloadFileRequestV1, FileServiceV1_DownloadFileServer) error
	DeleteFile(context.Context, *DeleteFileRequestV1) (*DeleteFileResponseV1, error)
	GetFiles(context.Context, *GetFilesRequestV1) (*GetFilesResponseV1, error)
	mustEmbedUnimplementedFileServiceV1Server()
}

// UnimplementedFileServiceV1Server must be embedded to have forward compatible implementations.
type UnimplementedFileServiceV1Server struct {
}

func (UnimplementedFileServiceV1Server) UploadFile(FileServiceV1_UploadFileServer) error {
	return status.Errorf(codes.Unimplemented, "method UploadFile not implemented")
}
func (UnimplementedFileServiceV1Server) DownloadFile(*DownloadFileRequestV1, FileServiceV1_DownloadFileServer) error {
	return status.Errorf(codes.Unimplemented, "method DownloadFile not implemented")
}
func (UnimplementedFileServiceV1Server) DeleteFile(context.Context, *DeleteFileRequestV1) (*DeleteFileResponseV1, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteFile not implemented")
}
func (UnimplementedFileServiceV1Server) GetFiles(context.Context, *GetFilesRequestV1) (*GetFilesResponseV1, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFiles not implemented")
}
func (UnimplementedFileServiceV1Server) mustEmbedUnimplementedFileServiceV1Server() {}

// UnsafeFileServiceV1Server may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FileServiceV1Server will
// result in compilation errors.
type UnsafeFileServiceV1Server interface {
	mustEmbedUnimplementedFileServiceV1Server()
}

func RegisterFileServiceV1Server(s grpc.ServiceRegistrar, srv FileServiceV1Server) {
	s.RegisterService(&FileServiceV1_ServiceDesc, srv)
}

func _FileServiceV1_UploadFile_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(FileServiceV1Server).UploadFile(&fileServiceV1UploadFileServer{ServerStream: stream})
}

type FileServiceV1_UploadFileServer interface {
	SendAndClose(*UploadFileResponseV1) error
	Recv() (*UploadFileRequestV1, error)
	grpc.ServerStream
}

type fileServiceV1UploadFileServer struct {
	grpc.ServerStream
}

func (x *fileServiceV1UploadFileServer) SendAndClose(m *UploadFileResponseV1) error {
	return x.ServerStream.SendMsg(m)
}

func (x *fileServiceV1UploadFileServer) Recv() (*UploadFileRequestV1, error) {
	m := new(UploadFileRequestV1)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _FileServiceV1_DownloadFile_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(DownloadFileRequestV1)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(FileServiceV1Server).DownloadFile(m, &fileServiceV1DownloadFileServer{ServerStream: stream})
}

type FileServiceV1_DownloadFileServer interface {
	Send(*DownloadFileResponseV1) error
	grpc.ServerStream
}

type fileServiceV1DownloadFileServer struct {
	grpc.ServerStream
}

func (x *fileServiceV1DownloadFileServer) Send(m *DownloadFileResponseV1) error {
	return x.ServerStream.SendMsg(m)
}

func _FileServiceV1_DeleteFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteFileRequestV1)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileServiceV1Server).DeleteFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FileServiceV1_DeleteFile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileServiceV1Server).DeleteFile(ctx, req.(*DeleteFileRequestV1))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileServiceV1_GetFiles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFilesRequestV1)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileServiceV1Server).GetFiles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FileServiceV1_GetFiles_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileServiceV1Server).GetFiles(ctx, req.(*GetFilesRequestV1))
	}
	return interceptor(ctx, in, info, handler)
}

// FileServiceV1_ServiceDesc is the grpc.ServiceDesc for FileServiceV1 service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FileServiceV1_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gophkeeper.gophkeeperapi.v1.FileServiceV1",
	HandlerType: (*FileServiceV1Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DeleteFile",
			Handler:    _FileServiceV1_DeleteFile_Handler,
		},
		{
			MethodName: "GetFiles",
			Handler:    _FileServiceV1_GetFiles_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "UploadFile",
			Handler:       _FileServiceV1_UploadFile_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "DownloadFile",
			Handler:       _FileServiceV1_DownloadFile_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "gophkeeper/gophkeeperapi/v1/file_service.proto",
}