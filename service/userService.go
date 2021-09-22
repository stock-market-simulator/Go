package service

import (
	"github.com/stock-market-simulator/Go/db/table"
)

func (g *gormHandler) SaveUser(token string) *table.User {
	//var user []*table.User
	g.db.AutoMigrate(&table.User{}, &table.Bookmark{})
	user := &table.User{Token: token}
	err := g.db.Model(&table.User{}).Create(user).Error
	if err != nil {
		panic(err)
	}
	//g.db.Model(&user).Related(&user.Bookmark, "ID")

	// create

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
