package main

import (
	_ "github.com/freelifer/coolgo/config"
)

// GOOS=linux GOARCH=amd64 go build -o coolgo_linux github.com/freelifer/coolgo/*.go
func main() {
	r := initRouter()
	r.Run() // listen and serve on 0.0.0.0:8080
	// r.Run(":8000")
}
