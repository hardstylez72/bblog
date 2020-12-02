package storage

import (
	"context"
	"github.com/hardstylez72/bblog/internal/storage"
	"gorm.io/gorm"
)

type Route struct {
	gorm.Model
	Id          int
	Path        string
	Method      string
	Entity      string
	Description string
	storage.TimeTamps
}

type Storage interface {
	GetRoutes(ctx context.Context) ([]Route, error)
}
