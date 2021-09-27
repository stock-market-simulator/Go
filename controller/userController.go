package controller

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/stock-market-simulator/Go/controller/dto"
)

func (a *AppHandler) saveUserHandler(c echo.Context) error {
	body := new(dto.UserDto)
	if err := c.Bind(body); err != nil {
		panic(err)
	}
	user, check := a.db.SaveUser(body.Token)

	if !check {
		return c.JSON(http.StatusInternalServerError, "이미 존재하는 토큰입니다.")
	}

	return c.JSON(http.StatusOK, user)
}

func (a *AppHandler) saveBookmark(c echo.Context) error {
	body := new(dto.BookmarkRequestDto)
	if err := c.Bind(body); err != nil {
		panic(err)
	}

	bookmark := a.db.SaveBookmark(body.Token, body.Name)

	return c.JSON(http.StatusOK, bookmark)
}

func (a *AppHandler) getUserBookmarkHandler(c echo.Context) error {
	token := c.Param("token")
	user := a.db.GetUserBookmarkData(token)

	return c.JSON(http.StatusOK, user)
}
