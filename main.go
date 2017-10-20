package main

import (
	_ "github.com/freelifer/coolgo/config"
)

func main() {
	r := initRouter()
	r.Run() // listen and serve on 0.0.0.0:8080
}
