package repository

import (
	"context"
	"time"
)

type KeeperRepository interface {
	SetSecureRecord(context.Context, *SecureRecord) error
	GetSecureRecord(context.Context, string, string) (*SecureRecord, error)
	GetSecureRecords(context.Context, string, SecureRecordType) ([]*SecureRecord, error)
}

const (
	TypeCredentials SecureRecordType = iota + 1
	TypeNote
	TypeRawNote
	TypeCard
)

type SecureRecordType int

type SecureRecord struct {
	ID           string
	Type         SecureRecordType
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
