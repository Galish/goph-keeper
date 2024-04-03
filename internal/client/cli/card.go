package cli

import (
	"errors"
	"fmt"

	"github.com/Galish/goph-keeper/internal/client/cli/ui"
	"github.com/Galish/goph-keeper/internal/client/usecase/keeper"
	"github.com/Galish/goph-keeper/internal/entity"
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
			Run:   a.selectCategory,
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

	var commands = []*ui.SelectOption{
		{
			Label: "Edit",
			Run: func() {
				a.editCard(id)
			},
		},
		{
			Label: "Delete",
			Run: func() {
				a.deleteCard(id)
			},
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
		for {
			err := a.keeper.AddCard(&card)
			if ok := a.ui.Retry(err); !ok {
				break
			}
		}
	}

	a.viewCardsList()
}

func (a *App) editCard(id string) {
	card, err := a.keeper.GetCard(id)
	if err != nil {
		a.ui.Error(err)
		return
	}

	var (
		overwrite bool
		updated   = &entity.Card{
			ID:      id,
			Version: card.Version + 1,
		}
	)

	updated.Title = a.ui.Edit("Title", card.Title, true)
	updated.Description = a.ui.Edit("Description", card.Description, false)
	updated.Number = a.ui.Edit("Card number", card.Number, true)
	updated.Holder = a.ui.Edit("Card holder", card.Holder, true)
	updated.CVC = a.ui.Edit("CVC code", card.CVC, true)
	updated.SetExpiry(a.ui.Edit("Expiration date", card.GetExpiry(), true))

	if ok := a.ui.Confirm("Update card details"); ok {
		for {
			err := a.keeper.UpdateCard(updated, overwrite)
			if errors.Is(err, keeper.ErrVersionConflict) {
				if ok := a.ui.Confirm("Card details have already been updated. Do you want to overwrite"); ok {
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

	a.viewCardsList()
}

func (a *App) deleteCard(id string) {
	if ok := a.ui.Confirm("Are you sure"); ok {
		for {
			err := a.keeper.DeleteCard(id)
			if ok := a.ui.Retry(err); !ok {
				break
			}
		}

		a.viewCardsList()
	} else {
		a.viewCard(id)
	}
}
