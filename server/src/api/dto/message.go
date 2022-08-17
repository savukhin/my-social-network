package dto

import "time"

type Message struct {
	AuthorID     int           `json:"author_id"`
	CreatedAt    time.Time     `json:"time"`
	Text         string        `json:"text"`
	Attachements []Attachement `json:"attachements"`
}

func GetMessageByID() {

}
