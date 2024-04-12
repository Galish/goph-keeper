package app

import (
	"context"

	"github.com/Galish/goph-keeper/internal/client/app/cli"
	"github.com/Galish/goph-keeper/internal/client/auth"
	"github.com/Galish/goph-keeper/internal/client/usecase"
)

type UserInterface interface {
	Break()
	Close() error
	Confirm(string) bool
	Edit(string, string, bool) string
	Error(error)
	Exit()
	Input(string, bool) string
	InputPassword(string, bool) string
	Print(string)
	ReadFile(string, bool) []byte
	Retry(error) bool
	Select(string, []*cli.SelectOption)
	WriteFile(string, []byte, bool)
}

type App struct {
	auth   *auth.Manager
	user   usecase.User
	notes  usecase.SecureNotes
	health usecase.HealthCheck
	ui     UserInterface
}

func New(
	auth *auth.Manager,
	user usecase.User,
	notes usecase.SecureNotes,
	health usecase.HealthCheck,
) *App {
	return &App{
		auth:   auth,
		user:   user,
		notes:  notes,
		health: health,
		ui:     cli.New(),
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
			a.ui.Exit()
		}

		a.ui.Break()
	}

	if a.auth.IsAuthorized() {
		a.selectCategory(ctx)
	} else {
		a.viewAuthScreen(ctx)
	}
}

func (a *App) Close() error {
	return a.ui.Close()
}
