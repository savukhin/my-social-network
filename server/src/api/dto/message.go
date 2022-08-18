package dto

import "time"

type MessageInput struct {
	Token        string        `json:"token"`
	AuthorID     int           `json:"author_id"`
	Text         string        `json:"text"`
	Attachements []Attachement `json:"attachements"`
	ChatID       int           `json:"chat_id"`
}

type Message struct {
	AuthorID     int           `json:"author_id"`
	CreatedAt    time.Time     `json:"time"`
	Text         string        `json:"text"`
	Attachements []Attachement `json:"attachements"`
	ChatID       int           `json:"chat_id"`
}
