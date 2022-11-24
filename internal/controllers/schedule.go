package controllers

import (
	"Are-you-free/internal/db"
	"Are-you-free/internal/models"
	"fmt"
	"net/http"
	"strconv"

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

	schedule_id := c.Param("id")
	strconv.Atoi(schedule_id)

	sqlStatement := "SELECT id, Year, Month, Day, StartHour, StartMinute, EndHour, EndMinute FROM schedule WHERE id = ? LIMIT 1"

	schedule := models.Schedule{}

	rows := con.QueryRow(sqlStatement, schedule_id)
	err2 := rows.Scan(&schedule.ID, &schedule.Year, &schedule.Month, &schedule.Day, &schedule.StartHour, &schedule.StartMinute, &schedule.EndHour, &schedule.EndMinute)
	if err2 != nil {
		fmt.Println(err2)
	}

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
