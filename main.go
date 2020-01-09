package main

import (
	"go-module/web"
)

func main() {
	server := &web.Server{
		Config: &web.Config{
			Port: 8000,
		},
	}
	server.Run()
}
