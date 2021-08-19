package service

import (
	"github.com/stock-market-simulator/Go/db/table"
)

func (g *gormHandler) GetDbTest() []*table.User {
	var user []*table.User
	result := g.db.Find(&user)
	if result.Error != nil {
		panic(result.Error)
	}

	return user
}

func (g *gormHandler) CreateDbTest() *table.User {
	// db 마이그레이션(User란 테이블이 없으면 생성)
	g.db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&(table.User{}))
	u := table.User{
		Name:     "test",
		Password: "success",
	}
	// User 테이블에 데이터 추가
	g.db.Create(&u)

	return &u
}
