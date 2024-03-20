package interceptors

import (
	"context"
	"errors"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type contextKey string

var (
	UserContextKey = contextKey("user")

	ErrMissingUserID = errors.New("missing user identifier")
)

var authMethods = map[string]bool{
	"/service.Keeper/SignIn": true,
	"/service.Keeper/SignUp": true,
}

// UserCheckInterceptor serves authentication error if user identifier not provided.
func UserCheckInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	if isAuth := authMethods[info.FullMethod]; isAuth {
		return handler(ctx, req)
	}

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, ErrMissingUserID.Error())
	}

	values := md.Get("user")
	if len(values) == 0 {
		return nil, status.Error(codes.Unauthenticated, ErrMissingUserID.Error())
	}

	if values[0] == "" {
		return nil, status.Error(codes.Unauthenticated, ErrMissingUserID.Error())
	}

	ctx = context.WithValue(ctx, UserContextKey, values[0])

	return handler(ctx, req)

}
