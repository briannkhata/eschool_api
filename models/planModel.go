package models

import "time"

type Plan struct {
	Id        int       `gorm:"primaryKey;autoIncrement" json:"_id"`
	Name      string    `json:"name'`
	Duration  string    `json:"duration'`
	Price     float64   `json:"price'`
	Deleted   bool      `json:"deleted"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}
