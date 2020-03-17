package model

import (
	"fmt"
	"log"
	"time"
)

type StockItem struct {
	ID            int    `json:"id"`
	ItemName      string `json:"itemname"`
	ItemQuantity  int    `json:"itemquantity"`
	ItemStockyard string `json:"itemstockyard"`
	CreatedAt     time.Time
	UpdatedAt     string `json:"updated_at"`
	CreatedBy     string `json:"created_by"`
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
	fmt.Println("nice")
	// if err := db.Model(s).Updates(map[string]interface{}{
	// 	"itemname": s.ItemName,
	// }); err != nil {
	// 	log.Println(err)
	// }
	if err := db.Model(s).Updates(StockItem{
		ItemName:      s.ItemName,
		ItemQuantity:  s.ItemQuantity,
		ItemStockyard: s.ItemStockyard,
		UpdatedAt:     s.UpdatedAt,
		CreatedBy:     s.CreatedBy,
	}); err != nil {
		log.Println(err)
	}
	// if rows == 0 {
	// 	return fmt.Errorf("%v はアップデートできませんでした")
	// }
	return nil
}
