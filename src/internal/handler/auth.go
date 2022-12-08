package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func HandleMyPage(c echo.Context) error {
	return c.Render(http.StatusOK, "mypage", map[string]interface{}{
		"title": "MyPage",
	})
}

func HandleSignUp(c echo.Context) error {
	return c.Render(http.StatusOK, "signup", map[string]interface{}{
		"title": "Login / SignUp Form",
	})
}
