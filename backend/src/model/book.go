package model

import (
	"gorm.io/gorm"
	"time"
)

type Book struct {
	gorm.Model
	Title        string    `json:"title"`
	AuthorName   string    `json:"authorName"`
	IssuanceDate time.Time `json:"issuanceDate"`
	DateDTO      DateDTO   `json:"date_dto" gorm:"embedded"`
	Genre        string    `json:"genre"`
}
