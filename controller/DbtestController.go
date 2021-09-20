package controller

import (
	"net/http"

	"github.com/labstack/echo"
)

func (a *AppHandler) getDbTestHandler(c echo.Context) error {
	user := a.db.GetDbTest()

	return c.JSON(http.StatusOK, user)
}
