package model

import "time"

type Chat struct {
	ID          int       `json:"id"`
	User1       User      `json:"user1"`
	User2       User      `json:"user2"`
	CreatedTime time.Time `json:"created_time"`
}
