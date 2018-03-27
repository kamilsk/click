package chi

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func NewRouter(_ interface{}) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Recoverer)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)

	r.Get("/*", func(rw http.ResponseWriter, req *http.Request) { rw.WriteHeader(http.StatusNotImplemented) })

	return r
}
