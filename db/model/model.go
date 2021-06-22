package model

import "github.com/stock-market-simulator/Go/db/table"

type DBHandler interface {
	// 이 두 메소드가 gormHandler를 참조하고 있으므로 이 인터페이스로 gormHandler 구조체 사용 가능
	GetDbTest() []*table.User
	CreateDbTest() *table.User
}

func NewDBHandler(database string) DBHandler {
	return mysqlHandler(database)
}
