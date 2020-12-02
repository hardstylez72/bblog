package user

import (
	"context"
	"database/sql"
	//_ "github.com/jackc/pgx/stdlib"
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
		insert into ad.users
		(
		  "id",
		  "external_id",
		  "auth_type",
		 
		  "login",
		  "name",
		  "email",
		  
		  created_at
		)
		  values
		(
		    :id,
		    :external_id,
		    :auth_type,
		 
		    :login,
		    :name,
		    :email,
		 
		    default
	    );`

	_, err := pg.db.NamedExecContext(ctx, query, u)
	return err
}

func (pg *pgStore) GetUserById(ctx context.Context, id int) (*User, error) {

	query := `select id,
				   auth_type,
				   external_id,
       
				   "email",
				   "login",
				   "name",
       
      			   created_at,
				   updated_at,
				   deleted_at
			    from ad.users u 
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
				   auth_type,
				   external_id,
       
				   "email",
				   "login",
				   "name",
       
      			   created_at,
				   updated_at,
				   deleted_at
 			    from ad.users u
		       where u.external_id = $1
		         and u.auth_type = $2;`

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

func (pg *pgStore) GetUsersAmount(ctx context.Context) (int, error) {

	query := `select count(*) from ad.users`
	var amount int

	err := pg.db.GetContext(ctx, amount, query)
	if err != nil {
		return -1, err
	}

	return amount, nil
}
