package router

import (
	"net/http"
	"server/globals"
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
	w.Header().Set("Access-Control-Allow-Origin", globals.CLIENT_URL)
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization")

	route := routes.NewRoute(w, r)
	var routeHandler func(*routes.Route) = nil

	switch route.Method {
	case "OPTIONS":
		w.WriteHeader(http.StatusOK)
		return
	case "GET":
		routeHandler = GetGETRoute(route.Path)
		break
	case "POST":
		routeHandler = GetPOSTRoute(route.Path)
		break
	}
	if routeHandler != nil {
		routeHandler(route)
	} else {
		go logger.LogWarning("Route not found: "+route.Path, "router.ServeHTTP", globals.ERROR_ROUTE_NOT_FOUND)
		route.ResponseWriter.WriteError(http.StatusNotFound, globals.ERROR_ROUTE_NOT_FOUND)
	}
}
