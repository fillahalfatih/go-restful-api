package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/google/uuid"

	"go-restful-api/models"
	"go-restful-api/config"
)

// Get All Books
func GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var books []models.Book
	if err := config.DB.Preload("Author").Find(&books).Error; err != nil {
		http.Error(w, "Failed to retrieve books", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(books)
}

// Get Single Book
func GetBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var book models.Book
	if err := config.DB.Preload("Author").First(&book, "id = ?", params["id"]).Error; err != nil {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(book)
}

// Create a New Book
func CreateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book models.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	book.Id = uuid.New().String()
	if err := config.DB.Create(&book).Error; err != nil {
		http.Error(w, "Failed to create book", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(book)
}

// Update a Book
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var book models.Book

	if err := config.DB.First(&book, "id = ?", params["id"]).Error; err != nil {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// GORM akan update berdasarkan primary key yang ditemukan tadi
	if err := config.DB.Save(&book).Error; err != nil {
		http.Error(w, "Failed to update book", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(book)
}

// Delete a Book
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var book models.Book

	if err := config.DB.First(&book, "id = ?", params["id"]).Error; err != nil {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	if err := config.DB.Delete(&book).Error; err != nil {
		http.Error(w, "Failed to delete book", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "Book deleted"})
}
