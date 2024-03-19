package keeper_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"gotest.tools/assert"

	"github.com/Galish/goph-keeper/internal/server/entity"
	"github.com/Galish/goph-keeper/internal/server/repository"
	"github.com/Galish/goph-keeper/internal/server/repository/mocks"
	"github.com/Galish/goph-keeper/internal/server/usecase/keeper"
)

var errWriteToRepo = errors.New("failed to write to repo")

func TestAddNote(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockKeeperRepository(ctrl)

	m.EXPECT().
		CreateSecureRecord(gomock.Any(), gomock.Any()).
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
		note *entity.Note
		err  error
	}{
		{
			"empty input",
			nil,
			keeper.ErrInvalidEntity,
		},
		{
			"invalid entity",
			&entity.Note{
				Value: "Text note...",
			},
			keeper.ErrInvalidEntity,
		},
		{
			"text note",
			&entity.Note{
				ID:    "#12345",
				Title: "Text note",
				Value: "Text note...",
			},
			nil,
		},
		{
			"binary note",
			&entity.Note{
				ID:       "#12345",
				Title:    "Binary note",
				RawValue: []byte("Binary note..."),
			},
			nil,
		},
		{
			"write to repo error",
			&entity.Note{
				ID:    "#765432",
				Title: "Text note",
				Value: "Text note...",
			},
			errWriteToRepo,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := uc.AddNote(context.Background(), tt.note)

			assert.Equal(t, tt.err, err)
		})
	}
}

func TestAddCard(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockKeeperRepository(ctrl)

	m.EXPECT().
		CreateSecureRecord(gomock.Any(), gomock.Any()).
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
		card *entity.Card
		err  error
	}{
		{
			"empty input",
			nil,
			keeper.ErrInvalidEntity,
		},
		{
			"invalid entity",
			&entity.Card{
				ID:     "#12345",
				Number: "1234 5678 9012 4453",
			},
			keeper.ErrInvalidEntity,
		},
		{
			"valid card",
			&entity.Card{
				ID:     "#12345",
				Title:  "Credit card",
				Number: "1234 5678 9012 4453",
				Holder: "John Daw",
				CVC:    "123",
				Expiry: time.Date(2025, time.Month(6), 11, 0, 0, 0, 0, time.UTC),
			},
			nil,
		},
		{
			"write to repo error",
			&entity.Card{
				ID:     "#765432",
				Title:  "Credit card",
				Number: "1234 5678 9012 4453",
				Holder: "John Daw",
				CVC:    "123",
				Expiry: time.Date(2025, time.Month(6), 11, 0, 0, 0, 0, time.UTC),
			},
			errWriteToRepo,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := uc.AddCard(context.Background(), tt.card)

			assert.Equal(t, tt.err, err)
		})
	}
}

func TestAddCredentials(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockKeeperRepository(ctrl)

	m.EXPECT().
		CreateSecureRecord(gomock.Any(), gomock.Any()).
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
			"valid card",
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
