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

type EntityPropertyApi struct {
	router                *gin.Engine
	authenticator         *auth.Authenticator
	entityPropertyService *service.EntityPropertyService
}

func NewEntityPropertyApi(router *gin.Engine, authenticator *auth.Authenticator, entityPropertyService *service.EntityPropertyService) *EntityPropertyApi {
	return &EntityPropertyApi{router: router, authenticator: authenticator, entityPropertyService: entityPropertyService}
}

func (a *EntityPropertyApi) Register() {

	a.router.POST("/api/v1/entities/:entityID/properties/:propertyID/instances", func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c, 5*time.Second)
		defer cancel()

		au, err := a.authenticator.ValidateUserContext(c)
		if err != nil {
			api.Error(c, http.StatusUnauthorized, err)
			return
		}

		entityID := c.Param("entityID")
		propertyID := c.Param("propertyID")

		var req request.EntityPropertyRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			api.Error(c, http.StatusBadRequest, err)
			return
		}

		entityProperty, err := a.entityPropertyService.Create(ctx, entityID, propertyID, &req, au.ID, au.OrganizationID)

		if err != nil {
			api.Error(c, http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusCreated, entityProperty)
	})

	a.router.GET("/api/v1/entities/:entityID/properties/:propertyID/instances", func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c, 5*time.Second)
		defer cancel()

		au, err := a.authenticator.ValidateUserContext(c)
		if err != nil {
			api.Error(c, http.StatusUnauthorized, err)
			return
		}

		entityID := c.Param("entityID")
		propertyID := c.Param("propertyID")
		page, size := api.Page(c)

		entityProperties, err := a.entityPropertyService.List(ctx, entityID, propertyID, au.OrganizationID, page, size)

		if err != nil {
			api.Error(c, http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, entityProperties)
	})

	a.router.GET("/api/v1/entities/:entityID/properties/:propertyID/instances/:entityPropertyID", func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c, 5*time.Second)
		defer cancel()

		au, err := a.authenticator.ValidateUserContext(c)
		if err != nil {
			api.Error(c, http.StatusUnauthorized, err)
			return
		}

		entityID := c.Param("entityID")
		propertyID := c.Param("propertyID")
		entityPropertyID := c.Param("entityPropertyID")

		entityProperty, err := a.entityPropertyService.Get(ctx, entityPropertyID, entityID, propertyID, au.OrganizationID)

		if err != nil {
			api.Error(c, http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, entityProperty)
	})

	a.router.PUT("/api/v1/entities/:entityID/properties/:propertyID/instances/:entityPropertyID", func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c, 5*time.Second)
		defer cancel()

		au, err := a.authenticator.ValidateUserContext(c)
		if err != nil {
			api.Error(c, http.StatusUnauthorized, err)
			return
		}

		entityID := c.Param("entityID")
		propertyID := c.Param("propertyID")
		entityPropertyID := c.Param("entityPropertyID")

		var req request.EntityPropertyRequest

		if err := c.ShouldBindJSON(&req); err != nil {
			api.Error(c, http.StatusBadRequest, err)
			return
		}

		entityProperty, err := a.entityPropertyService.Update(ctx, entityPropertyID, &req, entityID, propertyID, au.OrganizationID)

		if err != nil {
			api.Error(c, http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, entityProperty)
	})

	a.router.DELETE("/api/v1/entities/:entityID/properties/:propertyID/instances/:entityPropertyID", func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c, 5*time.Second)
		defer cancel()

		au, err := a.authenticator.ValidateUserContext(c)
		if err != nil {
			api.Error(c, http.StatusUnauthorized, err)
			return
		}

		entityID := c.Param("entityID")
		propertyID := c.Param("propertyID")
		entityPropertyID := c.Param("entityPropertyID")

		entityProperty, err := a.entityPropertyService.Delete(ctx, entityPropertyID, entityID, propertyID, au.OrganizationID)

		if err != nil {
			api.Error(c, http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, entityProperty)
	})
}
