// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             (unknown)
// source: gophkeeper/gophkeeperapi/v1/card_service.proto

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
	CardServiceV1_AddCard_FullMethodName    = "/gophkeeper.gophkeeperapi.v1.CardServiceV1/AddCard"
	CardServiceV1_DeleteCard_FullMethodName = "/gophkeeper.gophkeeperapi.v1.CardServiceV1/DeleteCard"
	CardServiceV1_GetCards_FullMethodName   = "/gophkeeper.gophkeeperapi.v1.CardServiceV1/GetCards"
)

// CardServiceV1Client is the client API for CardServiceV1 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CardServiceV1Client interface {
	AddCard(ctx context.Context, in *AddCardRequestV1, opts ...grpc.CallOption) (*AddCardResponseV1, error)
	DeleteCard(ctx context.Context, in *DeleteCardRequestV1, opts ...grpc.CallOption) (*DeleteCardResponseV1, error)
	GetCards(ctx context.Context, in *GetCardsRequestV1, opts ...grpc.CallOption) (*GetCardsResponseV1, error)
}

type cardServiceV1Client struct {
	cc grpc.ClientConnInterface
}

func NewCardServiceV1Client(cc grpc.ClientConnInterface) CardServiceV1Client {
	return &cardServiceV1Client{cc}
}

func (c *cardServiceV1Client) AddCard(ctx context.Context, in *AddCardRequestV1, opts ...grpc.CallOption) (*AddCardResponseV1, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddCardResponseV1)
	err := c.cc.Invoke(ctx, CardServiceV1_AddCard_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cardServiceV1Client) DeleteCard(ctx context.Context, in *DeleteCardRequestV1, opts ...grpc.CallOption) (*DeleteCardResponseV1, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteCardResponseV1)
	err := c.cc.Invoke(ctx, CardServiceV1_DeleteCard_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cardServiceV1Client) GetCards(ctx context.Context, in *GetCardsRequestV1, opts ...grpc.CallOption) (*GetCardsResponseV1, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetCardsResponseV1)
	err := c.cc.Invoke(ctx, CardServiceV1_GetCards_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CardServiceV1Server is the server API for CardServiceV1 service.
// All implementations must embed UnimplementedCardServiceV1Server
// for forward compatibility
type CardServiceV1Server interface {
	AddCard(context.Context, *AddCardRequestV1) (*AddCardResponseV1, error)
	DeleteCard(context.Context, *DeleteCardRequestV1) (*DeleteCardResponseV1, error)
	GetCards(context.Context, *GetCardsRequestV1) (*GetCardsResponseV1, error)
	mustEmbedUnimplementedCardServiceV1Server()
}

// UnimplementedCardServiceV1Server must be embedded to have forward compatible implementations.
type UnimplementedCardServiceV1Server struct {
}

func (UnimplementedCardServiceV1Server) AddCard(context.Context, *AddCardRequestV1) (*AddCardResponseV1, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddCard not implemented")
}
func (UnimplementedCardServiceV1Server) DeleteCard(context.Context, *DeleteCardRequestV1) (*DeleteCardResponseV1, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCard not implemented")
}
func (UnimplementedCardServiceV1Server) GetCards(context.Context, *GetCardsRequestV1) (*GetCardsResponseV1, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCards not implemented")
}
func (UnimplementedCardServiceV1Server) mustEmbedUnimplementedCardServiceV1Server() {}

// UnsafeCardServiceV1Server may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CardServiceV1Server will
// result in compilation errors.
type UnsafeCardServiceV1Server interface {
	mustEmbedUnimplementedCardServiceV1Server()
}

func RegisterCardServiceV1Server(s grpc.ServiceRegistrar, srv CardServiceV1Server) {
	s.RegisterService(&CardServiceV1_ServiceDesc, srv)
}

func _CardServiceV1_AddCard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddCardRequestV1)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CardServiceV1Server).AddCard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CardServiceV1_AddCard_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CardServiceV1Server).AddCard(ctx, req.(*AddCardRequestV1))
	}
	return interceptor(ctx, in, info, handler)
}

func _CardServiceV1_DeleteCard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCardRequestV1)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CardServiceV1Server).DeleteCard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CardServiceV1_DeleteCard_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CardServiceV1Server).DeleteCard(ctx, req.(*DeleteCardRequestV1))
	}
	return interceptor(ctx, in, info, handler)
}

func _CardServiceV1_GetCards_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCardsRequestV1)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CardServiceV1Server).GetCards(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CardServiceV1_GetCards_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CardServiceV1Server).GetCards(ctx, req.(*GetCardsRequestV1))
	}
	return interceptor(ctx, in, info, handler)
}

// CardServiceV1_ServiceDesc is the grpc.ServiceDesc for CardServiceV1 service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CardServiceV1_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gophkeeper.gophkeeperapi.v1.CardServiceV1",
	HandlerType: (*CardServiceV1Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddCard",
			Handler:    _CardServiceV1_AddCard_Handler,
		},
		{
			MethodName: "DeleteCard",
			Handler:    _CardServiceV1_DeleteCard_Handler,
		},
		{
			MethodName: "GetCards",
			Handler:    _CardServiceV1_GetCards_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gophkeeper/gophkeeperapi/v1/card_service.proto",
}