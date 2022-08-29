package models

import (
	"database/sql"
	"time"
)

type Chat struct {
	ID         int            `json:"id"`
	Title      string         `json:"title"`
	IsPersonal bool           `json:"is_personal"`
	PhotoID    sql.NullString `json:"photo_id"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"deleted_at"`
	DeletedAt  sql.NullTime   `json:"deleted_at"`
}
