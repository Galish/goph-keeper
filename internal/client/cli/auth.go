package cli

func (a *App) signInUser() {
	username := a.ui.Input("Enter username", true)
	password := a.ui.InputPassword("Enter password", true)

	err := a.user.SignIn(username, password)
	if ok := a.ui.Retry(err); ok {
		a.signInUser()
	}

	if err != nil {
		a.viewAuthScreen()
		return
	}

	a.ui.Print("\nWelcome to Goph Keeper!\n\n")
	a.selectCategory()
}

func (a *App) signUpUser() {
	username := a.ui.Input("Enter username", true)
	password := a.ui.InputPassword("Enter password", true)

	err := a.user.SignUp(username, password)
	if ok := a.ui.Retry(err); ok {
		a.signUpUser()
	}

	if err != nil {
		a.viewAuthScreen()
		return
	}

	a.ui.Print("\nWelcome to Goph Keeper!\n\n")
	a.selectCategory()
}
