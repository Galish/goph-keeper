package cli

import (
	"fmt"

	"github.com/Galish/goph-keeper/internal/client/cli/ui"
	"github.com/Galish/goph-keeper/internal/client/entity"
)

func (a *App) viewCredentialsList() {
	creds, err := a.keeper.GetCredentialsList()
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
		id := c.ID

		commands = append(
			commands,
			&ui.SelectOption{
				Label: fmt.Sprintf("%d. %s \t %s", i+1, c.Title, c.Description),
				Run: func() {
					a.viewCredentials(id)
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
			if err := a.keeper.DeleteCredentials(id); err != nil {
				a.ui.Error(err)
			}

			a.viewCredentialsList()
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
			Run:   a.viewCredentialsList,
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

	a.viewCredentialsList()
}

func (a *App) editCredentials(id string) {
	creds := &entity.Credentials{
		ID: id,
	}

	creds.Title = a.ui.Input("Title", true)
	creds.Description = a.ui.Input("Description", false)
	creds.Username = a.ui.Input("Username", true)
	creds.Password = a.ui.Input("Password", true)

	a.keeper.UpdateCredentials(creds)

	a.viewCredentialsList()
}
