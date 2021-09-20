package service

import (
	"github.com/stock-market-simulator/Go/db"
	"github.com/stock-market-simulator/Go/db/table"
	"gorm.io/gorm"
)

// 메소드들이 gormHandler를 리시버로 가지고 있으므로 이 인터페이스도 gormHandler 구조체를 암시적으로 사용 가능
type DBHandler interface {
	GetMajorData() []*table.Kospi
	GetDbTest() []*table.User
}

// gormHandler db에 직접 접근하므로 소문자로 사용해 private으로 보호하고 외부에서는 DBHandler를 이용해 gormHandler를 리시버로 가지고 있는 메소드에 접근(약간 getter 느낌??)
type gormHandler struct {
	db *gorm.DB
}

func NewDBHandler(database string) DBHandler {
	db, err := db.Connect(database)

	if err != nil {
		panic(err)
	}

	return &gormHandler{db: db}
}
