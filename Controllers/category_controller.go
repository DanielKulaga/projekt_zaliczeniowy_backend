package Controllers

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"myapp/Database"
	"myapp/Models"
	"net/http"
)

func GetRestaurantCategories(c echo.Context) error {
	var categoryList []Models.Category
	fmt.Printf("Get category list\n")

	if result := Database.Database.Find(&categoryList); result.Error != nil {
		return c.String(http.StatusNotFound, "Database Error")
	}

	return c.JSON(http.StatusOK, categoryList)
}

func PostRestaurantCategories(c echo.Context) error {
	category := new(Models.Category)

	fmt.Printf("Add new category \n")

	if err := c.Bind(category); err != nil {
		return c.String(http.StatusBadRequest, "Bad user "+err.Error())
	}

	Database.Database.Create(category)

	return c.JSON(http.StatusOK, category)
}
