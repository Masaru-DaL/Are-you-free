package models

import (
	"Are-you-free/internal/db"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
)

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

// func GetSchedules() Schedules {
// 	con := db.CreateConnection()
// 	// db.CreateConnection()
// 	sqlStatement := "SELECT id, Year, Month, Day, StartHour, StartMinute, EndHour, EndMinute FROM schedule order by id"

// 	// .Query: レコードの取得
// 	rows, err := con.Query(sqlStatement)
// 	fmt.Println(rows)
// 	fmt.Println(err)
// 	if err != nil {
// 		fmt.Println(err)
// 		// return c.JSON(http.StatusCreated, u);
// 	}
// 	defer rows.Close()
// 	result := Schedules{}

// 	// .Next: 各レコードに対して操作する
// 	for rows.Next() {
// 		schedule := Schedule{}
// 		// .Scan: 引数に渡したポインタにレコードの内容を読み込ませる
// 		err2 := rows.Scan(&schedule.ID, &schedule.Year, &schedule.Month, &schedule.Day, &schedule.StartHour, &schedule.StartMinute, &schedule.EndHour, &schedule.EndMinute)

// 		// エラーが発生した場合、終了する
// 		if err2 != nil {
// 			fmt.Println(err2)
// 		}
// 		result.Schedules = append(result.Schedules, schedule)
// 	}
// 	return result

// }

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
