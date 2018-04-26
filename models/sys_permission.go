package models

import (
	"time"
)

/* 菜单权限 */
type PermissionMenu struct {
	Id          int64     `json:"id"`
	Name        string    `json:"name"`
	ParentId    int64     `json:"parent_id"`
	Route       string    `json:"route"`
	Description string    `json:"description"`
	Enable      bool      `json:"enable"`
	CreatedAt   time.Time `xorm:"created" json:"-"`
	UpdatedAt   time.Time `xorm:"updated" json:"-"`
}

func CreatePermissionMenu(name string, parentId int64, route string, description string) (int64, error) {
	menu := new(PermissionMenu)
	menu.Name = name
	menu.ParentId = parentId
	menu.Route = route
	menu.Description = description
	affected, err := engine.Insert(menu)
	return affected, err
}

func GetPermissionMenus() (menus []PermissionMenu, err error) {
	menus = make([]PermissionMenu, 0)
	if err = engine.Find(&menus); err != nil {
		return nil, err
	}

	return menus, nil
}
