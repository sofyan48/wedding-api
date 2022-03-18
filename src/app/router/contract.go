// Package router
package router

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/orn-id/wedding-api/src/app/appctx"
	"github.com/orn-id/wedding-api/src/app/ucase/contract"
)

// httpHandlerFunc is a contract http handler for router
type httpHandlerFunc func(request *http.Request, svc contract.UseCase) appctx.Response

// Router is a contract router and must implement this interface
type Router interface {
	Route() *mux.Router
}
