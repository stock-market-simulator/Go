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

func (g *gormHandler) SaveBookmark(token string, name string) *table.Bookmark {
	// 토큰이 실제 존재하는지 유효성 검사 추가 필요
	user := &table.User{Token: token}
	// code는 db에서 name을 통해 알아낼 수 있다.
	bookmark := &table.Bookmark{Code: "000", Name: name}

	err := g.db.Model(user).Association("Bookmark").Append(bookmark)
	if err != nil {
		panic(err)
	}

	return bookmark
}

func (g *gormHandler) GetUserBookmarkData(token string) []*table.User {
	var user []*table.User
	result := g.db.Find(&user)
	if result.Error != nil {
		panic(result.Error)
	}

	return user
}
