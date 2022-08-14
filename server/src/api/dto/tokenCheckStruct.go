package dto

type TokenCheckStruct struct {
	Token   string `json:"id_token"`
	ExpiresAt string `json:"expires_at"`
	UserProfile
}
