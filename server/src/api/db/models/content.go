package models

import (
	"api/db"
	"database/sql"
	"fmt"
	"io/ioutil"
	"time"
)

type ContentType string

const (
	PostType    ContentType = "post"
	MessageType ContentType = "message"
	CommentType ContentType = "comment"
	PhotoType   ContentType = "photo"
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

func (content *Content) GetText() (string, error) {
	b, err := ioutil.ReadFile(content.Filepath)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

func getContentsFromQuery(rows *sql.Rows) ([]*Content, error) {
	result := make([]*Content, 0)
	for rows.Next() {
		content := &Content{}
		err := rows.Scan(&content.ID, &content.Filepath, &content.Type, &content.UserID, &content.CreatedAt)
		if err != nil {
			return nil, err
		}

		result = append(result, content)
	}

	return result, nil
}

func (content *Content) GetAttachements() ([]*Content, error) {
	sql := fmt.Sprintf(`
		SELECT id, filepath, content_type, user_id, created_at 
		FROM contents
		WHERE content_type = 'photo' AND parent_id = %d
		ORDER BY created_at DESC
	`, content.ID)

	rows, err := db.DB.Query(sql)
	if err != nil {
		return nil, err
	}

	result, err := getContentsFromQuery(rows)

	return result, err
}

func GetMessages(offset int, count int, chat_id int) ([]*Content, error) {
	sql := fmt.Sprintf(`
		SELECT id, filepath, content_type, user_id, created_at 
		FROM contents
		WHERE content_type = 'message' AND parent_id = %d
		ORDER BY created_at DESC
	`, chat_id)

	rows, err := db.DB.Query(sql)
	if err != nil {
		return nil, err
	}

	result, err := getContentsFromQuery(rows)
	return result, err
}

func GetPostsByUserID(user_id int) ([]*Content, error) {
	sql := fmt.Sprintf(`
		SELECT id, filepath, content_type, user_id, created_at 
		FROM contents
		WHERE content_type = 'post' AND user_id = %d
		ORDER BY created_at DESC
	`, user_id)

	rows, err := db.DB.Query(sql)
	if err != nil {
		return nil, err
	}

	result, err := getContentsFromQuery(rows)
	return result, err
}

func GetPostByID(post_id int) (*Content, error) {
	sql := fmt.Sprintf(`
		SELECT id, filepath, content_type, user_id, created_at 
		FROM contents
		WHERE content_type = 'post' AND id = %d
		ORDER BY created_at DESC
	`, post_id)

	content := &Content{}

	err := db.DB.
		QueryRow(sql).
		Scan(&content.ID, &content.Filepath, &content.Type, &content.UserID, &content.CreatedAt)

	if err != nil {
		return nil, err
	}

	return content, err
}

func CreateAvatarContent(path string, user_id int) *Content {
	content := &Content{
		Filepath:    path,
		Type:        PhotoType,
		ParentID:    sql.NullInt32{Int32: int32(user_id), Valid: true},
		UserID:      user_id,
		AttachOrder: 1,
	}

	return content
}

func (content *Content) Save() (int, error) {
	sql := fmt.Sprintf(`
	INSERT INTO contents (filepath, content_type, parent_id, user_id, attach_order)
	VALUES ('%s', '%s', %d, %d, %d)
	RETURNING id
	`, content.Filepath, content.Type, content.ParentID.Int32, content.UserID, content.AttachOrder)

	err := db.DB.QueryRow(sql).Scan(&content.ID)

	return content.ID, err
}
