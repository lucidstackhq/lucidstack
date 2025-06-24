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

type ModelApi struct {
	router        *gin.Engine
	authenticator *auth.Authenticator
	modelService  *service.ModelService
}

func NewModelApi(router *gin.Engine, authenticator *auth.Authenticator, modelService *service.ModelService) *ModelApi {
	return &ModelApi{router: router, authenticator: authenticator, modelService: modelService}
}

func (a *ModelApi) Register() {

	a.router.POST("/api/v1/models", func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c, 5*time.Second)
		defer cancel()

		au, err := a.authenticator.ValidateUserContext(c)
		if err != nil {
			api.Error(c, http.StatusUnauthorized, err)
			return
		}

		var req request.ModelCreationRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			api.Error(c, http.StatusBadRequest, err)
			return
		}

		model, err := a.modelService.Create(ctx, &req, au.ID, au.OrganizationID)

		if err != nil {
			api.Error(c, http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusCreated, model)
	})

	a.router.GET("/api/v1/models", func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c, 5*time.Second)
		defer cancel()

		au, err := a.authenticator.ValidateUserContext(c)
		if err != nil {
			api.Error(c, http.StatusUnauthorized, err)
			return
		}

		page, size := api.Page(c)

		models, err := a.modelService.List(ctx, au.OrganizationID, page, size)

		if err != nil {
			api.Error(c, http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, models)
	})

	a.router.GET("/api/v1/models/:modelID", func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c, 5*time.Second)
		defer cancel()

		au, err := a.authenticator.ValidateUserContext(c)
		if err != nil {
			api.Error(c, http.StatusUnauthorized, err)
			return
		}

		modelID := c.Param("modelID")

		model, err := a.modelService.Get(ctx, modelID, au.OrganizationID)

		if err != nil {
			api.Error(c, http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, model)
	})

	a.router.PUT("/api/v1/models/:modelID", func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c, 5*time.Second)
		defer cancel()

		au, err := a.authenticator.ValidateUserContext(c)
		if err != nil {
			api.Error(c, http.StatusUnauthorized, err)
			return
		}

		modelID := c.Param("modelID")

		var req request.ModelUpdateRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			api.Error(c, http.StatusBadRequest, err)
			return
		}

		model, err := a.modelService.Update(ctx, modelID, &req, au.OrganizationID)

		if err != nil {
			api.Error(c, http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, model)
	})

	a.router.DELETE("/api/v1/models/:modelID", func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c, 5*time.Second)
		defer cancel()

		au, err := a.authenticator.ValidateUserContext(c)
		if err != nil {
			api.Error(c, http.StatusUnauthorized, err)
			return
		}

		modelID := c.Param("modelID")

		model, err := a.modelService.Delete(ctx, modelID, au.OrganizationID)

		if err != nil {
			api.Error(c, http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, model)
	})
}
