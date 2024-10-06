package model

import "time"

type Chat struct {
	ID          int
	User1Id     int
	User2Id     int
	CreatedTime time.Time
}
