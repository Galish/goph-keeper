package cli

import (
	"context"
	"syscall"

	"github.com/Galish/goph-keeper/internal/client/cli/ui"
)

func (a *App) viewAuthScreen(ctx context.Context) {
	a.ui.Select(
		"You need to log in or sign up before continuing",
		[]*ui.SelectOption{
			{
				Label: "Already have an account? Log in",
				Run: func() {
					a.signInUser(ctx)
				},
			},
			{
				Label: "Create account",
				Run: func() {
					a.signUpUser(ctx)
				},
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

func (a *App) selectCategory(ctx context.Context) {
	a.ui.Select(
		"Select category",
		[]*ui.SelectOption{
			{
				Label: "Credentials",
				Run: func() {
					a.viewCredentialsList(ctx)
				},
			},
			{
				Label: "Bank cards",
				Run: func() {
					a.viewCardsList(ctx)
				},
			},
			{
				Label: "Text notes",
				Run: func() {
					a.viewTextNotesList(ctx)
				},
			},
			{
				Label: "Binary notes",
				Run: func() {
					a.viewRawNotesList(ctx)
				},
			},
		},
	)
}
