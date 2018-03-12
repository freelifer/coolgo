package models

import (
	"github.com/freelifer/coolgo/pkg/config"
	"github.com/freelifer/coolgo/pkg/dbs"
	"github.com/freelifer/coolgo/pkg/redis"

	"github.com/go-xorm/xorm"
)

var engine *xorm.Engine

func init() {

	dbDriver := config.String("app::db")
	engine = dbs.DbsInit(dbDriver)

	// engine.Logger().SetLevel(core.LOG_DEBUG)
	engine.ShowSQL(true)

	engine.Sync2(new(WxUser), new(Password), new(WxUserPassword))

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
