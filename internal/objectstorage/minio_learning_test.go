package objectstorage

import (
	"bytes"
	"context"
	"io/ioutil"
	"os"
	path2 "path"
	"testing"
	"time"
)

var config = Config{
	Host:            "localhost:9000",
	AccessKeyID:     "OR7DZ0AQRAEP1EGIXBXG",
	SecretAccessKey: "dJSQvadUHWz6rxuRwipsQHBm3Z1XacIkdKqYFyUP",
	UseSSL:          false,
}

func TestNewMinioStorage(t *testing.T) {
	client, err := NewMinioClient(config)
	if err != nil {
		t.Fatal(err)
	}

	storage := NewMinioStorage(client)

	imageSource, err := readImageFromDisc("/testdata/e8531975a9a2f49caabe3dad4d12470c787f3cf5r1-720-759v2_uhq.jpg")
	if err != nil {
		t.Fatal(err)
	}

	image := Image{
		Name:      "test-image.jpeg",
		CreatedAt: time.Now(),
		MimeType:  "image/png",
		Source:    bytes.NewBuffer(imageSource),
	}

	_, err = storage.Upload(context.TODO(), image)
	if err != nil {
		t.Fatal(err)
	}
}

func readImageFromDisc(path string) ([]byte, error) {
	curDir, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	f, err := os.Open(path2.Join(curDir, path))
	if err != nil {
		return nil, err
	}

	fb, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}
	return fb, nil
}
