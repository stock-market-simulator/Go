package service

import (
	"time"

	"github.com/stock-market-simulator/Go/db/table"
)

type majorList struct {
	Kospi  []*table.Kospi  `json:"kospi"`
	Kosdaq []*table.Kosdaq `json:"kosdaq"`
}

func (g *gormHandler) GetMajorData() majorList {
	now := time.Now()
	year := now.Year() - 1
	month := int(now.Month())
	day := now.Day()
	start := convert(year, month-2, day, "/")
	end := convert(year, month, day, "/")

	var kospi []*table.Kospi
	var kosdaq []*table.Kosdaq

	kospiResult := g.db.Where("date BETWEEN ? AND ?", start, end).Find(&kospi)
	if kospiResult.Error != nil {
		panic(kospiResult.Error)
	}

	kosdaqResult := g.db.Where("date BETWEEN ? AND ?", start, end).Find(&kosdaq)
	if kosdaqResult.Error != nil {
		panic(kosdaqResult.Error)
	}

	result := majorList{kospi, kosdaq}

	return result
}
