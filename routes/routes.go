package routes

import (
	"github.com/gorilla/mux"
	"go-restful-api/handlers"
)

func SetupRoutes() *mux.Router {
	router := mux.NewRouter()

	// Route list
	router.HandleFunc("/api/books", handlers.GetBooks).Methods("GET")
	router.HandleFunc("/api/books/{id}", handlers.GetBook).Methods("GET")
	router.HandleFunc("/api/books", handlers.CreateBook).Methods("POST")
	router.HandleFunc("/api/books/{id}", handlers.UpdateBook).Methods("PUT")
	router.HandleFunc("/api/books/{id}", handlers.DeleteBook).Methods("DELETE")

	return router
}
