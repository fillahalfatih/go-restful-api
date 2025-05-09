package main

import (
	"encoding/json"
	"net/http"
	"log"
	"github.com/gorilla/mux"
)

// Book Struct (Model)
type Book struct {
	Id   	string	`json:"id"`
	Isbn	string	`json:"isbn"`
	Title 	string	`json:"title"`
	Year	int 	`json:"year"`
	Author  *Author `json:"author"`
}

// Author Struct (Model)
type Author struct {
	FirstName	string	`json:"first_name"`
	LastName	string	`json:"last_name"`
}

// Init books var as a slice Book struct
var books []Book

// Get All Books
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

// Get Single Book
func getBook(w http.ResponseWriter, r *http.Request) {

}

// Create a New Book
func createBook(w http.ResponseWriter, r *http.Request) {

}

// Update a Book
func updateBook(w http.ResponseWriter, r *http.Request) {

}

// Delete a Book
func deleteBook(w http.ResponseWriter, r *http.Request) {

}

func main() {
	// Initialize the router
	r := mux.NewRouter()

	// Mock Data - @todo: implement DB
	books = append(books, Book{Id: "1", Isbn: "448-1234567890", Title: "Book One", Year: 2021, Author: &Author{FirstName: "John", LastName: "Doe"}})
	books = append(books, Book{Id: "2", Isbn: "448-1234567891", Title: "Book Two", Year: 2022, Author: &Author{FirstName: "Jane", LastName: "Smith"}})
	books = append(books, Book{Id: "3", Isbn: "448-1234567892", Title: "Book Three", Year: 2023, Author: &Author{FirstName: "Bob", LastName: "Johnson"}})
	books = append(books, Book{Id: "4", Isbn: "448-1234567893", Title: "Book Four", Year: 2024, Author: &Author{FirstName: "Alice", LastName: "Brown"}})
	books = append(books, Book{Id: "5", Isbn: "448-1234567894", Title: "Book Five", Year: 2025, Author: &Author{FirstName: "Charlie", LastName: "Davis"}})

	// Route handlers / Endpoints
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", r))
}