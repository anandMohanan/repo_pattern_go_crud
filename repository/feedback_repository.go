package repository

import (
	"context"
	"fmt"
	"hello/models"
	"hello/response"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type FeedbackRepository struct {
	collection *mongo.Collection
}

func GetFeedbackRepository() *FeedbackRepository {

	err := godotenv.Load()
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(os.Getenv("MONGO_URI")).SetServerAPIOptions(serverAPI)
	// Create a new client and connect to the server
	client, err := mongo.Connect(context.Background(), opts)
	if err != nil {
		panic(err)
	}
	database := client.Database("golang_test")
	collection := database.Collection("feedback")

	fmt.Println("repo connect")
	return &FeedbackRepository{
		collection: collection,
	}
}

func (fr *FeedbackRepository) GetFeedbacks(request models.FeedbackRequest) ([]response.FeedbackResponse, error) {

	fmt.Println(request)
	filter := bson.M{}

	// Add additional filters based on request parameters
	if request.FeedbackFilterObj.ProjectName != "" {
		filter["project_name"] = request.FeedbackFilterObj.ProjectName
	}
	if request.FeedbackFilterObj.OrganizationName != "" {
		filter["organization_name"] = request.FeedbackFilterObj.OrganizationName
	}
	if request.FeedbackFilterObj.Rating > 0 {
		filter["feedback_rating"] = request.FeedbackFilterObj.Rating
	}
	// Add date range filter if required

	options := options.Find()
	limit := int64(0)
	offset := int64(0)
	options.SetLimit(limit)
	options.SetSkip(offset)

	// Sorting
	sortField := request.FeedbackSort.Column
	sortOrder := 1 // Default to ascending order
	if request.FeedbackSort.Order == "desc" {
		sortOrder = -1
	}
	options.SetSort(bson.D{{Key: sortField, Value: sortOrder}})
	fmt.Println(filter)
	fmt.Println(options)
	// Execute the query
	cursor, err := fr.collection.Find(context.Background(), bson.D{})
	fmt.Println(cursor)
	fmt.Println(err)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())
	fmt.Println(cursor)
	// Decode the results
	var feedbacks []response.FeedbackResponse
	for cursor.Next(context.Background()) {
		var feedback response.FeedbackResponse
		if err := cursor.Decode(&feedback); err != nil {
			return nil, err
		}
		fmt.Println(feedback)
		feedbacks = append(feedbacks, feedback)
	}

	return feedbacks, nil
}
