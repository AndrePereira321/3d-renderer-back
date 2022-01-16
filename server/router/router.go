package router

import (
	"net/http"
	"server/logger"
	"server/server/router/routes"
)

type Router struct {
	http.Handler
}

func NewRouter() *Router {
	return &Router{}
}

func (router Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	route := routes.NewRoute(w, r)
	var routeHandler func(*routes.Route) = nil

	switch route.Method {
	case "OPTIONS":
		logger.Debug("!!!TODO!!! OPTIONS REQUEST")
	case "GET":
		routeHandler = GetGETRoute(route.Path)
	case "POST":
		routeHandler = GetPOSTRoute(route.Path)
	}

	if routeHandler != nil {
		routeHandler(route)
	} else {
		logger.LogWarning("Route not found: "+route.Path, "router.ServeHTTP", "TODO")
	}
}
