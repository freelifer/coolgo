package main

import (
	_ "github.com/freelifer/coolgo/config"
	"github.com/freelifer/coolgo/dao/redis"
	"os"
)

// GOOS=linux GOARCH=amd64 go build -o coolgo_linux github.com/freelifer/coolgo/*.go
// GOOS=window GOARCH=amd64 go build -o coolgo_win github.com/freelifer/coolgo/*.go
func main() {
	// redis init
	err := redis.NewRedisConn()
	if err != nil {
		os.Exit(0)
	}
	r := initRouter()
	r.Run() // listen and serve on 0.0.0.0:8080
	// r.Run(":8000")
}
