package cli

import (
	"os"

	"github.com/Galish/goph-keeper/internal/client/cli/ui"
)

func (a *App) renderHomeView() {
	a.ui.Select(
		"You need to log in or sign up before continuing",
		[]*ui.SelectOption{
			{
				Label: "Already have an account? Log in",
				Run:   a.renderSignInView,
			},
			{
				Label: "Create account",
				Run:   a.renderSignUpView,
			},
			{
				Label: "Exit",
				Run: func() {
					os.Exit(0)
				},
			},
		},
	)
}
