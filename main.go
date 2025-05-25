package main

import (
	"log"
	"net/http"
	"go-restful-api/routes"
	"go-restful-api/config"
	"go-restful-api/models"
)

func main() {
	// Mock data
	// handlers.Books = append(handlers.Books, models.Book{
	// 	Id:    "1",
	// 	Isbn:  "123-456",
	// 	Title: "Book One",
	// 	Year:  2021,
	// 	Author: &models.Author{
	// 		FirstName: "John",
	// 		LastName:  "Doe",
	// 	},
	// })

	config.ConnectDB()

	// Auto-migrate table Book
	config.DB.AutoMigrate(&models.Book{}, &models.Author{})

	// Setup server
	router := routes.SetupRoutes()
	log.Println("Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
