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

func (pg *pgStore) SaveUserWithEmail(ctx context.Context, u *UserAgr) error {
	tx, err := pg.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}
	err = pg.SaveUser(ctx, tx, &u.User)
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	for _, email := range u.Emails {
		email.UserId.Valid = true
		email.UserId.String = u.User.Id

		err = pg.SaveEmail(ctx, tx, &email)
		if err != nil {
			_ = tx.Rollback()
			return err
		}
	}

	return tx.Commit()
}

func (pg *pgStore) SaveEmail(ctx context.Context, tx *sqlx.Tx, u *Email) error {

	query := `
		insert into bb.user_emails
		(
		  "id",
		  address,
		  user_id
		)
		  values
		(
		    :id,
		    :address,
		    :user_id
	    );`

	_, err := tx.NamedExecContext(ctx, query, u)
	return err
}

func (pg *pgStore) SaveUser(ctx context.Context, tx *sqlx.Tx, u *User) error {

	query := `
		insert into bb.users
		(
		  "id",
		  "registered_at",
		  "external_id",
		  "external_auth_type",
		  "login",
		  "first_name",
		  "last_name",
		  "middle_name",
		  "is_banned"
		)
		  values
		(
		    :id,
		    default,
		    :external_id,
		    :external_auth_type,
		    :login,
		    :first_name,
		    :last_name,
		    :middle_name,
		    :is_banned
	    );`

	_, err := tx.NamedExecContext(ctx, query, u)
	return err
}

func (pg *pgStore) GetUserById(ctx context.Context, id string) (*User, error) {

	query := `select * from bb.users u where u.id = $1`

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

	query := `select *
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
