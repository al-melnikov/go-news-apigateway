package api

import (
	"api-gateway/pkg/models"

	"github.com/google/uuid"
)

type BadResponse struct {
	Success   bool   `json:"success"`
	Error     string `json:"error"`
	RequestID string `json:"request_id"`
}

type NewsByIDRequest struct {
	ID uuid.UUID `json:"id"`
}

type NewsByIDResponse struct {
	Success   bool             `json:"success"`
	Post      models.Post      `json:"post"`
	Comments  []models.Comment `json:"comments"`
	RequestID string           `json:"request_id"`
}

type NewsDetailedResponse struct {
	Success   bool             `json:"success"`
	Post      models.Post      `json:"post"`
	Comments  []models.Comment `json:"comments"`
	RequestID string           `json:"request_id"`
}

type NewsTreeDetailedResponse struct {
	Success   bool                 `json:"success"`
	Post      models.Post          `json:"post"`
	Comments  []models.CommentTree `json:"comments"`
	RequestID string               `json:"request_id"`
}

type CommentsByIDResponse struct {
	Success   bool             `json:"success"`
	Comments  []models.Comment `json:"comments"`
	RequestID string           `json:"request_id"`
}

type CommentsByIDRequest struct {
	ID uuid.UUID `json:"news_id"`
}

type CommentsTreeByIDRequest struct {
	ID uuid.UUID `json:"news_id"`
}

type CommentsTreeByIDResponse struct {
	Success   bool                 `json:"success"`
	Comments  []models.CommentTree `json:"comments"`
	RequestID string               `json:"request_id"`
}

type NewsRegExpRequest struct {
	Page   int    `json:"page"`
	RegExp string `json:"reg_exp"`
}

type NewsRegExpResponse struct {
	Success    bool              `json:"success"`
	Posts      []models.Post     `json:"posts"`
	Pagination models.Pagination `json:"pagination"`
	RequestID  string            `json:"request_id"`
}

type AddCommentRequest struct {
	NewsID   uuid.UUID     `json:"news_id"`
	ParentID uuid.NullUUID `json:"parent_id"`
	Content  string        `json:"content"`
}

type AddCommentCensorRequest struct {
	Content string `json:"content"`
}

type CensorResponse struct {
	Success    bool   `json:"success"`
	IsCensored bool   `json:"is_censored"`
	RequestID  string `json:"request_id"`
}

type AddCommentResponse struct {
	Success   bool      `json:"success"`
	Error     string    `json:"error"`
	CommentID uuid.UUID `json:"comment_id"`
	RequestID string    `json:"request_id"`
}
