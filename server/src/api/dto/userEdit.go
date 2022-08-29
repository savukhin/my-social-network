package dto

import (
	"api/db"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

type UserEdit struct {
	Status    string `json:"status,omitempty"`
	Name      string `json:"name,omitempty"`
	BirthDate string `json:"birth_date,omitempty"`
	City      string `json:"city,omitempty"`
}

func (user UserEdit) ApllyChanges(id int) error {
	sql := `
		UPDATE
			users
		SET`

	queries := make([]string, 0)
	tmp, _ := json.Marshal(user)

	var response map[string]interface{}
	json.Unmarshal(tmp, &response)

	for key, value := range response {
		queries = append(queries, fmt.Sprintf(" %s = '%s'", key, value.(string)))
	}
	sql += strings.Join(queries[:], ",")

	sql += fmt.Sprintf(` WHERE id = %d RETURNING id`, id)

	idUser := 0
	err := db.DB.QueryRow(sql).Scan(&idUser)
	if err != nil {
		return err
	}
	if err != nil || idUser == 0 {
		return errors.New("unknow error inserting :(")
	}

	return nil
}
