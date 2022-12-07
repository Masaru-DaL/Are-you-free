package models

import (
	"Are-you-free/internal/db"
	"fmt"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
)

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

type Schedule struct {
	ID          int `json:"id"`
	Year        int `json:"year"`
	Month       int `json:"month"`
	Day         int `Json:"day"`
	StartHour   int `json:"starthour"`
	StartMinute int `json:"startminute"`
	EndHour     int `json:"endhour"`
	EndMinute   int `json:"endminute"`
}

type Schedules struct {
	Schedules []Schedule `json:"Schedule"`
}

/* 1件取得の関数 */
func GetOneSchedule(c echo.Context) error {
	con := db.CreateConnection()

	schedule_id := c.Param("id")
	strconv.Atoi(schedule_id)

	sqlStatement := "SELECT id, Year, Month, Day, StartHour, StartMinute, EndHour, EndMinute FROM schedule WHERE id = ? LIMIT 1"

	schedule := Schedule{}

	rows := con.QueryRow(sqlStatement, schedule_id)
	err2 := rows.Scan(&schedule.ID, &schedule.Year, &schedule.Month, &schedule.Day, &schedule.StartHour, &schedule.StartMinute, &schedule.EndHour, &schedule.EndMinute)
	if err2 != nil {
		fmt.Println(err2)
	}

	return c.Render(http.StatusOK, "schedule", map[string]interface{}{
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

/* 全件取得の関数 */
func GetAllSchedules(c echo.Context) error {
	con := db.CreateConnection()

	sqlStatement := "SELECT ID, Year, Month, Day, StartHour, StartMinute, EndHour, EndMinute FROM schedule"

	// sqlStatement := "SELECT ID, Year, Month, Day, StartHour, StartMinute, EndHour, EndMinute FROM schedule where id = ?"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		fmt.Println(err)
	}

	// rows, err := stmt.Query(1)
	rows, err := stmt.Query()
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	schedules := Schedules{}
	for rows.Next() {
		schedule := Schedule{}
		err := rows.Scan(&schedule.ID, &schedule.Year, &schedule.Month, &schedule.Day, &schedule.StartHour, &schedule.StartMinute, &schedule.EndHour, &schedule.EndMinute)

		if err != nil {
			fmt.Println(err)
		}
		schedules.Schedules = append(schedules.Schedules, schedule)
	}
	return c.Render(http.StatusOK, "schedules", schedules.Schedules)
}

/* POSTリクエスト */
func PostSchedule(c echo.Context) error {
	con := db.CreateConnection()
	sch := new(Schedule)
	if err := c.Bind(sch); err != nil {
		return err
	}

	sqlStatement := "INSERT INTO schedule(id, year, month, day, starthour, startminute, endhour, endminute) VALUES(?, ?, ?, ?, ?, ?, ?, ?)"
	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		fmt.Print(err.Error())
	}
	defer stmt.Close()

	result, err2 := stmt.Exec(sch.ID, sch.Year, sch.Month, sch.Day, sch.StartHour, sch.StartMinute, sch.EndHour, sch.EndMinute)

	// エラーが発生したら終了する
	if err2 != nil {
		panic(err2)
	}
	fmt.Println(result.LastInsertId())

	return c.JSON(http.StatusCreated, sch.ID)
}

func PutSchedule(c echo.Context) error {
	con := db.CreateConnection()
	sch := new(Schedule)
	if err := c.Bind(sch); err != nil {
		return err
	}

	sqlStatement := "UPDATE schedule SET year=?, month=?, day=?, starthour=?, startminute=?, endhour=?, endminute=? WHERE id=?"
	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		fmt.Print(err.Error())
	}
	defer stmt.Close()

	result, err := stmt.Exec(sch.Year, sch.Month, sch.Day, sch.StartHour, sch.StartMinute, sch.EndHour, sch.EndMinute, sch.ID)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
		return c.JSON(http.StatusCreated, sch)
	}

	return c.JSON(http.StatusOK, sch.ID)
}

func DeleteSchedule(c echo.Context) error {
	con := db.CreateConnection()
	request_id := c.Param("id")
	sqlStatement := "DELETE FROM schedule where id = ?"
	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		fmt.Println(err.Error())
	}
	defer stmt.Close()

	result, err2 := stmt.Exec(request_id)
	if err2 != nil {
		panic(err2)
	}
	fmt.Println(result.RowsAffected())
	return c.JSON(http.StatusOK, "Deleted")
}
