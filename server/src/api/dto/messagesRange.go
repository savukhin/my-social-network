package dto

type MessageRangeInput struct {
	Offset int `json:"offset"`
	Count  int `json:"count"`
	ChatID int `json:"chat_id"`
}

type MessageRangeOutput struct {
	Messages []Message `json:"messages"`
}
