//go:build unit
// +build unit

package notes_test

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
	"github.com/Galish/goph-keeper/internal/client/usecase/notes"
	"github.com/Galish/goph-keeper/internal/entity"
)

func TestAddCredentials(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockKeeperClient(ctrl)

	m.EXPECT().
		AddCredentials(gomock.Any(), gomock.Any()).
		Return(nil, status.Error(codes.Unauthenticated, errors.New("authorization required").Error()))

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

	uc := notes.New(m)

	tests := []struct {
		name        string
		Credentials *entity.Credentials
		err         error
	}{
		{
			"missing entity",
			nil,
			notes.ErrInvalidEntity,
		},
		{
			"invalid entity",
			&entity.Credentials{},
			notes.ErrInvalidEntity,
		},
		{
			"authorization required",
			&entity.Credentials{
				Title:    "Gmail",
				Username: "john.doe",
				Password: "qwe123456",
			},
			notes.ErrAuthRequired,
		},
		{
			"failed validation",
			&entity.Credentials{
				Title:    "Gmail",
				Username: "john.doe",
				Password: "qwe123456",
			},
			notes.ErrInvalidEntity,
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
			notes.ErrNoConnection,
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
				assert.Equal(t, err, tt.err)
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
		Return(nil, status.Error(codes.Unauthenticated, errors.New("authorization required").Error()))

	m.EXPECT().
		UpdateCredentials(gomock.Any(), gomock.Any()).
		Return(nil, status.Error(codes.InvalidArgument, errors.New("failed entity validation").Error()))

	m.EXPECT().
		UpdateCredentials(gomock.Any(), gomock.Any()).
		Return(nil, status.Error(codes.FailedPrecondition, errors.New("entity version conflict").Error()))

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

	uc := notes.New(m)

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
			notes.ErrInvalidEntity,
		},
		{
			"invalid entity",
			&entity.Credentials{
				ID:    "#12345678",
				Title: "Credit Credentials",
			},
			false,
			notes.ErrInvalidEntity,
		},
		{
			"missing id",
			&entity.Credentials{
				Title:    "Gmail",
				Username: "john.doe",
				Password: "qwe123456",
			},
			false,
			notes.ErrInvalidEntity,
		},
		{
			"authorization required",
			&entity.Credentials{
				ID:       "#12345678",
				Title:    "Gmail",
				Username: "john.doe",
				Password: "qwe123456",
			},
			false,
			notes.ErrAuthRequired,
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
			notes.ErrInvalidEntity,
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
			notes.ErrVersionConflict,
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
			notes.ErrNotFound,
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
			notes.ErrNoConnection,
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
				assert.Equal(t, err, tt.err)
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
		Return(nil, status.Error(codes.Unauthenticated, errors.New("authorization required").Error()))

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

	uc := notes.New(m)

	type want struct {
		creds *entity.Credentials
		err   error
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
				notes.ErrMissingArgument,
			},
		},
		{
			"authorization required",
			"#12345",
			&want{
				nil,
				notes.ErrAuthRequired,
			},
		},
		{
			"nothing found",
			"#12345",
			&want{
				nil,
				notes.ErrNotFound,
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
				notes.ErrNoConnection,
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
			creds, err := uc.GetCredentials(context.Background(), tt.id)

			assert.Equal(t, tt.want.creds, creds)

			if err != nil {
				assert.Equal(t, err, tt.want.err)
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
		Return(nil, status.Error(codes.Unauthenticated, errors.New("authorization required").Error()))

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

	uc := notes.New(m)

	type want struct {
		creds []*entity.Credentials
		err   error
	}

	tests := []struct {
		name string
		want *want
	}{
		{
			"authorization required",
			&want{
				nil,
				notes.ErrAuthRequired,
			},
		},
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
				notes.ErrNoConnection,
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
			creds, err := uc.GetCredentialsList(context.Background())

			assert.Equal(t, tt.want.creds, creds)

			if err != nil {
				assert.Equal(t, err, tt.want.err)
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
		Return(nil, status.Error(codes.Unauthenticated, errors.New("authorization required").Error()))

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

	uc := notes.New(m)

	tests := []struct {
		name string
		id   string
		err  error
	}{
		{
			"missing argument",
			"",
			notes.ErrMissingArgument,
		},
		{
			"authorization required",
			"#12345",
			notes.ErrAuthRequired,
		},
		{
			"nothing found",
			"#12345",
			notes.ErrNotFound,
		},
		{
			"reading from repo error",
			"#12345",
			errReadFromRepo,
		},
		{
			"no internet connection",
			"#12345",
			notes.ErrNoConnection,
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
				assert.Equal(t, err, tt.err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
