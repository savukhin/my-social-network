package dto

import (
	"time"
)

type AttachementType string

const (
	PostType    AttachementType = "post"
	CommentType AttachementType = "comment"
)

type PhotoAttachement struct {
	ID        int       `json:"id"`
	AuthorID  int       `json:"author_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	URL       string    `json:"url"`
}

type Attachement struct {
	ID        int             `json:"id"`
	Type      AttachementType `json:"contentType"`
	UserID    int             `json:"user_id"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
}
