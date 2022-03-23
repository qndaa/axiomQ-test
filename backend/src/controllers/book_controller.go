package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/qndaa/backend/src/model"
	"github.com/qndaa/backend/src/services"
	"net/http"
	"time"
)

type BookController struct {
	bookService *services.BookService
}

var BookControllerInstance BookController

func init() {
	BookControllerInstance = BookController{
		bookService: &services.BookInstanceService,
	}
}

func (c *BookController) FetchAll(w http.ResponseWriter, r *http.Request) {
	c.bookService.FetchAll(w)
}

func (c *BookController) FetchOne(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	c.bookService.FetchOne(w, vars["id"])

}

func (c *BookController) Create(w http.ResponseWriter, r *http.Request) {
	var book model.Book
	err := json.NewDecoder(r.Body).Decode(&book)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	book.IssuanceDate = time.Time(book.DateDTO)
	c.bookService.Create(w, &book)
}
