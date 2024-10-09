package model

import (
	"time"
)

type ChatMessage struct {
	ID          int       `json:"id"`
	Chat        Chat      `json:"chat"`
	User        User      `json:"user"`
	Text        string    `json:"text"`
	CreatedTime time.Time `json:"created_time"`
}
