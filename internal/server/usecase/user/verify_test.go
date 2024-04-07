package user_test

/*
func TestVerify(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockUserRepository(ctrl)

	uc := user.New(m, auth.NewJWTManager(secretKey))

	type want struct {
		user *entity.User
		err  error
	}

	tests := []struct {
		name  string
		token string
		want  *want
	}{
		// {
		// 	"missing credentials",
		// 	"",
		// 	"",
		// 	user.ErrMissingCredentials,
		// },
		// {
		// 	"missing username",
		// 	"",
		// 	"qwe123456",
		// 	user.ErrMissingCredentials,
		// },
		// {
		// 	"missing password",
		// 	"john.doe",
		// 	"",
		// 	user.ErrMissingCredentials,
		// },
		// {
		// 	"user already exists",
		// 	"johnn.doe",
		// 	"qwe123456",
		// 	user.ErrConflict,
		// },
		// {
		// 	"valid credentials",
		// 	"john.doe",
		// 	"qwe123456",
		// 	nil,
		// },
		// {
		// 	"write to repo error",
		// 	"johny.doe",
		// 	"qwe123456",
		// 	errWriteToRepo,
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			token, err := uc.Verify(tt.token)

			assert.Equal(t, tt.err, err)

			if tt.err == nil {
				_, err := uc.Verify(token)

				require.NoError(t, err)
			}
		})
	}
}
*/
