package cli

import (
	"os"

	"github.com/Galish/goph-keeper/internal/client/cli/ui"
)

func (a *App) viewHomeScreen() {
	a.ui.Select(
		"You need to log in or sign up before continuing",
		[]*ui.SelectOption{
			{
				Label: "Already have an account? Log in",
				Run:   a.signInUser,
			},
			{
				Label: "Create account",
				Run:   a.signUpUser,
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

func (a *App) selectCategory() {
	a.ui.Print("Welcome to Goph Keeper!")

	a.ui.Select(
		"Select category",
		[]*ui.SelectOption{
			{
				Label: "Credentials",
				Run:   a.viewCredentialsList,
			},
			{
				Label: "Bank cards",
				Run:   a.viewCardsList,
			},
			{
				Label: "Text notes",
				Run:   a.viewTextNotesList,
			},
			{
				Label: "Binary notes",
			},
		},
	)
}
