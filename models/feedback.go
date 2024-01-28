package models

import (
	"time"
)

type FeedbackRequest struct {
	UserID              string    `json:"userId"`
	Limit               int64       `json:"limit"`
	Offset              int64       `json:"offset"`
	FeedbackSearchField string    `json:"feedbackSearchField"`
	FeedbackSort        Sort      `json:"feedbackSort"`
	FeedbackFilterObj   FilterObj `json:"feedbackFilterObj"`
}

type Sort struct {
	Column string `json:"column"`
	Order  string `json:"order"`
}

type FilterObj struct {
	ProjectName      string    `json:"projectName"`
	OrganizationName string    `json:"organizationName"`
	Rating           int       `json:"rating"`
	FromDate         time.Time `json:"fromDate"`
	EndDate          time.Time `json:"endDate"`
}

