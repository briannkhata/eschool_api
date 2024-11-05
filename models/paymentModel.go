package models

import "time"

type Payment struct {
	Id            int       `gorm:"primaryKey;autoIncrement" json:"id"`
	Name          int       `json:"name"`
	Type          string    `json:"type"`
	PaymentTypeId int       `json:"payment_type_id"`
	Logo          string    `json:"logo"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type PaymentType struct {
	Id        int       `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      int       `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}
