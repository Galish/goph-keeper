package grpc

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *KeeperServer) HealthCheck(_ context.Context, _ *emptypb.Empty) (*emptypb.Empty, error) {
	return nil, nil
}
