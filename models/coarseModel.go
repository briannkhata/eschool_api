package models

import "time"

type Coarse struct {
	Id           int       `gorm:"primaryKey;autoIncrement" json:"_id"`
	Title         string    `json:"name'`
	URL          string    `json:"url'`
	GradeLevelID int       `json:"grade_level_id'`
	Description  string    `json:"description'`
	Deleted      bool      `json:"deleted"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	DeletedAt    time.Time `json:"deleted_at"`
}
