package redis

import (
	"testing"
)

func Test_init(t *testing.T) {

}

func Test_Set(t *testing.T) {
	err := Set("key", "value")

	if err != nil {
		t.Error(err)
	} else {
		t.Log("put Success")
	}
}
