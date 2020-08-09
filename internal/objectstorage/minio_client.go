package objectstorage

import (
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type Config struct {
	Host            string
	AccessKeyID     string
	SecretAccessKey string
	UseSSL          bool
}

func NewMinioClient(config Config) (*minio.Client, error) {
	return minio.New(config.Host, &minio.Options{
		Creds:  credentials.NewStaticV4(config.AccessKeyID, config.SecretAccessKey, ""),
		Secure: config.UseSSL,
	})
}
