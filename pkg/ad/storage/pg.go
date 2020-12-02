package storage

import (
	"errors"
	"gorm.io/gorm"
)

type store struct {
	db *gorm.DB
}

var (
	ErrCantMigrateStruct = func(err error) error {
		return errors.New("error while doing migration: " + err.Error())
	}
)

func NewStorage(db *gorm.DB) *store {
	return &store{db: db}
}

func (s *store) Init() error {

	conn, err := s.db.DB()
	if err != nil {
		return ErrCantMigrateStruct(err)
	}

	_, err = conn.Exec("create schema if not exists ad;")
	if err != nil {
		return ErrCantMigrateStruct(err)
	}

	err = s.db.Migrator().CreateTable(&Route{}, "ad.routes")
	if err != nil {
		return ErrCantMigrateStruct(err)
	}
	return nil
}

/*


 */
