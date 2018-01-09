package dao

import ()

/* 用户 */
type User struct {
	Id       int64
	Nickname string
	Password string
	WxOpenid string
	Phone    string
}

/* 角色 */
type Role struct {
	Id   int64
	Name string
	Desc string
}

/* 权限 */
type Permission struct {
	Id   int64
	Name string
	Desc string
}

/* 账单 */
type Bill struct {
	Id   int64
	Name string
}
