package keeper_test

import (
	"context"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"gotest.tools/assert"

	"github.com/Galish/goph-keeper/internal/server/entity"
	"github.com/Galish/goph-keeper/internal/server/repository"
	"github.com/Galish/goph-keeper/internal/server/repository/mocks"
	"github.com/Galish/goph-keeper/internal/server/usecase/keeper"
)

func TestGetTextNotes(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockKeeperRepository(ctrl)

	m.EXPECT().
		GetSecureRecords(gomock.Any(), gomock.Any(), gomock.Eq(repository.TypeNote)).
		Return(nil, nil).
		Times(1)

	m.EXPECT().
		GetSecureRecords(gomock.Any(), gomock.Any(), gomock.Eq(repository.TypeNote)).
		Return([]*repository.SecureRecord{
			{
				ID:       "#12345",
				Type:     repository.TypeNote,
				Title:    "Text note",
				TextNote: "Text note...",
			},
			{
				ID:       "#23456",
				Type:     repository.TypeNote,
				Title:    "Another text note",
				TextNote: "Another text note...",
			},
		}, nil).
		Times(1)

	m.EXPECT().
		GetSecureRecords(gomock.Any(), gomock.Any(), gomock.Eq(repository.TypeNote)).
		Return(nil, errReadFromRepo).
		Times(1)

	uc := keeper.New(m)

	type want struct {
		notes []*entity.Note
		err   error
	}

	tests := []struct {
		name string
		user string
		want *want
	}{
		{
			"empty list",
			"user#12345",
			&want{
				[]*entity.Note{},
				nil,
			},
		},
		{
			"list of notes",
			"user#12345",
			&want{
				[]*entity.Note{
					{
						ID:    "#12345",
						Title: "Text note",
						Value: "Text note...",
					},
					{
						ID:    "#23456",
						Title: "Another text note",
						Value: "Another text note...",
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
			notes, err := uc.GetTextNotes(context.Background(), tt.user)

			assert.DeepEqual(t, tt.want.notes, notes)
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
		notes []*entity.Note
		err   error
	}

	tests := []struct {
		name string
		user string
		want *want
	}{
		{
			"empty list",
			"user#12345",
			&want{
				[]*entity.Note{},
				nil,
			},
		},
		{
			"list of notes",
			"user#12345",
			&want{
				[]*entity.Note{
					{
						ID:       "#12345",
						Title:    "Binary note",
						RawValue: []byte("Binary note..."),
					},
					{
						ID:       "#23456",
						Title:    "Another binary note",
						RawValue: []byte("Another binary note..."),
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

			assert.DeepEqual(t, tt.want.notes, notes)
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
			"empty list",
			"user#12345",
			&want{
				[]*entity.Card{},
				nil,
			},
		},
		{
			"list of notes",
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
			"empty list",
			"user#12345",
			&want{
				[]*entity.Credentials{},
				nil,
			},
		},
		{
			"list of notes",
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

			assert.DeepEqual(t, tt.want.creds, creds)
			assert.Equal(t, tt.want.err, err)
		})
	}
}
