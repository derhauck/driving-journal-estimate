package main

import (
	"driving-journal-estimate/service/factory"
	"github.com/gin-gonic/gin"
)

var router = gin.Default()

func init() {
	router.GET("/", factory.NewRandomController().Random)
}
