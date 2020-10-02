package objectstorage

import (
	"context"
	"io"
	"time"
)

type Object struct {
	Bucket string
	Name   string
}

type Image struct {
	Name      string
	CreatedAt time.Time
	MimeType  string
	IsPublic  bool
	Source    io.Reader
}

type Storage interface {
	UploadImage(ctx context.Context, image Image) (url string, err error)
}
