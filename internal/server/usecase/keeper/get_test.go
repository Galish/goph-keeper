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

var errReadFromRepo = errors.New("failed to read from repo")

func TestGetNote(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockKeeperRepository(ctrl)

	m.EXPECT().
		GetSecureRecord(gomock.Any(), gomock.Any(), gomock.Any()).
		DoAndReturn(func(_ context.Context, user, id string) (*repository.SecureRecord, error) {
			switch id {
			case "#12345":
				return &repository.SecureRecord{
					ID:       "#12345",
					Type:     repository.TypeNote,
					Title:    "Text note",
					TextNote: "Text note...",
				}, nil

			case "#23456":
				return &repository.SecureRecord{
					ID:      "#23456",
					Type:    repository.TypeRawNote,
					Title:   "Binary note",
					RawNote: []byte("Binary note..."),
				}, nil

			case "#34567":
				return &repository.SecureRecord{}, nil

			default:
				return nil, errReadFromRepo
			}
		}).
		AnyTimes()

	uc := keeper.New(m)

	type want struct {
		note *entity.Note
		err  error
	}

	tests := []struct {
		name string
		user string
		id   string
		want *want
	}{
		{
			"invalid type error",
			"user#12345",
			"#34567",
			&want{
				nil,
				keeper.ErrInvalidType,
			},
		},
		{
			"text note",
			"user#12345",
			"#12345",
			&want{
				&entity.Note{
					ID:    "#12345",
					Title: "Text note",
					Value: "Text note...",
				},
				nil,
			},
		},
		{
			"binary note",
			"user#12345",
			"#23456",
			&want{
				&entity.Note{
					ID:       "#23456",
					Title:    "Binary note",
					RawValue: []byte("Binary note..."),
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
			note, err := uc.GetNote(context.Background(), tt.user, tt.id)

			assert.DeepEqual(t, tt.want.note, note)
			assert.Equal(t, tt.want.err, err)
		})
	}
}

func TestGetCard(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockKeeperRepository(ctrl)

	m.EXPECT().
		GetSecureRecord(gomock.Any(), gomock.Any(), gomock.Any()).
		DoAndReturn(func(_ context.Context, user, id string) (*repository.SecureRecord, error) {
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
				return &repository.SecureRecord{}, nil

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
			"invalid type error",
			"user#12345",
			"#34567",
			&want{
				nil,
				keeper.ErrInvalidType,
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

func TestGetCredentials(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockKeeperRepository(ctrl)

	m.EXPECT().
		GetSecureRecord(gomock.Any(), gomock.Any(), gomock.Any()).
		DoAndReturn(func(_ context.Context, user, id string) (*repository.SecureRecord, error) {
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
				return &repository.SecureRecord{}, nil

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
			"invalid type error",
			"user#12345",
			"#34567",
			&want{
				nil,
				keeper.ErrInvalidType,
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

			assert.DeepEqual(t, tt.want.creds, creds)
			assert.Equal(t, tt.want.err, err)
		})
	}
}
