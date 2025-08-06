package main

import (
	"context"
	"fmt"
	"log"
	"maenews/backend/data"
	"maenews/backend/database"
	"maenews/backend/models"
	"math/rand"
	"regexp"
	"strings"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Fungsi helper untuk membuat slug dari judul
var nonAlphanumericRegex = regexp.MustCompile(`[^a-z0-9]+`)

func slugify(s string) string {
	return strings.Trim(nonAlphanumericRegex.ReplaceAllString(strings.ToLower(s), "-"), "-")
}

func main() {
	if err := godotenv.Load("../../.env"); err != nil {
		log.Fatal("Error loading .env file from root directory")
	}
	database.ConnectDB()
	fmt.Println("Seeding database...")
	clearCollections()
	seedAdminUser()
	seedArticles()
	seedTrendingItems()
	seedUpcomingEvents()
	fmt.Println("Database seeding completed successfully!")
}

func clearCollections() {
	fmt.Println("Clearing old data...")
	database.GetCollection("users").DeleteMany(context.Background(), primitive.M{})
	database.GetCollection("articles").DeleteMany(context.Background(), primitive.M{})
	database.GetCollection("trending_items").DeleteMany(context.Background(), primitive.M{})
	database.GetCollection("upcoming_events").DeleteMany(context.Background(), primitive.M{})
}

func seedAdminUser() {
	fmt.Println("Seeding admin user...")
	adminUser := models.User{
		Email:    "admin@maenews.com",
		Password: "passwordyangaman",
	}
	_, err := data.CreateUser(adminUser)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key error") {
			fmt.Println("Admin user already exists.")
		} else {
			log.Fatalf("Failed to seed admin user: %v", err)
		}
	} else {
		fmt.Println("Admin user created successfully: admin@maenews.com")
	}
}

func seedArticles() {
	fmt.Println("Seeding articles...")
	articleCollection := database.GetCollection("articles")
	var articlesToSeed []interface{}

	titles := []string{
		"Attack on Titan Final Season: Pengumuman Tanggal Rilis", "Studio Ghibli Umumkan Proyek Anime Terbaru 2026",
		"VTuber Kobo Kanaeru Pecahkan Rekor Subscribers", "Industri Gaming Indonesia: Perkembangan Game Anime",
		"Cosplay Competition World Championship 2025", "Game 'Genshin Impact' Rilis Update Besar Bertema Nusantara",
		"Review Manga 'Kagurabachi': Hype yang Terbukti?", "Tips & Trik Menjadi Cosplayer Profesional dari Hakken",
		"Developer Game 'Toge Productions' Umumkan Proyek Baru", "Film Live Action 'One Piece' Season 2 Mulai Syuting",
		"Panduan Lengkap Menghadiri Comifuro untuk Pemula", "Kaela Kovalskia dari Hololive ID Menjadi Streamer Teratas",
		"Workshop Membuat Properti Cosplay dari EVA Foam", "Valorant Champions Tour 2025: Tim Indonesia Lolos ke Final",
		"Anime 'Jujutsu Kaisen' Season 3 Dikonfirmasi", "Mengenal Seiyuu di Balik Karakter Populer",
		"Figure Skala 1/7 Gojo Satoru Dirilis, Dompet Menangis!", "Turnamen Mobile Legends Tingkat Asia Tenggara Dimulai",
		"Event Jejepangan Terbesar di Bandung Kembali Hadir", "Review Anime 'Frieren: Beyond Journey's End'",
	}

	categories := []string{"Anime", "Gaming", "Cosplay", "Event", "Content Creator"}
	authors := []string{"Admin", "Redaksi", "Tim Reporter", "Gaming Desk", "Cosplay News"}
	statuses := []string{"Published", "Published", "Published", "Draft"} // Mayoritas Published

	for i := 0; i < 20; i++ {
		title := titles[i%len(titles)]
		article := models.Article{
			Title:       title,
			Slug:        slugify(title),
			Description: fmt.Sprintf("Ini adalah deskripsi lengkap untuk artikel '%s'.", title),
			Excerpt:     fmt.Sprintf("Kutipan singkat dari artikel '%s'...", title),
			Category:    categories[i%len(categories)],
			Author:      authors[i%len(authors)],
			PublishedAt: time.Now().AddDate(0, 0, -i),
			ImageURL:    fmt.Sprintf("https://placehold.co/600x400/1E293B/FFFFFF?text=Artikel+%d", i+1),
			Tags:        []string{categories[i%len(categories)], "Update"},
			Featured:    i == 0,
			Views:       rand.Intn(5000) + 100,
			Status:      statuses[i%len(statuses)],
			CreatedAt:   time.Now().AddDate(0, 0, -i),
			UpdatedAt:   time.Now().AddDate(0, 0, -i),
		}
		articlesToSeed = append(articlesToSeed, article)
	}

	_, err := articleCollection.InsertMany(context.Background(), articlesToSeed)
	if err != nil {
		log.Fatalf("Failed to seed articles: %v", err)
	}
}

func seedTrendingItems() {
	fmt.Println("Seeding trending items (legacy)...")
	trendingCollection := database.GetCollection("trending_items")
	var itemsToSeed []interface{}

	titles := []string{
		"Demon Slayer Season 4 Dikonfirmasi", "Content Creator Terbaru dari Hololive",
	}
	categories := []string{"Anime", "Gaming", "Event", "Content Creator"}

	for i := 1; i <= 5; i++ {
		title := titles[i%len(titles)]
		item := models.TrendingItem{
			Title:       fmt.Sprintf("%s #%d", title, i),
			Description: fmt.Sprintf("Deskripsi singkat untuk topik trending ke-%d.", i),
			Category:    categories[i%len(categories)],
			PublishedAt: time.Now().AddDate(0, 0, -i),
		}
		itemsToSeed = append(itemsToSeed, item)
	}

	_, err := trendingCollection.InsertMany(context.Background(), itemsToSeed)
	if err != nil {
		log.Fatalf("Failed to seed trending items: %v", err)
	}
}

func seedUpcomingEvents() {
	fmt.Println("Seeding upcoming events...")
	eventCollection := database.GetCollection("upcoming_events")
	var eventsToSeed []interface{}

	locations := []string{"JCC, Jakarta", "ICE BSD, Tangerang", "Balai Kartini", "JIExpo Kemayoran", "Trans Studio, Bandung", "Pakuwon Mall, Surabaya"}
	eventNames := []string{"Pop Culture Fest", "Anime Convention", "Gaming Expo", "Cosplay Gathering", "Comic Frontier", "Creators Super Fest"}

	// PERBAIKAN: Loop diubah untuk membuat 30 data event
	for i := 1; i <= 30; i++ {
		title := fmt.Sprintf("%s %d", eventNames[i%len(eventNames)], 2025+i/6)
		event := models.Event{
			Title:       title,
			Slug:        slugify(title),
			Location:    locations[i%len(locations)],
			Date:        time.Now().AddDate(0, 1, i*5), // Event di masa depan dengan jarak 5 hari
			Category:    "Convention",
			Description: fmt.Sprintf("Deskripsi untuk event pop culture ke-%d.", i),
			ImageURL:    fmt.Sprintf("https://placehold.co/400x300/7C2D12/FFFFFF?text=Event+%d", i),
		}
		eventsToSeed = append(eventsToSeed, event)
	}

	_, err := eventCollection.InsertMany(context.Background(), eventsToSeed)
	if err != nil {
		log.Fatalf("Failed to seed upcoming events: %v", err)
	}
}
