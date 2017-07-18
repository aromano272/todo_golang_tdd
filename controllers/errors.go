package controllers

import (
	"fmt"
	"net/http"
	"encoding/json"
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
	error string
}


func (error *ApiError) Serve(w http.ResponseWriter, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)
	fmt.Fprintln(w, error)
	json.NewEncoder(w).Encode(error)
}
