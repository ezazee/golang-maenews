package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Struct untuk Artikel, disesuaikan untuk MongoDB
type Article struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Title       string             `json:"title" bson:"title"`
	Slug        string             `json:"slug" bson:"slug"`
	Description string             `json:"description" bson:"description"`
	Excerpt     string             `json:"excerpt" bson:"excerpt"`
	Category    string             `json:"category" bson:"category"`
	Author      string             `json:"author" bson:"author"`
	PublishedAt time.Time          `json:"publishedAt" bson:"publishedAt"`
	ImageURL    string             `json:"imageUrl" bson:"imageUrl"`
	Tags        []string           `json:"tags" bson:"tags"`
	Featured    bool               `json:"featured" bson:"featured"`
	Views       int                `json:"views" bson:"views"`
	CreatedAt   time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt   time.Time          `json:"updatedAt" bson:"updatedAt"`
}
type TrendingItem struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Title       string             `json:"title" bson:"title"`
	Slug        string             `json:"slug" bson:"slug"`
	Description string             `json:"description" bson:"description"`
	Category    string             `json:"category" bson:"category"`
	PublishedAt time.Time          `json:"publishedAt" bson:"publishedAt"`
}
type Event struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Title       string             `json:"title" bson:"title"`
	Slug        string             `json:"slug" bson:"slug"`
	Location    string             `json:"location" bson:"location"`
	Date        time.Time          `json:"date" bson:"date"`
	Category    string             `json:"category" bson:"category"`
	Description string             `json:"description" bson:"description"`
	ImageURL    string             `json:"imageUrl" bson:"imageUrl"`
}

type User struct {
	ID       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Email    string             `json:"email" bson:"email"`
	Password string             `json:"-" bson:"password"`
}
