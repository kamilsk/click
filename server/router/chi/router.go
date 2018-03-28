package chi

import (
	"net/http"

	common "github.com/kamilsk/click/server/middleware"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/kamilsk/click/server/router"
)

// NewRouter returns configured `github.com/go-chi/chi` router.
func NewRouter(api router.Server) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Recoverer)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)

	notImplemented := func(rw http.ResponseWriter, req *http.Request) { rw.WriteHeader(http.StatusNotImplemented) }

	r.Route("/api/v1", func(r chi.Router) {
		r.Post("/", notImplemented)

		r.Route("/{UUID}", func(r chi.Router) {
			r.Use(ctxPacker(common.Link, "UUID"))

			r.Get("/", api.GetV1)
			r.Put("/", notImplemented)
		})
	})

	r.Get("/*", api.Redirect)

	return r
}
