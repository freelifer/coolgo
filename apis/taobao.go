package apis

import (
	"fmt"
	"github.com/freelifer/coolgo/dao/redis"
	"github.com/freelifer/coolgo/service"
	"github.com/freelifer/coolgo/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func TbkCoupon(c *gin.Context) {
	page_size := c.DefaultQuery("page_size", "")
	page_no := c.DefaultQuery("page_no", "")

	key := utils.Md5(fmt.Sprintf("page_size_%s_page_no_%s", page_size, page_no))

	value, err := redis.Get(key)
	if err == nil {
		c.String(http.StatusOK, value)
		return
	}
	resp, _ := service.TbkCoupon(page_size, page_no)
	redis.PutCoupons(key, resp)
	c.String(http.StatusOK, resp)
}

func Tbktpwd(c *gin.Context) {
	resp, _ := service.Tbktpwd("", "", "")
	c.String(http.StatusOK, resp)
}

func TbkItemInfo(c *gin.Context) {
	num_iids := c.DefaultQuery("num_iids", "")
	resp, _ := service.TbkItemInfo(num_iids)
	c.String(http.StatusOK, resp)
}

func TbkShopGet(c *gin.Context) {
	resp, _ := service.TbkShopGet()
	c.String(http.StatusOK, resp)
}
