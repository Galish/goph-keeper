package interceptors

import (
	"context"

	"github.com/Galish/goph-keeper/internal/client/auth"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func NewAuthInterceptor(authClient *auth.AuthManager) grpc.UnaryClientInterceptor {
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
