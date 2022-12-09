package route

import (
	"io"
	"net/http"
	"src/internal/handler"
	"src/internal/models"
	"text/template"

	"github.com/labstack/echo/v4"
)

func InitRouting() *echo.Echo {
	e := echo.New()

	// html/template非対応
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusCreated, "Hello, World!!")
	})
	e.POST("/create", models.PostSchedule)
	e.PUT("/put", models.PutSchedule)
	e.DELETE("/schedule/delete/:id", models.DeleteSchedule)

	// html/template対応
	initTemplate(e)
	e.GET("/schedules", models.GetAllSchedules)
	e.GET("/schedule/:id", models.GetOneSchedule)
	e.GET("/signup", handler.HandleSignUp)
	e.GET("/mypage", handler.HandleMyPage)

	return e
}

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func initTemplate(e *echo.Echo) {
	templateList, err := template.New("t").ParseGlob("internal/public/auth/*.html")
	templateList.ParseGlob("internal/public/schedule/*.html")
	templateList.ParseGlob("internal/public/schedules/*.html")
	t := &Template{
		templates: template.Must(templateList, err),
	}
	e.Renderer = t
	e.Pre(MethodOverride)
}

/* PUTやDELETEにも対応させるメソッド */
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
