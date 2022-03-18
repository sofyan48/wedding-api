// Package bootstrap
package bootstrap

import (
	"os"
	"strconv"

	"github.com/orn-id/wedding-api/src/pkg/logger"
	"github.com/orn-id/wedding-api/src/pkg/util"
)

func RegistryLogger() {
	debug, _ := strconv.ParseBool(os.Getenv("APP_DEBUG"))
	lc := logger.Config{
		URL:         os.Getenv("LOGGER_URL"),
		Environment: util.EnvironmentTransform(os.Getenv("APP_ENVIRONMENT")),
		Debug:       debug,
		Level:       os.Getenv("LOGGER_LEVEL"),
	}

	h, e := logger.NewSentryHook(lc)

	if e != nil {
		logger.Fatal("log sentry failed to initialize Sentry")
	}

	logger.Setup(lc, h)
}
