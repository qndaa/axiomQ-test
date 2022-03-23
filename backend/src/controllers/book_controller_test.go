package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/qndaa/backend/src/db"
	"github.com/qndaa/backend/src/model"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestFetchOneBook(t *testing.T) {
	tt := []struct {
		name   string
		code   int
		method string
		want   string
		id     int
	}{
		{
			name:   "Get entity",
			code:   200,
			method: http.MethodGet,
			want:   "Marko",
			id:     1,
		}, {
			name:   "No entity",
			code:   400,
			method: http.MethodGet,
			want:   "",
			id:     0,
		},
	}

	db.Connect()
	db.AutoMigration()
	a := App{}
	a.Init()

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			r := httptest.NewRequest(tc.method, fmt.Sprintf("/api/book/%v", tc.id), nil)
			rr := httptest.NewRecorder()
			a.GetRouter().ServeHTTP(rr, r)
			if rr.Code != tc.code {
				t.Errorf("Want status '%d', got '%d'", tc.code, rr.Code)
			}
		})
	}
}

func TestFetchAllBooks(t *testing.T) {
	tt := struct {
		name   string
		code   int
		method string
		want   int
	}{
		name:   "Get Two entities",
		code:   200,
		method: http.MethodGet,
		want:   2,
	}

	db.Connect()
	db.AutoMigration()
	a := App{}
	a.Init()

	t.Run(tt.name, func(t *testing.T) {
		r := httptest.NewRequest(tt.method, fmt.Sprintf("/api/book"), nil)
		rr := httptest.NewRecorder()
		a.GetRouter().ServeHTTP(rr, r)
		if rr.Code != tt.code {
			t.Errorf("Want status '%d', got '%d'", tt.code, rr.Code)
		}
		fmt.Println(rr.Body)
		var books []model.Book
		err := json.NewDecoder(rr.Body).Decode(&books)

		if err != nil {
			t.Errorf("Problems with response body! %v", err)
		}

		if len(books) != tt.want {
			t.Errorf("Want length '%d', got '%d'", tt.want, len(books))
		}
	})

}

func TestCreateBook(t *testing.T) {

	book := model.Book{
		Title:        "Nova knjiga",
		AuthorName:   "Novi autor",
		IssuanceDate: time.Time{},
		DateDTO:      model.DateDTO{},
		Genre:        "Novi zanr!",
	}

	tt := struct {
		name   string
		code   int
		method string
		want   uint
		body   *model.Book
	}{
		name:   "Create Entity!",
		code:   201,
		method: http.MethodPost,
		want:   3,
		body:   &book,
	}

	db.Connect()
	db.AutoMigration()
	a := App{}
	a.Init()

	t.Run(tt.name, func(t *testing.T) {
		b, _ := json.Marshal(tt.body)
		r := httptest.NewRequest(tt.method, fmt.Sprintf("/api/book"), bytes.NewBuffer(b))
		rr := httptest.NewRecorder()
		a.GetRouter().ServeHTTP(rr, r)
		if rr.Code != tt.code {
			t.Errorf("Want status '%d', got '%d'", tt.code, rr.Code)
		}
		fmt.Println(rr.Body)
		var book model.Book
		err := json.NewDecoder(rr.Body).Decode(&book)

		if err != nil {
			t.Errorf("Problems with response body! %v", err)
		}

		if book.ID != tt.want {
			t.Errorf("Want ID '%d', got '%d'", tt.want, book.ID)
		}
	})
}
