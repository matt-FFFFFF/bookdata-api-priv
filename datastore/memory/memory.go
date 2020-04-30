package memory

import (
	"os"

	"github.com/matt-FFFFFF/bookstore"
)

// A BookService is an in-memory store for the bookstore & implements the bookstore.BookService interface
type BookService struct {
	Store   []*bookstore.Book
	CSVPath string
}

// Init loads the book data from disk into memory an returns a pointer to the store
func Init(csvPath string) (bookstore.BookService, error) {
	var b BookService
	csvpath := "./books.csv"
	file, err := os.Open(csvpath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	b.Store, err = LoadData(file)
	return &b, err
}

// Books is part of the BookSevice interface and returns all the books in the store
func (b *BookService) Books() ([]*bookstore.Book, error) {
	return b.Store, nil
}
