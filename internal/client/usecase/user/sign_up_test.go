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
	errWriteToRepo = errors.New("failed to write to repo")
)

func TestSignUp(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockKeeperClient(ctrl)

	m.EXPECT().
		SignUp(gomock.Any(), gomock.Any()).
		DoAndReturn(func(_ context.Context, in *pb.AuthRequest, opts ...grpc.CallOption) (*pb.AuthResponse, error) {
			if in.Username == "unavailable" {
				return nil, status.Error(codes.Unavailable, errors.New("no internet connection").Error())
			}

			if in.Username == "joe.doeeee" {
				return nil, status.Error(codes.Internal, errWriteToRepo.Error())
			}

			if in.Username == "johny.doe" {
				return nil, status.Error(codes.AlreadyExists, errors.New("user already exists").Error())
			}

			if in.Password != "qwe123456" {
				return nil, status.Error(codes.InvalidArgument, errors.New("missing login/password").Error())
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
			"johny.doe",
			"qwe123456",
			user.ErrAlreadyExists,
		},
		{
			"writing from repo error",
			"joe.doeeee",
			"qwe123456",
			errWriteToRepo,
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
			err := uc.SignUp(context.Background(), tt.username, tt.password)

			if err != nil {
				assert.Error(t, err, tt.err.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
