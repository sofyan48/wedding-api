package router

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/orn-id/wedding-api/src/app/ucase"
	"github.com/orn-id/wedding-api/src/app/ucase/invitation"
	"github.com/orn-id/wedding-api/src/handler"
	_ "github.com/orn-id/wedding-api/src/swagger"
	"github.com/orn-id/wedding-api/src/validator"
	httpSwagger "github.com/swaggo/http-swagger"
)

// Route preparing http router and will return mux router object
func (rtr *router) Route() *mux.Router {
	root := rtr.router.PathPrefix("/").Subrouter()
	in := root.PathPrefix("/in/").Subrouter()
	// //internal version sub router
	//inV1 := in.PathPrefix("/v1/").Subrouter()

	// swagger routes
	in.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)

	// healthy
	healthy := ucase.NewHealthCheck()
	in.HandleFunc("/health", rtr.handle(
		handler.HttpRequest,
		healthy,
	)).Methods(http.MethodGet)

	// // external subrouter
	ex := root.PathPrefix("/ex/").Subrouter()
	exV1 := ex.PathPrefix("/v1/").Subrouter()

	importsInvitation := invitation.NewImports()
	exV1.HandleFunc("/invitation/import", rtr.handle(
		handler.HttpRequest,
		importsInvitation,
	)).Methods(http.MethodPost)

	sendInvitation := invitation.NewSendInvitation(validator.New())
	exV1.HandleFunc("/invitation", rtr.handle(
		handler.HttpRequest,
		sendInvitation,
	)).Methods(http.MethodPost)

	return rtr.router

}
