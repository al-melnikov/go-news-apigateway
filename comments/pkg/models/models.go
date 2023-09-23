package models

import (
	"time"

	"github.com/google/uuid"
)

type Comment struct {
	ID        uuid.UUID     `json:"id"`
	NewsID    uuid.UUID     `json:"news_id"`
	ParentID  uuid.NullUUID `json:"parent_id"`
	CreatedAt time.Time     `json:"created_at"`
	Content   string        `json:"content"`
}

type CommentTree struct {
	ID               uuid.UUID     `json:"id"`
	NewsID           uuid.UUID     `json:"news_id"`
	CreatedAt        time.Time     `json:"created_at"`
	Content          string        `json:"content"`
	ThreadedComments []CommentTree `json:"thread"`
}
