package models

import (
	"database/sql"
	"time"
)

type Chat struct {
	id         int            `json:"id"`
	title      string         `json:"title"`
	photo_id   sql.NullString `json:"photo_id"`
	created_at time.Time      `json:"created_at"`
	updated_at time.Time      `json:"deleted_at"`
	deleted_at sql.NullTime   `json:"deleted_at"`
}
