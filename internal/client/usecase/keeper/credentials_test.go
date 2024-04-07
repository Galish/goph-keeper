package keeper_test

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/Galish/goph-keeper/api/proto"
	mocks "github.com/Galish/goph-keeper/internal/client/infrastructure/grpc/mock"
	"github.com/Galish/goph-keeper/internal/client/usecase/keeper"
	"github.com/Galish/goph-keeper/internal/entity"
)

func TestAddCredentials(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockKeeperClient(ctrl)

	m.EXPECT().
		AddCredentials(gomock.Any(), gomock.Any()).
		Return(nil, status.Error(codes.InvalidArgument, errors.New("failed entity validation").Error()))

	m.EXPECT().
		AddCredentials(gomock.Any(), gomock.Any()).
		Return(nil, status.Error(codes.Internal, errors.New("failed to write to repo").Error()))

	m.EXPECT().
		AddCredentials(gomock.Any(), gomock.Any()).
		Return(nil, status.Error(codes.Unavailable, errors.New("no connection").Error()))

	m.EXPECT().
		AddCredentials(gomock.Any(), gomock.Any()).
		Return(&pb.AddCredentialsResponse{Id: "#12345"}, nil)

	uc := keeper.New(m)

	tests := []struct {
		name        string
		Credentials *entity.Credentials
		err         error
	}{
		{
			"missing entity",
			nil,
			keeper.ErrInvalidEntity,
		},
		{
			"invalid entity",
			&entity.Credentials{},
			keeper.ErrInvalidEntity,
		},
		{
			"failed validation",
			&entity.Credentials{
				Title:    "Gmail",
				Username: "john.doe",
				Password: "qwe123456",
			},
			keeper.ErrInvalidEntity,
		},
		{
			"writing to repo error",
			&entity.Credentials{
				Title:    "Gmail",
				Username: "john.doe",
				Password: "qwe123456",
			},
			errWriteToRepo,
		},
		{
			"no internet connection",
			&entity.Credentials{
				Title:    "Gmail",
				Username: "john.doe",
				Password: "qwe123456",
			},
			keeper.ErrNoConnection,
		},
		{
			"valid entity",
			&entity.Credentials{
				Title:    "Gmail",
				Username: "john.doe",
				Password: "qwe123456",
			},
			nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := uc.AddCredentials(context.Background(), tt.Credentials)

			if err != nil {
				assert.Error(t, err, tt.err.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestUpdateCredentials(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockKeeperClient(ctrl)

	m.EXPECT().
		UpdateCredentials(gomock.Any(), gomock.Any()).
		Return(nil, status.Error(codes.InvalidArgument, errors.New("failed entity validation").Error()))

	m.EXPECT().
		UpdateCredentials(gomock.Any(), gomock.Any()).
		Return(nil, status.Error(codes.FailedPrecondition, errors.New("record version conflict").Error()))

	m.EXPECT().
		UpdateCredentials(gomock.Any(), gomock.Any()).
		Return(nil, status.Error(codes.NotFound, errors.New("no entity found").Error()))

	m.EXPECT().
		UpdateCredentials(gomock.Any(), gomock.Any()).
		Return(nil, status.Error(codes.Internal, errors.New("failed to write to repo").Error()))

	m.EXPECT().
		UpdateCredentials(gomock.Any(), gomock.Any()).
		Return(nil, status.Error(codes.Unavailable, errors.New("no connection").Error()))

	m.EXPECT().
		UpdateCredentials(gomock.Any(), gomock.Any()).
		Return(nil, nil)

	uc := keeper.New(m)

	tests := []struct {
		name        string
		Credentials *entity.Credentials
		overwrite   bool
		err         error
	}{
		{
			"missing entity",
			nil,
			false,
			keeper.ErrInvalidEntity,
		},
		{
			"invalid entity",
			&entity.Credentials{
				ID:    "#12345678",
				Title: "Credit Credentials",
			},
			false,
			keeper.ErrInvalidEntity,
		},
		{
			"missing id",
			&entity.Credentials{
				Title:    "Gmail",
				Username: "john.doe",
				Password: "qwe123456",
			},
			false,
			keeper.ErrInvalidEntity,
		},
		{
			"failed validation",
			&entity.Credentials{
				ID:       "#12345678",
				Title:    "Gmail",
				Username: "john.doe",
				Password: "qwe123456",
			},
			true,
			keeper.ErrInvalidEntity,
		},
		{
			"version conflict",
			&entity.Credentials{
				ID:       "#12345678",
				Title:    "Gmail",
				Username: "john.doe",
				Password: "qwe123456",
			},
			false,
			keeper.ErrVersionConflict,
		},
		{
			"nothing found",
			&entity.Credentials{
				ID:       "#12345678",
				Title:    "Gmail",
				Username: "john.doe",
				Password: "qwe123456",
			},
			false,
			keeper.ErrNotFound,
		},
		{
			"writing to repo error",
			&entity.Credentials{
				ID:       "#12345678",
				Title:    "Gmail",
				Username: "john.doe",
				Password: "qwe123456",
			},
			true,
			errWriteToRepo,
		},
		{
			"no internet connection",
			&entity.Credentials{
				ID:       "#12345678",
				Title:    "Gmail",
				Username: "john.doe",
				Password: "qwe123456",
			},
			false,
			keeper.ErrNoConnection,
		},
		{
			"valid entity",
			&entity.Credentials{
				ID:       "#12345678",
				Title:    "Gmail",
				Username: "john.doe",
				Password: "qwe123456",
			},
			true,
			nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := uc.UpdateCredentials(context.Background(), tt.Credentials, tt.overwrite)

			if err != nil {
				assert.Error(t, err, tt.err.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestGetCredentials(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockKeeperClient(ctrl)

	m.EXPECT().
		GetCredentials(gomock.Any(), gomock.Any()).
		Return(nil, status.Error(codes.NotFound, errors.New("no entity found").Error()))

	m.EXPECT().
		GetCredentials(gomock.Any(), gomock.Any()).
		Return(nil, status.Error(codes.Internal, errReadFromRepo.Error()))

	m.EXPECT().
		GetCredentials(gomock.Any(), gomock.Any()).
		Return(nil, status.Error(codes.Unavailable, errors.New("no connection").Error()))

	m.EXPECT().
		GetCredentials(gomock.Any(), gomock.Any()).
		Return(&pb.GetCredentialsResponse{
			Credentials: &pb.Credentials{
				Title:    "Gmail",
				Username: "john.doe",
				Password: "qwe123456",
			},
			Version: 10,
		}, nil)

	uc := keeper.New(m)

	type want struct {
		Credentials *entity.Credentials
		err         error
	}

	tests := []struct {
		name string
		id   string
		want *want
	}{
		{
			"missing argument",
			"",
			&want{
				nil,
				keeper.ErrMissingArgument,
			},
		},
		{
			"nothing found",
			"#12345",
			&want{
				nil,
				keeper.ErrNotFound,
			},
		},
		{
			"reading from repo error",
			"#12345",
			&want{
				nil,
				errReadFromRepo,
			},
		},
		{
			"no internet connection",
			"#12345",
			&want{
				nil,
				keeper.ErrNoConnection,
			},
		},
		{
			"valid entity",
			"#12345",
			&want{
				&entity.Credentials{
					Title:    "Gmail",
					Username: "john.doe",
					Password: "qwe123456",
					Version:  10,
				},
				nil,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Credentials, err := uc.GetCredentials(context.Background(), tt.id)

			assert.Equal(t, tt.want.Credentials, Credentials)

			if err != nil {
				assert.Error(t, err, tt.want.err.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestGetCredentialsList(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockKeeperClient(ctrl)

	m.EXPECT().
		GetCredentialsList(gomock.Any(), gomock.Any()).
		Return(&pb.GetListResponse{}, nil)

	m.EXPECT().
		GetCredentialsList(gomock.Any(), gomock.Any()).
		Return(nil, status.Error(codes.Internal, errReadFromRepo.Error()))

	m.EXPECT().
		GetCredentialsList(gomock.Any(), gomock.Any()).
		Return(nil, status.Error(codes.Unavailable, errors.New("no connection").Error()))

	m.EXPECT().
		GetCredentialsList(gomock.Any(), gomock.Any()).
		Return(&pb.GetListResponse{
			List: []*pb.ListItem{
				{
					Id:          "#12345",
					Title:       "Gmail",
					Description: "My main account",
				},
				{
					Id:          "#23456",
					Title:       "Yandex",
					Description: "Work account",
				},
			},
		}, nil)

	uc := keeper.New(m)

	type want struct {
		Credentialss []*entity.Credentials
		err          error
	}

	tests := []struct {
		name string
		want *want
	}{
		{
			"nothing found",
			&want{
				[]*entity.Credentials{},
				nil,
			},
		},
		{
			"reading from repo error",
			&want{
				nil,
				errReadFromRepo,
			},
		},
		{
			"no internet connection",
			&want{
				nil,
				keeper.ErrNoConnection,
			},
		},
		{
			"valid entity",
			&want{
				[]*entity.Credentials{
					{
						ID:          "#12345",
						Title:       "Gmail",
						Description: "My main account",
					},
					{
						ID:          "#23456",
						Title:       "Yandex",
						Description: "Work account",
					},
				},
				nil,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Credentialss, err := uc.GetCredentialsList(context.Background())

			assert.Equal(t, tt.want.Credentialss, Credentialss)

			if err != nil {
				assert.Error(t, err, tt.want.err.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestDeleteCredentials(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockKeeperClient(ctrl)

	m.EXPECT().
		DeleteCredentials(gomock.Any(), gomock.Any()).
		Return(nil, status.Error(codes.NotFound, errors.New("no entity found").Error()))

	m.EXPECT().
		DeleteCredentials(gomock.Any(), gomock.Any()).
		Return(nil, status.Error(codes.Internal, errReadFromRepo.Error()))

	m.EXPECT().
		DeleteCredentials(gomock.Any(), gomock.Any()).
		Return(nil, status.Error(codes.Unavailable, errors.New("no connection").Error()))

	m.EXPECT().
		DeleteCredentials(gomock.Any(), gomock.Any()).
		Return(nil, nil)

	uc := keeper.New(m)

	tests := []struct {
		name string
		id   string
		err  error
	}{
		{
			"missing argument",
			"",
			keeper.ErrMissingArgument,
		},
		{
			"nothing found",
			"#12345",
			keeper.ErrNotFound,
		},
		{
			"reading from repo error",
			"#12345",
			errReadFromRepo,
		},
		{
			"no internet connection",
			"#12345",
			keeper.ErrNoConnection,
		},
		{
			"successfully deleted",
			"#12345",
			nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := uc.DeleteCredentials(context.Background(), tt.id)

			if err != nil {
				assert.Error(t, err, tt.err.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
