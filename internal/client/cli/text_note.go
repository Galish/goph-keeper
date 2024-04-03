package cli

import (
	"errors"
	"fmt"

	"github.com/Galish/goph-keeper/internal/client/cli/ui"
	"github.com/Galish/goph-keeper/internal/client/usecase/keeper"
	"github.com/Galish/goph-keeper/internal/entity"
)

func (a *App) viewTextNotesList() {
	notes, err := a.keeper.GetTextNotesList()
	if err != nil {
		a.ui.Error(err)
		return
	}

	commands := []*ui.SelectOption{
		{
			Label: "+ Add new",
			Run:   a.addTextNote,
		},
		{
			Label: "  Cancel",
			Run:   a.selectCategory,
		},
	}

	for i, n := range notes {
		id := n.ID

		commands = append(
			commands,
			&ui.SelectOption{
				Label: fmt.Sprintf("%d. %s \t %s", i+1, n.Title, n.Description),
				Run: func() {
					a.viewTextNote(id)
				},
			},
		)
	}

	a.ui.Select("Add new text note or select existing", commands)
}

func (a *App) viewTextNote(id string) {
	note, err := a.keeper.GetTextNote(id)
	if err != nil {
		a.ui.Error(err)
		return
	}

	a.ui.Print(note.String())

	var commands = []*ui.SelectOption{
		{
			Label: "Edit",
			Run: func() {
				a.editTextNote(id)
			},
		},
		{
			Label: "Delete",
			Run: func() {
				a.deleteTextNote(id)
			},
		},
		{
			Label: "Cancel",
			Run:   a.viewTextNotesList,
		},
	}

	a.ui.Select("Select action", commands)
}

func (a *App) addTextNote() {
	note := entity.TextNote{}

	note.Title = a.ui.Input("Title", true)
	note.Description = a.ui.Input("Description", false)
	note.Value = a.ui.Input("Note", true)

	if ok := a.ui.Confirm("Add text note"); ok {
		for {
			err := a.keeper.AddTextNote(&note)
			if ok := a.ui.Retry(err); !ok {
				break
			}
		}
	}

	a.viewTextNotesList()
}

func (a *App) editTextNote(id string) {
	note, err := a.keeper.GetTextNote(id)
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

	if ok := a.ui.Confirm("Update text note"); ok {
		for {
			err := a.keeper.UpdateTextNote(updated, overwrite)
			if errors.Is(err, keeper.ErrVersionConflict) {
				if ok := a.ui.Confirm("Text note have already been updated. Do you want to overwrite"); ok {
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

	a.viewTextNotesList()
}

func (a *App) deleteTextNote(id string) {
	if ok := a.ui.Confirm("Are you sure"); ok {
		for {
			err := a.keeper.DeleteTextNote(id)
			if ok := a.ui.Retry(err); !ok {
				break
			}
		}

		a.viewTextNotesList()
	} else {
		a.viewTextNote(id)
	}
}
