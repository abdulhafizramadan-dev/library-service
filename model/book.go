package model

import (
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title         string  `json:"title"`
	Price         float64 `json:"price"`
	PublishedDate string  `json:"publishedDate"`
}

type BookResponse struct {
	Id            uint    `json:"id"`
	Title         string  `json:"title"`
	Price         float64 `json:"price"`
	PublishedDate string  `json:"publishedDate"`
}

type UpdateBookRequest struct {
	Title string `json:"title"`
}

func (b Book) MapToResponse() BookResponse {
	return BookResponse{
		Id:            b.ID,
		Title:         b.Title,
		Price:         b.Price,
		PublishedDate: b.PublishedDate,
	}
}
