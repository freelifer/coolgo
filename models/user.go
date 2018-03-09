package models

import ()

/* 用户 */
type User struct {
	Id        int64
	Nickname  string
	Password  string
	WxOpenid  string `xorm:"unique"`
	Phone     string
	CreatedAt int64 `xorm:"created"`
	UpdatedAt int64 `xorm:"updated"`
}
