package model

import "time"

type ChatMessage struct {
	ID          int       `json:"id"`
	ChatID      int       `json:"chatid"`
	UserID      int       `json:"userid"`
	Text        string    `json:"text"`
	CreatedTime time.Time `json:"createdtime"`
}
