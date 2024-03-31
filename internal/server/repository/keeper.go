package repository

import (
	"context"
	"time"
)

type KeeperRepository interface {
	AddSecureRecord(context.Context, *SecureRecord) error
	UpdateSecureRecord(context.Context, *SecureRecord) error
	DeleteSecureRecord(context.Context, string, string, SecureRecordType) error
	GetSecureRecord(context.Context, string, string, SecureRecordType) (*SecureRecord, error)
	GetSecureRecords(context.Context, string, SecureRecordType) ([]*SecureRecord, error)
}

const (
	TypeCredentials SecureRecordType = iota + 1
	TypeTextNote
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
