package service

import (
	"github.com/stock-market-simulator/Go/db/table"
)

func (g *gormHandler) SaveUser(token string) *table.User {
	user := &table.User{Token: token}

	g.db.AutoMigrate(&table.User{}, &table.Bookmark{})
	err := g.db.Model(&table.User{}).Create(user).Error
	if err != nil {
		panic(err)
	}

	return user
}

func (g *gormHandler) GetUserBookmarkData(token string) []*table.User {
	var user []*table.User
	result := g.db.Find(&user)
	if result.Error != nil {
		panic(result.Error)
	}

	return user
}
