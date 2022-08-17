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
	ID          int           `json:"id"`
	Filepath    string        `json:"filepath"`
	Type        ContentType   `json:"contentType"`
	ParentID    sql.NullInt32 `json:"parent_id"`
	UserID      int           `json:"user_id"`
	AttachOrder int           `json:"attach_order"`
	CreatedAt   time.Time     `json:"created_at"`
	UpdatedAt   time.Time     `json:"updated_at"`
	DeletedAt   sql.NullTime  `json:"deleted_at"`
}

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

func (content *Content) Save() (int, error) {
	sql := fmt.Sprintf(`
	INSERT INTO contents (filepath, content_type, parent_id, user_id, attach_order)
	VALUES ('%s', '%s', %d, %d, %d)
	RETURNING id
	`, content.Filepath, content.Type, content.ParentID.Int32, content.UserID, content.AttachOrder)

	content_id := 0
	err := db.DB.QueryRow(sql).Scan(&content_id)

	return content_id, err
}
