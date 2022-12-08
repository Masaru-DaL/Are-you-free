package main

import (
	"src/internal/route"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := route.InitRouting()
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

	e.Logger.Fatal(e.Start(":8080"))
}
