package mappers

import (
	"api/db/models"
	"api/dto"
	"database/sql"
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
)

func ToMessage(content *models.Content) (*dto.Message, error) {
	result := &dto.Message{}

	b, err := ioutil.ReadFile(content.Filepath)
	if err != nil {
		return nil, err
	}

	result.Text = string(b)
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

func MessageToContent(message *dto.MessageInput) (*models.Content, error) {
	content := &models.Content{
		Type:        models.Message,
		ParentID:    sql.NullInt32{Int32: int32(message.ChatID), Valid: true},
		UserID:      message.AuthorID,
		AttachOrder: 1,
	}

	path := filepath.Join("uploads", "messages", strconv.Itoa(message.ChatID))
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return nil, err
	}

	files, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}
	filename := filepath.Join(path, strconv.Itoa(len(files)))

	err = ioutil.WriteFile(filename, []byte(message.Text), os.ModePerm)
	if err != nil {
		return nil, err
	}

	content.Filepath = filename

	return content, nil
}

func ToChatDTO(chat_model *models.Chat, participants []models.User) (*dto.Chat, error) {
	chat := &dto.Chat{
		ID:         chat_model.ID,
		Title:      chat_model.Title,
		IsPersonal: chat_model.IsPersonal,
	}

	chat.PhotoURL = ""
	chat.Participants = make([]dto.UserCompressed, 0)
	for _, participant := range participants {
		user := ToUserCompressed(&participant)
		chat.Participants = append(chat.Participants, *user)
	}

	return chat, nil
}
