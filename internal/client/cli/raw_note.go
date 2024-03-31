package cli

import (
	"fmt"

	"github.com/Galish/goph-keeper/internal/client/cli/ui"
	"github.com/Galish/goph-keeper/internal/client/entity"
)

func (a *App) viewRawNotesList() {
	notes, err := a.keeper.GetRawNotesList()
	if err != nil {
		a.ui.Error(err)
		return
	}

	commands := []*ui.SelectOption{
		{
			Label: "+ Add new",
			Run:   a.addRawNote,
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
					a.viewRawNote(id)
				},
			},
		)
	}

	a.ui.Select("Add new binary note or select existing", commands)
}

func (a *App) viewRawNote(id string) {
	note, err := a.keeper.GetRawNote(id)
	if err != nil {
		a.ui.Error(err)
		return
	}

	a.ui.Print(note.String())

	handleDelete := func() {
		if ok := a.ui.Confirm("Are you sure"); ok {
			if err := a.keeper.DeleteRawNote(id); err != nil {
				a.ui.Error(err)
			}

			a.viewRawNotesList()
			return
		} else {
			a.viewRawNote(id)
		}
	}

	handleEdit := func() {
		a.editRawNote(id)
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
			Run:   a.viewRawNotesList,
		},
	}

	a.ui.Select("Select action", commands)
}

func (a *App) addRawNote() {
	note := entity.RawNote{}

	note.Title = a.ui.Input("Title", true)
	note.Description = a.ui.Input("Description", false)

	value := a.ui.Input("Note", true)
	if err := note.SetValue(value); err != nil {
		a.ui.Error(err)
	}

	if ok := a.ui.Confirm("Add text note"); ok {
		if err := a.keeper.AddRawNote(&note); err != nil {
			a.ui.Error(err)
		}
	}

	a.viewRawNotesList()
}

func (a *App) editRawNote(id string) {
	note := &entity.RawNote{
		ID: id,
	}

	note.Title = a.ui.Input("Title", true)
	note.Description = a.ui.Input("Description", false)

	value := a.ui.Input("Note", true)
	if err := note.SetValue(value); err != nil {
		a.ui.Error(err)
	}

	a.keeper.UpdateRawNote(note)

	a.viewRawNotesList()
}