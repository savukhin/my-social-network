package dto

type UserProfile struct {
	ID         int    `json:"id"`
	Username   string `json:"username"`
	Name       string `json:"name"`
	IsOnline   bool   `json:"isOnline"`
	Status     string `json:"status"`
	BirthDate  string `json:"birthDate"`
	City       string `json:"city"`
	Avatar_URL string `json:"avatar_id"`
}
