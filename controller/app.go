package controller

import (
	"net/http"

	"github.com/labstack/echo"
	model "github.com/stock-market-simulator/Go/service"
)

// db 객체
type AppHandler struct {
	db model.DBHandler
}

func MakeHandler() *echo.Echo {
	e := echo.New()
	handler := &AppHandler{
		db: model.NewDBHandler("trading"),
	}
	// db 테스트
	e.GET("/test/db/read", handler.getDbTestHandler)
	e.GET("/test/db/create", handler.createDbTestHandler)

	// api 테스트
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World!")
	})

	return e
}
