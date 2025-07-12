package api

import (
	"shop-search/internal/handlers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewRouter() *chi.Mux {
	r := chi.NewRouter()
	h := handlers.NewHandler()

	// Middleware
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)

	// Health check
	r.Get("/health", h.HealthCheck)

	// API routes
	r.Get("/api/v1/find", h.FindInAll)

	return r
}
