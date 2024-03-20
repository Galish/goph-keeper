package interceptors

import (
	"context"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

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
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		values := md.Get("token")
		if len(values) > 0 {
			user = values[0]
		}
	}

	logger.WithFields(logger.Fields{
		"duration": time.Since(start),
		"method":   info.FullMethod,
		"user":     user,
	}).Info("incoming request")

	return handler(ctx, req)
}
