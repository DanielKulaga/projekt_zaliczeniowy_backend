package Models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username   string `json:"username"`
	Email      string `json:"email"`
	Oauthtoken string `json:"oauthtoken"`
	Usertoken  string `json:"usertoken"`
}
