package models

import "time"

type GradeLevel struct {
	Id        int       `gorm:"primaryKey;autoIncrement" json:"_id"`
	Name      string    `json:"name'`
	Deleted   bool      `json:"deleted"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}
