package web

import (
	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
}

func (s *Server) Run() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
