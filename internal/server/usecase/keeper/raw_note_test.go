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

func TestAddRawNote(t *testing.T) {
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
		name string
		note *entity.RawNote
		err  error
	}{
		{
			"empty input",
			nil,
			keeper.ErrInvalidEntity,
		},
		{
			"invalid entity",
			&entity.RawNote{
				Value: []byte("Hello world!"),
			},
			keeper.ErrInvalidEntity,
		},
		{
			"valid entity",
			&entity.RawNote{
				ID:    "#12345",
				Title: "Binary note",
				Value: []byte("Hello world!"),
			},
			nil,
		},
		{
			"write to repo error",
			&entity.RawNote{
				ID:    "#765432",
				Title: "Text note",
				Value: []byte("Hello world!"),
			},
			errWriteToRepo,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := uc.AddRawNote(context.Background(), tt.note)

			assert.Equal(t, tt.err, err)
		})
	}
}

func TestGetRawNote(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockKeeperRepository(ctrl)

	m.EXPECT().
		GetSecureRecord(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Eq(repository.TypeRawNote)).
		DoAndReturn(func(_ context.Context, user, id string, t repository.SecureRecordType) (*repository.SecureRecord, error) {
			switch id {
			case "#23456":
				return &repository.SecureRecord{
					ID:      "#23456",
					Type:    repository.TypeRawNote,
					Title:   "Binary note",
					RawNote: []byte("Hello world!"),
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
		note *entity.RawNote
		err  error
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
			"valid binary note",
			"user#12345",
			"#23456",
			&want{
				&entity.RawNote{
					ID:    "#23456",
					Title: "Binary note",
					Value: []byte("Hello world!"),
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
			note, err := uc.GetRawNote(context.Background(), tt.user, tt.id)

			assert.Equal(t, tt.want.note, note)
			assert.Equal(t, tt.want.err, err)
		})
	}
}

func TestGetRawNotes(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockKeeperRepository(ctrl)

	m.EXPECT().
		GetSecureRecords(gomock.Any(), gomock.Any(), gomock.Eq(repository.TypeRawNote)).
		Return(nil, nil).
		Times(1)

	m.EXPECT().
		GetSecureRecords(gomock.Any(), gomock.Any(), gomock.Eq(repository.TypeRawNote)).
		Return([]*repository.SecureRecord{
			{
				ID:      "#12345",
				Type:    repository.TypeRawNote,
				Title:   "Binary note",
				RawNote: []byte("Binary note..."),
			},
			{
				ID:      "#23456",
				Type:    repository.TypeRawNote,
				Title:   "Another binary note",
				RawNote: []byte("Another binary note..."),
			},
		}, nil).
		Times(1)

	m.EXPECT().
		GetSecureRecords(gomock.Any(), gomock.Any(), gomock.Eq(repository.TypeRawNote)).
		Return(nil, errReadFromRepo).
		Times(1)

	uc := keeper.New(m)

	type want struct {
		notes []*entity.RawNote
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
				[]*entity.RawNote{},
				nil,
			},
		},
		{
			"list of notes",
			"user#12345",
			&want{
				[]*entity.RawNote{
					{
						ID:    "#12345",
						Title: "Binary note",
						Value: []byte("Binary note..."),
					},
					{
						ID:    "#23456",
						Title: "Another binary note",
						Value: []byte("Another binary note..."),
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
			notes, err := uc.GetRawNotes(context.Background(), tt.user)

			assert.Equal(t, tt.want.notes, notes)
			assert.Equal(t, tt.want.err, err)
		})
	}
}

func TestUpdateRawNote(t *testing.T) {
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
		name      string
		RawNote   *entity.RawNote
		overwrite bool
		err       error
	}{
		{
			"empty input",
			nil,
			false,
			keeper.ErrInvalidEntity,
		},
		{
			"invalid entity",
			&entity.RawNote{
				ID: "#12345",
			},
			false,
			keeper.ErrInvalidEntity,
		},
		{
			"invalid entity",
			&entity.RawNote{
				ID:    "#12345",
				Title: "Binary note",
			},
			false,
			keeper.ErrInvalidEntity,
		},
		{
			"invalid entity",
			&entity.RawNote{
				ID:    "#12345",
				Value: []byte("Hello world!"),
			},
			false,
			keeper.ErrInvalidEntity,
		},
		{
			"nothing found",
			&entity.RawNote{
				ID:    "#12345",
				Title: "Binary note",
				Value: []byte("Hello world!"),
			},
			true,
			keeper.ErrNotFound,
		},
		{
			"version required",
			&entity.RawNote{
				ID:    "#12345",
				Title: "Binary note",
				Value: []byte("Hello world!"),
			},
			false,
			keeper.ErrVersionRequired,
		},
		{
			"updated version",
			&entity.RawNote{
				ID:      "#789012",
				Title:   "Binary note",
				Value:   []byte("Hello world!"),
				Version: 10,
			},
			false,
			nil,
		},
		{
			"overwritten",
			&entity.RawNote{
				ID:    "#789012",
				Title: "Binary note",
				Value: []byte("Hello world!"),
			},
			true,
			nil,
		},
		{
			"version conflict",
			&entity.RawNote{
				ID:    "#34567",
				Title: "Binary note",
				Value: []byte("Hello world!"),
			},
			true,
			keeper.ErrVersionConflict,
		},
		{
			"write to repo error",
			&entity.RawNote{
				ID:    "#23456",
				Title: "Binary note",
				Value: []byte("Hello world!"),
			},
			true,
			errWriteToRepo,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := uc.UpdateRawNote(context.Background(), tt.RawNote, tt.overwrite)

			assert.Equal(t, tt.err, err)
		})
	}
}

func TestDeleteRawNote(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockKeeperRepository(ctrl)

	m.EXPECT().
		DeleteSecureRecord(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Eq(repository.TypeRawNote)).
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
			err := uc.DeleteRawNote(context.Background(), tt.user, tt.id)

			assert.Equal(t, tt.err, err)
		})
	}
}
