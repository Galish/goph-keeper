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

	var commands = []*ui.SelectOption{
		{
			Label: "Edit",
			Run: func() {
				a.editCredentials(id)
			},
		},
		{
			Label: "Delete",
			Run: func() {
				a.deleteCredentials(id)
			},
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
		for {
			err := a.keeper.AddCredentials(&creds)
			if ok := a.ui.Retry(err); !ok {
				break
			}
		}
	}

	a.viewCredentialsList()
}

func (a *App) editCredentials(id string) {
	creds, err := a.keeper.GetCredentials(id)
	if err != nil {
		a.ui.Error(err)
		return
	}

	var updated = &entity.Credentials{
		ID: id,
	}

	updated.Title = a.ui.Edit("Title", creds.Title, true)
	updated.Description = a.ui.Edit("Description", creds.Description, false)
	updated.Username = a.ui.Edit("Username", creds.Username, true)
	updated.Password = a.ui.Edit("Password", creds.Password, true)

	if ok := a.ui.Confirm("Update credentials"); ok {
		for {
			err := a.keeper.UpdateCredentials(updated)
			if ok := a.ui.Retry(err); !ok {
				break
			}
		}
	}

	a.viewCredentialsList()
}

func (a *App) deleteCredentials(id string) {
	if ok := a.ui.Confirm("Are you sure"); ok {
		for {
			err := a.keeper.DeleteCredentials(id)
			if ok := a.ui.Retry(err); !ok {
				break
			}
		}

		a.viewCredentialsList()
	} else {
		a.viewCredentials(id)
	}
}
