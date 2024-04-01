package cli

import (
	"syscall"

	"github.com/Galish/goph-keeper/internal/client/cli/ui"
)

func (a *App) viewAuthScreen() {
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
					syscall.Kill(syscall.Getpid(), syscall.SIGINT)
				},
			},
		},
	)
}

func (a *App) selectCategory() {
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
				Run:   a.viewRawNotesList,
			},
		},
	)
}
