package config

import (
	"github.com/Unknwon/goconfig"
	"strings"
)

var c *goconfig.ConfigFile

func init() {
	c, _ = goconfig.LoadConfigFile("conf/app.conf")
}

func MustValue(section, key string) string {
	if c != nil {
		return c.MustValue(section, key, "")
	}
	return ""
}

func MustBool(section, key string) bool {
	if c != nil {
		return c.MustBool(section, key, false)
	}
	return false
}

func String(key string) string {
	pos := strings.Index(key, "::")
	if pos > -1 {
		section := key[:pos]
		next := key[pos+2:]
		return MustValue(section, next)
	}

	return ""
}

func Bool(key string) bool {
	pos := strings.Index(key, "::")
	if pos > -1 {
		section := key[:pos]
		next := key[pos+2:]
		return MustBool(section, next)
	}

	return false
}
