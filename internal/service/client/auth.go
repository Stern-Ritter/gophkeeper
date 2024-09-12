package client

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

const (
	AuthorizationMetadataKey = "authorization"
)

// AuthInterceptor is a gRPC unary interceptor that adds an authorization token to the outgoing context metadata.
func (c *ClientImpl) AuthInterceptor(
	ctx context.Context,
	method string,
	req interface{},
	reply interface{},
	cc *grpc.ClientConn,
	invoker grpc.UnaryInvoker,
	opts ...grpc.CallOption,
) error {
	md := metadata.Pairs(AuthorizationMetadataKey, c.authToken)
	ctx = metadata.NewOutgoingContext(ctx, md)

	return invoker(ctx, method, req, reply, cc, opts...)
}

// AuthStreamInterceptor is a gRPC stream interceptor that adds an authorization token to the outgoing context metadata.
func (c *ClientImpl) AuthStreamInterceptor(
	ctx context.Context,
	desc *grpc.StreamDesc,
	cc *grpc.ClientConn,
	method string,
	streamer grpc.Streamer,
	opts ...grpc.CallOption,
) (grpc.ClientStream, error) {
	md := metadata.Pairs(AuthorizationMetadataKey, c.authToken)
	ctx = metadata.NewOutgoingContext(ctx, md)

	return streamer(ctx, desc, cc, method, opts...)
}
