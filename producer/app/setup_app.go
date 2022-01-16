package app

import (
	"github.com/gin-gonic/gin"
	"github.com/lakhinsu/rabbitmq-go-example/producer/middlewares"
	"github.com/lakhinsu/rabbitmq-go-example/producer/routers"
	"github.com/rs/zerolog/log"
)

// Function to setup the app object
func SetupApp() *gin.Engine {
	log.Info().Msg("Initializing service")

	// Create barebone engine
	app := gin.New()
	// Add default recovery middleware
	app.Use(gin.Recovery())

	// disabling the trusted proxy feature
	app.SetTrustedProxies(nil)

	// Add cors, request ID and request logging middleware
	log.Info().Msg("Adding cors, request id and request logging middleware")
	app.Use(middlewares.CORSMiddleware(), middlewares.RequestID(), middlewares.RequestLogger())

	// Setup routers
	log.Info().Msg("Setting up routers")
	routers.SetupRouters(app)

	return app
}
