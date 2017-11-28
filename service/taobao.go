package service

import (
	"bytes"
	"crypto/hmac"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"
)

const (
	SIGN_METHOD_MD5  string = "md5"
	SIGN_METHOD_HMAC string = "hmac"
)

// "Taobao_Config": {
//    "tb_app_key":"24659164",
//    "tb_app_secret":"cbe2b136be37cd2b66fd4490b8fbfb94",
//    "tb_adzone_id":"148758292",
//    "tb_pid": "mm_128081258_38438020_148758292",
//    "tb_favorites_id":["14267841", "14246808", "14246784"],
//    "tb_cat_ids":""
//  },
// mm_128081258_39714707_150048462
func TbkCoupon(page_size, page_no string) (string, error) {
	tbs := NewTbService()
	tbs.putPublicData("taobao.tbk.dg.item.coupon.get", "24659164", "")
	tbs.putPrivateData("adzone_id", "148758292")
	tbs.putPrivateData("platform", "2")
	tbs.putPrivateData("page_size", page_size)
	tbs.putPrivateData("page_no", page_no)
	tbs.signTopRequest("cbe2b136be37cd2b66fd4490b8fbfb94", SIGN_METHOD_HMAC)

	// s := fmt.Sprintf("https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code",
	// "wxcb7e15ce8102e6d3", "5852a3a293252fe02d2b83dbc3f8ec36", "code")
	// s := fmt.Sprintf("http://int.dpool.sina.com.cn/iplookup/iplookup.php?format=json&ip=%s", "218.4.255.255")
	resp, err := http.Get(tbs.createUrl())
	if err != nil {
		return "a", err
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	return string(body), nil
}

func Tbktpwd(text, url, logo_url string) (string, error) {
	tbs := NewTbService()
	tbs.putPublicData("taobao.tbk.tpwd.create", "24659164", "")
	tbs.putPrivateData("text", text)
	tbs.putPrivateData("url", url)
	tbs.putPrivateData("logo", logo_url)
	tbs.signTopRequest("cbe2b136be37cd2b66fd4490b8fbfb94", SIGN_METHOD_HMAC)

	// s := fmt.Sprintf("https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code",
	// "wxcb7e15ce8102e6d3", "5852a3a293252fe02d2b83dbc3f8ec36", "code")
	// s := fmt.Sprintf("http://int.dpool.sina.com.cn/iplookup/iplookup.php?format=json&ip=%s", "218.4.255.255")
	resp, err := http.Get(tbs.createUrl())
	if err != nil {
		return "a", err
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	return string(body), nil
}

// TbService Object and Method

type TbService struct {
	table map[string]string
}

func NewTbService() *TbService {
	var tbService TbService
	tbService.table = make(map[string]string)
	return &tbService
}

func (this *TbService) putPublicData(method string, app_key string, session string) {
	this.table["method"] = method
	this.table["app_key"] = app_key
	this.putPrivateData("session", session)
	this.table["timestamp"] = currentTime()
	this.table["format"] = "json"
	this.table["v"] = "2.0"
	this.table["sign_method"] = "hmac"
}

func (this *TbService) putPrivateData(key, value string) {
	if areNotEmpty(key, value) {
		this.table[key] = value
	}
}

func (this *TbService) signTopRequest(secret string, signMethod string) {
	// 第一步：检查参数是否已经排序
	keys := make([]string, len(this.table))
	i := 0
	for k, _ := range this.table {
		keys[i] = k
		i++
	}
	sort.Strings(keys)

	// 第二步：把所有参数名和参数值串在一起
	var buffer bytes.Buffer
	if strings.Compare(signMethod, SIGN_METHOD_MD5) == 0 {
		buffer.WriteString(secret)
	}

	for _, k := range keys {
		value := this.table[k]
		if areNotEmpty(k, value) {
			buffer.WriteString(k)
			buffer.WriteString(value)
		}
	}

	var sign string
	// 第三步：使用MD5/HMAC加密
	if strings.Compare(signMethod, SIGN_METHOD_HMAC) == 0 {
		sign = Hmac(secret, buffer.String())
	} else {
		buffer.WriteString(secret)
		sign = Md5(buffer.String())
	}

	this.table["sign"] = sign
}

func (this *TbService) createUrl() string {
	var buffer bytes.Buffer
	buffer.WriteString("http://gw.api.taobao.com/router/rest?")
	i := 0
	tableLen := len(this.table)
	for k, v := range this.table {
		i++
		buffer.WriteString(url.QueryEscape(k))
		buffer.WriteString("=")
		buffer.WriteString(url.QueryEscape(v))
		if tableLen != i {
			buffer.WriteString("&")
		}
	}
	return buffer.String()
}

func currentTime() string {
	t := time.Now()
	return t.Format("2006-01-02 15:04:05")
}

func areNotEmpty(key, value string) bool {
	if len(key) == 0 {
		return false
	}

	if len(value) == 0 {
		return false
	}
	return true

}

func Hmac(key, data string) string {
	hmac := hmac.New(md5.New, []byte(key))
	hmac.Write([]byte(data))
	return strings.ToUpper(hex.EncodeToString(hmac.Sum([]byte(""))))
}

func Md5(data string) string {
	md5 := md5.New()
	md5.Write([]byte(data))
	md5Data := md5.Sum([]byte(""))
	return hex.EncodeToString(md5Data)
}

func Md5_2(data string) string {
	md5 := md5.New()
	md5.Write([]byte(data))
	md5Data := md5.Sum([]byte(""))
	return strings.ToUpper(fmt.Sprintf("%x", md5Data))
}
