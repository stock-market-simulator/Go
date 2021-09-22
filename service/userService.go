package service

import (
	"github.com/stock-market-simulator/Go/db/table"
)

func (g *gormHandler) SaveUser(token string) *table.User {
	user := &table.User{Token: token}

	g.db.AutoMigrate(&table.User{}, &table.Bookmark{})
	err := g.db.Debug().Model(&table.User{}).Create(user).Error
	if err != nil {
		panic(err)
	}

	return user
}

func (g *gormHandler) SaveBookmark(token string, name string) *table.Bookmark {
	var user *table.User
	g.db.Debug().Where("Token=?", token).Find(&user)

	// code는 db에서 name을 통해 알아낼 수 있다.
	bookmark := &table.Bookmark{Code: "000", Name: name, UserID: user.UserID}

	err := g.db.Debug().Model(user).Association("Bookmarks").Append(bookmark)
	if err != nil {
		panic(err)
	}

	return bookmark
}

func (g *gormHandler) GetUserBookmarkData(token string) []table.Bookmark {
	result := table.User{}
	g.db.Debug().Where("Token=?", token).Preload("Bookmarks").Find(&result)

	return result.Bookmarks
}
