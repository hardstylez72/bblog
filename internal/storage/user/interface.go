package user

import (
	"context"
	"database/sql"
	"time"
)

type User struct {
	Id               string         `db:"id"`
	RegisteredAt     time.Time      `db:"registered_at"`
	ExternalId       string         `db:"external_id"`
	ExternalAuthType string         `db:"external_auth_type"`
	Login            sql.NullString `db:"login"`
	Name             sql.NullString `db:"name"`
	Email            sql.NullString `db:"email"`
	IsBanned         bool           `db:"is_banned"`
	RoleCode         string         `db:"role_code"`
}

type Storage interface {
	SaveUser(ctx context.Context, u *User) error
	GetUserById(ctx context.Context, id string) (*User, error)
	GetUserByExternalId(ctx context.Context, id, authTypeId string) (*User, error)
}
