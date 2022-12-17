package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

/* htmlのformにもPUTやDELETEにも対応させるメソッド */
func MethodOverride(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if c.Request().Method == "POST" {
			method := c.Request().PostFormValue("_method")
			if method == "PUT" || method == "PATCH" || method == "DELETE" {
				c.Request().Method = method
			}
		}
		return next(c)
	}
}

func LoginTemplate(c echo.Context) error {
	return c.Render(http.StatusOK, "login", "")
}

func SignUpTemplate(c echo.Context) error {
	return c.Render(http.StatusOK, "signup", "")
}
