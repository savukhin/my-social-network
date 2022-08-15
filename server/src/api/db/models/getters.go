package models

import (
	"api/db"
	"fmt"
)

func GetUserByID(id int) (*User, error) {
	sql := fmt.Sprintf("SELECT id, username, name, email, status, city, birthdate, avatar_id, isonline FROM users WHERE id = %d", id)
	row := db.DB.QueryRow(sql)

	temp := &User{}
	err := row.Scan(
		&temp.ID, &temp.Username, &temp.Name, &temp.Email,
		&temp.Status, &temp.City, &temp.BirthDate, &temp.Avatar_ID, &temp.IsOnline,
	)

	return temp, err
}

func GetChatsByUserID(id int) (*Chat, error) {

	sql := fmt.Sprintf(`
		SELECT id, title, photo_id, created_at, updated_at 
		JOIN user_to_chat AS u ON c.id == u.chat_id
		WHERE u.user_id = %d AND deleted_at == NULL 
	`, id)

	row := db.DB.QueryRow(sql)

	temp := &Chat{}
	err := row.Scan(&temp.id, &temp.title, &temp.photo_id, &temp.created_at, &temp.updated_at)
	if err != nil {
		return nil, err
	}

	return temp, nil
}
