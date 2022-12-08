package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func HandleSignUp(c echo.Context) error {
	return c.Render(http.StatusOK, "signup", "")
}

func HandleMyPage(c echo.Context) error {
	return c.Render(http.StatusOK, "mypage", "")
}
