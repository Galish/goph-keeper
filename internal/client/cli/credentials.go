package cli

import (
	"fmt"

	"github.com/Galish/goph-keeper/internal/client/cli/ui"
	"github.com/Galish/goph-keeper/internal/client/entity"
)

func (a *App) viewAllCredentials() {
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
		},
	}

	for i, c := range creds {
		commands = append(
			commands,
			&ui.SelectOption{
				Label: fmt.Sprintf("%d. %s \t %s", i+1, c.Title, c.Description),
				Run: func() {
					a.viewCredentials(c.ID)
				},
			},
		)
	}

	a.ui.Select("Add new credentials or select existing", commands)
}

func (a *App) viewCredentials(id string) {
	creds, err := a.keeper.GetCredentials(id)
	if err != nil {
		a.ui.Error(err)
		return
	}

	a.ui.Print(creds.String())

	handleDelete := func() {
		if ok := a.ui.Confirm("Are you sure"); ok {
			a.keeper.DeleteCredentials(id)
			a.viewAllCredentials()
			return
		} else {
			a.viewCredentials(id)
		}
	}

	handleEdit := func() {
		a.editCredentials(id)
	}

	var commands = []*ui.SelectOption{
		{
			Label: "Edit",
			Run:   handleEdit,
		},
		{
			Label: "Delete",
			Run:   handleDelete,
		},
		{
			Label: "Cancel",
			Run:   a.viewAllCredentials,
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

	if ok := a.ui.Confirm("Add credentials"); ok {
		if err := a.keeper.AddCredentials(&creds); err != nil {
			a.ui.Error(err)
		}
	}

	a.viewAllCredentials()
}

func (a *App) editCredentials(id string) {
	creds := entity.Credentials{
		ID: id,
	}

	creds.Title = a.ui.Input("Title", true)
	creds.Description = a.ui.Input("Description", false)
	creds.Username = a.ui.Input("Username", true)
	creds.Password = a.ui.Input("Password", true)

	a.keeper.UpdateCredentials(&creds)

	a.viewAllCredentials()
}
