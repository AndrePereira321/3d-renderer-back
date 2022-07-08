package router

import (
	"errors"
	"net/http"
	"os"
	"server/database/repositories"
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

func checkActiveSession(r *routes.Route) bool {
	sessionCode := r.Request.Header.Get(globals.SESSION_CODE_HEADER)
	if len(sessionCode) < 8 {
		r.ResponseWriter.WriteError(http.StatusForbidden, globals.ERROR_INVALID_SESSION)
		return false
	}
	active, err := repositories.IsActiveSession(sessionCode)
	if err != nil {
		go logger.LogError("Error retrieving user session "+sessionCode+" : "+err.Error(), "Router.checkActiveSession", globals.ERROR_DATABASE_ERROR)
		r.ResponseWriter.WriteErrorMessage(http.StatusInternalServerError, globals.ERROR_DATABASE_ERROR, "Error retrieving user session: "+err.Error())
		return false
	}
	if !active {
		r.ResponseWriter.WriteError(http.StatusForbidden, globals.ERROR_INVALID_SESSION)
		return false
	}
	return true
}

func (router Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", globals.CLIENT_URL)
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization, "+globals.SESSION_CODE_HEADER)

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
		if IsAuthRoute(route.Path) && !checkActiveSession(route) {
			return
		}
		routeHandler(route)
	} else {
		path := globals.CLIENT_FOLDER + route.Path
		if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
			go logger.LogWarning("Route not found: "+route.Path, "router.ServeHTTP", globals.ERROR_ROUTE_NOT_FOUND)
			route.ResponseWriter.WriteError(http.StatusNotFound, globals.ERROR_ROUTE_NOT_FOUND)
			return
		}
		http.ServeFile(w, r, path)
	}
}
