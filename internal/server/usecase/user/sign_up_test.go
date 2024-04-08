package user_test

import (
	"context"
	"errors"
	"regexp"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/Galish/goph-keeper/internal/entity"
	"github.com/Galish/goph-keeper/internal/server/repository"
	"github.com/Galish/goph-keeper/internal/server/repository/mocks"
	"github.com/Galish/goph-keeper/internal/server/usecase/user"
	"github.com/Galish/goph-keeper/pkg/auth"
)

const secretKey = "secret_key"

var errWriteToRepo = errors.New("failed to write to repo")

func TestSignUp(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockUserRepository(ctrl)

	m.EXPECT().
		AddUser(gomock.Any(), gomock.Any()).
		DoAndReturn(func(_ context.Context, user *entity.User) error {
			switch user.Login {
			case "john.doe":
				return nil

			case "johnn.doe":
				return repository.ErrConflict

			default:
				return errWriteToRepo
			}
		}).
		AnyTimes()

	uc := user.New(m, auth.NewJWTManager(secretKey))

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
			"missing credentials",
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
			"user already exists",
			"johnn.doe",
			"qwe123456",
			&want{
				"",
				user.ErrConflict,
			},
		},
		{
			"valid credentials",
			"john.doe",
			"qwe123456",
			&want{
				"",
				nil,
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
			token, err := uc.SignUp(context.Background(), tt.username, tt.password)

			if tt.want.err != nil {
				assert.Equal(t, err, tt.want.err)
				return
			}

			assert.Regexp(
				t,
				regexp.MustCompile("^[A-Za-z0-9-_]*.[A-Za-z0-9-_]*.[A-Za-z0-9-_]*$"),
				token,
			)
			assert.NoError(t, err)
		})
	}
}
