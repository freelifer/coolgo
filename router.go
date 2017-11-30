package main

import (
	. "github.com/freelifer/coolgo/apis"
	"github.com/gin-gonic/gin"
)

func initRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/", HelloWorld)
	router.GET("/ping", Ping)
	router.GET("/WXLogin", WXLogin)
	router.GET("/tb/coupons", TbkCoupon)
	router.GET("/tb/tpwd", Tbktpwd)
	router.GET("/tb/item_info", TbkItemInfo)
	router.GET("/tb/shop_get", TbkShopGet)

	return router
}
