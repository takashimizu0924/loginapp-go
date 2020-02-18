package model

import (
	"fmt"
	"time"
)

type StockItem struct {
	ID            int    `json:"id"`
	ItemName      string `json:"itemname"`
	ItemQuantity  int    `json:"itemquantity"`
	ItemStockyard string `json:"itemstockyard"`
	CreatedAt     time.Time
}

type StockItems []StockItem

func AddStockItem(stockitem *StockItem) {
	db.Create(stockitem)
}

func FindAllIStockItems(s *StockItem) StockItems {
	var stockitems StockItems
	db.Where(s).Find(&stockitems)
	return stockitems
}

func FindStockItem(s *StockItem) StockItems {
	var stockitems StockItems
	db.Where(s).First(&stockitems)
	return stockitems
}

func DeleteStockItem(s *StockItem) error {
	if rows := db.Where(s).Delete(&StockItem{}).RowsAffected; rows == 0 {
		return fmt.Errorf("%v は削除できませんでした", s)
	}
	return nil
}

func UpdateStockItem(s *StockItem) error {
	rows := db.Model(s).Update(map[string]interface{}{
		"name": s.ItemName,
	}).RowsAffected
	if rows == 0 {
		return fmt.Errorf("%v はアップデートできませんでした")
	}
	return nil
}
