package main

import (
	"github.com/labstack/echo/v4"
	"myapp/Database"
	"net/http"
)

func main() {
	Database.ConnectDataBase()
	
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
