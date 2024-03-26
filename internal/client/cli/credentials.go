package cli

import (
	"fmt"
)

func (a *App) renderCredentialsOverview() {
	creds, err := a.keeper.GetAllCredentials()
	if err != nil {
		fmt.Println("Err:", err)
		return
	}

	commands := []*SelectOption{
		{
			Label: "+ Add new",
			Run: func() {
				fmt.Println("___add new___")
			},
		},
		{
			Label: "  Cancel",
			Run:   a.renderHomeView,
		},
	}

	for i, c := range creds {
		commands = append(
			commands,
			&SelectOption{
				Label: fmt.Sprintf("%d. %s \t %s", i+1, c.Title, c.Description),
				Run: func() {
					a.renderCredentials(c.ID)
				},
			},
		)
	}

	a.ui.Select("Add new credentials or select existing", commands)
}

func (a *App) renderCredentials(id string) {
	creds, err := a.keeper.GetCredentials(id)
	if err != nil {
		fmt.Println("Err:", err)
		return
	}

	fmt.Printf(
		"View credentials\nTitle: %s\nDescription: %s\nUsername: %s\nPassword: %s\n\n",
		creds.Title,
		creds.Description,
		creds.Username,
		creds.Password,
	)

	var commands = []*SelectOption{
		{
			Label: "Edit",
			Run: func() {
				fmt.Println("___edit___")
			},
		},
		{
			Label: "Delete",
			Run: func() {
				if ok := a.ui.Confirm("Are you sure?"); ok {
					a.keeper.DeleteCredentials(id)
					a.renderCredentialsOverview()
				} else {
					a.renderCredentials(id)
				}
			},
		},
		{
			Label: "Cancel",
			Run:   a.renderCredentialsOverview,
		},
	}

	a.ui.Select("Select action", commands)
}
