package app

import (
	"context"
	"os"

	"github.com/orn-id/wedding-api/src/pkg/logger"
	"github.com/orn-id/wedding-api/src/server"
)

// Start function handler starting http listener from app
func Start(ctx context.Context) {
	serve := server.NewHTTPServer()
	defer serve.Done()
	logger.Info(logger.SetMessageFormat("starting services... " + os.Getenv("SERVER_PORT")))

	if err := serve.Run(ctx); err != nil {
		logger.Warn(logger.SetMessageFormat("service stopped, err:%s", err.Error()))

	}
	return
}
