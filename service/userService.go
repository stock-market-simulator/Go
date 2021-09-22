package service

import (
	"github.com/stock-market-simulator/Go/db/table"
)

func (g *gormHandler) SaveUser() []*table.User {
	var user []*table.User
	return user
}

func (g *gormHandler) GetUserData() []*table.User {
	var user []*table.User
	result := g.db.Find(&user)
	if result.Error != nil {
		panic(result.Error)
	}

	return user
}
