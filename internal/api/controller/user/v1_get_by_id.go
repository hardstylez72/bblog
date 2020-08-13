package user

import (
	"errors"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/hardstylez72/bblog/internal/api/controller"
	view "github.com/hardstylez72/bblog/internal/api/model/user"
	"net/http"
)

func (c userController) GetUserById(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	userId := chi.URLParam(r, "user_id")
	if userId == "" {
		err := errors.New("userId is missed")
		controller.ResponseWithError(controller.ErrInvalidInputParams(err), http.StatusBadRequest, w)
		return
	}

	user, err := c.userStorage.GetUserById(ctx, userId)
	if err != nil {
		controller.ResponseWithError(controller.ErrInternal(err), http.StatusInternalServerError, w)
		return
	}

	render.JSON(w, r, view.NewGetUserByIdResponse(user))

}
