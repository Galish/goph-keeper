package cli

import (
	"fmt"
)

func (a *App) renderSignInView() {
	username := a.ui.Input("Enter username")
	password := a.ui.Input("Enter password")

	a.ui.LineBreak()

	if err := a.user.SignIn(username, password); err != nil {
		fmt.Println("An error occured:", err)

		if ok := a.ui.Confirm("Want to try again?"); ok {
			a.renderSignInView()
		}
	}
}
