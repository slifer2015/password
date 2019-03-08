package controller

import (
	"gopkg.in/go-playground/validator.v9"
	"net/http"
	"reflect"
	"test.com/test/services/assert"
)

// TODO : all more functionality
type Base struct {
}

type ValidationErr map[string]string

func Validate(i interface{}) (map[string]string, bool) {
	var validationErr = make(map[string]string)
	var failed bool
	v := validator.New()
	err := v.Struct(i)
	if err != nil {
		failed = true
		for _, err := range err.(validator.ValidationErrors) {
			field, ok := reflect.TypeOf(i).Elem().FieldByName(err.Field())
			assert.True(ok)
			validationErr[field.Tag.Get("json")] = field.Tag.Get("message")
		}

	}
	return validationErr, failed
}

// Mix try to mix all middleware with the calling route
func Mix(final http.HandlerFunc, all ...func(c http.HandlerFunc) http.HandlerFunc) http.HandlerFunc {
	res := final
	for i := len(all) - 1; i >= 0; i-- {
		res = all[i](res)
	}

	return res
}