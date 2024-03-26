package cli

import (
	"fmt"
)

func (a *App) renderSignUpView() {
	username := a.ui.Input("Enter username")
	password := a.ui.Input("Enter password")

	err := a.user.SignUp(username, password)
	if err != nil {
		fmt.Println("An error occured:", err)

		if ok := a.ui.Confirm("Want to try again?"); ok {
			a.renderSignUpView()
		}
	}
}
