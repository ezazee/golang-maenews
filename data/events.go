package data

import (
	"context"
	"maenews/backend/database"
	"maenews/backend/models"
	"maenews/backend/utils"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateEvent(event models.Event) (models.Event, error) {
	eventCollection := database.GetCollection("upcoming_events")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	event.ID = primitive.NewObjectID()
	event.Slug = utils.Slugify(event.Title)

	_, err := eventCollection.InsertOne(ctx, event)
	return event, err
}
