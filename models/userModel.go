package models

import "time"

type User struct {
	Id         int        `gorm:"primaryKey;autoIncrement" json:"_id"`
	Firstname  string     `json:"first_name'`
	Lastname   string     `json:"last_name'`
	Username   string     `json:"username'`
	Password   string     `json:"password'`
	Role       string     `json:"role'` //Admin/Student/Teacher
	GradeLevel GradeLevel `json:"grade_level'`
	Address    string     `json:"address'`
	Online     bool       `json:"online'`
	Deleted    bool       `json:"deleted"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
	DeletedAt  time.Time  `json:"deleted_at"`
}
