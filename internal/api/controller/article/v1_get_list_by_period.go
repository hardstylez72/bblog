package article

import (
	"github.com/go-chi/render"
	"github.com/hardstylez72/bblog/internal/api/controller"
	ma "github.com/hardstylez72/bblog/internal/api/model/article"
	"github.com/hardstylez72/bblog/internal/storage/user"
	"net/http"
	"time"
)

func (c articleController) GetArticlesByPeriodHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	fromString := r.URL.Query().Get("from")
	from, err := time.Parse(time.RFC3339, fromString)
	if err != nil {
		controller.ResponseWithError(controller.ErrInvalidInputParams(err), http.StatusBadRequest, w)
		return
	}

	toString := r.URL.Query().Get("to")
	to, err := time.Parse(time.RFC3339, toString)
	if err != nil {
		controller.ResponseWithError(controller.ErrInvalidInputParams(err), http.StatusBadRequest, w)
		return
	}

	articles, err := c.articleStorage.GetArticlesByPeriod(ctx, from, to)
	if err != nil {
		if err == user.ErrNotFound {
			controller.ResponseWithError(ErrArticleNotFound(err), http.StatusNotFound, w)
		} else {
			controller.ResponseWithError(controller.ErrInternal(err), http.StatusInternalServerError, w)
		}
		return
	}

	render.JSON(w, r, ma.NewGetArticlesByPeriodResponse(articles))
}
