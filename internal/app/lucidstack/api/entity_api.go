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

type EntityApi struct {
	router        *gin.Engine
	authenticator *auth.Authenticator
	entityService *service.EntityService
}

func NewEntityApi(router *gin.Engine, authenticator *auth.Authenticator, entityService *service.EntityService) *EntityApi {
	return &EntityApi{router: router, authenticator: authenticator, entityService: entityService}
}

func (a *EntityApi) Register() {

	a.router.POST("/api/v1/entities", func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c, 5*time.Second)
		defer cancel()

		au, err := a.authenticator.ValidateUserContext(c)
		if err != nil {
			api.Error(c, http.StatusUnauthorized, err)
			return
		}

		var req request.EntityCreationRequest

		if err := c.ShouldBindJSON(&req); err != nil {
			api.Error(c, http.StatusBadRequest, err)
			return
		}

		entity, err := a.entityService.Create(ctx, &req, au.ID, au.OrganizationID)

		if err != nil {
			api.Error(c, http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, entity)
	})

	a.router.GET("/api/v1/models/:modelID/environments/:environmentID/entities", func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c, 5*time.Second)
		defer cancel()

		au, err := a.authenticator.ValidateUserContext(c)
		if err != nil {
			api.Error(c, http.StatusUnauthorized, err)
			return
		}

		modelID := c.Param("modelID")
		environmentID := c.Param("environmentID")
		page, size := api.Page(c)

		entities, err := a.entityService.ListForModel(ctx, modelID, environmentID, au.OrganizationID, page, size)

		if err != nil {
			api.Error(c, http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, entities)
	})

	a.router.GET("/api/v1/environments/:environmentID/entities/:entityID/children", func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c, 5*time.Second)
		defer cancel()

		au, err := a.authenticator.ValidateUserContext(c)
		if err != nil {
			api.Error(c, http.StatusUnauthorized, err)
			return
		}

		entityID := c.Param("entityID")
		environmentID := c.Param("environmentID")
		page, size := api.Page(c)

		entities, err := a.entityService.ListForParent(ctx, entityID, environmentID, au.OrganizationID, page, size)

		if err != nil {
			api.Error(c, http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, entities)
	})

	a.router.GET("/api/v1/entities/:entityID", func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c, 5*time.Second)
		defer cancel()

		au, err := a.authenticator.ValidateUserContext(c)
		if err != nil {
			api.Error(c, http.StatusUnauthorized, err)
			return
		}

		entityID := c.Param("entityID")

		entity, err := a.entityService.Get(ctx, entityID, au.OrganizationID)

		if err != nil {
			api.Error(c, http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, entity)
	})

	a.router.PUT("/api/v1/entities/:entityID", func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c, 5*time.Second)
		defer cancel()

		au, err := a.authenticator.ValidateUserContext(c)
		if err != nil {
			api.Error(c, http.StatusUnauthorized, err)
			return
		}

		entityID := c.Param("entityID")
		var req request.EntityUpdateRequest

		if err := c.ShouldBindJSON(&req); err != nil {
			api.Error(c, http.StatusBadRequest, err)
			return
		}

		entity, err := a.entityService.Update(ctx, entityID, &req, au.OrganizationID)

		if err != nil {
			api.Error(c, http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, entity)
	})

	a.router.DELETE("/api/v1/entities/:entityID", func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c, 5*time.Second)
		defer cancel()

		au, err := a.authenticator.ValidateUserContext(c)
		if err != nil {
			api.Error(c, http.StatusUnauthorized, err)
			return
		}

		entityID := c.Param("entityID")

		entity, err := a.entityService.Delete(ctx, entityID, au.OrganizationID)

		if err != nil {
			api.Error(c, http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, entity)
	})

	a.router.POST("/api/v1/entities/:entityID/parents", func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c, 5*time.Second)
		defer cancel()

		au, err := a.authenticator.ValidateUserContext(c)
		if err != nil {
			api.Error(c, http.StatusUnauthorized, err)
			return
		}

		entityID := c.Param("entityID")
		var req request.EntityParentRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			api.Error(c, http.StatusBadRequest, err)
			return
		}

		err = a.entityService.AddParent(ctx, entityID, &req, au.OrganizationID)

		if err != nil {
			api.Error(c, http.StatusInternalServerError, err)
			return
		}

		api.Success(c, http.StatusOK, "parent added successfully")
	})

	a.router.DELETE("/api/v1/entities/:entityID/parents", func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c, 5*time.Second)
		defer cancel()

		au, err := a.authenticator.ValidateUserContext(c)
		if err != nil {
			api.Error(c, http.StatusUnauthorized, err)
			return
		}

		entityID := c.Param("entityID")
		var req request.EntityParentRequest

		if err := c.ShouldBindJSON(&req); err != nil {
			api.Error(c, http.StatusBadRequest, err)
			return
		}

		err = a.entityService.DeleteParent(ctx, entityID, &req, au.OrganizationID)

		if err != nil {
			api.Error(c, http.StatusInternalServerError, err)
			return
		}

		api.Success(c, http.StatusOK, "parent deleted successfully")
	})

	a.router.GET("/api/v1/entities/:entityID/parents", func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c, 5*time.Second)
		defer cancel()

		au, err := a.authenticator.ValidateUserContext(c)
		if err != nil {
			api.Error(c, http.StatusUnauthorized, err)
			return
		}

		entityID := c.Param("entityID")

		parents, err := a.entityService.ListParents(ctx, entityID, au.OrganizationID)

		if err != nil {
			api.Error(c, http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, parents)
	})
}
