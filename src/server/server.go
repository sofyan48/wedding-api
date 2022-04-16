package server

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/orn-id/wedding-api/src/app/router"
	"github.com/orn-id/wedding-api/src/pkg/logger"
)

// NewHTTPServer creates http server instance
// returns: Server instance
func NewHTTPServer() Server {
	return &httpServer{
		router: router.NewRouter(),
	}
}

// httpServer as HTTP server implementation
type httpServer struct {
	router router.Router
}

// Run runs the http server gracefully
// returns:
//	err: error operation
func (h *httpServer) Run(ctx context.Context, port string) error {
	var err error

	server := http.Server{
		Addr:         fmt.Sprintf("0.0.0.0:%s", port),
		Handler:      h.router.Route(),
		ReadTimeout:  time.Duration(5) * time.Second,
		WriteTimeout: time.Duration(15) * time.Second,
	}

	go func() {
		err = server.ListenAndServe()
		if err != http.ErrServerClosed {
			logger.Error(logger.SetMessageFormat("http server got %#v", err))
		}
	}()

	<-ctx.Done()

	ctxShutDown, cancel := context.WithTimeout(context.Background(), 6*time.Second)
	defer func() {
		cancel()
	}()

	if err = server.Shutdown(ctxShutDown); err != nil {
		logger.Fatal(logger.SetMessageFormat("server Shutdown Failed:%+s", err))
	}

	logger.Info("server exited properly")

	if err == http.ErrServerClosed {
		err = nil
	}

	return err
}

// Done runs event wheen service stopped
func (h *httpServer) Done() {
	logger.Info("service http stopped")
}
