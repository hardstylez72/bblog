package route

import (
	"context"
	"errors"
	"github.com/hardstylez72/bblog/ad/pkg/routetag"
	"github.com/jmoiron/sqlx"
)

var (
	ErrRouteAlreadyExists = errors.New("route already exist")
)

type repository struct {
	conn *sqlx.DB
}

func NewRepository(conn *sqlx.DB) *repository {
	return &repository{conn: conn}
}

func (r *repository) Update(ctx context.Context, route *Route) (*Route, error) {
	query := `
	
			update ad.routes
			   set route = :route,
				   method = :method,
				   description = :description,
				   updated_at = now()
			where id = :id returning id,
						   route,
						   method,
						   description,
						   created_at,
						   updated_at,
						   deleted_at
`

	rows, err := r.conn.NamedQueryContext(ctx, query, route)
	if err != nil {
		return nil, err
	}

	var g Route
	for rows.Next() {
		err = rows.StructScan(&g)
		if err != nil {
			return nil, err
		}
	}

	return &g, nil
}

func (r *repository) Insert(ctx context.Context, route *Route) (*Route, error) {
	query := `
insert into ad.routes (
                       route,
                       method,
                       description,
                       created_at,
                       updated_at,
                       deleted_at
                       )
                   values (
                       :route,
                       :method,
                       :description,
                       now(),
                       null,
                       null
                   ) returning id,
                               route,
                       		   method,
                               description,
                               created_at,
                               updated_at,
                               deleted_at;
`

	rows, err := r.conn.NamedQueryContext(ctx, query, route)
	if err != nil {
		return nil, err
	}

	var g Route
	for rows.Next() {
		err = rows.StructScan(&g)
		if err != nil {
			return nil, err
		}
	}

	return &g, nil
}

func (r *repository) InsertWithTags(ctx context.Context, route *Route, tagNames []string) (*RouteWithTags, error) {
	tx, err := r.conn.BeginTxx(ctx, nil)
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		} else {
			_ = tx.Commit()
		}
	}()
	newRoute, err := InsertTx(ctx, tx, route)
	if err != nil {
		return nil, err
	}

	tagNames, err = routetag.Merge(ctx, r.conn, tx, newRoute.Id, tagNames)
	if err != nil {
		return nil, err
	}
	return &RouteWithTags{
		Route: *newRoute,
		Tags:  tagNames,
	}, nil
}

func InsertTx(ctx context.Context, tx *sqlx.Tx, route *Route) (*Route, error) {
	query := `
insert into ad.routes (
                       route,
                       method,
                       description,
                       created_at,
                       updated_at,
                       deleted_at
                       )
                   values (
                       :route,
                       :method,
                       :description,
                       now(),
                       null,
                       null
                   ) returning id,
                               route,
                       		   method,
                               description,
                               created_at,
                               updated_at,
                               deleted_at;
`
	var g Route
	rows, err := tx.NamedQuery(query, route)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err = rows.StructScan(&g)
		if err != nil {
			return nil, err
		}
	}
	if g.Id == 0 {
		return nil, ErrRouteAlreadyExists
	}

	return &g, nil
}

func (r *repository) List(ctx context.Context) ([]Route, error) {
	query := `
		select id,
			   route,
		       method,
			   description,
			   created_at,
			   updated_at,
			   deleted_at
		from ad.routes
	   where deleted_at is null;
`
	groups := make([]Route, 0)
	err := r.conn.SelectContext(ctx, &groups, query)
	if err != nil {
		return nil, err
	}

	return groups, nil
}

func (r *repository) Delete(ctx context.Context, id int) error {
	query := `
		update ad.routes 
			set deleted_at = now()
		where id = $1;
`
	_, err := r.conn.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}
