package service

import (
	"testing"
)

func Test_TbkCoupon(t *testing.T) {
	body, _ := TbkCoupon()

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
