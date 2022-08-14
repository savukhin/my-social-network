package mappers

import (
	"api/db/models"
	"api/dto"
	"strconv"
	"time"
)

func ToUserProfile(user *models.User) *dto.UserProfile {
	response := &dto.UserProfile{}

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
		response.Avatar_URL = strconv.Itoa(int(user.Avatar_ID.Int64))
	}

	return response
}

func ToTokenCheckStruct(user *models.User, id_token string, expires_at time.Time) *dto.TokenCheckStruct {
	res2 := ToUserProfile(user)
	response := &dto.TokenCheckStruct{Token: id_token, ExpiresAt: expires_at.String(), UserProfile: *res2}
	return response
}
