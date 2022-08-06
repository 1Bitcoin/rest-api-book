package domain

import (
	"errors"
	"time"
)

var (
	ErrBookNotFound = errors.New("book not found")
)

type Book struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Author      string    `json:"author"`
	PublishDate time.Time `json:"publish_date"`
	Rating      int       `json:"rating"`
}

type UpdateBookInput struct {
	Title       *string    `json:"title"`
	Author      *string    `json:"author"`
	PublishDate *time.Time `json:"publish_date"`
	Rating      *int       `json:"rating"`
}
