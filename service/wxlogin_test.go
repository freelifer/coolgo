package service

import (
	"testing"
)

func Test_weiXinLogin(t *testing.T) {
	data, err := weiXinLogin("", "", "")

	t.Log(data)
	t.Error(err)
}
