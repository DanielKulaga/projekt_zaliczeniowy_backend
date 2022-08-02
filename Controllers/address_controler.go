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
