package controllers

import (
	"Are-you-free/internal/db"
	"Are-you-free/internal/models"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

/* メソッドはmain.goのハンドラとして使用する */

/* 全てのスケジュールを取得する */
// func GetSchedules(c echo.Context) error {
// 	// modelsに定義された関数を実行する
// 	result := models.GetSchedule()
// 	println("Get All Schedules")
// 	return c.Render(http.StatusOK, "hello", result)
// }

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

func GetOneSchedule(c echo.Context) error {
	con := db.CreateConnection()
	sqlStatement := "SELECT id, Year, Month, Day, StartHour, StartMinute, EndHour, EndMinute FROM schedule order by id"
	rows, err := con.Query(sqlStatement)
	if err != nil {
		fmt.Println(err)
	}
	schedule := models.Schedule{}
	defer rows.Close()

	return c.Render(http.StatusOK, "hello", map[string]interface{}{
		"title":       "Get Schedule",
		"id":          schedule.ID,
		"year":        schedule.Year,
		"month":       schedule.Month,
		"day":         schedule.Day,
		"starthour":   schedule.StartHour,
		"startminute": schedule.StartMinute,
		"endhour":     schedule.EndHour,
		"endminute":   schedule.EndMinute,
	})
}

// template
// func Hello(c echo.Context) error {
// 	data := c.QueryParam("id")
// 	return c.Render(http.StatusOK, "hello", data)
// }
