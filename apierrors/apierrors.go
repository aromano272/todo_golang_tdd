package apierrors

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	InvalidId             = "The provided id is invalid"
	IdNotFound            = "This item doesn't exist"
	IdFieldMissing        = "The request is missing the id"
	TodoTitleFieldMissing = "The title field cannot be empty"
)

type ApiError interface {
	Serve(w http.ResponseWriter)
	ServeAndLog(w http.ResponseWriter, err error)

	GetError() string
	GetStatusCode() int
}

type simpleError struct {
	Error string `json:"error"`
	statusCode int
}

func NewApiError(str string, statusCode int) ApiError {
	return &simpleError{str, statusCode}
}

func (apierr *simpleError) Serve(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(apierr.statusCode)
	json.NewEncoder(w).Encode(apierr)
}

func (apierr *simpleError) ServeAndLog(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(apierr.statusCode)
	json.NewEncoder(w).Encode(apierr)
	fmt.Println(err)
}

func (apierr *simpleError) GetError() string {
	return apierr.Error
}

func (apierr *simpleError) GetStatusCode() int {
	return apierr.statusCode
}
