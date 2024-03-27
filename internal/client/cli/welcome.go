package cli

import (
	"github.com/Galish/goph-keeper/internal/client/cli/ui"
)

func (a *App) renderWelcomeView() {
	a.ui.Print("Welcome to Goph Keeper!")

	a.ui.Select(
		"Select category",
		[]*ui.SelectOption{
			{
				Label: "Credentials",
				Run:   a.viewAllCredentials,
			},
			{
				Label: "Text notes",
			},
			{
				Label: "Binary notes",
			},
			{
				Label: "Bank cards",
			},
		},
	)
}
