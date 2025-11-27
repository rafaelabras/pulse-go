package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/rafaelabras/pulse-go/internal/app"
)

// The route has the responsibility to route everything in the application

func SetupRoutes(app *app.Application) *chi.Mux {
	r := chi.NewRouter()

	r.Get("/health", app.HealthCheck)
	r.Get("/users/{id}", app.UserHandler.HandlerGetUserByID)
	r.Post("/users", app.UserHandler.HandlerCreateUser)
	return r
}
