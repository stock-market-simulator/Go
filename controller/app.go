package controller

import (
	"github.com/labstack/echo"
	model "github.com/stock-market-simulator/Go/service"
)

// db 객체
// AppHandler를 리시버로 가지고 있는 메소드들에 접근 가능
type AppHandler struct {
	db model.DBHandler
}

func Controller() *echo.Echo {
	e := echo.New()

	// db 연결
	handler := &AppHandler{
		db: model.NewDBHandler("trading"), // DBHandler의 메소드를 가지고 있음
	}

	// db 테스트
	e.GET("/test/db/read", handler.getDbTestHandler)

	// 코스피 코스닥 지수
	e.GET("/major", handler.getMajorHandler)

	// 클라이언트 websocket 연결 부분
	// e.GET("/ws", handler.handleConnections)
	// go handleMessages()

	return e
}
