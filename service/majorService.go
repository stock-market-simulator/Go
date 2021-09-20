package service

import (
	"github.com/stock-market-simulator/Go/db/table"
)

func (g *gormHandler) GetMajorData() []*table.Kospi {
	var kospi []*table.Kospi
	result := g.db.Find(&kospi)
	if result.Error != nil {
		panic(result.Error)
	}

	return kospi
}
