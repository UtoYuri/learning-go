package main

import (
	"fmt"
	_ "go-module/docs"
	"go-module/pkg/config"
	"go-module/web"
	"os"
	"strconv"
)

// @title Gin swagger
// @version 1.0
// @description Learning Go

// @contact.name yuri
// @contact.url https://utohub.com
// @contact.email yuri@utohub.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	if err := config.Init(); err != nil {
		fmt.Println("Init failed,", err)
		return
	}

	var port int64
	port, err := strconv.ParseInt(os.Getenv("PORT"), 10, 64)
	if err != nil {
		port = 8000
	}

	server := &web.Server{
		Config: &web.Config{
			Port: int(port),
		},
	}
	if err := server.Run(); err != nil {
		fmt.Println("Run server failed,", err)
	}
}
