package route

import (
	"io"
	"net/http"
	"src/internal/auth"
	"src/internal/handler"
	"text/template"

	"github.com/labstack/echo/v4"
)

func InitRouting() *echo.Echo {
	e := echo.New()

	/* html/template非対応 */
	// schedulesを操作する
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusCreated, "Hello, World!!")
	})
	e.POST("/create/schedule", handler.PostSchedule)
	e.PUT("/put", handler.PutSchedule)
	e.DELETE("/schedule/delete/:id", handler.DeleteSchedule)

	// usersを操作する
	e.POST("/login", handler.SignIn)
	e.POST("/signup", handler.SignUp)
	e.PUT("/user/update", handler.UpdateUser)

	admin := e.Group("/admin")
	admin.GET("/user/:id", handler.GetUser)
	admin.GET("/users", handler.GetUsers)
	admin.PUT("/user/update", handler.UpdateUserByAdmin)
	admin.DELETE("/user/delete", handler.DeleteUserByAdmin)

	// html/template対応
	initTemplate(e)
	e.GET("/signup", handler.SignUpTemplate)
	e.GET("/login", handler.LoginTemplate)

	e.GET("/schedule/:id", handler.GetOneSchedule)

	authGroup := e.Group("/auth")
	authGroup.Use(auth.CheckCookieMiddleware)
	authGroup.GET("/schedules", handler.GetAllSchedules)

	return e
}

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func initTemplate(e *echo.Echo) {
	templateList, err := template.New("t").ParseGlob("internal/public/template/*.html")
	// templateList.ParseGlob("internal/public/schedule/*.html")
	t := &Template{
		templates: template.Must(templateList, err),
	}
	e.Renderer = t
	e.Pre(handler.MethodOverride)
}
