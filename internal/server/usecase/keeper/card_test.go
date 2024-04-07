package keeper_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/Galish/goph-keeper/internal/entity"
	"github.com/Galish/goph-keeper/internal/server/repository"
	"github.com/Galish/goph-keeper/internal/server/repository/mocks"
	"github.com/Galish/goph-keeper/internal/server/usecase/keeper"
	"github.com/golang/mock/gomock"
	"gotest.tools/assert"
)

var (
	errWriteToRepo  = errors.New("failed to write to repo")
	errReadFromRepo = errors.New("failed to read from repo")
)

func TestAddCard(t *testing.T) {
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
				ID: "#12345",
			},
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
			"invalid entity",
			&entity.Card{
				ID:     "#12345",
				Holder: "John Daw",
			},
			keeper.ErrInvalidEntity,
		},
		{
			"invalid entity",
			&entity.Card{
				ID:  "#12345",
				CVC: "123",
			},
			keeper.ErrInvalidEntity,
		},
		{
			"invalid entity",
			&entity.Card{
				ID:     "#12345",
				Expiry: time.Date(2025, time.Month(6), 11, 0, 0, 0, 0, time.UTC),
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

func TestGetCard(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockKeeperRepository(ctrl)

	m.EXPECT().
		GetSecureRecord(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Eq(repository.TypeCard)).
		DoAndReturn(func(_ context.Context, user, id string, t repository.SecureRecordType) (*repository.SecureRecord, error) {
			switch id {
			case "#12345":
				return &repository.SecureRecord{
					ID:         "#12345",
					Type:       repository.TypeCard,
					Title:      "Credit card",
					CardNumber: "1234 5678 9012 4453",
					CardHolder: "John Daw",
					CardCVC:    "123",
					CardExpiry: time.Date(2025, time.Month(6), 11, 0, 0, 0, 0, time.UTC),
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
		card *entity.Card
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
			"valid credentials",
			"user#12345",
			"#12345",
			&want{
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
			card, err := uc.GetCard(context.Background(), tt.user, tt.id)

			assert.DeepEqual(t, tt.want.card, card)
			assert.Equal(t, tt.want.err, err)
		})
	}
}

func TestGetCards(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockKeeperRepository(ctrl)

	m.EXPECT().
		GetSecureRecords(gomock.Any(), gomock.Any(), gomock.Eq(repository.TypeCard)).
		Return(nil, nil).
		Times(1)

	m.EXPECT().
		GetSecureRecords(gomock.Any(), gomock.Any(), gomock.Eq(repository.TypeCard)).
		Return([]*repository.SecureRecord{
			{
				ID:         "#12345",
				Type:       repository.TypeCard,
				Title:      "Credit card",
				CardNumber: "1234 5678 9012 4453",
				CardHolder: "John Daw",
				CardCVC:    "123",
				CardExpiry: time.Date(2025, time.Month(6), 11, 0, 0, 0, 0, time.UTC),
			},
			{
				ID:         "#23456",
				Type:       repository.TypeCard,
				Title:      "Debit card",
				CardNumber: "1234 5678 9012 4453",
				CardHolder: "Johny Deep",
				CardCVC:    "234",
				CardExpiry: time.Date(2027, time.Month(6), 11, 0, 0, 0, 0, time.UTC),
			},
		}, nil).
		Times(1)

	m.EXPECT().
		GetSecureRecords(gomock.Any(), gomock.Any(), gomock.Eq(repository.TypeCard)).
		Return(nil, errReadFromRepo).
		Times(1)

	uc := keeper.New(m)

	type want struct {
		cards []*entity.Card
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
				[]*entity.Card{},
				nil,
			},
		},
		{
			"list of cards",
			"user#12345",
			&want{
				[]*entity.Card{
					{
						ID:     "#12345",
						Title:  "Credit card",
						Number: "1234 5678 9012 4453",
						Holder: "John Daw",
						CVC:    "123",
						Expiry: time.Date(2025, time.Month(6), 11, 0, 0, 0, 0, time.UTC),
					},
					{
						ID:     "#23456",
						Title:  "Debit card",
						Number: "1234 5678 9012 4453",
						Holder: "Johny Deep",
						CVC:    "234",
						Expiry: time.Date(2027, time.Month(6), 11, 0, 0, 0, 0, time.UTC),
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
			cards, err := uc.GetCards(context.Background(), tt.user)

			assert.DeepEqual(t, tt.want.cards, cards)
			assert.Equal(t, tt.want.err, err)
		})
	}
}

func TestUpdateCard(t *testing.T) {
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
		card      *entity.Card
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
			&entity.Card{
				ID: "#12345",
			},
			false,
			keeper.ErrInvalidEntity,
		},
		{
			"invalid entity",
			&entity.Card{
				ID:     "#12345",
				Number: "1234 5678 9012 4453",
			},
			false,
			keeper.ErrInvalidEntity,
		},
		{
			"invalid entity",
			&entity.Card{
				ID:     "#12345",
				Holder: "John Daw",
			},
			false,
			keeper.ErrInvalidEntity,
		},
		{
			"invalid entity",
			&entity.Card{
				ID:  "#12345",
				CVC: "123",
			},
			false,
			keeper.ErrInvalidEntity,
		},
		{
			"invalid entity",
			&entity.Card{
				ID:     "#12345",
				Expiry: time.Date(2025, time.Month(6), 11, 0, 0, 0, 0, time.UTC),
			},
			false,
			keeper.ErrInvalidEntity,
		},
		{
			"nothing found",
			&entity.Card{
				ID:     "#12345",
				Title:  "Credit card",
				Number: "1234 5678 9012 4453",
				Holder: "John Daw",
				CVC:    "123",
				Expiry: time.Date(2025, time.Month(6), 11, 0, 0, 0, 0, time.UTC),
			},
			true,
			keeper.ErrNotFound,
		},
		{
			"version required",
			&entity.Card{
				ID:     "#12345",
				Title:  "Credit card",
				Number: "1234 5678 9012 4453",
				Holder: "John Daw",
				CVC:    "123",
				Expiry: time.Date(2025, time.Month(6), 11, 0, 0, 0, 0, time.UTC),
			},
			false,
			keeper.ErrVersionRequired,
		},
		{
			"updated version",
			&entity.Card{
				ID:      "#789012",
				Title:   "Credit card",
				Number:  "1234 5678 9012 4453",
				Holder:  "John Daw",
				CVC:     "123",
				Expiry:  time.Date(2025, time.Month(6), 11, 0, 0, 0, 0, time.UTC),
				Version: 10,
			},
			false,
			nil,
		},
		{
			"overwritten",
			&entity.Card{
				ID:     "#789012",
				Title:  "Credit card",
				Number: "1234 5678 9012 4453",
				Holder: "John Daw",
				CVC:    "123",
				Expiry: time.Date(2025, time.Month(6), 11, 0, 0, 0, 0, time.UTC),
			},
			true,
			nil,
		},
		{
			"version conflict",
			&entity.Card{
				ID:     "#34567",
				Title:  "Credit card",
				Number: "1234 5678 9012 4453",
				Holder: "John Daw",
				CVC:    "123",
				Expiry: time.Date(2025, time.Month(6), 11, 0, 0, 0, 0, time.UTC),
			},
			true,
			keeper.ErrVersionConflict,
		},
		{
			"write to repo error",
			&entity.Card{
				ID:     "#23456",
				Title:  "Credit card",
				Number: "1234 5678 9012 4453",
				Holder: "John Daw",
				CVC:    "123",
				Expiry: time.Date(2025, time.Month(6), 11, 0, 0, 0, 0, time.UTC),
			},
			true,
			errWriteToRepo,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := uc.UpdateCard(context.Background(), tt.card, tt.overwrite)

			assert.Equal(t, tt.err, err)
		})
	}
}

func TestDeleteCard(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockKeeperRepository(ctrl)

	m.EXPECT().
		DeleteSecureRecord(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Eq(repository.TypeCard)).
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
			err := uc.DeleteCard(context.Background(), tt.user, tt.id)

			assert.Equal(t, tt.err, err)
		})
	}
}
