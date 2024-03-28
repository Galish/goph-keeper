package cli

import (
	"fmt"

	"github.com/Galish/goph-keeper/internal/client/cli/ui"
	"github.com/Galish/goph-keeper/internal/client/entity"
)

func (a *App) viewCardsList() {
	cards, err := a.keeper.GetCardsList()
	if err != nil {
		a.ui.Error(err)
		return
	}

	commands := []*ui.SelectOption{
		{
			Label: "+ Add new",
			Run:   a.addCard,
		},
		{
			Label: "  Cancel",
		},
	}

	for i, c := range cards {
		id := c.ID

		commands = append(
			commands,
			&ui.SelectOption{
				Label: fmt.Sprintf("%d. %s \t %s", i+1, c.Title, c.Description),
				Run: func() {
					a.viewCard(id)
				},
			},
		)
	}

	a.ui.Select("Add new card details or select existing", commands)
}

func (a *App) viewCard(id string) {
	card, err := a.keeper.GetCard(id)
	if err != nil {
		a.ui.Error(err)
		return
	}

	a.ui.Print(card.String())

	handleDelete := func() {
		if ok := a.ui.Confirm("Are you sure"); ok {
			if err := a.keeper.DeleteCard(id); err != nil {
				a.ui.Error(err)
			}

			a.viewCardsList()
			return
		} else {
			a.viewCard(id)
		}
	}

	handleEdit := func() {
		a.editCard(id)
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
			Run:   a.viewCardsList,
		},
	}

	a.ui.Select("Select action", commands)
}

func (a *App) addCard() {
	card := entity.Card{}

	card.Title = a.ui.Input("Title", true)
	card.Description = a.ui.Input("Description", false)
	card.Number = a.ui.Input("Card number", true)
	card.Holder = a.ui.Input("Card holder", true)
	card.CVC = a.ui.Input("CVC code", true)
	card.SetExpiry(a.ui.Input("Expiration date", true))

	if ok := a.ui.Confirm("Add card details"); ok {
		if err := a.keeper.AddCard(&card); err != nil {
			a.ui.Error(err)
		}
	}

	a.viewCardsList()
}

func (a *App) editCard(id string) {
	card := &entity.Card{
		ID: id,
	}

	card.Title = a.ui.Input("Title", true)
	card.Description = a.ui.Input("Description", false)
	card.Number = a.ui.Input("Card number", true)
	card.Holder = a.ui.Input("Card holder", true)
	card.CVC = a.ui.Input("CVC code", true)
	card.SetExpiry(a.ui.Input("Expiration date", true))

	a.keeper.UpdateCard(card)

	a.viewCardsList()
}
