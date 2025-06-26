package lucidstack

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/kamva/mgm/v3"
	"github.com/lucidstackhq/lucidstack/internal/app/lucidstack/api"
	"github.com/lucidstackhq/lucidstack/internal/app/lucidstack/service"
	"github.com/lucidstackhq/lucidstack/internal/pkg/auth"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
)

type Server struct {
	config *ServerConfig
}

func NewServer(config *ServerConfig) *Server {
	return &Server{config: config}
}

type ServerConfig struct {
	Host          string
	Port          string
	MongoEndpoint string
	MongoDatabase string
	JwtSigningKey string
}

func (s *Server) Start() {
	err := mgm.SetDefaultConfig(nil, s.config.MongoDatabase, options.Client().ApplyURI(s.config.MongoEndpoint))

	if err != nil {
		log.Fatal("error connecting to mongodb: ", err)
	}

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"*"},
		AllowCredentials: true,
	}))

	authenticator := auth.NewAuthenticator(s.config.JwtSigningKey)
	organizationService := service.NewOrganizationService()
	userService := service.NewUserService(organizationService, authenticator)
	modelService := service.NewModelService()
	propertyService := service.NewPropertyService(modelService)
	eventService := service.NewEventService(modelService)
	actionService := service.NewActionService(modelService)
	appService := service.NewAppService()
	environmentService := service.NewEnvironmentService()
	entityService := service.NewEntityService(modelService, environmentService)
	entityPropertyService := service.NewEntityPropertyService(entityService, propertyService)

	api.NewHealthCheckApi(router).Register()
	api.NewUserApi(router, authenticator, userService).Register()
	api.NewOrganizationApi(router, authenticator, organizationService).Register()
	api.NewModelApi(router, authenticator, modelService).Register()
	api.NewPropertyApi(router, authenticator, propertyService).Register()
	api.NewEventApi(router, authenticator, eventService).Register()
	api.NewActionApi(router, authenticator, actionService).Register()
	api.NewAppApi(router, authenticator, appService).Register()
	api.NewEnvironmentApi(router, authenticator, environmentService).Register()
	api.NewEntityApi(router, authenticator, entityService).Register()
	api.NewEntityPropertyApi(router, authenticator, entityPropertyService).Register()

	router.Delims("[[", "]]")
	router.LoadHTMLGlob("templates/*")
	s.registerPages(router)
	router.Static("/static", "./static")

	log.Println("starting lucidstack server")
	err = router.Run(fmt.Sprintf("%s:%s", s.config.Host, s.config.Port))

	if err != nil {
		log.Fatal("error starting server", err)
	}
}

func (s *Server) registerPages(router *gin.Engine) {
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	router.GET("/logout", func(c *gin.Context) {
		c.HTML(http.StatusOK, "logout.html", gin.H{})
	})

	router.GET("/join", func(c *gin.Context) {
		c.HTML(http.StatusOK, "join.html", gin.H{})
	})

	router.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", gin.H{})
	})

	router.GET("/home", func(c *gin.Context) {
		c.HTML(http.StatusOK, "home.html", gin.H{})
	})

	router.GET("/account", func(c *gin.Context) {
		c.HTML(http.StatusOK, "account.html", gin.H{})
	})

	router.GET("/organization", func(c *gin.Context) {
		c.HTML(http.StatusOK, "organization.html", gin.H{})
	})

	router.GET("/users", func(c *gin.Context) {
		c.HTML(http.StatusOK, "users.html", gin.H{})
	})

	router.GET("/users/:userID", func(c *gin.Context) {
		c.HTML(http.StatusOK, "user.html", gin.H{
			"user_id": c.Param("userID"),
		})
	})

	router.GET("/apps", func(c *gin.Context) {
		c.HTML(http.StatusOK, "apps.html", gin.H{})
	})

	router.GET("/apps/:appID", func(c *gin.Context) {
		c.HTML(http.StatusOK, "app.html", gin.H{
			"app_id": c.Param("appID"),
		})
	})
}
