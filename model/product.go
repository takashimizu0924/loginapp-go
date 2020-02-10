package model

import (
	"fmt"
)

type Product struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Quantity   int    `json:"quantity"`
	Created_by string `json:"created_by`
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
		"name":     p.Name,
		"quantity": p.Quantity,
	}).RowsAffected
	if rows == 0 {
		return fmt.Errorf("%v はアップデートできませんでした")
	}
	return nil
}
