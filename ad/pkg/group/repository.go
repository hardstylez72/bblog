package group

import (
	"context"
	"github.com/jmoiron/sqlx"
)

type repository struct {
	conn *sqlx.DB
}

func NewRepository(conn *sqlx.DB) *repository {
	return &repository{conn: conn}
}

func (r *repository) Insert(ctx context.Context, group *Group) (*Group, error) {
	query := `
insert into ad.groups (
                       code,
                       description,
                       created_at,
                       updated_at,
                       deleted_at
                       )
                   values (
                       :code,
                       :description,
                       now(),
                       null,
                       null
                   ) returning id,
                               code,
                               description,
                               created_at,
                               updated_at,
                               deleted_at;
`

	rows, err := r.conn.NamedQueryContext(ctx, query, group)
	if err != nil {
		return nil, err
	}

	var g Group
	for rows.Next() {
		err = rows.StructScan(&g)
		if err != nil {
			return nil, err
		}
	}

	return &g, nil
}

func (r *repository) List(ctx context.Context) ([]Group, error) {
	query := `
		select id,
			   code,
			   description,
			   created_at,
			   updated_at,
			   deleted_at
		from ad.groups;
`
	groups := make([]Group, 0)
	err := r.conn.SelectContext(ctx, &groups, query)
	if err != nil {
		return nil, err
	}

	return groups, nil
}
