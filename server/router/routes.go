package router

import (
	"server/server/router/routes"
	"server/server/router/routes/grid"
)

var get_routes = map[string]func(*routes.Route){
	"/references": routes.Reference,
	"/grid":       grid.Grid,
	"/test":       routes.Test,
}

var post_routes = map[string]func(*routes.Route){
	"/register":    routes.Register,
	"/login":       routes.Login,
	"/isConnected": routes.IsConnected,
	"/disconnect":  routes.Disconnect,
}

var auth_routes = map[string]bool{
	"/disconnect": true,
	"/grid":       true,
}

func GetGETRoute(path string) func(*routes.Route) {
	route, ok := get_routes[path]
	if !ok {
		return nil
	}
	return route
}

func GetPOSTRoute(path string) func(*routes.Route) {
	route, ok := post_routes[path]
	if !ok {
		return nil
	}
	return route
}

func IsAuthRoute(path string) bool {
	result, ok := auth_routes[path]
	if !ok {
		return false
	}
	return result
}
