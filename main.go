package main

import (
	"github.com/freelifer/coolgo/config"
	_ "github.com/freelifer/coolgo/dao/db"
	"github.com/freelifer/coolgo/dao/redis"
	"log"
	"os"
)

// GOOS=linux GOARCH=amd64 go build -o coolgo_linux github.com/freelifer/coolgo/*.go
// GOOS=windows GOARCH=amd64 go build -o coolgo_win github.com/freelifer/coolgo/*.go

// GOOS=linux GOARCH=amd64 go build -o coolgo_linux *.go
// GOOS=windows GOARCH=amd64 go build -o coolgo_win *.go

//go test -bench=WeiXinLogin github.com/freelifer/coolgo/service
func main() {
	redis_status, _ := config.Config.Bool("app::redis_status")
	// redis init
	err := redis.NewRedisConn(redis_status)
	if err != nil {
		log.Println("redis new err", err)
		os.Exit(0)
	}

	r := initRouter()
	r.Run() // listen and serve on 0.0.0.0:8080
	// r.Run(":8000")
}
