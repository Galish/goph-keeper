package cli

import (
	"context"

	"github.com/Galish/goph-keeper/internal/client/auth"
	"github.com/Galish/goph-keeper/internal/client/cli/ui"
	"github.com/Galish/goph-keeper/internal/client/usecase"
)

type App struct {
	auth   *auth.AuthManager
	user   usecase.User
	keeper usecase.Keeper
	ui     ui.UserInterface
}

func NewApp(ui ui.UserInterface, auth *auth.AuthManager, user usecase.User, keeper usecase.Keeper) *App {
	return &App{
		auth:   auth,
		user:   user,
		keeper: keeper,
		ui:     ui,
	}
}

func (a *App) Run() {
	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	if a.auth.IsAuthorized() {
		a.selectCategory(ctx)
	} else {
		a.viewAuthScreen(ctx)
	}
}
