package models

import (
	"encoding/json"
	"time"
)

/* 用户 */
type WxUser struct {
	Id        int64
	WxOpenid  string    `xorm:"unique"`
	CreatedAt time.Time `xorm:"created" json:"-"`
	UpdatedAt time.Time `xorm:"updated" json:"-"`
}

func CreateWxUser(openid string) (int64, error) {
	user := new(WxUser)
	user.WxOpenid = openid
	affected, err := engine.Insert(user)
	return affected, err
}

func GetWxUser(openid string) (*WxUser, bool, error) {
	user := &WxUser{WxOpenid: openid}
	has, err := engine.Get(user)
	return user, has, err
}

func GetOrCreateWxUser(openid string) (*WxUser, error) {
	// Get WxUser First
	wxUser, has, err := GetWxUser(openid)
	if err != nil {
		return nil, err
	}

	if !has {
		// Create WxUser Second
		_, err := CreateWxUser(openid)
		if err != nil {
			return nil, err
		}
		// Get WxUser Second
		wxUser, _, err := GetWxUser(openid)
		if err != nil {
			return nil, err
		}
		return wxUser, nil
	}
	return wxUser, nil
}

func WxUserToJson(s *WxUser) (string, error) {
	b, err := json.Marshal(s)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func JsonToWxUser(s string) (*WxUser, error) {
	var wxUser WxUser
	err := json.Unmarshal([]byte(s), &wxUser)
	if err != nil {
		return nil, err
	}
	return &wxUser, nil
}
