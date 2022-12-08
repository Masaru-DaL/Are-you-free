package route

import (
	"io"
	"net/http"
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
	e.Pre(models.MethodOverride)
	e.GET("/schedules", models.GetAllSchedules)
	e.GET("/schedule/:id", models.GetOneSchedule)

	return e
}

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func initTemplate(e *echo.Echo) {
	templateList, err := template.New("t").ParseGlob("internal/public/views/*.html")
	templateList.ParseGlob("internal/public/view/*.html")
	t := &Template{
		templates: template.Must(templateList, err),
	}
	e.Renderer = t
	e.Pre(models.MethodOverride)
}
