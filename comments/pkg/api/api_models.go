package api

import (
	"comments/pkg/models"

	"github.com/google/uuid"
)

type PostCommentResponse struct {
	Success   bool      `json:"success"`
	Error     string    `json:"error"`
	CommentID uuid.UUID `json:"comment_id"`
	RequestID string    `json:"request_id"`
}

type CreateCommentRequest struct {
	NewsID   uuid.UUID     `json:"news_id"`
	ParentID uuid.NullUUID `json:"parent_id"`
	Content  string        `json:"content"`
}

type GetCommentsRequest struct {
	NewsID uuid.UUID `json:"news_id"`
}

type GetCommentsResponse struct {
	Success   bool             `json:"success"`
	Error     string           `json:"error"`
	Comments  []models.Comment `json:"comments"`
	RequestID string           `json:"request_id"`
}

type GetCommentsTreeRequest struct {
	NewsID uuid.UUID `json:"news_id"`
}

type GetCommentsTreeResponse struct {
	Success   bool                 `json:"success"`
	Error     string               `json:"error"`
	Comments  []models.CommentTree `json:"comments"`
	RequestID string               `json:"request_id"`
}
