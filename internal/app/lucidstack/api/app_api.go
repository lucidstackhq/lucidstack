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

type AppApi struct {
	router        *gin.Engine
	authenticator *auth.Authenticator
	appService    *service.AppService
}

func NewAppApi(router *gin.Engine, authenticator *auth.Authenticator, appService *service.AppService) *AppApi {
	return &AppApi{router: router, authenticator: authenticator, appService: appService}
}

func (a *AppApi) Register() {

	a.router.POST("/api/v1/apps", func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c, 5*time.Second)
		defer cancel()

		au, err := a.authenticator.ValidateUserContext(c)
		if err != nil {
			api.Error(c, http.StatusUnauthorized, err)
			return
		}

		var req request.AppCreationRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			api.Error(c, http.StatusBadRequest, err)
			return
		}

		app, err := a.appService.Create(ctx, &req, au.ID, au.OrganizationID)

		if err != nil {
			api.Error(c, http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusCreated, app)
	})

	a.router.GET("/api/v1/apps", func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c, 5*time.Second)
		defer cancel()

		au, err := a.authenticator.ValidateUserContext(c)
		if err != nil {
			api.Error(c, http.StatusUnauthorized, err)
			return
		}

		page, size := api.Page(c)

		apps, err := a.appService.List(ctx, au.OrganizationID, page, size)

		if err != nil {
			api.Error(c, http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, apps)
	})

	a.router.GET("/api/v1/apps/:appID", func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c, 5*time.Second)
		defer cancel()

		au, err := a.authenticator.ValidateUserContext(c)
		if err != nil {
			api.Error(c, http.StatusUnauthorized, err)
			return
		}

		appID := c.Param("appID")

		app, err := a.appService.Get(ctx, appID, au.OrganizationID)

		if err != nil {
			api.Error(c, http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, app)
	})

	a.router.PUT("/api/v1/apps/:appID", func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c, 5*time.Second)
		defer cancel()

		au, err := a.authenticator.ValidateUserContext(c)
		if err != nil {
			api.Error(c, http.StatusUnauthorized, err)
			return
		}

		appID := c.Param("appID")

		var req request.AppUpdateRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			api.Error(c, http.StatusBadRequest, err)
			return
		}

		app, err := a.appService.Update(ctx, appID, &req, au.OrganizationID)

		if err != nil {
			api.Error(c, http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, app)
	})

	a.router.DELETE("/api/v1/apps/:appID", func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c, 5*time.Second)
		defer cancel()

		au, err := a.authenticator.ValidateUserContext(c)
		if err != nil {
			api.Error(c, http.StatusUnauthorized, err)
			return
		}

		appID := c.Param("appID")

		app, err := a.appService.Delete(ctx, appID, au.OrganizationID)

		if err != nil {
			api.Error(c, http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, app)
	})

	a.router.PUT("/api/v1/apps/:appID/secret", func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c, 5*time.Second)
		defer cancel()

		au, err := a.authenticator.ValidateUserContext(c)
		if err != nil {
			api.Error(c, http.StatusUnauthorized, err)
			return
		}

		appID := c.Param("appID")

		app, err := a.appService.ResetSecret(ctx, appID, au.OrganizationID)

		if err != nil {
			api.Error(c, http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, app)
	})

	a.router.GET("/api/v1/apps/:appID/secret", func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c, 5*time.Second)
		defer cancel()

		au, err := a.authenticator.ValidateUserContext(c)
		if err != nil {
			api.Error(c, http.StatusUnauthorized, err)
			return
		}

		appID := c.Param("appID")

		app, err := a.appService.GetSecret(ctx, appID, au.OrganizationID)

		if err != nil {
			api.Error(c, http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, app)
	})
}
