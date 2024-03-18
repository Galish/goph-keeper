package user_test

import (
	"context"
	"testing"

	"github.com/Galish/goph-keeper/internal/server/entity"
	"github.com/Galish/goph-keeper/internal/server/repository/mocks"
	"github.com/Galish/goph-keeper/internal/server/usecase/user"
	"github.com/golang/mock/gomock"
	"gotest.tools/assert"
)

func TestSignIn(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockUserRepository(ctrl)

	m.EXPECT().
		GetByLogin(gomock.Any(), gomock.Any()).
		DoAndReturn(func(_ context.Context, username string) (*entity.User, error) {
			switch username {
			case "john.doe":
				return &entity.User{
						ID:       "#12345",
						Login:    "john.doe",
						Password: "$2a$10$3S997zQF4Fh2MSmo5gIdwu7OUg3Q21WXe77dgmJGTxNYU7Y/rAdtK",
					},
					nil
			default:
				return nil, errWriteToRepo
			}
		}).
		AnyTimes()

	uc := user.New(m, "secret_key")

	type want struct {
		token string
		err   error
	}

	tests := []struct {
		name     string
		username string
		password string
		want     *want
	}{
		{
			"empty input",
			"",
			"",
			&want{
				"",
				user.ErrMissingCredentials,
			},
		},
		{
			"missing username",
			"",
			"qwe123456",
			&want{
				"",
				user.ErrMissingCredentials,
			},
		},
		{
			"missing password",
			"john.doe",
			"",
			&want{
				"",
				user.ErrMissingCredentials,
			},
		},
		{
			"valid credentials",
			"john.doe",
			"qwe123456",
			&want{
				"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOiIjMTIzNDUifQ.FnPcRyLXm11AqObgLd1HR-OB7FmsPtcbsUg31IUW6Ss",
				nil,
			},
		},
		{
			"invalid credentials",
			"john.doe",
			"qwe12345678",
			&want{
				"",
				user.ErrInvalidCredentials,
			},
		},
		{
			"write to repo error",
			"johny.doe",
			"qwe123456",
			&want{
				"",
				errWriteToRepo,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			token, err := uc.SignIn(context.Background(), tt.username, tt.password)

			assert.Equal(t, tt.want.token, token)
			assert.Equal(t, tt.want.err, err)
		})
	}
}
