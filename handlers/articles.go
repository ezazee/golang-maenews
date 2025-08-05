package handlers

import (
	"encoding/json"
	"maenews/backend/data"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
)

// Helper function untuk mengirim respons JSON
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// GetAllArticlesHandler sekarang menangani error dari database
func GetAllArticlesHandler(w http.ResponseWriter, r *http.Request) {
	articles, err := data.GetAllArticles()
	if err != nil {
		http.Error(w, "Failed to fetch articles", http.StatusInternalServerError)
		return
	}
	respondWithJSON(w, http.StatusOK, articles)
}

// GetArticleByIDHandler sekarang menangani error dari database
func GetArticleBySlugHandler(w http.ResponseWriter, r *http.Request) {
	slug := chi.URLParam(r, "slug") // Mengambil 'slug' dari URL
	article, err := data.GetArticleBySlug(slug)
	if err != nil {
		http.Error(w, "Article not found", http.StatusNotFound)
		return
	}
	respondWithJSON(w, http.StatusOK, article)
}

// PERBAIKAN: Menambahkan kembali handler GetArticlesByCategoryHandler
func GetArticlesByCategoryHandler(w http.ResponseWriter, r *http.Request) {
	categoryName := chi.URLParam(r, "categoryName")
	categoryName = strings.ReplaceAll(categoryName, "-", " ")
	articles, err := data.GetArticlesByCategory(categoryName)
	if err != nil {
		http.Error(w, "Failed to fetch articles by category", http.StatusInternalServerError)
		return
	}
	respondWithJSON(w, http.StatusOK, articles)
}

// PERBAIKAN: Menambahkan kembali handler GetArticlesByTagHandler
func GetArticlesByTagHandler(w http.ResponseWriter, r *http.Request) {
	tagName := chi.URLParam(r, "tagName")
	tagName = strings.ReplaceAll(tagName, "-", " ")
	articles, err := data.GetArticlesByTag(tagName)
	if err != nil {
		http.Error(w, "Failed to fetch articles by tag", http.StatusInternalServerError)
		return
	}
	respondWithJSON(w, http.StatusOK, articles)
}

func IncrementViewHandler(w http.ResponseWriter, r *http.Request) {
	slug := chi.URLParam(r, "slug")
	err := data.IncrementArticleView(slug)
	if err != nil {
		// Kita tidak perlu mengirim error ke client, cukup log di server
		// karena ini adalah operasi latar belakang.
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func SearchArticlesHandler(w http.ResponseWriter, r *http.Request) {
	query := chi.URLParam(r, "query")
	articles, err := data.SearchArticles(query)
	if err != nil {
		http.Error(w, "Failed to search articles", http.StatusInternalServerError)
		return
	}
	respondWithJSON(w, http.StatusOK, articles)
}
