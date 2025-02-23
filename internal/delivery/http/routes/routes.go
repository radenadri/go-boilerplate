package routes

import (
	"fmt"
	"net/http"
	"time"

	"github.com/getsentry/sentry-go"
	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-gonic/gin"
	"github.com/radenadri/go-boilerplate/config"
	_ "github.com/radenadri/go-boilerplate/docs"
	"github.com/radenadri/go-boilerplate/internal/delivery/http/controllers"
	"github.com/radenadri/go-boilerplate/internal/delivery/http/middlewares"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"
)

// @title           Go Boilerplate API
// @version         1.0
// @description     A robust and scalable Go boilerplate for building modern web applications.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  MIT
// @license.url   https://opensource.org/licenses/MIT

// @host      localhost:6500
// @BasePath  /api/v1

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func InitRouter() *gin.Engine {

	if err := sentry.Init(sentry.ClientOptions{
		Dsn:              config.SentryDSN,
		EnableTracing:    true,
		TracesSampleRate: 1.0,
		BeforeSend: func(event *sentry.Event, hint *sentry.EventHint) *sentry.Event {
			if hint.Context != nil {
				if req, ok := hint.Context.Value(sentry.RequestContextKey).(*http.Request); ok {
					config.Logger.Info("Sending error to Sentry",
						zap.String("method", req.Method),
						zap.String("path", req.URL.Path),
						zap.String("ip", req.RemoteAddr),
					)
				}
			}

			return event
		},
	}); err != nil {
		fmt.Printf("Sentry initialization failed: %v\n", err)
	}

	r := gin.Default()

	// Initialize Sentry's handler
	r.Use(sentrygin.New(sentrygin.Options{
		Repanic: true,
	}))

	// Override default error handlers
	r.NoRoute(middlewares.NotFoundHandler())
	r.Use(middlewares.ErrorHandler())

	// Swagger documentation endpoint
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Configure rate limiter
	rateLimiterConfig := middlewares.RateLimiterConfig{
		RedisClient: config.RedisClient,
		MaxRequests: 100,         // 100 requests
		Window:      time.Minute, // per minute
	}

	// Apply rate limiter to all routes
	r.Use(middlewares.RateLimiter(rateLimiterConfig))
	// Apply CORS middleware with default configuration
	r.Use(middlewares.CORS(middlewares.DefaultCORSConfig()))

	api := r.Group(fmt.Sprintf("/api/%s", config.AppApiVersion))

	// Initialize controllers
	userController := controllers.NewUserController()

	// Public routes
	public := api.Group("")
	{
		public.POST("/login", userController.Login)
		public.POST("/register", userController.Register)

		// Test Sentry
		public.GET("/foo", func(ctx *gin.Context) {
			panic("y tho")
		})
	}

	// Protected routes
	protected := api.Group("")
	protected.Use(middlewares.AuthenticateJWT())
	{
		protected.GET("/users", userController.GetAllUsers)
	}

	return r
}
