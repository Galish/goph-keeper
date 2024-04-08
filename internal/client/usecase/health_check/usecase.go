package healthcheck

import (
	"context"
	"errors"
	"time"

	pb "github.com/Galish/goph-keeper/api/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

var (
	defaultTimeout = 10 * time.Second

	ErrNoConnection = errors.New("check your connection and try again")
)

type HealthCheckUseCase struct {
	client pb.KeeperClient
}

func New(client pb.KeeperClient) *HealthCheckUseCase {
	return &HealthCheckUseCase{
		client: client,
	}
}

func (hc *HealthCheckUseCase) Check(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()

	_, err := hc.client.HealthCheck(ctx, &emptypb.Empty{})
	if err != nil {
		e, ok := status.FromError(err)
		if !ok {
			return err
		}

		if e.Code() == codes.Unavailable {
			return ErrNoConnection
		}

		return errors.New(e.Message())
	}

	return nil
}
