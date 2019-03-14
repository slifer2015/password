package router

import (
	"context"
	"net/http"
	"sync"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/go-chi/chi/middleware"

	"test.com/test/services/initializer"

	"test.com/test/services/assert"

	"github.com/go-chi/chi"
)

var (
	once = sync.Once{}
)

type initer struct {
}

func (initer) Initial(ctx context.Context) {
	once.Do(func() {
		r := chi.NewRouter()
		r.Use(middleware.RequestID)
		r.Use(middleware.RealIP)
		r.Use(middleware.Logger)
		r.Use(middleware.Recoverer)
		for i := range all {
			all[i].Routes(r)
		}
		go func() {
			assert.Nil(http.ListenAndServe(viper.GetString("port"), r))
		}()
		logrus.Warnf("listen on port %s", viper.GetString("port"))

	})
}

func init() {
	initializer.Register(initer{}, 1000)
}
