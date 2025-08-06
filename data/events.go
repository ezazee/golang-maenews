package data

import (
	"context"
	"maenews/backend/database"
	"maenews/backend/models"
	"maenews/backend/utils"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// --- FUNGSI CRUD LENGKAP UNTUK EVENT ---

// CreateEvent membuat event baru di database
func CreateEvent(event models.Event) (models.Event, error) {
	eventCollection := database.GetCollection("upcoming_events")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	event.ID = primitive.NewObjectID()
	event.Slug = utils.Slugify(event.Title)

	_, err := eventCollection.InsertOne(ctx, event)
	return event, err
}

// UpdateEventBySlug memperbarui event yang ada berdasarkan slug-nya
func UpdateEventBySlug(slug string, event models.Event) (models.Event, error) {
	eventCollection := database.GetCollection("upcoming_events")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	newSlug := utils.Slugify(event.Title)
	update := bson.M{
		"$set": bson.M{
			"title":       event.Title,
			"slug":        newSlug,
			"location":    event.Location,
			"date":        event.Date,
			"category":    event.Category,
			"description": event.Description,
			"imageUrl":    event.ImageURL,
		},
	}

	filter := bson.M{"slug": slug}
	_, err := eventCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return models.Event{}, err
	}

	// Ambil kembali data yang sudah diupdate untuk dikembalikan
	var updatedEvent models.Event
	err = eventCollection.FindOne(ctx, bson.M{"slug": newSlug}).Decode(&updatedEvent)
	return updatedEvent, err
}

// DeleteEventBySlug menghapus event dari database berdasarkan slug-nya
func DeleteEventBySlug(slug string) error {
	eventCollection := database.GetCollection("upcoming_events")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"slug": slug}
	_, err := eventCollection.DeleteOne(ctx, filter)
	return err
}
