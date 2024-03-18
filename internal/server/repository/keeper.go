package repository

import (
	"context"
	"time"
)

type KeeperRepository interface {
	Get(context.Context, string, string) (*Record, error)
	GetByType(context.Context, RecordType) ([]*Record, error)
	Set(context.Context, *Record) error
}

const (
	TypeCreds RecordType = iota + 1
	TypeNote
	TypeRawNote
	TypeCard
)

type RecordType int

type Record struct {
	ID           string
	Type         RecordType
	Title        string
	Description  string
	Username     string
	Password     string
	TextNote     string
	RawNote      []byte
	CardNumber   string
	CardHolder   string
	CardCVC      string
	CardExpiry   time.Time
	CreatedBy    string
	CreatedAt    time.Time
	LastEditedAt time.Time
}
