package group

import (
	"context"
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"net/http"
)

type Repository interface {
	List(ctx context.Context) ([]Group, error)
	Insert(ctx context.Context, group *Group) (*Group, error)
}

type controller struct {
	rep Repository
}

func NewController(rep Repository) *controller {
	return &controller{rep: rep}
}

func (c *controller) addGroup(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var req insertGroupRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	group, err := c.rep.Insert(ctx, insertGroupRequestConvert(&req))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	render.JSON(w, r, newInsertGroupResponse(group))
}

func (c *controller) getGroups(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	list, err := c.rep.List(ctx)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	render.JSON(w, r, newGetGroupsResponse(list))
}

func (c *controller) Mount(r chi.Router) {
	r.Post("/v1/groups/get", c.getGroups)
	r.Post("/v1/group/create", c.addGroup)
}
