package Controllers

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"myapp/Database"
	"myapp/Models"
	"net/http"
)

func AddAddressToDatabase(c echo.Context) error {
	addressDetails := new(Models.Address)

	fmt.Printf("Add new address details \n")

	if err := c.Bind(addressDetails); err != nil {
		return c.String(http.StatusBadRequest, "Bad address details "+err.Error())
	}

	Database.Database.Create(addressDetails)

	return c.JSON(http.StatusOK, addressDetails)
}

func GetAddressFromDatabase(c echo.Context) error {
	id := c.Param("id")
	var address Models.Address

	fmt.Printf("Get address with order_id: " + id + "\n")

	if result := Database.Database.Where("order_id", id).Find(&address); result.Error != nil {
		return c.String(http.StatusNotFound, "Database Error")
	}

	return c.JSON(http.StatusOK, address)
}
