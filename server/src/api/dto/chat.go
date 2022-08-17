package dto

type Chat struct {
	ID           int              `json:"id"`
	Title        string           `json:"title"`
	IsPersonal   bool             `json:"is_personal"`
	PhotoURL     string           `json:"photo_url"`
	Participants []UserCompressed `json:"participants"`
}
