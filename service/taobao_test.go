package service

import (
	"testing"
)

func Test_TbkCoupon(t *testing.T) {
	body, _ := TbkCoupon("0", "0")

	t.Log(body)
}

func Test_Tbktpwd(t *testing.T) {
	body, _ := Tbktpwd("红心剃毛机剃毛球修剪器充电式吸刮毛器衣服除毛球器去球机打毛器",
		"https://uland.taobao.com/coupon/edetail?e=QrweHOsuj2sGQASttHIRqa59KcrVlGwiINX7EmZ0YliAnVXUwZx8IhEJwQnkdjDZ7aNoMgWMAfl%2FJ9rOMHOUiMPWjN00SJO8DfqEFBOhTcyABFd8ZHrxxMbGHnPVW3ogxXxVCXdmdxYzhukA9d4NUmPfrr0N2WBeCqEIqV4SxV7k92%2BM7h46c6J7%2BkHL3AEW&traceId=0ab013ac15116183416294882e",
		"http://img.alicdn.com/tfscom/i2/1048721401/TB1d1S4b4PI8KJjSspfXXcCFXXa_!!0-item_pic.jpg")

	t.Log(body)

}
func Test_buildPublicData(t *testing.T) {
	tb := NewTbService()
	tb.putPublicData("taobao.tbk.item.get", "123456", "abcdefg")
	t.Log(tb.table)
}

func Test_currentTime(t *testing.T) {
	t.Log(currentTime())
}

func Test_areNotEmpty(t *testing.T) {
	if areNotEmpty("", "") {
		t.Error("   为空")
	}
	if areNotEmpty("", "a") {
		t.Error("  a 为空")
	}
	if !areNotEmpty("a", "a") {
		t.Error("a a 为空")
	}
}

func Test_Md5(t *testing.T) {
	t.Log(Md5_2("abcdefg"))
	t.Log(Md5("abcdefg"))
}

func Test_createUrl(t *testing.T) {
	tbs := NewTbService()
	tbs.putPublicData("taobao.tbk.dg.item.coupon.get", "24659164", "")
	tbs.putPrivateData("adzone_id", "148758292")
	tbs.signTopRequest("cbe2b136be37cd2b66fd4490b8fbfb94", SIGN_METHOD_HMAC)

	t.Log(tbs.createUrl())
}
