package logging

import (
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func init() {
	// Set time format
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	// Set global log level
	zerolog.SetGlobalLevel(zerolog.DebugLevel)

	// Configure LogStyle. Defaults to JSON logs, unless LOGSTYLE is set to 'console'
	logStyle := os.Getenv("LOGSTYLE")
	if logStyle == "console" {
		configurePrettyConsoleLogger()
	} else {
		configureJSONLogger()
	}
}

func configurePrettyConsoleLogger() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339})
}

func configureJSONLogger() {
	log.Logger = zerolog.New(os.Stderr).With().Timestamp().Logger()
}
