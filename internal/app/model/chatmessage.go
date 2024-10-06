package model

import "time"

type ChatMessage struct {
	ID          int
	ChatID      int
	UserID      int
	Text        string
	CreatedTime time.Time
}
