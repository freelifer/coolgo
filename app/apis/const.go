package apis

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var ParamErr = gin.H{"errcode": 40001, "errmsg": "param error"}
var InnerErr = gin.H{"errcode": 40002, "errmsg": "inner err"}

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
