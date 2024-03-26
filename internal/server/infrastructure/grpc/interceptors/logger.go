package interceptors

import (
	"context"
	"time"

	"google.golang.org/grpc"

	"github.com/Galish/goph-keeper/pkg/logger"
)

// LoggerInterceptor provides request logging.
func LoggerInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	start := time.Now()

	var user string
	if ctx.Value(UserContextKey) != nil {
		user = ctx.Value(UserContextKey).(string)
	}

	logger.WithFields(logger.Fields{
		"duration": time.Since(start),
		"method":   info.FullMethod,
		"user":     user,
	}).Info("incoming request")

	return handler(ctx, req)
}
