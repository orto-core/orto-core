package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/orto-core/server/auth-service/internal/models"
	"github.com/orto-core/server/auth-service/internal/service"
)

type AuthController interface {
	Login(*gin.Context)
	Register(*gin.Context)
}

type authController struct {
	service service.AuthService
}

func NewAuthController(authService service.AuthService) AuthController {
	return &authController{
		service: authService,
	}
}

func (c *authController) Login(ctx *gin.Context) {
	body := models.User{}
	if err := ctx.Bind(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": "provide username or password",
		})
		return
	}

	resp, err := c.service.Login(&body)
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

func (c *authController) Register(ctx *gin.Context) {
	body := models.User{}
	if err := ctx.Bind(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": "provide username or password",
		})
		return
	}

	resp, err := c.service.Register(&body)
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
