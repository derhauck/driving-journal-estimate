package controller

import (
	"driving-journal-estimate/public/calendar"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CalendarParam struct {
	Days  int     `form:"days"`
	Total float32 `form:"total"`
}
type RandomController struct {
	Month *calendar.Month
}

func (r *RandomController) Random(c *gin.Context) {
	var calendarParam CalendarParam
	if err := c.ShouldBind(&calendarParam); err == nil {
		r.Month.RandomDays(calendarParam.Days)
		err = r.Month.Calculate(calendarParam.Total)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": r.Month,
		})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

}
