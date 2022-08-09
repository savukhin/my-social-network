package models

type User struct {
	ID        int    `json:"id,omitempty"`
	Username  string `json:"username,omitempty"`
	Name      string `json:"name,omitempty"`
	Email     string `json:"email,omitempty"`
	Password  string `json:"password,omitempty"`
	IsOnline  bool   `json:"isOnline,omitempty"`
	Status    string `json:"status,omitempty"`
	BirthDate string `json:"birthDate,omitempty"`
	City      string `json:"city,omitempty"`
	Avatar_ID int    `json:"avatar_id,omitempty"`
}
