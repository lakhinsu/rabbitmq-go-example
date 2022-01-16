package utils

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// Modify the init function if you change logging library
// Find and replace the imports in the rest of the app
func init() {
	// Read log level
	logLevel := GetEnvVar("LOG_LEVEL")

	// Replace with switch or add more levels as per need
	if logLevel == "debug" {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	} else {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}

	// Set time format
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	// Set Global fields
	host, err := os.Hostname()
	if err != nil {
		log.Logger = log.With().Str("host", "unknown").Logger()
	} else {
		log.Logger = log.With().Str("host", host).Logger()
	}

	log.Logger = log.With().Str("service", "rabbitmq-producer").Logger()

	log.Logger = log.With().Caller().Logger()
}
