package cli

import (
	"context"
	"errors"
	"fmt"

	"github.com/Galish/goph-keeper/internal/client/cli/ui"
	"github.com/Galish/goph-keeper/internal/client/usecase/keeper"
	"github.com/Galish/goph-keeper/internal/entity"
)

func (a *App) viewCredentialsList(ctx context.Context) {
	creds, err := a.keeper.GetCredentialsList(ctx)
	if err != nil {
		a.ui.Error(err)
		return
	}

	commands := []*ui.SelectOption{
		{
			Label: "+ Add new",
			Run: func() {
				a.addCredentials(ctx)
			},
		},
		{
			Label: "  Cancel",
			Run: func() {
				a.selectCategory(ctx)
			},
		},
	}

	for i, c := range creds {
		id := c.ID

		commands = append(
			commands,
			&ui.SelectOption{
				Label: fmt.Sprintf("%d. %s \t %s", i+1, c.Title, c.Description),
				Run: func() {
					a.viewCredentials(ctx, id)
				},
			},
		)
	}

	a.ui.Select("Add new credentials or select existing", commands)
}

func (a *App) viewCredentials(ctx context.Context, id string) {
	creds, err := a.keeper.GetCredentials(ctx, id)
	if err != nil {
		a.ui.Error(err)
		return
	}

	a.ui.Print(creds.String())

	var commands = []*ui.SelectOption{
		{
			Label: "Edit",
			Run: func() {
				a.editCredentials(ctx, id)
			},
		},
		{
			Label: "Delete",
			Run: func() {
				a.deleteCredentials(ctx, id)
			},
		},
		{
			Label: "Cancel",
			Run: func() {
				a.viewCredentialsList(ctx)
			},
		},
	}

	a.ui.Select("Select action", commands)
}

func (a *App) addCredentials(ctx context.Context) {
	creds := entity.Credentials{}

	creds.Title = a.ui.Input("Title", true)
	creds.Description = a.ui.Input("Description", false)
	creds.Username = a.ui.Input("Username", true)
	creds.Password = a.ui.Input("Password", true)

	if ok := a.ui.Confirm("Add credentials"); ok {
		for {
			err := a.keeper.AddCredentials(ctx, &creds)
			if ok := a.ui.Retry(err); !ok {
				break
			}
		}
	}

	a.viewCredentialsList(ctx)
}

func (a *App) editCredentials(ctx context.Context, id string) {
	creds, err := a.keeper.GetCredentials(ctx, id)
	if err != nil {
		a.ui.Error(err)
		return
	}

	var (
		overwrite bool
		updated   = &entity.Credentials{
			ID:      id,
			Version: creds.Version + 1,
		}
	)

	updated.Title = a.ui.Edit("Title", creds.Title, true)
	updated.Description = a.ui.Edit("Description", creds.Description, false)
	updated.Username = a.ui.Edit("Username", creds.Username, true)
	updated.Password = a.ui.Edit("Password", creds.Password, true)

	if ok := a.ui.Confirm("Update credentials"); ok {
		for {
			err := a.keeper.UpdateCredentials(ctx, updated, overwrite)
			if errors.Is(err, keeper.ErrVersionConflict) {
				if ok := a.ui.Confirm("Credentials have already been updated. Want to overwrite"); ok {
					overwrite = true
					continue
				}

				break
			}

			if ok := a.ui.Retry(err); !ok {
				break
			}
		}
	}

	a.viewCredentialsList(ctx)
}

func (a *App) deleteCredentials(ctx context.Context, id string) {
	if ok := a.ui.Confirm("Are you sure"); ok {
		for {
			err := a.keeper.DeleteCredentials(ctx, id)
			if ok := a.ui.Retry(err); !ok {
				break
			}
		}

		a.viewCredentialsList(ctx)
	} else {
		a.viewCredentials(ctx, id)
	}
}
