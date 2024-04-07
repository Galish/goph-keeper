package repository

import (
	"context"
	"time"
)

type KeeperRepository interface {
	AddSecureRecord(ctx context.Context, record *SecureRecord) error
	UpdateSecureRecord(ctx context.Context, record *SecureRecord) error
	DeleteSecureRecord(ctx context.Context, user, id string, recordType SecureRecordType) error
	GetSecureRecord(ctx context.Context, user, id string, recordType SecureRecordType) (*SecureRecord, error)
	GetSecureRecords(ctx context.Context, user string, recordType SecureRecordType) ([]*SecureRecord, error)
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
	Version      int32
}
