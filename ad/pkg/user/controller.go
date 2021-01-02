package user

import (
	"context"
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/hardstylez72/bblog/ad/pkg/util"
	"net/http"
)

type Repository interface {
	GetById(ctx context.Context, id int) (*User, error)
	List(ctx context.Context) ([]User, error)
	Insert(ctx context.Context, group *User) (*User, error)
	Delete(ctx context.Context, id int) error
}

type controller struct {
	rep Repository
}

func NewController(rep Repository) *controller {
	return &controller{rep: rep}
}

func (c *controller) create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var req insertRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	group, err := c.rep.Insert(ctx, insertRequestConvert(&req))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	render.JSON(w, r, newInsertResponse(group))
}

func (c *controller) getById(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var req getRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		util.NewResponse(w).WithError(err).WithStatus(http.StatusBadRequest).Send()
		return
	}

	user, err := c.rep.GetById(ctx, req.Id)
	if err != nil {
		util.NewResponse(w).WithError(err).WithStatus(http.StatusInternalServerError).Send()
		return
	}

	util.NewResponse(w).WithJson(user).WithStatus(http.StatusOK).Send()
}

func (c *controller) list(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	list, err := c.rep.List(ctx)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	render.JSON(w, r, newListResponse(list))
}

func (c *controller) delete(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req deleteRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err := c.rep.Delete(ctx, req.Id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (c *controller) Mount(r chi.Router) {
	r.Post("/v1/user/list", c.list)
	r.Post("/v1/user/get", c.getById)
	r.Post("/v1/user/create", c.create)
	r.Post("/v1/user/delete", c.delete)

}
