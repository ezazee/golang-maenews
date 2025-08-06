package data

import (
	"context"
	"maenews/backend/database"
	"maenews/backend/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// PERBAIKAN: Variabel global dihapus. Koleksi akan diambil di dalam setiap fungsi.

// GetTrendingItems mengambil data dari MongoDB
func GetTrendingItems() ([]models.TrendingItem, error) {
	articleCollection := database.GetCollection("articles")
	var trendingItems []models.TrendingItem
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Filter: artikel yang dipublikasikan dalam 30 hari terakhir
	thirtyDaysAgo := time.Now().AddDate(0, 0, -30)
	filter := bson.M{"publishedAt": bson.M{"$gte": thirtyDaysAgo}}

	// Opsi: urutkan berdasarkan 'views' menurun, dan batasi 5 hasil
	opts := options.Find().SetSort(bson.D{{Key: "views", Value: -1}}).SetLimit(5)

	cursor, err := articleCollection.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	// Kita hanya perlu beberapa field untuk trending, jadi kita decode ke struct TrendingItem
	if err = cursor.All(ctx, &trendingItems); err != nil {
		return nil, err
	}

	return trendingItems, nil
}

// GetUpcomingEvents mengambil data dari MongoDB
func GetUpcomingEvents() ([]models.Event, error) {
	eventCollection := database.GetCollection("upcoming_events")
	var events []models.Event
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	today := time.Now()
	startOfDay := time.Date(today.Year(), today.Month(), today.Day(), 0, 0, 0, 0, today.Location())
	filter := bson.M{"date": bson.M{"$gte": startOfDay}}

	opts := options.Find().SetSort(bson.D{
		{Key: "date", Value: 1},
		{Key: "_id", Value: 1},
	})

	cursor, err := eventCollection.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	if err = cursor.All(ctx, &events); err != nil {
		return nil, err
	}

	return events, nil
}

func GetEventBySlug(slug string) (models.Event, error) {
	eventCollection := database.GetCollection("upcoming_events")
	var event models.Event
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := eventCollection.FindOne(ctx, bson.M{"slug": slug}).Decode(&event)
	if err != nil {
		return models.Event{}, err
	}

	return event, nil
}
