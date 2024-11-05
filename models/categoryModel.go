package models

import "time"

type Category struct {
	Id        int       `gorm:"primaryKey;autoIncrement" json:"id"` // Auto-incremented unique identifier
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
