package main

import (
	"Are-you-free/internal/controllers"
	"Are-you-free/internal/models"
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

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
	// e.GET("/schedules", controllers.GetSchedules)
	e.POST("/create", models.PostSchedule)
	e.PUT("/put", models.PutSchedule)
	e.DELETE("/schedule/delete/:id", models.DeleteSchedule)

	// template
	// t := &Template{
	// 	templates: template.Must(template.ParseGlob("public/views/*.html")),
	// }
	templateList, err := template.New("t").ParseGlob("public/views/*.html")
	templateList.ParseGlob("public/view/*.html")
	t := &Template{
		templates: template.Must(templateList, err),
	}
	e.Renderer = t
	e.Pre(controllers.MethodOverride)
	e.GET("/index/:id", controllers.GetOneSchedule)

	e.GET("/index", controllers.GetSchedules)
	// e.GET("/index", controllers.GetAllSchedules)
	// e.GET("/hello", controllers.GetSchedules)

	e.Logger.Fatal(e.Start(":8080"))
}
