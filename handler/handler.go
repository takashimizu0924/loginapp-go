package handler

import (
	"log"
	"net/http"
	"strconv"

	"../model"
	"github.com/labstack/echo"
)

func AddProduct(c echo.Context) error {
	product := new(model.Product)
	if err := c.Bind(product); err != nil {
		return err
	}
	if product.Name == "" {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "商品名が不正です",
		}
	}
	model.AddProduct(product)
	return c.JSON(http.StatusCreated, product)
}

func AddSales(c echo.Context) error {
	sales := new(model.Sales)
	if err := c.Bind(sales); err != nil {
		return err
	}
	if sales.Name == "" {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "売上名が不正です",
		}
	}
	model.AddSales(sales)
	return c.JSON(http.StatusCreated, sales)
}

func GetAllProduct(c echo.Context) error {
	user := c.QueryParam("created_by")
	// if err != nil {
	// 	return echo.ErrNotFound
	// }
	product := model.FindAllProduct(&model.Product{Created_by: user})
	return c.JSON(http.StatusOK, product)
}

func GetProduct(c echo.Context) error {
	id, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		return echo.ErrNotFound
	}
	product := model.FindProduct(&model.Product{ID: id})
	return c.JSON(http.StatusOK, product)
}

func GetSales(c echo.Context) error {
	id, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		return echo.ErrNotFound
	}
	sales := model.FindSales(&model.Sales{ID: id})
	return c.JSON(http.StatusOK, sales)
}

func DeleteProduct(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Fatal("first")
		return echo.ErrNotFound
	}

	if err := model.DeleteProduct(&model.Product{ID: id}); err != nil {
		log.Fatal("here", err)
		return echo.ErrNotFound
	}
	return c.NoContent(http.StatusNoContent)
}

func DeleteSales(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.ErrNotFound
	}
	if err := model.DeleteSales(&model.Sales{ID: id}); err != nil {
		return echo.ErrNotFound
	}
	return c.NoContent(http.StatusNoContent)
}

// func UpdateProduct(c echo.Context) error {
// 	id,err := strconv.Atoi(c.Param("id"))
// 	if err != nil {
// 		return echo.HTTPError
// 	}
// 	if err := model.UpdateProduct(&model.Product{ID: id});err != nil {
// 		return echo.HTTPError
// 	}
// 	return c.JSON(http.status)
// }

// func UpdateSales(c echo.Context) error {

// }
