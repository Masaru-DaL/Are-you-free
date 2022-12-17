package models

import "time"

type User struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	Password   string    `json:"password"`
	Is_admin   bool      `json:"is_admin"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
}

type Users struct {
	Users []User `json:"users"`
}

type Schedule struct {
	ID          int       `json:"id"`
	Year        int       `json:"year"`
	Month       int       `json:"month"`
	Day         int       `Json:"day"`
	StartHour   int       `json:"starthour"`
	StartMinute int       `json:"startminute"`
	EndHour     int       `json:"endhour"`
	EndMinute   int       `json:"endminute"`
	Created_at  time.Time `json:"created_at"`
	Updated_at  time.Time `json:"updated_at"`
}

type Schedules struct {
	Schedules []Schedule `json:"Schedule"`
}
