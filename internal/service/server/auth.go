package server

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	"github.com/Stern-Ritter/gophkeeper/internal/auth"
)

// ContextKey is a custom type used for context key to store authentication token
type ContextKey string

const (
	AuthorizationMetadataKey                = "authorization" // AuthorizationMetadataKey is the key used to retrieve the authorization token from gRPC metadata.
	AuthorizationTokenContextKey ContextKey = "user"          // AuthorizationTokenContextKey is the context key used to store the authentication token claims in the context.
)

// AuthInterceptor is a gRPC UnaryServerInterceptor that handles authentication for unary RPC calls.
// It checks the authorization token in the incoming request metadata. If the token is valid, the authentication
// claims are stored in the context for further use by the handler. If the token is missing or invalid,
// the request is rejected with an "Unauthenticated" status.
func (s *Server) AuthInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (interface{}, error) {
	exemptedMethods := map[string]bool{
		"/gophkeeper.gophkeeperapi.v1.AuthServiceV1/SignUp": true,
		"/gophkeeper.gophkeeperapi.v1.AuthServiceV1/SignIn": true,
	}

	if _, exempt := exemptedMethods[info.FullMethod]; exempt {
		return handler(ctx, req)
	}

	var tokenStr string

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		s.Logger.Error("Auth interceptor: missing request metadata")
		return nil, status.Errorf(codes.Unauthenticated, "missing auth token")
	}

	values := md.Get(AuthorizationMetadataKey)
	if len(values) == 0 {
		s.Logger.Error("Auth interceptor: missing request auth token")
		return nil, status.Errorf(codes.Unauthenticated, "missing auth token")
	}
	tokenStr = values[0]

	if len(tokenStr) == 0 {
		s.Logger.Error("Auth interceptor: missing request auth token")
		return nil, status.Errorf(codes.Unauthenticated, "missing auth token")
	}

	token, err := auth.ValidateToken(tokenStr, s.Config.AuthenticationKey)
	if err != nil {
		s.Logger.Error("Auth interceptor: invalid request auth token")
		return nil, status.Errorf(codes.Unauthenticated, "invalid auth token")
	}

	ctx = context.WithValue(ctx, AuthorizationTokenContextKey, token.Claims)
	return handler(ctx, req)
}

// AuthStreamInterceptor is a gRPC StreamServerInterceptor that handles authentication for streaming RPC calls.
// It checks the authorization token in the incoming stream's metadata. If the token is valid, the authentication
// claims are stored in the context for further use by the stream handler. If the token is missing or invalid,
// the stream is rejected with an "Unauthenticated" status.
func (s *Server) AuthStreamInterceptor(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo,
	handler grpc.StreamHandler) error {
	md, ok := metadata.FromIncomingContext(ss.Context())
	if !ok {
		s.Logger.Error("Auth stream interceptor: missing request metadata")
		return status.Errorf(codes.Unauthenticated, "missing auth token")
	}

	values := md.Get("authorization")
	if len(values) == 0 {
		s.Logger.Error("Auth stream interceptor: missing request auth token")
		return status.Errorf(codes.Unauthenticated, "missing auth token")
	}
	tokenStr := values[0]

	if len(tokenStr) == 0 {
		s.Logger.Error("Auth stream interceptor: missing request auth token")
		return status.Errorf(codes.Unauthenticated, "missing auth token")
	}

	token, err := auth.ValidateToken(tokenStr, s.Config.AuthenticationKey)
	if err != nil {
		s.Logger.Error("Auth stream interceptor: invalid request auth token")
		return status.Errorf(codes.Unauthenticated, "invalid auth token")
	}

	ctx := context.WithValue(ss.Context(), AuthorizationTokenContextKey, token.Claims)

	wrappedStream := &wrappedServerStream{
		ServerStream: ss,
		ctx:          ctx,
	}

	return handler(srv, wrappedStream)
}

type wrappedServerStream struct {
	grpc.ServerStream
	ctx context.Context
}

func (w *wrappedServerStream) Context() context.Context {
	return w.ctx
}
