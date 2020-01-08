package main

import (
	"go-module/web"
)

func main() {
	server := &web.Server{}
	server.Run()
}
