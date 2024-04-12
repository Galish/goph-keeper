package app

import "context"

func (a *App) signInUser(ctx context.Context) {
	a.ui.Break()

	username := a.ui.Input("Enter username", true)
	password := a.ui.InputPassword("Enter password", true)

	err := a.user.SignIn(ctx, username, password)
	if ok := a.ui.Retry(err); ok {
		a.signInUser(ctx)
	}

	if err != nil {
		a.viewAuthScreen(ctx)

		return
	}

	a.ui.Print("\nWelcome to Goph Keeper!\n")

	a.selectCategory(ctx)
}

func (a *App) signUpUser(ctx context.Context) {
	a.ui.Break()

	username := a.ui.Input("Enter username", true)
	password := a.ui.InputPassword("Enter password", true)

	err := a.user.SignUp(ctx, username, password)
	if ok := a.ui.Retry(err); ok {
		a.signUpUser(ctx)
	}

	if err != nil {
		a.viewAuthScreen(ctx)

		return
	}

	a.ui.Print("\nWelcome to Goph Keeper!\n")

	a.selectCategory(ctx)
}
