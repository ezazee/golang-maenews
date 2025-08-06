package handlers

import (
	"encoding/json"
	"maenews/backend/data"
	"maenews/backend/models"
	"maenews/backend/utils"
	"net/http"

	"github.com/go-chi/chi/v5"
)

// --- HANDLER LENGKAP UNTUK CRUD EVENT ---

func CreateEventHandler(w http.ResponseWriter, r *http.Request) {
	var event models.Event
	if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	createdEvent, err := data.CreateEvent(event)
	if err != nil {
		http.Error(w, "Failed to create event", http.StatusInternalServerError)
		return
	}
	utils.RespondWithJSON(w, http.StatusCreated, createdEvent)
}

func UpdateEventHandler(w http.ResponseWriter, r *http.Request) {
	slug := chi.URLParam(r, "slug")
	var event models.Event
	if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	updatedEvent, err := data.UpdateEventBySlug(slug, event)
	if err != nil {
		http.Error(w, "Failed to update event", http.StatusInternalServerError)
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, updatedEvent)
}

func DeleteEventHandler(w http.ResponseWriter, r *http.Request) {
	slug := chi.URLParam(r, "slug")
	if err := data.DeleteEventBySlug(slug); err != nil {
		http.Error(w, "Failed to delete event", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent) // Respons 204 No Content menandakan sukses
}
