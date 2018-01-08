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

	// release
	router.GET("/wx/login", WeiXinLogin)

	v1 := router.Group("/v1")
	{
		v1.GET("/tb/coupons", TbkCoupon)
		v1.GET("/tb/tpwd", Tbktpwd)
		v1.GET("/tb/item_info", TbkItemInfo)
		v1.GET("/tb/shop_get", TbkShopGet)
	}

	return router
}
