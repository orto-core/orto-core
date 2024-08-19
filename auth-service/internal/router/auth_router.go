package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/orto-core/server/auth-service/internal/controller"
	"github.com/orto-core/server/auth-service/internal/repository"
	"github.com/orto-core/server/auth-service/internal/service"
	"github.com/orto-core/server/auth-service/internal/store"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func RegisterRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/_status/healthz", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "healthy",
		})
	})

	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

	authRepository := repository.NewAuthRepository(store.DB)
	authService := service.NewAuthService(authRepository)
	authController := controller.NewAuthController(authService)

	r.POST("/login", authController.Login)
	r.POST("/register", authController.Register)

	return r
}
