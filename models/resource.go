package models

import (
	"time"
)

/**
 * 菜单
 *
 * id、菜单名称、菜单路由、父菜单、
 * 描述、开关、创建时间、更新时间
 */
type MenuResource struct {
	Id          int64
	Name        string
	Route       string
	ParentId    int64
	Description string    `json:"description"`
	Enable      bool      `json:"enable"`
	CreatedAt   time.Time `xorm:"created" json:"-"`
	UpdatedAt   time.Time `xorm:"updated" json:"-"`
}

/**
 * 资源
 *
 * id、类型(MENU、API)
 */
type Resource struct {
	Id     int64
	TypeId int64
}

/**
 * 资源类型
 *
 * id、类型名称
 */
type ResourceType struct {
	Id        int64
	Name      string    `xorm:"varchar(25) notnull unique"`
	CreatedAt time.Time `xorm:"created" json:"-"`
	UpdatedAt time.Time `xorm:"updated" json:"-"`
}

// 增
func CreateResourceType(name string) (int64, error) {
	resType := new(ResourceType)
	resType.Name = name
	affected, err := engine.Insert(resType)
	return affected, err
}

// 删

// 改

// 查
func GetResourceTypes() (menus []ResourceType, err error) {
	types = make([]ResourceType, 0)
	if err = engine.Find(&types); err != nil {
		return nil, err
	}

	return types, nil
}

func GetResourceType(id int64) (*ResourceType, bool, error) {
	resType := &ResourceType{Id: id}
	has, err := engine.Get(resType)
	return resType, has, err
}
