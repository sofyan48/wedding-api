package middleware

import (
	"net/http"
	"os"

	"github.com/orn-id/wedding-api/src/consts"
)

// BasicAuth ...
func BasicAuth(r *http.Request) string {
	// check if dev environtment not check basic auth
	if os.Getenv("APP_ENVIRONMENT") == "dev" {
		return consts.MiddlewarePassed
	}
	// if not dev check basic auth
	username, password, ok := r.BasicAuth()
	if !ok {
		return consts.AuthenticationFailure
	}

	isValid := (username == os.Getenv("APP_USERNAME")) && (password == os.Getenv("APP_PASSWORD"))
	if !isValid {
		return consts.AuthenticationFailure
	}

	return consts.MiddlewarePassed
}
