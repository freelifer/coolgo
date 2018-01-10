package redis

import (
	"testing"
)

func Test_NewRedisConn(t *testing.T) {
	// err := NewRedisConn()

	// if err != nil {
	// 	t.Error(err)
	// } else {
	// 	t.Log("NewRedisConn 通过")
	// }
}

func Test_Set(t *testing.T) {

	err := Set("key", "value")

	if err != nil {
		t.Error(err)
	} else {
		t.Log("put Success")
	}
}
