package services

import (
	"github.com/qndaa/backend/src/model"
	"github.com/qndaa/backend/src/repostitories"
	"github.com/qndaa/backend/src/utils"
	"net/http"
	"strconv"
)

type BookService struct {
	bookRepository *repostitories.BookRepository
}

var BookInstanceService BookService

func init() {
	BookInstanceService = BookService{
		bookRepository: &repostitories.BookInstanceRepository,
	}
}

func (s *BookService) FetchAll(w http.ResponseWriter) {
	books := s.bookRepository.FetchAll()
	utils.RenderJSON(w, books)
	w.WriteHeader(200)
}

func (s *BookService) FetchOne(w http.ResponseWriter, paramId string) {
	id, err := strconv.Atoi(paramId)

	if err != nil {
		utils.BadRequest(w, "Invalid ID!")
		return
	}

	book, err := s.bookRepository.FetchOne(id)
	if err != nil {
		utils.BadRequest(w, "Invalid ID!")
		return
	}

	if book.ID == 0 {
		utils.BadRequest(w, "Invalid ID!")
		return
	}

	utils.RenderJSON(w, book)
	w.WriteHeader(200)
}

func (s *BookService) Create(w http.ResponseWriter, b *model.Book) {
	book, err := s.bookRepository.Create(b)

	if err != nil {
		utils.BadRequest(w, "Invalid data!")
		return
	}

	w.WriteHeader(201)
	utils.RenderJSON(w, book)
}
