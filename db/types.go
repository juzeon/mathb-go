package db

import "time"

type Paste struct {
	ID        int       `json:"id"`
	UUID      string    `json:"uuid" gorm:"uniqueIndex"`
	Title     string    `json:"title"`
	Name      string    `json:"name"`
	Code      string    `json:"code"`
	CreatedAt time.Time `json:"created_at"`
}
