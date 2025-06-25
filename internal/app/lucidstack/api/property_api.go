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

type PropertyApi struct {
	router          *gin.Engine
	authenticator   *auth.Authenticator
	propertyService *service.PropertyService
}

func NewPropertyApi(router *gin.Engine, authenticator *auth.Authenticator, propertyService *service.PropertyService) *PropertyApi {
	return &PropertyApi{router: router, authenticator: authenticator, propertyService: propertyService}
}

func (a *PropertyApi) Register() {

	a.router.POST("/api/v1/models/:modelID/properties", func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c, 5*time.Second)
		defer cancel()

		au, err := a.authenticator.ValidateUserContext(c)
		if err != nil {
			api.Error(c, http.StatusUnauthorized, err)
			return
		}

		modelID := c.Param("modelID")

		var req request.PropertyCreationRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			api.Error(c, http.StatusBadRequest, err)
			return
		}

		property, err := a.propertyService.Create(ctx, modelID, &req, au.ID, au.OrganizationID)

		if err != nil {
			api.Error(c, http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusCreated, property)
	})

	a.router.GET("/api/v1/models/:modelID/properties", func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c, 5*time.Second)
		defer cancel()

		au, err := a.authenticator.ValidateUserContext(c)
		if err != nil {
			api.Error(c, http.StatusUnauthorized, err)
			return
		}

		modelID := c.Param("modelID")
		page, size := api.Page(c)

		properties, err := a.propertyService.List(ctx, modelID, au.OrganizationID, page, size)

		if err != nil {
			api.Error(c, http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, properties)
	})

	a.router.GET("/api/v1/models/:modelID/properties/:propertyID", func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c, 5*time.Second)
		defer cancel()

		au, err := a.authenticator.ValidateUserContext(c)
		if err != nil {
			api.Error(c, http.StatusUnauthorized, err)
			return
		}

		modelID := c.Param("modelID")
		propertyID := c.Param("propertyID")

		property, err := a.propertyService.Get(ctx, propertyID, modelID, au.OrganizationID)

		if err != nil {
			api.Error(c, http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, property)
	})

	a.router.PUT("/api/v1/models/:modelID/properties/:propertyID", func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c, 5*time.Second)
		defer cancel()

		au, err := a.authenticator.ValidateUserContext(c)
		if err != nil {
			api.Error(c, http.StatusUnauthorized, err)
			return
		}

		modelID := c.Param("modelID")
		propertyID := c.Param("propertyID")

		var req request.PropertyUpdateRequest

		if err := c.ShouldBindJSON(&req); err != nil {
			api.Error(c, http.StatusBadRequest, err)
			return
		}

		property, err := a.propertyService.Update(ctx, propertyID, &req, modelID, au.OrganizationID)

		if err != nil {
			api.Error(c, http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, property)
	})

	a.router.DELETE("/api/v1/models/:modelID/properties/:propertyID", func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c, 5*time.Second)
		defer cancel()

		au, err := a.authenticator.ValidateUserContext(c)
		if err != nil {
			api.Error(c, http.StatusUnauthorized, err)
			return
		}

		modelID := c.Param("modelID")
		propertyID := c.Param("propertyID")

		property, err := a.propertyService.Delete(ctx, propertyID, modelID, au.OrganizationID)

		if err != nil {
			api.Error(c, http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, property)
	})

	a.router.PUT("/api/v1/models/:modelID/properties/:propertyID/default-value", func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c, 5*time.Second)
		defer cancel()

		au, err := a.authenticator.ValidateUserContext(c)
		if err != nil {
			api.Error(c, http.StatusUnauthorized, err)
			return
		}

		modelID := c.Param("modelID")
		propertyID := c.Param("propertyID")

		var req request.PropertyDefaultValueRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			api.Error(c, http.StatusBadRequest, err)
			return
		}

		property, err := a.propertyService.UpdateDefaultValue(ctx, propertyID, &req, modelID, au.OrganizationID)

		if err != nil {
			api.Error(c, http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, property)
	})
}
