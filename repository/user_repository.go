package repository

import (
	"context"
	"fmt"
	"hello/models"
	"hello/validation"
	"os"
	"time"

	"hello/response"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserRepository struct {
	collection *mongo.Collection
}

func NewUserRepository() *UserRepository {
	err := godotenv.Load()
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(os.Getenv("MONGO_URI")).SetServerAPIOptions(serverAPI)
	// Create a new client and connect to the server
	client, err := mongo.Connect(context.Background(), opts)
	if err != nil {
		panic(err)
	}
	database := client.Database("golang_test")
	collection := database.Collection("users")
	return &UserRepository{
		collection: collection,
	}

}

func (ur *UserRepository) CreateUser(user *models.User) response.CreateUserResponse {
	if err := validation.CreateUserValidate(user, ur.collection); err != nil {
		return response.CreateUserResponse{
			UserID: "", // Empty UserID indicates failure
			Err:    err,
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := ur.collection.InsertOne(ctx, user)
	if err != nil {
		fmt.Println("error in create user repo", err)
	}
	return response.CreateUserResponse{
		UserID: user.ID,
		Err:    err,
	}
}

// func (ur *UserRepository) UpdateUser(user *models.User) error {
// }
//
// func (ur *UserRepository) DeleteUser(userID int) error {
// }
//
// func (ur *UserRepository) GetUserByID(userID int) (*models.User, error) {
// }
//
