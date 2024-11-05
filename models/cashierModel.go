package models

import "time"

type Cashier struct {
	Id        int       `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string    `json:"name'`
	Passcode  string    `json:"passcode"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
