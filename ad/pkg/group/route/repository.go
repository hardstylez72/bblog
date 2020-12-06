package route

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

func (r *repository) Insert(ctx context.Context, params []insertParams) ([]Route, error) {

	tx, err := r.conn.BeginTxx(ctx, nil)
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		} else {
			_ = tx.Commit()
		}
	}()

	routes := make([]Route, 0)
	for _, pair := range params {

		route, err := r.insertPair(ctx, tx, pair.GroupId, pair.RouteId)
		if err != nil {
			return nil, err
		}
		routes = append(routes, *route)
	}

	return routes, nil
}

func (r *repository) insertPair(ctx context.Context, tx *sqlx.Tx, groupId, routeId int) (*Route, error) {
	query := `
		with insert_row as (
			insert into ad.routes_groups (
					   route_id,
					   group_id
					   )
				   values (
					   $1,
					   $2
				   )
		)
		select r.id,
			   r.route,
		       r.method,
			   r.description,
			   r.created_at,
			   r.updated_at,
			   r.deleted_at
		from ad.routes r where r.id = $1;
`

	rows := tx.QueryRowxContext(ctx, query, routeId, groupId)
	var route Route
	err := rows.StructScan(&route)
	if err != nil {
		return nil, err
	}

	return &route, nil
}

func (r *repository) List(ctx context.Context, groupId int) ([]Route, error) {
	query := `
		select rg.route_id as id,
			   r.route,
		       r.method,
			   r.description,
			   r.created_at,
			   r.updated_at,
			   r.deleted_at
		from ad.routes r
    left join ad.routes_groups rg on rg.route_id = r.id
        where rg.group_id = $1 
          and deleted_at is null
`
	routes := make([]Route, 0)
	err := r.conn.SelectContext(ctx, &routes, query, groupId)
	if err != nil {
		return nil, err
	}

	return routes, nil
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
