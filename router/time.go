package router

// 使用するライブラリをimport
import (
	"Are-you-free/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

// 関数 GetTasksHandlerは引数がecho.Context型のc で、戻り値はerror型である
func GetTimesHandler(c echo.Context) error {

	// model(package)の関数GetTasksを実行し、戻り値をtimes,errと定義する。
	times, err := model.GetTimes()

	// errが空でない時は StatusBadRequest(*5) を返す
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad Request")
	}

	// StasusOK と timesを返す
	return c.JSON(http.StatusOK, times)
}
