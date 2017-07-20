package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ApiError interface {
	Serve(w http.ResponseWriter)
	ServeAndLog(w http.ResponseWriter, err error)
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
