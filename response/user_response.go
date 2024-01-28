package response

import "hello/models"

type CreateUserResponse struct {
	UserID string `json:"userId,omitempty"`
	Err    error  `json:"error,omitempty"`
}

type UpdateUserResponse struct {
	ModifiedCount int64 `json:"modifiedCount,omitempty"`
	Err           error `json:"error,omitempty"`
}

type DeleteUserResponse struct {
	DeletedCount int64 `json:"deletedCount,omitempty"`
	Err          error `json:"error,omitempty"`
}

type GetUserResponse struct {
	User models.User `json:"user,omitempty"`
	Err  error       `json:"error,omitempty"`
}

