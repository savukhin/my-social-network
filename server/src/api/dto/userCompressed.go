package dto

type UserCompressed struct {
	ID        int    `json:"id"`
	Username  string `json:"username"`
	Name      string `json:"name"`
	IsOnline  bool   `json:"is_online"`
	Status    string `json:"status"`
	AvatarURL string `json:"avatar_url"`
}
