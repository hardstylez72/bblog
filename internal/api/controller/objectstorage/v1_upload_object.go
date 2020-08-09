package objectstorage

import (
	"errors"
	"github.com/go-chi/render"
	"github.com/hardstylez72/bblog/internal/api/controller"
	"github.com/hardstylez72/bblog/internal/objectstorage"
	"mime/multipart"
	"net/http"
	"strconv"
	"time"
)

const (
	numberOfMegabytes = 1
	megabyte          = 1024 * 1024
	maxObjectSize     = megabyte * numberOfMegabytes
	maxFileNumber     = 5
)

type UploadedFile struct {
	Url  string `json:"url"`
	Name string `json:"name"`
}
type UploadObjectResponse struct {
	Files []UploadedFile `json:"files"`
}

func ErrObjectTooBig(err error) controller.Error {
	return controller.Error{
		Inner:   err,
		Message: "object is too big, max size is: " + strconv.Itoa(numberOfMegabytes) + " mb",
	}
}

func (c objectStorageController) UploadObjectHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	err := r.ParseMultipartForm(maxObjectSize)
	if err != nil {
		controller.ResponseWithError(ErrObjectTooBig(err), http.StatusBadRequest, w)
		return
	}

	multipartForm := r.MultipartForm
	if multipartForm == nil {
		err = errors.New("request type is not multipartForm")
		controller.ResponseWithError(controller.ErrInvalidInputParams(err), http.StatusBadRequest, w)
		return
	}

	files := multipartForm.File["files"]

	if len(files) > maxFileNumber {
		err = errors.New("number of files is over max number: " + strconv.Itoa(maxFileNumber))
		controller.ResponseWithError(controller.ErrInvalidInputParams(err), http.StatusBadRequest, w)
		return
	}

	uploadedFiles := make([]UploadedFile, 0)

	for _, file := range files {

		f, err := file.Open()
		if err != nil {
			continue
		}
		defer f.Close()

		url, err := c.objectStorage.Upload(ctx, convertAs(file, f))
		if err != nil {
			continue
		}

		uploadedFiles = append(uploadedFiles, UploadedFile{
			Url:  url,
			Name: file.Filename,
		})
	}

	responseBody := &UploadObjectResponse{
		Files: uploadedFiles,
	}

	render.JSON(w, r, responseBody)
}

func convertAs(fileHeader *multipart.FileHeader, source multipart.File) objectstorage.Image {
	return objectstorage.Image{
		Name:      fileHeader.Filename,
		CreatedAt: time.Now(),
		MimeType:  fileHeader.Header.Get("dd"),
		IsPublic:  false,
		Source:    source,
	}
}
