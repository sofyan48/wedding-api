// Package router
package server

import (
	"context"
	"net/http"

	"github.com/orn-id/wedding-api/src/app/appctx"
	ucase "github.com/orn-id/wedding-api/src/app/ucase/contract"
)

// httpHandlerFunc abstraction for http handler
type httpHandlerFunc func(request *http.Request, svc ucase.UseCase) appctx.Response

// Server contract
type Server interface {
	Run(ctx context.Context, port string) error
	Done()
}
