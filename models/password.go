package models

import (
	"time"
)

/* 密码 */
type Password struct {
	Id           int64
	UserName     string
	UserPassword string
	Title        string
	CreatedAt    time.Time `xorm:"created" json:"-"`
	UpdatedAt    time.Time `xorm:"updated" json:"-"`
}

type WxUserPassword struct {
	Id         int64
	WxUserId   int64
	PasswordId int64
	Status     int       `xorm:"default 0"` // 0:表示未使用 1：表示使用 2：表示删除
	CreatedAt  time.Time `xorm:"created" json:"-"`
	UpdatedAt  time.Time `xorm:"updated" json:"-"`
}

func CreatePassword(userName, userPassword string) (int64, error) {
	password := new(Password)
	password.UserName = userName
	password.UserPassword = userPassword
	affected, err := engine.Insert(password)
	return affected, err
}

func GetPassword(id int64) (*Password, bool, error) {
	password := &Password{Id: id}
	has, err := engine.Get(password)
	return password, has, err
}

func CreateWxUserPassword(wxUserId, passwordId int64) (int64, error) {
	wxUserPassword := new(WxUserPassword)
	wxUserPassword.WxUserId = wxUserId
	wxUserPassword.PasswordId = passwordId
	wxUserPassword.Status = 1
	affected, err := engine.Insert(wxUserPassword)
	return affected, err
}

func GetWxUserPasswords(wxUserId int64) ([]WxUserPassword, error) {
	wxUserPasswords := make([]WxUserPassword, 0)
	err := engine.Where("wx_user_id = ?", wxUserId).Find(&wxUserPasswords)
	return wxUserPasswords, err
}
