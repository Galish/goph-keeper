package app

import (
	"context"
	"errors"
	"fmt"

	"github.com/Galish/goph-keeper/internal/client/app/cli"
	"github.com/Galish/goph-keeper/internal/client/usecase/notes"
	"github.com/Galish/goph-keeper/internal/entity"
)

func (a *App) viewRawNotesList(ctx context.Context) {
	a.ui.Break()

	notes, err := a.notes.GetRawNotesList(ctx)
	if err != nil {
		a.ui.Error(err)

		return
	}

	commands := []*cli.SelectOption{
		{
			Label: "+ Add new",
			Run: func() {
				a.addRawNote(ctx)
			},
		},
		{
			Label: "x Cancel",
			Run: func() {
				a.selectCategory(ctx)
			},
		},
	}

	for i, n := range notes {
		id := n.ID

		commands = append(
			commands,
			&cli.SelectOption{
				Label: fmt.Sprintf("%d. %s \t %s", i+1, n.Title, n.Description),
				Run: func() {
					a.viewRawNote(ctx, id)
				},
			},
		)
	}

	a.ui.Select("Add new binary note or select existing", commands)
}

func (a *App) viewRawNote(ctx context.Context, id string) {
	a.ui.Break()

	note, err := a.notes.GetRawNote(ctx, id)
	if err != nil {
		a.ui.Error(err)

		return
	}

	a.ui.Print(note.String())
	a.ui.WriteFile("Enter the path to save the file", note.Value, false)
	a.ui.Break()

	var commands = []*cli.SelectOption{
		{
			Label: "Edit",
			Run: func() {
				a.editRawNote(ctx, id)
			},
		},
		{
			Label: "Delete",
			Run: func() {
				a.deleteRawNote(ctx, id)
			},
		},
		{
			Label: "Cancel",
			Run: func() {
				a.viewRawNotesList(ctx)
			},
		},
	}

	a.ui.Select("Select action", commands)
}

func (a *App) addRawNote(ctx context.Context) {
	a.ui.Break()

	var note = new(entity.RawNote)

	note.Title = a.ui.Input("Title", true)
	note.Description = a.ui.Input("Description", false)
	note.Value = a.ui.ReadFile("Enter file path", true)

	a.ui.Break()

	if ok := a.ui.Confirm("Add binary note"); ok {
		for {
			err := a.notes.AddRawNote(ctx, note)
			if ok := a.ui.Retry(err); !ok {
				break
			}
		}
	}

	a.viewRawNotesList(ctx)
}

func (a *App) editRawNote(ctx context.Context, id string) {
	a.ui.Break()

	note, err := a.notes.GetRawNote(ctx, id)
	if err != nil {
		a.ui.Error(err)

		return
	}

	var (
		overwrite bool
		updated   = &entity.RawNote{
			ID:      id,
			Version: note.Version + 1,
		}
	)

	updated.Title = a.ui.Edit("Title", note.Title, true)
	updated.Description = a.ui.Edit("Description", note.Description, false)

	if value := a.ui.ReadFile("Enter file path", false); value != nil {
		updated.Value = value
	} else {
		updated.Value = note.Value
	}

	a.ui.Break()

	if ok := a.ui.Confirm("Update binary note"); ok {
		for {
			err := a.notes.UpdateRawNote(ctx, updated, overwrite)
			if errors.Is(err, notes.ErrVersionConflict) {
				if ok := a.ui.Confirm("Binary note has already been updated. Want to overwrite"); ok {
					overwrite = true

					continue
				}

				break
			}

			a.ui.Break()

			if ok := a.ui.Retry(err); !ok {
				break
			}
		}
	}

	a.viewRawNotesList(ctx)
}

func (a *App) deleteRawNote(ctx context.Context, id string) {
	a.ui.Break()

	if ok := a.ui.Confirm("Are you sure"); ok {
		for {
			err := a.notes.DeleteRawNote(ctx, id)
			if ok := a.ui.Retry(err); !ok {
				break
			}
		}

		a.viewRawNotesList(ctx)

		return
	}

	a.viewRawNote(ctx, id)
}
