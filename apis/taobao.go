package apis

import (
	"fmt"
	"github.com/freelifer/coolgo/dao/redis"
	"github.com/freelifer/coolgo/service"
	"github.com/freelifer/coolgo/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Tb(c *gin.Context) {
	page_size := c.DefaultQuery("page_size", "")
	page_no := c.DefaultQuery("page_no", "")

	key := utils.Md5(fmt.Sprintf("page_size_%s_page_no_%s", page_size, page_no))

	value, err := redis.Get(key)
	if err == nil {
		c.String(http.StatusOK, value)
		return
	}
	resp, _ := service.TbkCoupon(page_size, page_no)
	redis.Set(key, resp)
	c.String(http.StatusOK, resp)
}
