package redis

import (
	"errors"
	"github.com/freelifer/coolgo/config"
	"github.com/garyburd/redigo/redis"
	"time"
)

var REDIS_STATUS_CLOSE error = errors.New("redis status close")
var c redis.Conn
var EXPIRE_TIME int = 1 * 24 * 60 * 60
var COUPONS_EXPIRE_TIME = 4 * 60 * 60

var (
	// 定义常量
	RedisClient *redis.Pool
	REDIS_HOST  string
	REDIS_DB    int
)

func init() {
	status := config.Bool("app::redis_status")
	if !status {
		return
	}
	// 从配置文件获取redis的ip以及db
	REDIS_HOST = config.String("redis::host")
	REDIS_DB = config.Int("redis::db")
	// 建立连接池
	RedisClient = &redis.Pool{
		// 从配置文件获取maxidle以及maxactive，取不到则用后面的默认值
		MaxIdle:     config.DefaultInt("redis::maxidle", 1),
		MaxActive:   config.DefaultInt("redis::maxactive", 10),
		IdleTimeout: 180 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", REDIS_HOST)
			if err != nil {
				return nil, err
			}
			// 选择db
			c.Do("SELECT", REDIS_DB)
			return c, nil
		},
	}
}

func SetSessionExpireTime(time int) {
	EXPIRE_TIME = time
}
func GetSession(sessionId string) (string, error) {
	if RedisClient == nil {
		return "", REDIS_STATUS_CLOSE
	}
	rc := RedisClient.Get()
	// 用完后将连接放回连接池
	defer rc.Close()
	value, err := redis.String(rc.Do("GET", sessionId))
	return value, err
}

func PutSession(sessionId, value string) error {
	if RedisClient == nil {
		return REDIS_STATUS_CLOSE
	}
	rc := RedisClient.Get()
	// 用完后将连接放回连接池
	defer rc.Close()
	_, err := rc.Do("SET", sessionId, value, "EX", EXPIRE_TIME)
	return err
}

func Set(key, value string) error {
	if RedisClient == nil {
		return REDIS_STATUS_CLOSE
	}
	rc := RedisClient.Get()
	// 用完后将连接放回连接池
	defer rc.Close()
	_, err := rc.Do("SET", key, value)
	return err
}

func Get(key string) (string, error) {
	if RedisClient == nil {
		return "", REDIS_STATUS_CLOSE
	}
	rc := RedisClient.Get()
	// 用完后将连接放回连接池
	defer rc.Close()
	value, err := redis.String(rc.Do("GET", key))
	return value, err
}

func PutCoupons(key, value string) error {
	if RedisClient == nil {
		return REDIS_STATUS_CLOSE
	}
	rc := RedisClient.Get()
	// 用完后将连接放回连接池
	defer rc.Close()
	// _, err = c.Do("SET", "password", "123456", "EX", "10")
	_, err := rc.Do("SET", key, value, "EX", COUPONS_EXPIRE_TIME)
	return err
}
