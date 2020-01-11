package web

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go-module/web/handlers"
)

type Config struct {
	Port int
}

type Server struct {
	Config *Config
}

func (s *Server) Run() error {
	if s.Config == nil {
		return fmt.Errorf("no configuration provided")
	}

	if os.Getenv("DEBUG") != "true" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	mediaGroup := router.Group("/media")
	{
		mediaGroup.GET("/:id", handlers.GetMedia)
		mediaGroup.POST("/", handlers.CreateMedia)
	}


	httpServer := &http.Server{
		Addr:           fmt.Sprintf(":%d", s.Config.Port),
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	return httpServer.ListenAndServe()
}
