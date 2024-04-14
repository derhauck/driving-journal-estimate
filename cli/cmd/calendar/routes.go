package calendar

import (
	"derhauck/driving-journal-estimate/factory"
	"net/http"

	"github.com/gin-gonic/gin"
)

func initRoutes(router *gin.Engine) {
	router.GET("/health", func(c *gin.Context) {
		c.Status(http.StatusNoContent)
	})
	router.GET("/random", factory.NewRandomController().Random)
	router.POST("/config", factory.NewRandomController().Configuration)
}
