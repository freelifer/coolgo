package models

import (
	"testing"
)

func Test_CreateWxUser(t *testing.T) {

	affected, err := CreateWxUser("654321")

	if err != nil {
		t.Error(err)
		return
	}

	t.Log("CreateWxUser Success ", affected)
}

func Test_GetWxUser(t *testing.T) {
	user, has, err := GetWxUser("134567")

	if err != nil {
		t.Error(err)
		return
	}

	if !has {
		t.Log("Get Wx User not has")
		return
	}

	t.Log(user)
}

func Test_GetWxUser1(t *testing.T) {
	user, has, err := GetWxUser("654321")

	if err != nil {
		t.Error(err)
		return
	}

	if !has {
		t.Log("Get Wx User not has")
		return
	}

	t.Log(user)
}

func Test_GetOrCreateWxUser(t *testing.T) {
	user, err := GetOrCreateWxUser("654321")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(user)

	j, err := WxUserToJson(user)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(j)

	u, err := JsonToWxUser(j)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(u)
}
