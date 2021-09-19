package controller

import (
	"net/http"

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
	e.GET("/test/db/create", handler.createDbTestHandler)

	// api 테스트
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World!")
	})

	// 클라이언트 websocket 연결
	// e.GET("/ws", handler.handleConnections)
	// go handleMessages()

	return e
}
