package dto

import (
	"time"
)

type PostCreate struct {
	Text string `json:"text"`
}

type Post struct {
	ID        int                 `json:"id"`
	Text      string              `json:"text"`
	AuthodID  int                 `json:"author_id"`
	CreatedAt time.Time           `json:"created_at"`
	UpdatedAt time.Time           `json:"updated_at"`
	Photos    []*PhotoAttachement `json:"photos"`
}
