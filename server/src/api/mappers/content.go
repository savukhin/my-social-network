package mappers

import (
	"api/db/models"
	"api/dto"
	"errors"
	"strconv"
)

func ToMessage(content *models.Content) (*dto.Message, error) {
	result := &dto.Message{}

	result.Text = content.Filepath
	result.AuthorID = content.UserID
	result.CreatedAt = content.CreatedAt

	return result, nil
}

func ToMessageRange(contents []models.Content) (*dto.MessageRangeOutput, error) {
	result := &dto.MessageRangeOutput{}

	for i, content := range contents {
		if content.Type != "message" {
			return nil, errors.New("content [" + strconv.Itoa(i) + "] is not a message")
		}
		message, err := ToMessage(&content)
		if err != nil {
			return nil, err
		}

		result.Messages = append(result.Messages, *message)
	}

	return result, nil
}
