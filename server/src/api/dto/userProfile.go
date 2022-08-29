package dto

type UserProfile struct {
	ID            int              `json:"id"`
	Username      string           `json:"username"`
	Name          string           `json:"name"`
	IsOnline      bool             `json:"is_online"`
	Status        string           `json:"status"`
	BirthDate     string           `json:"birth_date"`
	City          string           `json:"city"`
	AvatarURL     string           `json:"avatar_url"`
	Friends       []UserCompressed `json:"friends"`
	AddedToFriend bool             `json:"added_to_friends"`
}
