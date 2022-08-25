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

func GetUserAvatarURL(id int) (string, error) {
	filename := ""

	sql := fmt.Sprintf(`
		SELECT filepath
		FROM contents as c
		INNER JOIN users as u ON u.avatar_id = c.id AND c.content_type = 'photo' AND u.id = %d
	`, id)

	err := db.DB.QueryRow(sql).Scan(&filename)

	return filename, err
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

func GetChatParticipants(chat_id int) ([]int, error) {
	sql := fmt.Sprintf(`
		SELECT user_id FROM user_to_chat WHERE chat_id = %d
	`, chat_id)

	rows, err := db.DB.Query(sql)
	if err != nil {
		return nil, err
	}

	result := make([]int, 0)

	for rows.Next() {
		var user_id int
		err = rows.Scan(&user_id)
		if err != nil {
			return nil, err
		}

		result = append(result, user_id)
	}

	return result, nil
}

func GetChatsByUserID(id int) ([]*Chat, error) {
	sql := fmt.Sprintf(`
		SELECT c.id, c.title, c.photo_id, c.is_personal, c.created_at, c.updated_at 
		FROM chats as c
		JOIN user_to_chat AS u ON c.id = u.chat_id 
		AND u.user_id = %d AND c.deleted_at IS NULL 
	`, id)

	rows, err := db.DB.Query(sql)

	if err != nil {
		return nil, err
	}

	chats := make([]*Chat, 0)

	if rows.Next() {
		chat := &Chat{}
		err := rows.Scan(&chat.ID, &chat.Title, &chat.PhotoID, &chat.IsPersonal, &chat.CreatedAt, &chat.UpdatedAt)
		if err != nil {
			return nil, err
		}

		if chat.IsPersonal {
			participants, err := GetChatParticipants(chat.ID)
			if err != nil || len(participants) != 2 {
				fmt.Println("is pers 1 ", err, " len = ", len(participants))
				return nil, err
			}

			processing_id := 0
			if participants[0] == id {
				processing_id = 1
			}

			other, err := GetUserByID(processing_id)

			if err != nil {
				fmt.Println("is pers 2 ", err)
				return nil, err
			}

			chat.Title = other.Name
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
		AND c.deleted_at IS NULL
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

	user2, _ := GetUserByID(user2_id)
	chat.Title = user2.Name

	return chat, nil
}

func GetChatByID(chat_id int) (*Chat, error) {
	sql := fmt.Sprintf(`
		SELECT id, title, is_personal, photo_id, created_at, updated_at, deleted_at 
		FROM chats
		WHERE id = %d
	`, chat_id)

	chat := &Chat{}
	err :=
		db.DB.QueryRow(sql).
			Scan(&chat.ID, &chat.Title, &chat.IsPersonal, &chat.PhotoID, &chat.CreatedAt, &chat.UpdatedAt, &chat.DeletedAt)

	if err != nil {
		return nil, err
	}

	return chat, nil
}
