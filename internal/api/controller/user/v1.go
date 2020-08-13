package user

import (
	"github.com/go-chi/chi"
	"github.com/hardstylez72/bblog/internal/storage/user"
)

type userController struct {
	userStorage user.Storage
}

func NewUserController(userStorage user.Storage) *userController {
	return &userController{
		userStorage: userStorage,
	}
}

func (c userController) Mount(r chi.Router) {
	r.Get("/v1/user/{user_id}", c.GetUserById)
}
