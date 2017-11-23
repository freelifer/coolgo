package main

import (
	. "github.com/freelifer/coolgo/apis"
	"github.com/gin-gonic/gin"
)

func initRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/ping", Ping)
	router.GET("/WXLogin", WXLogin)
	router.GET("/Tb", Tb)

	return router
}
