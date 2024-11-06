package models

import "time"

type Comment struct {
	Id          int       `gorm:"primaryKey;autoIncrement" json:"_id"`
	UserId      string    `json:"user_id'`
	CommentBody string    `json:"comment_body'`
	Deleted     bool      `json:"deleted"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   time.Time `json:"deleted_at"`
}
