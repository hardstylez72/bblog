package route

import (
	"context"
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"net/http"
)

type insertParams struct {
	GroupId int
	RouteId int
}

type Repository interface {
	List(ctx context.Context, groupId int) ([]Route, error)
	Insert(ctx context.Context, params []insertParams) ([]Route, error)
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

	list, err := c.rep.List(ctx, req.GroupId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	render.JSON(w, r, newListResponse(list))
}

func (c *controller) Mount(r chi.Router) {
	r.Post("/v1/group/route/List", c.list)
	r.Post("/v1/group/route/create", c.create)
}
