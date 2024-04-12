package healthcheck

import (
	"context"
	"errors"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	pb "github.com/Galish/goph-keeper/api/proto"
)

var (
	defaultTimeout = 10 * time.Second

	ErrNoConnection = errors.New("check your connection and try again")
)

type UseCase struct {
	client pb.KeeperClient
}

func New(client pb.KeeperClient) *UseCase {
	return &UseCase{
		client: client,
	}
}

func (u *UseCase) Check(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()

	_, err := u.client.HealthCheck(ctx, &emptypb.Empty{})
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
