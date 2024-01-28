// service/user_service.go
package service

import (
	"fmt"
	"hello/models"
	"hello/repository"
	"hello/response"

	"github.com/google/uuid"
)

type UserService struct {
	userRepository *repository.UserRepository
}

func NewUserService(userRepository *repository.UserRepository) *UserService {
	return &UserService{userRepository: userRepository}
}

func (us *UserService) CreateUser(user *models.User) response.CreateUserResponse {
	// user.ID = 447567576
	user.ID = uuid.New().String()
	fmt.Println("create user service")
	return us.userRepository.CreateUser(user)
}

// func (us *UserService) UpdateUser(user *models.User) error {
// 	// Implement business logic for updating a user
// }
//
// func (us *UserService) DeleteUser(userID int) error {
// 	// Implement business logic for deleting a user
// }
//
// func (us *UserService) GetUserByID(userID int) (*models.User, error) {
// 	// Implement business logic for getting a user by ID
// }
