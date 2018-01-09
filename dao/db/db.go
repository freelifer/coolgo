package db

import (
	"github.com/freelifer/coolgo/config"
	"github.com/freelifer/coolgo/dao"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	_ "github.com/mattn/go-sqlite3"
)

var dbDriver string
var engine *xorm.Engine

func init() {
	dbDriver := config.String("app::db")
	if "sqlite3" == dbDriver {
		engine, _ = xorm.NewEngine("sqlite3", "./data.db")
	} else {
		engine, _ = xorm.NewEngine("mysql", "root:123@/test?charset=utf8")
	}

	// engine.Logger().SetLevel(core.LOG_DEBUG)
	engine.ShowSQL(true)

	engine.Sync2(new(dao.User), new(dao.Role), new(dao.Permission), new(dao.Bill))
}
