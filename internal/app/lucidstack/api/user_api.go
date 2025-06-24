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

type UserApi struct {
	router        *gin.Engine
	authenticator *auth.Authenticator
	userService   *service.UserService
}

func NewUserApi(router *gin.Engine, authenticator *auth.Authenticator, userService *service.UserService) *UserApi {
	return &UserApi{router: router, authenticator: authenticator, userService: userService}
}

func (a *UserApi) Register() {

	a.router.POST("/api/v1/users/signup", func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c, 5*time.Second)
		defer cancel()

		var req request.UserSignUpRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			api.Error(c, http.StatusBadRequest, err)
			return
		}

		user, err := a.userService.SignUp(ctx, &req)

		if err != nil {
			api.Error(c, http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusCreated, user)
	})

	a.router.POST("/api/v1/users/token", func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c, 5*time.Second)
		defer cancel()

		var req request.UserTokenRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			api.Error(c, http.StatusBadRequest, err)
			return
		}

		token, err := a.userService.GetToken(ctx, &req)

		if err != nil {
			api.Error(c, http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, token)
	})

	a.router.GET("/api/v1/users/me", func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c, 5*time.Second)
		defer cancel()

		au, err := a.authenticator.ValidateUserContext(c)
		if err != nil {
			api.Error(c, http.StatusUnauthorized, err)
			return
		}

		user, err := a.userService.Get(ctx, au.ID, au.OrganizationID)

		if err != nil {
			api.Error(c, http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, user)
	})

	a.router.PUT("/api/v1/users/me/password", func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c, 5*time.Second)
		defer cancel()

		au, err := a.authenticator.ValidateUserContext(c)
		if err != nil {
			api.Error(c, http.StatusUnauthorized, err)
			return
		}

		var req request.UserPasswordChangeRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			api.Error(c, http.StatusBadRequest, err)
			return
		}

		user, err := a.userService.ChangePassword(ctx, au.ID, &req, au.OrganizationID)

		if err != nil {
			api.Error(c, http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, user)
	})

	a.router.POST("/api/v1/users", func(c *gin.Context) {
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

		var req request.UserAdditionRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			api.Error(c, http.StatusBadRequest, err)
			return
		}

		user, err := a.userService.Add(ctx, &req, au.ID, au.OrganizationID)

		if err != nil {
			api.Error(c, http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusCreated, user)
	})

	a.router.GET("/api/v1/users", func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c, 5*time.Second)
		defer cancel()

		au, err := a.authenticator.ValidateUserContext(c)
		if err != nil {
			api.Error(c, http.StatusUnauthorized, err)
			return
		}

		page, size := api.Page(c)

		users, err := a.userService.List(ctx, au.OrganizationID, page, size)

		if err != nil {
			api.Error(c, http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, users)
	})

	a.router.GET("/api/v1/users/:userID", func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c, 5*time.Second)
		defer cancel()

		au, err := a.authenticator.ValidateUserContext(c)
		if err != nil {
			api.Error(c, http.StatusUnauthorized, err)
			return
		}

		userID := c.Param("userID")

		user, err := a.userService.Get(ctx, userID, au.OrganizationID)

		if err != nil {
			api.Error(c, http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, user)
	})

	a.router.PUT("/api/v1/users/:userID/password", func(c *gin.Context) {
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

		userID := c.Param("userID")

		user, err := a.userService.ResetPassword(ctx, userID, au.OrganizationID)

		if err != nil {
			api.Error(c, http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, user)
	})

	a.router.PUT("/api/v1/users/:userID/admin", func(c *gin.Context) {
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

		userID := c.Param("userID")
		var req request.UserAdminUpdateRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			api.Error(c, http.StatusBadRequest, err)
			return
		}

		user, err := a.userService.UpdateAdmin(ctx, userID, &req, au.OrganizationID)

		if err != nil {
			api.Error(c, http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, user)
	})

	a.router.DELETE("/api/v1/users/:userID", func(c *gin.Context) {
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

		userID := c.Param("userID")
		user, err := a.userService.Delete(ctx, userID, au.OrganizationID)
		if err != nil {
			api.Error(c, http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, user)
	})
}
