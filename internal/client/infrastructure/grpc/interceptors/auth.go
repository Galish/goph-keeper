// Package contains gRPC client interceptors.
package interceptors

import (
	"context"

	"github.com/Galish/goph-keeper/internal/client/auth"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// New AuthInterceptor returns a new interceptor that handles authentication by adding an access token to the request metadata.
func NewAuthInterceptor(authClient *auth.Manager) grpc.UnaryClientInterceptor {
	return func(
		ctx context.Context,
		method string,
		req, reply interface{},
		cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker,
		opts ...grpc.CallOption,
	) error {
		ctx = metadata.AppendToOutgoingContext(
			ctx,
			"authorization",
			authClient.GetToken(),
		)

		return invoker(ctx, method, req, reply, cc, opts...)
	}
}
