package storage

import (
	"contrib.go.opencensus.io/integrations/ocsql"
	"database/sql"
	"github.com/jmoiron/sqlx"
	"time"
)

type TimeTamps struct {
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deletedAt"`
}

func NewPGConnection(connString string) (*sqlx.DB, error) {
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

	sqlxDB := sqlx.NewDb(sqlDB, postgresDriverName)

	if err = sqlxDB.Ping(); err != nil {
		return nil, err
	}

	return sqlxDB, nil
}
