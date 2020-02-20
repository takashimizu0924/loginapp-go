package model

import (
	"fmt"
	"time"
)

type Sales struct {
	ID            int    `json:"id"`
	Code          int    `json:"code"`
	SalesName     string `json:"salesname"`
	SalesPrice    int    `json:"salesprice"`
	SalesQuantity int    `json:"saleequantity"`
	Created       time.Time
}

type AllSales []Sales

func AddSales(sales *Sales) {
	db.Create(sales)
}

func FindSales(s *Sales) AllSales {
	var allsales AllSales
	db.Where(s).First(&allsales)
	return allsales
}

func DeleteSales(s *Sales) error {
	if rows := db.Where(s).Delete(&Sales{}).RowsAffected; rows == 0 {
		return fmt.Errorf("売上 %v の削除はできませんでした", s)
	}
	return nil
}

func UpdateSales(s *Sales) error {
	rows := db.Where(s).Update(map[string]interface{}{
		"name":       s.SalesName,
		"price":      s.SalesPrice,
		"created_at": s.Created,
	}).RowsAffected
	if rows == 0 {
		return fmt.Errorf("%v は更新できませんでした", s)
	}
	return nil
}
