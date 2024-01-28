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
	//
	// // Initialize service
	userService := service.NewUserService(userRepository)
	//
	// // Initialize controller
	userController := controller.NewUserController(userService)

	// Initialize router

	// Define routes
	// router.HandleFunc("/users/{id}", userController.GetUserByID).Methods("GET")
	// router.HandleFunc("/users/{id}", userController.UpdateUser).Methods("PUT")
	// router.HandleFunc("/users/{id}", userController.DeleteUser).Methods("DELETE")

	feedbackRepository := repository.GetFeedbackRepository()

	// Create service instance
	feedbackService := service.NewFeedbackService(feedbackRepository)

	// Create controller instance
	feedbackController := controller.NewFeedbackController(feedbackService)

	router := mux.NewRouter()
    router.HandleFunc("/users", userController.CreateUser).Methods("POST")
	// Define routes using your favorite router (e.g., http.HandleFunc with the default HTTP server or mux)
	router.HandleFunc("/feedbacks", feedbackController.GetFeedbacksHandler).Methods("POST")

	// Start the server
	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
