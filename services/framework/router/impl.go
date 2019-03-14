package router

import "github.com/go-chi/chi"

var (
	all []Router
)

// Router every routing should implement
type Router interface {
	Routes(*chi.Mux)
}

// Register register router
func Register(r Router) {
	all = append(all, r)
}
