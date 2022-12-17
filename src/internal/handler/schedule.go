package handler

import (
	"fmt"
	"net/http"
	"src/internal/db"
	"src/internal/models"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
)

/* 1件取得の関数 */
func GetOneSchedule(c echo.Context) error {
	con := db.CreateConnection()

	schedule_id := c.Param("id")
	strconv.Atoi(schedule_id)

	sqlStatement := "SELECT id, Year, Month, Day, StartHour, StartMinute, EndHour, EndMinute, Created_at, Updated_at FROM schedules WHERE id = ? LIMIT 1"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		fmt.Println(err)
	}
	defer stmt.Close()

	schedule := models.Schedule{}
	err2 := stmt.QueryRow(schedule_id).Scan(&schedule.ID, &schedule.Year, &schedule.Month, &schedule.Day, &schedule.StartHour, &schedule.StartMinute, &schedule.EndHour, &schedule.EndMinute, &schedule.Created_at, &schedule.Updated_at)
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
		"created_at":  schedule.Created_at,
		"updated_at":  schedule.Updated_at,
	})
}

/* 全件取得の関数 */
func GetAllSchedules(c echo.Context) error {
	con := db.CreateConnection()

	sqlStatement := "SELECT ID, Year, Month, Day, StartHour, StartMinute, EndHour, EndMinute, Created_at, Updated_at FROM schedules"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		fmt.Println(err)
	}
	defer stmt.Close()

	// rows, err := stmt.Query(1)
	rows, err := stmt.Query()
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	schedules := models.Schedules{}
	for rows.Next() {
		schedule := models.Schedule{}
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
	sch := new(models.Schedule)
	if err := c.Bind(sch); err != nil {
		return err
	}

	sqlStatement := "INSERT INTO schedules(year, month, day, starthour, startminute, endhour, endminute) VALUES(?, ?, ?, ?, ?, ?, ?, ?)"
	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		fmt.Print(err.Error())
	}
	defer stmt.Close()

	result, err2 := stmt.Exec(sch.Year, sch.Month, sch.Day, sch.StartHour, sch.StartMinute, sch.EndHour, sch.EndMinute)

	// エラーが発生したら終了する
	if err2 != nil {
		panic(err2)
	}
	fmt.Println(result.LastInsertId())

	return c.JSON(http.StatusCreated, sch.ID)
}

func PutSchedule(c echo.Context) error {
	con := db.CreateConnection()
	sch := new(models.Schedule)
	if err := c.Bind(sch); err != nil {
		return err
	}

	sqlStatement := "UPDATE schedules SET year=?, month=?, day=?, starthour=?, startminute=?, endhour=?, endminute=? WHERE id=?"
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
	sqlStatement := "DELETE FROM schedules where id = ?"
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
