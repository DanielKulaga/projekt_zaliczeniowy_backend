package Controllers

import (
	"fmt"
	"log"
	"myapp/Database"
	"myapp/Models"
)

func FindUserInDB(email string) bool {
	var user Models.User
	Database.Database.Find(&user, "email = ?", email)
	if user.Email == "" {
		return false
	}
	return true
}

func AddUser(user Models.User) {
	fmt.Printf("Add new user \n")
	log.Printf("Add new user: %v", user)
	Database.Database.Create(&user)
}

func EditUserToken(user string, token string) {
	log.Printf("Edit user: %v", user)
	log.Printf("New token user: %v", token)

	userInDB := Models.User{}

	Database.Database.Where("email = ?", user).Find(&userInDB)

	userInDB.Usertoken = token

	Database.Database.Save(userInDB)
}
