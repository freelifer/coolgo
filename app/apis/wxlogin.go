package apis

import (
	"github.com/freelifer/coolgo/app/service"
	"github.com/freelifer/coolgo/models"
	"github.com/freelifer/coolgo/pkg/redis"
	"github.com/freelifer/coolgo/pkg/utils"
	"github.com/gin-gonic/gin"
)

func WeiXinLogin(c *gin.Context) {
	code := c.Query("code")
	wxData, err := service.WeiXinLogin(code)

	if err != nil {
		errorJSON(c, 40002, err.Error())
		return
	}

	wxUer, err := models.GetOrCreateWxUser(wxData.Openid)
	if err != nil {
		errorJSON(c, 40002, err.Error())
		return
	}

	json, err := models.WxUserToJson(wxUer)
	if err != nil {
		errorJSON(c, 40002, err.Error())
		return
	}

	sessionId := utils.NewSessionID()
	e := redis.PutSession(sessionId, json)
	if e != nil {
		errorJSON(c, 40002, e.Error())
		return
	}

	successJSON(c, gin.H{
		"sessionid": sessionId,
	})
}
