package validation

import (
	"context"
	"errors"
	"hello/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateUserValidate(user *models.User, collection *mongo.Collection) error {
	if user.Email == "" {
		return errors.New("email cannot be empty")
	}

	existingUser := models.User{}
	err := collection.FindOne(context.Background(), bson.M{"email": user.Email}).Decode(&existingUser)
	if err == nil {
		return errors.New("email already exists")
	} else if err != mongo.ErrNoDocuments {
		return err
	}
	return nil
}
