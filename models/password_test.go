package models

import (
	"testing"
)

func Test_GetWxUserPasswords(t *testing.T) {

	wxUserPasswords, err := GetWxUserPasswords(666)

	if err != nil {
		t.Error(err)
		return
	}

	t.Log("Test_GetWxUserPasswords Success ", wxUserPasswords)
}
