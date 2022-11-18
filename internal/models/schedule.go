package models

import (
	"Are-you-free/internal/db"
	"database/sql"
	"fmt"
)

type Schedule struct {
	ID          string `json:"id"`
	Year        int    `json:"year"`
	Month       int    `json:"month"`
	Day         int    `Json:"day"`
	StartHour   int    `json:"starthour"`
	StartMinute int    `json:"startminute"`
	EndHour     int    `json:"endhour"`
	EndMinute   int    `json:"endminute"`
}

type Schedules struct {
	Schedules []Schedule `json:"Schedule"`
}

var con *sql.DB

func GetSchedule() Schedules {
	con := db.CreateConnection()
	// db.CreateConnection()
	sqlStatement := "SELECT id, Year, Month, Day, StartHour, StartMinute, EndHour, EndMinute FROM schedule order by id"

	rows, err := con.Query(sqlStatement)
	fmt.Println(rows)
	fmt.Println(err)
	if err != nil {
		fmt.Println(err)
		// return c.JSON(http.StatusCreated, u);
	}
	defer rows.Close()
	result := Schedules{}

	for rows.Next() {
		schedule := Schedule{}

		err2 := rows.Scan(&schedule.ID, &schedule.Year, &schedule.Month, &schedule.Day, &schedule.StartHour, &schedule.StartMinute, &schedule.EndHour, &schedule.EndMinute)

		// エラーが発生した場合、終了する
		if err2 != nil {
			fmt.Println(err2)
		}
		result.Schedules = append(result.Schedules, schedule)
	}
	return result

}
