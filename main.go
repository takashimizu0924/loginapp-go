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

	// e.POST("/login", handler.Login)
	e.POST("/signup", handler.Signup)
	e.POST("/product", handler.AddProduct)
	e.POST("/sales", handler.AddSales)
	e.GET("/product", handler.GetProduct)
	e.GET("/allproduct", handler.GetAllProduct)
	e.GET("/sales", handler.GetSales)
	e.DELETE("/product/:id", handler.DeleteProduct)
	e.DELETE("/sales/:id", handler.DeleteSales)
	// e.PUT("/product/:id", handler.UpdateProduct)
	// e.PUT("/sales/:id", handler.UpdateSales)

	e.Logger.Fatal(e.Start(":9090"))
}
