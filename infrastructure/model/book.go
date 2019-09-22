package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Book presents the book entity
type Book struct {
	gorm.Model
	SID         string    `gorm:"type:varchar(16);not null;column:sid"`
	Title       string    `gorm:"type:varchar(128);not null;column:title"`
	SubTitle    string    `gorm:"type:varchar(128);not null;column:sub_title"`
	ImageURL    string    `gorm:"type:varchar(256);column:image_url"`
	Author      string    `gorm:"type:varchar(128);column:author"`
	Publisher   string    `gorm:"type:varchar(128);column:publisher"`
	Producer    string    `gorm:"type:varchar(128);column:producer"`
	Published   time.Time `gorm:"type:varchar(64);column:published"`
	OriginTitle string    `gorm:"type:varchar(128);column:origin_title"`
	Translator  string    `gorm:"type:varchar(128);column:translator"`
	PageNumber  int       `gorm:"type:integer;column:page_number"`
	Price       float64   `gorm:"type:decimal;column:price"`
	CoverType   string    `gorm:"type:varchar(64);column:cover_type"`
	Series      string    `gorm:"type:varchar(128);column:series"`
	ISBN        string    `gorm:"type:varchar(16);column:isbn"`
	Rate        float64   `gorm:"type:decimal;column:rate"`
	Intro       string    `gorm:"type:varchar(1024);column:intro"`
}

// BookModel presents the book model and serializer
type BookModel struct {
	ID          uint       `json:"id"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
	SID         string     `json:"sid"`
	Title       string     `json:"title"`
	SubTitle    string     `json:"sub_title"`
	ImageURL    string     `json:"image_url"`
	Author      string     `json:"author"`
	Publisher   string     `json:"publisher"`
	Producer    string     `json:"producer"`
	Published   time.Time  `json:"published"`
	OriginTitle string     `json:"origin_title"`
	Translator  string     `json:"translator"`
	PageNumber  int        `json:"page_number"`
	Price       float64    `json:"price"`
	CoverType   string     `json:"cover_type"`
	Series      string     `json:"series"`
	ISBN        string     `json:"isbn"`
	Rate        float64    `json:"rate"`
	Intro       string     `json:"intro"`
}

// Serializer transfers entity to model
func (book *Book) Serializer() BookModel {
	return BookModel{
		ID:          book.ID,
		CreatedAt:   book.CreatedAt,
		UpdatedAt:   book.UpdatedAt,
		DeletedAt:   book.DeletedAt,
		SID:         book.SID,
		Title:       book.Title,
		SubTitle:    book.SubTitle,
		Author:      book.Author,
		ImageURL:    book.ImageURL,
		Publisher:   book.Publisher,
		Producer:    book.Producer,
		Published:   book.Published,
		OriginTitle: book.OriginTitle,
		Translator:  book.Translator,
		PageNumber:  book.PageNumber,
		Price:       book.Price,
		CoverType:   book.CoverType,
		Series:      book.Series,
		ISBN:        book.ISBN,
		Rate:        book.Rate,
		Intro:       book.Intro,
	}
}
