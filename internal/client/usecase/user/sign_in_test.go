//go:build unit
// +build unit

package user_test

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/Galish/goph-keeper/api/proto"
	mocks "github.com/Galish/goph-keeper/internal/client/infrastructure/grpc/mock"
	"github.com/Galish/goph-keeper/internal/client/usecase/user"
)

var (
	errReadFromRepo = errors.New("failed to read from repo")
)

func TestSignIn(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockKeeperClient(ctrl)

	m.EXPECT().
		SignIn(gomock.Any(), gomock.Any()).
		DoAndReturn(func(_ context.Context, in *pb.AuthRequest, _ ...grpc.CallOption) (*pb.AuthResponse, error) {
			if in.GetUsername() == "unavailable" {
				return nil, status.Error(codes.Unavailable, errors.New("no internet connection").Error())
			}

			if in.GetUsername() == "joe.doeeee" {
				return nil, status.Error(codes.Internal, errReadFromRepo.Error())
			}

			if in.GetUsername() != "john.doe" {
				return nil, status.Error(codes.NotFound, errors.New("user not found").Error())
			}

			if in.GetPassword() != "qwe123456" {
				return nil, status.Error(codes.InvalidArgument, errors.New("incorrect login/password pair").Error())
			}

			return &pb.AuthResponse{AccessToken: "access_token"}, nil
		}).
		AnyTimes()

	uc := user.New(m)

	tests := []struct {
		name     string
		username string
		password string
		err      error
	}{
		{
			"missing input",
			"",
			"",
			user.ErrInvalidCredentials,
		},
		{
			"missing username",
			"",
			"qwe123456",
			user.ErrInvalidCredentials,
		},
		{
			"missing password",
			"john.doe",
			"",
			user.ErrInvalidCredentials,
		},
		{
			"invalid credentials",
			"john.doe",
			"qwe1234",
			user.ErrInvalidCredentials,
		},
		{
			"user not found",
			"joe.doe",
			"qwe123456",
			user.ErrNotFound,
		},
		{
			"reading from repo error",
			"joe.doeeee",
			"qwe123456",
			errReadFromRepo,
		},
		{
			"no connection",
			"unavailable",
			"unavailable",
			user.ErrNoConnection,
		},
		{
			"successful login",
			"john.doe",
			"qwe123456",
			nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := uc.SignIn(context.Background(), tt.username, tt.password)

			if err != nil {
				assert.Equal(t, err, tt.err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
