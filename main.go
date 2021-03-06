package main

import (
	"./handler"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// e.POST("/login", handler.Login)
	e.POST("/signup", handler.Signup)
	e.POST("/product", handler.AddProduct)
	e.POST("/sales", handler.AddSales)
	e.POST("/stockitem", handler.AddStockItem)
	e.GET("/product", handler.GetProduct)
	e.GET("/allproduct", handler.GetAllProduct)
	e.GET("/productname", handler.GetProductName)
	e.GET("/productcode", handler.GetProductCode)
	e.GET("/sales", handler.GetSales)
	e.GET("salescode", handler.GetSalesCode)
	e.GET("allsales", handler.GetAllSales)
	e.GET("/allstockitems", handler.GetAllStockItem)
	e.GET("/stockitem", handler.GetStockItem)
	e.GET("/periodsales", handler.FindPeriodSales)
	e.DELETE("/product/:id", handler.DeleteProduct)
	e.DELETE("/stockitem/:id", handler.DeleteStockItem)
	e.DELETE("/sales/:id", handler.DeleteSales)
	// e.PUT("/product/:id", handler.UpdateProduct)
	// e.PUT("/sales/:id", handler.UpdateSales)
	e.PATCH("/updateitem/:id", handler.UpdateItem)
	e.PATCH("/updatesales/:id", handler.UpdateSales)
	e.Logger.Fatal(e.Start(":9090"))
}
