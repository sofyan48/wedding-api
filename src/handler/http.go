// Package handler
package handler

import (
	"net/http"

	"github.com/orn-id/wedding-api/src/app/appctx"
	"github.com/orn-id/wedding-api/src/consts"
	"github.com/orn-id/wedding-api/src/app/ucase/contract"
)

// HttpRequest handler func wrapper
func HttpRequest(request *http.Request, svc contract.UseCase) appctx.Response {
	data := &appctx.Data{
		Request:     request,
		ServiceType: consts.ServiceTypeHTTP,
	}

	return svc.Serve(data)
}
