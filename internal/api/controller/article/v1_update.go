package article

import (
	"encoding/json"
	"github.com/go-chi/render"
	"github.com/hardstylez72/bblog/internal/api/controller"
	ma "github.com/hardstylez72/bblog/internal/api/model/article"
	"github.com/hardstylez72/bblog/internal/storage/article"
	"net/http"
)

func (c articleController) UpdateArticleHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var a ma.UpdateArticleRequest
	err := json.NewDecoder(r.Body).Decode(&a)
	if err != nil {
		controller.ResponseWithError(controller.ErrInvalidInputParams(err), http.StatusBadRequest, w)
		return
	}

	err = c.validator.Struct(&a)
	if err != nil {
		controller.ResponseWithError(controller.ErrInvalidInputParams(err), http.StatusBadRequest, w)
		return
	}

	id, err := c.articleStorage.UpdateArticle(ctx, convertUpdateArticleRequestToDb(&a))
	if err != nil {
		controller.ResponseWithError(controller.ErrInternal(err), http.StatusInternalServerError, w)
		return
	}
	render.JSON(w, r, ma.NewUpdateArticleResponse(id))
}

func convertUpdateArticleRequestToDb(a *ma.UpdateArticleRequest) *article.Article {

	return &article.Article{
		Id:      a.Id,
		Body:    a.Body,
		Title:   a.Title,
		UserId:  a.UserId,
		Preface: a.Preface,
	}
}
