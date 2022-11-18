package controllers

import (
	"net/http"

	"github.com/labstack/echo"
)

/* 全てのスケジュールを取得する */
func GetSchedules(c echo.Context) error {
	result := models.GetSchedule()
	println("Get All Schedules")
	return c.JSON(http.StatusOK, result)
}
