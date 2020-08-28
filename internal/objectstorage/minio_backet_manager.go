package objectstorage

import (
	"context"
	"github.com/minio/minio-go/v7"
	"time"
)

type minioBucketManager struct {
	client *minio.Client
}

func NewMinioBucketManager(client *minio.Client) *minioBucketManager {
	return &minioBucketManager{
		client: client,
	}
}

func (m *minioBucketManager) ResolveBucket(ctx context.Context, t time.Time) (string, error) {
	return m.resolveBucket(ctx, formBucketName(t))
}

func (m *minioBucketManager) resolveBucket(ctx context.Context, bucketName string) (string, error) {
	isBucketExists, err := m.client.BucketExists(ctx, bucketName)
	if err != nil {
		return "", err
	}

	if isBucketExists {
		return bucketName, nil
	}

	bucketOptions := minio.MakeBucketOptions{
		Region:        "",
		ObjectLocking: false,
	}

	err = m.client.MakeBucket(ctx, bucketName, bucketOptions)
	if err != nil {
		return "", err
	}

	return bucketName, nil
}

func (m *minioBucketManager) ResolvePublicBucket(ctx context.Context, t time.Time) (string, error) {
	bucketName := formPublicBucketName(t)

	bucketName, err := m.resolveBucket(ctx, bucketName)
	if err != nil {
		return "", err
	}

	policy := `{"Version":"2012-10-17","Statement":[{"Effect":"Allow","Principal":{"AWS":["*"]},"Action":["s3:GetBucketLocation","s3:ListBucket"],"Resource":["arn:aws:s3:::` + bucketName + `"]},{"Effect":"Allow","Principal":{"AWS":["*"]},"Action":["s3:GetObject"],"Resource":["arn:aws:s3:::` + bucketName + `/*"]}]}`

	err = m.client.SetBucketPolicy(ctx, bucketName, policy)
	if err != nil {
		return "", err
	}
	return bucketName, nil
}

func formBucketName(t time.Time) string {
	return t.Format("2006")
}

func formPublicBucketName(t time.Time) string {
	return t.Format("2006-public")
}
