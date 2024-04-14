// Package implements the user authentication and authorization logic.
package user

import (
	"time"

	pb "github.com/Galish/goph-keeper/api/proto"
)

var defaultTimeout = 1 * time.Minute

// UseCase represents the user usecase.
type UseCase struct {
	client pb.KeeperClient
}

// New returns a new user usecase.
func New(client pb.KeeperClient) *UseCase {
	return &UseCase{
		client: client,
	}
}
