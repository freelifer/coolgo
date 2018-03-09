package apis

import (
	"github.com/freelifer/coolgo/app/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func WeiXinLogin(c *gin.Context) {
	code := c.Query("code")
	wxData, err := service.WeiXinLogin(code)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"errcode": 40001,
			"errmsg":  err.Error(),
		})
		return
	}
	// openid 读取数据库
	c.JSON(http.StatusOK, gin.H{
		"errcode": 0,
		"errmsg":  "",
		"data": gin.H{
			"sessionid": wxData.SessionKey,
		},
	})
}
