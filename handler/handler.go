package handler

import (
	"fmt"
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
	product := c.QueryParam("productsection")
	productname := model.FindAllProduct(&model.Product{Productsection: product})
	return c.JSON(http.StatusOK, productname)
}

func GetProductCode(c echo.Context) error {
	code, _ := strconv.Atoi(c.QueryParam("productcode"))

	product := model.FindAllProduct(&model.Product{ProductCode: code})
	return c.JSON(http.StatusOK, product)
}

func GetProductName(c echo.Context) error {
	product := c.QueryParam("productname")
	productname := model.FindAllProduct(&model.Product{Name: product})
	return c.JSON(http.StatusOK, productname)
}

func GetAllSales(c echo.Context) error {
	createdby := c.QueryParam("created_by")
	sales := model.FindAllSales(&model.Sales{CreatedBy: createdby})
	return c.JSON(http.StatusOK, sales)
}

func GetSalesCode(c echo.Context) error {
	code, _ := strconv.Atoi(c.QueryParam("code"))

	sales := model.FindAllSales(&model.Sales{Code: code})
	return c.JSON(http.StatusOK, sales)
}

func GetSales(c echo.Context) error {
	salesprice, _ := strconv.Atoi(c.QueryParam("salesprice"))

	sales := model.FindSales(&model.Sales{SalesPrice: salesprice})
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

func DeleteStockItem(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Fatal("first")
		return echo.ErrNotFound
	}

	if err := model.DeleteStockItem(&model.StockItem{ID: id}); err != nil {
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

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.ErrNotFound
	}

	if err := model.FindProduct(&model.Product{ID: id}); err != nil {
		return echo.ErrNotFound
	}

	// if err := model.UpdateProduct(update); err != nil {
	// 	return echo.ErrNotFound
	// }

	return c.NoContent(http.StatusNoContent)
}

func UpdateSales(c echo.Context) error {
	update := new(model.Sales)
	if err := c.Bind(update); err != nil {
		return echo.ErrNotFound
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.ErrNotFound
	}

	salesdata := model.FindSales(&model.Sales{ID: id})

	updateData := salesdata[0]
	updateData.ReceiptNumber = update.ReceiptNumber
	updateData.Code = update.Code
	updateData.CompletedDate = update.CompletedDate
	updateData.SectionName = update.SectionName
	updateData.GestName = update.GestName
	updateData.SalesName = update.SalesName
	updateData.SalesPrice = update.SalesPrice
	updateData.SalesQuantity = update.SalesQuantity
	updateData.CreatedBy = update.CreatedBy
	fmt.Println(salesdata)
	if err := model.UpdateSales(&updateData); err != nil {
		return echo.ErrNotFound
	}

	return c.NoContent(http.StatusNoContent)
}

func UpdateItem(c echo.Context) error {

	updateData := new(model.StockItem)
	if err := c.Bind(updateData); err != nil {
		return err
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.ErrNotFound
	}

	itemdata := model.FindStockItem(&model.StockItem{ID: id})
	if len(itemdata) == 0 {
		return echo.ErrNotFound
	}
	item := itemdata[0]
	item.ItemName = updateData.ItemName
	item.ItemQuantity = updateData.ItemQuantity
	item.ItemStockyard = updateData.ItemStockyard
	item.UpdatedAt = updateData.UpdatedAt
	item.CreatedBy = updateData.CreatedBy
	fmt.Println(item)
	if err := model.UpdateStockItem(&item); err != nil {
		return echo.ErrNotFound
	}

	return c.NoContent(http.StatusNoContent)
}

func FindPeriodSales(c echo.Context) error {
	period := new(model.Period)
	if err := c.Bind(period); err != nil {
		return err
	}

	list := model.FindCompletedSales(period.StartDate, period.EndDate)
	fmt.Println(list, len(list))
	return c.JSON(http.StatusOK, list)
}
