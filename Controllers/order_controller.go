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

	fmt.Printf("Add new order to databse \n")

	if err := c.Bind(orderDetails); err != nil {
		return c.String(http.StatusBadRequest, "Bad order details "+err.Error())
	}

	Database.Database.Create(orderDetails)

	return c.JSON(http.StatusOK, orderDetails)
}

func GetOrderFromDatabase(c echo.Context) error {
	id := c.Param("id")
	var order Models.Order

	fmt.Printf("Get order id: " + id + "\n")

	if result := Database.Database.First(&order, id); result.Error != nil {
		return c.String(http.StatusNotFound, "Database Error")
	}

	return c.JSON(http.StatusOK, order)
}
