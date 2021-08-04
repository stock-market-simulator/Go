package app

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/stock-market-simulator/Go/db/model"
)

// db 테스트
type AppHandler struct {
	db model.DBHandler
}

func (a *AppHandler) getDbTestHandler(c echo.Context) error {
	user := a.db.GetDbTest()

	return c.JSON(http.StatusOK, user)
}

func (a *AppHandler) createDbTestHandler(c echo.Context) error {
	u := a.db.CreateDbTest()

	return c.JSON(http.StatusOK, u)
}

func MakeHandler() *echo.Echo {
	e := echo.New()
	handler := &AppHandler{
		db: model.NewDBHandler("test"),
	}
	// db 테스트
	e.GET("/test/db/read", handler.getDbTestHandler)
	e.GET("/test/db/create", handler.createDbTestHandler)

	return e
}
