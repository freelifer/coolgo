package db

import (
	"github.com/astaxie/beego/orm"
	"github.com/freelifer/coolgo/config"
	"github.com/freelifer/coolgo/dao"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
)

var dbDriver string

func init() {
	orm.Debug = true
	dbDriver := config.Config.String("app::db")
	if "sqlite3" == dbDriver {
		orm.RegisterDriver("sqlite3", orm.DRSqlite)
		orm.RegisterDataBase("default", "sqlite3", "data.db")
	} else {
		orm.RegisterDriver("mysql", orm.DRMySQL)
		orm.RegisterDataBase("default", "mysql", "root:root@/orm_test?charset=utf8")
	}

	orm.RegisterModel(new(dao.User))
	orm.RegisterModel(new(dao.Role))
	orm.RegisterModel(new(dao.Permission))
	orm.RegisterModel(new(dao.Bill))

	orm.RunSyncdb("default", false, true)
}
