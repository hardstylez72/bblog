package article

import (
	"github.com/hardstylez72/bblog/internal/storage/article"
	"time"
)

type ArticleWithBody struct {
	Article
	Body string `json:"body" validate:"required"`
}

type Article struct {
	Id        string     `json:"id"`
	Title     string     `json:"title" validate:"required"`
	UserId    string     `json:"userId" validate:"required"`
	Preface   string     `json:"preface" validate:"required"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt"`
}

func NewGetArticleByIdResponse(in *article.Article) *ArticleWithBody {
	out := &ArticleWithBody{
		Article: Article{
			Preface:   in.Preface,
			Id:        in.Id,
			Title:     in.Title,
			UserId:    in.UserId,
			CreatedAt: in.CreatedAt,
		},
		Body: in.Body,
	}

	if in.UpdatedAt.Valid {
		out.UpdatedAt = &in.UpdatedAt.Time
	}

	if in.DeletedAt.Valid {
		out.DeletedAt = &in.DeletedAt.Time
	}

	return out
}
