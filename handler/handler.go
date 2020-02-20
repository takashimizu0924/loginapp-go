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
	// if err := time.Parse("2020-2-17", product.CreatedAt); err != nil {
	// 	return err
	// }
	model.AddProduct(product)
	return c.JSON(http.StatusCreated, product)
}

func AddSales(c echo.Context) error {
	sales := new(model.Sales)
	if err := c.Bind(sales); err != nil {
		return err
	}
	if sales.SalesName == "" {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "売上名が不正です",
		}
	}
	model.AddSales(sales)
	return c.JSON(http.StatusCreated, sales)
}

func AddStockItem(c echo.Context) error {
	item := new(model.StockItem)
	if err := c.Bind(item); err != nil {
		return err
	}
	if item.ItemName == "" {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "売上名が不正です",
		}
	}
	model.AddStockItem(item)
	return c.JSON(http.StatusCreated, item)
}

func GetAllProduct(c echo.Context) error {
	user, _ := strconv.Atoi(c.QueryParam("id"))

	product := model.FindAllProduct(&model.Product{ID: user})
	return c.JSON(http.StatusOK, product)
}

func GetProduct(c echo.Context) error {
	code, err := strconv.Atoi(c.QueryParam("productcode"))
	if err != nil {
		return echo.ErrNotFound
	}
	product := model.FindProduct(&model.Product{ProductCode: code})
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

func GetAllStockItem(c echo.Context) error {
	id, _ := strconv.Atoi(c.QueryParam("id"))

	item := model.FindAllIStockItems(&model.StockItem{ID: id})
	return c.JSON(http.StatusOK, item)
}

func GetStockItem(c echo.Context) error {
	id, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		return echo.ErrNotFound
	}
	item := model.FindStockItem(&model.StockItem{ID: id})
	return c.JSON(http.StatusOK, item)
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

func UpdateProduct(c echo.Context) error {
	update := new(model.Product)
	if err := c.Bind(update); err != nil {
		return echo.ErrNotFound
	}
	if len(update.Name) == 0 {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "入力が不正です",
		}
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.ErrNotFound
	}

	if err := model.FindProduct(&model.Product{ID: id}); err != nil {
		return echo.ErrNotFound
	}

	if err := model.UpdateProduct(update); err != nil {
		return echo.ErrNotFound
	}

	return c.NoContent(http.StatusNoContent)
}

func UpdateSales(c echo.Context) error {
	update := new(model.Sales)
	if err := c.Bind(update); err != nil {
		return echo.ErrNotFound
	}
	if len(update.SalesName) == 0 {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "入力が不正です",
		}
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.ErrNotFound
	}

	if err := model.FindSales(&model.Sales{ID: id}); err != nil {
		return echo.ErrNotFound
	}

	if err := model.UpdateSales(update); err != nil {
		return echo.ErrNotFound
	}

	return c.NoContent(http.StatusNoContent)
}
