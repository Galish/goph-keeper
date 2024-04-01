package cli

func (a *App) signInUser() {
	username := a.ui.Input("Enter username", true)
	password := a.ui.InputPassword("Enter password", true)

	if err := a.user.SignIn(username, password); err != nil {
		a.ui.Error(err)
		a.ui.Break()

		if ok := a.ui.Confirm("Want to try again"); ok {
			a.signInUser()
		}
	}
}

func (a *App) signUpUser() {
	username := a.ui.Input("Enter username", true)
	password := a.ui.InputPassword("Enter password", true)

	err := a.user.SignUp(username, password)
	if err != nil {
		a.ui.Error(err)

		if ok := a.ui.Confirm("Want to try again"); ok {
			a.signUpUser()
		}
	}
}
