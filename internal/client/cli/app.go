package cli

import (
	"context"
	"syscall"

	"github.com/Galish/goph-keeper/internal/client/auth"
	"github.com/Galish/goph-keeper/internal/client/cli/ui"
	"github.com/Galish/goph-keeper/internal/client/usecase"
)

type App struct {
	auth   *auth.AuthManager
	user   usecase.User
	notes  usecase.SecureNotes
	health usecase.HealthCheck
	ui     ui.UserInterface
}

func NewApp(
	auth *auth.AuthManager,
	user usecase.User,
	notes usecase.SecureNotes,
	health usecase.HealthCheck,
) *App {
	return &App{
		auth:   auth,
		user:   user,
		notes:  notes,
		health: health,
		ui:     ui.New(),
	}
}

func (a *App) Run(ctx context.Context) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	for {
		err := a.health.Check(ctx)
		if err == nil {
			break
		}

		if ok := a.ui.Retry(err); !ok {
			a.Stop()
		}

		a.ui.Break()
	}

	if a.auth.IsAuthorized() {
		a.selectCategory(ctx)
	} else {
		a.viewAuthScreen(ctx)
	}
}

func (a *App) Stop() {
	syscall.Kill(syscall.Getpid(), syscall.SIGINT)
}

func (a *App) Close() error {
	return a.ui.Close()
}
