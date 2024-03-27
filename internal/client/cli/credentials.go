package cli

import (
	"fmt"

	"github.com/Galish/goph-keeper/internal/client/cli/ui"
	"github.com/Galish/goph-keeper/internal/client/entity"
)

func (a *App) renderAllCredentials() {
	creds, err := a.keeper.GetAllCredentials()
	if err != nil {
		a.ui.Error(err)
		return
	}

	commands := []*ui.SelectOption{
		{
			Label: "+ Add new",
			Run:   a.addCredentials,
		},
		{
			Label: "  Cancel",
			Run:   a.renderHomeView,
		},
	}

	for i, c := range creds {
		commands = append(
			commands,
			&ui.SelectOption{
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
		a.ui.Error(err)
		return
	}

	a.ui.Print(creds.String())

	handleDelete := func() {
		if ok := a.ui.Confirm("Are you sure"); ok {
			a.keeper.DeleteCredentials(id)
			a.renderAllCredentials()
			return
		} else {
			a.renderCredentials(id)
		}
	}

	var commands = []*ui.SelectOption{
		{
			Label: "Edit",
			Run:   a.editCredentials,
		},
		{
			Label: "Delete",
			Run:   handleDelete,
		},
		{
			Label: "Cancel",
			Run:   a.renderAllCredentials,
		},
	}

	a.ui.Select("Select action", commands)
}

func (a *App) addCredentials() {
	creds := entity.Credentials{}

	creds.Title = a.ui.Input("Title", true)
	creds.Description = a.ui.Input("Description", false)
	creds.Username = a.ui.Input("Username", true)
	creds.Password = a.ui.Input("Password", true)

	ok := a.ui.Confirm("Add credentials")
	if ok {
		fmt.Println("-save-")
	}

	a.renderAllCredentials()
}

func (a *App) editCredentials() {
	creds := entity.Credentials{}

	creds.Title = a.ui.Input("Title", true)
	creds.Description = a.ui.Input("Description", false)
	creds.Username = a.ui.Input("Username", true)
	creds.Password = a.ui.Input("Password", true)

	fmt.Printf(
		"Edit credentials\nTitle: %s\nDescription: %s\nUsername: %s\nPassword: %s\n\n",
		creds.Title,
		creds.Description,
		creds.Username,
		creds.Password,
	)

	a.renderAllCredentials()
}
