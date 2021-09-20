package controller

import (
	"net/http"

	"github.com/labstack/echo"
)

func (a *AppHandler) getMajorHandler(c echo.Context) error {
	user := a.db.GetMajorData()

	return c.JSON(http.StatusOK, user)
}
