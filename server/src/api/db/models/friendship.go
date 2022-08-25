package models

import (
	"api/db"
	"database/sql"
	"fmt"
	"time"
)

type Friendship struct {
	ID        int          `json:"id"`
	User1ID   int          `json:"user1_id"`
	User2ID   int          `json:"user2_id"`
	CreatedAt time.Time    `json:"created_at,omitempty"`
	UpdatedAt time.Time    `json:"updated_at,omitempty"`
	DeletedAt sql.NullTime `json:"deleted_at,omitempty"`
}

func (friendship *Friendship) Save() (int, error) {
	sql := fmt.Sprintf(`
		INSERT INTO friendships (user1_id, user2_id)
		VALUES (%d, %d)
		RETURNING id
		`, friendship.User1ID, friendship.User2ID)

	err := db.DB.QueryRow(sql).Scan(&friendship.ID)

	return friendship.ID, err
}

func (friendship *Friendship) Delete() error {
	sql := fmt.Sprintf(`
		UPDATE friendships 
		SET deleted_at = now()
		WHERE id = %d
		RETURNING id
		`, friendship.ID)

	err := db.DB.QueryRow(sql).Scan(&friendship.ID)

	return err
}

func GetFriendships(user_id int) ([]*Friendship, error) {
	sql := fmt.Sprintf(`
		SELECT id, user1_id, user2_id, created_at, updated_at, deleted_at
		FROM friendships
		WHERE (user1_id = %d OR user2_id = %d) AND deleted_at IS NULL
		ORDER BY created_at DESC
	`, user_id, user_id)

	rows, err := db.DB.Query(sql)
	if err != nil {
		return nil, err
	}

	result := make([]*Friendship, 0)

	for rows.Next() {
		friendship := &Friendship{}
		err := rows.Scan(&friendship.ID, &friendship.User1ID, &friendship.User2ID, &friendship.CreatedAt, &friendship.UpdatedAt, &friendship.DeletedAt)

		if err != nil {
			return nil, err
		}

		result = append(result, friendship)
	}

	return result, err
}

func GetFriendship(user1_id int, user2_id int) (*Friendship, error) {
	sql := fmt.Sprintf(`
		SELECT id, user1_id, user2_id, created_at, updated_at, deleted_at
		FROM friendships 
		WHERE (user1_id = %d AND user2_id = %d) OR (user2_id = %d AND user1_id = %d)
	`, user1_id, user2_id, user1_id, user2_id)

	friendship := &Friendship{}

	err := db.DB.QueryRow(sql).Scan(&friendship.ID, &friendship.User1ID, &friendship.User2ID, &friendship.CreatedAt, &friendship.UpdatedAt, &friendship.DeletedAt)

	return friendship, err
}
