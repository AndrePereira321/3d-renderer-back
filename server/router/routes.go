package router

import "server/server/router/routes"

var get_routes = map[string]func(*routes.Route){
	"/test": routes.Test,
}

var post_routes = map[string]func(*routes.Route){
	"/register": routes.Register,
	"/login":    routes.Login,
}

var auth_routes = map[string]bool{}

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
	_, ok := auth_routes[path]
	return ok
}
