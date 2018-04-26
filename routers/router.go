package routers

import (
	. "github.com/freelifer/coolgo/app/apis"
	// "github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	// r.Use(gzip.Gzip(gzip.BestCompression))

	r.GET("/", HelloWorld)
	r.GET("/ping", Ping)
	r.GET("/WXLogin", WXLogin)

	r.POST("/huobipro", Huobipro)

	// release
	r.GET("/wx/login", WeiXinLogin)

	v1 := r.Group("/v1")
	{
		v1.GET("/tb/coupons", TbkCoupon)
		v1.GET("/tb/tpwd", Tbktpwd)
		v1.GET("/tb/item_info", TbkItemInfo)
		v1.GET("/tb/shop_get", TbkShopGet)
	}

	// backend api
	{
		v1.POST("/login", BackendLogin)
		v1.GET("/signin", BackendSignin)
	}
	return r
}
