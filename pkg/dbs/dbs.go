package dbs

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	_ "github.com/mattn/go-sqlite3"
)

func DbsInit(dbDriver string) *xorm.Engine {
	var engine *xorm.Engine

	if "mysql" == dbDriver {
		engine, _ = xorm.NewEngine("mysql", "root:123@/test?charset=utf8")
	} else {
		//sqlite3
		engine, _ = xorm.NewEngine("sqlite3", "./data.db")
	}
	return engine
}
