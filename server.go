package main

import (
	"github.com/labstack/echo/v4"
	"myapp/Controllers"
	"myapp/Database"
	"net/http"
)

func main() {
	Database.ConnectDataBase()

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/items", Controllers.GetItems)
	e.GET("/items/:id", Controllers.GetItemsInCategory)
	e.POST("/items", Controllers.PostItems)
	e.GET("/category", Controllers.GetRestaurantCategories)
	e.POST("/category", Controllers.PostRestaurantCategories)
	e.Logger.Fatal(e.Start(":1323"))
}
