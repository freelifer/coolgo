package apis

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Huobipro(c *gin.Context) {
	name := c.DefaultPostForm("name", "")
	transactions := c.DefaultPostForm("transactions", "")

	c.Header("Access-Control-Allow-Origin", "*")
	c.String(http.StatusOK, "HelloWorld")
	c.JSON(http.StatusOK, gin.H{
		"name":         name,
		"transactions": transactions,
	})
}
