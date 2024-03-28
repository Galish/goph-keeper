package cli

import (
	"fmt"

	"github.com/Galish/goph-keeper/internal/client/cli/ui"
	"github.com/Galish/goph-keeper/internal/client/entity"
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

	handleDelete := func() {
		if ok := a.ui.Confirm("Are you sure"); ok {
			if err := a.keeper.DeleteTextNote(id); err != nil {
				a.ui.Error(err)
			}

			a.viewTextNotesList()
			return
		} else {
			a.viewTextNote(id)
		}
	}

	handleEdit := func() {
		a.editTextNote(id)
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
		if err := a.keeper.AddTextNote(&note); err != nil {
			a.ui.Error(err)
		}
	}

	a.viewTextNotesList()
}

func (a *App) editTextNote(id string) {
	note := &entity.TextNote{
		ID: id,
	}

	note.Title = a.ui.Input("Title", true)
	note.Description = a.ui.Input("Description", false)
	note.Value = a.ui.Input("Note", true)

	a.keeper.UpdateTextNote(note)

	a.viewTextNotesList()
}
