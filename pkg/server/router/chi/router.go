package chi

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/kamilsk/click/pkg/server/router"
	internal "github.com/kamilsk/click/pkg/server/router/chi/middleware"
)

// NewRouter returns configured `github.com/go-chi/chi` router.
func NewRouter(api router.Server) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Recoverer)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)

	r.Route("/api/v1", func(r chi.Router) {
		r.Route("/{ID}", func(r chi.Router) {
			r.Use(internal.Pack("ID", "id"))
			r.Get("/", api.GetV1)
		})
		r.Get("/pass", api.Pass)
	})
	r.Get("/*", api.Redirect)

	return r
}
