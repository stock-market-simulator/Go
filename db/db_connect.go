package db

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Info struct {
	User     string
	Password string
	Protocol string
	Endpoint string
	Port     string
}

func Connect(database string) (*gorm.DB, error) {
	info := Info{User: os.Getenv("TRADING_USER"), Password: os.Getenv("TRADING_PASSWORD"), Protocol: "tcp", Endpoint: os.Getenv("TRADING_ENDPOINT"), Port: "3306"}

	connect_info := fmt.Sprintf("%s:%s@%s(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		info.User, info.Password, info.Protocol, info.Endpoint, info.Port, database,
	)

	return gorm.Open(mysql.Open(connect_info), &gorm.Config{})
}
