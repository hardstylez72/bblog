package auth

import (
	"context"
	"github.com/hardstylez72/bblog/internal/storage/user"
	"net/http"
)

const (
	UserRoleCode = "USER"
)

type OAuth2 interface {
	GetToken(ctx context.Context, code string) (interface{}, error)
	GetUser(ctx context.Context, state string, code string) (*user.User, error)
	HandleLogin(w http.ResponseWriter, r *http.Request)
	HandleCallback(w http.ResponseWriter, r *http.Request)
}
