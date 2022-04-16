package app

import (
	"context"
	"os"

	"github.com/orn-id/wedding-api/src/pkg/logger"
	"github.com/orn-id/wedding-api/src/server"
	"github.com/spf13/cobra"
)

func Serve(ctx context.Context) *cobra.Command {
	serve := &cobra.Command{}
	serve.Flags().StringP("port", "", os.Getenv("SERVER_PORT"), "Add port")
	serve.Use = "serve"
	serve.Short = "Run HttpServer"
	serve.Run = func(cmd *cobra.Command, args []string) {
		port, _ := cmd.Flags().GetString("port")
		serve := server.NewHTTPServer()
		defer serve.Done()
		logger.Info(logger.SetMessageFormat("starting services... " + os.Getenv("SERVER_PORT")))
		if err := serve.Run(ctx, port); err != nil {
			logger.Warn(logger.SetMessageFormat("service stopped, err:%s", err.Error()))
		}
	}
	return serve
}
