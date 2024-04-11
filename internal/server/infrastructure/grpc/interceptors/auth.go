package interceptors

import (
	"context"
	"errors"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	"github.com/Galish/goph-keeper/internal/server/usecase"
)

type contextKey string

var (
	publicMethods = map[string]bool{
		"/service.Keeper/SignIn":      true,
		"/service.Keeper/SignUp":      true,
		"/service.Keeper/HealthCheck": true,
	}

	ErrInvalidAccessToken = errors.New("access token is invalid")
	ErrNoAccessToken      = errors.New("access token token is not provided")

	UserContextKey = contextKey("user")
)

type AuthInterceptor struct {
	user usecase.User
}

func NewAuthInterceptor(user usecase.User) *AuthInterceptor {
	return &AuthInterceptor{
		user: user,
	}
}

func (ai *AuthInterceptor) Unary() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		if isPublic := publicMethods[info.FullMethod]; isPublic {
			return handler(ctx, req)
		}

		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, status.Error(codes.Unauthenticated, ErrNoAccessToken.Error())
		}

		values := md.Get("authorization")
		if len(values) == 0 {
			return nil, status.Error(codes.Unauthenticated, ErrNoAccessToken.Error())
		}

		if values[0] == "" {
			return nil, status.Error(codes.Unauthenticated, ErrNoAccessToken.Error())
		}

		user, err := ai.user.Verify(values[0])
		if err != nil {
			return nil, status.Error(codes.Unauthenticated, ErrInvalidAccessToken.Error())
		}

		ctx = context.WithValue(ctx, UserContextKey, user.ID)

		return handler(ctx, req)
	}
}
