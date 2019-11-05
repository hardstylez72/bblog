package store

import (
	"context"
	"contrib.go.opencensus.io/integrations/ocsql"
	"database/sql"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
)

type Postgres string

type pgStore struct {
	DB *sqlx.DB
}

func NewPGStore(cfg Postgres) (Store, error) {
	var err error
	const postgresDriverName = "pgx"
	connString := string(cfg)

	driverName, err := ocsql.Register(postgresDriverName, ocsql.WithOptions(ocsql.TraceOptions{
		AllowRoot:         false,
		Ping:              false,
		RowsNext:          true,
		RowsClose:         true,
		RowsAffected:      false,
		LastInsertID:      false,
		Query:             true,
		QueryParams:       true,
		DefaultAttributes: nil,
	}))
	if err != nil {
		return nil, err
	}

	sqlDB, err := sql.Open(driverName, connString)
	if err != nil {
		return nil, err
	}

	sqlxDB := sqlx.NewDb(sqlDB, postgresDriverName)

	if err = sqlxDB.Ping(); err != nil {
		return nil, err
	}

	return &pgStore{
		DB: sqlxDB,
	}, nil
}


func (pg *pgStore) SaveUser(ctx context.Context, u *UserAgr) error {
	tx, err := pg.DB.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}
	err = pg._saveUser(ctx, tx, &u.User)
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	for _, phone := range u.Phones {
		phone.UserId.Valid = true
		phone.UserId.String = u.User.Id

		err = pg._savePhone(ctx, tx, &phone)
		if err != nil {
			_ = tx.Rollback()
			return err
		}
	}

	for _, email := range u.Emails {
		email.UserId.Valid = true
		email.UserId.String = u.User.Id

		err = pg._saveEmail(ctx, tx, &email)
		if err != nil {
			_ = tx.Rollback()
			return err
		}
	}

	return tx.Commit()
}

func (pg *pgStore) _saveEmail(ctx context.Context, tx *sqlx.Tx, u *Email) error {

	query := `
		insert into bb.emails
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

func (pg *pgStore) _savePhone(ctx context.Context, tx *sqlx.Tx, u *Phone) error {

	query := `
		insert into bb.phones
		(
		  "id",
		  number,
		  user_id
		)
		  values
		(
		    :id,
		    :number,
		    :user_id
	    );`

	_, err := tx.NamedExecContext(ctx, query, u)
	return err
}

func (pg *pgStore) _saveUser(ctx context.Context, tx *sqlx.Tx, u *User) error {

	query := `
		insert into bb.users
		(
		  "id",
		  "registered_at",
		  "external_id",
		  "external_auth_type",
		  "login",
		  "password",
		  "first_name",
		  "last_name",
		  "middle_name",
		  "is_banned",
		  "last_activity_time"
		)
		  values
		(
		    :id,
		    default,
		    :external_id,
		    :external_auth_type,
		    :login,
		    :password,
		    :first_name,
		    :last_name,
		    :middle_name,
		    :is_banned,
			default
	    );`

	_, err := tx.NamedExecContext(ctx, query, u)
	return err
}

func (pg *pgStore) GetUserById(ctx context.Context, id string) (*User, error) {

	query := `select * from bb.users u where u.id = $1`

	u := new(User)
	err := pg.DB.GetContext(ctx, u, query, id)
	if err != nil {
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
	err := pg.DB.GetContext(ctx, u, query, id, authTypeId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return u, nil
}