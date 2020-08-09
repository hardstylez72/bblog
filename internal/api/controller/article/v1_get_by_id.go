package article

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/hardstylez72/bblog/internal/api/controller"
	ma "github.com/hardstylez72/bblog/internal/api/model/article"
	"github.com/hardstylez72/bblog/internal/storage/user"
	"net/http"
)

func (c articleController) GetArticleByIdHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := chi.URLParam(r, "id")

	article, err := c.articleStorage.GetArticleById(ctx, id)
	if err != nil {
		if err == user.ErrNotFound {
			controller.ResponseWithError(ErrArticleNotFound(err), http.StatusNotFound, w)
		} else {
			controller.ResponseWithError(controller.ErrInternal(err), http.StatusInternalServerError, w)
		}
		return
	}

	render.JSON(w, r, ma.NewGetArticleByIdResponse(article))
}
