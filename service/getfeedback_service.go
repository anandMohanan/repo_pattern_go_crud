package service

import (
	"fmt"
	"hello/models"
	"hello/repository"
	"hello/response"
)

type FeedbackService struct {
	feedbackRepository *repository.FeedbackRepository
}

func NewFeedbackService(feedbackRepository *repository.FeedbackRepository) *FeedbackService {

fmt.Println("new feedback servic")
    return &FeedbackService{
		feedbackRepository: feedbackRepository,
	}
}

func (fs *FeedbackService) GetFeedbacks(request models.FeedbackRequest) ([]response.FeedbackResponse, error) {
	// You can add validation or additional logic here
	fmt.Println("feedback service")
	return fs.feedbackRepository.GetFeedbacks(request)
}
