package router

import (
	"github.com/gin-gonic/gin"
	"github.com/orto-core/server/api-gateway/internal/utils"
	"github.com/spf13/viper"
)

func RegisterRouter(rg *gin.RouterGroup) {
	authServiceURL := viper.GetString("services.auth_service.url")
	tenantServiceURL := viper.GetString("services.tenant_service.url")
	pageServiceURL := viper.GetString("services.page_service.url")

	rg.Any("/api/auth/*proxyPath", utils.CreateReverseProxy(authServiceURL))
	rg.Any("/api/tenant/*proxyPath", utils.CreateReverseProxy(tenantServiceURL))
	rg.Any("/api/page/*proxyPath", utils.CreateReverseProxy(pageServiceURL))
}
