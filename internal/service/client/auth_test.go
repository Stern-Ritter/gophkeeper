package client

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func TestAuthInterceptor(t *testing.T) {
	t.Run("successful unary call with authorization token", func(t *testing.T) {
		authToken := "token"
		client := ClientImpl{authToken: authToken}

		var capturedCtx context.Context
		invoker := func(ctx context.Context, method string, req interface{}, reply interface{}, cc *grpc.ClientConn, opts ...grpc.CallOption) error {
			capturedCtx = ctx
			return nil
		}

		err := client.AuthInterceptor(context.Background(), "/test.Service/Method", nil, nil, nil, invoker)
		assert.NoError(t, err, "Expected no error from the interceptor")

		md, ok := metadata.FromOutgoingContext(capturedCtx)
		assert.True(t, ok, "Expected metadata to be set in the outgoing context")
		assert.Contains(t, md[AuthorizationMetadataKey], authToken, "Expected authorization token in metadata")
	})

}

func TestAuthStreamInterceptor(t *testing.T) {
	t.Run("successful stream call with authorization token", func(t *testing.T) {
		authToken := "token"
		client := ClientImpl{authToken: authToken}

		var capturedCtx context.Context
		streamer := func(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, opts ...grpc.CallOption) (grpc.ClientStream, error) {
			capturedCtx = ctx
			return nil, nil
		}

		_, err := client.AuthStreamInterceptor(context.Background(), &grpc.StreamDesc{}, nil, "/test.Service/StreamMethod", streamer)
		assert.NoError(t, err, "Expected no error from the stream interceptor")

		md, ok := metadata.FromOutgoingContext(capturedCtx)
		assert.True(t, ok, "Expected metadata to be set in the outgoing context")
		assert.Contains(t, md[AuthorizationMetadataKey], authToken, "Expected authorization token in metadata")
	})
}
