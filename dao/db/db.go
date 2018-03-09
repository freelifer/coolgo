package db

import (
	"github.com/freelifer/coolgo/dao"
	"github.com/freelifer/coolgo/pkg/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	_ "github.com/mattn/go-sqlite3"
)

var dbDriver string
var engine *xorm.Engine

func init() {
	dbDriver := config.String("app::db")
	if "mysql" == dbDriver {
		engine, _ = xorm.NewEngine("mysql", "root:123@/test?charset=utf8")
	} else {
		//sqlite3
		engine, _ = xorm.NewEngine("sqlite3", "./data.db")
	}

	// engine.Logger().SetLevel(core.LOG_DEBUG)
	engine.ShowSQL(true)

	engine.Sync2(new(dao.User), new(dao.Role), new(dao.Permission), new(dao.Bill))
}

func CreateUser(openid string) (int64, error) {
	user := new(dao.User)
	user.WxOpenid = openid
	affected, err := engine.Insert(user)
	return affected, err
}

func FindUser(openid string) (*dao.User, error) {
	user := &dao.User{WxOpenid: openid}
	_, err := engine.Get(user)
	return user, err
}

func FindUser1(openid string) (*dao.User, bool, error) {
	user := &dao.User{WxOpenid: openid}
	has, err := engine.Get(user)
	return user, has, err
}
