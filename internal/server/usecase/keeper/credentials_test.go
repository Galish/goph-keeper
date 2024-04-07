package keeper_test

import (
	"context"
	"testing"

	"github.com/Galish/goph-keeper/internal/entity"
	"github.com/Galish/goph-keeper/internal/server/repository"
	"github.com/Galish/goph-keeper/internal/server/repository/mocks"
	"github.com/Galish/goph-keeper/internal/server/usecase/keeper"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestAddCredentials(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockKeeperRepository(ctrl)

	m.EXPECT().
		AddSecureRecord(gomock.Any(), gomock.Any()).
		DoAndReturn(func(_ context.Context, r *repository.SecureRecord) error {
			if r.ID == "#765432" {
				return errWriteToRepo
			}
			return nil
		}).
		AnyTimes()

	uc := keeper.New(m)

	tests := []struct {
		name  string
		creds *entity.Credentials
		err   error
	}{
		{
			"empty input",
			nil,
			keeper.ErrInvalidEntity,
		},
		{
			"invalid entity",
			&entity.Credentials{
				ID:       "#12345",
				Username: "john.doe",
			},
			keeper.ErrInvalidEntity,
		},
		{
			"valid entity",
			&entity.Credentials{
				ID:       "#12345",
				Title:    "Gmail",
				Username: "john.doe",
				Password: "qwe123456",
			},
			nil,
		},
		{
			"write to repo error",
			&entity.Credentials{
				ID:       "#765432",
				Title:    "Gmail",
				Username: "john.doe",
				Password: "qwe123456",
			},
			errWriteToRepo,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := uc.AddCredentials(context.Background(), tt.creds)

			assert.Equal(t, tt.err, err)
		})
	}
}

func TestGetCredentials(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockKeeperRepository(ctrl)

	m.EXPECT().
		GetSecureRecord(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Eq(repository.TypeCredentials)).
		DoAndReturn(func(_ context.Context, user, id string, t repository.SecureRecordType) (*repository.SecureRecord, error) {
			switch id {
			case "#12345":
				return &repository.SecureRecord{
					ID:       "#12345",
					Type:     repository.TypeCredentials,
					Title:    "Gmail",
					Username: "john.doe",
					Password: "qwe123456",
				}, nil

			case "#34567":
				return nil, repository.ErrNotFound

			default:
				return nil, errReadFromRepo
			}
		}).
		AnyTimes()

	uc := keeper.New(m)

	type want struct {
		creds *entity.Credentials
		err   error
	}

	tests := []struct {
		name string
		user string
		id   string
		want *want
	}{
		{
			"missing id",
			"user#12345",
			"",
			&want{
				nil,
				keeper.ErrMissingArgument,
			},
		},
		{
			"missing user",
			"",
			"#34567",
			&want{
				nil,
				keeper.ErrMissingArgument,
			},
		},
		{
			"nothing found",
			"user#12345",
			"#34567",
			&want{
				nil,
				keeper.ErrNotFound,
			},
		},
		{
			"valid credentials",
			"user#12345",
			"#12345",
			&want{
				&entity.Credentials{
					ID:       "#12345",
					Title:    "Gmail",
					Username: "john.doe",
					Password: "qwe123456",
				},
				nil,
			},
		},
		{
			"read from repo error",
			"user#12345",
			"#435214",
			&want{
				nil,
				errReadFromRepo,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			creds, err := uc.GetCredentials(context.Background(), tt.user, tt.id)

			assert.Equal(t, tt.want.creds, creds)
			assert.Equal(t, tt.want.err, err)
		})
	}
}

func TestGetAllCredentials(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockKeeperRepository(ctrl)

	m.EXPECT().
		GetSecureRecords(gomock.Any(), gomock.Any(), gomock.Eq(repository.TypeCredentials)).
		Return(nil, nil).
		Times(1)

	m.EXPECT().
		GetSecureRecords(gomock.Any(), gomock.Any(), gomock.Eq(repository.TypeCredentials)).
		Return([]*repository.SecureRecord{
			{
				ID:       "#12345",
				Type:     repository.TypeCredentials,
				Title:    "Gmail",
				Username: "john.doe",
				Password: "qwe123456",
			},
			{
				ID:       "#23456",
				Type:     repository.TypeCredentials,
				Title:    "Yandex",
				Username: "johny.doe",
				Password: "asd123456",
			},
		}, nil).
		Times(1)

	m.EXPECT().
		GetSecureRecords(gomock.Any(), gomock.Any(), gomock.Eq(repository.TypeCredentials)).
		Return(nil, errReadFromRepo).
		Times(1)

	uc := keeper.New(m)

	type want struct {
		creds []*entity.Credentials
		err   error
	}

	tests := []struct {
		name string
		user string
		want *want
	}{
		{
			"missing user",
			"",
			&want{
				nil,
				keeper.ErrMissingArgument,
			},
		},
		{
			"empty list",
			"user#12345",
			&want{
				[]*entity.Credentials{},
				nil,
			},
		},
		{
			"list of credentials",
			"user#12345",
			&want{
				[]*entity.Credentials{
					{
						ID:       "#12345",
						Title:    "Gmail",
						Username: "john.doe",
						Password: "qwe123456",
					},
					{
						ID:       "#23456",
						Title:    "Yandex",
						Username: "johny.doe",
						Password: "asd123456",
					},
				},
				nil,
			},
		},
		{
			"read from repo error",
			"user#12345",
			&want{
				nil,
				errReadFromRepo,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			creds, err := uc.GetAllCredentials(context.Background(), tt.user)

			assert.Equal(t, tt.want.creds, creds)
			assert.Equal(t, tt.want.err, err)
		})
	}
}

func TestUpdateCredentials(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockKeeperRepository(ctrl)

	m.EXPECT().
		UpdateSecureRecord(gomock.Any(), gomock.Any()).
		DoAndReturn(func(_ context.Context, r *repository.SecureRecord) error {
			switch r.ID {
			case "#12345":
				return repository.ErrNotFound

			case "#23456":
				return errWriteToRepo

			case "#34567":
				return repository.ErrVersionConflict

			default:
				return nil
			}
		}).
		AnyTimes()

	uc := keeper.New(m)

	tests := []struct {
		name        string
		Credentials *entity.Credentials
		overwrite   bool
		err         error
	}{
		{
			"empty input",
			nil,
			false,
			keeper.ErrInvalidEntity,
		},
		{
			"invalid entity",
			&entity.Credentials{
				ID: "#12345",
			},
			false,
			keeper.ErrInvalidEntity,
		},
		{
			"invalid entity",
			&entity.Credentials{
				ID:       "#12345",
				Title:    "Gmail",
				Username: "john.doe",
			},
			false,
			keeper.ErrInvalidEntity,
		},
		{
			"invalid entity",
			&entity.Credentials{
				ID:       "#12345",
				Title:    "Gmail",
				Password: "qwe123456",
			},
			false,
			keeper.ErrInvalidEntity,
		},
		{
			"nothing found",
			&entity.Credentials{
				ID:       "#12345",
				Title:    "Gmail",
				Username: "john.doe",
				Password: "qwe123456",
			},
			true,
			keeper.ErrNotFound,
		},
		{
			"version required",
			&entity.Credentials{
				ID:       "#12345",
				Title:    "Gmail",
				Username: "john.doe",
				Password: "qwe123456",
			},
			false,
			keeper.ErrVersionRequired,
		},
		{
			"updated version",
			&entity.Credentials{
				ID:       "#789012",
				Title:    "Gmail",
				Username: "john.doe",
				Password: "qwe123456",
				Version:  10,
			},
			false,
			nil,
		},
		{
			"overwritten",
			&entity.Credentials{
				ID:       "#789012",
				Title:    "Gmail",
				Username: "john.doe",
				Password: "qwe123456",
			},
			true,
			nil,
		},
		{
			"version conflict",
			&entity.Credentials{
				ID:       "#34567",
				Title:    "Gmail",
				Username: "john.doe",
				Password: "qwe123456",
			},
			true,
			keeper.ErrVersionConflict,
		},
		{
			"write to repo error",
			&entity.Credentials{
				ID:       "#23456",
				Title:    "Gmail",
				Username: "john.doe",
				Password: "qwe123456",
			},
			true,
			errWriteToRepo,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := uc.UpdateCredentials(context.Background(), tt.Credentials, tt.overwrite)

			assert.Equal(t, tt.err, err)
		})
	}
}

func TestDeleteCredentials(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockKeeperRepository(ctrl)

	m.EXPECT().
		DeleteSecureRecord(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Eq(repository.TypeCredentials)).
		DoAndReturn(func(_ context.Context, user, id string, _ repository.SecureRecordType) error {
			switch id {
			case "#12345":
				return repository.ErrNotFound

			case "#23456":
				return errWriteToRepo

			default:
				return nil
			}
		}).
		AnyTimes()

	uc := keeper.New(m)

	tests := []struct {
		name string
		user string
		id   string
		err  error
	}{
		{
			"missing id",
			"user#12345",
			"",
			keeper.ErrMissingArgument,
		},
		{
			"missing user",
			"",
			"#12345",
			keeper.ErrMissingArgument,
		},
		{
			"nothing found",
			"user#12345",
			"#12345",
			keeper.ErrNotFound,
		},
		{
			"write to repo error",
			"user#12345",
			"#23456",
			errWriteToRepo,
		},
		{
			"deleted",
			"user#12345",
			"#34567",
			nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := uc.DeleteCredentials(context.Background(), tt.user, tt.id)

			assert.Equal(t, tt.err, err)
		})
	}
}
