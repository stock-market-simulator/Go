package service

import (
	"time"

	"github.com/stock-market-simulator/Go/controller/dto"
	"github.com/stock-market-simulator/Go/db/table"
)

func (g *gormHandler) SaveUser(token string) (*table.User, bool) {
	user := &table.User{Token: token}
	var checkUser *table.User
	g.db.Debug().Where("Token=?", token).Find(&checkUser)

	if checkUser.UserID != 0 {
		return checkUser, false
	}
	g.db.AutoMigrate(&table.User{}, &table.Bookmark{})
	err := g.db.Debug().Model(&table.User{}).Create(user).Error
	if err != nil {
		panic(err)
	}

	return user, true
}

func (g *gormHandler) SaveBookmark(token string, name string) *table.Bookmark {
	var user *table.User
	g.db.Debug().Where("Token=?", token).Find(&user)

	stockInfo := table.StockInfo{}
	g.db.Debug().Where("name=?", name).Find(&stockInfo)

	bookmark := &table.Bookmark{Code: stockInfo.Code, Name: name, UserID: user.UserID}

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

		currentStock := table.Stock{}
		previousStock := table.Stock{}

		g.db.Debug().Table("stock_"+v.Code).Where("data=?", current).Find(&currentStock)
		bookmarkInfo.CurrentPrice = currentStock.MarketPrice

		g.db.Debug().Table("stock_"+v.Code).Where("data=?", previous).Find(&previousStock)
		bookmarkInfo.PreviousPrice = previousStock.MarketPrice

		res = append(res, bookmarkInfo)
	}

	return res
}
