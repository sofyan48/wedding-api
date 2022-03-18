package router

import (
	"context"
	"encoding/json"
	"net/http"
	"os"
	"runtime/debug"
	"time"

	"github.com/gorilla/mux"

	"github.com/orn-id/wedding-api/src/app/appctx"
	"github.com/orn-id/wedding-api/src/bootstrap"
	"github.com/orn-id/wedding-api/src/consts"
	"github.com/orn-id/wedding-api/src/middleware"
	"github.com/orn-id/wedding-api/src/pkg/logger"

	ucaseContract "github.com/orn-id/wedding-api/src/app/ucase/contract"
)

type router struct {
	router *mux.Router
}

// NewRouter initialize new router wil return Router Interface
func NewRouter() Router {
	bootstrap.RegistryMessage()
	bootstrap.RegistryLogger()

	return &router{
		router: mux.NewRouter(),
	}
}

func (rtr *router) handle(hfn httpHandlerFunc, svc ucaseContract.UseCase, mdws ...middleware.MiddlewareFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			err := recover()
			if err != nil {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusInternalServerError)
				res := appctx.Response{
					Name: consts.InternalFailure,
				}

				res.SetMessage()
				logger.Error(logger.SetMessageFormat("error %v", string(debug.Stack())))
				json.NewEncoder(w).Encode(res)
				return
			}
		}()
		
		w.Header().Set("Access-Control-Allow-Origin", os.Getenv("CORS_CONTROL_ALLOW_ORIGIN"))

		var st time.Time
		var lt time.Duration

		st = time.Now()

		ctx := context.WithValue(r.Context(), "access", map[string]interface{}{
			"path":      r.URL.Path,
			"remote_ip": r.RemoteAddr,
			"method":    r.Method,
		})

		req := r.WithContext(ctx)

		if status := middleware.FilterFunc(req, mdws); status != consts.MiddlewarePassed {
			rtr.response(w, appctx.Response{
				Name: status,
			})

			return
		}

		resp := hfn(req, svc)

		resp.Lang = rtr.defaultLang(req.Header.Get("X-Lang"))

		rtr.response(w, resp)

		lt = time.Since(st)
		logger.AccessLog("request",
			logger.Any("tag", "http-request"),
			logger.Any("http.path", req.URL.Path),
			logger.Any("http.method", req.Method),
			logger.Any("http.agent", req.UserAgent()),
			logger.Any("http.referer", req.Referer()),
			logger.Any("http.status", resp.GetCode()),
			logger.Any("http.latency", lt.Seconds()),
		)
	}
}

// func (rtr *router) handleWithCache(storage cache.Cacher, duration string, hfn httpHandlerFunc, svc ucaseContract.UseCase, mdws ...middleware.MiddlewareFunc) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		defer func() {
// 			err := recover()
// 			if err != nil {
// 				w.Header().Set("Content-Type", "application/json")
// 				w.WriteHeader(http.StatusInternalServerError)
// 				res := appctx.Response{
// 					Name: consts.InternalFailure,
// 				}

// 				res.SetMessage()
// 				logger.Error(logger.SetMessageFormat("error %v", string(debug.Stack())))
// 				json.NewEncoder(w).Encode(res)
// 				return
// 			}
// 		}()
// 		var lt time.Duration
// 		var st time.Time
// 		ctx := context.WithValue(r.Context(), "access", map[string]interface{}{
// 			"path":      r.URL.Path,
// 			"remote_ip": r.RemoteAddr,
// 			"method":    r.Method,
// 		})
// 		req := r.WithContext(ctx)

// 		if status := middleware.FilterFunc(req, mdws); status != consts.MiddlewarePassed {
// 			rtr.response(w, appctx.Response{
// 				Name: status,
// 			})
// 			return
// 		}

// 		st = time.Now()
// 		lt = time.Since(st)

// 		var resp appctx.Response
// 		content, _ := storage.Get(r.RequestURI)
// 		if len(content) == 0 {
// 			resp = hfn(req, svc)
// 			insertCache, _ := json.Marshal(resp)
// 			durationTime, _ := time.ParseDuration(duration)
// 			storage.Set(r.RequestURI, insertCache, durationTime)
// 		} else {
// 			json.Unmarshal(content, &resp)
// 		}

// 		resp.Lang = rtr.defaultLang(req.Header.Get("X-Lang"))
// 		rtr.response(w, resp)

// 		logger.AccessLog("request",
// 			logger.Any("tag", "http-request"),
// 			logger.Any("http.path", req.URL.Path),
// 			logger.Any("http.method", req.Method),
// 			logger.Any("http.agent", req.UserAgent()),
// 			logger.Any("http.referer", req.Referer()),
// 			logger.Any("http.status", resp.GetCode()),
// 			logger.Any("http.latency", lt.Seconds()),
// 		)

// 	}
// }

// response prints as a json and formatted string for DGP legacy
func (rtr *router) response(w http.ResponseWriter, resp appctx.Response) {

	w.Header().Set("Content-Type", "application/json")

	defer func() {
		resp.SetMessage()
		w.WriteHeader(resp.GetCode())
		json.NewEncoder(w).Encode(resp)
	}()

	return

}

func (rtr *router) defaultLang(l string) string {

	if len(l) == 0 {
		return os.Getenv("APP_DEFAULT_LANG")
	}

	return l
}
