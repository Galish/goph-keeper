package cli

import (
	"context"
	"errors"
	"fmt"

	"github.com/Galish/goph-keeper/internal/client/cli/ui"
	"github.com/Galish/goph-keeper/internal/client/usecase/notes"
	"github.com/Galish/goph-keeper/internal/entity"
)

func (a *App) viewTextNotesList(ctx context.Context) {
	a.ui.Break()

	notes, err := a.notes.GetTextNotesList(ctx)
	if err != nil {
		a.ui.Error(err)

		return
	}

	commands := []*ui.SelectOption{
		{
			Label: "+ Add new",
			Run: func() {
				a.addTextNote(ctx)
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
			&ui.SelectOption{
				Label: fmt.Sprintf("%d. %s \t %s", i+1, n.Title, n.Description),
				Run: func() {
					a.viewTextNote(ctx, id)
				},
			},
		)
	}

	a.ui.Select("Add new text note or select existing", commands)
}

func (a *App) viewTextNote(ctx context.Context, id string) {
	a.ui.Break()

	note, err := a.notes.GetTextNote(ctx, id)
	if err != nil {
		a.ui.Error(err)

		return
	}

	a.ui.Print(note.String())
	a.ui.Break()

	var commands = []*ui.SelectOption{
		{
			Label: "Edit",
			Run: func() {
				a.editTextNote(ctx, id)
			},
		},
		{
			Label: "Delete",
			Run: func() {
				a.deleteTextNote(ctx, id)
			},
		},
		{
			Label: "Cancel",
			Run: func() {
				a.viewTextNotesList(ctx)
			},
		},
	}

	a.ui.Select("Select action", commands)
}

func (a *App) addTextNote(ctx context.Context) {
	a.ui.Break()

	var note = new(entity.TextNote)

	note.Title = a.ui.Input("Title", true)
	note.Description = a.ui.Input("Description", false)
	note.Value = a.ui.Input("Note", true)

	a.ui.Break()

	if ok := a.ui.Confirm("Add text note"); ok {
		for {
			err := a.notes.AddTextNote(ctx, note)
			if ok := a.ui.Retry(err); !ok {
				break
			}
		}
	}

	a.viewTextNotesList(ctx)
}

func (a *App) editTextNote(ctx context.Context, id string) {
	a.ui.Break()

	note, err := a.notes.GetTextNote(ctx, id)
	if err != nil {
		a.ui.Error(err)

		return
	}

	var (
		overwrite bool
		updated   = &entity.TextNote{
			ID:      id,
			Version: note.Version + 1,
		}
	)

	updated.Title = a.ui.Edit("Title", note.Title, true)
	updated.Description = a.ui.Edit("Description", note.Description, false)
	updated.Value = a.ui.Edit("Note", note.Value, true)

	a.ui.Break()

	if ok := a.ui.Confirm("Update text note"); ok {
		for {
			err := a.notes.UpdateTextNote(ctx, updated, overwrite)
			if errors.Is(err, notes.ErrVersionConflict) {
				if ok := a.ui.Confirm("Text note has already been updated. Want to overwrite"); ok {
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

	a.viewTextNotesList(ctx)
}

func (a *App) deleteTextNote(ctx context.Context, id string) {
	a.ui.Break()

	if ok := a.ui.Confirm("Are you sure"); ok {
		for {
			err := a.notes.DeleteTextNote(ctx, id)
			if ok := a.ui.Retry(err); !ok {
				break
			}
		}

		a.viewTextNotesList(ctx)

		return
	}

	a.viewTextNote(ctx, id)
}
