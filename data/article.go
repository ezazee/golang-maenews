package data

import (
	"context"
	"maenews/backend/database"
	"maenews/backend/models"
	"maenews/backend/utils"
	"math"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// PERBAIKAN: Variabel global dihapus. Koleksi akan diambil di dalam setiap fungsi.

// GetAllArticles sekarang mengambil data dari MongoDB
func GetAllArticles(page, limit int) (models.PaginatedArticleResponse, error) {
	articleCollection := database.GetCollection("articles")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// 1. Hitung total dokumen untuk menentukan total halaman
	totalDocuments, err := articleCollection.CountDocuments(ctx, bson.M{})
	if err != nil {
		return models.PaginatedArticleResponse{}, err
	}

	// 2. Hitung total halaman
	totalPages := int(math.Ceil(float64(totalDocuments) / float64(limit)))
	if totalPages == 0 {
		totalPages = 1
	}

	// 3. Siapkan opsi untuk query: skip (lewatkan) dan limit (batasi)
	skip := (page - 1) * limit
	findOptions := options.Find()
	findOptions.SetSkip(int64(skip))
	findOptions.SetLimit(int64(limit))
	findOptions.SetSort(bson.D{{Key: "publishedAt", Value: -1}}) // Urutkan dari yang terbaru

	// 4. Lakukan query untuk mendapatkan artikel di halaman saat ini
	var articles []models.Article
	cursor, err := articleCollection.Find(ctx, bson.M{}, findOptions)
	if err != nil {
		return models.PaginatedArticleResponse{}, err
	}
	defer cursor.Close(ctx)

	if err = cursor.All(ctx, &articles); err != nil {
		return models.PaginatedArticleResponse{}, err
	}

	// 5. Buat struktur respons lengkap
	response := models.PaginatedArticleResponse{
		Data: articles,
		Pagination: models.PaginationData{
			CurrentPage: page,
			TotalPages:  totalPages,
		},
	}

	return response, nil
}

func GetArticleBySlug(slug string) (models.Article, error) {
	articleCollection := database.GetCollection("articles")
	var article models.Article
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Mencari berdasarkan field 'slug'
	err := articleCollection.FindOne(ctx, bson.M{"slug": slug}).Decode(&article)
	if err != nil {
		return models.Article{}, err
	}

	return article, nil
}

// GetArticlesByCategory mengambil data dari MongoDB
func GetArticlesByCategory(categoryName string) ([]models.Article, error) {
	articleCollection := database.GetCollection("articles")
	var articles []models.Article
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"category": bson.M{"$regex": primitive.Regex{Pattern: "^" + categoryName + "$", Options: "i"}}}
	cursor, err := articleCollection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	if err = cursor.All(ctx, &articles); err != nil {
		return nil, err
	}

	return articles, nil
}

// GetArticlesByTag mengambil data dari MongoDB
func GetArticlesByTag(tagName string) ([]models.Article, error) {
	articleCollection := database.GetCollection("articles")
	var articles []models.Article
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"tags": bson.M{"$regex": primitive.Regex{Pattern: "^" + tagName + "$", Options: "i"}}}
	cursor, err := articleCollection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	if err = cursor.All(ctx, &articles); err != nil {
		return nil, err
	}

	return articles, nil
}

func IncrementArticleView(slug string) error {
	articleCollection := database.GetCollection("articles")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"slug": slug}
	update := bson.M{"$inc": bson.M{"views": 1}}

	_, err := articleCollection.UpdateOne(ctx, filter, update)
	return err
}

func SearchArticles(query string) ([]models.Article, error) {
	articleCollection := database.GetCollection("articles")
	var articles []models.Article
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Membuat filter regex untuk pencarian case-insensitive di beberapa field
	filter := bson.M{
		"$or": []bson.M{
			{"title": bson.M{"$regex": primitive.Regex{Pattern: query, Options: "i"}}},
			{"excerpt": bson.M{"$regex": primitive.Regex{Pattern: query, Options: "i"}}},
			{"description": bson.M{"$regex": primitive.Regex{Pattern: query, Options: "i"}}},
			{"tags": bson.M{"$regex": primitive.Regex{Pattern: query, Options: "i"}}},
		},
	}

	cursor, err := articleCollection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	if err = cursor.All(ctx, &articles); err != nil {
		return nil, err
	}

	return articles, nil
}

func CreateArticle(article models.Article) (models.Article, error) {
	articleCollection := database.GetCollection("articles")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	article.ID = primitive.NewObjectID()
	article.Slug = utils.Slugify(article.Title)
	article.CreatedAt = time.Now()
	article.UpdatedAt = time.Now()
	// ... set nilai default lain jika perlu

	_, err := articleCollection.InsertOne(ctx, article)
	return article, err
}

func UpdateArticleBySlug(slug string, article models.Article) (models.Article, error) {
	articleCollection := database.GetCollection("articles")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	update := bson.M{
		"$set": bson.M{
			"title":       article.Title,
			"slug":        utils.Slugify(article.Title),
			"description": article.Description,
			"excerpt":     article.Excerpt,
			"category":    article.Category,
			"author":      article.Author,
			"imageUrl":    article.ImageURL,
			"tags":        article.Tags,
			"featured":    article.Featured,
			"updatedAt":   time.Now(),
		},
	}

	filter := bson.M{"slug": slug}
	_, err := articleCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return models.Article{}, err
	}

	// Ambil kembali data yang sudah diupdate
	var updatedArticle models.Article
	err = articleCollection.FindOne(ctx, bson.M{"slug": utils.Slugify(article.Title)}).Decode(&updatedArticle)
	return updatedArticle, err
}

func DeleteArticleBySlug(slug string) error {
	articleCollection := database.GetCollection("articles")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"slug": slug}
	_, err := articleCollection.DeleteOne(ctx, filter)
	return err
}
