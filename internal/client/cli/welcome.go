package cli

func (a *App) renderWelcomeView() {
	a.ui.Display("Welcome to Goph Keeper!")

	a.ui.Select(
		"Select category",
		[]*SelectOption{
			{
				"Credentials",
				a.renderCredentialsOverview,
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
