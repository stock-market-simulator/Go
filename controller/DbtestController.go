package controller

import (
	"net/http"

	"github.com/labstack/echo"
)

func (a *AppHandler) getDbTestHandler(c echo.Context) error {
	user := a.db.GetDbTest()

	return c.JSON(http.StatusOK, user)
}

func (a *AppHandler) createDbTestHandler(c echo.Context) error {
	u := a.db.CreateDbTest()

	return c.JSON(http.StatusOK, u)
}
