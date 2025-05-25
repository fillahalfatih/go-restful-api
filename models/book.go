package models

type Book struct {
	Id       string  `json:"id" gorm:"primaryKey"`
	Isbn     string  `json:"isbn"`
	Title    string  `json:"title"`
	Year     int     `json:"year"`
	AuthorID uint    `json:"author_id"` // foreign key ke Author
	Author   Author  `json:"author" gorm:"foreignKey:AuthorID"`
}

type Author struct {
	Id        uint   `json:"id" gorm:"primaryKey"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}