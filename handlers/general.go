package handlers

import (
	"maenews/backend/data"
	"net/http"

	"github.com/go-chi/chi/v5"
)

// PERBAIKAN: GetTrendingItemsHandler sekarang menangani error dari database
func GetTrendingItemsHandler(w http.ResponseWriter, r *http.Request) {
	items, err := data.GetTrendingItems()
	if err != nil {
		http.Error(w, "Failed to fetch trending items", http.StatusInternalServerError)
		return
	}
	respondWithJSON(w, http.StatusOK, items)
}

// PERBAIKAN: GetUpcomingEventsHandler sekarang menangani error dari database
func GetUpcomingEventsHandler(w http.ResponseWriter, r *http.Request) {
	events, err := data.GetUpcomingEvents()
	if err != nil {
		http.Error(w, "Failed to fetch upcoming events", http.StatusInternalServerError)
		return
	}
	respondWithJSON(w, http.StatusOK, events)
}

func GetEventBySlugHandler(w http.ResponseWriter, r *http.Request) {
	slug := chi.URLParam(r, "slug")
	event, err := data.GetEventBySlug(slug)
	if err != nil {
		http.Error(w, "Event not found", http.StatusNotFound)
		return
	}
	respondWithJSON(w, http.StatusOK, event)
}
