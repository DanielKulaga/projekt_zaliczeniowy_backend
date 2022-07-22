package Models

import "gorm.io/gorm"

type Item struct {
	gorm.Model
	Name       string `json:"name"`
	Price      int    `json:"price"`
	Photo      string `json:"photo"`
	CategoryID int
	Category   Category
}
