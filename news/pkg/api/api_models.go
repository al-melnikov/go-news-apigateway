package api

import (
	"news/pkg/models"

	"github.com/google/uuid"
)

type GetPostByIDRequest struct {
	ID uuid.UUID `json:"id"`
}

type GetPostByIDResponse struct {
	Success   bool        `json:"success"`
	Post      models.Post `json:"post"`
	RequestID string      `json:"request_id"`
}

type GetPostsByRegExpRequest struct {
	RegExp      string `json:"reg_exp"`
	CurrentPage int    `json:"page"`
}

type GetPostsByRegExpResponse struct {
	Success    bool              `json:"success"`
	Posts      []models.Post     `json:"posts"`
	Pagination models.Pagination `json:"pagination"`
	RequestID  string            `json:"request_id"`
}

type BadResponse struct {
	Success   bool   `json:"success"`
	Error     string `json:"error"`
	RequestID string `json:"request_id"`
}
