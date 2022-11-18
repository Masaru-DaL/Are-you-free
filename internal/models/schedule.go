package models

type Shedule struct {
	ID          string `json:"id"`
	Year        int    `json:"Year"`
	Month       int    `json:"Month"`
	Day         int    `Json:"Day"`
	StartHour   int    `json:"StartHour"`
	StartMinute int    `json:"StartMinute"`
	EndHour     int    `json:"EndHour"`
	EndMinute   int    `json:"EndMinute"`
}

type Schedules struct {
	Schedules []Schedule `json:"Schedule"`
}
