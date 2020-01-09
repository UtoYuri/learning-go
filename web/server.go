package web

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-module/web/handlers"
	"net/http"
	"time"
)

type Config struct {
	Port int
}

type Server struct {
	Config *Config
}

func (s *Server) Run() error {
	var port int

	if s.Config == nil {
		port = s.Config.Port
	}

	if port == 0 {
		port = 8000
	}


	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	getGroup := router.Group("/get")
	getGroup.GET("/params", handlers.GetByParams)
	getGroup.GET("/params/:field_a/:field_b", handlers.GetByUriParams)

	httpServer := &http.Server{
		Addr:           fmt.Sprintf(":%d", port),
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	return httpServer.ListenAndServe()
}
