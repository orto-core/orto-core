package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/orto-core/server/tenant-service/internal/models"
	"github.com/orto-core/server/tenant-service/internal/service"
)

type TenantController interface {
	AddTenant(*gin.Context)
}

type tenantController struct {
	service service.TenantService
}

func NewTenantController(tenantService service.TenantService) TenantController {
	return &tenantController{
		service: tenantService,
	}
}

func (c *tenantController) AddTenant(ctx *gin.Context) {
	body := models.Tenant{}
	if err := ctx.Bind(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": "provide tenant name",
		})
		return
	}

	resp, err := c.service.AddTenant(&body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": resp,
	})
}
