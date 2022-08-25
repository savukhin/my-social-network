package mappers

import (
	"api/db/models"
	"api/dto"
	"strconv"
	"time"
)

func ToUserProfile(user *models.User) *dto.UserProfile {
	response := &dto.UserProfile{
		Friends:       make([]dto.UserCompressed, 0),
		AddedToFriend: false,
	}

	response.ID = user.ID
	response.Username = user.Username
	response.Name = user.Name
	response.IsOnline = user.IsOnline
	if user.Status.Valid {
		response.Status = user.Status.String
	}
	if user.BirthDate.Valid {
		response.BirthDate = user.BirthDate.Time.Format("02.01.2006")
	}
	if user.City.Valid {
		response.City = user.City.String
	}
	if user.Avatar_ID.Valid {
		response.AvatarURL = strconv.Itoa(int(user.Avatar_ID.Int64))
	}

	return response
}

func ToUserCompressed(user *models.User) *dto.UserCompressed {
	result := &dto.UserCompressed{
		ID:        user.ID,
		Username:  user.Username,
		Name:      user.Name,
		IsOnline:  user.IsOnline,
		Status:    user.Status.String,
		AvatarURL: "",
	}

	return result
}

func ToTokenCheckStruct(user *models.User, id_token string, expires_at time.Time) *dto.TokenCheckStruct {
	res2 := ToUserProfile(user)
	response := &dto.TokenCheckStruct{Token: id_token, ExpiresAt: expires_at.String(), UserProfile: *res2}
	return response
}
