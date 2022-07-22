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
	db, err := gorm.Open(sqlite.Open("./database/restaurant.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	//Migrate the schema
	//db.AutoMigrate(&models.Order{})
	db.AutoMigrate(&Models.Category{})
	//db.AutoMigrate(&models.Payment{})
	//db.AutoMigrate(&models.Product{})
	//db.AutoMigrate(&models.Delivery{})
	//db.AutoMigrate(&models.User{})

	Database = db
	fmt.Printf("Database connected!")
}
