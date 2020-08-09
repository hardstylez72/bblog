package article

import (
	"github.com/hardstylez72/bblog/internal/storage/article"
)

func NewGetArticlesByPeriodResponse(in []article.Article) []Article {
	out := make([]Article, 0, len(in))
	for _, i := range in {
		el := Article{
			Id:        i.Id,
			Preface:   i.Preface,
			Title:     i.Title,
			UserId:    i.UserId,
			CreatedAt: i.CreatedAt,
			UpdatedAt: nil,
			DeletedAt: nil,
		}
		out = append(out, el)
	}
	return out
}
