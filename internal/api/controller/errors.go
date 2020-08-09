package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Error struct {
	Inner   error
	Message string
}

type ErrorResponse struct {
	Inner struct {
		Message string
		Stack   string
	}
	Message string
}

func ResponseWithError(err Error, httpStatusCode int, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatusCode)

	e := ErrorResponse{
		Inner: struct {
			Message string
			Stack   string
		}{
			Message: err.Inner.Error(),
			Stack:   FormatErr(err.Inner),
		},
		Message: err.Message,
	}
	_ = json.NewEncoder(w).Encode(e)
}

func ErrInternal(err error) Error {
	return Error{
		Inner:   err,
		Message: "Internal error",
	}
}

func ErrInvalidInputParams(err error) Error {
	return Error{
		Inner:   err,
		Message: "Invalid input params",
	}
}

func FormatErr(err error) string {
	return fmt.Sprintf("%+v", err)
}
