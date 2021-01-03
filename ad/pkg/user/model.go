package user

import (
	"github.com/hardstylez72/bblog/ad/pkg/util"
	"time"
)

type User struct {
	Id         int    `json:"id" db:"id"`
	ExternalId string `json:"externalId" db:"external_id"`
	IsSystem   bool   `json:"isSystem" db:"is_system"`

	Name        util.JsonNullString `json:"name" db:"name"`
	Description util.JsonNullString `json:"description" db:"description"`
	Email       util.JsonNullString `json:"email" db:"email"`
	Phone       util.JsonNullString `json:"hone" db:"phone"`

	CreatedAt time.Time         `json:"createdAt" db:"created_at"`
	UpdatedAt util.JsonNullTime `json:"updatedAt" db:"updated_at"`
	DeletedAt util.JsonNullTime `json:"deletedAt" db:"deleted_at"`
}
