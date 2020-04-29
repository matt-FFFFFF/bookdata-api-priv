package bookstore

import "github.com/google/uuid"

// A Book defines the data record for a single book.
type Book struct {
	BookID        uuid.UUID `json:"book_id"`
	Title         string    `json:"title"`
	Authors       []string  `json:"authors"`
	AverageRating float64   `json:"average_rating"`
	ISBN          string    `json:"isbn"`
	ISBN13        string    `json:"isbn_13"`
	LanguageCode  string    `json:"language_code"`
	NumPages      int       `json:"num_pages"`
	Ratings       int       `json:"ratings"`
	Reviews       int       `json:"reviews"`
}

// BookService describes the methods required to process books
type BookService interface {
	BookById(u uuid.UUID) (*Book, error)
	BookByIsbn(i string) (*Book, error)
	Books() ([]*Book, error)
	BooksByAuthor(a string) ([]*Book, error)
	BooksByTitle(t string) ([]*Book, error)
	NewBook(b *Book) (uuid.UUID, error)
	DeleteBookById(u uuid.UUID) error
	DeleteBookByIsbn(i string) error
}
