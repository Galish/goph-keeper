package cli

import (
	"errors"
	"fmt"

	"github.com/Galish/goph-keeper/internal/client/cli/ui"
	"github.com/Galish/goph-keeper/internal/client/usecase/keeper"
	"github.com/Galish/goph-keeper/internal/entity"
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
	a.ui.WriteFile("Enter the path to save the file", note.Value, false)

	var commands = []*ui.SelectOption{
		{
			Label: "Edit",
			Run: func() {
				a.editRawNote(id)
			},
		},
		{
			Label: "Delete",
			Run: func() {
				a.deleteRawNote(id)
			},
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
	note.Value = a.ui.ReadFile("Enter file path", true)

	if ok := a.ui.Confirm("Add binary note"); ok {
		for {
			err := a.keeper.AddRawNote(&note)
			if ok := a.ui.Retry(err); !ok {
				break
			}
		}
	}

	a.viewRawNotesList()
}

func (a *App) editRawNote(id string) {
	note, err := a.keeper.GetRawNote(id)
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

	if ok := a.ui.Confirm("Update binary note"); ok {
		for {
			err := a.keeper.UpdateRawNote(updated, overwrite)
			if errors.Is(err, keeper.ErrVersionConflict) {
				if ok := a.ui.Confirm("Binary note has already been updated. Want to overwrite"); ok {
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

	a.viewRawNotesList()
}

func (a *App) deleteRawNote(id string) {
	if ok := a.ui.Confirm("Are you sure"); ok {
		for {
			err := a.keeper.DeleteRawNote(id)
			if ok := a.ui.Retry(err); !ok {
				break
			}
		}

		a.viewRawNotesList()
	} else {
		a.viewRawNote(id)
	}
}
