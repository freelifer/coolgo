package main

import (
	_ "github.com/freelifer/coolgo/models"
	. "github.com/freelifer/coolgo/routers"
)

// glide mirror set https://golang.org/x/mobile https://github.com/golang/mobile --vcs git
// glide mirror set https://golang.org/x/crypto https://github.com/golang/crypto --vcs git
// glide mirror set https://golang.org/x/net https://github.com/golang/net --vcs git
// glide mirror set https://golang.org/x/tools https://github.com/golang/tools --vcs git
// glide mirror set https://golang.org/x/text https://github.com/golang/text --vcs git
// glide mirror set https://golang.org/x/image https://github.com/golang/image --vcs git
// glide mirror set https://golang.org/x/sys https://github.com/golang/sys --vcs git
// glide mirror set https://golang.org/x/net/context https://github.com/golang/net/context --vcs git
// glide mirror set golang.org/x/net/context github.com/golang/net/context --vcs git

// GOOS=linux GOARCH=amd64 go build -o coolgo_linux github.com/freelifer/coolgo/*.go
// GOOS=windows GOARCH=amd64 go build -o coolgo_win github.com/freelifer/coolgo/*.go

// GOOS=linux GOARCH=amd64 go build -o coolgo_linux *.go
// GOOS=windows GOARCH=386 go build -o coolgo_win *.go
// GOOS=windows GOARCH=386 go build -ldflags "-s -w" -o coolgo_win.exe main.go

//go test -bench=WeiXinLogin github.com/freelifer/coolgo/service
// go test -run='Test_GetWxUserPasswords' *.go -v

// GOOS=windows GOARCH=386 go build -tags 'mysql' -o coolgo_win main.go
func main() {
	r := InitRouter()
	r.Run() // listen and serve on 0.0.0.0:8080
	// r.Run(":8000")
}
