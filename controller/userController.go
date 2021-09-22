package controller

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/stock-market-simulator/Go/controller/dto"
)

func (a *AppHandler) saveUserHandler(c echo.Context) error {
	userBody := new(dto.UserDto)
	if err := c.Bind(userBody); err != nil {
		panic(err)
	}
	user := a.db.SaveUser(userBody.Token)

	return c.JSON(http.StatusOK, user)
}

func (a *AppHandler) getUserHandler(c echo.Context) error {
	user := a.db.GetUserData()

	return c.JSON(http.StatusOK, user)
}
