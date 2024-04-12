package app

import (
	"context"

	"github.com/Galish/goph-keeper/internal/client/app/cli"
)

func (a *App) viewAuthScreen(ctx context.Context) {
	a.ui.Break()

	a.ui.Select(
		"You need to log in or sign up before continuing",
		[]*cli.SelectOption{
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
				Run:   a.ui.Exit,
			},
		},
	)
}

func (a *App) selectCategory(ctx context.Context) {
	a.ui.Break()

	a.ui.Select(
		"Select category",
		[]*cli.SelectOption{
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
