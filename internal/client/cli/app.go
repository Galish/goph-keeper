package cli

import (
	"github.com/Galish/goph-keeper/internal/client/auth"
	"github.com/Galish/goph-keeper/internal/client/cli/ui"
	"github.com/Galish/goph-keeper/internal/client/usecase"
)

type App struct {
	authClient *auth.AuthClient
	user       usecase.User
	keeper     usecase.Keeper
	ui         *ui.UI
	done       chan struct{}
}

func NewApp(authClient *auth.AuthClient, user usecase.User, keeper usecase.Keeper) *App {
	return &App{
		authClient: authClient,
		user:       user,
		keeper:     keeper,
		ui:         ui.New(),
		done:       make(chan struct{}),
	}
}

func (a *App) Run() {
loop:
	for {
		select {
		case <-a.done:
			break loop

		default:
			if a.authClient.IsAuthorized() {
				a.renderWelcomeView()
			} else {
				a.renderHomeView()
			}
		}
	}
}

func (a *App) Close() error {
	close(a.done)

	return a.ui.Close()
}
