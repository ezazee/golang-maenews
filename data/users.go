package data

import (
	"context"
	"maenews/backend/database"
	"maenews/backend/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

// CreateUser membuat pengguna baru dengan password yang di-hash
func CreateUser(user models.User) (models.User, error) {
	userCollection := database.GetCollection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return models.User{}, err
	}
	user.Password = string(hashedPassword)
	user.ID = primitive.NewObjectID()

	_, err = userCollection.InsertOne(ctx, user)
	if err != nil {
		return models.User{}, err
	}
	user.Password = "" // Hapus password sebelum dikembalikan
	return user, nil
}

// GetUserByUsername mengambil pengguna berdasarkan username
func GetUserByUsername(username string) (models.User, error) {
	userCollection := database.GetCollection("users")
	var user models.User
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := userCollection.FindOne(ctx, bson.M{"username": username}).Decode(&user)
	return user, err
}
