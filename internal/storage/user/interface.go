package user

import (
	"context"
	"database/sql"
	"time"
)

type User struct {
	Id         int    `db:"id"`
	AuthType   string `db:"auth_type"`
	ExternalId string `db:"external_id"`

	Email sql.NullString `db:"email"`
	Login sql.NullString `db:"login"`
	Name  sql.NullString `db:"name"`

	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at"`
	DeletedAt sql.NullTime `db:"deleted_at"`
}

type Storage interface {
	SaveUser(ctx context.Context, user *User) error
	GetUserById(ctx context.Context, id int) (*User, error)
	GetUserByExternalId(ctx context.Context, userExternalId, authTypeId string) (*User, error)
	GetUsersAmount(ctx context.Context) (int, error)
}
