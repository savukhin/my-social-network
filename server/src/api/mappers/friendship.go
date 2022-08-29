package mappers

import (
	"api/db/models"
	"api/dto"
)

func AddToFriendsToFriendship(form *dto.AddToFriends, user_id int) *models.Friendship {
	friendship := &models.Friendship{
		User1ID: user_id,
		User2ID: form.AddingUserID,
	}

	return friendship
}
