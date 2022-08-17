package dto

import (
	"time"
)

type AttachementType string

const (
	Post    AttachementType = "post"
	Comment AttachementType = "comment"
)

type Attachement struct {
	ID        int             `json:"id"`
	Type      AttachementType `json:"contentType"`
	UserID    int             `json:"user_id"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
}
