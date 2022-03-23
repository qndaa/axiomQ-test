package db

import (
	"github.com/qndaa/backend/src/model"
	"gorm.io/gorm"
	"time"
)

var Books = []model.Book{
	{
		Model:      gorm.Model{},
		Title:      "Na Drini cuprija",
		AuthorName: "Ivo Andric",
		IssuanceDate: time.Date(2020, time.April,
			34, 0, 0, 0, 0, time.UTC),
		Genre: "Roman",
	}, {
		Model:      gorm.Model{},
		Title:      "Tvrdjava",
		AuthorName: "Mesa Selimovic",
		IssuanceDate: time.Date(2020, time.April,
			34, 0, 0, 0, 0, time.UTC),

		Genre: "Roman",
	},
}

var Users = []model.User{
	{
		Model:    gorm.Model{},
		Username: "marko",
		Password: "marko",
	}, {
		Model:    gorm.Model{},
		Username: "janko",
		Password: "janko",
	},
}
