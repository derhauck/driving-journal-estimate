package controller

import (
	"driving-journal-estimate/cli/public/calendar"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CalendarParam struct {
	Days  int     `form:"days"`
	Total float32 `form:"total"`
}
type RandomController struct {
	// TODO: Add your controller struct fields here
}

func (r *RandomController) Random(c *gin.Context) {
	var calendarParam CalendarParam
	if c.ShouldBind(&calendarParam) == nil {
		month := calendar.NewRandomMonth(calendarParam.Days)
		month.Calculate(calendarParam.Total)
		month.Print()

		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World!!!",
		})
	}

}
