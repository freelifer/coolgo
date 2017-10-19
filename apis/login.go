package apis

import (
	"fmt"
	"github.com/bitly/go-simplejson"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func WXLogin(c *gin.Context) {
	s := fmt.Sprintf("https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code",
		"wxcb7e15ce8102e6d3", "5852a3a293252fe02d2b83dbc3f8ec36", "011EZl8Q0itqI72VlK5Q0qiA8Q0EZl8Y")
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

	// openid, err := js.Get("openid")
	// session_key, err := js.Get("session_key")
	// unionid, err := js.Get("unionid")

	fmt.Println(js)
	c.String(http.StatusOK, string(body))
}
