package cli

import (
	"context"
	"errors"
	"fmt"

	"github.com/Galish/goph-keeper/internal/client/cli/ui"
	"github.com/Galish/goph-keeper/internal/client/usecase/notes"
	"github.com/Galish/goph-keeper/internal/entity"
	"github.com/Galish/goph-keeper/pkg/logger"
)

func (a *App) viewCardsList(ctx context.Context) {
	a.ui.Break()

	cards, err := a.notes.GetCardsList(ctx)
	if err != nil {
		a.ui.Error(err)

		return
	}

	commands := []*ui.SelectOption{
		{
			Label: "+ Add new",
			Run: func() {
				a.addCard(ctx)
			},
		},
		{
			Label: "x Cancel",
			Run: func() {
				a.selectCategory(ctx)
			},
		},
	}

	for i, c := range cards {
		id := c.ID

		commands = append(
			commands,
			&ui.SelectOption{
				Label: fmt.Sprintf("%d. %s \t %s", i+1, c.Title, c.Description),
				Run: func() {
					a.viewCard(ctx, id)
				},
			},
		)
	}

	a.ui.Select("Add new card details or select existing", commands)
}

func (a *App) viewCard(ctx context.Context, id string) {
	a.ui.Break()

	card, err := a.notes.GetCard(ctx, id)
	if err != nil {
		a.ui.Error(err)

		return
	}

	a.ui.Print(card.String())
	a.ui.Break()

	var commands = []*ui.SelectOption{
		{
			Label: "Edit",
			Run: func() {
				a.editCard(ctx, id)
			},
		},
		{
			Label: "Delete",
			Run: func() {
				a.deleteCard(ctx, id)
			},
		},
		{
			Label: "Cancel",
			Run: func() {
				a.viewCardsList(ctx)
			},
		},
	}

	a.ui.Select("Select action", commands)
}

func (a *App) addCard(ctx context.Context) {
	a.ui.Break()

	var card = new(entity.Card)

	card.Title = a.ui.Input("Title", true)
	card.Description = a.ui.Input("Description", false)
	card.Number = a.ui.Input("Card number", true)
	card.Holder = a.ui.Input("Card holder", true)
	card.CVC = a.ui.Input("CVC code", true)

	expiry := a.ui.Input("Expiration date", true)
	if err := card.SetExpiry(expiry); err != nil {
		logger.WithError(err).Debug("failed to set card expiration date")
	}

	a.ui.Break()

	if ok := a.ui.Confirm("Add card details"); ok {
		for {
			err := a.notes.AddCard(ctx, card)
			if ok := a.ui.Retry(err); !ok {
				break
			}
		}
	}

	a.viewCardsList(ctx)
}

func (a *App) editCard(ctx context.Context, id string) {
	a.ui.Break()

	card, err := a.notes.GetCard(ctx, id)
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

	expiry := a.ui.Edit("Expiration date", card.GetExpiry(), true)
	if err := updated.SetExpiry(expiry); err != nil {
		logger.WithError(err).Debug("failed to set card expiration date")
	}

	a.ui.Break()

	if ok := a.ui.Confirm("Update card details"); ok {
		for {
			err := a.notes.UpdateCard(ctx, updated, overwrite)
			if errors.Is(err, notes.ErrVersionConflict) {
				if ok := a.ui.Confirm("Card details have already been updated. Want to overwrite"); ok {
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

	a.viewCardsList(ctx)
}

func (a *App) deleteCard(ctx context.Context, id string) {
	a.ui.Break()

	if ok := a.ui.Confirm("Are you sure"); ok {
		for {
			err := a.notes.DeleteCard(ctx, id)
			if ok := a.ui.Retry(err); !ok {
				break
			}
		}

		a.viewCardsList(ctx)

		return
	}

	a.viewCard(ctx, id)
}
