package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/orto-core/server/api-gateway/config"
	"github.com/orto-core/server/api-gateway/internal/router"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func init() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
}

func main() {
	if err := config.LoadConfig(); err != nil {
		panic(err)
	}

	r := gin.Default()

	r.GET("/_status/healthz", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "healthy",
		})
	})

	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

	rg := r.Group("v1")
	router.RegisterRouter(rg)

	port := fmt.Sprintf(":%s", os.Getenv("PORT"))

	if err := r.Run(port); err != nil {
		panic(err)
	}
}
