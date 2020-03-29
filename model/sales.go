package model

import (
	"fmt"
	"log"
	"time"
)

type Sales struct {
	ID            int    `json:"id"`
	ReceiptNumber int    `json:"receiptnumber"`
	Code          int    `json:"code"`
	CompletedDate string `json:"completeddate"`
	SectionName   string `json:"sectionname"`
	GestName      string `json:"gestname"`
	SalesName     string `json:"salesname"`
	SalesPrice    int    `json:"salesprice"`
	SalesQuantity int    `json:"salesquantity"`
	CreatedBy     string `json:"created_by"`
	CreatedAt     time.Time
}

type Period struct {
	StartDate string `json:"startdate"`
	EndDate   string `json:"enddate"`
}

type AllSales []Sales

func AddSales(sales *Sales) {
	db.Create(sales)
}

func FindSales(s *Sales) AllSales {
	var findsale AllSales
	db.Where(s).First(&findsale)
	return findsale
}
func FindAllSales(s *Sales) AllSales {
	var allsales AllSales
	db.Where(s).Find(&allsales)
	return allsales
}

func DeleteSales(s *Sales) error {
	if rows := db.Where(s).Delete(&Sales{}).RowsAffected; rows == 0 {
		return fmt.Errorf("売上 %v の削除はできませんでした", s)
	}
	return nil
}

func UpdateSales(s *Sales) error {
	fmt.Println("ok")
	if err := db.Model(s).Updates(Sales{
		ReceiptNumber: s.ReceiptNumber,
		Code:          s.Code,
		CompletedDate: s.CompletedDate,
		SectionName:   s.SectionName,
		GestName:      s.GestName,
		SalesName:     s.SalesName,
		SalesPrice:    s.SalesPrice,
		SalesQuantity: s.SalesQuantity,
		CreatedBy:     s.CreatedBy,
		// "created_at": s.Created,
	}); err != nil {
		log.Println(err)
	}
	return nil
}

func FindCompletedSales(s string, e string) AllSales {
	var allsales AllSales
	db.Where("completed_date BETWEEN ? AND ?", s, e).Find(&allsales)
	return allsales
}
