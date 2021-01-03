package group

import (
	"context"
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/hardstylez72/bblog/ad/pkg/util"
	"net/http"
)

type Repository interface {
	GetById(ctx context.Context, id int) (*Group, error)
	List(ctx context.Context) ([]Group, error)
	Insert(ctx context.Context, group *Group) (*Group, error)
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
		util.NewResponse(w).WithError(err).WithStatus(http.StatusBadRequest).Send()
		return
	}

	group, err := c.rep.Insert(ctx, insertRequestConvert(&req))
	if err != nil {
		util.NewResponse(w).WithError(err).WithStatus(http.StatusInternalServerError).Send()
		return
	}
	util.NewResponse(w).WithStatus(http.StatusOK).WithJson(newInsertResponse(group)).Send()
}

func (c *controller) list(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	list, err := c.rep.List(ctx)
	if err != nil {
		util.NewResponse(w).WithError(err).WithStatus(http.StatusInternalServerError).Send()
		return
	}
	util.NewResponse(w).WithStatus(http.StatusOK).WithJson(newListResponse(list)).Send()
}
func (c *controller) getById(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req getRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		util.NewResponse(w).WithError(err).WithStatus(http.StatusBadRequest).Send()
		return
	}

	group, err := c.rep.GetById(ctx, req.Id)
	if err != nil {
		util.NewResponse(w).WithError(err).WithStatus(http.StatusInternalServerError).Send()
		return
	}

	util.NewResponse(w).WithStatus(http.StatusOK).WithJson(group).Send()
}

func (c *controller) delete(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req deleteRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		util.NewResponse(w).WithError(err).WithStatus(http.StatusBadRequest).Send()
		return
	}

	err := c.rep.Delete(ctx, req.Id)
	if err != nil {
		util.NewResponse(w).WithError(err).WithStatus(http.StatusInternalServerError).Send()
		return
	}

	util.NewResponse(w).WithStatus(http.StatusOK).Send()
}

func (c *controller) Mount(r chi.Router) {
	r.Post("/v1/group/list", c.list)
	r.Post("/v1/group/get", c.getById)
	r.Post("/v1/group/create", c.create)
	r.Post("/v1/group/delete", c.delete)

}
