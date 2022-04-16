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

	http := &cobra.Command{
		Use:   "http",
		Short: "Run http server",
	}
	http.AddCommand(app.Serve(ctx))

	wa := &cobra.Command{
		Use:   "wa",
		Short: "WA Client",
	}

	wa.AddCommand(app.Scan())

	cmd := []*cobra.Command{
		http,
		wa,
	}

	rootCmd.AddCommand(cmd...)
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
