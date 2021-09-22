package controller

import (
	"net/http"

	"github.com/labstack/echo"
)

func (a *AppHandler) saveUserHandler(c echo.Context) error {
	user := a.db.SaveUser()

	return c.JSON(http.StatusOK, user)
}

func (a *AppHandler) getUserHandler(c echo.Context) error {
	user := a.db.GetUserData()

	return c.JSON(http.StatusOK, user)
}
