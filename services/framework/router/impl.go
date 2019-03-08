package router

import "github.com/go-chi/chi"

var (
	all []Router
)

type Router interface {
	Routes(*chi.Mux)
}

func Register(r Router) {
	all = append(all, r)
}
