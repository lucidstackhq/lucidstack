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

type EventApi struct {
	router        *gin.Engine
	authenticator *auth.Authenticator
	eventService  *service.EventService
}

func NewEventApi(router *gin.Engine, authenticator *auth.Authenticator, eventService *service.EventService) *EventApi {
	return &EventApi{router: router, authenticator: authenticator, eventService: eventService}
}

func (a *EventApi) Register() {

	a.router.POST("/api/v1/models/:modelID/events", func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c, 5*time.Second)
		defer cancel()

		au, err := a.authenticator.ValidateUserContext(c)
		if err != nil {
			api.Error(c, http.StatusUnauthorized, err)
			return
		}

		modelID := c.Param("modelID")
		var req request.EventCreationRequest

		if err := c.ShouldBindJSON(&req); err != nil {
			api.Error(c, http.StatusBadRequest, err)
			return
		}

		event, err := a.eventService.Create(ctx, modelID, &req, au.ID, au.OrganizationID)

		if err != nil {
			api.Error(c, http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusCreated, event)
	})

	a.router.GET("/api/v1/models/:modelID/events", func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c, 5*time.Second)
		defer cancel()

		au, err := a.authenticator.ValidateUserContext(c)
		if err != nil {
			api.Error(c, http.StatusUnauthorized, err)
			return
		}

		modelID := c.Param("modelID")
		page, size := api.Page(c)

		events, err := a.eventService.List(ctx, modelID, au.OrganizationID, page, size)

		if err != nil {
			api.Error(c, http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, events)
	})

	a.router.GET("/api/v1/models/:modelID/events/:eventID", func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c, 5*time.Second)
		defer cancel()

		au, err := a.authenticator.ValidateUserContext(c)
		if err != nil {
			api.Error(c, http.StatusUnauthorized, err)
			return
		}

		modelID := c.Param("modelID")
		eventID := c.Param("eventID")

		event, err := a.eventService.Get(ctx, eventID, modelID, au.OrganizationID)

		if err != nil {
			api.Error(c, http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, event)
	})

	a.router.PUT("/api/v1/models/:modelID/events/:eventID", func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c, 5*time.Second)
		defer cancel()

		au, err := a.authenticator.ValidateUserContext(c)
		if err != nil {
			api.Error(c, http.StatusUnauthorized, err)
			return
		}

		modelID := c.Param("modelID")
		eventID := c.Param("eventID")

		var req request.EventUpdateRequest

		if err := c.ShouldBindJSON(&req); err != nil {
			api.Error(c, http.StatusBadRequest, err)
			return
		}

		event, err := a.eventService.Update(ctx, eventID, &req, modelID, au.OrganizationID)

		if err != nil {
			api.Error(c, http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, event)
	})

	a.router.DELETE("/api/v1/models/:modelID/events/:eventID", func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c, 5*time.Second)
		defer cancel()

		au, err := a.authenticator.ValidateUserContext(c)
		if err != nil {
			api.Error(c, http.StatusUnauthorized, err)
			return
		}

		modelID := c.Param("modelID")
		eventID := c.Param("eventID")

		event, err := a.eventService.Delete(ctx, eventID, modelID, au.OrganizationID)

		if err != nil {
			api.Error(c, http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, event)
	})
}
