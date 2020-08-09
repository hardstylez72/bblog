package user

import (
	"context"
	"database/sql"
)

type UserAgr struct {
	User
	Emails []Email
}

type User struct {
	Id               string         `db:"id"`
	RegisteredAt     sql.NullTime   `db:"registered_at"`
	ExternalId       sql.NullString `db:"external_id"`
	ExternalAuthType sql.NullString `db:"external_auth_type"`
	Login            sql.NullString `db:"login"`
	FirstName        sql.NullString `db:"first_name"`
	LastName         sql.NullString `db:"last_name"`
	MiddleName       sql.NullString `db:"middle_name"`
	IsBanned         sql.NullBool   `db:"is_banned"`
}

type Email struct {
	Id            string         `db:"id"`
	UserId        sql.NullString `db:"user_id"`
	CreatedAt     sql.NullTime   `db:"created_at"`
	IsActive      sql.NullBool   `db:"is_active"`
	Address       sql.NullString `db:"address"`
}

type Storage interface {
	SaveUserWithEmail(ctx context.Context, u *UserAgr) error
	GetUserById(ctx context.Context, id string) (*User, error)
	GetUserByExternalId(ctx context.Context, id, authTypeId string) (*User, error)
	//GetEmailsByUserId(ctx context.Context, id string) ([]Email, error)
	//AddEmail(ctx context.Context, u *Email) error
}
