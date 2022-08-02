package Models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	Dishes string `json:"dishes"`
	Email  int    `json:"email"`
}
