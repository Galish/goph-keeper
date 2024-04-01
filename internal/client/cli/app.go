package cli

import (
	"github.com/Galish/goph-keeper/internal/client/auth"
	"github.com/Galish/goph-keeper/internal/client/cli/ui"
	"github.com/Galish/goph-keeper/internal/client/usecase"
)

type App struct {
	auth   *auth.AuthManager
	user   usecase.User
	keeper usecase.Keeper
	ui     ui.UserInterface
	done   chan struct{}
}

func NewApp(ui ui.UserInterface, auth *auth.AuthManager, user usecase.User, keeper usecase.Keeper) *App {
	return &App{
		auth:   auth,
		user:   user,
		keeper: keeper,
		ui:     ui,
		done:   make(chan struct{}),
	}
}

func (a *App) Run() {
loop:
	for {
		select {
		case <-a.done:
			break loop

		default:
			if a.auth.IsAuthorized() {
				a.selectCategory()
			} else {
				a.viewHomeScreen()
			}
		}
	}
}

func (a *App) Close() error {
	close(a.done)

	return nil
}
