package cli

import (
	"context"

	"github.com/Galish/goph-keeper/internal/client/auth"
	"github.com/Galish/goph-keeper/internal/client/cli/ui"
	"github.com/Galish/goph-keeper/internal/client/usecase"
)

type App struct {
	auth  *auth.AuthManager
	user  usecase.User
	notes usecase.SecureNotes
	ui    ui.UserInterface
}

func NewApp(auth *auth.AuthManager, user usecase.User, notes usecase.SecureNotes) *App {
	return &App{
		auth:  auth,
		user:  user,
		notes: notes,
		ui:    ui.New(),
	}
}

func (a *App) Run(ctx context.Context) {
	ctx, cancel := context.WithCancel(ctx)

	defer cancel()

	if a.auth.IsAuthorized() {
		a.selectCategory(ctx)
	} else {
		a.viewAuthScreen(ctx)
	}
}

func (a *App) Close() error {
	return a.ui.Close()
}
