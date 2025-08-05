package handlers

import (
	"encoding/json"
	"maenews/backend/data"
	"maenews/backend/models"
	"maenews/backend/utils"
	"net/http"
)

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
