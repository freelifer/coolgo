package apis

import (
	"github.com/freelifer/coolgo/models"
	"github.com/gin-gonic/gin"

	"encoding/base64"
	"strings"
)

func BackendLogin(c *gin.Context) {
	name := c.DefaultPostForm("name", "")
	password := c.DefaultPostForm("password", "")

	successJSON(c, gin.H{
		"token": name + password,
	})
}

func BackendSignin(c *gin.Context) {
	if !checkAuth(c, "pig", "123456") {
		c.JSON(401, "")
		return
	}
	menus, err := models.GetPermissionMenus()
	if err != nil {
		errorJSON(c, 40002, err.Error())
		return
	}
	c.Header("Access-Control-Allow-Origin", "*")
	// c.JSON(http.StatusOK, gin.H{
	// 	"menus": menus,
	// })

	successJSON(c, gin.H{
		"menus": menus,
	})
}

func checkAuth(c *gin.Context, user string, password string) bool {

	s := strings.SplitN(c.GetHeader("Authorization"), " ", 2)
	if len(s) != 2 {
		return false
	}

	b, err := base64.StdEncoding.DecodeString(s[1])
	if err != nil {
		return false
	}

	pair := strings.SplitN(string(b), ":", 2)
	if len(pair) != 2 {
		return false
	}

	return pair[0] == user && pair[1] == password
}
