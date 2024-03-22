package calendar

import (
	"derhauck/driving-journal-estimate/factory"

	"github.com/gin-gonic/gin"
)

func initRoutes(router *gin.Engine) {
	router.GET("/random", factory.NewRandomController().Random)
	router.POST("/config", factory.NewRandomController().Configuration)
}
