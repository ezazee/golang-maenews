package database

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Client adalah instance dari koneksi database MongoDB
var Client *mongo.Client
// PERBAIKAN: Nama database didefinisikan langsung di sini
const dbName = "maenewsDB"

// ConnectDB menghubungkan aplikasi ke MongoDB
func ConnectDB() {
	uri := os.Getenv("MONGO_URI")
	if uri == "" {
		log.Fatal("You must set your 'MONGO_URI' environment variable.")
	}

	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// Ping database untuk memastikan koneksi berhasil
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	Client = client
	log.Println("Successfully connected to MongoDB!")
}

// GetCollection mengembalikan sebuah collection dari database
func GetCollection(collectionName string) *mongo.Collection {
	// PERBAIKAN: Menggunakan konstanta dbName yang sudah didefinisikan
	return Client.Database(dbName).Collection(collectionName)
}