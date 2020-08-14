package article

import (
	"github.com/go-chi/chi"
	"github.com/go-playground/validator/v10"
	"github.com/hardstylez72/bblog/internal/api/controller"
	"github.com/hardstylez72/bblog/internal/storage/article"
)

type articleController struct {
	articleStorage article.Storage
	validator      *validator.Validate
}

func ErrArticleNotFound(err error) controller.Error {
	return controller.Error{
		Inner:   err,
		Message: "Article is not found",
	}
}

func NewArticleController(articleStorage article.Storage) *articleController {
	return &articleController{
		articleStorage: articleStorage,
		validator:      validator.New(),
	}
}

func (c articleController) Mount(r chi.Router) {
	r.Route("/v1/article", func(r chi.Router) {
		r.Get("/{id}", c.GetArticleByIdHandler)
		r.Delete("/{id}", c.DeleteArticleByIdHandler)
		r.Post("/", c.SaveArticleHandler)
		r.Put("/", c.UpdateArticleHandler)
	})
	r.Get("/v1/articles", c.GetArticlesByPeriodHandler)
}
