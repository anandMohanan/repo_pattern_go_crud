package response

import (
    "time"
)

type FeedbackResponse struct {
    ID                string    `json:"id"`
    ProjectID         string    `json:"project_id"`
    ProjectName       string    `json:"project_name"`
    UserID            string    `json:"user_id"`
    OrganizationName  string    `json:"organization_name"`
    FeedbackRating    int       `json:"feedback_rating"`
    FeedbackComment   string    `json:"feedback_comment"`
    CreatedBy         string    `json:"created_by"`
    CreatedOn         time.Time `json:"created_on"`
    LastModifiedBy    string    `json:"last_modified_by"`
    LastModifiedOn    time.Time `json:"last_modified_on"`
    IsActive          int       `json:"is_active"`
}

