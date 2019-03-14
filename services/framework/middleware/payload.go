package middleware

import (
	"context"
	"encoding/json"
	"net/http"
	"reflect"
	"strings"

	"test.com/test/services/array"
	"test.com/test/services/assert"
	"test.com/test/services/framework"
)

type contextKey string

const (
	// ContextBody is the context key for the body unmarshalled object
	ContextBody contextKey = "_body"
)

// PayloadUnmarshallerGenerator unmarshaller middleware generator
func PayloadUnmarshallerGenerator(pattern interface{}) func(http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			method := strings.ToUpper(r.Method)
			ok := array.StringInArray(method, "DELETE", "GET")
			assert.True(!ok)
			cp := reflect.New(reflect.TypeOf(pattern)).Elem().Addr().Interface()
			decoder := json.NewDecoder(r.Body)
			err := decoder.Decode(cp)
			if err != nil {
				w.Header().Set("error", "invalid request body")
				e := struct {
					Error string `json:"error"`
				}{
					Error: "invalid request body",
				}
				framework.JSON(w, http.StatusBadRequest, e)
				return
			}
			ctx := context.WithValue(r.Context(), ContextBody, cp)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

// GetPayload from the request
func GetPayload(r *http.Request) (interface{}, bool) {
	t := r.Context().Value(ContextBody)
	return t, t != nil
}
