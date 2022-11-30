package main

import (
	"Are-you-free/internal/models"
	"io"
	"net/http"
	"text/template"

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
	initRouting(e)

	e.Logger.Fatal(e.Start(":8080"))
}

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func initRouting(e *echo.Echo) {
	// html/template非対応
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusCreated, "Hello, World!!")
	})
	e.POST("/create", models.PostSchedule)
	e.PUT("/put", models.PutSchedule)
	e.DELETE("/schedule/delete/:id", models.DeleteSchedule)

	// html/template対応
	initTemplate(e)
	e.Pre(models.MethodOverride)
	e.GET("/schedules", models.GetAllSchedules)
	e.GET("/schedule/:id", models.GetOneSchedule)
}

func initTemplate(e *echo.Echo) {
	templateList, err := template.New("t").ParseGlob("public/views/*.html")
	templateList.ParseGlob("public/view/*.html")
	t := &Template{
		templates: template.Must(templateList, err),
	}
	e.Renderer = t
	e.Pre(models.MethodOverride)
}
