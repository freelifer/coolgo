package db

import (
	"testing"
)

func Test_FindUser(t *testing.T) {

	user, has, err := FindUser1("654321")

	if err != nil {
		t.Error(err)
		return
	}
	if !has {
		t.Log("has not User")
		return
	}

	t.Log(user.Phone)
}

func Test_CreateUser(t *testing.T) {

	affected, err := CreateUser("134567")

	if err != nil {
		t.Error(err)
		return
	}

	t.Log(affected)
}
