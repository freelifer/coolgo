package apis

import (
	"github.com/freelifer/coolgo/dao/redis"
	"github.com/freelifer/coolgo/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Tb(c *gin.Context) {
	page_size := c.DefaultQuery("page_size", "")
	page_no := c.DefaultQuery("page_no", "")
	resp, _ := service.TbkCoupon(page_size, page_no)
	c.String(http.StatusOK, resp)
}
