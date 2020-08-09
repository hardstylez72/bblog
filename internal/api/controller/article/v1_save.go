package article

import (
	"encoding/json"
	"github.com/go-chi/render"
	"github.com/google/uuid"
	"github.com/hardstylez72/bblog/internal/api/controller"
	ma "github.com/hardstylez72/bblog/internal/api/model/article"
	"github.com/hardstylez72/bblog/internal/storage/article"
	"net/http"
)

func (c articleController) SaveArticleHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var a ma.SaveArticleRequest
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

	id, err := c.articleStorage.SaveArticle(ctx, convertSaveArticleRequestToDb(&a))
	if err != nil {
		controller.ResponseWithError(controller.ErrInternal(err), http.StatusInternalServerError, w)
		return
	}
	render.JSON(w, r, ma.NewSaveArticleResponse(id))
}

func convertSaveArticleRequestToDb(a *ma.SaveArticleRequest) *article.Article {

	article := &article.Article{
		Body:    a.Body,
		Title:   a.Title,
		UserId:  a.UserId,
		Preface: a.Preface,
	}

	if a.Id != "" {
		article.Id = a.Id
	} else {
		article.Id = uuid.New().String()
	}

	return article
}
