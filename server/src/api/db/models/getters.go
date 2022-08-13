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
