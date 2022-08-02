package Controllers

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"myapp/Database"
	"myapp/Models"
	"net/http"
)

func CreateOrderToDatabase(c echo.Context) error {
	orderDetails := new(Models.Order)

	fmt.Printf("Add new order details \n")

	if err := c.Bind(orderDetails); err != nil {
		return c.String(http.StatusBadRequest, "Bad order details "+err.Error())
	}

	Database.Database.Create(orderDetails)

	return c.JSON(http.StatusOK, orderDetails)
}
