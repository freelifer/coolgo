package apis

import (
	"fmt"
	"github.com/bitly/go-simplejson"
	"github.com/freelifer/coolgo/pkg/utils"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

func HelloWorld(c *gin.Context) {
	code := c.Query("q")
	c.String(http.StatusOK, code)
}

func Ping(c *gin.Context) {
	c.String(http.StatusOK, "Hello World!")
}

func WXLogin(c *gin.Context) {
	code := c.Query("code")
	s := fmt.Sprintf("https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code",
		"wxcb7e15ce8102e6d3", "5852a3a293252fe02d2b83dbc3f8ec36", code)
	// s := fmt.Sprintf("http://int.dpool.sina.com.cn/iplookup/iplookup.php?format=json&ip=%s", "218.4.255.255")
	resp, err := http.Get(s)
	if err != nil {
		// handle error
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	js, err := simplejson.NewJson(body)

	if err != nil {
		panic(err.Error())
	}
	// if errcode != nil {
	// c.String(http.StatusOK, string(body))
	// return
	// }

	fmt.Println(js)
	openid := js.Get("openid").MustString()
	if len(openid) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"errcode": 40001,
			"errmsg":  "openid has not",
		})
		return
	}

	sessionid := utils.NewSessionID()
	// session_key := js.Get("session_key").MustString()
	// unionid := js.Get("unionid").MustString()
	// config.Cache.Put(sessionid, gin.H{
	// 	"session_key": session_key,
	// 	"unionid":     unionid,
	// 	"openid":      openid,
	// }, 20*time.Second)
	// c.String(http.StatusOK, string(body))
	c.JSON(http.StatusOK, gin.H{
		"errcode": 0,
		"errmsg":  "",
		"data": gin.H{
			"sessionid": sessionid,
		},
	})
	// session_key, err := js.Get("session_key")
	// unionid, err := js.Get("unionid")

	// c.JSON(http.StatusOK, gin.H{
	// "openid": openid,
	// })
}
