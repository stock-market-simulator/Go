package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/stock-market-simulator/Go/db"
)

type User struct {
	ID       uint   //기본적으로 ID가 primaryKey
	Name     string `json:"Name" gorm:"column:Name"`
	Password string `json:"Password" gorm:"column:Password"`
}

// 테이블 이름을 따로 설정해주지 않으면 default로 구조체이름을 소문자로 만든 후 s가 붙은 테이블을 참조함
// ex) User구조체면 users의 테이블을 참조함. 따라서 밑에 TableName처럼 테이블 이름을 설정해야함
func (User) TableName() string {
	return "User"
}

func main() {
	e := echo.New()

	// db 테스트
	/*
		e.GET("/test/db/create", func(c echo.Context) error {
			db, err := db.Connect("test")
			if err != nil {
				return err
			}

			// db 마이그레이션(User란 테이블이 없으면 생성)
			db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&(User{}))
			u := User{
				Name:     "test",
				Password: "성공",
			}
			// User 테이블에 데이터 추가
			db.Create(&u)

			return c.JSON(http.StatusOK, u)
		})
	*/
	e.GET("/test/db/read", func(c echo.Context) error {
		db, err := db.Connect("test")
		if err != nil {
			return err
		}

		// User 테이블에 있는 데이터 가져오기
		var user []User
		result := db.Find(&user)
		if result.Error != nil {
			return result.Error
		}

		return c.JSON(http.StatusOK, user)
	})

	e.Logger.Fatal(e.Start(":5000"))
}
