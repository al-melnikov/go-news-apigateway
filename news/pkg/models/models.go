package models

import (
	"time"

	"github.com/google/uuid"
)

type Post struct {
	ID        uuid.UUID `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	Link      string    `json:"link"`
}

type Pagination struct {
	CurrentPage int `json:"current_page"`
	PagesNumber int `json:"pages_number"`
	ItemsOnPage int `json:"items_on_page"`
}
