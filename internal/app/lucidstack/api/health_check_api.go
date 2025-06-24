package api

import (
	"github.com/gin-gonic/gin"
	"github.com/lucidstackhq/lucidstack/internal/app/lucidstack/response"
	"net/http"
)

type HealthCheckApi struct {
	router *gin.Engine
}

func NewHealthCheckApi(router *gin.Engine) *HealthCheckApi {
	return &HealthCheckApi{router: router}
}

func (a *HealthCheckApi) Register() {
	a.router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, &response.HealthCheckResponse{Status: "ok"})
	})
}
