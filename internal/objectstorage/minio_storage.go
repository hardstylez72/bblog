package objectstorage

import (
	"context"
	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
)

const objectStreamType = -1

type minioStorage struct {
	client        *minio.Client
	bucketManager *minioBucketManager
}

func NewMinioStorage(client *minio.Client) *minioStorage {
	return &minioStorage{
		client:        client,
		bucketManager: NewMinioBucketManager(client),
	}
}

func (s *minioStorage) Upload(ctx context.Context, image Image) (url string, err error) {

	var bucketName string
	if image.IsPublic {
		bucketName, err = s.bucketManager.ResolvePublicBucket(ctx, image.CreatedAt)
		if err != nil {
			return "", nil
		}
	} else {
		bucketName, err = s.bucketManager.ResolveBucket(ctx, image.CreatedAt)
		if err != nil {
			return "", nil
		}
	}

	options := minio.PutObjectOptions{
		ContentType:  image.MimeType,
		UserMetadata: map[string]string{}, // todo: refill
	}
	objectName := uuid.New().String()

	uploadedObject, err := s.client.PutObject(ctx, bucketName, objectName, image.Source, objectStreamType, options)
	if err != nil {
		return "", nil
	}

	return uploadedObject.Location, err
}
