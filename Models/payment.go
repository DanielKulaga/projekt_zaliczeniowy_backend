package Models

import "gorm.io/gorm"

type Payment struct {
	gorm.Model
	Amount int `json:"amount"`
}
