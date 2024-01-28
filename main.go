// main.go
package main

import (
	"github.com/gorilla/mux"
	"hello/controller"
	"hello/repository"
	"hello/service"
	"log"
	"net/http"
)

func main() {
	// Initialize repository
	userRepository := repository.NewUserRepository()

	// Initialize service
	userService := service.NewUserService(userRepository)

	// Initialize controller
	userController := controller.NewUserController(userService)

	// Initialize router
	router := mux.NewRouter()

	// Define routes
	router.HandleFunc("/users", userController.CreateUser).Methods("POST")
	// router.HandleFunc("/users/{id}", userController.GetUserByID).Methods("GET")
	// router.HandleFunc("/users/{id}", userController.UpdateUser).Methods("PUT")
	// router.HandleFunc("/users/{id}", userController.DeleteUser).Methods("DELETE")

	// Start the server
	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

