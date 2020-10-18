package user

import (
	"github.com/hardstylez72/bblog/internal/storage/user"
	"time"
)

type User struct {
	Id               int       `json:"id"`
	RegisteredAt     time.Time `json:"registeredAt"`
	ExternalId       string    `json:"externalId"`
	ExternalAuthType string    `json:"externalAuthType"`
	Login            *string   `json:"login"`
	Name             *string   `json:"name"`
	Email            *string   `json:"email"`
}

func NewGetUserByIdResponse(u *user.User) *User {
	out := &User{
		Id:               u.Id,
		RegisteredAt:     u.CreatedAt,
		ExternalId:       u.ExternalId,
		ExternalAuthType: u.AuthType,
		Login:            nil,
		Name:             nil,
		Email:            nil,
	}

	if u.Name.Valid {
		out.Name = &u.Name.String
	}
	if u.Login.Valid {
		out.Login = &u.Login.String
	}
	if u.Email.Valid {
		out.Email = &u.Email.String
	}

	return out
}
