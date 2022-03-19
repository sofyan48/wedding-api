package cmd

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"

	"github.com/orn-id/wedding-api/src/cmd/app"
	"github.com/orn-id/wedding-api/src/pkg/logger"
)

// Start handler registering service command
func Start() {

	rootCmd := &cobra.Command{}
	logger.SetJSONFormatter()
	ctx, cancel := context.WithCancel(context.Background())

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-quit
		cancel()
	}()

	cmd := []*cobra.Command{
		{
			Use:   "serve",
			Short: "Run HTTP Server",
			Run: func(cmd *cobra.Command, args []string) {
				app.Start(ctx)
			},
		},
		{
			Use:   "scan",
			Short: "Login whatsapp web client",
			Run: func(cmd *cobra.Command, args []string) {
				app.LoginClient()
			},
		},
	}

	rootCmd.AddCommand(cmd...)
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
