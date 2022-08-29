package models

import (
	"api/db"
	"database/sql"
	"fmt"
	"time"
)

type Like struct {
	ID        int          `json:"id"`
	UserID    int          `json:"user_id"`
	ContentID int          `json:"content_id"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
	DeletedAt sql.NullTime `json:"deleted_at,omitempty"`
}

func (like *Like) Save() (int, error) {
	sql := fmt.Sprintf(`
	INSERT INTO likes (user_id, content_id)
	VALUES (%d, %d)
	RETURNING id
	`, like.UserID, like.ContentID)

	err := db.DB.QueryRow(sql).Scan(&like.ID)

	return like.ID, err
}

func GetLikesByContent(content_id int) ([]*Like, error) {
	sql := fmt.Sprintf(`
		SELECT id, user_id, content_id, created_at, updated_at, deleted_at
		FROM likes
		WHERE content_id = %d AND deleted_at IS NULL
		ORDER BY created_at DESC
	`, content_id)

	rows, err := db.DB.Query(sql)
	if err != nil {
		return nil, err
	}

	result := make([]*Like, 0)
	for rows.Next() {
		like := &Like{}
		err := rows.Scan(&like.ID, &like.UserID, &like.ContentID, &like.CreatedAt, &like.UpdatedAt, &like.DeletedAt)
		if err != nil {
			return nil, err
		}

		result = append(result, like)
	}

	return result, nil
}

func GetLike(content_id int, user_id int) (*Like, error) {
	sql := fmt.Sprintf(`
		SELECT id, user_id, content_id, created_at, updated_at, deleted_at
		FROM likes
		WHERE content_id = %d AND user_id = %d AND deleted_at IS NULL
		ORDER BY created_at DESC
	`, content_id, user_id)

	like := &Like{}
	err := db.DB.QueryRow(sql).Scan(&like.ID, &like.UserID, &like.ContentID, &like.CreatedAt, &like.UpdatedAt, &like.DeletedAt)

	return like, err
}

func DeleteLike(like_id int) error {
	sql := fmt.Sprintf(`
		UPDATE likes
		SET deleted_at = NOW()
		WHERE id = %d
		RETURNING id
	`, like_id)

	id := 0
	err := db.DB.QueryRow(sql).Scan(&id)
	return err
}
