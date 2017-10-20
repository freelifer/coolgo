package config

import (
	"fmt"
	"github.com/astaxie/beego/config"
	"github.com/freelifer/coolgo/utils"
	"os"
	"path/filepath"
)

var Config config.Configer

func init() {
	var err error
	var AppPath string
	if AppPath, err = filepath.Abs(filepath.Dir(os.Args[0])); err != nil {
		panic(err)
	}
	workPath, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	var appConfigPath = filepath.Join(workPath, "github.com/freelifer/coolgo/conf", "app.conf")
	fmt.Println(appConfigPath)
	if !utils.FileExists(appConfigPath) {
		appConfigPath = filepath.Join(AppPath, "conf", "app.conf")
		fmt.Println(appConfigPath)
		if !utils.FileExists(appConfigPath) {
			return
		}
	}
	Config, err = config.NewConfig("ini", appConfigPath)

	appName := Config.String("app")

	fmt.Println("appName", appName)
	if err != nil {
		panic(err.Error())
	}
}
