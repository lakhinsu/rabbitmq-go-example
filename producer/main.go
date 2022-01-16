package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/lakhinsu/rabbitmq-go-example/producer/app"
	"github.com/lakhinsu/rabbitmq-go-example/producer/utils"
	"github.com/rs/zerolog/log"
)

func init() {
	// Set gin mode
	mode := utils.GetEnvVar("GIN_MODE")
	gin.SetMode(mode)
}

func main() {
	// Setup the app
	app := app.SetupApp()

	// Read ADDR and port
	addr := utils.GetEnvVar("GIN_ADDR")
	port := utils.GetEnvVar("GIN_PORT")

	https := utils.GetEnvVar("GIN_HTTPS")
	// HTTPS mode
	if https == "true" {
		certFile := utils.GetEnvVar("GIN_CERT")
		certKey := utils.GetEnvVar("GIN_CERT_KEY")
		log.Info().Msgf("Starting service on https//:%s:%s", addr, port)

		if err := app.RunTLS(fmt.Sprintf("%s:%s", addr, port), certFile, certKey); err != nil {
			log.Fatal().Err(err).Msg("Error occurred while setting up the server in HTTPS mode")
		}
	}
	// HTTP mode
	log.Info().Msgf("Starting service on http//:%s:%s", addr, port)
	if err := app.Run(fmt.Sprintf("%s:%s", addr, port)); err != nil {
		log.Fatal().Err(err).Msg("Error occurred while setting up the server")
	}
}
