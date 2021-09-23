package service

import (
	"time"

	"github.com/stock-market-simulator/Go/controller/dto"
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
	bookmark := &table.Bookmark{Code: "003545", Name: name, UserID: user.UserID}

	err := g.db.Debug().Model(user).Association("Bookmarks").Append(bookmark)
	if err != nil {
		panic(err)
	}

	return bookmark
}

func (g *gormHandler) GetUserBookmarkData(token string) []dto.BookmarkResponseDto {
	user := table.User{}
	g.db.Debug().Where("Token=?", token).Preload("Bookmarks").Find(&user)

	now := time.Now()
	year := now.Year() - 1
	month := int(now.Month())
	day := now.Day()
	current := convert(year, month, day, "")
	previous := convert(year, month, day-1, "")

	var res []dto.BookmarkResponseDto
	for _, v := range user.Bookmarks {
		var bookmarkInfo dto.BookmarkResponseDto
		bookmarkInfo.Name = v.Name

		stockInfo := table.Stock{}

		g.db.Debug().Table("stock_"+v.Code).Where("data=?", current).Find(&stockInfo)
		bookmarkInfo.CurrentPrice = stockInfo.MarketPrice

		g.db.Debug().Table("stock_"+v.Code).Where("data=?", previous).Find(&stockInfo)
		bookmarkInfo.PreviousPrice = stockInfo.MarketPrice

		res = append(res, bookmarkInfo)
	}

	return res
}
