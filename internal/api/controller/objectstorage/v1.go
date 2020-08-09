package objectstorage

import (
	"github.com/go-chi/chi"
	"github.com/hardstylez72/bblog/internal/objectstorage"
)

type objectStorageController struct {
	objectStorage objectstorage.Storage
}

func NewObjectStorageController(objectStorage objectstorage.Storage) *objectStorageController {
	return &objectStorageController{
		objectStorage: objectStorage,
	}
}

func (c objectStorageController) Mount(r chi.Router) {
	r.Post("/v1/object-storage/upload", c.UploadObjectHandler)
}
