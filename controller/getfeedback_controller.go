package controller

import (
	"encoding/json"
	"fmt"
	"hello/models"
	"hello/service"
	"io"
	"net/http"
)

type FeedbackController struct {
	feedbackService *service.FeedbackService
}

func NewFeedbackController(feedbackService *service.FeedbackService) *FeedbackController {

	fmt.Println("get feeback conollernew")
	return &FeedbackController{
		feedbackService: feedbackService,
	}
}

func (fc *FeedbackController) GetFeedbacksHandler(w http.ResponseWriter, r *http.Request) {
	var request models.FeedbackRequest
	fmt.Println(r.Body)
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil && err != io.EOF {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println("request", request)
	feedbacks, err := fc.feedbackService.GetFeedbacks(request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Encode feedbacks array as JSON and send as HTTP response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(feedbacks)
}
