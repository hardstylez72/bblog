package user

import (
	"context"
	"database/sql"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
)

type pgStore struct {
	db *sqlx.DB
}

func NewPGStorage(db *sqlx.DB) *pgStore {
	return &pgStore{
		db: db,
	}
}

func (pg *pgStore) SaveUser(ctx context.Context, u *User) error {

	query := `
		insert into bb.users
		(
		  "id",
		  "registered_at",
		  "external_id",
		  "external_auth_type",
		  "login",
		  "name",
		  "is_banned",
		  "role_id",
		  "email"
		)
		  values
		(
		    :id,
		    default,
		    :external_id,
		    :external_auth_type,
		    :login,
		    :name,
		    :is_banned,
		    (select id from bb.roles where code = :role_code),
		    :email
	    );`

	_, err := pg.db.NamedExecContext(ctx, query, u)
	return err
}

func (pg *pgStore) GetUserById(ctx context.Context, id string) (*User, error) {

	query := `select id,
				   registered_at,
				   external_auth_type,
				   external_id,
				   email,
				   login,
				   name,
				   is_banned,
				   (select code from bb.roles where id = role_id)  as role_code
			    from bb.users u 
  			   where u.id = $1`

	u := new(User)
	err := pg.db.GetContext(ctx, u, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrNotFound
		}
		return nil, err
	}

	return u, nil
}

func (pg *pgStore) GetUserByExternalId(ctx context.Context, id, authTypeId string) (*User, error) {

	query := `select id,
				   registered_at,
				   external_auth_type,
				   external_id,
				   email,
				   login,
				   name,
				   is_banned,
				   (select code from bb.roles where id = role_id)  as role_code
 			    from bb.users u
		       where u.external_id = $1
		         and u.external_auth_type = $2;`

	u := new(User)
	err := pg.db.GetContext(ctx, u, query, id, authTypeId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrNotFound
		}

		return nil, err
	}

	return u, nil
}
