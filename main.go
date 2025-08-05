package main

import (
	"log"
	"net/http"
	"maenews/backend/database"
	"maenews/backend/router"

	"github.com/joho/godotenv"
)

func main() {
	// Muat variabel dari file .env
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// Hubungkan ke database
	database.ConnectDB()
	
	r := router.SetupRouter()

	log.Println("MaeNews backend server is starting on port 8080...")
	
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("Could not start server: %s\n", err)
	}
}