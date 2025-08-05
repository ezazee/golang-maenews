package router

import (
	"maenews/backend/handlers"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func SetupRouter() http.Handler {
	r := chi.NewRouter()

	// Middleware dasar
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Mengatur CORS agar frontend Next.js (dari port 3000) bisa mengakses API ini
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	r.Route("/api/v1", func(r chi.Router) {
		// Rute Artikel
		r.Get("/articles", handlers.GetAllArticlesHandler)
		r.Get("/articles/{slug}", handlers.GetArticleBySlugHandler)
		r.Post("/articles/{slug}/view", handlers.IncrementViewHandler)
		r.Get("/category/{categoryName}", handlers.GetArticlesByCategoryHandler)
		r.Get("/tag/{tagName}", handlers.GetArticlesByTagHandler)
		r.Get("/trending", handlers.GetTrendingItemsHandler)
		r.Get("/events/upcoming", handlers.GetUpcomingEventsHandler)
		r.Get("/events/{slug}", handlers.GetEventBySlugHandler)
		r.Get("/search/{query}", handlers.SearchArticlesHandler)
	})

	return r
}
