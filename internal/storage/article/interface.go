package article

import (
	"context"
	"database/sql"
	"time"
)

type Article struct {
	Id        string       `db:"id"`
	Body      string       `db:"body"`
	Title     string       `db:"title"`
	Preface   string       `db:"preface"`
	UserId    string       `db:"user_id"`
	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at"`
	DeletedAt sql.NullTime `db:"deleted_at"`
}

type Storage interface {
	SaveArticle(ctx context.Context, article *Article) (id string, err error)
	UpdateArticle(ctx context.Context, article *Article) (id string, err error)

	GetArticleById(ctx context.Context, id string) (*Article, error)
	GetArticleIdsByPeriod(ctx context.Context, from, to time.Time) ([]string, error)
	GetArticlesByPeriod(ctx context.Context, from, to time.Time) ([]Article, error)

	DeleteArticleById(ctx context.Context, id string) error
}
