package mappers

import (
	"api/db/models"
	"api/dto"
	"database/sql"
	"errors"
	"time"
)

func LikePostToLike(like *dto.LikePost, user_id int) *models.Like {
	model := &models.Like{
		UserID:    user_id,
		ContentID: like.PostID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		DeletedAt: sql.NullTime{Valid: false},
	}

	return model
}

func ToLikePost(like *models.Like) (*dto.LikePostSend, error) {
	content, err := models.GetPostByID(like.ContentID)
	if err != nil {
		return nil, err
	}

	if content.Type != models.PostType {
		return nil, errors.New("content is not post")
	}

	data := &dto.LikePostSend{
		UserID: like.UserID,
		PostID: like.ContentID,
	}

	return data, nil
}
