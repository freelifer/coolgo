package apis

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Tb(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
