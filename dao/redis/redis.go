package redis

import (
	"github.com/garyburd/redigo/redis"
)

var c redis.Conn
var EXPIRE_TIME int = 1 * 24 * 60 * 60

func NewRedisConn() error {
	var err error
	c, err = redis.Dial("tcp", "localhost:6379")
	defer c.Close()

	return err
}

func SetSessionExpireTime(time int) {
	EXPIRE_TIME = time
}
func GetSession(sessionId string) (string, error) {
	value, err := redis.String(c.Do("GET", sessionId))
	return value, err
}

func PutSession(sessionId, value string) error {
	_, err = c.Do("SET", sessionId, value, "EX", EXPIRE_TIME)
	return err
}

func Set(key, value string) error {
	// _, err = c.Do("SET", "password", "123456", "EX", "10")
	_, err := c.Do("SET", "mykey", "superWang")
	return err
}

func Get(key string) (string, error) {
	value, err := redis.String(c.Do("GET", key))
	return value, err
}
