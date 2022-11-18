package main

import (
	"Are-you-free/internal/controllers"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	/*
		Logger: リクエスト単位のログを出力する
		Recover: 予期せぬpanicを起こしてもサーバを落とさない
		CORS: アクセスを許可するオリジン(デフォルト)とメソッドの設定
	*/
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusCreated, "Hello, World!!")
	})
	e.GET("/schedules", controllers.GetSchedules)

	e.Logger.Fatal(e.Start(":8080"))
}
