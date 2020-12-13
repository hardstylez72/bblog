package grouproute

import (
	"context"
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type params struct {
	GroupId int `json:"groupId" validate:"required"`
	RouteId int `json:"routeId" validate:"required"`
}

type Repository interface {
	List(ctx context.Context, groupId int) ([]Route, error)
	ListNotInGroup(ctx context.Context, groupId int) ([]Route, error)
	Insert(ctx context.Context, params []params) ([]Route, error)
	Delete(ctx context.Context, params []params) error
}

type controller struct {
	rep       Repository
	validator *validator.Validate
}

func NewController(rep Repository) *controller {
	return &controller{
		rep:       rep,
		validator: validator.New(),
	}
}

func (c *controller) create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var req insertRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	routes, err := c.rep.Insert(ctx, insertRequestConvert(req))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	render.JSON(w, r, routes)
}

func (c *controller) list(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req listRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := c.validator.Struct(req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	var list []Route
	var err error

	if req.BelongToGroup {
		list, err = c.rep.List(ctx, req.GroupId)
	} else {
		list, err = c.rep.ListNotInGroup(ctx, req.GroupId)
	}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	render.JSON(w, r, newListResponse(list))
}

func (c *controller) delete(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req []params

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	rr := deleteRequest{Params: req}

	if err := c.validator.Struct(rr); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	err := c.rep.Delete(ctx, req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (c *controller) Mount(r chi.Router) {
	r.Post("/v1/group/route/list", c.list)
	r.Post("/v1/group/route/create", c.create)
	r.Post("/v1/group/route/delete", c.delete)
}
