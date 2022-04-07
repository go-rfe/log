package log

import (
	"net/http"
	"strings"

	"github.com/go-chi/httplog"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func Level(level string) {
	switch level {
	case "DEBUG":
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	case "INFO":
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	case "WARNING":
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
	case "ERROR":
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	default:
		log.Fatal().Msgf("Unsupported log Level: %s", level)
	}
}

func HTTPRequestLogger(level string) func(next http.Handler) http.Handler {
	logger := httplog.NewLogger("metrics", httplog.Options{
		JSON:     false,
		LogLevel: strings.ToLower(level),
	})

	return httplog.RequestLogger(logger)
}
