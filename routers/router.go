package routers

import (
	. "github.com/freelifer/coolgo/app/apis"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/", HelloWorld)
	r.GET("/ping", Ping)
	r.GET("/WXLogin", WXLogin)

	// release
	r.GET("/wx/login", WeiXinLogin)

	v1 := r.Group("/v1")
	{
		v1.GET("/tb/coupons", TbkCoupon)
		v1.GET("/tb/tpwd", Tbktpwd)
		v1.GET("/tb/item_info", TbkItemInfo)
		v1.GET("/tb/shop_get", TbkShopGet)
	}

	return r
}
