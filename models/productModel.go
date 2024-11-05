package models

import "time"

type Product struct {
	Id               int       `gorm:"primaryKey;autoIncrement" json:"id"`
	SKU              string    `json:"sku"`
	Name             string    `json:"name"`
	Stock            int       `json:"stock"`
	Price            float64   `json:"price"`
	Image            string    `json:"image"`
	TotalFinalPrice  int       `json:"total_final_price"`
	TotalNormalPrice int       `json:"total_normal_price"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
	CategoryId       int       `json:"categoryId"`
	DiscountId       int       `json:"discountId"`
}
