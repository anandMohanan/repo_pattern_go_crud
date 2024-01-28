// controller/user_controller.go
package controller

import (
	"encoding/json"
	"fmt"
	"hello/models"
	"hello/service"
	"net/http"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController(userService *service.UserService) *UserController {
	return &UserController{userService: userService}
}

func (uc *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	// Parse request body
	fmt.Println("create user controller")
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Call service to create user
	data := uc.userService.CreateUser(&user)
	if data.Err != nil {
		fmt.Println("data err", data.Err)
		http.Error(w, data.Err.Error(), http.StatusBadRequest)

		return
	}
	fmt.Println("data", data.UserID)
	// Respond with success message
	jsonResponse := map[string]string{"userID": data.UserID}
	jsonResponseBytes, err := json.Marshal(jsonResponse)
	w.WriteHeader(http.StatusCreated)
	w.Write(jsonResponseBytes)
}

// Implement other controller methods for user CRUD operations
