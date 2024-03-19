package user_test

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"gotest.tools/assert"

	"github.com/Galish/goph-keeper/internal/server/entity"
	"github.com/Galish/goph-keeper/internal/server/repository/mocks"
	"github.com/Galish/goph-keeper/internal/server/usecase/user"
)

func TestSignIn(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockUserRepository(ctrl)

	m.EXPECT().
		GetUserByLogin(gomock.Any(), gomock.Any()).
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
		name  string
		creds *entity.User
		want  *want
	}{
		{
			"empty input",
			nil,
			&want{
				"",
				user.ErrMissingCredentials,
			},
		},
		{
			"missing username",
			&entity.User{
				Login:    "",
				Password: "qwe123456",
			},
			&want{
				"",
				user.ErrMissingCredentials,
			},
		},
		{
			"missing password",
			&entity.User{
				Login:    "john.doe",
				Password: "",
			},
			&want{
				"",
				user.ErrMissingCredentials,
			},
		},
		{
			"valid credentials",
			&entity.User{
				Login:    "john.doe",
				Password: "qwe123456",
			},
			&want{
				"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOiIjMTIzNDUifQ.FnPcRyLXm11AqObgLd1HR-OB7FmsPtcbsUg31IUW6Ss",
				nil,
			},
		},
		{
			"invalid credentials",
			&entity.User{
				Login:    "john.doe",
				Password: "qwe12345678",
			},
			&want{
				"",
				user.ErrInvalidCredentials,
			},
		},
		{
			"write to repo error",
			&entity.User{
				Login:    "johny.doe",
				Password: "qwe123456",
			},
			&want{
				"",
				errWriteToRepo,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			token, err := uc.SignIn(context.Background(), tt.creds)

			assert.Equal(t, tt.want.token, token)
			assert.Equal(t, tt.want.err, err)
		})
	}
}
