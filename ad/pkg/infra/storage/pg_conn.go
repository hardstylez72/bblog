package storage

import (
	"contrib.go.opencensus.io/integrations/ocsql"
	"database/sql"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"time"
)

type TimeTamps struct {
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	DeletedAt time.Time `db:"deleted_at"`
}

func NewPGConnection(connString string) (*sql.DB, error) {
	var err error
	const postgresDriverName = "pgx"

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

	return sqlDB, nil
}

func WrapPgConnWithSqlx(conn *sql.DB) (*sqlx.DB, error) {
	const postgresDriverName = "pgx"
	sqlxDB := sqlx.NewDb(conn, postgresDriverName)

	if err := sqlxDB.Ping(); err != nil {
		return nil, err
	}
	return sqlxDB, nil
}
