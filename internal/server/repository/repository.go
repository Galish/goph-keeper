package repository

import (
	"context"
	"time"
)

type Repository interface {
	Get(context.Context, string) (*Record, error)
	GetByType(context.Context, RecordType) ([]*Record, error)
	Set(context.Context, *Record) error
}

const (
	TypeCreds RecordType = iota
	TypeNote
	TypeCardDetails
)

type RecordType int

type Record struct {
	ID             string
	Type           RecordType
	Title          string
	Description    string
	Username       string
	Password       string
	TextNote       string
	RawNote        []byte
	CardNumber     string
	CardHolderName string
	CardCVC        string
	CardExpiryDate time.Time
	CreatedAt      time.Time
	LastEditedAt   time.Time
}
