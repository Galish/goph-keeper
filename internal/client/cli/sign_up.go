package cli

func (a *App) renderSignUpView() {
	username := a.ui.Input("Enter username", true)
	password := a.ui.Input("Enter password", true)

	err := a.user.SignUp(username, password)
	if err != nil {
		a.ui.Error(err)

		if ok := a.ui.Confirm("Want to try again?"); ok {
			a.renderSignUpView()
		}
	}
}
