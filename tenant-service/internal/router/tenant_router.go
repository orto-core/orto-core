package router

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/orto-core/server/tenant-service/internal/controller"
	"github.com/orto-core/server/tenant-service/internal/repository"
	"github.com/orto-core/server/tenant-service/internal/service"
	"github.com/orto-core/server/tenant-service/internal/store"
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

	tenantRepository := repository.NewTenantRepository(store.DB)
	tenantService := service.NewTenantService(tenantRepository)
	tenantController := controller.NewTenantController(tenantService)

	r.POST("/tenant", tenantController.AddTenant)

	return r
}
