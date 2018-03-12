package apis

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func errorJSON(c *gin.Context, errcode int, errmsg string) {
	c.JSON(http.StatusOK, gin.H{
		"errcode": errcode,
		"errmsg":  errmsg,
	})
}

func successJSON(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"errcode": 0,
		"errmsg":  "",
		"data":    data,
	})
}
