package main

import (
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"myapp/Controllers"
	"myapp/Database"
	"net/http"
)

func main() {
	Database.ConnectDataBase()
	err := godotenv.Load(".env")
	if err != nil {
		print("Error loading .env file")
	}
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000", "https://restaurant-ruczaj.azurewebsites.net"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/items", Controllers.GetItems)
	e.GET("/items/:id", Controllers.GetItemsInCategory)
	e.POST("/items", Controllers.PostItems)
	e.GET("/category", Controllers.GetRestaurantCategories)
	e.POST("/category", Controllers.PostRestaurantCategories)
	e.POST("/create-payment-intent", Controllers.HandleCreatePaymentIntent)
	e.GET("/login/google", Controllers.GoogleLogin)
	e.GET("/login/github", Controllers.GithubLogin)
	e.GET("/google/callback", Controllers.GoogleCallback)
	e.GET("/github/callback", Controllers.GithubCallback)
	e.Logger.Fatal(e.Start(":1323"))

}
