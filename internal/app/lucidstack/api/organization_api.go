package api

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/lucidstackhq/lucidstack/internal/app/lucidstack/request"
	"github.com/lucidstackhq/lucidstack/internal/app/lucidstack/service"
	"github.com/lucidstackhq/lucidstack/internal/pkg/api"
	"github.com/lucidstackhq/lucidstack/internal/pkg/auth"
	"net/http"
	"time"
)

type OrganizationApi struct {
	router              *gin.Engine
	authenticator       *auth.Authenticator
	organizationService *service.OrganizationService
}

func NewOrganizationApi(router *gin.Engine, authenticator *auth.Authenticator, organizationService *service.OrganizationService) *OrganizationApi {
	return &OrganizationApi{router: router, authenticator: authenticator, organizationService: organizationService}
}

func (a *OrganizationApi) Register() {

	a.router.GET("/api/v1/organization", func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c, 5*time.Second)
		defer cancel()

		au, err := a.authenticator.ValidateUserContext(c)
		if err != nil {
			api.Error(c, http.StatusUnauthorized, err)
			return
		}

		organization, err := a.organizationService.Get(ctx, au.OrganizationID)

		if err != nil {
			api.Error(c, http.StatusUnauthorized, err)
			return
		}

		c.JSON(http.StatusOK, organization)
	})

	a.router.PUT("/api/v1/organization", func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c, 5*time.Second)
		defer cancel()

		au, err := a.authenticator.ValidateUserContext(c)
		if err != nil {
			api.Error(c, http.StatusUnauthorized, err)
			return
		}

		if !au.Admin {
			api.Forbidden(c)
			return
		}

		var req request.OrganizationUpdateRequest

		if err := c.ShouldBindJSON(&req); err != nil {
			api.Error(c, http.StatusBadRequest, err)
			return
		}

		organization, err := a.organizationService.Update(ctx, au.OrganizationID, &req)

		if err != nil {
			api.Error(c, http.StatusUnauthorized, err)
			return
		}

		c.JSON(http.StatusOK, organization)
	})
}
