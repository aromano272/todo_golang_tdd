package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func jsonErr(err error, descr string) string {
	return fmt.Sprintf(`{"error":"%s", "description":"%s"}`, err.Error(), descr)
}

func jsonErrStr(descr string) string {
	return fmt.Sprintf(`{"error":"%s"}`, descr)
}

func jsonKVP(key, val string) string {
	return fmt.Sprintf(`{"%s":"%s"}`, key, val)
}

type ApiError struct {
	Error string `json:"error"`
}

func (error ApiError) Serve(w http.ResponseWriter, code int) ApiError {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(error)

	return error
}
