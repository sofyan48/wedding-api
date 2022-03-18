// Package middleware
package middleware

import (
	"net/http"

	"github.com/orn-id/wedding-api/src/consts"
)

// MiddlewareFunc is contract for middleware and must implement this type for http if need middleware http request
type MiddlewareFunc func(r *http.Request) string

// FilterFunc is a iterator resolver in each middleware registered
func FilterFunc(r *http.Request, mfs []MiddlewareFunc) string {
	for _, mf := range mfs {
		if status := mf(r); status != consts.MiddlewarePassed {
			return status
		}
	}

	return consts.MiddlewarePassed
}
