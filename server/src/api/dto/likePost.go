package dto

type LikePostSend struct {
	PostID int `json:"post_id"`
	UserID int `json:"user_id"`
}

type LikePost struct {
	PostID int `json:"post_id"`
}
