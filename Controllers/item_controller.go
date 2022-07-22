package Controllers

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"myapp/Database"
	"myapp/Models"
	"net/http"
)

func GetItems(c echo.Context) error {
	var items []Models.Item
	fmt.Printf("Get items\n")

	if result := Database.Database.Find(&items); result.Error != nil {
		return c.String(http.StatusNotFound, "Get all products Database Error")
	}

	return c.JSON(http.StatusOK, items)
}

func GetItemsInCategory(c echo.Context) error {
	var products []Models.Item
	id := c.Param("id")
	fmt.Printf("Get products with category id " + id + "\n")

	if result := Database.Database.Where("category_id", id).Find(&products); result.Error != nil {
		return c.String(http.StatusNotFound, "Get Products In Category Database Error")
	}

	return c.JSON(http.StatusOK, products)
}

func PostItems(c echo.Context) error {
	product := new(Models.Item)

	fmt.Printf("Add new product \n")

	if err := c.Bind(product); err != nil {
		return c.String(http.StatusBadRequest, "Bad product"+err.Error())
	}

	Database.Database.Create(product)

	return c.JSON(http.StatusOK, product)
}
