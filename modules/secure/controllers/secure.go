package controllers

import (
	"github.com/go-chi/chi"
	"test.com/test/services/framework/controller"
	"test.com/test/services/framework/middleware"
	"test.com/test/services/framework/router"
)

type ctrl struct {
	controller.Base
}

func (ctrl) Routes(m *chi.Mux) {
	m.Post("/generate", controller.Mix(
		generate,
		middleware.PayloadUnmarshallerGenerator(generatePayload{}),
	))
}

func init() {
	router.Register(ctrl{})
}
