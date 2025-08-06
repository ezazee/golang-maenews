package router

import (
	"maenews/backend/handlers"
	auth "maenews/backend/middleware" // Memberi nama alias 'auth' pada middleware kustom
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware" // Middleware bawaan dari Chi
	"github.com/go-chi/cors"
)

func SetupRouter() http.Handler {
	r := chi.NewRouter()

	// Middleware dasar dari package 'middleware' bawaan Chi
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Mengatur CORS
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	// --- Rute Publik ---
	r.Route("/api/v1", func(r chi.Router) {
		r.Get("/articles", handlers.GetAllArticlesHandler)
		r.Get("/articles/{slug}", handlers.GetArticleBySlugHandler)
		r.Post("/articles/{slug}/view", handlers.IncrementViewHandler)
		r.Get("/category/{categoryName}", handlers.GetArticlesByCategoryHandler)
		r.Get("/tag/{tagName}", handlers.GetArticlesByTagHandler)
		r.Get("/search/{query}", handlers.SearchArticlesHandler)
		r.Get("/trending", handlers.GetTrendingItemsHandler)
		r.Get("/events/upcoming", handlers.GetUpcomingEventsHandler)

		// PERBAIKAN: Memastikan rute untuk mengambil detail satu event sudah ada
		r.Get("/events/{slug}", handlers.GetEventBySlugHandler)
	})

	// --- Rute Otentikasi ---
	r.Post("/api/v1/register", handlers.RegisterHandler)
	r.Post("/api/v1/login", handlers.LoginHandler)

	// --- Rute Admin yang Dilindungi ---
	r.Route("/api/v1/admin", func(r chi.Router) {
		r.Use(auth.JWTMiddleware)

		// Rute CRUD untuk Artikel
		r.Post("/articles", handlers.CreateArticleHandler)
		r.Put("/articles/{slug}", handlers.UpdateArticleHandler)
		r.Delete("/articles/{slug}", handlers.DeleteArticleHandler)

		// Rute CRUD untuk Event
		r.Post("/events", handlers.CreateEventHandler)
		r.Put("/events/{slug}", handlers.UpdateEventHandler)
		r.Delete("/events/{slug}", handlers.DeleteEventHandler)
	})

	return r
}
