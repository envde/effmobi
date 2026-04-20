package transport

import (
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	httpSwagger "github.com/swaggo/http-swagger"

	_ "github.com/envde/effmobi/interanl/docs"
	db "github.com/envde/effmobi/interanl/pkg/postgres"
	"github.com/envde/effmobi/interanl/service"
	"github.com/envde/effmobi/interanl/transport/handlers"
)

func NewRouter(queries *db.Queries, log *slog.Logger) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/swagger/*", httpSwagger.Handler())

	svc := service.NewSubscriptionService(queries, log)
	h := handlers.NewSubscriptionHandler(svc)

	r.Route("/subscriptions", func(r chi.Router) {
		r.Post("/", h.Create)
		r.Get("/", h.List)
		r.Get("/{id}", h.Get)
		r.Put("/{id}", h.Update)
		r.Delete("/{id}", h.Delete)
		r.Post("/sum", h.Sum)
	})

	return r
}
