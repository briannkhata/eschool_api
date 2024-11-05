package models

import "time"

type Discount struct {
	Id              int       `gorm:"primaryKey;autoIncrement" json:"id"`
	Qty             int       `json:"qty"`
	Type            string    `json:"type"`
	Result          int       `json:"result"`
	ExpiredAt       int       `json:"expiredAt"`
	ExpiredAtFormat string    `json:"axpiredAtFormat"`
	StringFormat    string    `json:"StringFormat"`
	CreatedAt       time.Time `json:"created_at'`
	UpdatedAt       time.Time `json:"updated_at'`
}
