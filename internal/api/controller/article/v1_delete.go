package article

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/hardstylez72/bbckend/internal/api/controller"
	ma "github.com/hardstylez72/bbckend/internal/api/model/article"
	"net/http"
)

func (c articleController) DeleteArticleByIdHandler(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	id := chi.URLParam(r, "id")

	err := c.articleStorage.DeleteArticleById(ctx, id)
	if err != nil {
		controller.ResponseWithError(controller.ErrInternal(err), http.StatusInternalServerError, w)
		return
	}

	render.JSON(w, r, ma.NewDeleteArticleResponse(id))
}
