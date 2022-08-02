package Models

import "gorm.io/gorm"

type Address struct {
	gorm.Model
	OrderId int    `json:"orderId"`
	Address string `json:"address"`
}
