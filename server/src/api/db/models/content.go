package models

import (
	"database/sql"
	"time"
)

type ContentType string

const (
	Post    ContentType = "post"
	Message ContentType = "message"
	Comment ContentType = "comment"
	Photo   ContentType = "photo"
)

type Content struct {
	id                int           `json:"id"`
	filepath          string        `json:"filepath"`
	content_type      ContentType   `json:"contentType"`
	parent_content_id sql.NullInt32 `json:"parent_content_id"`
	user_id           int           `json:"user_id"`
	attach_order      int           `json:"attach_order"`
	created_at        time.Time     `json:"created_at"`
	updated_at        time.Time     `json:"updated_at"`
	deleted_at        sql.NullTime  `json:"deleted_at"`
}
