package controllers

import (
	"Are-you-free/internal/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

/* メソッドはmain.goのハンドラとして使用する */

/* 全てのスケジュールを取得する */
func GetSchedules(c echo.Context) error {
	// modelsに定義された関数を実行する
	result := models.GetSchedule()
	println("Get All Schedules")
	return c.JSON(http.StatusOK, result)
}
