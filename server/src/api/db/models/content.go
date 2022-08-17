package models

import (
	"api/db"
	"database/sql"
	"fmt"
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
	ID              int           `json:"id"`
	Filepath        string        `json:"filepath"`
	Type            ContentType   `json:"contentType"`
	ParentContentID sql.NullInt32 `json:"parent_content_id"`
	UserID          int           `json:"user_id"`
	AttachOrder     int           `json:"attach_order"`
	CreatedAt       time.Time     `json:"created_at"`
	UpdatedAt       time.Time     `json:"updated_at"`
	DeletedAt       sql.NullTime  `json:"deleted_at"`
}

// func GetMessageByID(message_id)

func GetMessages(offset int, count int, chat_id int) ([]Content, error) {
	sql := fmt.Sprintf(`
		SELECT id, filepath, content_type, user_id, created_at 
		FROM contents
		WHERE content_type = 'message' AND parent_id = %d
	`, chat_id)

	rows, err := db.DB.Query(sql)
	if err != nil {
		return nil, err
	}

	result := make([]Content, 0)
	for rows.Next() {
		message := &Content{}
		err := rows.Scan(&message.ID, &message.Filepath, &message.Type, &message.UserID, &message.CreatedAt)
		if err != nil {
			return nil, err
		}

		result = append(result, *message)
	}
	return result, nil
}
