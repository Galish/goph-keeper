package cli

import (
	"os"
)

func (a *App) renderHomeView() {
	a.ui.Select(
		"You need to log in or sign up before continuing",
		[]*SelectOption{
			{
				"Already have an account? Log in",
				a.renderSignInView,
			},
			{
				"Create account",
				a.renderSignUpView,
			},
			{
				"Exit",
				func() {
					os.Exit(0)
				},
			},
		},
	)
}
