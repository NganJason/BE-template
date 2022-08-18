package main

import (
	"context"
	"net/http"

	"github.com/NganJason/BE-template/internal/handler"
	"github.com/NganJason/BE-template/internal/middleware"
	"github.com/NganJason/BE-template/internal/vo"
	"github.com/NganJason/BE-template/pkg/clog"
	"github.com/NganJason/BE-template/pkg/server"
)

func main() {
	router := server.NewMainRouter()

	for _, middleware := range allMiddlewares() {
		router.AddMiddleware(middleware)
	}

	for _, route := range allRoutes() {
		router.AddRoute(route)
	}

	server := &http.Server{
		Addr:    ":" + "8082",
		Handler: router,
	}

	clog.Info(context.Background(), "serving at 8082")

	server.ListenAndServe()
}

func allMiddlewares() []server.Middleware {
	middlewares := []server.Middleware{
		&middleware.ValidateBodyMiddleware{},
	}

	return middlewares
}

func allRoutes() []*server.Route {
	routes := []*server.Route{
		{
			Name:    vo.CmdHealthCheck,
			Method:  http.MethodGet,
			Path:    vo.PathHealthCheck,
			Handler: handler.HealthCheck,
			Req:     &vo.HealthCheckRequest{},
		},
	}

	return routes
}
