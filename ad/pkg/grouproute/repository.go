package grouproute

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

func (r *repository) deletePair(ctx context.Context, tx *sqlx.Tx, groupId, routeId int) error {
	query := `delete from ad.groups_routes where route_id = $1 and group_id = $2`

	_, err := tx.ExecContext(ctx, query, routeId, groupId)
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) Delete(ctx context.Context, params []params) error {
	tx, err := r.conn.BeginTxx(ctx, nil)
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		} else {
			_ = tx.Commit()
		}
	}()

	for _, pair := range params {

		err := r.deletePair(ctx, tx, pair.GroupId, pair.RouteId)
		if err != nil {
			return err
		}

	}
	return nil
}

func (r *repository) insertPair(ctx context.Context, tx *sqlx.Tx, groupId, routeId int) (*Route, error) {
	query := `
		with insert_row as (
			insert into ad.groups_routes (
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

func (r *repository) Insert(ctx context.Context, params []params) ([]Route, error) {

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

func (r *repository) ListNotInGroup(ctx context.Context, groupId int) ([]Route, error) {
	query := `
		select r.id,
			   r.route,
		       r.method,
			   r.description,
			   r.created_at,
			   r.updated_at,
			   r.deleted_at
		from ad.routes r
    left join ad.groups_routes rg on rg.route_id = r.id
        where r.id not in (select route_id from ad.groups_routes where group_id = $1)
          and deleted_at is null
`
	routes := make([]Route, 0)
	err := r.conn.SelectContext(ctx, &routes, query, groupId)
	if err != nil {
		return nil, err
	}

	return routes, nil
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
    left join ad.groups_routes rg on rg.route_id = r.id
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
