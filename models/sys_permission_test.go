package models

import (
	"encoding/json"
	"fmt"
	"testing"
)

func Test_CreatePermissionMenu(t *testing.T) {

	//name string, parentId int64, route string, description string
	affected, err := CreatePermissionMenu("首页", 0, "/", "")

	if err != nil {
		t.Error(err)
		return
	}

	t.Log("CreateWxUser Success ", affected)
}

func Test_GetPermissionMenus(t *testing.T) {

	menus, err := GetPermissionMenus()

	if err != nil {
		t.Error(err)
		return
	}

	for i, d := range menus {
		fmt.Printf("DataIndex : %d        DataContent : %#v\n", i, d)
		t.Log("Test_GetPermissionMenus Success ", d)
	}
	b, err := json.Marshal(menus)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("Test_GetPermissionMenus Success ", string(b))
}
