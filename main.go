package main

import ()

func main() {

	r := initRouter()
	r.Run() // listen and serve on 0.0.0.0:8080
}
