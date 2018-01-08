package dao

import ()

/* 用户 */
type User struct {
	Id       int
	Nickname string
	Password string
	WxOpenid string
	Phone    string
}

/* 角色 */
type Role struct {
	Id   int
	Name string
	Desc string
}

/* 权限 */
type Permission struct {
	Id   int
	Name string
	Desc string
}

/* 账单 */
type Bill struct {
	Id   int
	name string
}
