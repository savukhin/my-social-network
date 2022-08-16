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

func CreatePersonalChat(user1_id int, user2_id int) (*Chat, error) {
	sql := fmt.Sprintf(`
		INSERT INTO chats (title) 
		VALUES ('PERSONAL')
		RETURNING id, photo_id
	`)

	chat := &Chat{}

	err := db.DB.QueryRow(sql).Scan(&chat.ID, &chat.PhotoID)
	if err != nil {
		return nil, err
	}

	sql = fmt.Sprintf(`
		INSERT INTO user_to_chat (user_id, chat_id) VALUES (%d, %d) RETURNING id
	`, user1_id, chat.ID)
	id := 0
	err = db.DB.QueryRow(sql).Scan(&id)
	if err != nil {
		return nil, err
	}

	sql = fmt.Sprintf(`
		INSERT INTO user_to_chat (user_id, chat_id) VALUES (%d, %d) RETURNING id
	`, user2_id, chat.ID)
	id = 0
	err = db.DB.QueryRow(sql).Scan(&id)
	if err != nil {
		return nil, err
	}

	return chat, nil
}

func GetChatsByUserID(id int) ([]*Chat, error) {
	sql := fmt.Sprintf(`
		SELECT id, title, photo_id, created_at, updated_at 
		FROM chats as c
		JOIN user_to_chat AS u ON c.id = u.chat_id
		WHERE u.user_id = %d AND deleted_at = NULL 
	`, id)

	rows, err := db.DB.Query(sql)

	if err != nil {
		return nil, err
	}

	chats := make([]*Chat, 0)

	if rows.Next() {
		chat := &Chat{}
		err := rows.Scan(&chat.ID, &chat.Title, &chat.PhotoID, &chat.CreatedAt, &chat.UpdatedAt)
		if err != nil {
			return nil, err
		}

		chats = append(chats, chat)
	}

	return chats, nil
}

func GetPersonalChat(user1_id int, user2_id int) (*Chat, error) {
	sql := fmt.Sprintf(`
		SELECT c.id, c.title, c.photo_id, c.created_at, c.updated_at 
		FROM chats as c
		JOIN user_to_chat AS u1 ON c.id = u1.chat_id AND u1.user_id = %d
		JOIN user_to_chat AS u2 ON c.id = u2.chat_id AND u2.user_id = %d
	`, user1_id, user2_id)

	chat := &Chat{}
	err :=
		db.DB.QueryRow(sql).
			Scan(&chat.ID, &chat.Title, &chat.PhotoID, &chat.CreatedAt, &chat.UpdatedAt)

	if err != nil {
		chat, err = CreatePersonalChat(user1_id, user2_id)

		if err != nil {
			return nil, err
		}
	}

	return chat, nil
}
