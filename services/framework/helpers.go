package framework

import (
	"encoding/json"
	"net/http"
)

const (
	headerContentType = "Content-Type"
	jsonMIME          = "application/json;charset=UTF-8"
)

type simpleErr struct {
	Err string `json:"error"`
}

// JSON is a helper function to write an json in output
func JSON(w http.ResponseWriter, code int, i interface{}) {
	w.Header().Set(headerContentType, jsonMIME)
	w.WriteHeader(code)
	enc := json.NewEncoder(w)
	_ = enc.Encode(i)
}

// JSON is a helper function to write an json in output
func JSONErr(w http.ResponseWriter, code int, err error) {
	w.Header().Set(headerContentType, jsonMIME)
	w.WriteHeader(code)
	enc := json.NewEncoder(w)
	if err != nil {
		_ = enc.Encode(simpleErr{err.Error()})
	}
}
