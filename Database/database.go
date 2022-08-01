package Database

import "fmt"
import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"myapp/Models"
)

var Database *gorm.DB = nil

func ConnectDataBase() {
	// GORM configuration
	db, err := gorm.Open(sqlite.Open("./Database/restaurant.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	//Migrate the schema
	//db.AutoMigrate(&models.Order{})
	db.AutoMigrate(&Models.Category{})
	//db.AutoMigrate(&models.Payment{})
	db.AutoMigrate(&Models.Item{})
	//db.AutoMigrate(&models.Delivery{})
	db.AutoMigrate(&Models.User{})

	Database = db
	fmt.Printf("Database connected!")
}
