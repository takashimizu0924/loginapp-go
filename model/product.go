package model

import (
	"fmt"
	"time"
)

type Product struct {
	ID             int    `json:"id"`
	ProductCode    int    `json:"productcode"`
	Productsection string `json:"productsection"`
	Name           string `json:"productname"`
	CreatedAt      time.Time
}

type Products []Product

func AddProduct(product *Product) {
	db.Create(product)
}

func FindAllProduct(p *Product) Products {
	var products Products
	db.Where(p).Find(&products)
	return products
}

func FindProduct(p *Product) Products {
	var products Products
	db.Where(p).First(&products)
	return products
}

func DeleteProduct(p *Product) error {
	if rows := db.Where(p).Delete(&Product{}).RowsAffected; rows == 0 {
		return fmt.Errorf("%v は削除できませんでした", p)
	}
	return nil
}

func UpdateProduct(p *Product) error {
	rows := db.Model(p).Update(map[string]interface{}{
		"name": p.Name,
	}).RowsAffected
	if rows == 0 {
		return fmt.Errorf("%v はアップデートできませんでした")
	}
	return nil
}
