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
	// TODO: Add your controller struct fields here
}

func (r *RandomController) Random(c *gin.Context) {
	var calendarParam CalendarParam
	if err := c.ShouldBind(&calendarParam); err == nil {
		month := calendar.NewRandomMonth(calendarParam.Days)
		err = month.Calculate(calendarParam.Total)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": month,
		})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

}
