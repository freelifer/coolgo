package redis

import (
	"errors"
	"github.com/garyburd/redigo/redis"
)

var REDIS_STATUS_CLOSE error = errors.New("redis status close")
var c redis.Conn
var EXPIRE_TIME int = 1 * 24 * 60 * 60
var COUPONS_EXPIRE_TIME = 4 * 60 * 60

func NewRedisConn(status bool) error {
	if status {
		return nil
	}
	var err error
	c, err = redis.Dial("tcp", "localhost:6379")
	// defer c.Close()

	return err
}

func SetSessionExpireTime(time int) {
	EXPIRE_TIME = time
}
func GetSession(sessionId string) (string, error) {
	if c == nil {
		return "", REDIS_STATUS_CLOSE
	}
	value, err := redis.String(c.Do("GET", sessionId))
	return value, err
}

func PutSession(sessionId, value string) error {
	if c == nil {
		return REDIS_STATUS_CLOSE
	}
	_, err := c.Do("SET", sessionId, value, "EX", EXPIRE_TIME)
	return err
}

func Set(key, value string) error {
	if c == nil {
		return REDIS_STATUS_CLOSE
	}
	_, err := c.Do("SET", key, value)
	return err
}

func Get(key string) (string, error) {
	if c == nil {
		return "", REDIS_STATUS_CLOSE
	}
	value, err := redis.String(c.Do("GET", key))
	return value, err
}

func PutCoupons(key, value string) error {
	if c == nil {
		return REDIS_STATUS_CLOSE
	}
	// _, err = c.Do("SET", "password", "123456", "EX", "10")
	_, err := c.Do("SET", key, value, "EX", COUPONS_EXPIRE_TIME)
	return err
}
