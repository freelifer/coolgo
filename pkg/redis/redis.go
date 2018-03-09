package redis

import (
	"errors"
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
)

func RedisInit(redisHost string, redisDb int, redisMaxidle int, redisMaxActive int) {
	if redisMaxidle < 1 {
		redisMaxidle = 1
	}
	if redisMaxActive < 10 {
		redisMaxActive = 10
	}
	// 建立连接池
	RedisClient = &redis.Pool{
		// 从配置文件获取maxidle以及maxactive，取不到则用后面的默认值
		MaxIdle:     redisMaxidle,
		MaxActive:   redisMaxActive,
		IdleTimeout: 180 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", redisHost)
			if err != nil {
				return nil, err
			}
			// 选择db
			c.Do("SELECT", redisDb)
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
