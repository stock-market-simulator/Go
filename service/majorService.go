package service

import (
	"github.com/stock-market-simulator/Go/db/table"
)

type majorList struct {
	Kospi  []*table.Kospi  `json:"kospi"`
	Kosdaq []*table.Kosdaq `json:"kosdaq"`
}

func (g *gormHandler) GetMajorData() majorList {
	var kospi []*table.Kospi
	var kosdaq []*table.Kosdaq

	kospiResult := g.db.Find(&kospi)
	if kospiResult.Error != nil {
		panic(kospiResult.Error)
	}

	kosdaqResult := g.db.Find(&kosdaq)
	if kosdaqResult.Error != nil {
		panic(kosdaqResult.Error)
	}

	result := majorList{kospi, kosdaq}

	return result
}
