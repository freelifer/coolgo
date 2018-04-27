package models

import (
	"github.com/freelifer/coolgo/pkg/config"
	"github.com/freelifer/coolgo/pkg/redis"

	"fmt"
	"github.com/go-xorm/xorm"
)

var engine *xorm.Engine

func init() {

	engine = DbsInit()

	// engine.Logger().SetLevel(core.LOG_DEBUG)
	engine.ShowSQL(true)

	engine.Sync2(new(WxUser), new(Password), new(WxUserPassword), new(PermissionMenu))

	status := config.Bool("app::redis_status")
	if !status {
		return
	}
	// 从配置文件获取redis的ip以及db
	REDIS_HOST := config.String("redis::host")
	REDIS_DB := config.Int("redis::db")
	maxidle := config.DefaultInt("redis::maxidle", 1)
	maxactive := config.DefaultInt("redis::maxactive", 10)
	redis.RedisInit(REDIS_HOST, REDIS_DB, maxidle, maxactive)
}

func DbsInit() *xorm.Engine {
	var engine *xorm.Engine

	dbDriver := config.String("app::db")
	if "mysql" == dbDriver {
		dbuser := config.String("mysql::dbuser")
		dbpwd := config.String("mysql::dbpwd")
		dbname := config.String("mysql::dbname")

		if dbuser == "" {
			dbuser = "root"
		}
		if dbpwd == "" {
			dbpwd = "root"
		}
		if dbname == "" {
			dbname = "coolgo"
		}
		conn := fmt.Sprintf("%s:%s@/%s?charset=utf8", dbuser, dbpwd, dbname)
		engine, _ = xorm.NewEngine("mysql", conn)
	} else {
		//sqlite3
		dbname := config.String("sqlite3::dbname")
		if dbname == "" {
			dbname = "data.db"
		}
		engine, _ = xorm.NewEngine("sqlite3", dbname)
	}
	return engine
}
