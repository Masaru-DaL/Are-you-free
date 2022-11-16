package model

import (
	_ "gorm.io/gorm"
)

type Time struct {
	ID     int
	Year   int
	Month  int
	Day    int
	Hour   int
	Minute int
}

// type StartTime struct {
// 	ID     int
// 	Year   int
// 	Month  int
// 	Day    int
// 	hour   int
// 	minute int
// }

// type EndTime struct {
// 	ID     int
// 	Year   int
// 	Month  int
// 	Day    int
// 	hour   int
// 	minute int
// }

// 関数GetTasksは、引数はなく、戻り値は[]Task型（Task型のスライス）とerror型である
func GetTimes() ([]Time, error) {

	// 空のタスクのスライスである、tasksを定義する
	var times []Time

	// tasksにDBのタスク全てを代入する。その操作の成否をerrと定義する(*4)
	err := db.Find(&times).Error

	return times, err
}
