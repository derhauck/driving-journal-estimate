package main

import (
	"driving-journal-estimate/factory"
	"github.com/gin-gonic/gin"
)

func initRoutes(router *gin.Engine) {
	router.GET("/", factory.NewRandomController().Random)
}
