package ad

import (
	"context"
	"errors"
	"github.com/jmoiron/sqlx"
	"time"
)

type pgStore struct {
	db *sqlx.DB
}

func NewPgStorage(db *sqlx.DB) *pgStore {
	return &pgStore{db: db}
}

func (p pgStore) UpdateArticle(ctx context.Context, article *Article) (id string, err error) {
	query := `
	   update ad.articles
			set body = :body,
			    title = :title,
			    preface = :preface,
			    updated_at = now()
     	where id = :id
	`
	_, err = p.db.NamedExecContext(ctx, query, article)

	return article.Id, err
}

func (p pgStore) SaveArticle(ctx context.Context, article *Article) (id string, err error) {
	query := `
	    insert into ad.articles
	        (id, body, title, preface, user_id, created_at, updated_at, deleted_at) 
     values (:id, :body, :title, :preface, :user_id, default, null, null) 
     returning id
         
	`
	rows, err := p.db.NamedQueryContext(ctx, query, article)
	for rows.Next() {
		err = rows.Scan(&id)
	}

	if id == "" {
		return "", errors.New("error while adding new article")
	}
	return id, err
}

func (p pgStore) GetArticleById(ctx context.Context, id string) (*Article, error) {
	query := `
	    select id, body, title, preface, user_id, created_at, updated_at, deleted_at
	      from ad.articles
	     where id = $1
	`
	var article Article
	err := p.db.GetContext(ctx, &article, query, id)
	if err != nil {
		return nil, err
	}
	return &article, nil
}

func (p pgStore) GetArticleIdsByPeriod(ctx context.Context, from, to time.Time) ([]string, error) {
	query := `
	    select id
	      from ad.articles
	     where created_at between $1 and $2
	`
	var ids = make([]string, 0)
	err := p.db.SelectContext(ctx, &ids, query, from, to)
	if err != nil {
		return nil, err
	}
	return ids, nil
}

func (p pgStore) GetArticlesByPeriod(ctx context.Context, from, to time.Time) ([]Article, error) {
	query := `
	    select id, preface, title, user_id, created_at, updated_at, deleted_at
	      from ad.articles
	     where created_at between $1 and $2
	`
	var articles = make([]Article, 0)
	err := p.db.SelectContext(ctx, &articles, query, from, to)
	if err != nil {
		return nil, err
	}
	return articles, nil
}

func (p pgStore) DeleteArticleById(ctx context.Context, id string) error {
	query := `
	    update ad.articles
           set is_deleted = now()
         where id = ?
	`
	_, err := p.db.ExecContext(ctx, query, id)

	return err
}
