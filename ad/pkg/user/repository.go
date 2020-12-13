package user

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

func (r *repository) Insert(ctx context.Context, entity *User) (*User, error) {
	query := `
insert into ad.users (
                       external_id,
					   is_system,
					   name,
					   description,
					   email,
					   phone,
                       created_at,
					   updated_at,
					   deleted_at
                       )
                   values (
                       :external_id,
					   :is_system,
					   :name,
					   :description,
					   :email,
					   :phone,
					   now(),
					   :updated_at,
					   :deleted_at
                   ) returning id,
                               external_id,
							   is_system,
							   name,
							   description,
							   email,
							   phone,
							   created_at,
							   updated_at,
							   deleted_at;
`

	rows, err := r.conn.NamedQueryContext(ctx, query, entity)
	if err != nil {
		return nil, err
	}

	var g User
	for rows.Next() {
		err = rows.StructScan(&g)
		if err != nil {
			return nil, err
		}
	}

	return &g, nil
}

func (r *repository) List(ctx context.Context) ([]User, error) {
	query := `
		select id,
			   external_id,
			   is_system,
			   name,
			   description,
			   email,
			   phone,
		       created_at,
			   updated_at,
			   deleted_at
		from ad.users
	   where deleted_at is null;
`
	groups := make([]User, 0)
	err := r.conn.SelectContext(ctx, &groups, query)
	if err != nil {
		return nil, err
	}

	return groups, nil
}

func (r *repository) Delete(ctx context.Context, id int) error {
	query := `
		update ad.users 
			set deleted_at = now()
		where id = $1;
`
	_, err := r.conn.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}
