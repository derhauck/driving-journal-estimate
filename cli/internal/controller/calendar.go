package controller

import (
	"derhauck/driving-journal-estimate/public/calendar"
	"derhauck/driving-journal-estimate/public/config"
	"derhauck/driving-journal-estimate/public/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CalendarRandomParam struct {
	Days  int     `form:"days"`
	Total float32 `form:"total"`
}

type CalendarConfigurationParams struct {
	Days  int     `form:"days"`
	Total float32 `form:"total"`
}
type CalendarController struct {
	Month  *calendar.Month
	Logger logger.Inf
}

func (r *CalendarController) Random(c *gin.Context) {
	var calendarParam CalendarRandomParam
	if err := c.ShouldBind(&calendarParam); err == nil {
		r.Month.RandomDays(calendarParam.Days)
		r.Month.Calculate(calendarParam.Total)

		c.JSON(http.StatusOK, gin.H{
			"message": r.Month,
		})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

}

func (r *CalendarController) Configuration(c *gin.Context) {
	var calendarParam config.File
	if err := c.ShouldBind(&calendarParam); err == nil {
		if calendarParam.Total == 0 {
			calendarParam.Total = 10000
		}

		if calendarParam.Baseline == 0 {
			calendarParam.Baseline = 0.5
		}
		calendarParam.DayConfig()
		r.Month.Days = calendarParam.DayConfig()
		r.Month.Calculate(calendarParam.Total)

		c.JSON(http.StatusOK, gin.H{
			"message": r.Month,
		})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

}
