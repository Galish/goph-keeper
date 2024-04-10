// Package implements the logic for working with secure notes.
package notes

import (
	"time"

	pb "github.com/Galish/goph-keeper/api/proto"
)

var defaultTimeout = 1 * time.Minute

// UseCase represents the secure notes usecase.
type UseCase struct {
	client pb.KeeperClient
}

// New returns a new secure notes usecase.
func New(client pb.KeeperClient) *UseCase {
	return &UseCase{
		client: client,
	}
}
