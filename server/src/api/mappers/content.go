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

	text, err := content.GetText()
	if err != nil {
		return nil, err
	}

	result.Text = text
	result.AuthorID = content.UserID
	result.CreatedAt = content.CreatedAt
	result.ChatID = int(content.ParentID.Int32)

	return result, nil
}

func ToMessageRange(contents []*models.Content) (*dto.MessageRangeOutput, error) {
	result := &dto.MessageRangeOutput{
		Messages: make([]dto.Message, 0),
	}

	for i, content := range contents {
		if content.Type != "message" {
			return nil, errors.New("content [" + strconv.Itoa(i) + "] is not a message")
		}
		message, err := ToMessage(content)
		if err != nil {
			return nil, err
		}

		result.Messages = append(result.Messages, *message)
	}

	return result, nil
}

func saveText(text string, folder string) (string, error) {
	path := filepath.Join("uploads", folder)
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return "", err
	}

	files, err := ioutil.ReadDir(path)
	if err != nil {
		return "", err
	}
	filename := filepath.Join(path, strconv.Itoa(len(files)))

	err = ioutil.WriteFile(filename, []byte(text), os.ModePerm)
	if err != nil {
		return "", err
	}

	return filename, nil
}

func MessageToContent(message *dto.MessageInput) (*models.Content, error) {
	content := &models.Content{
		Type:        models.MessageType,
		ParentID:    sql.NullInt32{Int32: int32(message.ChatID), Valid: true},
		UserID:      message.AuthorID,
		AttachOrder: 1,
	}

	filename, err := saveText(message.Text, filepath.Join("messages", strconv.Itoa(message.ChatID)))
	if err != nil {
		return nil, err
	}

	content.Filepath = filename

	return content, nil
}

func ToChatDTO(chat_model *models.Chat, participants []*models.User, lastMessage *dto.Message) (*dto.Chat, error) {
	chat := &dto.Chat{
		ID:          chat_model.ID,
		Title:       chat_model.Title,
		IsPersonal:  chat_model.IsPersonal,
		LastMessage: lastMessage,
	}

	chat.PhotoURL = ""
	chat.Participants = make([]dto.UserCompressed, 0)
	for _, participant := range participants {
		user := ToUserCompressed(participant)
		chat.Participants = append(chat.Participants, *user)
	}

	return chat, nil
}

func ContentToPhotoAttachement(content *models.Content) (*dto.PhotoAttachement, error) {
	if content.Type != models.PhotoType {
		return nil, errors.New("content type is not photo")
	}

	photo := &dto.PhotoAttachement{
		ID:        content.ID,
		AuthorID:  content.UserID,
		CreatedAt: content.CreatedAt,
		UpdatedAt: content.UpdatedAt,
		URL:       content.Filepath,
	}

	return photo, nil
}

func ContentToPost(content *models.Content) (*dto.Post, error) {
	if content.Type != models.PostType {
		return nil, errors.New("content type is not post")
	}

	post := &dto.Post{
		ID:        content.ID,
		AuthodID:  content.UserID,
		CreatedAt: content.CreatedAt,
		UpdatedAt: content.UpdatedAt,
		Photos:    make([]*dto.PhotoAttachement, 0),
	}

	text, err := content.GetText()
	if err != nil {
		return nil, err
	}
	post.Text = text

	attachements, err := content.GetAttachements()
	if err != nil {
		return nil, err
	}

	for _, attach := range attachements {
		photo, err := ContentToPhotoAttachement(attach)
		if err == nil {
			post.Photos = append(post.Photos, photo)
			continue
		}
	}

	return post, nil
}

func PostCreateToContent(post *dto.PostCreate, user_id int) (*models.Content, error) {
	content := &models.Content{
		ID:          0,
		Type:        models.PostType,
		ParentID:    sql.NullInt32{Valid: false},
		UserID:      user_id,
		AttachOrder: 1,
	}

	filename, err := saveText(post.Text, filepath.Join("posts", strconv.Itoa(user_id)))
	if err != nil {
		return nil, err
	}

	content.Filepath = filename

	return content, nil
}
