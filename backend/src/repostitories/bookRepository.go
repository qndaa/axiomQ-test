package repostitories

import (
	"fmt"
	"github.com/qndaa/backend/src/db"
	"github.com/qndaa/backend/src/model"
)

type BookRepository struct {
}

var BookInstanceRepository BookRepository

func init() {
	BookInstanceRepository = BookRepository{}
}

func (r *BookRepository) FetchAll() []model.Book {
	var books []model.Book
	db.DB.Find(&books)
	return books
}

func (r *BookRepository) FetchOne(id int) (model.Book, error) {
	var book model.Book
	response := db.DB.Find(&book, "ID = ?", id)
	fmt.Println(book)
	return book, response.Error
}

func (r *BookRepository) Create(b *model.Book) (*model.Book, error) {
	response := db.DB.Create(b)
	return b, response.Error
}
