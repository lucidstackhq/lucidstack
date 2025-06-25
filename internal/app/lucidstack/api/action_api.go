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

type ActionApi struct {
	router        *gin.Engine
	authenticator *auth.Authenticator
	actionService *service.ActionService
}

func NewActionApi(router *gin.Engine, authenticator *auth.Authenticator, actionService *service.ActionService) *ActionApi {
	return &ActionApi{router: router, authenticator: authenticator, actionService: actionService}
}

func (a *ActionApi) Register() {

	a.router.POST("/api/v1/models/:modelID/actions", func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c, 5*time.Second)
		defer cancel()

		au, err := a.authenticator.ValidateUserContext(c)
		if err != nil {
			api.Error(c, http.StatusUnauthorized, err)
			return
		}

		modelID := c.Param("modelID")

		var req request.ActionCreationRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			api.Error(c, http.StatusBadRequest, err)
			return
		}

		action, err := a.actionService.Create(ctx, modelID, &req, au.ID, au.OrganizationID)

		if err != nil {
			api.Error(c, http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusCreated, action)
	})

	a.router.GET("/api/v1/models/:modelID/actions", func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c, 5*time.Second)
		defer cancel()

		au, err := a.authenticator.ValidateUserContext(c)
		if err != nil {
			api.Error(c, http.StatusUnauthorized, err)
			return
		}

		modelID := c.Param("modelID")
		page, size := api.Page(c)

		actions, err := a.actionService.List(ctx, modelID, au.OrganizationID, page, size)

		if err != nil {
			api.Error(c, http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, actions)
	})

	a.router.GET("/api/v1/models/:modelID/actions/:actionID", func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c, 5*time.Second)
		defer cancel()

		au, err := a.authenticator.ValidateUserContext(c)

		if err != nil {
			api.Error(c, http.StatusUnauthorized, err)
			return
		}

		modelID := c.Param("modelID")
		actionID := c.Param("actionID")

		action, err := a.actionService.Get(ctx, actionID, modelID, au.OrganizationID)

		if err != nil {
			api.Error(c, http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, action)
	})

	a.router.PUT("/api/v1/models/:modelID/actions/:actionID", func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c, 5*time.Second)
		defer cancel()

		au, err := a.authenticator.ValidateUserContext(c)
		if err != nil {
			api.Error(c, http.StatusUnauthorized, err)
			return
		}

		modelID := c.Param("modelID")
		actionID := c.Param("actionID")

		var req request.ActionUpdateRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			api.Error(c, http.StatusBadRequest, err)
			return
		}

		action, err := a.actionService.Update(ctx, actionID, &req, modelID, au.OrganizationID)

		if err != nil {
			api.Error(c, http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, action)
	})

	a.router.DELETE("/api/v1/models/:modelID/actions/:actionID", func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c, 5*time.Second)
		defer cancel()

		au, err := a.authenticator.ValidateUserContext(c)
		if err != nil {
			api.Error(c, http.StatusUnauthorized, err)
			return
		}

		modelID := c.Param("modelID")
		actionID := c.Param("actionID")

		action, err := a.actionService.Delete(ctx, actionID, modelID, au.OrganizationID)
		if err != nil {
			api.Error(c, http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, action)
	})
}
